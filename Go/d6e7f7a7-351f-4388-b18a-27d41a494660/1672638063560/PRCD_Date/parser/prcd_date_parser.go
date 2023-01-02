// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638063560/PRCD_Date.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PRCD_Date

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

type PRCD_DateParser struct {
	*antlr.BaseParser
}

var prcd_dateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func prcd_dateParserInit() {
	staticData := &prcd_dateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PRCD_DATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "prcd_date",
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

// PRCD_DateParserInit initializes any static state used to implement PRCD_DateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPRCD_DateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PRCD_DateParserInit() {
	staticData := &prcd_dateParserStaticData
	staticData.once.Do(prcd_dateParserInit)
}

// NewPRCD_DateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPRCD_DateParser(input antlr.TokenStream) *PRCD_DateParser {
	PRCD_DateParserInit()
	this := new(PRCD_DateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &prcd_dateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PRCD_Date.g4"

	return this
}

// PRCD_DateParser tokens.
const (
	PRCD_DateParserEOF       = antlr.TokenEOF
	PRCD_DateParserPRCD_DATE = 1
	PRCD_DateParserOWNKEY    = 2
	PRCD_DateParserSPLITKEY  = 3
	PRCD_DateParserWS        = 4
)

// PRCD_DateParser rules.
const (
	PRCD_DateParserRULE_expression = 0
	PRCD_DateParserRULE_prcd_date  = 1
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
	p.RuleIndex = PRCD_DateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRCD_DateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Prcd_date() IPrcd_dateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrcd_dateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrcd_dateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PRCD_DateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRCD_DateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRCD_DateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PRCD_DateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PRCD_DateParserRULE_expression)

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
		p.Prcd_date()
	}
	{
		p.SetState(5)
		p.Match(PRCD_DateParserEOF)
	}

	return localctx
}

// IPrcd_dateContext is an interface to support dynamic dispatch.
type IPrcd_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrcd_dateContext differentiates from other interfaces.
	IsPrcd_dateContext()
}

type Prcd_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrcd_dateContext() *Prcd_dateContext {
	var p = new(Prcd_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PRCD_DateParserRULE_prcd_date
	return p
}

func (*Prcd_dateContext) IsPrcd_dateContext() {}

func NewPrcd_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prcd_dateContext {
	var p = new(Prcd_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PRCD_DateParserRULE_prcd_date

	return p
}

func (s *Prcd_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Prcd_dateContext) PRCD_DATE() antlr.TerminalNode {
	return s.GetToken(PRCD_DateParserPRCD_DATE, 0)
}

func (s *Prcd_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prcd_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prcd_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRCD_DateListener); ok {
		listenerT.EnterPrcd_date(s)
	}
}

func (s *Prcd_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PRCD_DateListener); ok {
		listenerT.ExitPrcd_date(s)
	}
}

func (p *PRCD_DateParser) Prcd_date() (localctx IPrcd_dateContext) {
	this := p
	_ = this

	localctx = NewPrcd_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PRCD_DateParserRULE_prcd_date)

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
		p.Match(PRCD_DateParserPRCD_DATE)
	}

	return localctx
}
