package main

import (
	jsonParser "ANTLR/antlr/real-apps/go/json-xml/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"reflect"
	"strings"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := jsonParser.NewJSONLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := jsonParser.NewJSONParser(tokens)
	tree := parser.Json()
	fmt.Printf("Tree: %v\n", tree.ToStringTree(parser.RuleNames, parser))
	walker := antlr.NewParseTreeWalker()
	listener := &jsonListener{}
	listener.xml = make(map[antlr.ParseTree]string)
	walker.Walk(listener, tree)
	fmt.Println(listener.xml)
}

type jsonListener struct {
	*jsonParser.BaseJSONListener

	/*
		Create the annotations map:
		Associate a property with a parse tree node
		Useful with parse tree listeners that need to associate values with
		particular tree nodes, kind of like specifying a return value for the listener event method that visited a particular node
	*/
	// New field xml (annotations) map. keys: parseTree node and value is struct: new type with fields of what info we want to pass around (xml var: see page 125 and 134)
	xml map[antlr.ParseTree]string
}

func (l *jsonListener) getXML(ctx antlr.ParseTree) string {
	value, found := l.xml[ctx]
	if !found {
		return ""
	}
	return value
}

func (l *jsonListener) setXML(ctx antlr.ParseTree, s string) {
	l.xml[ctx] = s
}

// The Atom alternatives of value 'return' or annotate the Atom node with, the text of the matched token
func (l *jsonListener) ExitAtom(ctx *jsonParser.AtomContext) {
	l.setXML(ctx, ctx.GetText())
}

func (l *jsonListener) ExitString(ctx *jsonParser.StringContext) {
	l.setXML(ctx, stripQuotes(ctx.GetText()))
}

func (l *jsonListener) ExitObjectValue(ctx *jsonParser.ObjectValueContext) {
	l.setXML(ctx, l.getXML(ctx.Object()))
}

func (l *jsonListener) ExitArrayValue(ctx *jsonParser.ArrayValueContext) {
	l.setXML(ctx, l.getXML(ctx.Array()))
}

func (l *jsonListener) ExitPair(ctx *jsonParser.PairContext) {
	tag := stripQuotes(ctx.STRING().GetText())
	vctx := ctx.Value()
	x := fmt.Sprintf("<%s>%s</%s>\n", tag, l.getXML(vctx), tag)
	l.setXML(vctx, x)
}

func (l *jsonListener) ExitAnObject(ctx *jsonParser.AnObjectContext) {
	var output strings.Builder
	output.WriteString("\n")
	for _, pairContext := range ctx.AllPair() {
		output.WriteString(l.getXML(pairContext))
	}
	l.setXML(ctx, output.String())
}

func (l *jsonListener) ExitBlankObject(ctx *jsonParser.BlankObjectContext) {
	l.setXML(ctx, "")
}

func (l *jsonListener) ExitArrayOfValues(ctx *jsonParser.ArrayOfValuesContext) {
	var output strings.Builder
	output.WriteString("\n")
	for _, valueContext := range ctx.AllValue() {
		output.WriteString("<element>")
		output.WriteString(l.getXML(valueContext))
		output.WriteString("</element>")
		output.WriteString("\n")
	}
	l.setXML(ctx, output.String())
}

func (l *jsonListener) ExitBlankArray(ctx *jsonParser.BlankArrayContext) {
	l.setXML(ctx, "")
}

func (l *jsonListener) ExitJson(ctx *jsonParser.JsonContext) {
	var parseTree antlr.ParseTree
	fmt.Println(ctx.GetChildOfType(0, reflect.TypeOf(parseTree)))
	l.setXML(ctx, l.getXML(ctx.GetChildOfType(0, reflect.TypeOf(parseTree))))
}

func stripQuotes(quote string) string {
	out := []byte(quote)
	out = out[1 : len(out)-1]
	return string(out)
}
