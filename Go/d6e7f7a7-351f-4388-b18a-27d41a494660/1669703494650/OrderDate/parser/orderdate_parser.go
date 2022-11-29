// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669703494650/OrderDate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OrderDate

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

type OrderDateParser struct {
	*antlr.BaseParser
}

var orderdateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func orderdateParserInit() {
	staticData := &orderdateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ORDERDATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "orderdate",
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

// OrderDateParserInit initializes any static state used to implement OrderDateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOrderDateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OrderDateParserInit() {
	staticData := &orderdateParserStaticData
	staticData.once.Do(orderdateParserInit)
}

// NewOrderDateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOrderDateParser(input antlr.TokenStream) *OrderDateParser {
	OrderDateParserInit()
	this := new(OrderDateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &orderdateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "OrderDate.g4"

	return this
}

// OrderDateParser tokens.
const (
	OrderDateParserEOF       = antlr.TokenEOF
	OrderDateParserORDERDATE = 1
	OrderDateParserOWNKEY    = 2
	OrderDateParserSPLITKEY  = 3
	OrderDateParserWS        = 4
)

// OrderDateParser rules.
const (
	OrderDateParserRULE_expression = 0
	OrderDateParserRULE_orderdate  = 1
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
	p.RuleIndex = OrderDateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OrderDateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Orderdate() IOrderdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderdateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OrderDateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderDateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderDateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OrderDateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OrderDateParserRULE_expression)

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
		p.Orderdate()
	}
	{
		p.SetState(5)
		p.Match(OrderDateParserEOF)
	}

	return localctx
}

// IOrderdateContext is an interface to support dynamic dispatch.
type IOrderdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderdateContext differentiates from other interfaces.
	IsOrderdateContext()
}

type OrderdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderdateContext() *OrderdateContext {
	var p = new(OrderdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OrderDateParserRULE_orderdate
	return p
}

func (*OrderdateContext) IsOrderdateContext() {}

func NewOrderdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderdateContext {
	var p = new(OrderdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OrderDateParserRULE_orderdate

	return p
}

func (s *OrderdateContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderdateContext) ORDERDATE() antlr.TerminalNode {
	return s.GetToken(OrderDateParserORDERDATE, 0)
}

func (s *OrderdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderDateListener); ok {
		listenerT.EnterOrderdate(s)
	}
}

func (s *OrderdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderDateListener); ok {
		listenerT.ExitOrderdate(s)
	}
}

func (p *OrderDateParser) Orderdate() (localctx IOrderdateContext) {
	this := p
	_ = this

	localctx = NewOrderdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OrderDateParserRULE_orderdate)

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
		p.Match(OrderDateParserORDERDATE)
	}

	return localctx
}
