// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675673022752/Correlation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Correlation

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

type CorrelationParser struct {
	*antlr.BaseParser
}

var correlationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func correlationParserInit() {
	staticData := &correlationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CORRELATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "correlation",
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

// CorrelationParserInit initializes any static state used to implement CorrelationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCorrelationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CorrelationParserInit() {
	staticData := &correlationParserStaticData
	staticData.once.Do(correlationParserInit)
}

// NewCorrelationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCorrelationParser(input antlr.TokenStream) *CorrelationParser {
	CorrelationParserInit()
	this := new(CorrelationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &correlationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Correlation.g4"

	return this
}

// CorrelationParser tokens.
const (
	CorrelationParserEOF         = antlr.TokenEOF
	CorrelationParserCORRELATION = 1
	CorrelationParserOWNKEY      = 2
	CorrelationParserSPLITKEY    = 3
	CorrelationParserWS          = 4
)

// CorrelationParser rules.
const (
	CorrelationParserRULE_expression  = 0
	CorrelationParserRULE_correlation = 1
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
	p.RuleIndex = CorrelationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CorrelationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Correlation() ICorrelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICorrelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICorrelationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CorrelationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CorrelationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CorrelationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CorrelationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CorrelationParserRULE_expression)

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
		p.Correlation()
	}
	{
		p.SetState(5)
		p.Match(CorrelationParserEOF)
	}

	return localctx
}

// ICorrelationContext is an interface to support dynamic dispatch.
type ICorrelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCorrelationContext differentiates from other interfaces.
	IsCorrelationContext()
}

type CorrelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCorrelationContext() *CorrelationContext {
	var p = new(CorrelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CorrelationParserRULE_correlation
	return p
}

func (*CorrelationContext) IsCorrelationContext() {}

func NewCorrelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CorrelationContext {
	var p = new(CorrelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CorrelationParserRULE_correlation

	return p
}

func (s *CorrelationContext) GetParser() antlr.Parser { return s.parser }

func (s *CorrelationContext) CORRELATION() antlr.TerminalNode {
	return s.GetToken(CorrelationParserCORRELATION, 0)
}

func (s *CorrelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CorrelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CorrelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CorrelationListener); ok {
		listenerT.EnterCorrelation(s)
	}
}

func (s *CorrelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CorrelationListener); ok {
		listenerT.ExitCorrelation(s)
	}
}

func (p *CorrelationParser) Correlation() (localctx ICorrelationContext) {
	this := p
	_ = this

	localctx = NewCorrelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CorrelationParserRULE_correlation)

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
		p.Match(CorrelationParserCORRELATION)
	}

	return localctx
}
