// Code generated from Ambig.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Ambig

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 22, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 14,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 20, 10, 3, 3, 3, 2, 2, 4, 2, 4, 2,
	2, 2, 21, 2, 13, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8,
	7, 3, 2, 2, 8, 14, 3, 2, 2, 2, 9, 10, 7, 7, 2, 2, 10, 11, 7, 4, 2, 2, 11,
	12, 7, 5, 2, 2, 12, 14, 7, 3, 2, 2, 13, 6, 3, 2, 2, 2, 13, 9, 3, 2, 2,
	2, 14, 3, 3, 2, 2, 2, 15, 16, 7, 7, 2, 2, 16, 17, 7, 4, 2, 2, 17, 20, 7,
	5, 2, 2, 18, 20, 7, 6, 2, 2, 19, 15, 3, 2, 2, 2, 19, 18, 3, 2, 2, 2, 20,
	5, 3, 2, 2, 2, 4, 13, 19,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "';'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "INT", "ID", "WS",
}

var ruleNames = []string{
	"stat", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AmbigParser struct {
	*antlr.BaseParser
}

func NewAmbigParser(input antlr.TokenStream) *AmbigParser {
	this := new(AmbigParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Ambig.g4"

	return this
}

// AmbigParser tokens.
const (
	AmbigParserEOF  = antlr.TokenEOF
	AmbigParserT__0 = 1
	AmbigParserT__1 = 2
	AmbigParserT__2 = 3
	AmbigParserINT  = 4
	AmbigParserID   = 5
	AmbigParserWS   = 6
)

// AmbigParser rules.
const (
	AmbigParserRULE_stat = 0
	AmbigParserRULE_expr = 1
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AmbigParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbigParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbigParserID, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbigListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbigListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *AmbigParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AmbigParserRULE_stat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(4)
			p.Expr()
		}
		{
			p.SetState(5)
			p.Match(AmbigParserT__0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(7)
			p.Match(AmbigParserID)
		}
		{
			p.SetState(8)
			p.Match(AmbigParserT__1)
		}
		{
			p.SetState(9)
			p.Match(AmbigParserT__2)
		}
		{
			p.SetState(10)
			p.Match(AmbigParserT__0)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AmbigParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmbigParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(AmbigParserID, 0)
}

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(AmbigParserINT, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbigListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmbigListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *AmbigParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AmbigParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(17)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AmbigParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(13)
			p.Match(AmbigParserID)
		}
		{
			p.SetState(14)
			p.Match(AmbigParserT__1)
		}
		{
			p.SetState(15)
			p.Match(AmbigParserT__2)
		}

	case AmbigParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(16)
			p.Match(AmbigParserINT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
