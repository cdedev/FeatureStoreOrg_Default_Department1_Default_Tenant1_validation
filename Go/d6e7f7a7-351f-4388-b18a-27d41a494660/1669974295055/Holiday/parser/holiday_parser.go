// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669974295055/Holiday.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Holiday

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

type HolidayParser struct {
	*antlr.BaseParser
}

var holidayParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func holidayParserInit() {
	staticData := &holidayParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HOLIDAY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "holiday",
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

// HolidayParserInit initializes any static state used to implement HolidayParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHolidayParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HolidayParserInit() {
	staticData := &holidayParserStaticData
	staticData.once.Do(holidayParserInit)
}

// NewHolidayParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHolidayParser(input antlr.TokenStream) *HolidayParser {
	HolidayParserInit()
	this := new(HolidayParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &holidayParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Holiday.g4"

	return this
}

// HolidayParser tokens.
const (
	HolidayParserEOF      = antlr.TokenEOF
	HolidayParserHOLIDAY  = 1
	HolidayParserOWNKEY   = 2
	HolidayParserSPLITKEY = 3
	HolidayParserWS       = 4
)

// HolidayParser rules.
const (
	HolidayParserRULE_expression = 0
	HolidayParserRULE_holiday    = 1
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
	p.RuleIndex = HolidayParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HolidayParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Holiday() IHolidayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHolidayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHolidayContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HolidayParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HolidayListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HolidayListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HolidayParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HolidayParserRULE_expression)

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
		p.Holiday()
	}
	{
		p.SetState(5)
		p.Match(HolidayParserEOF)
	}

	return localctx
}

// IHolidayContext is an interface to support dynamic dispatch.
type IHolidayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHolidayContext differentiates from other interfaces.
	IsHolidayContext()
}

type HolidayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHolidayContext() *HolidayContext {
	var p = new(HolidayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HolidayParserRULE_holiday
	return p
}

func (*HolidayContext) IsHolidayContext() {}

func NewHolidayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HolidayContext {
	var p = new(HolidayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HolidayParserRULE_holiday

	return p
}

func (s *HolidayContext) GetParser() antlr.Parser { return s.parser }

func (s *HolidayContext) HOLIDAY() antlr.TerminalNode {
	return s.GetToken(HolidayParserHOLIDAY, 0)
}

func (s *HolidayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HolidayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HolidayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HolidayListener); ok {
		listenerT.EnterHoliday(s)
	}
}

func (s *HolidayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HolidayListener); ok {
		listenerT.ExitHoliday(s)
	}
}

func (p *HolidayParser) Holiday() (localctx IHolidayContext) {
	this := p
	_ = this

	localctx = NewHolidayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HolidayParserRULE_holiday)

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
		p.Match(HolidayParserHOLIDAY)
	}

	return localctx
}
