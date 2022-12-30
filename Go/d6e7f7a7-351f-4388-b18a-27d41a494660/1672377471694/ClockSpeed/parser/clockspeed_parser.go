// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377471694/ClockSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ClockSpeed

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

type ClockSpeedParser struct {
	*antlr.BaseParser
}

var clockspeedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clockspeedParserInit() {
	staticData := &clockspeedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLOCKSPEED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "clockspeed",
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

// ClockSpeedParserInit initializes any static state used to implement ClockSpeedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClockSpeedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClockSpeedParserInit() {
	staticData := &clockspeedParserStaticData
	staticData.once.Do(clockspeedParserInit)
}

// NewClockSpeedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClockSpeedParser(input antlr.TokenStream) *ClockSpeedParser {
	ClockSpeedParserInit()
	this := new(ClockSpeedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clockspeedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ClockSpeed.g4"

	return this
}

// ClockSpeedParser tokens.
const (
	ClockSpeedParserEOF        = antlr.TokenEOF
	ClockSpeedParserCLOCKSPEED = 1
	ClockSpeedParserOWNKEY     = 2
	ClockSpeedParserSPLITKEY   = 3
	ClockSpeedParserWS         = 4
)

// ClockSpeedParser rules.
const (
	ClockSpeedParserRULE_expression = 0
	ClockSpeedParserRULE_clockspeed = 1
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
	p.RuleIndex = ClockSpeedParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClockSpeedParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Clockspeed() IClockspeedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClockspeedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClockspeedContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClockSpeedParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClockSpeedListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClockSpeedListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ClockSpeedParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClockSpeedParserRULE_expression)

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
		p.Clockspeed()
	}
	{
		p.SetState(5)
		p.Match(ClockSpeedParserEOF)
	}

	return localctx
}

// IClockspeedContext is an interface to support dynamic dispatch.
type IClockspeedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClockspeedContext differentiates from other interfaces.
	IsClockspeedContext()
}

type ClockspeedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClockspeedContext() *ClockspeedContext {
	var p = new(ClockspeedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClockSpeedParserRULE_clockspeed
	return p
}

func (*ClockspeedContext) IsClockspeedContext() {}

func NewClockspeedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClockspeedContext {
	var p = new(ClockspeedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClockSpeedParserRULE_clockspeed

	return p
}

func (s *ClockspeedContext) GetParser() antlr.Parser { return s.parser }

func (s *ClockspeedContext) CLOCKSPEED() antlr.TerminalNode {
	return s.GetToken(ClockSpeedParserCLOCKSPEED, 0)
}

func (s *ClockspeedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClockspeedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClockspeedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClockSpeedListener); ok {
		listenerT.EnterClockspeed(s)
	}
}

func (s *ClockspeedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClockSpeedListener); ok {
		listenerT.ExitClockspeed(s)
	}
}

func (p *ClockSpeedParser) Clockspeed() (localctx IClockspeedContext) {
	this := p
	_ = this

	localctx = NewClockspeedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClockSpeedParserRULE_clockspeed)

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
		p.Match(ClockSpeedParserCLOCKSPEED)
	}

	return localctx
}
