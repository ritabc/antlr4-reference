// Code generated from Simple.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Simple

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 64, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 6, 2, 14,
	10, 2, 13, 2, 14, 2, 15, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 22, 10, 3, 13, 3,
	14, 3, 23, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 55, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 62, 10, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10,
	2, 2, 2, 63, 2, 13, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8,
	54, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 14, 5, 4, 3, 2, 13, 12, 3, 2, 2,
	2, 14, 15, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 3, 3,
	2, 2, 2, 17, 18, 7, 3, 2, 2, 18, 19, 7, 12, 2, 2, 19, 21, 7, 4, 2, 2, 20,
	22, 5, 6, 4, 2, 21, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 21, 3, 2, 2,
	2, 23, 24, 3, 2, 2, 2, 24, 25, 3, 2, 2, 2, 25, 26, 7, 5, 2, 2, 26, 27,
	8, 3, 1, 2, 27, 5, 3, 2, 2, 2, 28, 29, 7, 6, 2, 2, 29, 30, 7, 12, 2, 2,
	30, 31, 7, 7, 2, 2, 31, 43, 8, 4, 1, 2, 32, 33, 7, 6, 2, 2, 33, 34, 7,
	12, 2, 2, 34, 35, 7, 8, 2, 2, 35, 36, 7, 12, 2, 2, 36, 37, 7, 9, 2, 2,
	37, 38, 7, 4, 2, 2, 38, 39, 5, 8, 5, 2, 39, 40, 7, 5, 2, 2, 40, 41, 8,
	4, 1, 2, 41, 43, 3, 2, 2, 2, 42, 28, 3, 2, 2, 2, 42, 32, 3, 2, 2, 2, 43,
	7, 3, 2, 2, 2, 44, 45, 5, 10, 6, 2, 45, 46, 7, 7, 2, 2, 46, 47, 8, 5, 1,
	2, 47, 55, 3, 2, 2, 2, 48, 49, 7, 12, 2, 2, 49, 50, 7, 10, 2, 2, 50, 51,
	5, 10, 6, 2, 51, 52, 7, 7, 2, 2, 52, 53, 8, 5, 1, 2, 53, 55, 3, 2, 2, 2,
	54, 44, 3, 2, 2, 2, 54, 48, 3, 2, 2, 2, 55, 9, 3, 2, 2, 2, 56, 62, 7, 11,
	2, 2, 57, 58, 7, 12, 2, 2, 58, 59, 7, 8, 2, 2, 59, 60, 7, 11, 2, 2, 60,
	62, 7, 9, 2, 2, 61, 56, 3, 2, 2, 2, 61, 57, 3, 2, 2, 2, 62, 11, 3, 2, 2,
	2, 7, 15, 23, 42, 54, 61,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'class'", "'{'", "'}'", "'int'", "';'", "'('", "')'", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "INT", "ID", "WS",
}

var ruleNames = []string{
	"prog", "classDef", "member", "stat", "expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SimpleParser struct {
	*antlr.BaseParser
}

func NewSimpleParser(input antlr.TokenStream) *SimpleParser {
	this := new(SimpleParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Simple.g4"

	return this
}

// SimpleParser tokens.
const (
	SimpleParserEOF  = antlr.TokenEOF
	SimpleParserT__0 = 1
	SimpleParserT__1 = 2
	SimpleParserT__2 = 3
	SimpleParserT__3 = 4
	SimpleParserT__4 = 5
	SimpleParserT__5 = 6
	SimpleParserT__6 = 7
	SimpleParserT__7 = 8
	SimpleParserINT  = 9
	SimpleParserID   = 10
	SimpleParserWS   = 11
)

// SimpleParser rules.
const (
	SimpleParserRULE_prog     = 0
	SimpleParserRULE_classDef = 1
	SimpleParserRULE_member   = 2
	SimpleParserRULE_stat     = 3
	SimpleParserRULE_expr     = 4
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllClassDef() []IClassDefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassDefContext)(nil)).Elem())
	var tst = make([]IClassDefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassDefContext)
		}
	}

	return tst
}

func (s *ProgContext) ClassDef(i int) IClassDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassDefContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *SimpleParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SimpleParserRULE_prog)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SimpleParserT__0 {
		{
			p.SetState(10)
			p.ClassDef()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID    antlr.Token
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleParserRULE_classDef
	return p
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) Get_ID() antlr.Token { return s._ID }

func (s *ClassDefContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ClassDefContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleParserID, 0)
}

func (s *ClassDefContext) AllMember() []IMemberContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMemberContext)(nil)).Elem())
	var tst = make([]IMemberContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMemberContext)
		}
	}

	return tst
}

func (s *ClassDefContext) Member(i int) IMemberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMemberContext)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *SimpleParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SimpleParserRULE_classDef)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.Match(SimpleParserT__0)
	}
	{
		p.SetState(16)

		var _m = p.Match(SimpleParserID)

		localctx.(*ClassDefContext)._ID = _m
	}
	{
		p.SetState(17)
		p.Match(SimpleParserT__1)
	}
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SimpleParserT__3 {
		{
			p.SetState(18)
			p.Member()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(23)
		p.Match(SimpleParserT__2)
	}
	fmt.Println("class " + (func() string {
		if localctx.(*ClassDefContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ClassDefContext).Get_ID().GetText()
		}
	}()))

	return localctx
}

// IMemberContext is an interface to support dynamic dispatch.
type IMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetF returns the f token.
	GetF() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetF sets the f token.
	SetF(antlr.Token)

	// IsMemberContext differentiates from other interfaces.
	IsMemberContext()
}

type MemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID    antlr.Token
	f      antlr.Token
}

func NewEmptyMemberContext() *MemberContext {
	var p = new(MemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleParserRULE_member
	return p
}

func (*MemberContext) IsMemberContext() {}

func NewMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberContext {
	var p = new(MemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleParserRULE_member

	return p
}

func (s *MemberContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberContext) Get_ID() antlr.Token { return s._ID }

func (s *MemberContext) GetF() antlr.Token { return s.f }

func (s *MemberContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *MemberContext) SetF(v antlr.Token) { s.f = v }

func (s *MemberContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(SimpleParserID)
}

func (s *MemberContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(SimpleParserID, i)
}

func (s *MemberContext) Stat() IStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *MemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.EnterMember(s)
	}
}

func (s *MemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.ExitMember(s)
	}
}

func (p *SimpleParser) Member() (localctx IMemberContext) {
	localctx = NewMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SimpleParserRULE_member)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Match(SimpleParserT__3)
		}
		{
			p.SetState(27)

			var _m = p.Match(SimpleParserID)

			localctx.(*MemberContext)._ID = _m
		}
		{
			p.SetState(28)
			p.Match(SimpleParserT__4)
		}
		fmt.Println("var " + (func() string {
			if localctx.(*MemberContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*MemberContext).Get_ID().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(SimpleParserT__3)
		}
		{
			p.SetState(31)

			var _m = p.Match(SimpleParserID)

			localctx.(*MemberContext).f = _m
		}
		{
			p.SetState(32)
			p.Match(SimpleParserT__5)
		}
		{
			p.SetState(33)
			p.Match(SimpleParserID)
		}
		{
			p.SetState(34)
			p.Match(SimpleParserT__6)
		}
		{
			p.SetState(35)
			p.Match(SimpleParserT__1)
		}
		{
			p.SetState(36)
			p.Stat()
		}
		{
			p.SetState(37)
			p.Match(SimpleParserT__2)
		}
		fmt.Println("method: " + (func() string {
			if localctx.(*MemberContext).GetF() == nil {
				return ""
			} else {
				return localctx.(*MemberContext).GetF().GetText()
			}
		}()))

	}

	return localctx
}

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr returns the _expr rule contexts.
	Get_expr() IExprContext

	// Set_expr sets the _expr rule contexts.
	Set_expr(IExprContext)

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_expr  IExprContext
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SimpleParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Get_expr() IExprContext { return s._expr }

func (s *StatContext) Set_expr(v IExprContext) { s._expr = v }

func (s *StatContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StatContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleParserID, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *SimpleParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SimpleParserRULE_stat)

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

	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)

			var _x = p.Expr()

			localctx.(*StatContext)._expr = _x
		}
		{
			p.SetState(43)
			p.Match(SimpleParserT__4)
		}
		fmt.Println("found expr: " + (func() string {
			if localctx.(*StatContext).Get_expr() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*StatContext).Get_expr().GetStart(), localctx.(*StatContext)._expr.GetStop())
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(46)
			p.Match(SimpleParserID)
		}
		{
			p.SetState(47)
			p.Match(SimpleParserT__7)
		}
		{
			p.SetState(48)

			var _x = p.Expr()

			localctx.(*StatContext)._expr = _x
		}
		{
			p.SetState(49)
			p.Match(SimpleParserT__4)
		}
		fmt.Println("found assign: " + (func() string {
			if localctx.(*StatContext).Get_expr() == nil {
				return ""
			} else {
				return p.GetTokenStream().GetTextFromTokens(localctx.(*StatContext).Get_expr().GetStart(), localctx.(*StatContext)._expr.GetStop())
			}
		}()))

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
	p.RuleIndex = SimpleParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SimpleParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) INT() antlr.TerminalNode {
	return s.GetToken(SimpleParserINT, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(SimpleParserID, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SimpleListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *SimpleParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SimpleParserRULE_expr)

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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SimpleParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(54)
			p.Match(SimpleParserINT)
		}

	case SimpleParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(55)
			p.Match(SimpleParserID)
		}
		{
			p.SetState(56)
			p.Match(SimpleParserT__5)
		}
		{
			p.SetState(57)
			p.Match(SimpleParserINT)
		}
		{
			p.SetState(58)
			p.Match(SimpleParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
