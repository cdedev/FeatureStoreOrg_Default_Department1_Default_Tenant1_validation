// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603911590/Bohrium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bohrium

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

type BohriumParser struct {
	*antlr.BaseParser
}

var bohriumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bohriumParserInit() {
	staticData := &bohriumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BOHRIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bohrium",
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

// BohriumParserInit initializes any static state used to implement BohriumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBohriumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BohriumParserInit() {
	staticData := &bohriumParserStaticData
	staticData.once.Do(bohriumParserInit)
}

// NewBohriumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBohriumParser(input antlr.TokenStream) *BohriumParser {
	BohriumParserInit()
	this := new(BohriumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bohriumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bohrium.g4"

	return this
}

// BohriumParser tokens.
const (
	BohriumParserEOF      = antlr.TokenEOF
	BohriumParserBOHRIUM  = 1
	BohriumParserOWNKEY   = 2
	BohriumParserSPLITKEY = 3
	BohriumParserWS       = 4
)

// BohriumParser rules.
const (
	BohriumParserRULE_expression = 0
	BohriumParserRULE_bohrium    = 1
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
	p.RuleIndex = BohriumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BohriumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bohrium() IBohriumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBohriumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBohriumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BohriumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BohriumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BohriumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BohriumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BohriumParserRULE_expression)

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
		p.Bohrium()
	}
	{
		p.SetState(5)
		p.Match(BohriumParserEOF)
	}

	return localctx
}

// IBohriumContext is an interface to support dynamic dispatch.
type IBohriumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBohriumContext differentiates from other interfaces.
	IsBohriumContext()
}

type BohriumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBohriumContext() *BohriumContext {
	var p = new(BohriumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BohriumParserRULE_bohrium
	return p
}

func (*BohriumContext) IsBohriumContext() {}

func NewBohriumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BohriumContext {
	var p = new(BohriumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BohriumParserRULE_bohrium

	return p
}

func (s *BohriumContext) GetParser() antlr.Parser { return s.parser }

func (s *BohriumContext) BOHRIUM() antlr.TerminalNode {
	return s.GetToken(BohriumParserBOHRIUM, 0)
}

func (s *BohriumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BohriumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BohriumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BohriumListener); ok {
		listenerT.EnterBohrium(s)
	}
}

func (s *BohriumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BohriumListener); ok {
		listenerT.ExitBohrium(s)
	}
}

func (p *BohriumParser) Bohrium() (localctx IBohriumContext) {
	this := p
	_ = this

	localctx = NewBohriumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BohriumParserRULE_bohrium)

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
		p.Match(BohriumParserBOHRIUM)
	}

	return localctx
}
