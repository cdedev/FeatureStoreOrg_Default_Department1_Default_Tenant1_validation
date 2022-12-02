// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982174882/Fuel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fuel

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

type FuelParser struct {
	*antlr.BaseParser
}

var fuelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fuelParserInit() {
	staticData := &fuelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FUEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "fuel",
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

// FuelParserInit initializes any static state used to implement FuelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFuelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FuelParserInit() {
	staticData := &fuelParserStaticData
	staticData.once.Do(fuelParserInit)
}

// NewFuelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFuelParser(input antlr.TokenStream) *FuelParser {
	FuelParserInit()
	this := new(FuelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fuelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Fuel.g4"

	return this
}

// FuelParser tokens.
const (
	FuelParserEOF      = antlr.TokenEOF
	FuelParserFUEL     = 1
	FuelParserOWNKEY   = 2
	FuelParserSPLITKEY = 3
	FuelParserWS       = 4
)

// FuelParser rules.
const (
	FuelParserRULE_expression = 0
	FuelParserRULE_fuel       = 1
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
	p.RuleIndex = FuelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Fuel() IFuelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FuelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FuelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuelParserRULE_expression)

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
		p.Fuel()
	}
	{
		p.SetState(5)
		p.Match(FuelParserEOF)
	}

	return localctx
}

// IFuelContext is an interface to support dynamic dispatch.
type IFuelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuelContext differentiates from other interfaces.
	IsFuelContext()
}

type FuelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuelContext() *FuelContext {
	var p = new(FuelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuelParserRULE_fuel
	return p
}

func (*FuelContext) IsFuelContext() {}

func NewFuelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuelContext {
	var p = new(FuelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuelParserRULE_fuel

	return p
}

func (s *FuelContext) GetParser() antlr.Parser { return s.parser }

func (s *FuelContext) FUEL() antlr.TerminalNode {
	return s.GetToken(FuelParserFUEL, 0)
}

func (s *FuelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelListener); ok {
		listenerT.EnterFuel(s)
	}
}

func (s *FuelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelListener); ok {
		listenerT.ExitFuel(s)
	}
}

func (p *FuelParser) Fuel() (localctx IFuelContext) {
	this := p
	_ = this

	localctx = NewFuelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FuelParserRULE_fuel)

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
		p.Match(FuelParserFUEL)
	}

	return localctx
}
