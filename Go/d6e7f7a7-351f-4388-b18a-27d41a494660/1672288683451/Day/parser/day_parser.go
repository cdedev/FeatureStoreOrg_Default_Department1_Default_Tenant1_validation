// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288683451/Day.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Day

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

type DayParser struct {
	*antlr.BaseParser
}

var dayParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dayParserInit() {
	staticData := &dayParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DAY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "day",
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

// DayParserInit initializes any static state used to implement DayParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDayParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DayParserInit() {
	staticData := &dayParserStaticData
	staticData.once.Do(dayParserInit)
}

// NewDayParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDayParser(input antlr.TokenStream) *DayParser {
	DayParserInit()
	this := new(DayParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dayParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Day.g4"

	return this
}

// DayParser tokens.
const (
	DayParserEOF      = antlr.TokenEOF
	DayParserDAY      = 1
	DayParserOWNKEY   = 2
	DayParserSPLITKEY = 3
	DayParserWS       = 4
)

// DayParser rules.
const (
	DayParserRULE_expression = 0
	DayParserRULE_day        = 1
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
	p.RuleIndex = DayParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DayParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Day() IDayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DayParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DayListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DayListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DayParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DayParserRULE_expression)

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
		p.Day()
	}
	{
		p.SetState(5)
		p.Match(DayParserEOF)
	}

	return localctx
}

// IDayContext is an interface to support dynamic dispatch.
type IDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayContext differentiates from other interfaces.
	IsDayContext()
}

type DayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayContext() *DayContext {
	var p = new(DayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DayParserRULE_day
	return p
}

func (*DayContext) IsDayContext() {}

func NewDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayContext {
	var p = new(DayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DayParserRULE_day

	return p
}

func (s *DayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayContext) DAY() antlr.TerminalNode {
	return s.GetToken(DayParserDAY, 0)
}

func (s *DayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DayListener); ok {
		listenerT.EnterDay(s)
	}
}

func (s *DayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DayListener); ok {
		listenerT.ExitDay(s)
	}
}

func (p *DayParser) Day() (localctx IDayContext) {
	this := p
	_ = this

	localctx = NewDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DayParserRULE_day)

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
		p.Match(DayParserDAY)
	}

	return localctx
}
