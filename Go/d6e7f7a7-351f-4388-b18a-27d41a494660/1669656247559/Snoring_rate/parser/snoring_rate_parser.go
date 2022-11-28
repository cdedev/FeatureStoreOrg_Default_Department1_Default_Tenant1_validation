// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656247559/Snoring_rate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Snoring_rate

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

type Snoring_rateParser struct {
	*antlr.BaseParser
}

var snoring_rateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func snoring_rateParserInit() {
	staticData := &snoring_rateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SNORING_RATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "snoring_rate",
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

// Snoring_rateParserInit initializes any static state used to implement Snoring_rateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSnoring_rateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Snoring_rateParserInit() {
	staticData := &snoring_rateParserStaticData
	staticData.once.Do(snoring_rateParserInit)
}

// NewSnoring_rateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSnoring_rateParser(input antlr.TokenStream) *Snoring_rateParser {
	Snoring_rateParserInit()
	this := new(Snoring_rateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &snoring_rateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Snoring_rate.g4"

	return this
}

// Snoring_rateParser tokens.
const (
	Snoring_rateParserEOF          = antlr.TokenEOF
	Snoring_rateParserSNORING_RATE = 1
	Snoring_rateParserOWNKEY       = 2
	Snoring_rateParserSPLITKEY     = 3
	Snoring_rateParserWS           = 4
)

// Snoring_rateParser rules.
const (
	Snoring_rateParserRULE_expression   = 0
	Snoring_rateParserRULE_snoring_rate = 1
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
	p.RuleIndex = Snoring_rateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Snoring_rateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Snoring_rate() ISnoring_rateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISnoring_rateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISnoring_rateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Snoring_rateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Snoring_rateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Snoring_rateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Snoring_rateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Snoring_rateParserRULE_expression)

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
		p.Snoring_rate()
	}
	{
		p.SetState(5)
		p.Match(Snoring_rateParserEOF)
	}

	return localctx
}

// ISnoring_rateContext is an interface to support dynamic dispatch.
type ISnoring_rateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSnoring_rateContext differentiates from other interfaces.
	IsSnoring_rateContext()
}

type Snoring_rateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySnoring_rateContext() *Snoring_rateContext {
	var p = new(Snoring_rateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Snoring_rateParserRULE_snoring_rate
	return p
}

func (*Snoring_rateContext) IsSnoring_rateContext() {}

func NewSnoring_rateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Snoring_rateContext {
	var p = new(Snoring_rateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Snoring_rateParserRULE_snoring_rate

	return p
}

func (s *Snoring_rateContext) GetParser() antlr.Parser { return s.parser }

func (s *Snoring_rateContext) SNORING_RATE() antlr.TerminalNode {
	return s.GetToken(Snoring_rateParserSNORING_RATE, 0)
}

func (s *Snoring_rateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Snoring_rateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Snoring_rateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Snoring_rateListener); ok {
		listenerT.EnterSnoring_rate(s)
	}
}

func (s *Snoring_rateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Snoring_rateListener); ok {
		listenerT.ExitSnoring_rate(s)
	}
}

func (p *Snoring_rateParser) Snoring_rate() (localctx ISnoring_rateContext) {
	this := p
	_ = this

	localctx = NewSnoring_rateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Snoring_rateParserRULE_snoring_rate)

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
		p.Match(Snoring_rateParserSNORING_RATE)
	}

	return localctx
}
