// Code generated from JSON.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JSON

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 60, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3, 24,
	11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	7, 4, 36, 10, 4, 12, 4, 14, 4, 39, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	45, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 54, 10, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 2, 2, 7, 2, 4, 6, 8, 10, 2, 2, 2, 65, 2, 14,
	3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10,
	55, 3, 2, 2, 2, 12, 15, 5, 4, 3, 2, 13, 15, 5, 6, 4, 2, 14, 12, 3, 2, 2,
	2, 14, 13, 3, 2, 2, 2, 15, 3, 3, 2, 2, 2, 16, 17, 7, 3, 2, 2, 17, 22, 5,
	10, 6, 2, 18, 19, 7, 4, 2, 2, 19, 21, 5, 10, 6, 2, 20, 18, 3, 2, 2, 2,
	21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 25, 3,
	2, 2, 2, 24, 22, 3, 2, 2, 2, 25, 26, 7, 5, 2, 2, 26, 30, 3, 2, 2, 2, 27,
	28, 7, 3, 2, 2, 28, 30, 7, 5, 2, 2, 29, 16, 3, 2, 2, 2, 29, 27, 3, 2, 2,
	2, 30, 5, 3, 2, 2, 2, 31, 32, 7, 6, 2, 2, 32, 37, 5, 8, 5, 2, 33, 34, 7,
	4, 2, 2, 34, 36, 5, 8, 5, 2, 35, 33, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37,
	35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2, 39, 37, 3, 2, 2,
	2, 40, 41, 7, 7, 2, 2, 41, 45, 3, 2, 2, 2, 42, 43, 7, 6, 2, 2, 43, 45,
	7, 7, 2, 2, 44, 31, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 7, 3, 2, 2, 2,
	46, 54, 7, 12, 2, 2, 47, 54, 7, 13, 2, 2, 48, 54, 5, 4, 3, 2, 49, 54, 5,
	6, 4, 2, 50, 54, 7, 8, 2, 2, 51, 54, 7, 9, 2, 2, 52, 54, 7, 10, 2, 2, 53,
	46, 3, 2, 2, 2, 53, 47, 3, 2, 2, 2, 53, 48, 3, 2, 2, 2, 53, 49, 3, 2, 2,
	2, 53, 50, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 9, 3,
	2, 2, 2, 55, 56, 7, 12, 2, 2, 56, 57, 7, 11, 2, 2, 57, 58, 5, 8, 5, 2,
	58, 11, 3, 2, 2, 2, 8, 14, 22, 29, 37, 44, 53,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "','", "'}'", "'['", "']'", "'true'", "'false'", "'null'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "STRING", "NUMBER", "WS",
}

var ruleNames = []string{
	"json", "object", "array", "value", "pair",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type JSONParser struct {
	*antlr.BaseParser
}

func NewJSONParser(input antlr.TokenStream) *JSONParser {
	this := new(JSONParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "JSON.g4"

	return this
}

// JSONParser tokens.
const (
	JSONParserEOF    = antlr.TokenEOF
	JSONParserT__0   = 1
	JSONParserT__1   = 2
	JSONParserT__2   = 3
	JSONParserT__3   = 4
	JSONParserT__4   = 5
	JSONParserT__5   = 6
	JSONParserT__6   = 7
	JSONParserT__7   = 8
	JSONParserT__8   = 9
	JSONParserSTRING = 10
	JSONParserNUMBER = 11
	JSONParserWS     = 12
)

// JSONParser rules.
const (
	JSONParserRULE_json   = 0
	JSONParserRULE_object = 1
	JSONParserRULE_array  = 2
	JSONParserRULE_value  = 3
	JSONParserRULE_pair   = 4
)

// IJsonContext is an interface to support dynamic dispatch.
type IJsonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJsonContext differentiates from other interfaces.
	IsJsonContext()
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

func (s *JsonContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *JsonContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitJson(s)
	}
}

func (p *JSONParser) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JSONParserRULE_json)

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

	p.SetState(12)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSONParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.Object()
		}

	case JSONParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) CopyFrom(ctx *ObjectContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlankObjectContext struct {
	*ObjectContext
}

func NewBlankObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlankObjectContext {
	var p = new(BlankObjectContext)

	p.ObjectContext = NewEmptyObjectContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectContext))

	return p
}

func (s *BlankObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterBlankObject(s)
	}
}

func (s *BlankObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitBlankObject(s)
	}
}

type AnObjectContext struct {
	*ObjectContext
}

func NewAnObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnObjectContext {
	var p = new(AnObjectContext)

	p.ObjectContext = NewEmptyObjectContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ObjectContext))

	return p
}

func (s *AnObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnObjectContext) AllPair() []IPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPairContext)(nil)).Elem())
	var tst = make([]IPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPairContext)
		}
	}

	return tst
}

func (s *AnObjectContext) Pair(i int) IPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPairContext)
}

func (s *AnObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAnObject(s)
	}
}

func (s *AnObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAnObject(s)
	}
}

func (p *JSONParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JSONParserRULE_object)
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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAnObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Match(JSONParserT__0)
		}
		{
			p.SetState(15)
			p.Pair()
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JSONParserT__1 {
			{
				p.SetState(16)
				p.Match(JSONParserT__1)
			}
			{
				p.SetState(17)
				p.Pair()
			}

			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(23)
			p.Match(JSONParserT__2)
		}

	case 2:
		localctx = NewBlankObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(JSONParserT__0)
		}
		{
			p.SetState(26)
			p.Match(JSONParserT__2)
		}

	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) CopyFrom(ctx *ArrayContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayOfValuesContext struct {
	*ArrayContext
}

func NewArrayOfValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayOfValuesContext {
	var p = new(ArrayOfValuesContext)

	p.ArrayContext = NewEmptyArrayContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayContext))

	return p
}

func (s *ArrayOfValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayOfValuesContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayOfValuesContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayOfValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterArrayOfValues(s)
	}
}

func (s *ArrayOfValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitArrayOfValues(s)
	}
}

type BlankArrayContext struct {
	*ArrayContext
}

func NewBlankArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlankArrayContext {
	var p = new(BlankArrayContext)

	p.ArrayContext = NewEmptyArrayContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ArrayContext))

	return p
}

func (s *BlankArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterBlankArray(s)
	}
}

func (s *BlankArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitBlankArray(s)
	}
}

func (p *JSONParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JSONParserRULE_array)
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

	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewArrayOfValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(29)
			p.Match(JSONParserT__3)
		}
		{
			p.SetState(30)
			p.Value()
		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == JSONParserT__1 {
			{
				p.SetState(31)
				p.Match(JSONParserT__1)
			}
			{
				p.SetState(32)
				p.Value()
			}

			p.SetState(37)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(38)
			p.Match(JSONParserT__4)
		}

	case 2:
		localctx = NewBlankArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(40)
			p.Match(JSONParserT__3)
		}
		{
			p.SetState(41)
			p.Match(JSONParserT__4)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ObjectValueContext struct {
	*ValueContext
}

func NewObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) Object() IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

type StringValueContext struct {
	*ValueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSONParserSTRING, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitStringValue(s)
	}
}

type ArrayValueContext struct {
	*ValueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

type AtomContext struct {
	*ValueContext
}

func NewAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomContext {
	var p = new(AtomContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(JSONParserNUMBER, 0)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *JSONParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JSONParserRULE_value)

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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case JSONParserSTRING:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(JSONParserSTRING)
		}

	case JSONParserNUMBER:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(JSONParserNUMBER)
		}

	case JSONParserT__0:
		localctx = NewObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(46)
			p.Object()
		}

	case JSONParserT__3:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(47)
			p.Array()
		}

	case JSONParserT__5:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(48)
			p.Match(JSONParserT__5)
		}

	case JSONParserT__6:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(49)
			p.Match(JSONParserT__6)
		}

	case JSONParserT__7:
		localctx = NewAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(50)
			p.Match(JSONParserT__7)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPairContext is an interface to support dynamic dispatch.
type IPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPairContext differentiates from other interfaces.
	IsPairContext()
}

type PairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPairContext() *PairContext {
	var p = new(PairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JSONParserRULE_pair
	return p
}

func (*PairContext) IsPairContext() {}

func NewPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PairContext {
	var p = new(PairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JSONParserRULE_pair

	return p
}

func (s *PairContext) GetParser() antlr.Parser { return s.parser }

func (s *PairContext) STRING() antlr.TerminalNode {
	return s.GetToken(JSONParserSTRING, 0)
}

func (s *PairContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *PairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.EnterPair(s)
	}
}

func (s *PairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JSONListener); ok {
		listenerT.ExitPair(s)
	}
}

func (p *JSONParser) Pair() (localctx IPairContext) {
	localctx = NewPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JSONParserRULE_pair)

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
		p.SetState(53)
		p.Match(JSONParserSTRING)
	}
	{
		p.SetState(54)
		p.Match(JSONParserT__8)
	}
	{
		p.SetState(55)
		p.Value()
	}

	return localctx
}
