// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674623607607/Stock.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stock

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

type StockParser struct {
	*antlr.BaseParser
}

var stockParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stockParserInit() {
	staticData := &stockParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STOCK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "stock",
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

// StockParserInit initializes any static state used to implement StockParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStockParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StockParserInit() {
	staticData := &stockParserStaticData
	staticData.once.Do(stockParserInit)
}

// NewStockParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStockParser(input antlr.TokenStream) *StockParser {
	StockParserInit()
	this := new(StockParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &stockParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Stock.g4"

	return this
}

// StockParser tokens.
const (
	StockParserEOF      = antlr.TokenEOF
	StockParserSTOCK    = 1
	StockParserOWNKEY   = 2
	StockParserSPLITKEY = 3
	StockParserWS       = 4
)

// StockParser rules.
const (
	StockParserRULE_expression = 0
	StockParserRULE_stock      = 1
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
	p.RuleIndex = StockParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StockParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Stock() IStockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStockContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StockParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StockListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StockListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StockParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StockParserRULE_expression)

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
		p.Stock()
	}
	{
		p.SetState(5)
		p.Match(StockParserEOF)
	}

	return localctx
}

// IStockContext is an interface to support dynamic dispatch.
type IStockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStockContext differentiates from other interfaces.
	IsStockContext()
}

type StockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStockContext() *StockContext {
	var p = new(StockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StockParserRULE_stock
	return p
}

func (*StockContext) IsStockContext() {}

func NewStockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StockContext {
	var p = new(StockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StockParserRULE_stock

	return p
}

func (s *StockContext) GetParser() antlr.Parser { return s.parser }

func (s *StockContext) STOCK() antlr.TerminalNode {
	return s.GetToken(StockParserSTOCK, 0)
}

func (s *StockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StockListener); ok {
		listenerT.EnterStock(s)
	}
}

func (s *StockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StockListener); ok {
		listenerT.ExitStock(s)
	}
}

func (p *StockParser) Stock() (localctx IStockContext) {
	this := p
	_ = this

	localctx = NewStockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StockParserRULE_stock)

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
		p.Match(StockParserSTOCK)
	}

	return localctx
}
