package main

import (
	jsonParser "ANTLR/text-follow-along/real-apps/go/json-xml/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	// "github.com/sanity-io/litter"
	"os"
	// "strconv"
	"strings"
)

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := jsonParser.NewJSONLexer(input)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := jsonParser.NewJSONParser(tokens)
	tree := parser.Json()
	// fmt.Printf("Tree: %v\n", tree.ToStringTree(parser.RuleNames, parser))
	walker := antlr.NewParseTreeWalker()
	listener := &jsonListener{}
	walker.Walk(listener, tree)
	listener.printAllXML()
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

func (l *jsonListener) EnterJson(ctx *jsonParser.JsonContext) {
	l.xml = make(map[antlr.ParseTree]string)
}

// The Atom alternatives of value 'return' or annotate the Atom node with, the text of the matched token
func (l *jsonListener) ExitAtom(ctx *jsonParser.AtomContext) {
	l.setXML(ctx, ctx.GetText())
}

func (l *jsonListener) ExitStringValue(ctx *jsonParser.StringValueContext) {
	l.setXML(ctx, stripQuotes(ctx.GetText()))
}

func (l *jsonListener) ExitArrayValue(ctx *jsonParser.ArrayValueContext) {
	l.setXML(ctx, l.getXML(ctx.Array()))
}

func (l *jsonListener) ExitPair(ctx *jsonParser.PairContext) {
	tag := stripQuotes(ctx.STRING().GetText())
	vctx := ctx.Value()
	x := fmt.Sprintf("<%v>%v</%v>\n", tag, l.getXML(vctx), tag)
	l.setXML(ctx, x)
}

func (l *jsonListener) ExitAnObject(ctx *jsonParser.AnObjectContext) {
	var output strings.Builder
	output.WriteString("\n")
	for _, iPairContext := range ctx.AllPair() {
		pctx := iPairContext.GetRuleContext()
		write := l.getXML(pctx)
		output.WriteString(write)
	}
	l.setXML(ctx, output.String())
}

func (l *jsonListener) ExitObjectValue(ctx *jsonParser.ObjectValueContext) {
	l.setXML(ctx, l.getXML(ctx.Object()))
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
	l.setXML(ctx, l.getXML(ctx))
}

func stripQuotes(quote string) string {
	out := []byte(quote)
	if len(out) == 0 {
		return quote
	}
	out = out[1 : len(out)-1]
	return string(out)
}

// func (l jsonListener) printForDebugging() {
// 	var out strings.Builder
// 	out.WriteString("\nprintForDebugging Start\n")
// 	out.WriteString(strconv.Itoa(len(l.xml)))
// 	for ctx, value := range l.xml {
// 		pair := fmt.Sprintf("(Some Context with text) %v : %v (value)\n", ctx.GetText(), value)
// 		out.WriteString(pair)
// 	}
// 	out.WriteString("\nprintForDebugging End\n")
// 	fmt.Print(out.String())
// }

func (l jsonListener) printAllXML() {
	var out strings.Builder
	for _, value := range l.xml {
		out.WriteString(value)
	}
	fmt.Println(out.String())
}
