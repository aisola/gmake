package main

// Token Types
const T_EOF string = "T_EOF"
const T_DIRECT string = "T_DIRECT"
const T_CMDPART string = "T_CMDPART"

const T_LPAREN string = "T_LPAREN"
const T_RPAREN string = "T_RPAREN"

const T_LCBRAC string = "T_LCBRAC"
const T_RCBRAC string = "T_RCBRAC"

const T_COMMA string = "T_COMMA"
const T_SEMI string = "T_SEMI"

// asts

type Directive struct {
	Name string
	// Dependencies []string
	Commands []Command
}

type Command struct {
	Parts []string
}
