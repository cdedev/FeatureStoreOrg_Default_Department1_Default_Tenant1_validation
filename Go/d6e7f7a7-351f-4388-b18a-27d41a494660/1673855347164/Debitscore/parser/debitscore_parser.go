// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673855347164/Debitscore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Debitscore

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

type DebitscoreParser struct {
	*antlr.BaseParser
}

var debitscoreParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func debitscoreParserInit() {
	staticData := &debitscoreParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DEBITSCORE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "debitscore",
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

// DebitscoreParserInit initializes any static state used to implement DebitscoreParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDebitscoreParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DebitscoreParserInit() {
	staticData := &debitscoreParserStaticData
	staticData.once.Do(debitscoreParserInit)
}

// NewDebitscoreParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDebitscoreParser(input antlr.TokenStream) *DebitscoreParser {
	DebitscoreParserInit()
	this := new(DebitscoreParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &debitscoreParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Debitscore.g4"

	return this
}

// DebitscoreParser tokens.
const (
	DebitscoreParserEOF        = antlr.TokenEOF
	DebitscoreParserDEBITSCORE = 1
	DebitscoreParserOWNKEY     = 2
	DebitscoreParserSPLITKEY   = 3
	DebitscoreParserWS         = 4
)

// DebitscoreParser rules.
const (
	DebitscoreParserRULE_expression = 0
	DebitscoreParserRULE_debitscore = 1
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
	p.RuleIndex = DebitscoreParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DebitscoreParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Debitscore() IDebitscoreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDebitscoreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDebitscoreContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DebitscoreParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebitscoreListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebitscoreListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DebitscoreParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DebitscoreParserRULE_expression)

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
		p.Debitscore()
	}
	{
		p.SetState(5)
		p.Match(DebitscoreParserEOF)
	}

	return localctx
}

// IDebitscoreContext is an interface to support dynamic dispatch.
type IDebitscoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebitscoreContext differentiates from other interfaces.
	IsDebitscoreContext()
}

type DebitscoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebitscoreContext() *DebitscoreContext {
	var p = new(DebitscoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DebitscoreParserRULE_debitscore
	return p
}

func (*DebitscoreContext) IsDebitscoreContext() {}

func NewDebitscoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebitscoreContext {
	var p = new(DebitscoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DebitscoreParserRULE_debitscore

	return p
}

func (s *DebitscoreContext) GetParser() antlr.Parser { return s.parser }

func (s *DebitscoreContext) DEBITSCORE() antlr.TerminalNode {
	return s.GetToken(DebitscoreParserDEBITSCORE, 0)
}

func (s *DebitscoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebitscoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebitscoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebitscoreListener); ok {
		listenerT.EnterDebitscore(s)
	}
}

func (s *DebitscoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DebitscoreListener); ok {
		listenerT.ExitDebitscore(s)
	}
}

func (p *DebitscoreParser) Debitscore() (localctx IDebitscoreContext) {
	this := p
	_ = this

	localctx = NewDebitscoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DebitscoreParserRULE_debitscore)

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
		p.Match(DebitscoreParserDEBITSCORE)
	}

	return localctx
}
