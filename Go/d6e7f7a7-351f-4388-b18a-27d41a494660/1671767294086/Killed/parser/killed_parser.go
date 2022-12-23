// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671767294086/Killed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Killed

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

type KilledParser struct {
	*antlr.BaseParser
}

var killedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func killedParserInit() {
	staticData := &killedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "KILLED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "killed",
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

// KilledParserInit initializes any static state used to implement KilledParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKilledParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KilledParserInit() {
	staticData := &killedParserStaticData
	staticData.once.Do(killedParserInit)
}

// NewKilledParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKilledParser(input antlr.TokenStream) *KilledParser {
	KilledParserInit()
	this := new(KilledParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &killedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Killed.g4"

	return this
}

// KilledParser tokens.
const (
	KilledParserEOF      = antlr.TokenEOF
	KilledParserKILLED   = 1
	KilledParserOWNKEY   = 2
	KilledParserSPLITKEY = 3
	KilledParserWS       = 4
)

// KilledParser rules.
const (
	KilledParserRULE_expression = 0
	KilledParserRULE_killed     = 1
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
	p.RuleIndex = KilledParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KilledParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Killed() IKilledContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKilledContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKilledContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(KilledParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilledListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilledListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *KilledParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KilledParserRULE_expression)

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
		p.Killed()
	}
	{
		p.SetState(5)
		p.Match(KilledParserEOF)
	}

	return localctx
}

// IKilledContext is an interface to support dynamic dispatch.
type IKilledContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKilledContext differentiates from other interfaces.
	IsKilledContext()
}

type KilledContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKilledContext() *KilledContext {
	var p = new(KilledContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KilledParserRULE_killed
	return p
}

func (*KilledContext) IsKilledContext() {}

func NewKilledContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KilledContext {
	var p = new(KilledContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KilledParserRULE_killed

	return p
}

func (s *KilledContext) GetParser() antlr.Parser { return s.parser }

func (s *KilledContext) KILLED() antlr.TerminalNode {
	return s.GetToken(KilledParserKILLED, 0)
}

func (s *KilledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KilledContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KilledContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilledListener); ok {
		listenerT.EnterKilled(s)
	}
}

func (s *KilledContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilledListener); ok {
		listenerT.ExitKilled(s)
	}
}

func (p *KilledParser) Killed() (localctx IKilledContext) {
	this := p
	_ = this

	localctx = NewKilledContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KilledParserRULE_killed)

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
		p.Match(KilledParserKILLED)
	}

	return localctx
}
