// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528831254/LoanTerm.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LoanTerm

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

type LoanTermParser struct {
	*antlr.BaseParser
}

var loantermParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func loantermParserInit() {
	staticData := &loantermParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LOANTERM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "loanterm",
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

// LoanTermParserInit initializes any static state used to implement LoanTermParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLoanTermParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LoanTermParserInit() {
	staticData := &loantermParserStaticData
	staticData.once.Do(loantermParserInit)
}

// NewLoanTermParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLoanTermParser(input antlr.TokenStream) *LoanTermParser {
	LoanTermParserInit()
	this := new(LoanTermParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &loantermParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "LoanTerm.g4"

	return this
}

// LoanTermParser tokens.
const (
	LoanTermParserEOF      = antlr.TokenEOF
	LoanTermParserLOANTERM = 1
	LoanTermParserOWNKEY   = 2
	LoanTermParserSPLITKEY = 3
	LoanTermParserWS       = 4
)

// LoanTermParser rules.
const (
	LoanTermParserRULE_expression = 0
	LoanTermParserRULE_loanterm   = 1
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
	p.RuleIndex = LoanTermParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoanTermParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Loanterm() ILoantermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoantermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoantermContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LoanTermParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoanTermListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoanTermListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LoanTermParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LoanTermParserRULE_expression)

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
		p.Loanterm()
	}
	{
		p.SetState(5)
		p.Match(LoanTermParserEOF)
	}

	return localctx
}

// ILoantermContext is an interface to support dynamic dispatch.
type ILoantermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoantermContext differentiates from other interfaces.
	IsLoantermContext()
}

type LoantermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoantermContext() *LoantermContext {
	var p = new(LoantermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoanTermParserRULE_loanterm
	return p
}

func (*LoantermContext) IsLoantermContext() {}

func NewLoantermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoantermContext {
	var p = new(LoantermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoanTermParserRULE_loanterm

	return p
}

func (s *LoantermContext) GetParser() antlr.Parser { return s.parser }

func (s *LoantermContext) LOANTERM() antlr.TerminalNode {
	return s.GetToken(LoanTermParserLOANTERM, 0)
}

func (s *LoantermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoantermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoantermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoanTermListener); ok {
		listenerT.EnterLoanterm(s)
	}
}

func (s *LoantermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoanTermListener); ok {
		listenerT.ExitLoanterm(s)
	}
}

func (p *LoanTermParser) Loanterm() (localctx ILoantermContext) {
	this := p
	_ = this

	localctx = NewLoantermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LoanTermParserRULE_loanterm)

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
		p.Match(LoanTermParserLOANTERM)
	}

	return localctx
}
