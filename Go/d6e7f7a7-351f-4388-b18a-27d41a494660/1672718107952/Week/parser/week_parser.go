// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672718107952/Week.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Week

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

type WeekParser struct {
	*antlr.BaseParser
}

var weekParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func weekParserInit() {
	staticData := &weekParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WEEK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "week",
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

// WeekParserInit initializes any static state used to implement WeekParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWeekParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeekParserInit() {
	staticData := &weekParserStaticData
	staticData.once.Do(weekParserInit)
}

// NewWeekParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWeekParser(input antlr.TokenStream) *WeekParser {
	WeekParserInit()
	this := new(WeekParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &weekParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Week.g4"

	return this
}

// WeekParser tokens.
const (
	WeekParserEOF      = antlr.TokenEOF
	WeekParserWEEK     = 1
	WeekParserOWNKEY   = 2
	WeekParserSPLITKEY = 3
	WeekParserWS       = 4
)

// WeekParser rules.
const (
	WeekParserRULE_expression = 0
	WeekParserRULE_week       = 1
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
	p.RuleIndex = WeekParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeekParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Week() IWeekContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWeekContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWeekContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WeekParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WeekParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WeekParserRULE_expression)

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
		p.Week()
	}
	{
		p.SetState(5)
		p.Match(WeekParserEOF)
	}

	return localctx
}

// IWeekContext is an interface to support dynamic dispatch.
type IWeekContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekContext differentiates from other interfaces.
	IsWeekContext()
}

type WeekContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekContext() *WeekContext {
	var p = new(WeekContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WeekParserRULE_week
	return p
}

func (*WeekContext) IsWeekContext() {}

func NewWeekContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekContext {
	var p = new(WeekContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeekParserRULE_week

	return p
}

func (s *WeekContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekContext) WEEK() antlr.TerminalNode {
	return s.GetToken(WeekParserWEEK, 0)
}

func (s *WeekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekListener); ok {
		listenerT.EnterWeek(s)
	}
}

func (s *WeekContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekListener); ok {
		listenerT.ExitWeek(s)
	}
}

func (p *WeekParser) Week() (localctx IWeekContext) {
	this := p
	_ = this

	localctx = NewWeekContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WeekParserRULE_week)

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
		p.Match(WeekParserWEEK)
	}

	return localctx
}
