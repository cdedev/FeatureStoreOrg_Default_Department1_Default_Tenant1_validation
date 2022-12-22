// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690372200/BodilyInjuries.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BodilyInjuries

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

type BodilyInjuriesParser struct {
	*antlr.BaseParser
}

var bodilyinjuriesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bodilyinjuriesParserInit() {
	staticData := &bodilyinjuriesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BODILYINJURIES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bodilyinjuries",
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

// BodilyInjuriesParserInit initializes any static state used to implement BodilyInjuriesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBodilyInjuriesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BodilyInjuriesParserInit() {
	staticData := &bodilyinjuriesParserStaticData
	staticData.once.Do(bodilyinjuriesParserInit)
}

// NewBodilyInjuriesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBodilyInjuriesParser(input antlr.TokenStream) *BodilyInjuriesParser {
	BodilyInjuriesParserInit()
	this := new(BodilyInjuriesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bodilyinjuriesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BodilyInjuries.g4"

	return this
}

// BodilyInjuriesParser tokens.
const (
	BodilyInjuriesParserEOF            = antlr.TokenEOF
	BodilyInjuriesParserBODILYINJURIES = 1
	BodilyInjuriesParserOWNKEY         = 2
	BodilyInjuriesParserSPLITKEY       = 3
	BodilyInjuriesParserWS             = 4
)

// BodilyInjuriesParser rules.
const (
	BodilyInjuriesParserRULE_expression     = 0
	BodilyInjuriesParserRULE_bodilyinjuries = 1
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
	p.RuleIndex = BodilyInjuriesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BodilyInjuriesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bodilyinjuries() IBodilyinjuriesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBodilyinjuriesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBodilyinjuriesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BodilyInjuriesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BodilyInjuriesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BodilyInjuriesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BodilyInjuriesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BodilyInjuriesParserRULE_expression)

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
		p.Bodilyinjuries()
	}
	{
		p.SetState(5)
		p.Match(BodilyInjuriesParserEOF)
	}

	return localctx
}

// IBodilyinjuriesContext is an interface to support dynamic dispatch.
type IBodilyinjuriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBodilyinjuriesContext differentiates from other interfaces.
	IsBodilyinjuriesContext()
}

type BodilyinjuriesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBodilyinjuriesContext() *BodilyinjuriesContext {
	var p = new(BodilyinjuriesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BodilyInjuriesParserRULE_bodilyinjuries
	return p
}

func (*BodilyinjuriesContext) IsBodilyinjuriesContext() {}

func NewBodilyinjuriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BodilyinjuriesContext {
	var p = new(BodilyinjuriesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BodilyInjuriesParserRULE_bodilyinjuries

	return p
}

func (s *BodilyinjuriesContext) GetParser() antlr.Parser { return s.parser }

func (s *BodilyinjuriesContext) BODILYINJURIES() antlr.TerminalNode {
	return s.GetToken(BodilyInjuriesParserBODILYINJURIES, 0)
}

func (s *BodilyinjuriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BodilyinjuriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BodilyinjuriesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BodilyInjuriesListener); ok {
		listenerT.EnterBodilyinjuries(s)
	}
}

func (s *BodilyinjuriesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BodilyInjuriesListener); ok {
		listenerT.ExitBodilyinjuries(s)
	}
}

func (p *BodilyInjuriesParser) Bodilyinjuries() (localctx IBodilyinjuriesContext) {
	this := p
	_ = this

	localctx = NewBodilyinjuriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BodilyInjuriesParserRULE_bodilyinjuries)

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
		p.Match(BodilyInjuriesParserBODILYINJURIES)
	}

	return localctx
}
