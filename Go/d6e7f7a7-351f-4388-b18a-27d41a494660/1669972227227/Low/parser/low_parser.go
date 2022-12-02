// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972227227/Low.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Low

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

type LowParser struct {
	*antlr.BaseParser
}

var lowParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lowParserInit() {
	staticData := &lowParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LOW", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "low",
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

// LowParserInit initializes any static state used to implement LowParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLowParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LowParserInit() {
	staticData := &lowParserStaticData
	staticData.once.Do(lowParserInit)
}

// NewLowParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLowParser(input antlr.TokenStream) *LowParser {
	LowParserInit()
	this := new(LowParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lowParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Low.g4"

	return this
}

// LowParser tokens.
const (
	LowParserEOF      = antlr.TokenEOF
	LowParserLOW      = 1
	LowParserOWNKEY   = 2
	LowParserSPLITKEY = 3
	LowParserWS       = 4
)

// LowParser rules.
const (
	LowParserRULE_expression = 0
	LowParserRULE_low        = 1
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
	p.RuleIndex = LowParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LowParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Low() ILowContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILowContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILowContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LowParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LowListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LowListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LowParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LowParserRULE_expression)

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
		p.Low()
	}
	{
		p.SetState(5)
		p.Match(LowParserEOF)
	}

	return localctx
}

// ILowContext is an interface to support dynamic dispatch.
type ILowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLowContext differentiates from other interfaces.
	IsLowContext()
}

type LowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLowContext() *LowContext {
	var p = new(LowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LowParserRULE_low
	return p
}

func (*LowContext) IsLowContext() {}

func NewLowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LowContext {
	var p = new(LowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LowParserRULE_low

	return p
}

func (s *LowContext) GetParser() antlr.Parser { return s.parser }

func (s *LowContext) LOW() antlr.TerminalNode {
	return s.GetToken(LowParserLOW, 0)
}

func (s *LowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LowListener); ok {
		listenerT.EnterLow(s)
	}
}

func (s *LowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LowListener); ok {
		listenerT.ExitLow(s)
	}
}

func (p *LowParser) Low() (localctx ILowContext) {
	this := p
	_ = this

	localctx = NewLowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LowParserRULE_low)

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
		p.Match(LowParserLOW)
	}

	return localctx
}
