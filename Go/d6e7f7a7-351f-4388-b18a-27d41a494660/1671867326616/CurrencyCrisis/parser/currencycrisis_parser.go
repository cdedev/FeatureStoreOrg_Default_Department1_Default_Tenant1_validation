// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671867326616/CurrencyCrisis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CurrencyCrisis

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

type CurrencyCrisisParser struct {
	*antlr.BaseParser
}

var currencycrisisParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func currencycrisisParserInit() {
	staticData := &currencycrisisParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CURRENCYCRISIS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "currencycrisis",
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

// CurrencyCrisisParserInit initializes any static state used to implement CurrencyCrisisParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCurrencyCrisisParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CurrencyCrisisParserInit() {
	staticData := &currencycrisisParserStaticData
	staticData.once.Do(currencycrisisParserInit)
}

// NewCurrencyCrisisParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCurrencyCrisisParser(input antlr.TokenStream) *CurrencyCrisisParser {
	CurrencyCrisisParserInit()
	this := new(CurrencyCrisisParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &currencycrisisParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CurrencyCrisis.g4"

	return this
}

// CurrencyCrisisParser tokens.
const (
	CurrencyCrisisParserEOF            = antlr.TokenEOF
	CurrencyCrisisParserCURRENCYCRISIS = 1
	CurrencyCrisisParserOWNKEY         = 2
	CurrencyCrisisParserSPLITKEY       = 3
	CurrencyCrisisParserWS             = 4
)

// CurrencyCrisisParser rules.
const (
	CurrencyCrisisParserRULE_expression     = 0
	CurrencyCrisisParserRULE_currencycrisis = 1
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
	p.RuleIndex = CurrencyCrisisParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CurrencyCrisisParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Currencycrisis() ICurrencycrisisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICurrencycrisisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICurrencycrisisContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CurrencyCrisisParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyCrisisListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyCrisisListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CurrencyCrisisParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CurrencyCrisisParserRULE_expression)

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
		p.Currencycrisis()
	}
	{
		p.SetState(5)
		p.Match(CurrencyCrisisParserEOF)
	}

	return localctx
}

// ICurrencycrisisContext is an interface to support dynamic dispatch.
type ICurrencycrisisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurrencycrisisContext differentiates from other interfaces.
	IsCurrencycrisisContext()
}

type CurrencycrisisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurrencycrisisContext() *CurrencycrisisContext {
	var p = new(CurrencycrisisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CurrencyCrisisParserRULE_currencycrisis
	return p
}

func (*CurrencycrisisContext) IsCurrencycrisisContext() {}

func NewCurrencycrisisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurrencycrisisContext {
	var p = new(CurrencycrisisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CurrencyCrisisParserRULE_currencycrisis

	return p
}

func (s *CurrencycrisisContext) GetParser() antlr.Parser { return s.parser }

func (s *CurrencycrisisContext) CURRENCYCRISIS() antlr.TerminalNode {
	return s.GetToken(CurrencyCrisisParserCURRENCYCRISIS, 0)
}

func (s *CurrencycrisisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurrencycrisisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurrencycrisisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyCrisisListener); ok {
		listenerT.EnterCurrencycrisis(s)
	}
}

func (s *CurrencycrisisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CurrencyCrisisListener); ok {
		listenerT.ExitCurrencycrisis(s)
	}
}

func (p *CurrencyCrisisParser) Currencycrisis() (localctx ICurrencycrisisContext) {
	this := p
	_ = this

	localctx = NewCurrencycrisisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CurrencyCrisisParserRULE_currencycrisis)

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
		p.Match(CurrencyCrisisParserCURRENCYCRISIS)
	}

	return localctx
}
