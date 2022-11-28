// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669629685893/Churn.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Churn

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

type ChurnParser struct {
	*antlr.BaseParser
}

var churnParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func churnParserInit() {
	staticData := &churnParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHURN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "churn",
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

// ChurnParserInit initializes any static state used to implement ChurnParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewChurnParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChurnParserInit() {
	staticData := &churnParserStaticData
	staticData.once.Do(churnParserInit)
}

// NewChurnParser produces a new parser instance for the optional input antlr.TokenStream.
func NewChurnParser(input antlr.TokenStream) *ChurnParser {
	ChurnParserInit()
	this := new(ChurnParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &churnParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Churn.g4"

	return this
}

// ChurnParser tokens.
const (
	ChurnParserEOF      = antlr.TokenEOF
	ChurnParserCHURN    = 1
	ChurnParserOWNKEY   = 2
	ChurnParserSPLITKEY = 3
	ChurnParserWS       = 4
)

// ChurnParser rules.
const (
	ChurnParserRULE_expression = 0
	ChurnParserRULE_churn      = 1
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
	p.RuleIndex = ChurnParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChurnParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Churn() IChurnContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChurnContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChurnContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChurnParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChurnListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChurnListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ChurnParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChurnParserRULE_expression)

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
		p.Churn()
	}
	{
		p.SetState(5)
		p.Match(ChurnParserEOF)
	}

	return localctx
}

// IChurnContext is an interface to support dynamic dispatch.
type IChurnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChurnContext differentiates from other interfaces.
	IsChurnContext()
}

type ChurnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChurnContext() *ChurnContext {
	var p = new(ChurnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChurnParserRULE_churn
	return p
}

func (*ChurnContext) IsChurnContext() {}

func NewChurnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChurnContext {
	var p = new(ChurnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChurnParserRULE_churn

	return p
}

func (s *ChurnContext) GetParser() antlr.Parser { return s.parser }

func (s *ChurnContext) CHURN() antlr.TerminalNode {
	return s.GetToken(ChurnParserCHURN, 0)
}

func (s *ChurnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChurnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChurnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChurnListener); ok {
		listenerT.EnterChurn(s)
	}
}

func (s *ChurnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChurnListener); ok {
		listenerT.ExitChurn(s)
	}
}

func (p *ChurnParser) Churn() (localctx IChurnContext) {
	this := p
	_ = this

	localctx = NewChurnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChurnParserRULE_churn)

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
		p.Match(ChurnParserCHURN)
	}

	return localctx
}
