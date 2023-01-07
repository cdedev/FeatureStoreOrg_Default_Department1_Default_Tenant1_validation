// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083087516/SalesOutlet.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SalesOutlet

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

type SalesOutletParser struct {
	*antlr.BaseParser
}

var salesoutletParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func salesoutletParserInit() {
	staticData := &salesoutletParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SALESOUTLET", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "salesoutlet",
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

// SalesOutletParserInit initializes any static state used to implement SalesOutletParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSalesOutletParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SalesOutletParserInit() {
	staticData := &salesoutletParserStaticData
	staticData.once.Do(salesoutletParserInit)
}

// NewSalesOutletParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSalesOutletParser(input antlr.TokenStream) *SalesOutletParser {
	SalesOutletParserInit()
	this := new(SalesOutletParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &salesoutletParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SalesOutlet.g4"

	return this
}

// SalesOutletParser tokens.
const (
	SalesOutletParserEOF         = antlr.TokenEOF
	SalesOutletParserSALESOUTLET = 1
	SalesOutletParserOWNKEY      = 2
	SalesOutletParserSPLITKEY    = 3
	SalesOutletParserWS          = 4
)

// SalesOutletParser rules.
const (
	SalesOutletParserRULE_expression  = 0
	SalesOutletParserRULE_salesoutlet = 1
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
	p.RuleIndex = SalesOutletParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SalesOutletParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Salesoutlet() ISalesoutletContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISalesoutletContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISalesoutletContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SalesOutletParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesOutletListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesOutletListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SalesOutletParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SalesOutletParserRULE_expression)

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
		p.Salesoutlet()
	}
	{
		p.SetState(5)
		p.Match(SalesOutletParserEOF)
	}

	return localctx
}

// ISalesoutletContext is an interface to support dynamic dispatch.
type ISalesoutletContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalesoutletContext differentiates from other interfaces.
	IsSalesoutletContext()
}

type SalesoutletContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalesoutletContext() *SalesoutletContext {
	var p = new(SalesoutletContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SalesOutletParserRULE_salesoutlet
	return p
}

func (*SalesoutletContext) IsSalesoutletContext() {}

func NewSalesoutletContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalesoutletContext {
	var p = new(SalesoutletContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SalesOutletParserRULE_salesoutlet

	return p
}

func (s *SalesoutletContext) GetParser() antlr.Parser { return s.parser }

func (s *SalesoutletContext) SALESOUTLET() antlr.TerminalNode {
	return s.GetToken(SalesOutletParserSALESOUTLET, 0)
}

func (s *SalesoutletContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalesoutletContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalesoutletContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesOutletListener); ok {
		listenerT.EnterSalesoutlet(s)
	}
}

func (s *SalesoutletContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesOutletListener); ok {
		listenerT.ExitSalesoutlet(s)
	}
}

func (p *SalesOutletParser) Salesoutlet() (localctx ISalesoutletContext) {
	this := p
	_ = this

	localctx = NewSalesoutletContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SalesOutletParserRULE_salesoutlet)

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
		p.Match(SalesOutletParserSALESOUTLET)
	}

	return localctx
}
