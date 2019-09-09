// Code generated from Simple.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Simple

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleListener is a complete listener for a parse tree produced by SimpleParser.
type BaseSimpleListener struct{}

var _ SimpleListener = &BaseSimpleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseSimpleListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseSimpleListener) ExitProg(ctx *ProgContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseSimpleListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseSimpleListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterMember is called when production member is entered.
func (s *BaseSimpleListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseSimpleListener) ExitMember(ctx *MemberContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseSimpleListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseSimpleListener) ExitStat(ctx *StatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSimpleListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSimpleListener) ExitExpr(ctx *ExprContext) {}
