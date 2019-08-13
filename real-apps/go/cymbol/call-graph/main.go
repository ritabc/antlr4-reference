package main

import (
    cParser "Projects/src/ANTLR/text-follow-along/real-apps/go/cymbol/Cymbol"
    "fmt"
    "github.com/antlr/antlr4/runtime/Go/antlr"
    "os"
    "strings"
)

type functionListener struct {
    *cParser.BaseCymbolListener

    graph Graph
    currentFunctionName string
}

type Graph struct{
    functionNodes []string
    edges []functionCall // cant be a map b/c 1 func can call multiple functions
}

type functionCall struct {
    caller string
    callee string
}

func (g *Graph) toDOT() string {
    var buf strings.Builder
    buf.WriteString("digraph G{\n")
    buf.WriteString("  ranksep=.25;\n")
    buf.WriteString("  edge [arrowsize=.5]\n")
    buf.WriteString("  node [shape=circle, fontname=\"ArialNarrow\",\n")
    buf.WriteString("        fontsize=12, fixedsize=true, height=.45];\n")
    buf.WriteString("  ")

    for _, node := range g.functionNodes{
        buf.WriteString(node) // print all functionNodes first
        buf.WriteString("; ")
    }
    buf.WriteString("\n")
    for _, call := range g.edges {
        buf.WriteString("  ")
        buf.WriteString(call.caller)
        buf.WriteString(" -> ")
        buf.WriteString(call.callee)
        buf.WriteString(";\n")
    }
    buf.WriteString("}\n")
    return buf.String()
}

//func (fl *functionListener) EnterFile(ctx *cParser.FileContext) {
//}


func (fl *functionListener) EnterFunctionDecl(ctx *cParser.FunctionDeclContext) {
    fl.currentFunctionName = ctx.ID().GetText()
    fl.graph.functionNodes = append(fl.graph.functionNodes, fl.currentFunctionName)
}

func (fl *functionListener) ExitCall(ctx *cParser.CallContext) {
    // map current function to callee
    var funcCall functionCall
    funcCall.callee = ctx.ID().GetText()
    funcCall.caller = fl.currentFunctionName
    fl.graph.edges = append(fl.graph.edges, funcCall)
}

func main() {
    input, _ := antlr.NewFileStream(os.Args[1])
    lexer := cParser.NewCymbolLexer(input)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := cParser.NewCymbolParser(tokens)
    tree := parser.File()
    walker := antlr.NewParseTreeWalker()
    listener := &functionListener{}
    //listener.graph.edges = make(map[string]string)
    walker.Walk(listener, tree)
    fmt.Printf("%v", listener.graph.toDOT())
}
