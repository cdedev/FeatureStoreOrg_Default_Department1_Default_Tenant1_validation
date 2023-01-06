// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672977218378/WeekdayName.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeekdayName

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

type WeekdayNameParser struct {
	*antlr.BaseParser
}

var weekdaynameParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func weekdaynameParserInit() {
	staticData := &weekdaynameParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WEEKDAYNAME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "weekdayname",
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

// WeekdayNameParserInit initializes any static state used to implement WeekdayNameParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWeekdayNameParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeekdayNameParserInit() {
	staticData := &weekdaynameParserStaticData
	staticData.once.Do(weekdaynameParserInit)
}

// NewWeekdayNameParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWeekdayNameParser(input antlr.TokenStream) *WeekdayNameParser {
	WeekdayNameParserInit()
	this := new(WeekdayNameParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &weekdaynameParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "WeekdayName.g4"

	return this
}

// WeekdayNameParser tokens.
const (
	WeekdayNameParserEOF         = antlr.TokenEOF
	WeekdayNameParserWEEKDAYNAME = 1
	WeekdayNameParserOWNKEY      = 2
	WeekdayNameParserSPLITKEY    = 3
	WeekdayNameParserWS          = 4
)

// WeekdayNameParser rules.
const (
	WeekdayNameParserRULE_expression  = 0
	WeekdayNameParserRULE_weekdayname = 1
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
	p.RuleIndex = WeekdayNameParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeekdayNameParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Weekdayname() IWeekdaynameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWeekdaynameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWeekdaynameContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WeekdayNameParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekdayNameListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekdayNameListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WeekdayNameParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WeekdayNameParserRULE_expression)

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
		p.Weekdayname()
	}
	{
		p.SetState(5)
		p.Match(WeekdayNameParserEOF)
	}

	return localctx
}

// IWeekdaynameContext is an interface to support dynamic dispatch.
type IWeekdaynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekdaynameContext differentiates from other interfaces.
	IsWeekdaynameContext()
}

type WeekdaynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekdaynameContext() *WeekdaynameContext {
	var p = new(WeekdaynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WeekdayNameParserRULE_weekdayname
	return p
}

func (*WeekdaynameContext) IsWeekdaynameContext() {}

func NewWeekdaynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekdaynameContext {
	var p = new(WeekdaynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeekdayNameParserRULE_weekdayname

	return p
}

func (s *WeekdaynameContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekdaynameContext) WEEKDAYNAME() antlr.TerminalNode {
	return s.GetToken(WeekdayNameParserWEEKDAYNAME, 0)
}

func (s *WeekdaynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekdaynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekdaynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekdayNameListener); ok {
		listenerT.EnterWeekdayname(s)
	}
}

func (s *WeekdaynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeekdayNameListener); ok {
		listenerT.ExitWeekdayname(s)
	}
}

func (p *WeekdayNameParser) Weekdayname() (localctx IWeekdaynameContext) {
	this := p
	_ = this

	localctx = NewWeekdaynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WeekdayNameParserRULE_weekdayname)

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
		p.Match(WeekdayNameParserWEEKDAYNAME)
	}

	return localctx
}
