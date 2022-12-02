// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669980860573/Engine_Power.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Engine_Power

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

type Engine_PowerParser struct {
	*antlr.BaseParser
}

var engine_powerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func engine_powerParserInit() {
	staticData := &engine_powerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ENGINE_POWER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "engine_power",
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

// Engine_PowerParserInit initializes any static state used to implement Engine_PowerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEngine_PowerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Engine_PowerParserInit() {
	staticData := &engine_powerParserStaticData
	staticData.once.Do(engine_powerParserInit)
}

// NewEngine_PowerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEngine_PowerParser(input antlr.TokenStream) *Engine_PowerParser {
	Engine_PowerParserInit()
	this := new(Engine_PowerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &engine_powerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Engine_Power.g4"

	return this
}

// Engine_PowerParser tokens.
const (
	Engine_PowerParserEOF          = antlr.TokenEOF
	Engine_PowerParserENGINE_POWER = 1
	Engine_PowerParserOWNKEY       = 2
	Engine_PowerParserSPLITKEY     = 3
	Engine_PowerParserWS           = 4
)

// Engine_PowerParser rules.
const (
	Engine_PowerParserRULE_expression   = 0
	Engine_PowerParserRULE_engine_power = 1
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
	p.RuleIndex = Engine_PowerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Engine_PowerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Engine_power() IEngine_powerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_powerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_powerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Engine_PowerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Engine_PowerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Engine_PowerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Engine_PowerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Engine_PowerParserRULE_expression)

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
		p.Engine_power()
	}
	{
		p.SetState(5)
		p.Match(Engine_PowerParserEOF)
	}

	return localctx
}

// IEngine_powerContext is an interface to support dynamic dispatch.
type IEngine_powerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEngine_powerContext differentiates from other interfaces.
	IsEngine_powerContext()
}

type Engine_powerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_powerContext() *Engine_powerContext {
	var p = new(Engine_powerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Engine_PowerParserRULE_engine_power
	return p
}

func (*Engine_powerContext) IsEngine_powerContext() {}

func NewEngine_powerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_powerContext {
	var p = new(Engine_powerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Engine_PowerParserRULE_engine_power

	return p
}

func (s *Engine_powerContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_powerContext) ENGINE_POWER() antlr.TerminalNode {
	return s.GetToken(Engine_PowerParserENGINE_POWER, 0)
}

func (s *Engine_powerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_powerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Engine_powerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Engine_PowerListener); ok {
		listenerT.EnterEngine_power(s)
	}
}

func (s *Engine_powerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Engine_PowerListener); ok {
		listenerT.ExitEngine_power(s)
	}
}

func (p *Engine_PowerParser) Engine_power() (localctx IEngine_powerContext) {
	this := p
	_ = this

	localctx = NewEngine_powerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Engine_PowerParserRULE_engine_power)

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
		p.Match(Engine_PowerParserENGINE_POWER)
	}

	return localctx
}
