// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701718101/Shipping.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shipping

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

type ShippingParser struct {
	*antlr.BaseParser
}

var shippingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func shippingParserInit() {
	staticData := &shippingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SHIPPING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "shipping",
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

// ShippingParserInit initializes any static state used to implement ShippingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewShippingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ShippingParserInit() {
	staticData := &shippingParserStaticData
	staticData.once.Do(shippingParserInit)
}

// NewShippingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewShippingParser(input antlr.TokenStream) *ShippingParser {
	ShippingParserInit()
	this := new(ShippingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &shippingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Shipping.g4"

	return this
}

// ShippingParser tokens.
const (
	ShippingParserEOF      = antlr.TokenEOF
	ShippingParserSHIPPING = 1
	ShippingParserOWNKEY   = 2
	ShippingParserSPLITKEY = 3
	ShippingParserWS       = 4
)

// ShippingParser rules.
const (
	ShippingParserRULE_expression = 0
	ShippingParserRULE_shipping   = 1
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
	p.RuleIndex = ShippingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShippingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Shipping() IShippingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShippingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShippingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ShippingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShippingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShippingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ShippingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ShippingParserRULE_expression)

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
		p.Shipping()
	}
	{
		p.SetState(5)
		p.Match(ShippingParserEOF)
	}

	return localctx
}

// IShippingContext is an interface to support dynamic dispatch.
type IShippingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShippingContext differentiates from other interfaces.
	IsShippingContext()
}

type ShippingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShippingContext() *ShippingContext {
	var p = new(ShippingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShippingParserRULE_shipping
	return p
}

func (*ShippingContext) IsShippingContext() {}

func NewShippingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShippingContext {
	var p = new(ShippingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShippingParserRULE_shipping

	return p
}

func (s *ShippingContext) GetParser() antlr.Parser { return s.parser }

func (s *ShippingContext) SHIPPING() antlr.TerminalNode {
	return s.GetToken(ShippingParserSHIPPING, 0)
}

func (s *ShippingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShippingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShippingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShippingListener); ok {
		listenerT.EnterShipping(s)
	}
}

func (s *ShippingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShippingListener); ok {
		listenerT.ExitShipping(s)
	}
}

func (p *ShippingParser) Shipping() (localctx IShippingContext) {
	this := p
	_ = this

	localctx = NewShippingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ShippingParserRULE_shipping)

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
		p.Match(ShippingParserSHIPPING)
	}

	return localctx
}
