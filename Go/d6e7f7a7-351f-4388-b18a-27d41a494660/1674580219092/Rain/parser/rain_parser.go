// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674580219092/Rain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rain

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

type RainParser struct {
	*antlr.BaseParser
}

var rainParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rainParserInit() {
	staticData := &rainParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RAIN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rain",
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

// RainParserInit initializes any static state used to implement RainParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRainParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RainParserInit() {
	staticData := &rainParserStaticData
	staticData.once.Do(rainParserInit)
}

// NewRainParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRainParser(input antlr.TokenStream) *RainParser {
	RainParserInit()
	this := new(RainParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rainParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rain.g4"

	return this
}

// RainParser tokens.
const (
	RainParserEOF      = antlr.TokenEOF
	RainParserRAIN     = 1
	RainParserOWNKEY   = 2
	RainParserSPLITKEY = 3
	RainParserWS       = 4
)

// RainParser rules.
const (
	RainParserRULE_expression = 0
	RainParserRULE_rain       = 1
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
	p.RuleIndex = RainParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RainParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rain() IRainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRainContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RainParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RainListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RainListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RainParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RainParserRULE_expression)

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
		p.Rain()
	}
	{
		p.SetState(5)
		p.Match(RainParserEOF)
	}

	return localctx
}

// IRainContext is an interface to support dynamic dispatch.
type IRainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRainContext differentiates from other interfaces.
	IsRainContext()
}

type RainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRainContext() *RainContext {
	var p = new(RainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RainParserRULE_rain
	return p
}

func (*RainContext) IsRainContext() {}

func NewRainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RainContext {
	var p = new(RainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RainParserRULE_rain

	return p
}

func (s *RainContext) GetParser() antlr.Parser { return s.parser }

func (s *RainContext) RAIN() antlr.TerminalNode {
	return s.GetToken(RainParserRAIN, 0)
}

func (s *RainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RainListener); ok {
		listenerT.EnterRain(s)
	}
}

func (s *RainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RainListener); ok {
		listenerT.ExitRain(s)
	}
}

func (p *RainParser) Rain() (localctx IRainContext) {
	this := p
	_ = this

	localctx = NewRainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RainParserRULE_rain)

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
		p.Match(RainParserRAIN)
	}

	return localctx
}
