package main

import (
	parserL "Projects/src/ANTLR/text-follow-along/errors/go/ambig/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"strconv"
)

type ambigListener struct {
	*parserL.BaseAmbigListener
}

func main() {
	input := antlr.NewInputStream(`f();`)
	lexer := parserL.NewAmbigLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	ambigParser := parserL.NewAmbigParser(tokens)
	ambigParser.RemoveErrorListeners()
	ambigParser.AddErrorListener(&antlr.DiagnosticErrorListener{})
	ambigParser.AddErrorListener(&verboseErrorListener{})
	ambigParser.GetInterpreter().SetPredictionMode(antlr.PredictionModeLLExactAmbigDetection)
	tree := ambigParser.Stat()
	walker := antlr.NewParseTreeWalker()
	listener := &ambigListener{}
	walker.Walk(listener, tree)
}

type verboseErrorListener struct {
	*antlr.DefaultErrorListener
	//antlr.ErrorListener
}

func (vel *verboseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	//fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" at "+litter.Sdump(offendingSymbol)+": "+msg)
	fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+": "+msg)
}
