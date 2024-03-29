// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675150990145/Indicator.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Indicator

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

type IndicatorParser struct {
	*antlr.BaseParser
}

var indicatorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func indicatorParserInit() {
	staticData := &indicatorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INDICATOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "indicator",
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

// IndicatorParserInit initializes any static state used to implement IndicatorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIndicatorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IndicatorParserInit() {
	staticData := &indicatorParserStaticData
	staticData.once.Do(indicatorParserInit)
}

// NewIndicatorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIndicatorParser(input antlr.TokenStream) *IndicatorParser {
	IndicatorParserInit()
	this := new(IndicatorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &indicatorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Indicator.g4"

	return this
}

// IndicatorParser tokens.
const (
	IndicatorParserEOF       = antlr.TokenEOF
	IndicatorParserINDICATOR = 1
	IndicatorParserOWNKEY    = 2
	IndicatorParserSPLITKEY  = 3
	IndicatorParserWS        = 4
)

// IndicatorParser rules.
const (
	IndicatorParserRULE_expression = 0
	IndicatorParserRULE_indicator  = 1
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
	p.RuleIndex = IndicatorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IndicatorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Indicator() IIndicatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndicatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndicatorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IndicatorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndicatorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndicatorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IndicatorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IndicatorParserRULE_expression)

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
		p.Indicator()
	}
	{
		p.SetState(5)
		p.Match(IndicatorParserEOF)
	}

	return localctx
}

// IIndicatorContext is an interface to support dynamic dispatch.
type IIndicatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndicatorContext differentiates from other interfaces.
	IsIndicatorContext()
}

type IndicatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndicatorContext() *IndicatorContext {
	var p = new(IndicatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IndicatorParserRULE_indicator
	return p
}

func (*IndicatorContext) IsIndicatorContext() {}

func NewIndicatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndicatorContext {
	var p = new(IndicatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IndicatorParserRULE_indicator

	return p
}

func (s *IndicatorContext) GetParser() antlr.Parser { return s.parser }

func (s *IndicatorContext) INDICATOR() antlr.TerminalNode {
	return s.GetToken(IndicatorParserINDICATOR, 0)
}

func (s *IndicatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndicatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndicatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndicatorListener); ok {
		listenerT.EnterIndicator(s)
	}
}

func (s *IndicatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndicatorListener); ok {
		listenerT.ExitIndicator(s)
	}
}

func (p *IndicatorParser) Indicator() (localctx IIndicatorContext) {
	this := p
	_ = this

	localctx = NewIndicatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IndicatorParserRULE_indicator)

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
		p.Match(IndicatorParserINDICATOR)
	}

	return localctx
}
