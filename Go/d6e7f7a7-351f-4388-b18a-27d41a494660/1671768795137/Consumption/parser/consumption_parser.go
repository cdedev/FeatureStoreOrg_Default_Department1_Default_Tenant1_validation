// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671768795137/Consumption.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Consumption

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

type ConsumptionParser struct {
	*antlr.BaseParser
}

var consumptionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func consumptionParserInit() {
	staticData := &consumptionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONSUMPTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "consumption",
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

// ConsumptionParserInit initializes any static state used to implement ConsumptionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConsumptionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConsumptionParserInit() {
	staticData := &consumptionParserStaticData
	staticData.once.Do(consumptionParserInit)
}

// NewConsumptionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConsumptionParser(input antlr.TokenStream) *ConsumptionParser {
	ConsumptionParserInit()
	this := new(ConsumptionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &consumptionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Consumption.g4"

	return this
}

// ConsumptionParser tokens.
const (
	ConsumptionParserEOF         = antlr.TokenEOF
	ConsumptionParserCONSUMPTION = 1
	ConsumptionParserOWNKEY      = 2
	ConsumptionParserSPLITKEY    = 3
	ConsumptionParserWS          = 4
)

// ConsumptionParser rules.
const (
	ConsumptionParserRULE_expression  = 0
	ConsumptionParserRULE_consumption = 1
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
	p.RuleIndex = ConsumptionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumptionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Consumption() IConsumptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConsumptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConsumptionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConsumptionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConsumptionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConsumptionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConsumptionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConsumptionParserRULE_expression)

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
		p.Consumption()
	}
	{
		p.SetState(5)
		p.Match(ConsumptionParserEOF)
	}

	return localctx
}

// IConsumptionContext is an interface to support dynamic dispatch.
type IConsumptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConsumptionContext differentiates from other interfaces.
	IsConsumptionContext()
}

type ConsumptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConsumptionContext() *ConsumptionContext {
	var p = new(ConsumptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConsumptionParserRULE_consumption
	return p
}

func (*ConsumptionContext) IsConsumptionContext() {}

func NewConsumptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConsumptionContext {
	var p = new(ConsumptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConsumptionParserRULE_consumption

	return p
}

func (s *ConsumptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConsumptionContext) CONSUMPTION() antlr.TerminalNode {
	return s.GetToken(ConsumptionParserCONSUMPTION, 0)
}

func (s *ConsumptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConsumptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConsumptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConsumptionListener); ok {
		listenerT.EnterConsumption(s)
	}
}

func (s *ConsumptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConsumptionListener); ok {
		listenerT.ExitConsumption(s)
	}
}

func (p *ConsumptionParser) Consumption() (localctx IConsumptionContext) {
	this := p
	_ = this

	localctx = NewConsumptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConsumptionParserRULE_consumption)

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
		p.Match(ConsumptionParserCONSUMPTION)
	}

	return localctx
}
