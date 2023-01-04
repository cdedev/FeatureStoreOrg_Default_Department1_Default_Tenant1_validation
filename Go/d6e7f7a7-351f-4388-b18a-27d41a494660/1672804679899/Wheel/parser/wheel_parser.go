// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804679899/Wheel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wheel

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

type WheelParser struct {
	*antlr.BaseParser
}

var wheelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func wheelParserInit() {
	staticData := &wheelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WHEEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "wheel",
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

// WheelParserInit initializes any static state used to implement WheelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWheelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WheelParserInit() {
	staticData := &wheelParserStaticData
	staticData.once.Do(wheelParserInit)
}

// NewWheelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWheelParser(input antlr.TokenStream) *WheelParser {
	WheelParserInit()
	this := new(WheelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &wheelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Wheel.g4"

	return this
}

// WheelParser tokens.
const (
	WheelParserEOF      = antlr.TokenEOF
	WheelParserWHEEL    = 1
	WheelParserOWNKEY   = 2
	WheelParserSPLITKEY = 3
	WheelParserWS       = 4
)

// WheelParser rules.
const (
	WheelParserRULE_expression = 0
	WheelParserRULE_wheel      = 1
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
	p.RuleIndex = WheelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WheelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Wheel() IWheelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWheelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWheelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WheelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WheelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WheelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WheelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WheelParserRULE_expression)

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
		p.Wheel()
	}
	{
		p.SetState(5)
		p.Match(WheelParserEOF)
	}

	return localctx
}

// IWheelContext is an interface to support dynamic dispatch.
type IWheelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWheelContext differentiates from other interfaces.
	IsWheelContext()
}

type WheelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWheelContext() *WheelContext {
	var p = new(WheelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WheelParserRULE_wheel
	return p
}

func (*WheelContext) IsWheelContext() {}

func NewWheelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WheelContext {
	var p = new(WheelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WheelParserRULE_wheel

	return p
}

func (s *WheelContext) GetParser() antlr.Parser { return s.parser }

func (s *WheelContext) WHEEL() antlr.TerminalNode {
	return s.GetToken(WheelParserWHEEL, 0)
}

func (s *WheelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WheelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WheelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WheelListener); ok {
		listenerT.EnterWheel(s)
	}
}

func (s *WheelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WheelListener); ok {
		listenerT.ExitWheel(s)
	}
}

func (p *WheelParser) Wheel() (localctx IWheelContext) {
	this := p
	_ = this

	localctx = NewWheelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WheelParserRULE_wheel)

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
		p.Match(WheelParserWHEEL)
	}

	return localctx
}
