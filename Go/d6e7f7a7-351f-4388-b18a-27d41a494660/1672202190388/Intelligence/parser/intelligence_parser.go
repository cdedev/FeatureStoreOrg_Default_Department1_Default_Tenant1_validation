// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202190388/Intelligence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Intelligence

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

type IntelligenceParser struct {
	*antlr.BaseParser
}

var intelligenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func intelligenceParserInit() {
	staticData := &intelligenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INTELLIGENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "intelligence",
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

// IntelligenceParserInit initializes any static state used to implement IntelligenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIntelligenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IntelligenceParserInit() {
	staticData := &intelligenceParserStaticData
	staticData.once.Do(intelligenceParserInit)
}

// NewIntelligenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIntelligenceParser(input antlr.TokenStream) *IntelligenceParser {
	IntelligenceParserInit()
	this := new(IntelligenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &intelligenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Intelligence.g4"

	return this
}

// IntelligenceParser tokens.
const (
	IntelligenceParserEOF          = antlr.TokenEOF
	IntelligenceParserINTELLIGENCE = 1
	IntelligenceParserOWNKEY       = 2
	IntelligenceParserSPLITKEY     = 3
	IntelligenceParserWS           = 4
)

// IntelligenceParser rules.
const (
	IntelligenceParserRULE_expression   = 0
	IntelligenceParserRULE_intelligence = 1
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
	p.RuleIndex = IntelligenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntelligenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Intelligence() IIntelligenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntelligenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntelligenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IntelligenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntelligenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntelligenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IntelligenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IntelligenceParserRULE_expression)

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
		p.Intelligence()
	}
	{
		p.SetState(5)
		p.Match(IntelligenceParserEOF)
	}

	return localctx
}

// IIntelligenceContext is an interface to support dynamic dispatch.
type IIntelligenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntelligenceContext differentiates from other interfaces.
	IsIntelligenceContext()
}

type IntelligenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntelligenceContext() *IntelligenceContext {
	var p = new(IntelligenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IntelligenceParserRULE_intelligence
	return p
}

func (*IntelligenceContext) IsIntelligenceContext() {}

func NewIntelligenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntelligenceContext {
	var p = new(IntelligenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntelligenceParserRULE_intelligence

	return p
}

func (s *IntelligenceContext) GetParser() antlr.Parser { return s.parser }

func (s *IntelligenceContext) INTELLIGENCE() antlr.TerminalNode {
	return s.GetToken(IntelligenceParserINTELLIGENCE, 0)
}

func (s *IntelligenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntelligenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntelligenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntelligenceListener); ok {
		listenerT.EnterIntelligence(s)
	}
}

func (s *IntelligenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntelligenceListener); ok {
		listenerT.ExitIntelligence(s)
	}
}

func (p *IntelligenceParser) Intelligence() (localctx IIntelligenceContext) {
	this := p
	_ = this

	localctx = NewIntelligenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IntelligenceParserRULE_intelligence)

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
		p.Match(IntelligenceParserINTELLIGENCE)
	}

	return localctx
}
