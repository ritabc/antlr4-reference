// Code generated from JSON.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 14, 99, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11,
	60, 10, 11, 12, 11, 14, 11, 63, 11, 11, 3, 11, 3, 11, 3, 12, 5, 12, 68,
	10, 12, 3, 12, 3, 12, 3, 12, 6, 12, 73, 10, 12, 13, 12, 14, 12, 74, 3,
	12, 5, 12, 78, 10, 12, 3, 12, 5, 12, 81, 10, 12, 3, 13, 3, 13, 3, 13, 7,
	13, 86, 10, 13, 12, 13, 14, 13, 89, 11, 13, 5, 13, 91, 10, 13, 3, 14, 6,
	14, 94, 10, 14, 13, 14, 14, 14, 95, 3, 14, 3, 14, 2, 2, 15, 3, 3, 5, 4,
	7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 2,
	27, 14, 3, 2, 6, 4, 2, 36, 36, 94, 94, 3, 2, 50, 59, 3, 2, 51, 59, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 105, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2,
	2, 2, 7, 33, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 39,
	3, 2, 2, 2, 15, 44, 3, 2, 2, 2, 17, 50, 3, 2, 2, 2, 19, 55, 3, 2, 2, 2,
	21, 57, 3, 2, 2, 2, 23, 80, 3, 2, 2, 2, 25, 90, 3, 2, 2, 2, 27, 93, 3,
	2, 2, 2, 29, 30, 7, 125, 2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 46, 2, 2,
	32, 6, 3, 2, 2, 2, 33, 34, 7, 127, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 7,
	93, 2, 2, 36, 10, 3, 2, 2, 2, 37, 38, 7, 95, 2, 2, 38, 12, 3, 2, 2, 2,
	39, 40, 7, 118, 2, 2, 40, 41, 7, 116, 2, 2, 41, 42, 7, 119, 2, 2, 42, 43,
	7, 103, 2, 2, 43, 14, 3, 2, 2, 2, 44, 45, 7, 104, 2, 2, 45, 46, 7, 99,
	2, 2, 46, 47, 7, 110, 2, 2, 47, 48, 7, 117, 2, 2, 48, 49, 7, 103, 2, 2,
	49, 16, 3, 2, 2, 2, 50, 51, 7, 112, 2, 2, 51, 52, 7, 119, 2, 2, 52, 53,
	7, 110, 2, 2, 53, 54, 7, 110, 2, 2, 54, 18, 3, 2, 2, 2, 55, 56, 7, 60,
	2, 2, 56, 20, 3, 2, 2, 2, 57, 61, 7, 36, 2, 2, 58, 60, 10, 2, 2, 2, 59,
	58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2,
	2, 62, 64, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 36, 2, 2, 65, 22,
	3, 2, 2, 2, 66, 68, 7, 47, 2, 2, 67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2,
	68, 69, 3, 2, 2, 2, 69, 70, 5, 25, 13, 2, 70, 72, 7, 48, 2, 2, 71, 73,
	9, 3, 2, 2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	74, 75, 3, 2, 2, 2, 75, 81, 3, 2, 2, 2, 76, 78, 7, 47, 2, 2, 77, 76, 3,
	2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81, 5, 25, 13, 2,
	80, 67, 3, 2, 2, 2, 80, 77, 3, 2, 2, 2, 81, 24, 3, 2, 2, 2, 82, 91, 7,
	50, 2, 2, 83, 87, 9, 4, 2, 2, 84, 86, 9, 3, 2, 2, 85, 84, 3, 2, 2, 2, 86,
	89, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 91, 3, 2, 2,
	2, 89, 87, 3, 2, 2, 2, 90, 82, 3, 2, 2, 2, 90, 83, 3, 2, 2, 2, 91, 26,
	3, 2, 2, 2, 92, 94, 9, 5, 2, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 8,
	14, 2, 2, 98, 28, 3, 2, 2, 2, 11, 2, 61, 67, 74, 77, 80, 87, 90, 95, 3,
	8, 2, 2,
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
	"", "'{'", "','", "'}'", "'['", "']'", "'true'", "'false'", "'null'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"STRING", "NUMBER", "INT", "WS",
}

type JSONLexer struct {
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

func NewJSONLexer(input antlr.CharStream) *JSONLexer {

	l := new(JSONLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "JSON.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JSONLexer tokens.
const (
	JSONLexerT__0   = 1
	JSONLexerT__1   = 2
	JSONLexerT__2   = 3
	JSONLexerT__3   = 4
	JSONLexerT__4   = 5
	JSONLexerT__5   = 6
	JSONLexerT__6   = 7
	JSONLexerT__7   = 8
	JSONLexerT__8   = 9
	JSONLexerSTRING = 10
	JSONLexerNUMBER = 11
	JSONLexerWS     = 12
)
