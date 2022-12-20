// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671526013217/Income.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Income

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

type IncomeParser struct {
	*antlr.BaseParser
}

var incomeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func incomeParserInit() {
	staticData := &incomeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INCOME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "income",
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

// IncomeParserInit initializes any static state used to implement IncomeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIncomeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IncomeParserInit() {
	staticData := &incomeParserStaticData
	staticData.once.Do(incomeParserInit)
}

// NewIncomeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIncomeParser(input antlr.TokenStream) *IncomeParser {
	IncomeParserInit()
	this := new(IncomeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &incomeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Income.g4"

	return this
}

// IncomeParser tokens.
const (
	IncomeParserEOF      = antlr.TokenEOF
	IncomeParserINCOME   = 1
	IncomeParserOWNKEY   = 2
	IncomeParserSPLITKEY = 3
	IncomeParserWS       = 4
)

// IncomeParser rules.
const (
	IncomeParserRULE_expression = 0
	IncomeParserRULE_income     = 1
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
	p.RuleIndex = IncomeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IncomeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Income() IIncomeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncomeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncomeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IncomeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IncomeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IncomeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IncomeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IncomeParserRULE_expression)

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
		p.Income()
	}
	{
		p.SetState(5)
		p.Match(IncomeParserEOF)
	}

	return localctx
}

// IIncomeContext is an interface to support dynamic dispatch.
type IIncomeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncomeContext differentiates from other interfaces.
	IsIncomeContext()
}

type IncomeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncomeContext() *IncomeContext {
	var p = new(IncomeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IncomeParserRULE_income
	return p
}

func (*IncomeContext) IsIncomeContext() {}

func NewIncomeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncomeContext {
	var p = new(IncomeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IncomeParserRULE_income

	return p
}

func (s *IncomeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncomeContext) INCOME() antlr.TerminalNode {
	return s.GetToken(IncomeParserINCOME, 0)
}

func (s *IncomeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncomeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncomeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IncomeListener); ok {
		listenerT.EnterIncome(s)
	}
}

func (s *IncomeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IncomeListener); ok {
		listenerT.ExitIncome(s)
	}
}

func (p *IncomeParser) Income() (localctx IIncomeContext) {
	this := p
	_ = this

	localctx = NewIncomeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IncomeParserRULE_income)

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
		p.Match(IncomeParserINCOME)
	}

	return localctx
}
