// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669702959119/Order.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Order

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

type OrderParser struct {
	*antlr.BaseParser
}

var orderParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func orderParserInit() {
	staticData := &orderParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ORDER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "order",
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

// OrderParserInit initializes any static state used to implement OrderParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOrderParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OrderParserInit() {
	staticData := &orderParserStaticData
	staticData.once.Do(orderParserInit)
}

// NewOrderParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOrderParser(input antlr.TokenStream) *OrderParser {
	OrderParserInit()
	this := new(OrderParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &orderParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Order.g4"

	return this
}

// OrderParser tokens.
const (
	OrderParserEOF      = antlr.TokenEOF
	OrderParserORDER    = 1
	OrderParserOWNKEY   = 2
	OrderParserSPLITKEY = 3
	OrderParserWS       = 4
)

// OrderParser rules.
const (
	OrderParserRULE_expression = 0
	OrderParserRULE_order      = 1
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
	p.RuleIndex = OrderParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OrderParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Order() IOrderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OrderParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OrderParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OrderParserRULE_expression)

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
		p.Order()
	}
	{
		p.SetState(5)
		p.Match(OrderParserEOF)
	}

	return localctx
}

// IOrderContext is an interface to support dynamic dispatch.
type IOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderContext differentiates from other interfaces.
	IsOrderContext()
}

type OrderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderContext() *OrderContext {
	var p = new(OrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OrderParserRULE_order
	return p
}

func (*OrderContext) IsOrderContext() {}

func NewOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderContext {
	var p = new(OrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OrderParserRULE_order

	return p
}

func (s *OrderContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderContext) ORDER() antlr.TerminalNode {
	return s.GetToken(OrderParserORDER, 0)
}

func (s *OrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderListener); ok {
		listenerT.EnterOrder(s)
	}
}

func (s *OrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrderListener); ok {
		listenerT.ExitOrder(s)
	}
}

func (p *OrderParser) Order() (localctx IOrderContext) {
	this := p
	_ = this

	localctx = NewOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OrderParserRULE_order)

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
		p.Match(OrderParserORDER)
	}

	return localctx
}
