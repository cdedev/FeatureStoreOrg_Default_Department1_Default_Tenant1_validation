// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376971814/Levy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Levy

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

type LevyParser struct {
	*antlr.BaseParser
}

var levyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func levyParserInit() {
	staticData := &levyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LEVY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "levy",
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

// LevyParserInit initializes any static state used to implement LevyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLevyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LevyParserInit() {
	staticData := &levyParserStaticData
	staticData.once.Do(levyParserInit)
}

// NewLevyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLevyParser(input antlr.TokenStream) *LevyParser {
	LevyParserInit()
	this := new(LevyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &levyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Levy.g4"

	return this
}

// LevyParser tokens.
const (
	LevyParserEOF      = antlr.TokenEOF
	LevyParserLEVY     = 1
	LevyParserOWNKEY   = 2
	LevyParserSPLITKEY = 3
	LevyParserWS       = 4
)

// LevyParser rules.
const (
	LevyParserRULE_expression = 0
	LevyParserRULE_levy       = 1
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
	p.RuleIndex = LevyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LevyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Levy() ILevyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILevyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILevyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LevyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LevyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LevyParserRULE_expression)

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
		p.Levy()
	}
	{
		p.SetState(5)
		p.Match(LevyParserEOF)
	}

	return localctx
}

// ILevyContext is an interface to support dynamic dispatch.
type ILevyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLevyContext differentiates from other interfaces.
	IsLevyContext()
}

type LevyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLevyContext() *LevyContext {
	var p = new(LevyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LevyParserRULE_levy
	return p
}

func (*LevyContext) IsLevyContext() {}

func NewLevyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LevyContext {
	var p = new(LevyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LevyParserRULE_levy

	return p
}

func (s *LevyContext) GetParser() antlr.Parser { return s.parser }

func (s *LevyContext) LEVY() antlr.TerminalNode {
	return s.GetToken(LevyParserLEVY, 0)
}

func (s *LevyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LevyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LevyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevyListener); ok {
		listenerT.EnterLevy(s)
	}
}

func (s *LevyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevyListener); ok {
		listenerT.ExitLevy(s)
	}
}

func (p *LevyParser) Levy() (localctx ILevyContext) {
	this := p
	_ = this

	localctx = NewLevyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LevyParserRULE_levy)

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
		p.Match(LevyParserLEVY)
	}

	return localctx
}
