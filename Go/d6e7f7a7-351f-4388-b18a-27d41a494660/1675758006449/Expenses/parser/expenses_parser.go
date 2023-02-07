// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758006449/Expenses.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expenses

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

type ExpensesParser struct {
	*antlr.BaseParser
}

var expensesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expensesParserInit() {
	staticData := &expensesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EXPENSES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "expenses",
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

// ExpensesParserInit initializes any static state used to implement ExpensesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpensesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpensesParserInit() {
	staticData := &expensesParserStaticData
	staticData.once.Do(expensesParserInit)
}

// NewExpensesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpensesParser(input antlr.TokenStream) *ExpensesParser {
	ExpensesParserInit()
	this := new(ExpensesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &expensesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Expenses.g4"

	return this
}

// ExpensesParser tokens.
const (
	ExpensesParserEOF      = antlr.TokenEOF
	ExpensesParserEXPENSES = 1
	ExpensesParserOWNKEY   = 2
	ExpensesParserSPLITKEY = 3
	ExpensesParserWS       = 4
)

// ExpensesParser rules.
const (
	ExpensesParserRULE_expression = 0
	ExpensesParserRULE_expenses   = 1
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
	p.RuleIndex = ExpensesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpensesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Expenses() IExpensesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpensesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpensesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpensesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpensesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpensesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExpensesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpensesParserRULE_expression)

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
		p.Expenses()
	}
	{
		p.SetState(5)
		p.Match(ExpensesParserEOF)
	}

	return localctx
}

// IExpensesContext is an interface to support dynamic dispatch.
type IExpensesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpensesContext differentiates from other interfaces.
	IsExpensesContext()
}

type ExpensesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpensesContext() *ExpensesContext {
	var p = new(ExpensesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpensesParserRULE_expenses
	return p
}

func (*ExpensesContext) IsExpensesContext() {}

func NewExpensesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpensesContext {
	var p = new(ExpensesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpensesParserRULE_expenses

	return p
}

func (s *ExpensesContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpensesContext) EXPENSES() antlr.TerminalNode {
	return s.GetToken(ExpensesParserEXPENSES, 0)
}

func (s *ExpensesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpensesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpensesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpensesListener); ok {
		listenerT.EnterExpenses(s)
	}
}

func (s *ExpensesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpensesListener); ok {
		listenerT.ExitExpenses(s)
	}
}

func (p *ExpensesParser) Expenses() (localctx IExpensesContext) {
	this := p
	_ = this

	localctx = NewExpensesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpensesParserRULE_expenses)

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
		p.Match(ExpensesParserEXPENSES)
	}

	return localctx
}
