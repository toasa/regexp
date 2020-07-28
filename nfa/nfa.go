package nfa

import (
	"fmt"
	"regexp/parser"
)

const ε byte = '\000'

type State struct {
	ID    int
	Nexts map[byte][]State
}

func newState(id int) State {
	return State{
		ID:    id,
		Nexts: make(map[byte][]State),
	}
}

// NFA is non-deterministic finite automaton
type NFA struct {
	States       []State
	StartState   State
	AcceptStates []State
}

func newNFA(states []State, start State, accepts []State) NFA {
	return NFA{
		States:       states,
		StartState:   start,
		AcceptStates: accepts,
	}
}

func removeDuplicate(states []State) []State {
	dup_map := make(map[int]bool)
	newStates := make([]State, 0)

	for _, state := range states {
		_, ok := dup_map[state.ID]
		if !ok {
			dup_map[state.ID] = true
			newStates = append(newStates, state)
		}
	}

	return newStates
}

func contain(state State, states []State) bool {
	for _, s := range states {
		if s.ID == state.ID {
			return true
		}
	}
	return false
}

func (nfa *NFA) accept(input string) bool {
	curStates := []State{nfa.StartState}
	nextStates := []State{}
	for i := 0; i < len(input); i++ {
		c := input[i]
		for _, state := range curStates {
			next, ok := state.Nexts[c]
			if ok {
				nextStates = append(nextStates, next...)
			}
		}
		curStates = removeDuplicate(nextStates)
	}

	for _, state := range curStates {
		if contain(state, nfa.AcceptStates) {
			return true
		}
	}
	return false
}

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

func (g *Generator) newStateID() int {
	id := g.StateCount
	g.StateCount++
	return id
}

func (g *Generator) genSymbolNFA(symbol byte) NFA {
	srcID := g.newStateID()
	dstID := g.newStateID()

	src := newState(srcID)
	dst := newState(dstID)

	tmp := src.Nexts[symbol]
	tmp = append(tmp, dst)
	src.Nexts[symbol] = tmp

	states := []State{src, dst}
	accepts := []State{dst}
	return newNFA(states, src, accepts)
}

func CreateNFA(node *parser.Node) NFA {
	generator := newGenerator()

	var nfa NFA
	if node.Type == parser.ND_SYMBOL {
		nfa = generator.genSymbolNFA(node.Value)
	}
	return nfa
}
