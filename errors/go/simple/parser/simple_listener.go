// Code generated from Simple.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Simple

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SimpleListener is a complete listener for a parse tree produced by SimpleParser.
type SimpleListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterClassDef is called when entering the classDef production.
	EnterClassDef(c *ClassDefContext)

	// EnterMember is called when entering the member production.
	EnterMember(c *MemberContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitClassDef is called when exiting the classDef production.
	ExitClassDef(c *ClassDefContext)

	// ExitMember is called when exiting the member production.
	ExitMember(c *MemberContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
