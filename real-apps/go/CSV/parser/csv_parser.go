// Code generated from CSV.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // CSV

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 52, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 6,
	2, 15, 10, 2, 13, 2, 14, 2, 16, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 24,
	10, 4, 12, 4, 14, 4, 27, 11, 4, 3, 4, 5, 4, 30, 10, 4, 3, 4, 3, 4, 3, 5,
	7, 5, 35, 10, 5, 12, 5, 14, 5, 38, 11, 5, 3, 5, 3, 5, 7, 5, 42, 10, 5,
	12, 5, 14, 5, 45, 11, 5, 3, 6, 3, 6, 3, 6, 5, 6, 50, 10, 6, 3, 6, 2, 2,
	7, 2, 4, 6, 8, 10, 2, 2, 2, 53, 2, 12, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6,
	20, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 49, 3, 2, 2, 2, 12, 14, 5, 4, 3,
	2, 13, 15, 5, 6, 4, 2, 14, 13, 3, 2, 2, 2, 15, 16, 3, 2, 2, 2, 16, 14,
	3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 3, 3, 2, 2, 2, 18, 19, 5, 6, 4, 2,
	19, 5, 3, 2, 2, 2, 20, 25, 5, 8, 5, 2, 21, 22, 7, 7, 2, 2, 22, 24, 5, 8,
	5, 2, 23, 21, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26,
	3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 30, 7, 3, 2, 2,
	29, 28, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 7,
	4, 2, 2, 32, 7, 3, 2, 2, 2, 33, 35, 7, 8, 2, 2, 34, 33, 3, 2, 2, 2, 35,
	38, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 39, 3, 2, 2,
	2, 38, 36, 3, 2, 2, 2, 39, 43, 5, 10, 6, 2, 40, 42, 7, 8, 2, 2, 41, 40,
	3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2,
	44, 9, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 50, 7, 5, 2, 2, 47, 50, 7, 6,
	2, 2, 48, 50, 3, 2, 2, 2, 49, 46, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 48,
	3, 2, 2, 2, 50, 11, 3, 2, 2, 2, 8, 16, 25, 29, 36, 43, 49,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'\r'", "'\n'", "", "", "','",
}
var symbolicNames = []string{
	"", "", "", "TEXT", "STRING", "COMMA", "WS",
}

var ruleNames = []string{
	"file", "hdr", "row", "fieldWWS", "field",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CSVParser struct {
	*antlr.BaseParser
}

func NewCSVParser(input antlr.TokenStream) *CSVParser {
	this := new(CSVParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CSV.g4"

	return this
}

// CSVParser tokens.
const (
	CSVParserEOF    = antlr.TokenEOF
	CSVParserT__0   = 1
	CSVParserT__1   = 2
	CSVParserTEXT   = 3
	CSVParserSTRING = 4
	CSVParserCOMMA  = 5
	CSVParserWS     = 6
)

// CSVParser rules.
const (
	CSVParserRULE_file     = 0
	CSVParserRULE_hdr      = 1
	CSVParserRULE_row      = 2
	CSVParserRULE_fieldWWS = 3
	CSVParserRULE_field    = 4
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) Hdr() IHdrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHdrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHdrContext)
}

func (s *FileContext) AllRow() []IRowContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRowContext)(nil)).Elem())
	var tst = make([]IRowContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRowContext)
		}
	}

	return tst
}

func (s *FileContext) Row(i int) IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *CSVParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CSVParserRULE_file)
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
		p.SetState(10)
		p.Hdr()
	}
	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CSVParserT__0)|(1<<CSVParserT__1)|(1<<CSVParserTEXT)|(1<<CSVParserSTRING)|(1<<CSVParserCOMMA)|(1<<CSVParserWS))) != 0) {
		{
			p.SetState(11)
			p.Row()
		}

		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHdrContext is an interface to support dynamic dispatch.
type IHdrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHdrContext differentiates from other interfaces.
	IsHdrContext()
}

type HdrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHdrContext() *HdrContext {
	var p = new(HdrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_hdr
	return p
}

func (*HdrContext) IsHdrContext() {}

func NewHdrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HdrContext {
	var p = new(HdrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_hdr

	return p
}

func (s *HdrContext) GetParser() antlr.Parser { return s.parser }

func (s *HdrContext) Row() IRowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRowContext)
}

func (s *HdrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HdrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HdrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterHdr(s)
	}
}

func (s *HdrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitHdr(s)
	}
}

func (p *CSVParser) Hdr() (localctx IHdrContext) {
	localctx = NewHdrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CSVParserRULE_hdr)

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
		p.SetState(16)
		p.Row()
	}

	return localctx
}

// IRowContext is an interface to support dynamic dispatch.
type IRowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRowContext differentiates from other interfaces.
	IsRowContext()
}

type RowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRowContext() *RowContext {
	var p = new(RowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_row
	return p
}

func (*RowContext) IsRowContext() {}

func NewRowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RowContext {
	var p = new(RowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_row

	return p
}

func (s *RowContext) GetParser() antlr.Parser { return s.parser }

func (s *RowContext) AllFieldWWS() []IFieldWWSContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldWWSContext)(nil)).Elem())
	var tst = make([]IFieldWWSContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldWWSContext)
		}
	}

	return tst
}

func (s *RowContext) FieldWWS(i int) IFieldWWSContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldWWSContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldWWSContext)
}

func (s *RowContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CSVParserCOMMA)
}

func (s *RowContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CSVParserCOMMA, i)
}

func (s *RowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterRow(s)
	}
}

func (s *RowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitRow(s)
	}
}

func (p *CSVParser) Row() (localctx IRowContext) {
	localctx = NewRowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CSVParserRULE_row)
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
		p.SetState(18)
		p.FieldWWS()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CSVParserCOMMA {
		{
			p.SetState(19)
			p.Match(CSVParserCOMMA)
		}
		{
			p.SetState(20)
			p.FieldWWS()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CSVParserT__0 {
		{
			p.SetState(26)
			p.Match(CSVParserT__0)
		}

	}
	{
		p.SetState(29)
		p.Match(CSVParserT__1)
	}

	return localctx
}

// IFieldWWSContext is an interface to support dynamic dispatch.
type IFieldWWSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldWWSContext differentiates from other interfaces.
	IsFieldWWSContext()
}

type FieldWWSContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldWWSContext() *FieldWWSContext {
	var p = new(FieldWWSContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_fieldWWS
	return p
}

func (*FieldWWSContext) IsFieldWWSContext() {}

func NewFieldWWSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldWWSContext {
	var p = new(FieldWWSContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_fieldWWS

	return p
}

func (s *FieldWWSContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldWWSContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *FieldWWSContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(CSVParserWS)
}

func (s *FieldWWSContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(CSVParserWS, i)
}

func (s *FieldWWSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldWWSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldWWSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterFieldWWS(s)
	}
}

func (s *FieldWWSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitFieldWWS(s)
	}
}

func (p *CSVParser) FieldWWS() (localctx IFieldWWSContext) {
	localctx = NewFieldWWSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CSVParserRULE_fieldWWS)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(31)
				p.Match(CSVParserWS)
			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}
	{
		p.SetState(37)
		p.Field()
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CSVParserWS {
		{
			p.SetState(38)
			p.Match(CSVParserWS)
		}

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CSVParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CSVParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) CopyFrom(ctx *FieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringContext struct {
	*FieldContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(CSVParserSTRING, 0)
}

func (s *StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterString(s)
	}
}

func (s *StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitString(s)
	}
}

type TextContext struct {
	*FieldContext
}

func NewTextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextContext {
	var p = new(TextContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(CSVParserTEXT, 0)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitText(s)
	}
}

type EmptyContext struct {
	*FieldContext
}

func NewEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyContext {
	var p = new(EmptyContext)

	p.FieldContext = NewEmptyFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldContext))

	return p
}

func (s *EmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.EnterEmpty(s)
	}
}

func (s *EmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CSVListener); ok {
		listenerT.ExitEmpty(s)
	}
}

func (p *CSVParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CSVParserRULE_field)

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

	p.SetState(47)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CSVParserTEXT:
		localctx = NewTextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(44)
			p.Match(CSVParserTEXT)
		}

	case CSVParserSTRING:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(45)
			p.Match(CSVParserSTRING)
		}

	case CSVParserT__0, CSVParserT__1, CSVParserCOMMA, CSVParserWS:
		localctx = NewEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
