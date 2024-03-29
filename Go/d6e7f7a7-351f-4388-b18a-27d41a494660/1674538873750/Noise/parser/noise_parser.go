// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538873750/Noise.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Noise

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

type NoiseParser struct {
	*antlr.BaseParser
}

var noiseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func noiseParserInit() {
	staticData := &noiseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NOISE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "noise",
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

// NoiseParserInit initializes any static state used to implement NoiseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNoiseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NoiseParserInit() {
	staticData := &noiseParserStaticData
	staticData.once.Do(noiseParserInit)
}

// NewNoiseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNoiseParser(input antlr.TokenStream) *NoiseParser {
	NoiseParserInit()
	this := new(NoiseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &noiseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Noise.g4"

	return this
}

// NoiseParser tokens.
const (
	NoiseParserEOF      = antlr.TokenEOF
	NoiseParserNOISE    = 1
	NoiseParserOWNKEY   = 2
	NoiseParserSPLITKEY = 3
	NoiseParserWS       = 4
)

// NoiseParser rules.
const (
	NoiseParserRULE_expression = 0
	NoiseParserRULE_noise      = 1
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
	p.RuleIndex = NoiseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NoiseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Noise() INoiseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoiseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoiseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NoiseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoiseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoiseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NoiseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NoiseParserRULE_expression)

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
		p.Noise()
	}
	{
		p.SetState(5)
		p.Match(NoiseParserEOF)
	}

	return localctx
}

// INoiseContext is an interface to support dynamic dispatch.
type INoiseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoiseContext differentiates from other interfaces.
	IsNoiseContext()
}

type NoiseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoiseContext() *NoiseContext {
	var p = new(NoiseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NoiseParserRULE_noise
	return p
}

func (*NoiseContext) IsNoiseContext() {}

func NewNoiseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoiseContext {
	var p = new(NoiseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NoiseParserRULE_noise

	return p
}

func (s *NoiseContext) GetParser() antlr.Parser { return s.parser }

func (s *NoiseContext) NOISE() antlr.TerminalNode {
	return s.GetToken(NoiseParserNOISE, 0)
}

func (s *NoiseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoiseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoiseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoiseListener); ok {
		listenerT.EnterNoise(s)
	}
}

func (s *NoiseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoiseListener); ok {
		listenerT.ExitNoise(s)
	}
}

func (p *NoiseParser) Noise() (localctx INoiseContext) {
	this := p
	_ = this

	localctx = NewNoiseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NoiseParserRULE_noise)

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
		p.Match(NoiseParserNOISE)
	}

	return localctx
}
