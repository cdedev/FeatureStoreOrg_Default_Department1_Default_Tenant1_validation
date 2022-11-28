// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669629590914/Calls.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calls

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

type CallsParser struct {
	*antlr.BaseParser
}

var callsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func callsParserInit() {
	staticData := &callsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CALLS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "calls",
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

// CallsParserInit initializes any static state used to implement CallsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCallsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CallsParserInit() {
	staticData := &callsParserStaticData
	staticData.once.Do(callsParserInit)
}

// NewCallsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCallsParser(input antlr.TokenStream) *CallsParser {
	CallsParserInit()
	this := new(CallsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &callsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Calls.g4"

	return this
}

// CallsParser tokens.
const (
	CallsParserEOF      = antlr.TokenEOF
	CallsParserCALLS    = 1
	CallsParserOWNKEY   = 2
	CallsParserSPLITKEY = 3
	CallsParserWS       = 4
)

// CallsParser rules.
const (
	CallsParserRULE_expression = 0
	CallsParserRULE_calls      = 1
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
	p.RuleIndex = CallsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CallsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Calls() ICallsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CallsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CallsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CallsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CallsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CallsParserRULE_expression)

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
		p.Calls()
	}
	{
		p.SetState(5)
		p.Match(CallsParserEOF)
	}

	return localctx
}

// ICallsContext is an interface to support dynamic dispatch.
type ICallsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallsContext differentiates from other interfaces.
	IsCallsContext()
}

type CallsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallsContext() *CallsContext {
	var p = new(CallsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CallsParserRULE_calls
	return p
}

func (*CallsContext) IsCallsContext() {}

func NewCallsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallsContext {
	var p = new(CallsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CallsParserRULE_calls

	return p
}

func (s *CallsContext) GetParser() antlr.Parser { return s.parser }

func (s *CallsContext) CALLS() antlr.TerminalNode {
	return s.GetToken(CallsParserCALLS, 0)
}

func (s *CallsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CallsListener); ok {
		listenerT.EnterCalls(s)
	}
}

func (s *CallsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CallsListener); ok {
		listenerT.ExitCalls(s)
	}
}

func (p *CallsParser) Calls() (localctx ICallsContext) {
	this := p
	_ = this

	localctx = NewCallsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CallsParserRULE_calls)

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
		p.Match(CallsParserCALLS)
	}

	return localctx
}
