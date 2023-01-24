// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538930547/Trails.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Trails

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

type TrailsParser struct {
	*antlr.BaseParser
}

var trailsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func trailsParserInit() {
	staticData := &trailsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TRAILS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "trails",
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

// TrailsParserInit initializes any static state used to implement TrailsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTrailsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TrailsParserInit() {
	staticData := &trailsParserStaticData
	staticData.once.Do(trailsParserInit)
}

// NewTrailsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTrailsParser(input antlr.TokenStream) *TrailsParser {
	TrailsParserInit()
	this := new(TrailsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &trailsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Trails.g4"

	return this
}

// TrailsParser tokens.
const (
	TrailsParserEOF      = antlr.TokenEOF
	TrailsParserTRAILS   = 1
	TrailsParserOWNKEY   = 2
	TrailsParserSPLITKEY = 3
	TrailsParserWS       = 4
)

// TrailsParser rules.
const (
	TrailsParserRULE_expression = 0
	TrailsParserRULE_trails     = 1
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
	p.RuleIndex = TrailsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TrailsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Trails() ITrailsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrailsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrailsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TrailsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrailsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrailsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TrailsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TrailsParserRULE_expression)

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
		p.Trails()
	}
	{
		p.SetState(5)
		p.Match(TrailsParserEOF)
	}

	return localctx
}

// ITrailsContext is an interface to support dynamic dispatch.
type ITrailsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrailsContext differentiates from other interfaces.
	IsTrailsContext()
}

type TrailsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrailsContext() *TrailsContext {
	var p = new(TrailsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TrailsParserRULE_trails
	return p
}

func (*TrailsContext) IsTrailsContext() {}

func NewTrailsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrailsContext {
	var p = new(TrailsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TrailsParserRULE_trails

	return p
}

func (s *TrailsContext) GetParser() antlr.Parser { return s.parser }

func (s *TrailsContext) TRAILS() antlr.TerminalNode {
	return s.GetToken(TrailsParserTRAILS, 0)
}

func (s *TrailsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrailsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrailsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrailsListener); ok {
		listenerT.EnterTrails(s)
	}
}

func (s *TrailsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrailsListener); ok {
		listenerT.ExitTrails(s)
	}
}

func (p *TrailsParser) Trails() (localctx ITrailsContext) {
	this := p
	_ = this

	localctx = NewTrailsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TrailsParserRULE_trails)

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
		p.Match(TrailsParserTRAILS)
	}

	return localctx
}
