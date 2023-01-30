// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072586303/AnnualGrowthRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AnnualGrowthRate

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

type AnnualGrowthRateParser struct {
	*antlr.BaseParser
}

var annualgrowthrateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func annualgrowthrateParserInit() {
	staticData := &annualgrowthrateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ANNUALGROWTHRATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "annualgrowthrate",
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

// AnnualGrowthRateParserInit initializes any static state used to implement AnnualGrowthRateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAnnualGrowthRateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AnnualGrowthRateParserInit() {
	staticData := &annualgrowthrateParserStaticData
	staticData.once.Do(annualgrowthrateParserInit)
}

// NewAnnualGrowthRateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAnnualGrowthRateParser(input antlr.TokenStream) *AnnualGrowthRateParser {
	AnnualGrowthRateParserInit()
	this := new(AnnualGrowthRateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &annualgrowthrateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AnnualGrowthRate.g4"

	return this
}

// AnnualGrowthRateParser tokens.
const (
	AnnualGrowthRateParserEOF              = antlr.TokenEOF
	AnnualGrowthRateParserANNUALGROWTHRATE = 1
	AnnualGrowthRateParserOWNKEY           = 2
	AnnualGrowthRateParserSPLITKEY         = 3
	AnnualGrowthRateParserWS               = 4
)

// AnnualGrowthRateParser rules.
const (
	AnnualGrowthRateParserRULE_expression       = 0
	AnnualGrowthRateParserRULE_annualgrowthrate = 1
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
	p.RuleIndex = AnnualGrowthRateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnualGrowthRateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Annualgrowthrate() IAnnualgrowthrateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnnualgrowthrateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnnualgrowthrateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AnnualGrowthRateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnualGrowthRateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnualGrowthRateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AnnualGrowthRateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AnnualGrowthRateParserRULE_expression)

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
		p.Annualgrowthrate()
	}
	{
		p.SetState(5)
		p.Match(AnnualGrowthRateParserEOF)
	}

	return localctx
}

// IAnnualgrowthrateContext is an interface to support dynamic dispatch.
type IAnnualgrowthrateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnualgrowthrateContext differentiates from other interfaces.
	IsAnnualgrowthrateContext()
}

type AnnualgrowthrateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnualgrowthrateContext() *AnnualgrowthrateContext {
	var p = new(AnnualgrowthrateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnualGrowthRateParserRULE_annualgrowthrate
	return p
}

func (*AnnualgrowthrateContext) IsAnnualgrowthrateContext() {}

func NewAnnualgrowthrateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnualgrowthrateContext {
	var p = new(AnnualgrowthrateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnualGrowthRateParserRULE_annualgrowthrate

	return p
}

func (s *AnnualgrowthrateContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnualgrowthrateContext) ANNUALGROWTHRATE() antlr.TerminalNode {
	return s.GetToken(AnnualGrowthRateParserANNUALGROWTHRATE, 0)
}

func (s *AnnualgrowthrateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnualgrowthrateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnualgrowthrateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnualGrowthRateListener); ok {
		listenerT.EnterAnnualgrowthrate(s)
	}
}

func (s *AnnualgrowthrateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnualGrowthRateListener); ok {
		listenerT.ExitAnnualgrowthrate(s)
	}
}

func (p *AnnualGrowthRateParser) Annualgrowthrate() (localctx IAnnualgrowthrateContext) {
	this := p
	_ = this

	localctx = NewAnnualgrowthrateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AnnualGrowthRateParserRULE_annualgrowthrate)

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
		p.Match(AnnualGrowthRateParserANNUALGROWTHRATE)
	}

	return localctx
}
