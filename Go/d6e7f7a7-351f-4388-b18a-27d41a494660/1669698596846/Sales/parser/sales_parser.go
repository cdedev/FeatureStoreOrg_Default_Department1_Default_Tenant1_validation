// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669698596846/Sales.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sales

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

type SalesParser struct {
	*antlr.BaseParser
}

var salesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func salesParserInit() {
	staticData := &salesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SALES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sales",
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

// SalesParserInit initializes any static state used to implement SalesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSalesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SalesParserInit() {
	staticData := &salesParserStaticData
	staticData.once.Do(salesParserInit)
}

// NewSalesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSalesParser(input antlr.TokenStream) *SalesParser {
	SalesParserInit()
	this := new(SalesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &salesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sales.g4"

	return this
}

// SalesParser tokens.
const (
	SalesParserEOF      = antlr.TokenEOF
	SalesParserSALES    = 1
	SalesParserOWNKEY   = 2
	SalesParserSPLITKEY = 3
	SalesParserWS       = 4
)

// SalesParser rules.
const (
	SalesParserRULE_expression = 0
	SalesParserRULE_sales      = 1
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
	p.RuleIndex = SalesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SalesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sales() ISalesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISalesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISalesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SalesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SalesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SalesParserRULE_expression)

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
		p.Sales()
	}
	{
		p.SetState(5)
		p.Match(SalesParserEOF)
	}

	return localctx
}

// ISalesContext is an interface to support dynamic dispatch.
type ISalesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalesContext differentiates from other interfaces.
	IsSalesContext()
}

type SalesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalesContext() *SalesContext {
	var p = new(SalesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SalesParserRULE_sales
	return p
}

func (*SalesContext) IsSalesContext() {}

func NewSalesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalesContext {
	var p = new(SalesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SalesParserRULE_sales

	return p
}

func (s *SalesContext) GetParser() antlr.Parser { return s.parser }

func (s *SalesContext) SALES() antlr.TerminalNode {
	return s.GetToken(SalesParserSALES, 0)
}

func (s *SalesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesListener); ok {
		listenerT.EnterSales(s)
	}
}

func (s *SalesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalesListener); ok {
		listenerT.ExitSales(s)
	}
}

func (p *SalesParser) Sales() (localctx ISalesContext) {
	this := p
	_ = this

	localctx = NewSalesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SalesParserRULE_sales)

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
		p.Match(SalesParserSALES)
	}

	return localctx
}
