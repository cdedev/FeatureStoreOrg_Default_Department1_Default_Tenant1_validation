// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638016123/TransTotal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TransTotal

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

type TransTotalParser struct {
	*antlr.BaseParser
}

var transtotalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func transtotalParserInit() {
	staticData := &transtotalParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TRANSTOTAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "transtotal",
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

// TransTotalParserInit initializes any static state used to implement TransTotalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTransTotalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TransTotalParserInit() {
	staticData := &transtotalParserStaticData
	staticData.once.Do(transtotalParserInit)
}

// NewTransTotalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTransTotalParser(input antlr.TokenStream) *TransTotalParser {
	TransTotalParserInit()
	this := new(TransTotalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &transtotalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "TransTotal.g4"

	return this
}

// TransTotalParser tokens.
const (
	TransTotalParserEOF        = antlr.TokenEOF
	TransTotalParserTRANSTOTAL = 1
	TransTotalParserOWNKEY     = 2
	TransTotalParserSPLITKEY   = 3
	TransTotalParserWS         = 4
)

// TransTotalParser rules.
const (
	TransTotalParserRULE_expression = 0
	TransTotalParserRULE_transtotal = 1
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
	p.RuleIndex = TransTotalParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransTotalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Transtotal() ITranstotalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITranstotalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITranstotalContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TransTotalParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransTotalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransTotalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TransTotalParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TransTotalParserRULE_expression)

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
		p.Transtotal()
	}
	{
		p.SetState(5)
		p.Match(TransTotalParserEOF)
	}

	return localctx
}

// ITranstotalContext is an interface to support dynamic dispatch.
type ITranstotalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranstotalContext differentiates from other interfaces.
	IsTranstotalContext()
}

type TranstotalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranstotalContext() *TranstotalContext {
	var p = new(TranstotalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TransTotalParserRULE_transtotal
	return p
}

func (*TranstotalContext) IsTranstotalContext() {}

func NewTranstotalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranstotalContext {
	var p = new(TranstotalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TransTotalParserRULE_transtotal

	return p
}

func (s *TranstotalContext) GetParser() antlr.Parser { return s.parser }

func (s *TranstotalContext) TRANSTOTAL() antlr.TerminalNode {
	return s.GetToken(TransTotalParserTRANSTOTAL, 0)
}

func (s *TranstotalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranstotalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranstotalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransTotalListener); ok {
		listenerT.EnterTranstotal(s)
	}
}

func (s *TranstotalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TransTotalListener); ok {
		listenerT.ExitTranstotal(s)
	}
}

func (p *TransTotalParser) Transtotal() (localctx ITranstotalContext) {
	this := p
	_ = this

	localctx = NewTranstotalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TransTotalParserRULE_transtotal)

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
		p.Match(TransTotalParserTRANSTOTAL)
	}

	return localctx
}
