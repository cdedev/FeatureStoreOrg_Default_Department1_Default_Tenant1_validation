// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878210754/Hypoacusis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hypoacusis

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

type HypoacusisParser struct {
	*antlr.BaseParser
}

var hypoacusisParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hypoacusisParserInit() {
	staticData := &hypoacusisParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HYPOACUSIS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hypoacusis",
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

// HypoacusisParserInit initializes any static state used to implement HypoacusisParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHypoacusisParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HypoacusisParserInit() {
	staticData := &hypoacusisParserStaticData
	staticData.once.Do(hypoacusisParserInit)
}

// NewHypoacusisParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHypoacusisParser(input antlr.TokenStream) *HypoacusisParser {
	HypoacusisParserInit()
	this := new(HypoacusisParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hypoacusisParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Hypoacusis.g4"

	return this
}

// HypoacusisParser tokens.
const (
	HypoacusisParserEOF        = antlr.TokenEOF
	HypoacusisParserHYPOACUSIS = 1
	HypoacusisParserOWNKEY     = 2
	HypoacusisParserSPLITKEY   = 3
	HypoacusisParserWS         = 4
)

// HypoacusisParser rules.
const (
	HypoacusisParserRULE_expression = 0
	HypoacusisParserRULE_hypoacusis = 1
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
	p.RuleIndex = HypoacusisParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HypoacusisParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hypoacusis() IHypoacusisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHypoacusisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHypoacusisContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HypoacusisParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypoacusisListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypoacusisListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HypoacusisParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HypoacusisParserRULE_expression)

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
		p.Hypoacusis()
	}
	{
		p.SetState(5)
		p.Match(HypoacusisParserEOF)
	}

	return localctx
}

// IHypoacusisContext is an interface to support dynamic dispatch.
type IHypoacusisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHypoacusisContext differentiates from other interfaces.
	IsHypoacusisContext()
}

type HypoacusisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHypoacusisContext() *HypoacusisContext {
	var p = new(HypoacusisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HypoacusisParserRULE_hypoacusis
	return p
}

func (*HypoacusisContext) IsHypoacusisContext() {}

func NewHypoacusisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HypoacusisContext {
	var p = new(HypoacusisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HypoacusisParserRULE_hypoacusis

	return p
}

func (s *HypoacusisContext) GetParser() antlr.Parser { return s.parser }

func (s *HypoacusisContext) HYPOACUSIS() antlr.TerminalNode {
	return s.GetToken(HypoacusisParserHYPOACUSIS, 0)
}

func (s *HypoacusisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HypoacusisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HypoacusisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypoacusisListener); ok {
		listenerT.EnterHypoacusis(s)
	}
}

func (s *HypoacusisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HypoacusisListener); ok {
		listenerT.ExitHypoacusis(s)
	}
}

func (p *HypoacusisParser) Hypoacusis() (localctx IHypoacusisContext) {
	this := p
	_ = this

	localctx = NewHypoacusisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HypoacusisParserRULE_hypoacusis)

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
		p.Match(HypoacusisParserHYPOACUSIS)
	}

	return localctx
}
