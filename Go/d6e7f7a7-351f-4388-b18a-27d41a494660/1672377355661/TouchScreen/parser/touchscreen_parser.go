// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377355661/TouchScreen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TouchScreen

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

type TouchScreenParser struct {
	*antlr.BaseParser
}

var touchscreenParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func touchscreenParserInit() {
	staticData := &touchscreenParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TOUCHSCREEN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "touchscreen",
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

// TouchScreenParserInit initializes any static state used to implement TouchScreenParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTouchScreenParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TouchScreenParserInit() {
	staticData := &touchscreenParserStaticData
	staticData.once.Do(touchscreenParserInit)
}

// NewTouchScreenParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTouchScreenParser(input antlr.TokenStream) *TouchScreenParser {
	TouchScreenParserInit()
	this := new(TouchScreenParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &touchscreenParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "TouchScreen.g4"

	return this
}

// TouchScreenParser tokens.
const (
	TouchScreenParserEOF         = antlr.TokenEOF
	TouchScreenParserTOUCHSCREEN = 1
	TouchScreenParserOWNKEY      = 2
	TouchScreenParserSPLITKEY    = 3
	TouchScreenParserWS          = 4
)

// TouchScreenParser rules.
const (
	TouchScreenParserRULE_expression  = 0
	TouchScreenParserRULE_touchscreen = 1
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
	p.RuleIndex = TouchScreenParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TouchScreenParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Touchscreen() ITouchscreenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITouchscreenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITouchscreenContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TouchScreenParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TouchScreenListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TouchScreenListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TouchScreenParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TouchScreenParserRULE_expression)

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
		p.Touchscreen()
	}
	{
		p.SetState(5)
		p.Match(TouchScreenParserEOF)
	}

	return localctx
}

// ITouchscreenContext is an interface to support dynamic dispatch.
type ITouchscreenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTouchscreenContext differentiates from other interfaces.
	IsTouchscreenContext()
}

type TouchscreenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTouchscreenContext() *TouchscreenContext {
	var p = new(TouchscreenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TouchScreenParserRULE_touchscreen
	return p
}

func (*TouchscreenContext) IsTouchscreenContext() {}

func NewTouchscreenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TouchscreenContext {
	var p = new(TouchscreenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TouchScreenParserRULE_touchscreen

	return p
}

func (s *TouchscreenContext) GetParser() antlr.Parser { return s.parser }

func (s *TouchscreenContext) TOUCHSCREEN() antlr.TerminalNode {
	return s.GetToken(TouchScreenParserTOUCHSCREEN, 0)
}

func (s *TouchscreenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TouchscreenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TouchscreenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TouchScreenListener); ok {
		listenerT.EnterTouchscreen(s)
	}
}

func (s *TouchscreenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TouchScreenListener); ok {
		listenerT.ExitTouchscreen(s)
	}
}

func (p *TouchScreenParser) Touchscreen() (localctx ITouchscreenContext) {
	this := p
	_ = this

	localctx = NewTouchscreenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TouchScreenParserRULE_touchscreen)

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
		p.Match(TouchScreenParserTOUCHSCREEN)
	}

	return localctx
}
