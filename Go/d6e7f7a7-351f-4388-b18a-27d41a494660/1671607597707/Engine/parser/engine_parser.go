// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671607597707/Engine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Engine

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

type EngineParser struct {
	*antlr.BaseParser
}

var engineParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func engineParserInit() {
	staticData := &engineParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ENGINE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "engine",
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

// EngineParserInit initializes any static state used to implement EngineParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEngineParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EngineParserInit() {
	staticData := &engineParserStaticData
	staticData.once.Do(engineParserInit)
}

// NewEngineParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEngineParser(input antlr.TokenStream) *EngineParser {
	EngineParserInit()
	this := new(EngineParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &engineParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Engine.g4"

	return this
}

// EngineParser tokens.
const (
	EngineParserEOF      = antlr.TokenEOF
	EngineParserENGINE   = 1
	EngineParserOWNKEY   = 2
	EngineParserSPLITKEY = 3
	EngineParserWS       = 4
)

// EngineParser rules.
const (
	EngineParserRULE_expression = 0
	EngineParserRULE_engine     = 1
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
	p.RuleIndex = EngineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EngineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Engine() IEngineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngineContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EngineParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EngineParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EngineParserRULE_expression)

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
		p.Engine()
	}
	{
		p.SetState(5)
		p.Match(EngineParserEOF)
	}

	return localctx
}

// IEngineContext is an interface to support dynamic dispatch.
type IEngineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEngineContext differentiates from other interfaces.
	IsEngineContext()
}

type EngineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngineContext() *EngineContext {
	var p = new(EngineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EngineParserRULE_engine
	return p
}

func (*EngineContext) IsEngineContext() {}

func NewEngineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EngineContext {
	var p = new(EngineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EngineParserRULE_engine

	return p
}

func (s *EngineContext) GetParser() antlr.Parser { return s.parser }

func (s *EngineContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(EngineParserENGINE, 0)
}

func (s *EngineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EngineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EngineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineListener); ok {
		listenerT.EnterEngine(s)
	}
}

func (s *EngineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineListener); ok {
		listenerT.ExitEngine(s)
	}
}

func (p *EngineParser) Engine() (localctx IEngineContext) {
	this := p
	_ = this

	localctx = NewEngineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EngineParserRULE_engine)

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
		p.Match(EngineParserENGINE)
	}

	return localctx
}
