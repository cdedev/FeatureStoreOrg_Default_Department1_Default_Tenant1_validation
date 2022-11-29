// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700956001/Steps.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Steps

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

type StepsParser struct {
	*antlr.BaseParser
}

var stepsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stepsParserInit() {
	staticData := &stepsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STEPS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "steps",
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

// StepsParserInit initializes any static state used to implement StepsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStepsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StepsParserInit() {
	staticData := &stepsParserStaticData
	staticData.once.Do(stepsParserInit)
}

// NewStepsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStepsParser(input antlr.TokenStream) *StepsParser {
	StepsParserInit()
	this := new(StepsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &stepsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Steps.g4"

	return this
}

// StepsParser tokens.
const (
	StepsParserEOF      = antlr.TokenEOF
	StepsParserSTEPS    = 1
	StepsParserOWNKEY   = 2
	StepsParserSPLITKEY = 3
	StepsParserWS       = 4
)

// StepsParser rules.
const (
	StepsParserRULE_expression = 0
	StepsParserRULE_steps      = 1
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
	p.RuleIndex = StepsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StepsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Steps() IStepsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStepsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStepsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StepsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StepsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StepsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StepsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StepsParserRULE_expression)

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
		p.Steps()
	}
	{
		p.SetState(5)
		p.Match(StepsParserEOF)
	}

	return localctx
}

// IStepsContext is an interface to support dynamic dispatch.
type IStepsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStepsContext differentiates from other interfaces.
	IsStepsContext()
}

type StepsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStepsContext() *StepsContext {
	var p = new(StepsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StepsParserRULE_steps
	return p
}

func (*StepsContext) IsStepsContext() {}

func NewStepsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StepsContext {
	var p = new(StepsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StepsParserRULE_steps

	return p
}

func (s *StepsContext) GetParser() antlr.Parser { return s.parser }

func (s *StepsContext) STEPS() antlr.TerminalNode {
	return s.GetToken(StepsParserSTEPS, 0)
}

func (s *StepsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StepsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StepsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StepsListener); ok {
		listenerT.EnterSteps(s)
	}
}

func (s *StepsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StepsListener); ok {
		listenerT.ExitSteps(s)
	}
}

func (p *StepsParser) Steps() (localctx IStepsContext) {
	this := p
	_ = this

	localctx = NewStepsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StepsParserRULE_steps)

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
		p.Match(StepsParserSTEPS)
	}

	return localctx
}
