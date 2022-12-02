// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983780063/CO2Emissions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CO2Emissions

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

type CO2EmissionsParser struct {
	*antlr.BaseParser
}

var co2emissionsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func co2emissionsParserInit() {
	staticData := &co2emissionsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CO2EMISSIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "co2emissions",
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

// CO2EmissionsParserInit initializes any static state used to implement CO2EmissionsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCO2EmissionsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CO2EmissionsParserInit() {
	staticData := &co2emissionsParserStaticData
	staticData.once.Do(co2emissionsParserInit)
}

// NewCO2EmissionsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCO2EmissionsParser(input antlr.TokenStream) *CO2EmissionsParser {
	CO2EmissionsParserInit()
	this := new(CO2EmissionsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &co2emissionsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CO2Emissions.g4"

	return this
}

// CO2EmissionsParser tokens.
const (
	CO2EmissionsParserEOF          = antlr.TokenEOF
	CO2EmissionsParserCO2EMISSIONS = 1
	CO2EmissionsParserOWNKEY       = 2
	CO2EmissionsParserSPLITKEY     = 3
	CO2EmissionsParserWS           = 4
)

// CO2EmissionsParser rules.
const (
	CO2EmissionsParserRULE_expression   = 0
	CO2EmissionsParserRULE_co2emissions = 1
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
	p.RuleIndex = CO2EmissionsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CO2EmissionsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Co2emissions() ICo2emissionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICo2emissionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICo2emissionsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CO2EmissionsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CO2EmissionsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CO2EmissionsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CO2EmissionsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CO2EmissionsParserRULE_expression)

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
		p.Co2emissions()
	}
	{
		p.SetState(5)
		p.Match(CO2EmissionsParserEOF)
	}

	return localctx
}

// ICo2emissionsContext is an interface to support dynamic dispatch.
type ICo2emissionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCo2emissionsContext differentiates from other interfaces.
	IsCo2emissionsContext()
}

type Co2emissionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCo2emissionsContext() *Co2emissionsContext {
	var p = new(Co2emissionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CO2EmissionsParserRULE_co2emissions
	return p
}

func (*Co2emissionsContext) IsCo2emissionsContext() {}

func NewCo2emissionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Co2emissionsContext {
	var p = new(Co2emissionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CO2EmissionsParserRULE_co2emissions

	return p
}

func (s *Co2emissionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Co2emissionsContext) CO2EMISSIONS() antlr.TerminalNode {
	return s.GetToken(CO2EmissionsParserCO2EMISSIONS, 0)
}

func (s *Co2emissionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Co2emissionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Co2emissionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CO2EmissionsListener); ok {
		listenerT.EnterCo2emissions(s)
	}
}

func (s *Co2emissionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CO2EmissionsListener); ok {
		listenerT.ExitCo2emissions(s)
	}
}

func (p *CO2EmissionsParser) Co2emissions() (localctx ICo2emissionsContext) {
	this := p
	_ = this

	localctx = NewCo2emissionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CO2EmissionsParserRULE_co2emissions)

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
		p.Match(CO2EmissionsParserCO2EMISSIONS)
	}

	return localctx
}
