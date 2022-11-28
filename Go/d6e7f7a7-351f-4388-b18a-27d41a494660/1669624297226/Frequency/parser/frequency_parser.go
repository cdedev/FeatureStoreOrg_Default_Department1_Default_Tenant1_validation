// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624297226/Frequency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Frequency

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

type FrequencyParser struct {
	*antlr.BaseParser
}

var frequencyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func frequencyParserInit() {
	staticData := &frequencyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FREQUENCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "frequency",
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

// FrequencyParserInit initializes any static state used to implement FrequencyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFrequencyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FrequencyParserInit() {
	staticData := &frequencyParserStaticData
	staticData.once.Do(frequencyParserInit)
}

// NewFrequencyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFrequencyParser(input antlr.TokenStream) *FrequencyParser {
	FrequencyParserInit()
	this := new(FrequencyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &frequencyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Frequency.g4"

	return this
}

// FrequencyParser tokens.
const (
	FrequencyParserEOF       = antlr.TokenEOF
	FrequencyParserFREQUENCY = 1
	FrequencyParserOWNKEY    = 2
	FrequencyParserSPLITKEY  = 3
	FrequencyParserWS        = 4
)

// FrequencyParser rules.
const (
	FrequencyParserRULE_expression = 0
	FrequencyParserRULE_frequency  = 1
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
	p.RuleIndex = FrequencyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FrequencyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Frequency() IFrequencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrequencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrequencyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FrequencyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrequencyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrequencyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FrequencyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FrequencyParserRULE_expression)

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
		p.Frequency()
	}
	{
		p.SetState(5)
		p.Match(FrequencyParserEOF)
	}

	return localctx
}

// IFrequencyContext is an interface to support dynamic dispatch.
type IFrequencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrequencyContext differentiates from other interfaces.
	IsFrequencyContext()
}

type FrequencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrequencyContext() *FrequencyContext {
	var p = new(FrequencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FrequencyParserRULE_frequency
	return p
}

func (*FrequencyContext) IsFrequencyContext() {}

func NewFrequencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrequencyContext {
	var p = new(FrequencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FrequencyParserRULE_frequency

	return p
}

func (s *FrequencyContext) GetParser() antlr.Parser { return s.parser }

func (s *FrequencyContext) FREQUENCY() antlr.TerminalNode {
	return s.GetToken(FrequencyParserFREQUENCY, 0)
}

func (s *FrequencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrequencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrequencyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrequencyListener); ok {
		listenerT.EnterFrequency(s)
	}
}

func (s *FrequencyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrequencyListener); ok {
		listenerT.ExitFrequency(s)
	}
}

func (p *FrequencyParser) Frequency() (localctx IFrequencyContext) {
	this := p
	_ = this

	localctx = NewFrequencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FrequencyParserRULE_frequency)

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
		p.Match(FrequencyParserFREQUENCY)
	}

	return localctx
}
