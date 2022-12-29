// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672285317914/Detail.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Detail

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

type DetailParser struct {
	*antlr.BaseParser
}

var detailParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func detailParserInit() {
	staticData := &detailParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DETAIL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "detail",
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

// DetailParserInit initializes any static state used to implement DetailParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDetailParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DetailParserInit() {
	staticData := &detailParserStaticData
	staticData.once.Do(detailParserInit)
}

// NewDetailParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDetailParser(input antlr.TokenStream) *DetailParser {
	DetailParserInit()
	this := new(DetailParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &detailParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Detail.g4"

	return this
}

// DetailParser tokens.
const (
	DetailParserEOF      = antlr.TokenEOF
	DetailParserDETAIL   = 1
	DetailParserOWNKEY   = 2
	DetailParserSPLITKEY = 3
	DetailParserWS       = 4
)

// DetailParser rules.
const (
	DetailParserRULE_expression = 0
	DetailParserRULE_detail     = 1
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
	p.RuleIndex = DetailParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DetailParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Detail() IDetailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDetailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDetailContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DetailParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DetailListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DetailListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DetailParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DetailParserRULE_expression)

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
		p.Detail()
	}
	{
		p.SetState(5)
		p.Match(DetailParserEOF)
	}

	return localctx
}

// IDetailContext is an interface to support dynamic dispatch.
type IDetailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDetailContext differentiates from other interfaces.
	IsDetailContext()
}

type DetailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDetailContext() *DetailContext {
	var p = new(DetailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DetailParserRULE_detail
	return p
}

func (*DetailContext) IsDetailContext() {}

func NewDetailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DetailContext {
	var p = new(DetailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DetailParserRULE_detail

	return p
}

func (s *DetailContext) GetParser() antlr.Parser { return s.parser }

func (s *DetailContext) DETAIL() antlr.TerminalNode {
	return s.GetToken(DetailParserDETAIL, 0)
}

func (s *DetailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DetailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DetailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DetailListener); ok {
		listenerT.EnterDetail(s)
	}
}

func (s *DetailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DetailListener); ok {
		listenerT.ExitDetail(s)
	}
}

func (p *DetailParser) Detail() (localctx IDetailContext) {
	this := p
	_ = this

	localctx = NewDetailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DetailParserRULE_detail)

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
		p.Match(DetailParserDETAIL)
	}

	return localctx
}
