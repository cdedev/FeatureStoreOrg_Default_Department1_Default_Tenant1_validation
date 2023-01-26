// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674722356661/Habitat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Habitat

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

type HabitatParser struct {
	*antlr.BaseParser
}

var habitatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func habitatParserInit() {
	staticData := &habitatParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HABITAT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "habitat",
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

// HabitatParserInit initializes any static state used to implement HabitatParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHabitatParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HabitatParserInit() {
	staticData := &habitatParserStaticData
	staticData.once.Do(habitatParserInit)
}

// NewHabitatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHabitatParser(input antlr.TokenStream) *HabitatParser {
	HabitatParserInit()
	this := new(HabitatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &habitatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Habitat.g4"

	return this
}

// HabitatParser tokens.
const (
	HabitatParserEOF      = antlr.TokenEOF
	HabitatParserHABITAT  = 1
	HabitatParserOWNKEY   = 2
	HabitatParserSPLITKEY = 3
	HabitatParserWS       = 4
)

// HabitatParser rules.
const (
	HabitatParserRULE_expression = 0
	HabitatParserRULE_habitat    = 1
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
	p.RuleIndex = HabitatParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HabitatParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Habitat() IHabitatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHabitatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHabitatContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HabitatParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HabitatListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HabitatListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HabitatParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HabitatParserRULE_expression)

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
		p.Habitat()
	}
	{
		p.SetState(5)
		p.Match(HabitatParserEOF)
	}

	return localctx
}

// IHabitatContext is an interface to support dynamic dispatch.
type IHabitatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHabitatContext differentiates from other interfaces.
	IsHabitatContext()
}

type HabitatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHabitatContext() *HabitatContext {
	var p = new(HabitatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HabitatParserRULE_habitat
	return p
}

func (*HabitatContext) IsHabitatContext() {}

func NewHabitatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HabitatContext {
	var p = new(HabitatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HabitatParserRULE_habitat

	return p
}

func (s *HabitatContext) GetParser() antlr.Parser { return s.parser }

func (s *HabitatContext) HABITAT() antlr.TerminalNode {
	return s.GetToken(HabitatParserHABITAT, 0)
}

func (s *HabitatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HabitatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HabitatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HabitatListener); ok {
		listenerT.EnterHabitat(s)
	}
}

func (s *HabitatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HabitatListener); ok {
		listenerT.ExitHabitat(s)
	}
}

func (p *HabitatParser) Habitat() (localctx IHabitatContext) {
	this := p
	_ = this

	localctx = NewHabitatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HabitatParserRULE_habitat)

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
		p.Match(HabitatParserHABITAT)
	}

	return localctx
}
