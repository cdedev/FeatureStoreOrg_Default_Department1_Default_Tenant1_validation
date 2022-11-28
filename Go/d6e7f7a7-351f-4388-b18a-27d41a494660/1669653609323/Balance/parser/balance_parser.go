// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653609323/Balance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Balance

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

type BalanceParser struct {
	*antlr.BaseParser
}

var balanceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func balanceParserInit() {
	staticData := &balanceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BALANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "balance",
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

// BalanceParserInit initializes any static state used to implement BalanceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBalanceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BalanceParserInit() {
	staticData := &balanceParserStaticData
	staticData.once.Do(balanceParserInit)
}

// NewBalanceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBalanceParser(input antlr.TokenStream) *BalanceParser {
	BalanceParserInit()
	this := new(BalanceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &balanceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Balance.g4"

	return this
}

// BalanceParser tokens.
const (
	BalanceParserEOF      = antlr.TokenEOF
	BalanceParserBALANCE  = 1
	BalanceParserOWNKEY   = 2
	BalanceParserSPLITKEY = 3
	BalanceParserWS       = 4
)

// BalanceParser rules.
const (
	BalanceParserRULE_expression = 0
	BalanceParserRULE_balance    = 1
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
	p.RuleIndex = BalanceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BalanceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Balance() IBalanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBalanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBalanceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BalanceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BalanceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BalanceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BalanceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BalanceParserRULE_expression)

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
		p.Balance()
	}
	{
		p.SetState(5)
		p.Match(BalanceParserEOF)
	}

	return localctx
}

// IBalanceContext is an interface to support dynamic dispatch.
type IBalanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBalanceContext differentiates from other interfaces.
	IsBalanceContext()
}

type BalanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBalanceContext() *BalanceContext {
	var p = new(BalanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BalanceParserRULE_balance
	return p
}

func (*BalanceContext) IsBalanceContext() {}

func NewBalanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BalanceContext {
	var p = new(BalanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BalanceParserRULE_balance

	return p
}

func (s *BalanceContext) GetParser() antlr.Parser { return s.parser }

func (s *BalanceContext) BALANCE() antlr.TerminalNode {
	return s.GetToken(BalanceParserBALANCE, 0)
}

func (s *BalanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BalanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BalanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BalanceListener); ok {
		listenerT.EnterBalance(s)
	}
}

func (s *BalanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BalanceListener); ok {
		listenerT.ExitBalance(s)
	}
}

func (p *BalanceParser) Balance() (localctx IBalanceContext) {
	this := p
	_ = this

	localctx = NewBalanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BalanceParserRULE_balance)

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
		p.Match(BalanceParserBALANCE)
	}

	return localctx
}
