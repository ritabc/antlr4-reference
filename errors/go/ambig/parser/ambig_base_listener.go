// Code generated from Ambig.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Ambig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAmbigListener is a complete listener for a parse tree produced by AmbigParser.
type BaseAmbigListener struct{}

var _ AmbigListener = &BaseAmbigListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAmbigListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAmbigListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAmbigListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAmbigListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseAmbigListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseAmbigListener) ExitStat(ctx *StatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseAmbigListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseAmbigListener) ExitExpr(ctx *ExprContext) {}
