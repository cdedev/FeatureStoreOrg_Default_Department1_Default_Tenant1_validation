// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669973257050/Gross_Amount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gross_Amount

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

type Gross_AmountParser struct {
	*antlr.BaseParser
}

var gross_amountParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gross_amountParserInit() {
	staticData := &gross_amountParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GROSS_AMOUNT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gross_amount",
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

// Gross_AmountParserInit initializes any static state used to implement Gross_AmountParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGross_AmountParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Gross_AmountParserInit() {
	staticData := &gross_amountParserStaticData
	staticData.once.Do(gross_amountParserInit)
}

// NewGross_AmountParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGross_AmountParser(input antlr.TokenStream) *Gross_AmountParser {
	Gross_AmountParserInit()
	this := new(Gross_AmountParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gross_amountParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Gross_Amount.g4"

	return this
}

// Gross_AmountParser tokens.
const (
	Gross_AmountParserEOF          = antlr.TokenEOF
	Gross_AmountParserGROSS_AMOUNT = 1
	Gross_AmountParserOWNKEY       = 2
	Gross_AmountParserSPLITKEY     = 3
	Gross_AmountParserWS           = 4
)

// Gross_AmountParser rules.
const (
	Gross_AmountParserRULE_expression   = 0
	Gross_AmountParserRULE_gross_amount = 1
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
	p.RuleIndex = Gross_AmountParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Gross_AmountParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gross_amount() IGross_amountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGross_amountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGross_amountContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Gross_AmountParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Gross_AmountListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Gross_AmountListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Gross_AmountParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Gross_AmountParserRULE_expression)

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
		p.Gross_amount()
	}
	{
		p.SetState(5)
		p.Match(Gross_AmountParserEOF)
	}

	return localctx
}

// IGross_amountContext is an interface to support dynamic dispatch.
type IGross_amountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGross_amountContext differentiates from other interfaces.
	IsGross_amountContext()
}

type Gross_amountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGross_amountContext() *Gross_amountContext {
	var p = new(Gross_amountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Gross_AmountParserRULE_gross_amount
	return p
}

func (*Gross_amountContext) IsGross_amountContext() {}

func NewGross_amountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Gross_amountContext {
	var p = new(Gross_amountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Gross_AmountParserRULE_gross_amount

	return p
}

func (s *Gross_amountContext) GetParser() antlr.Parser { return s.parser }

func (s *Gross_amountContext) GROSS_AMOUNT() antlr.TerminalNode {
	return s.GetToken(Gross_AmountParserGROSS_AMOUNT, 0)
}

func (s *Gross_amountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Gross_amountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Gross_amountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Gross_AmountListener); ok {
		listenerT.EnterGross_amount(s)
	}
}

func (s *Gross_amountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Gross_AmountListener); ok {
		listenerT.ExitGross_amount(s)
	}
}

func (p *Gross_AmountParser) Gross_amount() (localctx IGross_amountContext) {
	this := p
	_ = this

	localctx = NewGross_amountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Gross_AmountParserRULE_gross_amount)

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
		p.Match(Gross_AmountParserGROSS_AMOUNT)
	}

	return localctx
}
