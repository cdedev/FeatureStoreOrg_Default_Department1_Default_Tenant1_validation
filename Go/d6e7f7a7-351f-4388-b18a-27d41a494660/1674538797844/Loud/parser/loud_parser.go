// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538797844/Loud.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loud

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

type LoudParser struct {
	*antlr.BaseParser
}

var loudParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func loudParserInit() {
	staticData := &loudParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LOUD", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "loud",
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

// LoudParserInit initializes any static state used to implement LoudParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLoudParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LoudParserInit() {
	staticData := &loudParserStaticData
	staticData.once.Do(loudParserInit)
}

// NewLoudParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLoudParser(input antlr.TokenStream) *LoudParser {
	LoudParserInit()
	this := new(LoudParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &loudParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Loud.g4"

	return this
}

// LoudParser tokens.
const (
	LoudParserEOF      = antlr.TokenEOF
	LoudParserLOUD     = 1
	LoudParserOWNKEY   = 2
	LoudParserSPLITKEY = 3
	LoudParserWS       = 4
)

// LoudParser rules.
const (
	LoudParserRULE_expression = 0
	LoudParserRULE_loud       = 1
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
	p.RuleIndex = LoudParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoudParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Loud() ILoudContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoudContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoudContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LoudParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LoudParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LoudParserRULE_expression)

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
		p.Loud()
	}
	{
		p.SetState(5)
		p.Match(LoudParserEOF)
	}

	return localctx
}

// ILoudContext is an interface to support dynamic dispatch.
type ILoudContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoudContext differentiates from other interfaces.
	IsLoudContext()
}

type LoudContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoudContext() *LoudContext {
	var p = new(LoudContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoudParserRULE_loud
	return p
}

func (*LoudContext) IsLoudContext() {}

func NewLoudContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoudContext {
	var p = new(LoudContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoudParserRULE_loud

	return p
}

func (s *LoudContext) GetParser() antlr.Parser { return s.parser }

func (s *LoudContext) LOUD() antlr.TerminalNode {
	return s.GetToken(LoudParserLOUD, 0)
}

func (s *LoudContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoudContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoudContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudListener); ok {
		listenerT.EnterLoud(s)
	}
}

func (s *LoudContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudListener); ok {
		listenerT.ExitLoud(s)
	}
}

func (p *LoudParser) Loud() (localctx ILoudContext) {
	this := p
	_ = this

	localctx = NewLoudContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LoudParserRULE_loud)

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
		p.Match(LoudParserLOUD)
	}

	return localctx
}
