package main

import (
	"ANTLR/real-apps/go/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Setup input
const input = `Details,Month,Amount
	Mid Bonus,June,"$2,000"
	,January,"""zippo"""
	Total Bonuses, "","$%,000"`

const EMPTY = ""

// Load array of row maps. Each map maps field name to value
var rows = make([]map[string]string, len(input))

// List column names
var headers = []string{}

// List of fields in current row
var currentRowFieldValues = []string{}

type csvParser struct {
	*parser.BaseCSVListener
}

func main() {

}

func exitString(ctx parser.StringContext) {
	currentRowFieldValues := append(currentRowFieldValues, ctx.STRING().GetText())
}

func exitText(ctx parser.TextContext) {
	currentRowFieldValues := append(currentRowFieldValues, ctx.TEXT().GetText())
}

func exitEmpty(ctx parser.EmptyContext) {
	currentRowFieldValues := append(currentRowFieldValues, EMPTY)
}
