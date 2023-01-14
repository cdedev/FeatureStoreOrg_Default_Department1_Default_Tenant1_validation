// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673672257964/Currency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Currency

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

type CurrencyParser struct {
	*antlr.BaseParser
}

var currencyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func currencyParserInit() {
	staticData := &currencyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CURRENCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "currency",
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

// CurrencyParserInit initializes any static state used to implement CurrencyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCurrencyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CurrencyParserInit() {
	staticData := &currencyParserStaticData
	staticData.once.Do(currencyParserInit)
}

// NewCurrencyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCurrencyParser(input antlr.TokenStream) *CurrencyParser {
	CurrencyParserInit()
	this := new(CurrencyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &currencyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Currency.g4"

	return this
}

// CurrencyParser tokens.
const (
	CurrencyParserEOF      = antlr.TokenEOF
	CurrencyParserCURRENCY = 1
	CurrencyParserOWNKEY   = 2
	CurrencyParserSPLITKEY = 3
	CurrencyParserWS       = 4
)

// CurrencyParser rules.
const (
	CurrencyParserRULE_expression = 0
	CurrencyParserRULE_currency   = 1
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
	p.RuleIndex = CurrencyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CurrencyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Currency() ICurrencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurrencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurrencyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CurrencyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CurrencyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CurrencyParserRULE_expression)

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
		p.Currency()
	}
	{
		p.SetState(5)
		p.Match(CurrencyParserEOF)
	}

	return localctx
}

// ICurrencyContext is an interface to support dynamic dispatch.
type ICurrencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrencyContext differentiates from other interfaces.
	IsCurrencyContext()
}

type CurrencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrencyContext() *CurrencyContext {
	var p = new(CurrencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CurrencyParserRULE_currency
	return p
}

func (*CurrencyContext) IsCurrencyContext() {}

func NewCurrencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrencyContext {
	var p = new(CurrencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CurrencyParserRULE_currency

	return p
}

func (s *CurrencyContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrencyContext) CURRENCY() antlr.TerminalNode {
	return s.GetToken(CurrencyParserCURRENCY, 0)
}

func (s *CurrencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrencyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyListener); ok {
		listenerT.EnterCurrency(s)
	}
}

func (s *CurrencyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyListener); ok {
		listenerT.ExitCurrency(s)
	}
}

func (p *CurrencyParser) Currency() (localctx ICurrencyContext) {
	this := p
	_ = this

	localctx = NewCurrencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CurrencyParserRULE_currency)

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
		p.Match(CurrencyParserCURRENCY)
	}

	return localctx
}
