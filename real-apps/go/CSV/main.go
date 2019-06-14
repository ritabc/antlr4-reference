package main

import (
	"ANTLR/antlr/real-apps/go/CSV/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Setup input
var input = `Details,Month,Amount
Mid Bonus,June,"$2,000"
,January, """zippo"""
Total Bonuses,"","$5,000"
`

const EMPTY = ""

// Load array of row maps. Each map maps field name to value
var rows = make([]map[string]string, len(input))

// List column names
var headers = []string{}

// List of fields in current row
var currentRowFieldValues = []string{}

// type csvBaseParser parser.BaseCSVListener

type csvListener struct {
	*parser.BaseCSVListener
}

func main() {
	is := antlr.NewInputStream(input)
	lexer := parser.NewCSVLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	csvParser := parser.NewCSVParser(tokens)
	tree := csvParser.File()
	fmt.Println(tree.ToStringTree(csvParser.RuleNames, csvParser))
	walker := antlr.NewParseTreeWalker()
	walker.Walk(&csvListener{}, tree)

}

// At the end of text or string, add field to current row
func (s *csvListener) ExitString(ctx *parser.StringContext) {
	// fmt.Printf("reached\n")
	currentRowFieldValues = append(currentRowFieldValues, ctx.STRING().GetText())
}
func (s *csvListener) ExitText(ctx *parser.TextContext) {
	fmt.Println(ctx.TEXT().GetText())

	currentRowFieldValues = append(currentRowFieldValues, ctx.TEXT().GetText())
}

// func (s *csvListener) EnterField(ctx *parser.TextContext) {
// 	fmt.Printf("reached\n")
// 	currentRowFieldValues = append(currentRowFieldValues, ctx.TEXT().GetText())
// }

// If field is empty, add empty string to current row
func (s *csvListener) ExitEmpty(ctx *parser.EmptyContext) {
	// ctx.ToStringTree()
	currentRowFieldValues = append(currentRowFieldValues, EMPTY)
}

// Compose headers
func (s *csvListener) ExitHdr(ctx *parser.HdrContext) {

	for _, row := range currentRowFieldValues {
		headers = append(headers, row)
	}
}

// clear currentRowFieldValues
func (s *csvListener) EnterRow(ctx *parser.RowContext) {
	fmt.Printf("reached\n")
	currentRowFieldValues = []string{}
}

// After each row:
func (s *csvListener) ExitRow(ctx *parser.RowContext) {
	// If a header row, do nothing
	if ctx.BaseParserRuleContext.GetRuleContext().GetRuleIndex() == parser.CSVParserRULE_hdr {
		return
	}
	// If a data row
	m := make(map[string]string)
	for i, cellValue := range currentRowFieldValues {
		if len(headers)-1 >= i {
			m[headers[i]] = cellValue
		}
		// fmt.Println(cellValue)
	}
	rows = append(rows, m)
}
