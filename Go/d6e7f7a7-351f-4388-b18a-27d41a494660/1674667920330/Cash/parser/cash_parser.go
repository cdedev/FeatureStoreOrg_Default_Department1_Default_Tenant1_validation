// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674667920330/Cash.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cash

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

type CashParser struct {
	*antlr.BaseParser
}

var cashParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cashParserInit() {
	staticData := &cashParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CASH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cash",
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

// CashParserInit initializes any static state used to implement CashParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCashParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CashParserInit() {
	staticData := &cashParserStaticData
	staticData.once.Do(cashParserInit)
}

// NewCashParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCashParser(input antlr.TokenStream) *CashParser {
	CashParserInit()
	this := new(CashParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cashParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cash.g4"

	return this
}

// CashParser tokens.
const (
	CashParserEOF      = antlr.TokenEOF
	CashParserCASH     = 1
	CashParserOWNKEY   = 2
	CashParserSPLITKEY = 3
	CashParserWS       = 4
)

// CashParser rules.
const (
	CashParserRULE_expression = 0
	CashParserRULE_cash       = 1
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
	p.RuleIndex = CashParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CashParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cash() ICashContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICashContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICashContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CashParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CashListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CashListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CashParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CashParserRULE_expression)

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
		p.Cash()
	}
	{
		p.SetState(5)
		p.Match(CashParserEOF)
	}

	return localctx
}

// ICashContext is an interface to support dynamic dispatch.
type ICashContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCashContext differentiates from other interfaces.
	IsCashContext()
}

type CashContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCashContext() *CashContext {
	var p = new(CashContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CashParserRULE_cash
	return p
}

func (*CashContext) IsCashContext() {}

func NewCashContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CashContext {
	var p = new(CashContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CashParserRULE_cash

	return p
}

func (s *CashContext) GetParser() antlr.Parser { return s.parser }

func (s *CashContext) CASH() antlr.TerminalNode {
	return s.GetToken(CashParserCASH, 0)
}

func (s *CashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CashContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CashListener); ok {
		listenerT.EnterCash(s)
	}
}

func (s *CashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CashListener); ok {
		listenerT.ExitCash(s)
	}
}

func (p *CashParser) Cash() (localctx ICashContext) {
	this := p
	_ = this

	localctx = NewCashContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CashParserRULE_cash)

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
		p.Match(CashParserCASH)
	}

	return localctx
}
