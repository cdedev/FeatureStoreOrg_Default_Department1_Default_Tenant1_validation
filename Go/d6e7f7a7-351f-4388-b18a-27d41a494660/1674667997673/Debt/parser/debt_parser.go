// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674667997673/Debt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Debt

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

type DebtParser struct {
	*antlr.BaseParser
}

var debtParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func debtParserInit() {
	staticData := &debtParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DEBT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "debt",
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

// DebtParserInit initializes any static state used to implement DebtParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDebtParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DebtParserInit() {
	staticData := &debtParserStaticData
	staticData.once.Do(debtParserInit)
}

// NewDebtParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDebtParser(input antlr.TokenStream) *DebtParser {
	DebtParserInit()
	this := new(DebtParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &debtParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Debt.g4"

	return this
}

// DebtParser tokens.
const (
	DebtParserEOF      = antlr.TokenEOF
	DebtParserDEBT     = 1
	DebtParserOWNKEY   = 2
	DebtParserSPLITKEY = 3
	DebtParserWS       = 4
)

// DebtParser rules.
const (
	DebtParserRULE_expression = 0
	DebtParserRULE_debt       = 1
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
	p.RuleIndex = DebtParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DebtParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Debt() IDebtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDebtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDebtContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DebtParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebtListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebtListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DebtParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DebtParserRULE_expression)

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
		p.Debt()
	}
	{
		p.SetState(5)
		p.Match(DebtParserEOF)
	}

	return localctx
}

// IDebtContext is an interface to support dynamic dispatch.
type IDebtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebtContext differentiates from other interfaces.
	IsDebtContext()
}

type DebtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebtContext() *DebtContext {
	var p = new(DebtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DebtParserRULE_debt
	return p
}

func (*DebtContext) IsDebtContext() {}

func NewDebtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebtContext {
	var p = new(DebtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DebtParserRULE_debt

	return p
}

func (s *DebtContext) GetParser() antlr.Parser { return s.parser }

func (s *DebtContext) DEBT() antlr.TerminalNode {
	return s.GetToken(DebtParserDEBT, 0)
}

func (s *DebtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebtListener); ok {
		listenerT.EnterDebt(s)
	}
}

func (s *DebtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebtListener); ok {
		listenerT.ExitDebt(s)
	}
}

func (p *DebtParser) Debt() (localctx IDebtContext) {
	this := p
	_ = this

	localctx = NewDebtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DebtParserRULE_debt)

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
		p.Match(DebtParserDEBT)
	}

	return localctx
}
