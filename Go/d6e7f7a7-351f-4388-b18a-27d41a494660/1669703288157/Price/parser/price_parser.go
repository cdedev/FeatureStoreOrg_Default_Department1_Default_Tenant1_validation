// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669703288157/Price.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Price

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

type PriceParser struct {
	*antlr.BaseParser
}

var priceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func priceParserInit() {
	staticData := &priceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PRICE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "price",
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

// PriceParserInit initializes any static state used to implement PriceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPriceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PriceParserInit() {
	staticData := &priceParserStaticData
	staticData.once.Do(priceParserInit)
}

// NewPriceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPriceParser(input antlr.TokenStream) *PriceParser {
	PriceParserInit()
	this := new(PriceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &priceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Price.g4"

	return this
}

// PriceParser tokens.
const (
	PriceParserEOF      = antlr.TokenEOF
	PriceParserPRICE    = 1
	PriceParserOWNKEY   = 2
	PriceParserSPLITKEY = 3
	PriceParserWS       = 4
)

// PriceParser rules.
const (
	PriceParserRULE_expression = 0
	PriceParserRULE_price      = 1
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
	p.RuleIndex = PriceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PriceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Price() IPriceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPriceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPriceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PriceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PriceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PriceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PriceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PriceParserRULE_expression)

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
		p.Price()
	}
	{
		p.SetState(5)
		p.Match(PriceParserEOF)
	}

	return localctx
}

// IPriceContext is an interface to support dynamic dispatch.
type IPriceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPriceContext differentiates from other interfaces.
	IsPriceContext()
}

type PriceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPriceContext() *PriceContext {
	var p = new(PriceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PriceParserRULE_price
	return p
}

func (*PriceContext) IsPriceContext() {}

func NewPriceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PriceContext {
	var p = new(PriceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PriceParserRULE_price

	return p
}

func (s *PriceContext) GetParser() antlr.Parser { return s.parser }

func (s *PriceContext) PRICE() antlr.TerminalNode {
	return s.GetToken(PriceParserPRICE, 0)
}

func (s *PriceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PriceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PriceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PriceListener); ok {
		listenerT.EnterPrice(s)
	}
}

func (s *PriceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PriceListener); ok {
		listenerT.ExitPrice(s)
	}
}

func (p *PriceParser) Price() (localctx IPriceContext) {
	this := p
	_ = this

	localctx = NewPriceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PriceParserRULE_price)

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
		p.Match(PriceParserPRICE)
	}

	return localctx
}
