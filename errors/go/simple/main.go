package main

import (
	parserL "Projects/src/ANTLR/text-follow-along/errors/go/simple/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sanity-io/litter"
	"os"
	"strconv"
	"strings"
)

type simpleListener struct {
	*parserL.BaseSimpleListener
}

type verboseErrorListener struct {
	*antlr.DefaultErrorListener
	//antlr.ErrorListener
}

func (vel *verboseErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	//fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" at "+litter.Sdump(offendingSymbol)+": "+msg)
	fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+": "+msg)
}

type underlineErrorListener struct {
	//antlr.ErrorListener
	*antlr.DefaultErrorListener
}

func (uel *underlineErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	fmt.Fprintln(os.Stderr, "line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg)
	underlineError(recognizer, offendingSymbol.(antlr.Token), line, column, e)
}

func underlineError(recognizer antlr.Recognizer, offendingToken antlr.Token, line int, column int, e antlr.RecognitionException) {
	//lexer := offendingToken.GetInputStream().(antlr.CommonTokenStream)
	//lexer := e.GetInputStream().(antlr.CommonTokenStream)
	//tokens := e.GetInputStream().(antlr.TokenStream)
	//input := tokens.GetTokenSource().GetInputStream().GetText(0, len(tokens.GetAllText())+1)
	//recognizer.(antlr.TokenStream)
	//lines := strings.Split(input, "\n")
	//lines := offendingToken.GetText(offendingToken.GetInputStream())
	is := recognizer.(*antlr.BaseParser).GetTokenStream()
	input := is.GetTextFromInterval(&antlr.Interval{
		Start: 0,
		Stop:  column,
	})
	litter.Dump("input")
	litter.Dump(input)
	lines := strings.Split(input, "\n")
	litter.Dump("lines")
	litter.Dump(lines)
	errorLine := lines[line]
	fmt.Fprint(os.Stderr, errorLine)
	for i := 0; i < column; i++ {
		fmt.Fprint(os.Stderr, " ")
	}
	start := offendingToken.GetStart()
	stop := offendingToken.GetStop()
	if start >= 0 && stop >= 0 {
		for i := start; i <= stop; i++ {
			fmt.Fprint(os.Stderr, "^")
		}
		fmt.Fprintln(os.Stderr, "")
	}
}

func main() {
	input := antlr.NewInputStream(`class T XYZ {
int ;
}
`)
	lexer := parserL.NewSimpleLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	simpleParser := parserL.NewSimpleParser(tokens)
	simpleParser.RemoveErrorListeners()                      // remove ConsoleErrorListener
	simpleParser.AddErrorListener(&underlineErrorListener{}) // add ours
	//simpleParser.AddErrorListener(&verboseErrorListener{}) // add ours
	tree := simpleParser.Prog()
	walker := antlr.NewParseTreeWalker()
	listener := &simpleListener{}
	walker.Walk(listener, tree)
}
