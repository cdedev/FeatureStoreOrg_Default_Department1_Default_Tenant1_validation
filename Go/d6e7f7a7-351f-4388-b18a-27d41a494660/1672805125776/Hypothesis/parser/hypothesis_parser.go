// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672805125776/Hypothesis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hypothesis

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

type HypothesisParser struct {
	*antlr.BaseParser
}

var hypothesisParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hypothesisParserInit() {
	staticData := &hypothesisParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HYPOTHESIS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hypothesis",
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

// HypothesisParserInit initializes any static state used to implement HypothesisParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHypothesisParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HypothesisParserInit() {
	staticData := &hypothesisParserStaticData
	staticData.once.Do(hypothesisParserInit)
}

// NewHypothesisParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHypothesisParser(input antlr.TokenStream) *HypothesisParser {
	HypothesisParserInit()
	this := new(HypothesisParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hypothesisParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Hypothesis.g4"

	return this
}

// HypothesisParser tokens.
const (
	HypothesisParserEOF        = antlr.TokenEOF
	HypothesisParserHYPOTHESIS = 1
	HypothesisParserOWNKEY     = 2
	HypothesisParserSPLITKEY   = 3
	HypothesisParserWS         = 4
)

// HypothesisParser rules.
const (
	HypothesisParserRULE_expression = 0
	HypothesisParserRULE_hypothesis = 1
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
	p.RuleIndex = HypothesisParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HypothesisParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hypothesis() IHypothesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHypothesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHypothesisContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HypothesisParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypothesisListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypothesisListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HypothesisParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HypothesisParserRULE_expression)

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
		p.Hypothesis()
	}
	{
		p.SetState(5)
		p.Match(HypothesisParserEOF)
	}

	return localctx
}

// IHypothesisContext is an interface to support dynamic dispatch.
type IHypothesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHypothesisContext differentiates from other interfaces.
	IsHypothesisContext()
}

type HypothesisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHypothesisContext() *HypothesisContext {
	var p = new(HypothesisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HypothesisParserRULE_hypothesis
	return p
}

func (*HypothesisContext) IsHypothesisContext() {}

func NewHypothesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HypothesisContext {
	var p = new(HypothesisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HypothesisParserRULE_hypothesis

	return p
}

func (s *HypothesisContext) GetParser() antlr.Parser { return s.parser }

func (s *HypothesisContext) HYPOTHESIS() antlr.TerminalNode {
	return s.GetToken(HypothesisParserHYPOTHESIS, 0)
}

func (s *HypothesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HypothesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HypothesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypothesisListener); ok {
		listenerT.EnterHypothesis(s)
	}
}

func (s *HypothesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypothesisListener); ok {
		listenerT.ExitHypothesis(s)
	}
}

func (p *HypothesisParser) Hypothesis() (localctx IHypothesisContext) {
	this := p
	_ = this

	localctx = NewHypothesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HypothesisParserRULE_hypothesis)

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
		p.Match(HypothesisParserHYPOTHESIS)
	}

	return localctx
}
