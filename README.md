# regexp

演算子 `*`, `|` とカッコ `()` のみが使える正規表現エンジンです。

次の3機能に対応しています。

- パースした構文木の出力
- NFAの状態遷移図の出力
- 正規表現の受理判定

# DEMO

## 受理判定

```bash
$ go run main.go -regexp "(a|bc)*" -input "aabcabcbca"
(a|bc)* accepts aabcabcbca.
```

## 状態遷移図

`(a|bc)*`

![state diagram](state_diagram.png)

# Requirement

* Golang
* Graphviz処理ツール

# Usage

## パースした構文木の出力

```bash
$ go run main.go -regexp <Regular expression> -ast
```

## NFAの状態遷移図の出力

```bash
$ go run main.go -regexp <Regular expression> -state-diagram
```

## 正規表現の受理判定

```bash
$ go run main.go -regexp <Regular expression> -input <Input string>
```

# License

[MIT license](https://en.wikipedia.org/wiki/MIT_License)
