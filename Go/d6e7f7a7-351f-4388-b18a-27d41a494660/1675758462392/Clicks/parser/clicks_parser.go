// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758462392/Clicks.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clicks

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

type ClicksParser struct {
	*antlr.BaseParser
}

var clicksParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clicksParserInit() {
	staticData := &clicksParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLICKS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "clicks",
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

// ClicksParserInit initializes any static state used to implement ClicksParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClicksParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClicksParserInit() {
	staticData := &clicksParserStaticData
	staticData.once.Do(clicksParserInit)
}

// NewClicksParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClicksParser(input antlr.TokenStream) *ClicksParser {
	ClicksParserInit()
	this := new(ClicksParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clicksParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Clicks.g4"

	return this
}

// ClicksParser tokens.
const (
	ClicksParserEOF      = antlr.TokenEOF
	ClicksParserCLICKS   = 1
	ClicksParserOWNKEY   = 2
	ClicksParserSPLITKEY = 3
	ClicksParserWS       = 4
)

// ClicksParser rules.
const (
	ClicksParserRULE_expression = 0
	ClicksParserRULE_clicks     = 1
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
	p.RuleIndex = ClicksParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClicksParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Clicks() IClicksContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClicksContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClicksContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClicksParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClicksListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClicksListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ClicksParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClicksParserRULE_expression)

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
		p.Clicks()
	}
	{
		p.SetState(5)
		p.Match(ClicksParserEOF)
	}

	return localctx
}

// IClicksContext is an interface to support dynamic dispatch.
type IClicksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClicksContext differentiates from other interfaces.
	IsClicksContext()
}

type ClicksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClicksContext() *ClicksContext {
	var p = new(ClicksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClicksParserRULE_clicks
	return p
}

func (*ClicksContext) IsClicksContext() {}

func NewClicksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClicksContext {
	var p = new(ClicksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClicksParserRULE_clicks

	return p
}

func (s *ClicksContext) GetParser() antlr.Parser { return s.parser }

func (s *ClicksContext) CLICKS() antlr.TerminalNode {
	return s.GetToken(ClicksParserCLICKS, 0)
}

func (s *ClicksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClicksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClicksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClicksListener); ok {
		listenerT.EnterClicks(s)
	}
}

func (s *ClicksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClicksListener); ok {
		listenerT.ExitClicks(s)
	}
}

func (p *ClicksParser) Clicks() (localctx IClicksContext) {
	this := p
	_ = this

	localctx = NewClicksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClicksParserRULE_clicks)

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
		p.Match(ClicksParserCLICKS)
	}

	return localctx
}
