// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673670414326/WeatherType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherType

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

type WeatherTypeParser struct {
	*antlr.BaseParser
}

var weathertypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func weathertypeParserInit() {
	staticData := &weathertypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WEATHERTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "weathertype",
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

// WeatherTypeParserInit initializes any static state used to implement WeatherTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWeatherTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeatherTypeParserInit() {
	staticData := &weathertypeParserStaticData
	staticData.once.Do(weathertypeParserInit)
}

// NewWeatherTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWeatherTypeParser(input antlr.TokenStream) *WeatherTypeParser {
	WeatherTypeParserInit()
	this := new(WeatherTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &weathertypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "WeatherType.g4"

	return this
}

// WeatherTypeParser tokens.
const (
	WeatherTypeParserEOF         = antlr.TokenEOF
	WeatherTypeParserWEATHERTYPE = 1
	WeatherTypeParserOWNKEY      = 2
	WeatherTypeParserSPLITKEY    = 3
	WeatherTypeParserWS          = 4
)

// WeatherTypeParser rules.
const (
	WeatherTypeParserRULE_expression  = 0
	WeatherTypeParserRULE_weathertype = 1
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
	p.RuleIndex = WeatherTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeatherTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Weathertype() IWeathertypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWeathertypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWeathertypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WeatherTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WeatherTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WeatherTypeParserRULE_expression)

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
		p.Weathertype()
	}
	{
		p.SetState(5)
		p.Match(WeatherTypeParserEOF)
	}

	return localctx
}

// IWeathertypeContext is an interface to support dynamic dispatch.
type IWeathertypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeathertypeContext differentiates from other interfaces.
	IsWeathertypeContext()
}

type WeathertypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeathertypeContext() *WeathertypeContext {
	var p = new(WeathertypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WeatherTypeParserRULE_weathertype
	return p
}

func (*WeathertypeContext) IsWeathertypeContext() {}

func NewWeathertypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeathertypeContext {
	var p = new(WeathertypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeatherTypeParserRULE_weathertype

	return p
}

func (s *WeathertypeContext) GetParser() antlr.Parser { return s.parser }

func (s *WeathertypeContext) WEATHERTYPE() antlr.TerminalNode {
	return s.GetToken(WeatherTypeParserWEATHERTYPE, 0)
}

func (s *WeathertypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeathertypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeathertypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherTypeListener); ok {
		listenerT.EnterWeathertype(s)
	}
}

func (s *WeathertypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherTypeListener); ok {
		listenerT.ExitWeathertype(s)
	}
}

func (p *WeatherTypeParser) Weathertype() (localctx IWeathertypeContext) {
	this := p
	_ = this

	localctx = NewWeathertypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WeatherTypeParserRULE_weathertype)

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
		p.Match(WeatherTypeParserWEATHERTYPE)
	}

	return localctx
}
