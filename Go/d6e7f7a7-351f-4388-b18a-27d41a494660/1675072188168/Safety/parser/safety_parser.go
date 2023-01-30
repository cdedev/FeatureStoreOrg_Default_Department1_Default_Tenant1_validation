// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072188168/Safety.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Safety

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

type SafetyParser struct {
	*antlr.BaseParser
}

var safetyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func safetyParserInit() {
	staticData := &safetyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SAFETY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "safety",
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

// SafetyParserInit initializes any static state used to implement SafetyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSafetyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SafetyParserInit() {
	staticData := &safetyParserStaticData
	staticData.once.Do(safetyParserInit)
}

// NewSafetyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSafetyParser(input antlr.TokenStream) *SafetyParser {
	SafetyParserInit()
	this := new(SafetyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &safetyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Safety.g4"

	return this
}

// SafetyParser tokens.
const (
	SafetyParserEOF      = antlr.TokenEOF
	SafetyParserSAFETY   = 1
	SafetyParserOWNKEY   = 2
	SafetyParserSPLITKEY = 3
	SafetyParserWS       = 4
)

// SafetyParser rules.
const (
	SafetyParserRULE_expression = 0
	SafetyParserRULE_safety     = 1
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
	p.RuleIndex = SafetyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SafetyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Safety() ISafetyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISafetyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISafetyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SafetyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SafetyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SafetyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SafetyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SafetyParserRULE_expression)

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
		p.Safety()
	}
	{
		p.SetState(5)
		p.Match(SafetyParserEOF)
	}

	return localctx
}

// ISafetyContext is an interface to support dynamic dispatch.
type ISafetyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSafetyContext differentiates from other interfaces.
	IsSafetyContext()
}

type SafetyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySafetyContext() *SafetyContext {
	var p = new(SafetyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SafetyParserRULE_safety
	return p
}

func (*SafetyContext) IsSafetyContext() {}

func NewSafetyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SafetyContext {
	var p = new(SafetyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SafetyParserRULE_safety

	return p
}

func (s *SafetyContext) GetParser() antlr.Parser { return s.parser }

func (s *SafetyContext) SAFETY() antlr.TerminalNode {
	return s.GetToken(SafetyParserSAFETY, 0)
}

func (s *SafetyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SafetyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SafetyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SafetyListener); ok {
		listenerT.EnterSafety(s)
	}
}

func (s *SafetyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SafetyListener); ok {
		listenerT.ExitSafety(s)
	}
}

func (p *SafetyParser) Safety() (localctx ISafetyContext) {
	this := p
	_ = this

	localctx = NewSafetyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SafetyParserRULE_safety)

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
		p.Match(SafetyParserSAFETY)
	}

	return localctx
}
