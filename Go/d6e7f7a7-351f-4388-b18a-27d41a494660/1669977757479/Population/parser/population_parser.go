// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669977757479/Population.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Population

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

type PopulationParser struct {
	*antlr.BaseParser
}

var populationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func populationParserInit() {
	staticData := &populationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "POPULATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "population",
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

// PopulationParserInit initializes any static state used to implement PopulationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPopulationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PopulationParserInit() {
	staticData := &populationParserStaticData
	staticData.once.Do(populationParserInit)
}

// NewPopulationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPopulationParser(input antlr.TokenStream) *PopulationParser {
	PopulationParserInit()
	this := new(PopulationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &populationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Population.g4"

	return this
}

// PopulationParser tokens.
const (
	PopulationParserEOF        = antlr.TokenEOF
	PopulationParserPOPULATION = 1
	PopulationParserOWNKEY     = 2
	PopulationParserSPLITKEY   = 3
	PopulationParserWS         = 4
)

// PopulationParser rules.
const (
	PopulationParserRULE_expression = 0
	PopulationParserRULE_population = 1
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
	p.RuleIndex = PopulationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PopulationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Population() IPopulationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPopulationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPopulationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PopulationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopulationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopulationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PopulationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PopulationParserRULE_expression)

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
		p.Population()
	}
	{
		p.SetState(5)
		p.Match(PopulationParserEOF)
	}

	return localctx
}

// IPopulationContext is an interface to support dynamic dispatch.
type IPopulationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPopulationContext differentiates from other interfaces.
	IsPopulationContext()
}

type PopulationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPopulationContext() *PopulationContext {
	var p = new(PopulationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PopulationParserRULE_population
	return p
}

func (*PopulationContext) IsPopulationContext() {}

func NewPopulationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PopulationContext {
	var p = new(PopulationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PopulationParserRULE_population

	return p
}

func (s *PopulationContext) GetParser() antlr.Parser { return s.parser }

func (s *PopulationContext) POPULATION() antlr.TerminalNode {
	return s.GetToken(PopulationParserPOPULATION, 0)
}

func (s *PopulationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PopulationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PopulationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopulationListener); ok {
		listenerT.EnterPopulation(s)
	}
}

func (s *PopulationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopulationListener); ok {
		listenerT.ExitPopulation(s)
	}
}

func (p *PopulationParser) Population() (localctx IPopulationContext) {
	this := p
	_ = this

	localctx = NewPopulationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PopulationParserRULE_population)

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
		p.Match(PopulationParserPOPULATION)
	}

	return localctx
}
