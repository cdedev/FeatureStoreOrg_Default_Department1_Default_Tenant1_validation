// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675222188616/Month.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Month

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

type MonthParser struct {
	*antlr.BaseParser
}

var monthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monthParserInit() {
	staticData := &monthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MONTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "month",
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

// MonthParserInit initializes any static state used to implement MonthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMonthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonthParserInit() {
	staticData := &monthParserStaticData
	staticData.once.Do(monthParserInit)
}

// NewMonthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMonthParser(input antlr.TokenStream) *MonthParser {
	MonthParserInit()
	this := new(MonthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &monthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Month.g4"

	return this
}

// MonthParser tokens.
const (
	MonthParserEOF      = antlr.TokenEOF
	MonthParserMONTH    = 1
	MonthParserOWNKEY   = 2
	MonthParserSPLITKEY = 3
	MonthParserWS       = 4
)

// MonthParser rules.
const (
	MonthParserRULE_expression = 0
	MonthParserRULE_month      = 1
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
	p.RuleIndex = MonthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Month() IMonthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMonthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MonthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MonthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MonthParserRULE_expression)

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
		p.Month()
	}
	{
		p.SetState(5)
		p.Match(MonthParserEOF)
	}

	return localctx
}

// IMonthContext is an interface to support dynamic dispatch.
type IMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthContext differentiates from other interfaces.
	IsMonthContext()
}

type MonthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthContext() *MonthContext {
	var p = new(MonthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonthParserRULE_month
	return p
}

func (*MonthContext) IsMonthContext() {}

func NewMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthContext {
	var p = new(MonthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonthParserRULE_month

	return p
}

func (s *MonthContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthContext) MONTH() antlr.TerminalNode {
	return s.GetToken(MonthParserMONTH, 0)
}

func (s *MonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthListener); ok {
		listenerT.EnterMonth(s)
	}
}

func (s *MonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthListener); ok {
		listenerT.ExitMonth(s)
	}
}

func (p *MonthParser) Month() (localctx IMonthContext) {
	this := p
	_ = this

	localctx = NewMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonthParserRULE_month)

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
		p.Match(MonthParserMONTH)
	}

	return localctx
}
