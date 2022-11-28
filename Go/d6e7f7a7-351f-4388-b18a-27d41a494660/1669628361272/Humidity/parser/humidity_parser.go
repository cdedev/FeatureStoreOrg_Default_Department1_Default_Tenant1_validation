// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628361272/Humidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Humidity

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

type HumidityParser struct {
	*antlr.BaseParser
}

var humidityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func humidityParserInit() {
	staticData := &humidityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HUMIDITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "humidity",
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

// HumidityParserInit initializes any static state used to implement HumidityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHumidityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HumidityParserInit() {
	staticData := &humidityParserStaticData
	staticData.once.Do(humidityParserInit)
}

// NewHumidityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHumidityParser(input antlr.TokenStream) *HumidityParser {
	HumidityParserInit()
	this := new(HumidityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &humidityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Humidity.g4"

	return this
}

// HumidityParser tokens.
const (
	HumidityParserEOF      = antlr.TokenEOF
	HumidityParserHUMIDITY = 1
	HumidityParserOWNKEY   = 2
	HumidityParserSPLITKEY = 3
	HumidityParserWS       = 4
)

// HumidityParser rules.
const (
	HumidityParserRULE_expression = 0
	HumidityParserRULE_humidity   = 1
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
	p.RuleIndex = HumidityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HumidityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Humidity() IHumidityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHumidityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHumidityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HumidityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HumidityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HumidityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HumidityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HumidityParserRULE_expression)

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
		p.Humidity()
	}
	{
		p.SetState(5)
		p.Match(HumidityParserEOF)
	}

	return localctx
}

// IHumidityContext is an interface to support dynamic dispatch.
type IHumidityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHumidityContext differentiates from other interfaces.
	IsHumidityContext()
}

type HumidityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHumidityContext() *HumidityContext {
	var p = new(HumidityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HumidityParserRULE_humidity
	return p
}

func (*HumidityContext) IsHumidityContext() {}

func NewHumidityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HumidityContext {
	var p = new(HumidityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HumidityParserRULE_humidity

	return p
}

func (s *HumidityContext) GetParser() antlr.Parser { return s.parser }

func (s *HumidityContext) HUMIDITY() antlr.TerminalNode {
	return s.GetToken(HumidityParserHUMIDITY, 0)
}

func (s *HumidityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HumidityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HumidityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HumidityListener); ok {
		listenerT.EnterHumidity(s)
	}
}

func (s *HumidityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HumidityListener); ok {
		listenerT.ExitHumidity(s)
	}
}

func (p *HumidityParser) Humidity() (localctx IHumidityContext) {
	this := p
	_ = this

	localctx = NewHumidityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HumidityParserRULE_humidity)

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
		p.Match(HumidityParserHUMIDITY)
	}

	return localctx
}
