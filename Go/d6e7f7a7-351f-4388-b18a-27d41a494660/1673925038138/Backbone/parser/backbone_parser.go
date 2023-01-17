// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925038138/Backbone.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Backbone

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

type BackboneParser struct {
	*antlr.BaseParser
}

var backboneParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func backboneParserInit() {
	staticData := &backboneParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BACKBONE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "backbone",
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

// BackboneParserInit initializes any static state used to implement BackboneParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBackboneParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BackboneParserInit() {
	staticData := &backboneParserStaticData
	staticData.once.Do(backboneParserInit)
}

// NewBackboneParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBackboneParser(input antlr.TokenStream) *BackboneParser {
	BackboneParserInit()
	this := new(BackboneParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &backboneParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Backbone.g4"

	return this
}

// BackboneParser tokens.
const (
	BackboneParserEOF      = antlr.TokenEOF
	BackboneParserBACKBONE = 1
	BackboneParserOWNKEY   = 2
	BackboneParserSPLITKEY = 3
	BackboneParserWS       = 4
)

// BackboneParser rules.
const (
	BackboneParserRULE_expression = 0
	BackboneParserRULE_backbone   = 1
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
	p.RuleIndex = BackboneParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BackboneParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Backbone() IBackboneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBackboneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBackboneContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BackboneParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BackboneListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BackboneListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BackboneParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BackboneParserRULE_expression)

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
		p.Backbone()
	}
	{
		p.SetState(5)
		p.Match(BackboneParserEOF)
	}

	return localctx
}

// IBackboneContext is an interface to support dynamic dispatch.
type IBackboneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBackboneContext differentiates from other interfaces.
	IsBackboneContext()
}

type BackboneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBackboneContext() *BackboneContext {
	var p = new(BackboneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BackboneParserRULE_backbone
	return p
}

func (*BackboneContext) IsBackboneContext() {}

func NewBackboneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BackboneContext {
	var p = new(BackboneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BackboneParserRULE_backbone

	return p
}

func (s *BackboneContext) GetParser() antlr.Parser { return s.parser }

func (s *BackboneContext) BACKBONE() antlr.TerminalNode {
	return s.GetToken(BackboneParserBACKBONE, 0)
}

func (s *BackboneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackboneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BackboneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BackboneListener); ok {
		listenerT.EnterBackbone(s)
	}
}

func (s *BackboneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BackboneListener); ok {
		listenerT.ExitBackbone(s)
	}
}

func (p *BackboneParser) Backbone() (localctx IBackboneContext) {
	this := p
	_ = this

	localctx = NewBackboneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BackboneParserRULE_backbone)

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
		p.Match(BackboneParserBACKBONE)
	}

	return localctx
}
