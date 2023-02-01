// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675242465648/Discount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Discount

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

type DiscountParser struct {
	*antlr.BaseParser
}

var discountParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func discountParserInit() {
	staticData := &discountParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DISCOUNT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "discount",
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

// DiscountParserInit initializes any static state used to implement DiscountParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiscountParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiscountParserInit() {
	staticData := &discountParserStaticData
	staticData.once.Do(discountParserInit)
}

// NewDiscountParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiscountParser(input antlr.TokenStream) *DiscountParser {
	DiscountParserInit()
	this := new(DiscountParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &discountParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Discount.g4"

	return this
}

// DiscountParser tokens.
const (
	DiscountParserEOF      = antlr.TokenEOF
	DiscountParserDISCOUNT = 1
	DiscountParserOWNKEY   = 2
	DiscountParserSPLITKEY = 3
	DiscountParserWS       = 4
)

// DiscountParser rules.
const (
	DiscountParserRULE_expression = 0
	DiscountParserRULE_discount   = 1
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
	p.RuleIndex = DiscountParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiscountParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Discount() IDiscountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiscountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiscountContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiscountParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiscountListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiscountListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiscountParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiscountParserRULE_expression)

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
		p.Discount()
	}
	{
		p.SetState(5)
		p.Match(DiscountParserEOF)
	}

	return localctx
}

// IDiscountContext is an interface to support dynamic dispatch.
type IDiscountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiscountContext differentiates from other interfaces.
	IsDiscountContext()
}

type DiscountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiscountContext() *DiscountContext {
	var p = new(DiscountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiscountParserRULE_discount
	return p
}

func (*DiscountContext) IsDiscountContext() {}

func NewDiscountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiscountContext {
	var p = new(DiscountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiscountParserRULE_discount

	return p
}

func (s *DiscountContext) GetParser() antlr.Parser { return s.parser }

func (s *DiscountContext) DISCOUNT() antlr.TerminalNode {
	return s.GetToken(DiscountParserDISCOUNT, 0)
}

func (s *DiscountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiscountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiscountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiscountListener); ok {
		listenerT.EnterDiscount(s)
	}
}

func (s *DiscountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiscountListener); ok {
		listenerT.ExitDiscount(s)
	}
}

func (p *DiscountParser) Discount() (localctx IDiscountContext) {
	this := p
	_ = this

	localctx = NewDiscountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiscountParserRULE_discount)

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
		p.Match(DiscountParserDISCOUNT)
	}

	return localctx
}
