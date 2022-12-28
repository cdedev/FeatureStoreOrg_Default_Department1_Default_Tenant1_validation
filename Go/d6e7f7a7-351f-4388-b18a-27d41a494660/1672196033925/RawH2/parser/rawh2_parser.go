// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196033925/RawH2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RawH2

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RawH2Parser struct {
	*antlr.BaseParser
}

var rawh2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rawh2ParserInit() {
	staticData := &rawh2ParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RAWH2", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rawh2",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 4, 10, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		0, 0, 2, 0, 2, 0, 0, 7, 0, 4, 1, 0, 0, 0, 2, 7, 1, 0, 0, 0, 4, 5, 3, 2,
		1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 5, 1, 0, 0, 8, 3, 1, 0,
		0, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// RawH2ParserInit initializes any static state used to implement RawH2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRawH2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RawH2ParserInit() {
	staticData := &rawh2ParserStaticData
	staticData.once.Do(rawh2ParserInit)
}

// NewRawH2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewRawH2Parser(input antlr.TokenStream) *RawH2Parser {
	RawH2ParserInit()
	this := new(RawH2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rawh2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RawH2.g4"

	return this
}

// RawH2Parser tokens.
const (
	RawH2ParserEOF      = antlr.TokenEOF
	RawH2ParserRAWH2    = 1
	RawH2ParserOWNKEY   = 2
	RawH2ParserSPLITKEY = 3
	RawH2ParserWS       = 4
)

// RawH2Parser rules.
const (
	RawH2ParserRULE_expression = 0
	RawH2ParserRULE_rawh2      = 1
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RawH2ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RawH2ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rawh2() IRawh2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawh2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawh2Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RawH2ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawH2Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawH2Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RawH2Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RawH2ParserRULE_expression)

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
		p.SetState(4)
		p.Rawh2()
	}
	{
		p.SetState(5)
		p.Match(RawH2ParserEOF)
	}

	return localctx
}

// IRawh2Context is an interface to support dynamic dispatch.
type IRawh2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawh2Context differentiates from other interfaces.
	IsRawh2Context()
}

type Rawh2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawh2Context() *Rawh2Context {
	var p = new(Rawh2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RawH2ParserRULE_rawh2
	return p
}

func (*Rawh2Context) IsRawh2Context() {}

func NewRawh2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rawh2Context {
	var p = new(Rawh2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RawH2ParserRULE_rawh2

	return p
}

func (s *Rawh2Context) GetParser() antlr.Parser { return s.parser }

func (s *Rawh2Context) RAWH2() antlr.TerminalNode {
	return s.GetToken(RawH2ParserRAWH2, 0)
}

func (s *Rawh2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rawh2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rawh2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawH2Listener); ok {
		listenerT.EnterRawh2(s)
	}
}

func (s *Rawh2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawH2Listener); ok {
		listenerT.ExitRawh2(s)
	}
}

func (p *RawH2Parser) Rawh2() (localctx IRawh2Context) {
	this := p
	_ = this

	localctx = NewRawh2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RawH2ParserRULE_rawh2)

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
		p.SetState(7)
		p.Match(RawH2ParserRAWH2)
	}

	return localctx
}
