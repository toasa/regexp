package nfa

import (
	"fmt"
	"regexp/parser"
)

type State struct {
	ID    int
	Nexts map[rune][]*State
}

// NFA is non-deterministic finite automaton
type NFA struct {
	States       []*State
	StartState   *State
	AcceptStates []*State
}

func newNFA(states []*State, start *State, accepts []*State) *NFA {
	return &NFA{
		States:       states,
		StartState:   start,
		AcceptStates: accepts,
	}
}

func removeDuplicate(states []*State) []*State {
	dup_map := make(map[int]bool)
	newStates := make([]*State, 0)

	for _, state := range states {
		_, ok := dup_map[state.ID]
		if !ok {
			dup_map[state.ID] = true
			newStates = append(newStates, state)
		}
	}

	return newStates
}

func contain(state *State, states []*State) bool {
	for _, s := range states {
		if s.ID == state.ID {
			return true
		}
	}
	return false
}

func adaptEpsilonTransition(states []*State) []*State {
	nextStates := []*State{}
	for _, state := range states {
		if nexts, ok := state.Nexts['ε']; ok {
			nextStates = append(nextStates, nexts...)
		} else {
			nextStates = append(nextStates, state)
		}
	}
	return removeDuplicate(nextStates)
}

func containStateOfEpsilonTransitive(states []*State) bool {
	for _, state := range states {
		if _, ok := state.Nexts['ε']; ok {
			return true
		}
	}
	return false
}

// check that nfa accepts the string or not.
func (nfa *NFA) accept(str string) bool {
	curStates := []*State{nfa.StartState}
	for containStateOfEpsilonTransitive(curStates) {
		curStates = adaptEpsilonTransition(curStates)
	}

	for _, c := range str {
		nextStates := []*State{}
		for _, state := range curStates {
			next, ok := state.Nexts[c]
			if ok {
				nextStates = append(nextStates, next...)
			}
		}
		nextStates = removeDuplicate(nextStates)
		// adapt ε transition *after* each symbol is read.
		curStates = adaptEpsilonTransition(nextStates)
	}

	for _, state := range curStates {
		if contain(state, nfa.AcceptStates) {
			return true
		}
	}
	return false
}

// DumpDOT outputs a DOT. DOT is a graph description language.
// The start state forms square box and the accept states form double circle.
func (nfa *NFA) DumpDOT() {
	fmt.Printf("digraph G {\n")
	fmt.Printf("    %d [shape = box];\n", nfa.StartState.ID)
	for _, s := range nfa.AcceptStates {
		fmt.Printf("    %d [shape = doublecircle];\n", s.ID)
	}

	for _, src := range nfa.States {
		for symbol, dstStates := range src.Nexts {
			for _, dst := range dstStates {
				fmt.Printf("    %d -> %d [label=%c];\n", src.ID, dst.ID, symbol)
			}
		}
	}
	fmt.Print("}\n")
}

// NFA generator
type Generator struct {
	StateCount int
}

func newGenerator() *Generator {
	return &Generator{
		StateCount: 0,
	}
}

func (g *Generator) newState() *State {
	id := g.StateCount
	g.StateCount++
	return &State{
		ID:    id,
		Nexts: make(map[rune][]*State),
	}
}

func (g *Generator) genSymbolNFA(symbol rune) *NFA {
	src := g.newState()
	dst := g.newState()

	tmp := src.Nexts[symbol]
	tmp = append(tmp, dst)
	src.Nexts[symbol] = tmp

	states := []*State{src, dst}
	accepts := []*State{dst}

	return newNFA(states, src, accepts)
}

func (g *Generator) genUnionNFA(lhs, rhs *NFA) *NFA {
	start := g.newState()
	start.Nexts['ε'] = []*State{lhs.StartState, rhs.StartState}

	states := []*State{start}
	states = append(states, lhs.States...)
	states = append(states, rhs.States...)
	accepts := append(lhs.AcceptStates, rhs.AcceptStates...)
	return newNFA(states, start, accepts)
}

func (g *Generator) genConcateNFA(lhs, rhs *NFA) *NFA {
	for _, state := range lhs.AcceptStates {
		tmp := state.Nexts['ε']
		tmp = append(tmp, rhs.StartState)
		state.Nexts['ε'] = tmp
	}

	start := lhs.StartState
	accepts := rhs.AcceptStates
	states := append(lhs.States, rhs.States...)
	return newNFA(states, start, accepts)
}

func (g *Generator) gen(node *parser.Node) *NFA {
	switch node.Type {
	case parser.ND_SYMBOL:
		return g.genSymbolNFA(node.Value)
	case parser.ND_UNION:
		lhs := g.gen(node.Lhs)
		rhs := g.gen(node.Rhs)
		return g.genUnionNFA(lhs, rhs)
	case parser.ND_CONCAT:
		lhs := g.gen(node.Lhs)
		rhs := g.gen(node.Rhs)
		return g.genConcateNFA(lhs, rhs)
	}
	return nil
}

func CreateNFA(node *parser.Node) *NFA {
	generator := newGenerator()
	return generator.gen(node)
}
