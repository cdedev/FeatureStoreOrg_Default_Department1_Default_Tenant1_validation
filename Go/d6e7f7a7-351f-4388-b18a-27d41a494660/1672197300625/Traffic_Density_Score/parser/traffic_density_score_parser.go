// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672197300625/Traffic_Density_Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Traffic_Density_Score

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

type Traffic_Density_ScoreParser struct {
	*antlr.BaseParser
}

var traffic_density_scoreParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func traffic_density_scoreParserInit() {
	staticData := &traffic_density_scoreParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TRAFFIC_DENSITY_SCORE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "traffic_density_score",
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

// Traffic_Density_ScoreParserInit initializes any static state used to implement Traffic_Density_ScoreParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTraffic_Density_ScoreParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Traffic_Density_ScoreParserInit() {
	staticData := &traffic_density_scoreParserStaticData
	staticData.once.Do(traffic_density_scoreParserInit)
}

// NewTraffic_Density_ScoreParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTraffic_Density_ScoreParser(input antlr.TokenStream) *Traffic_Density_ScoreParser {
	Traffic_Density_ScoreParserInit()
	this := new(Traffic_Density_ScoreParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &traffic_density_scoreParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Traffic_Density_Score.g4"

	return this
}

// Traffic_Density_ScoreParser tokens.
const (
	Traffic_Density_ScoreParserEOF                   = antlr.TokenEOF
	Traffic_Density_ScoreParserTRAFFIC_DENSITY_SCORE = 1
	Traffic_Density_ScoreParserOWNKEY                = 2
	Traffic_Density_ScoreParserSPLITKEY              = 3
	Traffic_Density_ScoreParserWS                    = 4
)

// Traffic_Density_ScoreParser rules.
const (
	Traffic_Density_ScoreParserRULE_expression            = 0
	Traffic_Density_ScoreParserRULE_traffic_density_score = 1
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
	p.RuleIndex = Traffic_Density_ScoreParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Traffic_Density_ScoreParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Traffic_density_score() ITraffic_density_scoreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITraffic_density_scoreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITraffic_density_scoreContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Traffic_Density_ScoreParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Traffic_Density_ScoreListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Traffic_Density_ScoreListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Traffic_Density_ScoreParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Traffic_Density_ScoreParserRULE_expression)

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
		p.Traffic_density_score()
	}
	{
		p.SetState(5)
		p.Match(Traffic_Density_ScoreParserEOF)
	}

	return localctx
}

// ITraffic_density_scoreContext is an interface to support dynamic dispatch.
type ITraffic_density_scoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTraffic_density_scoreContext differentiates from other interfaces.
	IsTraffic_density_scoreContext()
}

type Traffic_density_scoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTraffic_density_scoreContext() *Traffic_density_scoreContext {
	var p = new(Traffic_density_scoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Traffic_Density_ScoreParserRULE_traffic_density_score
	return p
}

func (*Traffic_density_scoreContext) IsTraffic_density_scoreContext() {}

func NewTraffic_density_scoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Traffic_density_scoreContext {
	var p = new(Traffic_density_scoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Traffic_Density_ScoreParserRULE_traffic_density_score

	return p
}

func (s *Traffic_density_scoreContext) GetParser() antlr.Parser { return s.parser }

func (s *Traffic_density_scoreContext) TRAFFIC_DENSITY_SCORE() antlr.TerminalNode {
	return s.GetToken(Traffic_Density_ScoreParserTRAFFIC_DENSITY_SCORE, 0)
}

func (s *Traffic_density_scoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Traffic_density_scoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Traffic_density_scoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Traffic_Density_ScoreListener); ok {
		listenerT.EnterTraffic_density_score(s)
	}
}

func (s *Traffic_density_scoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Traffic_Density_ScoreListener); ok {
		listenerT.ExitTraffic_density_score(s)
	}
}

func (p *Traffic_Density_ScoreParser) Traffic_density_score() (localctx ITraffic_density_scoreContext) {
	this := p
	_ = this

	localctx = NewTraffic_density_scoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Traffic_Density_ScoreParserRULE_traffic_density_score)

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
		p.Match(Traffic_Density_ScoreParserTRAFFIC_DENSITY_SCORE)
	}

	return localctx
}
