// Code generated from Ambig.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 38, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 6, 5, 23, 10, 5, 13, 5, 14,
	5, 24, 3, 6, 6, 6, 28, 10, 6, 13, 6, 14, 6, 29, 3, 7, 6, 7, 33, 10, 7,
	13, 7, 14, 7, 34, 3, 7, 3, 7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 3, 2, 5, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 40, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 3, 15, 3, 2, 2, 2,
	5, 17, 3, 2, 2, 2, 7, 19, 3, 2, 2, 2, 9, 22, 3, 2, 2, 2, 11, 27, 3, 2,
	2, 2, 13, 32, 3, 2, 2, 2, 15, 16, 7, 61, 2, 2, 16, 4, 3, 2, 2, 2, 17, 18,
	7, 42, 2, 2, 18, 6, 3, 2, 2, 2, 19, 20, 7, 43, 2, 2, 20, 8, 3, 2, 2, 2,
	21, 23, 9, 2, 2, 2, 22, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 22, 3,
	2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 10, 3, 2, 2, 2, 26, 28, 9, 3, 2, 2, 27,
	26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2,
	2, 30, 12, 3, 2, 2, 2, 31, 33, 9, 4, 2, 2, 32, 31, 3, 2, 2, 2, 33, 34,
	3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 36, 3, 2, 2, 2,
	36, 37, 8, 7, 2, 2, 37, 14, 3, 2, 2, 2, 6, 2, 24, 29, 34, 3, 8, 2, 2,
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
	"", "';'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "INT", "ID", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "INT", "ID", "WS",
}

type AmbigLexer struct {
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

func NewAmbigLexer(input antlr.CharStream) *AmbigLexer {

	l := new(AmbigLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Ambig.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AmbigLexer tokens.
const (
	AmbigLexerT__0 = 1
	AmbigLexerT__1 = 2
	AmbigLexerT__2 = 3
	AmbigLexerINT  = 4
	AmbigLexerID   = 5
	AmbigLexerWS   = 6
)
