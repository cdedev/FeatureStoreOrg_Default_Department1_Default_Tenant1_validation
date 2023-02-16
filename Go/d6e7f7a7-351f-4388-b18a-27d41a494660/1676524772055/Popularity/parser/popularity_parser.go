// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524772055/Popularity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Popularity

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

type PopularityParser struct {
	*antlr.BaseParser
}

var popularityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func popularityParserInit() {
	staticData := &popularityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "POPULARITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "popularity",
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

// PopularityParserInit initializes any static state used to implement PopularityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPopularityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PopularityParserInit() {
	staticData := &popularityParserStaticData
	staticData.once.Do(popularityParserInit)
}

// NewPopularityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPopularityParser(input antlr.TokenStream) *PopularityParser {
	PopularityParserInit()
	this := new(PopularityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &popularityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Popularity.g4"

	return this
}

// PopularityParser tokens.
const (
	PopularityParserEOF        = antlr.TokenEOF
	PopularityParserPOPULARITY = 1
	PopularityParserOWNKEY     = 2
	PopularityParserSPLITKEY   = 3
	PopularityParserWS         = 4
)

// PopularityParser rules.
const (
	PopularityParserRULE_expression = 0
	PopularityParserRULE_popularity = 1
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
	p.RuleIndex = PopularityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PopularityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Popularity() IPopularityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPopularityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPopularityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PopularityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopularityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopularityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PopularityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PopularityParserRULE_expression)

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
		p.Popularity()
	}
	{
		p.SetState(5)
		p.Match(PopularityParserEOF)
	}

	return localctx
}

// IPopularityContext is an interface to support dynamic dispatch.
type IPopularityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPopularityContext differentiates from other interfaces.
	IsPopularityContext()
}

type PopularityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPopularityContext() *PopularityContext {
	var p = new(PopularityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PopularityParserRULE_popularity
	return p
}

func (*PopularityContext) IsPopularityContext() {}

func NewPopularityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PopularityContext {
	var p = new(PopularityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PopularityParserRULE_popularity

	return p
}

func (s *PopularityContext) GetParser() antlr.Parser { return s.parser }

func (s *PopularityContext) POPULARITY() antlr.TerminalNode {
	return s.GetToken(PopularityParserPOPULARITY, 0)
}

func (s *PopularityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PopularityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PopularityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopularityListener); ok {
		listenerT.EnterPopularity(s)
	}
}

func (s *PopularityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PopularityListener); ok {
		listenerT.ExitPopularity(s)
	}
}

func (p *PopularityParser) Popularity() (localctx IPopularityContext) {
	this := p
	_ = this

	localctx = NewPopularityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PopularityParserRULE_popularity)

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
		p.Match(PopularityParserPOPULARITY)
	}

	return localctx
}
