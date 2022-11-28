// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624790720/City.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // City

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

type CityParser struct {
	*antlr.BaseParser
}

var cityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cityParserInit() {
	staticData := &cityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "city",
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

// CityParserInit initializes any static state used to implement CityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CityParserInit() {
	staticData := &cityParserStaticData
	staticData.once.Do(cityParserInit)
}

// NewCityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCityParser(input antlr.TokenStream) *CityParser {
	CityParserInit()
	this := new(CityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "City.g4"

	return this
}

// CityParser tokens.
const (
	CityParserEOF      = antlr.TokenEOF
	CityParserCITY     = 1
	CityParserOWNKEY   = 2
	CityParserSPLITKEY = 3
	CityParserWS       = 4
)

// CityParser rules.
const (
	CityParserRULE_expression = 0
	CityParserRULE_city       = 1
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
	p.RuleIndex = CityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) City() ICityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CityParserRULE_expression)

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
		p.City()
	}
	{
		p.SetState(5)
		p.Match(CityParserEOF)
	}

	return localctx
}

// ICityContext is an interface to support dynamic dispatch.
type ICityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCityContext differentiates from other interfaces.
	IsCityContext()
}

type CityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCityContext() *CityContext {
	var p = new(CityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CityParserRULE_city
	return p
}

func (*CityContext) IsCityContext() {}

func NewCityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CityContext {
	var p = new(CityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CityParserRULE_city

	return p
}

func (s *CityContext) GetParser() antlr.Parser { return s.parser }

func (s *CityContext) CITY() antlr.TerminalNode {
	return s.GetToken(CityParserCITY, 0)
}

func (s *CityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CityListener); ok {
		listenerT.EnterCity(s)
	}
}

func (s *CityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CityListener); ok {
		listenerT.ExitCity(s)
	}
}

func (p *CityParser) City() (localctx ICityContext) {
	this := p
	_ = this

	localctx = NewCityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CityParserRULE_city)

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
		p.Match(CityParserCITY)
	}

	return localctx
}
