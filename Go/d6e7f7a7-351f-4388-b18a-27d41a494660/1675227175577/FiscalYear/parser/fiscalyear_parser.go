// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675227175577/FiscalYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FiscalYear

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

type FiscalYearParser struct {
	*antlr.BaseParser
}

var fiscalyearParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fiscalyearParserInit() {
	staticData := &fiscalyearParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FISCALYEAR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "fiscalyear",
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

// FiscalYearParserInit initializes any static state used to implement FiscalYearParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFiscalYearParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FiscalYearParserInit() {
	staticData := &fiscalyearParserStaticData
	staticData.once.Do(fiscalyearParserInit)
}

// NewFiscalYearParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFiscalYearParser(input antlr.TokenStream) *FiscalYearParser {
	FiscalYearParserInit()
	this := new(FiscalYearParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fiscalyearParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FiscalYear.g4"

	return this
}

// FiscalYearParser tokens.
const (
	FiscalYearParserEOF        = antlr.TokenEOF
	FiscalYearParserFISCALYEAR = 1
	FiscalYearParserOWNKEY     = 2
	FiscalYearParserSPLITKEY   = 3
	FiscalYearParserWS         = 4
)

// FiscalYearParser rules.
const (
	FiscalYearParserRULE_expression = 0
	FiscalYearParserRULE_fiscalyear = 1
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
	p.RuleIndex = FiscalYearParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FiscalYearParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Fiscalyear() IFiscalyearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFiscalyearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFiscalyearContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FiscalYearParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FiscalYearListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FiscalYearListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FiscalYearParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FiscalYearParserRULE_expression)

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
		p.Fiscalyear()
	}
	{
		p.SetState(5)
		p.Match(FiscalYearParserEOF)
	}

	return localctx
}

// IFiscalyearContext is an interface to support dynamic dispatch.
type IFiscalyearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFiscalyearContext differentiates from other interfaces.
	IsFiscalyearContext()
}

type FiscalyearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFiscalyearContext() *FiscalyearContext {
	var p = new(FiscalyearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FiscalYearParserRULE_fiscalyear
	return p
}

func (*FiscalyearContext) IsFiscalyearContext() {}

func NewFiscalyearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FiscalyearContext {
	var p = new(FiscalyearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FiscalYearParserRULE_fiscalyear

	return p
}

func (s *FiscalyearContext) GetParser() antlr.Parser { return s.parser }

func (s *FiscalyearContext) FISCALYEAR() antlr.TerminalNode {
	return s.GetToken(FiscalYearParserFISCALYEAR, 0)
}

func (s *FiscalyearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FiscalyearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FiscalyearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FiscalYearListener); ok {
		listenerT.EnterFiscalyear(s)
	}
}

func (s *FiscalyearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FiscalYearListener); ok {
		listenerT.ExitFiscalyear(s)
	}
}

func (p *FiscalYearParser) Fiscalyear() (localctx IFiscalyearContext) {
	this := p
	_ = this

	localctx = NewFiscalyearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FiscalYearParserRULE_fiscalyear)

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
		p.Match(FiscalYearParserFISCALYEAR)
	}

	return localctx
}
