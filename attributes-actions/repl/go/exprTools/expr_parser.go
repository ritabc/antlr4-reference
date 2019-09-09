// Code generated from Expr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Expr

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 46, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 17, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 29, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 7, 3, 41, 10, 3, 12, 3, 14, 3, 44, 11, 3, 3, 3, 2,
	3, 4, 4, 2, 4, 2, 4, 3, 2, 10, 11, 3, 2, 12, 13, 2, 49, 2, 16, 3, 2, 2,
	2, 4, 28, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 8, 2, 2, 8, 17, 3, 2,
	2, 2, 9, 10, 7, 6, 2, 2, 10, 11, 7, 3, 2, 2, 11, 12, 5, 4, 3, 2, 12, 13,
	7, 8, 2, 2, 13, 14, 8, 2, 1, 2, 14, 17, 3, 2, 2, 2, 15, 17, 7, 8, 2, 2,
	16, 6, 3, 2, 2, 2, 16, 9, 3, 2, 2, 2, 16, 15, 3, 2, 2, 2, 17, 3, 3, 2,
	2, 2, 18, 19, 8, 3, 1, 2, 19, 20, 7, 7, 2, 2, 20, 29, 8, 3, 1, 2, 21, 22,
	7, 6, 2, 2, 22, 29, 8, 3, 1, 2, 23, 24, 7, 4, 2, 2, 24, 25, 5, 4, 3, 2,
	25, 26, 7, 5, 2, 2, 26, 27, 8, 3, 1, 2, 27, 29, 3, 2, 2, 2, 28, 18, 3,
	2, 2, 2, 28, 21, 3, 2, 2, 2, 28, 23, 3, 2, 2, 2, 29, 42, 3, 2, 2, 2, 30,
	31, 12, 7, 2, 2, 31, 32, 9, 2, 2, 2, 32, 33, 5, 4, 3, 8, 33, 34, 8, 3,
	1, 2, 34, 41, 3, 2, 2, 2, 35, 36, 12, 6, 2, 2, 36, 37, 9, 3, 2, 2, 37,
	38, 5, 4, 3, 7, 38, 39, 8, 3, 1, 2, 39, 41, 3, 2, 2, 2, 40, 30, 3, 2, 2,
	2, 40, 35, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43,
	3, 2, 2, 2, 43, 5, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 6, 16, 28, 40, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'('", "')'", "", "", "", "", "'*'", "'/'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "ID", "INT", "NEWLINE", "WS", "MULT", "DIV", "PLUS", "MINUS",
}

var ruleNames = []string{
	"stat", "e",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ExprParser struct {
	*antlr.BaseParser
}

func NewExprParser(input antlr.TokenStream) *ExprParser {
	this := new(ExprParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Expr.g4"

	return this
}

/** "memory" for our calculator; variable/value pairs go here */
var memory = map[string]int{}

func eval(left, op, right int) int {
	switch op {
	case ExprParserMULT:
		return left * right
	case ExprParserDIV:
		return left / right
	case ExprParserPLUS:
		return left + right
	case ExprParserMINUS:
		return left - right
	default:
		return 0
	}
}

// ExprParser tokens.
const (
	ExprParserEOF     = antlr.TokenEOF
	ExprParserT__0    = 1
	ExprParserT__1    = 2
	ExprParserT__2    = 3
	ExprParserID      = 4
	ExprParserINT     = 5
	ExprParserNEWLINE = 6
	ExprParserWS      = 7
	ExprParserMULT    = 8
	ExprParserDIV     = 9
	ExprParserPLUS    = 10
	ExprParserMINUS   = 11
)

// ExprParser rules.
const (
	ExprParserRULE_stat = 0
	ExprParserRULE_e    = 1
)

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_e returns the _e rule contexts.
	Get_e() IEContext

	// Set_e sets the _e rule contexts.
	Set_e(IEContext)

	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_ID    antlr.Token
	_e     IEContext
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_stat
	return p
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) Get_ID() antlr.Token { return s._ID }

func (s *StatContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *StatContext) Get_e() IEContext { return s._e }

func (s *StatContext) Set_e(v IEContext) { s._e = v }

func (s *StatContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *StatContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ExprParserNEWLINE, 0)
}

func (s *StatContext) ID() antlr.TerminalNode {
	return s.GetToken(ExprParserID, 0)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterStat(s)
	}
}

func (s *StatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitStat(s)
	}
}

func (p *ExprParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExprParserRULE_stat)

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

	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(4)
			p.e(0)
		}
		{
			p.SetState(5)
			p.Match(ExprParserNEWLINE)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(7)

			var _m = p.Match(ExprParserID)

			localctx.(*StatContext)._ID = _m
		}
		{
			p.SetState(8)
			p.Match(ExprParserT__0)
		}
		{
			p.SetState(9)

			var _x = p.e(0)

			localctx.(*StatContext)._e = _x
		}
		{
			p.SetState(10)
			p.Match(ExprParserNEWLINE)
		}
		memory[(func() string {
			if localctx.(*StatContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*StatContext).Get_ID().GetText()
			}
		}())] = localctx.(*StatContext).Get_e().GetV()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(13)
			p.Match(ExprParserNEWLINE)
		}

	}

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INT returns the _INT token.
	Get_INT() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Set_INT sets the _INT token.
	Set_INT(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetA returns the a rule contexts.
	GetA() IEContext

	// Get_e returns the _e rule contexts.
	Get_e() IEContext

	// GetB returns the b rule contexts.
	GetB() IEContext

	// SetA sets the a rule contexts.
	SetA(IEContext)

	// Set_e sets the _e rule contexts.
	Set_e(IEContext)

	// SetB sets the b rule contexts.
	SetB(IEContext)

	// GetV returns the v attribute.
	GetV() int

	// SetV sets the v attribute.
	SetV(int)

	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	v      int
	a      IEContext
	_INT   antlr.Token
	_ID    antlr.Token
	_e     IEContext
	op     antlr.Token
	b      IEContext
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExprParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExprParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) Get_INT() antlr.Token { return s._INT }

func (s *EContext) Get_ID() antlr.Token { return s._ID }

func (s *EContext) GetOp() antlr.Token { return s.op }

func (s *EContext) Set_INT(v antlr.Token) { s._INT = v }

func (s *EContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *EContext) SetOp(v antlr.Token) { s.op = v }

func (s *EContext) GetA() IEContext { return s.a }

func (s *EContext) Get_e() IEContext { return s._e }

func (s *EContext) GetB() IEContext { return s.b }

func (s *EContext) SetA(v IEContext) { s.a = v }

func (s *EContext) Set_e(v IEContext) { s._e = v }

func (s *EContext) SetB(v IEContext) { s.b = v }

func (s *EContext) GetV() int { return s.v }

func (s *EContext) SetV(v int) { s.v = v }

func (s *EContext) INT() antlr.TerminalNode {
	return s.GetToken(ExprParserINT, 0)
}

func (s *EContext) ID() antlr.TerminalNode {
	return s.GetToken(ExprParserID, 0)
}

func (s *EContext) AllE() []IEContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEContext)(nil)).Elem())
	var tst = make([]IEContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEContext)
		}
	}

	return tst
}

func (s *EContext) E(i int) IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *EContext) MULT() antlr.TerminalNode {
	return s.GetToken(ExprParserMULT, 0)
}

func (s *EContext) DIV() antlr.TerminalNode {
	return s.GetToken(ExprParserDIV, 0)
}

func (s *EContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ExprParserPLUS, 0)
}

func (s *EContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ExprParserMINUS, 0)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.EnterE(s)
	}
}

func (s *EContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExprListener); ok {
		listenerT.ExitE(s)
	}
}

func (p *ExprParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *ExprParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ExprParserRULE_e, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExprParserINT:
		{
			p.SetState(17)

			var _m = p.Match(ExprParserINT)

			localctx.(*EContext)._INT = _m
		}
		localctx.(*EContext).v = (func() int {
			if localctx.(*EContext).Get_INT() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*EContext).Get_INT().GetText())
				return i
			}
		}())

	case ExprParserID:
		{
			p.SetState(19)

			var _m = p.Match(ExprParserID)

			localctx.(*EContext)._ID = _m
		}

		var id = (func() string {
			if localctx.(*EContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*EContext).Get_ID().GetText()
			}
		}())
		value, found := memory[id]
		if found {
			localctx.(*EContext).v = value
		} else {
			localctx.(*EContext).v = 0
		}

	case ExprParserT__1:
		{
			p.SetState(21)
			p.Match(ExprParserT__1)
		}
		{
			p.SetState(22)

			var _x = p.e(0)

			localctx.(*EContext)._e = _x
		}
		{
			p.SetState(23)
			p.Match(ExprParserT__2)
		}
		localctx.(*EContext).v = localctx.(*EContext).Get_e().GetV()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewEContext(p, _parentctx, _parentState)
				localctx.(*EContext).a = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_e)
				p.SetState(28)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(29)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserMULT || _la == ExprParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(30)

					var _x = p.e(6)

					localctx.(*EContext).b = _x
					localctx.(*EContext)._e = _x
				}
				localctx.(*EContext).v = eval(localctx.(*EContext).GetA().GetV(), (func() int {
					if localctx.(*EContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*EContext).GetOp().GetTokenType()
					}
				}()), localctx.(*EContext).GetB().GetV())

			case 2:
				localctx = NewEContext(p, _parentctx, _parentState)
				localctx.(*EContext).a = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ExprParserRULE_e)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(34)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ExprParserPLUS || _la == ExprParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(35)

					var _x = p.e(5)

					localctx.(*EContext).b = _x
					localctx.(*EContext)._e = _x
				}
				localctx.(*EContext).v = eval(localctx.(*EContext).GetA().GetV(), (func() int {
					if localctx.(*EContext).GetOp() == nil {
						return 0
					} else {
						return localctx.(*EContext).GetOp().GetTokenType()
					}
				}()), localctx.(*EContext).GetB().GetV())

			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ExprParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExprParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
