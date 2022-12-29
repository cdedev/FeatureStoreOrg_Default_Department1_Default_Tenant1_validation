// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672285713110/RevisedAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RevisedAmount

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

type RevisedAmountParser struct {
	*antlr.BaseParser
}

var revisedamountParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func revisedamountParserInit() {
	staticData := &revisedamountParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "REVISEDAMOUNT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "revisedamount",
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

// RevisedAmountParserInit initializes any static state used to implement RevisedAmountParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRevisedAmountParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RevisedAmountParserInit() {
	staticData := &revisedamountParserStaticData
	staticData.once.Do(revisedamountParserInit)
}

// NewRevisedAmountParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRevisedAmountParser(input antlr.TokenStream) *RevisedAmountParser {
	RevisedAmountParserInit()
	this := new(RevisedAmountParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &revisedamountParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RevisedAmount.g4"

	return this
}

// RevisedAmountParser tokens.
const (
	RevisedAmountParserEOF           = antlr.TokenEOF
	RevisedAmountParserREVISEDAMOUNT = 1
	RevisedAmountParserOWNKEY        = 2
	RevisedAmountParserSPLITKEY      = 3
	RevisedAmountParserWS            = 4
)

// RevisedAmountParser rules.
const (
	RevisedAmountParserRULE_expression    = 0
	RevisedAmountParserRULE_revisedamount = 1
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
	p.RuleIndex = RevisedAmountParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RevisedAmountParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Revisedamount() IRevisedamountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRevisedamountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRevisedamountContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RevisedAmountParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RevisedAmountListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RevisedAmountListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RevisedAmountParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RevisedAmountParserRULE_expression)

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
		p.Revisedamount()
	}
	{
		p.SetState(5)
		p.Match(RevisedAmountParserEOF)
	}

	return localctx
}

// IRevisedamountContext is an interface to support dynamic dispatch.
type IRevisedamountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRevisedamountContext differentiates from other interfaces.
	IsRevisedamountContext()
}

type RevisedamountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRevisedamountContext() *RevisedamountContext {
	var p = new(RevisedamountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RevisedAmountParserRULE_revisedamount
	return p
}

func (*RevisedamountContext) IsRevisedamountContext() {}

func NewRevisedamountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RevisedamountContext {
	var p = new(RevisedamountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RevisedAmountParserRULE_revisedamount

	return p
}

func (s *RevisedamountContext) GetParser() antlr.Parser { return s.parser }

func (s *RevisedamountContext) REVISEDAMOUNT() antlr.TerminalNode {
	return s.GetToken(RevisedAmountParserREVISEDAMOUNT, 0)
}

func (s *RevisedamountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RevisedamountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RevisedamountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RevisedAmountListener); ok {
		listenerT.EnterRevisedamount(s)
	}
}

func (s *RevisedamountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RevisedAmountListener); ok {
		listenerT.ExitRevisedamount(s)
	}
}

func (p *RevisedAmountParser) Revisedamount() (localctx IRevisedamountContext) {
	this := p
	_ = this

	localctx = NewRevisedamountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RevisedAmountParserRULE_revisedamount)

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
		p.Match(RevisedAmountParserREVISEDAMOUNT)
	}

	return localctx
}
