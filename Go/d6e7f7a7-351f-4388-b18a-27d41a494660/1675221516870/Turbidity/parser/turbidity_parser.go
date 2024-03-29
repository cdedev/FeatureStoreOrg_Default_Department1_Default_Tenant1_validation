// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675221516870/Turbidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Turbidity

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

type TurbidityParser struct {
	*antlr.BaseParser
}

var turbidityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func turbidityParserInit() {
	staticData := &turbidityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TURBIDITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "turbidity",
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

// TurbidityParserInit initializes any static state used to implement TurbidityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTurbidityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TurbidityParserInit() {
	staticData := &turbidityParserStaticData
	staticData.once.Do(turbidityParserInit)
}

// NewTurbidityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTurbidityParser(input antlr.TokenStream) *TurbidityParser {
	TurbidityParserInit()
	this := new(TurbidityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &turbidityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Turbidity.g4"

	return this
}

// TurbidityParser tokens.
const (
	TurbidityParserEOF       = antlr.TokenEOF
	TurbidityParserTURBIDITY = 1
	TurbidityParserOWNKEY    = 2
	TurbidityParserSPLITKEY  = 3
	TurbidityParserWS        = 4
)

// TurbidityParser rules.
const (
	TurbidityParserRULE_expression = 0
	TurbidityParserRULE_turbidity  = 1
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
	p.RuleIndex = TurbidityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurbidityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Turbidity() ITurbidityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITurbidityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITurbidityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TurbidityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurbidityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurbidityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TurbidityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TurbidityParserRULE_expression)

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
		p.Turbidity()
	}
	{
		p.SetState(5)
		p.Match(TurbidityParserEOF)
	}

	return localctx
}

// ITurbidityContext is an interface to support dynamic dispatch.
type ITurbidityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTurbidityContext differentiates from other interfaces.
	IsTurbidityContext()
}

type TurbidityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTurbidityContext() *TurbidityContext {
	var p = new(TurbidityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TurbidityParserRULE_turbidity
	return p
}

func (*TurbidityContext) IsTurbidityContext() {}

func NewTurbidityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TurbidityContext {
	var p = new(TurbidityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TurbidityParserRULE_turbidity

	return p
}

func (s *TurbidityContext) GetParser() antlr.Parser { return s.parser }

func (s *TurbidityContext) TURBIDITY() antlr.TerminalNode {
	return s.GetToken(TurbidityParserTURBIDITY, 0)
}

func (s *TurbidityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TurbidityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TurbidityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurbidityListener); ok {
		listenerT.EnterTurbidity(s)
	}
}

func (s *TurbidityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TurbidityListener); ok {
		listenerT.ExitTurbidity(s)
	}
}

func (p *TurbidityParser) Turbidity() (localctx ITurbidityContext) {
	this := p
	_ = this

	localctx = NewTurbidityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TurbidityParserRULE_turbidity)

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
		p.Match(TurbidityParserTURBIDITY)
	}

	return localctx
}
