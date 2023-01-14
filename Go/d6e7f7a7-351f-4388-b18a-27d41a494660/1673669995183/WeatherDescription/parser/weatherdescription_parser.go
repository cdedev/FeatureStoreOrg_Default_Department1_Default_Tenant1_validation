// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673669995183/WeatherDescription.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeatherDescription

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

type WeatherDescriptionParser struct {
	*antlr.BaseParser
}

var weatherdescriptionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func weatherdescriptionParserInit() {
	staticData := &weatherdescriptionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WEATHERDESCRIPTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "weatherdescription",
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

// WeatherDescriptionParserInit initializes any static state used to implement WeatherDescriptionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWeatherDescriptionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeatherDescriptionParserInit() {
	staticData := &weatherdescriptionParserStaticData
	staticData.once.Do(weatherdescriptionParserInit)
}

// NewWeatherDescriptionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWeatherDescriptionParser(input antlr.TokenStream) *WeatherDescriptionParser {
	WeatherDescriptionParserInit()
	this := new(WeatherDescriptionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &weatherdescriptionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "WeatherDescription.g4"

	return this
}

// WeatherDescriptionParser tokens.
const (
	WeatherDescriptionParserEOF                = antlr.TokenEOF
	WeatherDescriptionParserWEATHERDESCRIPTION = 1
	WeatherDescriptionParserOWNKEY             = 2
	WeatherDescriptionParserSPLITKEY           = 3
	WeatherDescriptionParserWS                 = 4
)

// WeatherDescriptionParser rules.
const (
	WeatherDescriptionParserRULE_expression         = 0
	WeatherDescriptionParserRULE_weatherdescription = 1
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
	p.RuleIndex = WeatherDescriptionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeatherDescriptionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Weatherdescription() IWeatherdescriptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWeatherdescriptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWeatherdescriptionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WeatherDescriptionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherDescriptionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherDescriptionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WeatherDescriptionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WeatherDescriptionParserRULE_expression)

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
		p.Weatherdescription()
	}
	{
		p.SetState(5)
		p.Match(WeatherDescriptionParserEOF)
	}

	return localctx
}

// IWeatherdescriptionContext is an interface to support dynamic dispatch.
type IWeatherdescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeatherdescriptionContext differentiates from other interfaces.
	IsWeatherdescriptionContext()
}

type WeatherdescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeatherdescriptionContext() *WeatherdescriptionContext {
	var p = new(WeatherdescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WeatherDescriptionParserRULE_weatherdescription
	return p
}

func (*WeatherdescriptionContext) IsWeatherdescriptionContext() {}

func NewWeatherdescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeatherdescriptionContext {
	var p = new(WeatherdescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeatherDescriptionParserRULE_weatherdescription

	return p
}

func (s *WeatherdescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *WeatherdescriptionContext) WEATHERDESCRIPTION() antlr.TerminalNode {
	return s.GetToken(WeatherDescriptionParserWEATHERDESCRIPTION, 0)
}

func (s *WeatherdescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeatherdescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeatherdescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherDescriptionListener); ok {
		listenerT.EnterWeatherdescription(s)
	}
}

func (s *WeatherdescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeatherDescriptionListener); ok {
		listenerT.ExitWeatherdescription(s)
	}
}

func (p *WeatherDescriptionParser) Weatherdescription() (localctx IWeatherdescriptionContext) {
	this := p
	_ = this

	localctx = NewWeatherdescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WeatherDescriptionParserRULE_weatherdescription)

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
		p.Match(WeatherDescriptionParserWEATHERDESCRIPTION)
	}

	return localctx
}
