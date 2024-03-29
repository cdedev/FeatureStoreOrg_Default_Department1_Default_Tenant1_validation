// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719079626/Housing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Housing

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

type HousingParser struct {
	*antlr.BaseParser
}

var housingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func housingParserInit() {
	staticData := &housingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HOUSING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "housing",
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

// HousingParserInit initializes any static state used to implement HousingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHousingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HousingParserInit() {
	staticData := &housingParserStaticData
	staticData.once.Do(housingParserInit)
}

// NewHousingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHousingParser(input antlr.TokenStream) *HousingParser {
	HousingParserInit()
	this := new(HousingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &housingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Housing.g4"

	return this
}

// HousingParser tokens.
const (
	HousingParserEOF      = antlr.TokenEOF
	HousingParserHOUSING  = 1
	HousingParserOWNKEY   = 2
	HousingParserSPLITKEY = 3
	HousingParserWS       = 4
)

// HousingParser rules.
const (
	HousingParserRULE_expression = 0
	HousingParserRULE_housing    = 1
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
	p.RuleIndex = HousingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HousingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Housing() IHousingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHousingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHousingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HousingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HousingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HousingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HousingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HousingParserRULE_expression)

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
		p.Housing()
	}
	{
		p.SetState(5)
		p.Match(HousingParserEOF)
	}

	return localctx
}

// IHousingContext is an interface to support dynamic dispatch.
type IHousingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHousingContext differentiates from other interfaces.
	IsHousingContext()
}

type HousingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHousingContext() *HousingContext {
	var p = new(HousingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HousingParserRULE_housing
	return p
}

func (*HousingContext) IsHousingContext() {}

func NewHousingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HousingContext {
	var p = new(HousingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HousingParserRULE_housing

	return p
}

func (s *HousingContext) GetParser() antlr.Parser { return s.parser }

func (s *HousingContext) HOUSING() antlr.TerminalNode {
	return s.GetToken(HousingParserHOUSING, 0)
}

func (s *HousingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HousingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HousingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HousingListener); ok {
		listenerT.EnterHousing(s)
	}
}

func (s *HousingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HousingListener); ok {
		listenerT.ExitHousing(s)
	}
}

func (p *HousingParser) Housing() (localctx IHousingContext) {
	this := p
	_ = this

	localctx = NewHousingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HousingParserRULE_housing)

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
		p.Match(HousingParserHOUSING)
	}

	return localctx
}
