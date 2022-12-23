// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671768897139/FuelConsumption.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FuelConsumption

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

type FuelConsumptionParser struct {
	*antlr.BaseParser
}

var fuelconsumptionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fuelconsumptionParserInit() {
	staticData := &fuelconsumptionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FUELCONSUMPTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "fuelconsumption",
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

// FuelConsumptionParserInit initializes any static state used to implement FuelConsumptionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFuelConsumptionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FuelConsumptionParserInit() {
	staticData := &fuelconsumptionParserStaticData
	staticData.once.Do(fuelconsumptionParserInit)
}

// NewFuelConsumptionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFuelConsumptionParser(input antlr.TokenStream) *FuelConsumptionParser {
	FuelConsumptionParserInit()
	this := new(FuelConsumptionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fuelconsumptionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FuelConsumption.g4"

	return this
}

// FuelConsumptionParser tokens.
const (
	FuelConsumptionParserEOF             = antlr.TokenEOF
	FuelConsumptionParserFUELCONSUMPTION = 1
	FuelConsumptionParserOWNKEY          = 2
	FuelConsumptionParserSPLITKEY        = 3
	FuelConsumptionParserWS              = 4
)

// FuelConsumptionParser rules.
const (
	FuelConsumptionParserRULE_expression      = 0
	FuelConsumptionParserRULE_fuelconsumption = 1
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
	p.RuleIndex = FuelConsumptionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuelConsumptionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Fuelconsumption() IFuelconsumptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuelconsumptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuelconsumptionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FuelConsumptionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelConsumptionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelConsumptionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FuelConsumptionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuelConsumptionParserRULE_expression)

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
		p.Fuelconsumption()
	}
	{
		p.SetState(5)
		p.Match(FuelConsumptionParserEOF)
	}

	return localctx
}

// IFuelconsumptionContext is an interface to support dynamic dispatch.
type IFuelconsumptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuelconsumptionContext differentiates from other interfaces.
	IsFuelconsumptionContext()
}

type FuelconsumptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuelconsumptionContext() *FuelconsumptionContext {
	var p = new(FuelconsumptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuelConsumptionParserRULE_fuelconsumption
	return p
}

func (*FuelconsumptionContext) IsFuelconsumptionContext() {}

func NewFuelconsumptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuelconsumptionContext {
	var p = new(FuelconsumptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuelConsumptionParserRULE_fuelconsumption

	return p
}

func (s *FuelconsumptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuelconsumptionContext) FUELCONSUMPTION() antlr.TerminalNode {
	return s.GetToken(FuelConsumptionParserFUELCONSUMPTION, 0)
}

func (s *FuelconsumptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuelconsumptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuelconsumptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelConsumptionListener); ok {
		listenerT.EnterFuelconsumption(s)
	}
}

func (s *FuelconsumptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelConsumptionListener); ok {
		listenerT.ExitFuelconsumption(s)
	}
}

func (p *FuelConsumptionParser) Fuelconsumption() (localctx IFuelconsumptionContext) {
	this := p
	_ = this

	localctx = NewFuelconsumptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FuelConsumptionParserRULE_fuelconsumption)

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
		p.Match(FuelConsumptionParserFUELCONSUMPTION)
	}

	return localctx
}
