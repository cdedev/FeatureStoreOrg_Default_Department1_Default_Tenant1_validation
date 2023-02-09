// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675921656060/Risk.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Risk

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

type RiskParser struct {
	*antlr.BaseParser
}

var riskParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func riskParserInit() {
	staticData := &riskParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RISK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "risk",
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

// RiskParserInit initializes any static state used to implement RiskParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRiskParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RiskParserInit() {
	staticData := &riskParserStaticData
	staticData.once.Do(riskParserInit)
}

// NewRiskParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRiskParser(input antlr.TokenStream) *RiskParser {
	RiskParserInit()
	this := new(RiskParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &riskParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Risk.g4"

	return this
}

// RiskParser tokens.
const (
	RiskParserEOF      = antlr.TokenEOF
	RiskParserRISK     = 1
	RiskParserOWNKEY   = 2
	RiskParserSPLITKEY = 3
	RiskParserWS       = 4
)

// RiskParser rules.
const (
	RiskParserRULE_expression = 0
	RiskParserRULE_risk       = 1
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
	p.RuleIndex = RiskParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiskParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Risk() IRiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRiskContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RiskParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiskListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiskListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RiskParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RiskParserRULE_expression)

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
		p.Risk()
	}
	{
		p.SetState(5)
		p.Match(RiskParserEOF)
	}

	return localctx
}

// IRiskContext is an interface to support dynamic dispatch.
type IRiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRiskContext differentiates from other interfaces.
	IsRiskContext()
}

type RiskContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRiskContext() *RiskContext {
	var p = new(RiskContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RiskParserRULE_risk
	return p
}

func (*RiskContext) IsRiskContext() {}

func NewRiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RiskContext {
	var p = new(RiskContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RiskParserRULE_risk

	return p
}

func (s *RiskContext) GetParser() antlr.Parser { return s.parser }

func (s *RiskContext) RISK() antlr.TerminalNode {
	return s.GetToken(RiskParserRISK, 0)
}

func (s *RiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RiskContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiskListener); ok {
		listenerT.EnterRisk(s)
	}
}

func (s *RiskContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RiskListener); ok {
		listenerT.ExitRisk(s)
	}
}

func (p *RiskParser) Risk() (localctx IRiskContext) {
	this := p
	_ = this

	localctx = NewRiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RiskParserRULE_risk)

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
		p.Match(RiskParserRISK)
	}

	return localctx
}
