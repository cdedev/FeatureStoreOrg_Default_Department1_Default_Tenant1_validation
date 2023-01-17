// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673924843070/Airborne.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airborne

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

type AirborneParser struct {
	*antlr.BaseParser
}

var airborneParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func airborneParserInit() {
	staticData := &airborneParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AIRBORNE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "airborne",
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

// AirborneParserInit initializes any static state used to implement AirborneParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAirborneParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AirborneParserInit() {
	staticData := &airborneParserStaticData
	staticData.once.Do(airborneParserInit)
}

// NewAirborneParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAirborneParser(input antlr.TokenStream) *AirborneParser {
	AirborneParserInit()
	this := new(AirborneParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &airborneParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Airborne.g4"

	return this
}

// AirborneParser tokens.
const (
	AirborneParserEOF      = antlr.TokenEOF
	AirborneParserAIRBORNE = 1
	AirborneParserOWNKEY   = 2
	AirborneParserSPLITKEY = 3
	AirborneParserWS       = 4
)

// AirborneParser rules.
const (
	AirborneParserRULE_expression = 0
	AirborneParserRULE_airborne   = 1
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
	p.RuleIndex = AirborneParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AirborneParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Airborne() IAirborneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAirborneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAirborneContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AirborneParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirborneListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirborneListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AirborneParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AirborneParserRULE_expression)

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
		p.Airborne()
	}
	{
		p.SetState(5)
		p.Match(AirborneParserEOF)
	}

	return localctx
}

// IAirborneContext is an interface to support dynamic dispatch.
type IAirborneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAirborneContext differentiates from other interfaces.
	IsAirborneContext()
}

type AirborneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAirborneContext() *AirborneContext {
	var p = new(AirborneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AirborneParserRULE_airborne
	return p
}

func (*AirborneContext) IsAirborneContext() {}

func NewAirborneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AirborneContext {
	var p = new(AirborneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AirborneParserRULE_airborne

	return p
}

func (s *AirborneContext) GetParser() antlr.Parser { return s.parser }

func (s *AirborneContext) AIRBORNE() antlr.TerminalNode {
	return s.GetToken(AirborneParserAIRBORNE, 0)
}

func (s *AirborneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AirborneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AirborneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirborneListener); ok {
		listenerT.EnterAirborne(s)
	}
}

func (s *AirborneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirborneListener); ok {
		listenerT.ExitAirborne(s)
	}
}

func (p *AirborneParser) Airborne() (localctx IAirborneContext) {
	this := p
	_ = this

	localctx = NewAirborneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AirborneParserRULE_airborne)

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
		p.Match(AirborneParserAIRBORNE)
	}

	return localctx
}
