// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterE is called when entering the e production.
	EnterE(c *EContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitE is called when exiting the e production.
	ExitE(c *EContext)
}
