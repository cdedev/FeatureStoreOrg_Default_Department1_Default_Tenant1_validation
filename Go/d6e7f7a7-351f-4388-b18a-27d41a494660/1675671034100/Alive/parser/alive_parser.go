// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675671034100/Alive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Alive

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

type AliveParser struct {
	*antlr.BaseParser
}

var aliveParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aliveParserInit() {
	staticData := &aliveParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ALIVE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "alive",
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

// AliveParserInit initializes any static state used to implement AliveParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAliveParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AliveParserInit() {
	staticData := &aliveParserStaticData
	staticData.once.Do(aliveParserInit)
}

// NewAliveParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAliveParser(input antlr.TokenStream) *AliveParser {
	AliveParserInit()
	this := new(AliveParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &aliveParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Alive.g4"

	return this
}

// AliveParser tokens.
const (
	AliveParserEOF      = antlr.TokenEOF
	AliveParserALIVE    = 1
	AliveParserOWNKEY   = 2
	AliveParserSPLITKEY = 3
	AliveParserWS       = 4
)

// AliveParser rules.
const (
	AliveParserRULE_expression = 0
	AliveParserRULE_alive      = 1
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
	p.RuleIndex = AliveParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliveParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Alive() IAliveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliveContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AliveParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliveListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliveListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AliveParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AliveParserRULE_expression)

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
		p.Alive()
	}
	{
		p.SetState(5)
		p.Match(AliveParserEOF)
	}

	return localctx
}

// IAliveContext is an interface to support dynamic dispatch.
type IAliveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliveContext differentiates from other interfaces.
	IsAliveContext()
}

type AliveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliveContext() *AliveContext {
	var p = new(AliveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AliveParserRULE_alive
	return p
}

func (*AliveContext) IsAliveContext() {}

func NewAliveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliveContext {
	var p = new(AliveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AliveParserRULE_alive

	return p
}

func (s *AliveContext) GetParser() antlr.Parser { return s.parser }

func (s *AliveContext) ALIVE() antlr.TerminalNode {
	return s.GetToken(AliveParserALIVE, 0)
}

func (s *AliveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliveListener); ok {
		listenerT.EnterAlive(s)
	}
}

func (s *AliveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AliveListener); ok {
		listenerT.ExitAlive(s)
	}
}

func (p *AliveParser) Alive() (localctx IAliveContext) {
	this := p
	_ = this

	localctx = NewAliveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AliveParserRULE_alive)

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
		p.Match(AliveParserALIVE)
	}

	return localctx
}
