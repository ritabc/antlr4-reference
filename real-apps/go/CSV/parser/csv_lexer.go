// Code generated from CSV.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 39, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 6, 4, 21, 10, 4, 13, 4, 14, 4, 22, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 29, 10, 5, 12, 5, 14, 5, 32, 11, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 3, 2, 5, 6, 2, 12, 12, 15, 15, 36, 36, 46, 46, 3, 2, 36, 36, 4, 2, 11,
	11, 34, 34, 2, 41, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 3, 15, 3, 2, 2,
	2, 5, 17, 3, 2, 2, 2, 7, 20, 3, 2, 2, 2, 9, 24, 3, 2, 2, 2, 11, 35, 3,
	2, 2, 2, 13, 37, 3, 2, 2, 2, 15, 16, 7, 15, 2, 2, 16, 4, 3, 2, 2, 2, 17,
	18, 7, 12, 2, 2, 18, 6, 3, 2, 2, 2, 19, 21, 10, 2, 2, 2, 20, 19, 3, 2,
	2, 2, 21, 22, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 8,
	3, 2, 2, 2, 24, 30, 7, 36, 2, 2, 25, 26, 7, 36, 2, 2, 26, 29, 7, 36, 2,
	2, 27, 29, 10, 3, 2, 2, 28, 25, 3, 2, 2, 2, 28, 27, 3, 2, 2, 2, 29, 32,
	3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2,
	32, 30, 3, 2, 2, 2, 33, 34, 7, 36, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7,
	46, 2, 2, 36, 12, 3, 2, 2, 2, 37, 38, 9, 4, 2, 2, 38, 14, 3, 2, 2, 2, 6,
	2, 22, 28, 30, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\r'", "'\n'", "", "", "','",
}

var lexerSymbolicNames = []string{
	"", "", "", "TEXT", "STRING", "COMMA", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "TEXT", "STRING", "COMMA", "WS",
}

type CSVLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCSVLexer(input antlr.CharStream) *CSVLexer {

	l := new(CSVLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CSV.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CSVLexer tokens.
const (
	CSVLexerT__0   = 1
	CSVLexerT__1   = 2
	CSVLexerTEXT   = 3
	CSVLexerSTRING = 4
	CSVLexerCOMMA  = 5
	CSVLexerWS     = 6
)
