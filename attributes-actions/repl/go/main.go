package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/antlr/antlr4/runtime/Go/antlr"
    parserL "Projects/src/ANTLR/text-follow-along/attributes-actions/tools/go/exprTools"

)

func printInvalidCmd(text string) {
    // We might have a panic here so we must defer & recover
    defer recoverExp(text)

    // \n will be ignored
    t := strings.TrimSuffix(text, "\n")
    if t != "" {
        expression := text
        fmt.Println("go-repl>", expression)
    }
}

func recoverExp(text string) {
    if r := recover(); r != nil {
        fmt.Println("go-repl> unknown command ", text)
    }
}

type exprListener struct {
    *parserL.BaseExprListener
}

func main() {
    // read input, print it back out
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome To go-repl")
    fmt.Print("go-repl> ")
    text, _ := reader.ReadString('\n')
    for ; ; text, _ = reader.ReadString('\n') {
        printInvalidCmd(text)
        input := antlr.NewInputStream(text+"\n")
        lexer := parserL.NewExprLexer(input)
        tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
        exprParser := parserL.NewExprParser(tokens)
        tree := exprParser.Stat()
        walker := antlr.NewParseTreeWalker()
        listener := &exprListener{}
        walker.Walk(listener, tree)
    }
}
