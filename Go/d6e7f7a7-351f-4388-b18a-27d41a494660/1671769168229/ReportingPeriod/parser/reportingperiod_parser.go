// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769168229/ReportingPeriod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReportingPeriod

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

type ReportingPeriodParser struct {
	*antlr.BaseParser
}

var reportingperiodParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func reportingperiodParserInit() {
	staticData := &reportingperiodParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "REPORTINGPERIOD", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "reportingperiod",
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

// ReportingPeriodParserInit initializes any static state used to implement ReportingPeriodParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewReportingPeriodParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ReportingPeriodParserInit() {
	staticData := &reportingperiodParserStaticData
	staticData.once.Do(reportingperiodParserInit)
}

// NewReportingPeriodParser produces a new parser instance for the optional input antlr.TokenStream.
func NewReportingPeriodParser(input antlr.TokenStream) *ReportingPeriodParser {
	ReportingPeriodParserInit()
	this := new(ReportingPeriodParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &reportingperiodParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ReportingPeriod.g4"

	return this
}

// ReportingPeriodParser tokens.
const (
	ReportingPeriodParserEOF             = antlr.TokenEOF
	ReportingPeriodParserREPORTINGPERIOD = 1
	ReportingPeriodParserOWNKEY          = 2
	ReportingPeriodParserSPLITKEY        = 3
	ReportingPeriodParserWS              = 4
)

// ReportingPeriodParser rules.
const (
	ReportingPeriodParserRULE_expression      = 0
	ReportingPeriodParserRULE_reportingperiod = 1
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
	p.RuleIndex = ReportingPeriodParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReportingPeriodParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Reportingperiod() IReportingperiodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReportingperiodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReportingperiodContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ReportingPeriodParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReportingPeriodListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReportingPeriodListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ReportingPeriodParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReportingPeriodParserRULE_expression)

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
		p.Reportingperiod()
	}
	{
		p.SetState(5)
		p.Match(ReportingPeriodParserEOF)
	}

	return localctx
}

// IReportingperiodContext is an interface to support dynamic dispatch.
type IReportingperiodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReportingperiodContext differentiates from other interfaces.
	IsReportingperiodContext()
}

type ReportingperiodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReportingperiodContext() *ReportingperiodContext {
	var p = new(ReportingperiodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReportingPeriodParserRULE_reportingperiod
	return p
}

func (*ReportingperiodContext) IsReportingperiodContext() {}

func NewReportingperiodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReportingperiodContext {
	var p = new(ReportingperiodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReportingPeriodParserRULE_reportingperiod

	return p
}

func (s *ReportingperiodContext) GetParser() antlr.Parser { return s.parser }

func (s *ReportingperiodContext) REPORTINGPERIOD() antlr.TerminalNode {
	return s.GetToken(ReportingPeriodParserREPORTINGPERIOD, 0)
}

func (s *ReportingperiodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReportingperiodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReportingperiodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReportingPeriodListener); ok {
		listenerT.EnterReportingperiod(s)
	}
}

func (s *ReportingperiodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReportingPeriodListener); ok {
		listenerT.ExitReportingperiod(s)
	}
}

func (p *ReportingPeriodParser) Reportingperiod() (localctx IReportingperiodContext) {
	this := p
	_ = this

	localctx = NewReportingperiodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReportingPeriodParserRULE_reportingperiod)

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
		p.Match(ReportingPeriodParserREPORTINGPERIOD)
	}

	return localctx
}
