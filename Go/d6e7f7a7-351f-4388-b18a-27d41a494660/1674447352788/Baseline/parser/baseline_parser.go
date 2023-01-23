// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447352788/Baseline.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Baseline

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

type BaselineParser struct {
	*antlr.BaseParser
}

var baselineParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func baselineParserInit() {
	staticData := &baselineParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BASELINE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "baseline",
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

// BaselineParserInit initializes any static state used to implement BaselineParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBaselineParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BaselineParserInit() {
	staticData := &baselineParserStaticData
	staticData.once.Do(baselineParserInit)
}

// NewBaselineParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBaselineParser(input antlr.TokenStream) *BaselineParser {
	BaselineParserInit()
	this := new(BaselineParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &baselineParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Baseline.g4"

	return this
}

// BaselineParser tokens.
const (
	BaselineParserEOF      = antlr.TokenEOF
	BaselineParserBASELINE = 1
	BaselineParserOWNKEY   = 2
	BaselineParserSPLITKEY = 3
	BaselineParserWS       = 4
)

// BaselineParser rules.
const (
	BaselineParserRULE_expression = 0
	BaselineParserRULE_baseline   = 1
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
	p.RuleIndex = BaselineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BaselineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Baseline() IBaselineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaselineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaselineContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BaselineParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BaselineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BaselineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BaselineParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BaselineParserRULE_expression)

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
		p.Baseline()
	}
	{
		p.SetState(5)
		p.Match(BaselineParserEOF)
	}

	return localctx
}

// IBaselineContext is an interface to support dynamic dispatch.
type IBaselineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaselineContext differentiates from other interfaces.
	IsBaselineContext()
}

type BaselineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaselineContext() *BaselineContext {
	var p = new(BaselineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BaselineParserRULE_baseline
	return p
}

func (*BaselineContext) IsBaselineContext() {}

func NewBaselineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaselineContext {
	var p = new(BaselineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BaselineParserRULE_baseline

	return p
}

func (s *BaselineContext) GetParser() antlr.Parser { return s.parser }

func (s *BaselineContext) BASELINE() antlr.TerminalNode {
	return s.GetToken(BaselineParserBASELINE, 0)
}

func (s *BaselineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaselineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaselineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BaselineListener); ok {
		listenerT.EnterBaseline(s)
	}
}

func (s *BaselineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BaselineListener); ok {
		listenerT.ExitBaseline(s)
	}
}

func (p *BaselineParser) Baseline() (localctx IBaselineContext) {
	this := p
	_ = this

	localctx = NewBaselineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BaselineParserRULE_baseline)

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
		p.Match(BaselineParserBASELINE)
	}

	return localctx
}
