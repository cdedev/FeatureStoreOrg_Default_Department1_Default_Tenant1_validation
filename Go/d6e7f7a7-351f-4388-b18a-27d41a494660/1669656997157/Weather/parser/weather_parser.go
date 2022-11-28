// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656997157/Weather.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weather

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

type WeatherParser struct {
	*antlr.BaseParser
}

var weatherParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func weatherParserInit() {
	staticData := &weatherParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WEATHER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "weather",
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

// WeatherParserInit initializes any static state used to implement WeatherParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWeatherParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeatherParserInit() {
	staticData := &weatherParserStaticData
	staticData.once.Do(weatherParserInit)
}

// NewWeatherParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWeatherParser(input antlr.TokenStream) *WeatherParser {
	WeatherParserInit()
	this := new(WeatherParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &weatherParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Weather.g4"

	return this
}

// WeatherParser tokens.
const (
	WeatherParserEOF      = antlr.TokenEOF
	WeatherParserWEATHER  = 1
	WeatherParserOWNKEY   = 2
	WeatherParserSPLITKEY = 3
	WeatherParserWS       = 4
)

// WeatherParser rules.
const (
	WeatherParserRULE_expression = 0
	WeatherParserRULE_weather    = 1
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
	p.RuleIndex = WeatherParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeatherParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Weather() IWeatherContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWeatherContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWeatherContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WeatherParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WeatherParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WeatherParserRULE_expression)

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
		p.Weather()
	}
	{
		p.SetState(5)
		p.Match(WeatherParserEOF)
	}

	return localctx
}

// IWeatherContext is an interface to support dynamic dispatch.
type IWeatherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeatherContext differentiates from other interfaces.
	IsWeatherContext()
}

type WeatherContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeatherContext() *WeatherContext {
	var p = new(WeatherContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WeatherParserRULE_weather
	return p
}

func (*WeatherContext) IsWeatherContext() {}

func NewWeatherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeatherContext {
	var p = new(WeatherContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeatherParserRULE_weather

	return p
}

func (s *WeatherContext) GetParser() antlr.Parser { return s.parser }

func (s *WeatherContext) WEATHER() antlr.TerminalNode {
	return s.GetToken(WeatherParserWEATHER, 0)
}

func (s *WeatherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeatherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeatherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherListener); ok {
		listenerT.EnterWeather(s)
	}
}

func (s *WeatherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherListener); ok {
		listenerT.ExitWeather(s)
	}
}

func (p *WeatherParser) Weather() (localctx IWeatherContext) {
	this := p
	_ = this

	localctx = NewWeatherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WeatherParserRULE_weather)

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
		p.Match(WeatherParserWEATHER)
	}

	return localctx
}
