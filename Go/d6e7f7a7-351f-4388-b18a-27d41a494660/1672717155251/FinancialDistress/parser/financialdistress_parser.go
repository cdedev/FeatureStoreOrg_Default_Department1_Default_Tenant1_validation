// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672717155251/FinancialDistress.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FinancialDistress

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

type FinancialDistressParser struct {
	*antlr.BaseParser
}

var financialdistressParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func financialdistressParserInit() {
	staticData := &financialdistressParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FINANCIALDISTRESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "financialdistress",
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

// FinancialDistressParserInit initializes any static state used to implement FinancialDistressParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFinancialDistressParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FinancialDistressParserInit() {
	staticData := &financialdistressParserStaticData
	staticData.once.Do(financialdistressParserInit)
}

// NewFinancialDistressParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFinancialDistressParser(input antlr.TokenStream) *FinancialDistressParser {
	FinancialDistressParserInit()
	this := new(FinancialDistressParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &financialdistressParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FinancialDistress.g4"

	return this
}

// FinancialDistressParser tokens.
const (
	FinancialDistressParserEOF               = antlr.TokenEOF
	FinancialDistressParserFINANCIALDISTRESS = 1
	FinancialDistressParserOWNKEY            = 2
	FinancialDistressParserSPLITKEY          = 3
	FinancialDistressParserWS                = 4
)

// FinancialDistressParser rules.
const (
	FinancialDistressParserRULE_expression        = 0
	FinancialDistressParserRULE_financialdistress = 1
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
	p.RuleIndex = FinancialDistressParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinancialDistressParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Financialdistress() IFinancialdistressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFinancialdistressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFinancialdistressContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FinancialDistressParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinancialDistressListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinancialDistressListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FinancialDistressParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FinancialDistressParserRULE_expression)

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
		p.Financialdistress()
	}
	{
		p.SetState(5)
		p.Match(FinancialDistressParserEOF)
	}

	return localctx
}

// IFinancialdistressContext is an interface to support dynamic dispatch.
type IFinancialdistressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFinancialdistressContext differentiates from other interfaces.
	IsFinancialdistressContext()
}

type FinancialdistressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinancialdistressContext() *FinancialdistressContext {
	var p = new(FinancialdistressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FinancialDistressParserRULE_financialdistress
	return p
}

func (*FinancialdistressContext) IsFinancialdistressContext() {}

func NewFinancialdistressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinancialdistressContext {
	var p = new(FinancialdistressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FinancialDistressParserRULE_financialdistress

	return p
}

func (s *FinancialdistressContext) GetParser() antlr.Parser { return s.parser }

func (s *FinancialdistressContext) FINANCIALDISTRESS() antlr.TerminalNode {
	return s.GetToken(FinancialDistressParserFINANCIALDISTRESS, 0)
}

func (s *FinancialdistressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinancialdistressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinancialdistressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinancialDistressListener); ok {
		listenerT.EnterFinancialdistress(s)
	}
}

func (s *FinancialdistressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FinancialDistressListener); ok {
		listenerT.ExitFinancialdistress(s)
	}
}

func (p *FinancialDistressParser) Financialdistress() (localctx IFinancialdistressContext) {
	this := p
	_ = this

	localctx = NewFinancialdistressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FinancialDistressParserRULE_financialdistress)

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
		p.Match(FinancialDistressParserFINANCIALDISTRESS)
	}

	return localctx
}
