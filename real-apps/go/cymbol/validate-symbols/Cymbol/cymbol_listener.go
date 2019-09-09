// Code generated from Cymbol.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Cymbol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CymbolListener is a complete listener for a parse tree produced by CymbolParser.
type CymbolListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterCymbolType is called when entering the cymbolType production.
	EnterCymbolType(c *CymbolTypeContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterCall is called when entering the Call production.
	EnterCall(c *CallContext)

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterMult is called when entering the Mult production.
	EnterMult(c *MultContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterEqual is called when entering the Equal production.
	EnterEqual(c *EqualContext)

	// EnterVar is called when entering the Var production.
	EnterVar(c *VarContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterIndex is called when entering the Index production.
	EnterIndex(c *IndexContext)

	// EnterNegate is called when entering the Negate production.
	EnterNegate(c *NegateContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitCymbolType is called when exiting the cymbolType production.
	ExitCymbolType(c *CymbolTypeContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitCall is called when exiting the Call production.
	ExitCall(c *CallContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitMult is called when exiting the Mult production.
	ExitMult(c *MultContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitEqual is called when exiting the Equal production.
	ExitEqual(c *EqualContext)

	// ExitVar is called when exiting the Var production.
	ExitVar(c *VarContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitIndex is called when exiting the Index production.
	ExitIndex(c *IndexContext)

	// ExitNegate is called when exiting the Negate production.
	ExitNegate(c *NegateContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)
}
