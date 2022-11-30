// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789612226/Probability.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Probability

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

type ProbabilityParser struct {
	*antlr.BaseParser
}

var probabilityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func probabilityParserInit() {
	staticData := &probabilityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PROBABILITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "probability",
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

// ProbabilityParserInit initializes any static state used to implement ProbabilityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProbabilityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProbabilityParserInit() {
	staticData := &probabilityParserStaticData
	staticData.once.Do(probabilityParserInit)
}

// NewProbabilityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProbabilityParser(input antlr.TokenStream) *ProbabilityParser {
	ProbabilityParserInit()
	this := new(ProbabilityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &probabilityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Probability.g4"

	return this
}

// ProbabilityParser tokens.
const (
	ProbabilityParserEOF         = antlr.TokenEOF
	ProbabilityParserPROBABILITY = 1
	ProbabilityParserOWNKEY      = 2
	ProbabilityParserSPLITKEY    = 3
	ProbabilityParserWS          = 4
)

// ProbabilityParser rules.
const (
	ProbabilityParserRULE_expression  = 0
	ProbabilityParserRULE_probability = 1
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
	p.RuleIndex = ProbabilityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProbabilityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Probability() IProbabilityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProbabilityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProbabilityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProbabilityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProbabilityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProbabilityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ProbabilityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProbabilityParserRULE_expression)

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
		p.Probability()
	}
	{
		p.SetState(5)
		p.Match(ProbabilityParserEOF)
	}

	return localctx
}

// IProbabilityContext is an interface to support dynamic dispatch.
type IProbabilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProbabilityContext differentiates from other interfaces.
	IsProbabilityContext()
}

type ProbabilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProbabilityContext() *ProbabilityContext {
	var p = new(ProbabilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProbabilityParserRULE_probability
	return p
}

func (*ProbabilityContext) IsProbabilityContext() {}

func NewProbabilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProbabilityContext {
	var p = new(ProbabilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProbabilityParserRULE_probability

	return p
}

func (s *ProbabilityContext) GetParser() antlr.Parser { return s.parser }

func (s *ProbabilityContext) PROBABILITY() antlr.TerminalNode {
	return s.GetToken(ProbabilityParserPROBABILITY, 0)
}

func (s *ProbabilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProbabilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProbabilityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProbabilityListener); ok {
		listenerT.EnterProbability(s)
	}
}

func (s *ProbabilityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProbabilityListener); ok {
		listenerT.ExitProbability(s)
	}
}

func (p *ProbabilityParser) Probability() (localctx IProbabilityContext) {
	this := p
	_ = this

	localctx = NewProbabilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProbabilityParserRULE_probability)

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
		p.Match(ProbabilityParserPROBABILITY)
	}

	return localctx
}
