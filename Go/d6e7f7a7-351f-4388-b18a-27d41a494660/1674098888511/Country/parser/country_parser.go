// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674098888511/Country.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Country

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

type CountryParser struct {
	*antlr.BaseParser
}

var countryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func countryParserInit() {
	staticData := &countryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COUNTRY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "country",
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

// CountryParserInit initializes any static state used to implement CountryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCountryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CountryParserInit() {
	staticData := &countryParserStaticData
	staticData.once.Do(countryParserInit)
}

// NewCountryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCountryParser(input antlr.TokenStream) *CountryParser {
	CountryParserInit()
	this := new(CountryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &countryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Country.g4"

	return this
}

// CountryParser tokens.
const (
	CountryParserEOF      = antlr.TokenEOF
	CountryParserCOUNTRY  = 1
	CountryParserOWNKEY   = 2
	CountryParserSPLITKEY = 3
	CountryParserWS       = 4
)

// CountryParser rules.
const (
	CountryParserRULE_expression = 0
	CountryParserRULE_country    = 1
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
	p.RuleIndex = CountryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CountryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Country() ICountryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CountryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CountryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CountryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CountryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CountryParserRULE_expression)

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
		p.Country()
	}
	{
		p.SetState(5)
		p.Match(CountryParserEOF)
	}

	return localctx
}

// ICountryContext is an interface to support dynamic dispatch.
type ICountryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountryContext differentiates from other interfaces.
	IsCountryContext()
}

type CountryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountryContext() *CountryContext {
	var p = new(CountryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CountryParserRULE_country
	return p
}

func (*CountryContext) IsCountryContext() {}

func NewCountryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountryContext {
	var p = new(CountryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CountryParserRULE_country

	return p
}

func (s *CountryContext) GetParser() antlr.Parser { return s.parser }

func (s *CountryContext) COUNTRY() antlr.TerminalNode {
	return s.GetToken(CountryParserCOUNTRY, 0)
}

func (s *CountryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CountryListener); ok {
		listenerT.EnterCountry(s)
	}
}

func (s *CountryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CountryListener); ok {
		listenerT.ExitCountry(s)
	}
}

func (p *CountryParser) Country() (localctx ICountryContext) {
	this := p
	_ = this

	localctx = NewCountryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CountryParserRULE_country)

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
		p.Match(CountryParserCOUNTRY)
	}

	return localctx
}
