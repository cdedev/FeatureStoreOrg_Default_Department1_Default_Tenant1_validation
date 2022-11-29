// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690287872/Carbohydrates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbohydrates

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

type CarbohydratesParser struct {
	*antlr.BaseParser
}

var carbohydratesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func carbohydratesParserInit() {
	staticData := &carbohydratesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CARBOHYDRATES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "carbohydrates",
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

// CarbohydratesParserInit initializes any static state used to implement CarbohydratesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCarbohydratesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CarbohydratesParserInit() {
	staticData := &carbohydratesParserStaticData
	staticData.once.Do(carbohydratesParserInit)
}

// NewCarbohydratesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCarbohydratesParser(input antlr.TokenStream) *CarbohydratesParser {
	CarbohydratesParserInit()
	this := new(CarbohydratesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &carbohydratesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Carbohydrates.g4"

	return this
}

// CarbohydratesParser tokens.
const (
	CarbohydratesParserEOF           = antlr.TokenEOF
	CarbohydratesParserCARBOHYDRATES = 1
	CarbohydratesParserOWNKEY        = 2
	CarbohydratesParserSPLITKEY      = 3
	CarbohydratesParserWS            = 4
)

// CarbohydratesParser rules.
const (
	CarbohydratesParserRULE_expression    = 0
	CarbohydratesParserRULE_carbohydrates = 1
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
	p.RuleIndex = CarbohydratesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CarbohydratesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Carbohydrates() ICarbohydratesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICarbohydratesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICarbohydratesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CarbohydratesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbohydratesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbohydratesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CarbohydratesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CarbohydratesParserRULE_expression)

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
		p.Carbohydrates()
	}
	{
		p.SetState(5)
		p.Match(CarbohydratesParserEOF)
	}

	return localctx
}

// ICarbohydratesContext is an interface to support dynamic dispatch.
type ICarbohydratesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCarbohydratesContext differentiates from other interfaces.
	IsCarbohydratesContext()
}

type CarbohydratesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCarbohydratesContext() *CarbohydratesContext {
	var p = new(CarbohydratesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CarbohydratesParserRULE_carbohydrates
	return p
}

func (*CarbohydratesContext) IsCarbohydratesContext() {}

func NewCarbohydratesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CarbohydratesContext {
	var p = new(CarbohydratesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CarbohydratesParserRULE_carbohydrates

	return p
}

func (s *CarbohydratesContext) GetParser() antlr.Parser { return s.parser }

func (s *CarbohydratesContext) CARBOHYDRATES() antlr.TerminalNode {
	return s.GetToken(CarbohydratesParserCARBOHYDRATES, 0)
}

func (s *CarbohydratesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CarbohydratesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CarbohydratesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbohydratesListener); ok {
		listenerT.EnterCarbohydrates(s)
	}
}

func (s *CarbohydratesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbohydratesListener); ok {
		listenerT.ExitCarbohydrates(s)
	}
}

func (p *CarbohydratesParser) Carbohydrates() (localctx ICarbohydratesContext) {
	this := p
	_ = this

	localctx = NewCarbohydratesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CarbohydratesParserRULE_carbohydrates)

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
		p.Match(CarbohydratesParserCARBOHYDRATES)
	}

	return localctx
}
