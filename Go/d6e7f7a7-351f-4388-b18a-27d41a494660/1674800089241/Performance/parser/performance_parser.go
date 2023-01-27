// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674800089241/Performance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Performance

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

type PerformanceParser struct {
	*antlr.BaseParser
}

var performanceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func performanceParserInit() {
	staticData := &performanceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PERFORMANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "performance",
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

// PerformanceParserInit initializes any static state used to implement PerformanceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPerformanceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PerformanceParserInit() {
	staticData := &performanceParserStaticData
	staticData.once.Do(performanceParserInit)
}

// NewPerformanceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPerformanceParser(input antlr.TokenStream) *PerformanceParser {
	PerformanceParserInit()
	this := new(PerformanceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &performanceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Performance.g4"

	return this
}

// PerformanceParser tokens.
const (
	PerformanceParserEOF         = antlr.TokenEOF
	PerformanceParserPERFORMANCE = 1
	PerformanceParserOWNKEY      = 2
	PerformanceParserSPLITKEY    = 3
	PerformanceParserWS          = 4
)

// PerformanceParser rules.
const (
	PerformanceParserRULE_expression  = 0
	PerformanceParserRULE_performance = 1
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
	p.RuleIndex = PerformanceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PerformanceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Performance() IPerformanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPerformanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPerformanceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PerformanceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerformanceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerformanceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PerformanceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PerformanceParserRULE_expression)

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
		p.Performance()
	}
	{
		p.SetState(5)
		p.Match(PerformanceParserEOF)
	}

	return localctx
}

// IPerformanceContext is an interface to support dynamic dispatch.
type IPerformanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPerformanceContext differentiates from other interfaces.
	IsPerformanceContext()
}

type PerformanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPerformanceContext() *PerformanceContext {
	var p = new(PerformanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PerformanceParserRULE_performance
	return p
}

func (*PerformanceContext) IsPerformanceContext() {}

func NewPerformanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PerformanceContext {
	var p = new(PerformanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PerformanceParserRULE_performance

	return p
}

func (s *PerformanceContext) GetParser() antlr.Parser { return s.parser }

func (s *PerformanceContext) PERFORMANCE() antlr.TerminalNode {
	return s.GetToken(PerformanceParserPERFORMANCE, 0)
}

func (s *PerformanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PerformanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PerformanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerformanceListener); ok {
		listenerT.EnterPerformance(s)
	}
}

func (s *PerformanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerformanceListener); ok {
		listenerT.ExitPerformance(s)
	}
}

func (p *PerformanceParser) Performance() (localctx IPerformanceContext) {
	this := p
	_ = this

	localctx = NewPerformanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PerformanceParserRULE_performance)

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
		p.Match(PerformanceParserPERFORMANCE)
	}

	return localctx
}
