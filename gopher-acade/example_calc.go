package main

import (
	"./parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(ctx *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch ctx.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected token: %s", ctx.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(ctx *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch ctx.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected token: %s", ctx.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(ctx *parser.NumberContext) {
	i, err := strconv.Atoi(ctx.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 +2*3")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// // Read all tokens
	// for {
	// 	t := lexer.NextToken()
	// 	if t.GetTokenType() == antlr.TokenEOF {
	// 		break
	// 	}
	// 	fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	// }
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)
	tree := p.Stmt()

	var listener calcListener

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	fmt.Printf("output: %v", listener.pop())
}
