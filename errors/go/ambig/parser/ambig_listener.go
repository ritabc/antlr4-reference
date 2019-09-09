// Code generated from Ambig.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Ambig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AmbigListener is a complete listener for a parse tree produced by AmbigParser.
type AmbigListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
