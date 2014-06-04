package main

import "fmt"

// parserState represents the state of the scanner
// as a function that returns the next state.
type parserState func(*parser) parserState

// Parse creates a new parser with the recommended
// parameters.
func Parse(name string, tokens []LexToken) []*Directive {
	p := &parser{
		name:   name,
		tokens: tokens,
		pos:    -1,
	}
	p.run()
	return p.ast
}

// the parser type
type parser struct {
	name   string
	tokens []LexToken
	pos    int

	ast              []*Directive // the ast
	currentDirective *Directive
	cmdparts         Command
}

func (p *parser) newDirective(name string) {
	p.currentDirective = &Directive{Name: name, Commands: make([]Command, 0)}
	p.cmdparts = Command{Parts: make([]string, 0)}
}

func (p *parser) closeDirective() {
	p.ast = append(p.ast, p.currentDirective)
}

func (p *parser) addCmdPart(part string) {
	p.cmdparts.Parts = append(p.cmdparts.Parts, part)
}

func (p *parser) flushCommand() {
	p.currentDirective.Commands = append(p.currentDirective.Commands, p.cmdparts)
	p.cmdparts = Command{Parts: make([]string, 0)}
}

// peek returns what the next token is but does NOT
// advance the position.
func (p *parser) peek() *LexToken {
	tok := p.next()
	p.backup()
	return tok
}

// nest returns what the next token AND
// advances p.pos.
func (p *parser) next() *LexToken {
	if p.pos >= len(p.tokens) {
		return nil
	}
	p.pos += 1
	return &p.tokens[p.pos]
}

// backup sets the position back one token
func (p *parser) backup() {
	p.pos -= 1
}

// run starts the statemachine
func (p *parser) run() {
	for state := initialParserState; state != nil; {
		state = state(p)
	}
}

// the starting state for parsing
func initialParserState(p *parser) parserState {

	for t := p.next(); t[0] != T_EOF; t = p.next() {

		if t[0] == T_DIRECT {
			p.newDirective(t[1])

		} else if t[0] == T_LCBRAC {
			return commandsState

		} else {
			fmt.Printf("gmake:%s: unexpected '%s' expecting T_DIRECT\n", t[2], t[0])
			return nil

		}

	}

	return nil
}

func commandsState(p *parser) parserState {

	for t := p.next(); t[0] != T_RCBRAC; t = p.next() {

		if t[0] == T_CMDPART {
			p.addCmdPart(t[1])

		} else if t[0] == T_SEMI {
			p.flushCommand()

		} else {
			fmt.Printf("gmake:%s: unexpected '%s' expecting T_CMDPART or T_SEMI\n", t[2], t[0])
			return nil

		}

	}
	p.closeDirective()
	return initialParserState
}
