// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742760524/Earnings.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Earnings

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

type EarningsParser struct {
	*antlr.BaseParser
}

var earningsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func earningsParserInit() {
	staticData := &earningsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EARNINGS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "earnings",
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

// EarningsParserInit initializes any static state used to implement EarningsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEarningsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EarningsParserInit() {
	staticData := &earningsParserStaticData
	staticData.once.Do(earningsParserInit)
}

// NewEarningsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEarningsParser(input antlr.TokenStream) *EarningsParser {
	EarningsParserInit()
	this := new(EarningsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &earningsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Earnings.g4"

	return this
}

// EarningsParser tokens.
const (
	EarningsParserEOF      = antlr.TokenEOF
	EarningsParserEARNINGS = 1
	EarningsParserOWNKEY   = 2
	EarningsParserSPLITKEY = 3
	EarningsParserWS       = 4
)

// EarningsParser rules.
const (
	EarningsParserRULE_expression = 0
	EarningsParserRULE_earnings   = 1
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
	p.RuleIndex = EarningsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EarningsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Earnings() IEarningsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEarningsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEarningsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EarningsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EarningsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EarningsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EarningsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EarningsParserRULE_expression)

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
		p.Earnings()
	}
	{
		p.SetState(5)
		p.Match(EarningsParserEOF)
	}

	return localctx
}

// IEarningsContext is an interface to support dynamic dispatch.
type IEarningsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEarningsContext differentiates from other interfaces.
	IsEarningsContext()
}

type EarningsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEarningsContext() *EarningsContext {
	var p = new(EarningsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EarningsParserRULE_earnings
	return p
}

func (*EarningsContext) IsEarningsContext() {}

func NewEarningsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EarningsContext {
	var p = new(EarningsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EarningsParserRULE_earnings

	return p
}

func (s *EarningsContext) GetParser() antlr.Parser { return s.parser }

func (s *EarningsContext) EARNINGS() antlr.TerminalNode {
	return s.GetToken(EarningsParserEARNINGS, 0)
}

func (s *EarningsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EarningsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EarningsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EarningsListener); ok {
		listenerT.EnterEarnings(s)
	}
}

func (s *EarningsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EarningsListener); ok {
		listenerT.ExitEarnings(s)
	}
}

func (p *EarningsParser) Earnings() (localctx IEarningsContext) {
	this := p
	_ = this

	localctx = NewEarningsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EarningsParserRULE_earnings)

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
		p.Match(EarningsParserEARNINGS)
	}

	return localctx
}
