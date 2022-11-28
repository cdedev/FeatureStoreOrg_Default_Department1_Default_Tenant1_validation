// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669630209027/Loss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loss

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

type LossParser struct {
	*antlr.BaseParser
}

var lossParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lossParserInit() {
	staticData := &lossParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LOSS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "loss",
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

// LossParserInit initializes any static state used to implement LossParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLossParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LossParserInit() {
	staticData := &lossParserStaticData
	staticData.once.Do(lossParserInit)
}

// NewLossParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLossParser(input antlr.TokenStream) *LossParser {
	LossParserInit()
	this := new(LossParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lossParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Loss.g4"

	return this
}

// LossParser tokens.
const (
	LossParserEOF      = antlr.TokenEOF
	LossParserLOSS     = 1
	LossParserOWNKEY   = 2
	LossParserSPLITKEY = 3
	LossParserWS       = 4
)

// LossParser rules.
const (
	LossParserRULE_expression = 0
	LossParserRULE_loss       = 1
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
	p.RuleIndex = LossParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LossParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Loss() ILossContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILossContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILossContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LossParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LossListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LossListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LossParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LossParserRULE_expression)

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
		p.Loss()
	}
	{
		p.SetState(5)
		p.Match(LossParserEOF)
	}

	return localctx
}

// ILossContext is an interface to support dynamic dispatch.
type ILossContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLossContext differentiates from other interfaces.
	IsLossContext()
}

type LossContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLossContext() *LossContext {
	var p = new(LossContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LossParserRULE_loss
	return p
}

func (*LossContext) IsLossContext() {}

func NewLossContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LossContext {
	var p = new(LossContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LossParserRULE_loss

	return p
}

func (s *LossContext) GetParser() antlr.Parser { return s.parser }

func (s *LossContext) LOSS() antlr.TerminalNode {
	return s.GetToken(LossParserLOSS, 0)
}

func (s *LossContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LossContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LossContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LossListener); ok {
		listenerT.EnterLoss(s)
	}
}

func (s *LossContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LossListener); ok {
		listenerT.ExitLoss(s)
	}
}

func (p *LossParser) Loss() (localctx ILossContext) {
	this := p
	_ = this

	localctx = NewLossContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LossParserRULE_loss)

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
		p.Match(LossParserLOSS)
	}

	return localctx
}
