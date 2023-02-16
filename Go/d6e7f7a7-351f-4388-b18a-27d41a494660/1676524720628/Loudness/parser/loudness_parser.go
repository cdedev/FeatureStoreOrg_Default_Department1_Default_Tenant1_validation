// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524720628/Loudness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loudness

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

type LoudnessParser struct {
	*antlr.BaseParser
}

var loudnessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func loudnessParserInit() {
	staticData := &loudnessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LOUDNESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "loudness",
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

// LoudnessParserInit initializes any static state used to implement LoudnessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLoudnessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LoudnessParserInit() {
	staticData := &loudnessParserStaticData
	staticData.once.Do(loudnessParserInit)
}

// NewLoudnessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLoudnessParser(input antlr.TokenStream) *LoudnessParser {
	LoudnessParserInit()
	this := new(LoudnessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &loudnessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Loudness.g4"

	return this
}

// LoudnessParser tokens.
const (
	LoudnessParserEOF      = antlr.TokenEOF
	LoudnessParserLOUDNESS = 1
	LoudnessParserOWNKEY   = 2
	LoudnessParserSPLITKEY = 3
	LoudnessParserWS       = 4
)

// LoudnessParser rules.
const (
	LoudnessParserRULE_expression = 0
	LoudnessParserRULE_loudness   = 1
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
	p.RuleIndex = LoudnessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoudnessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Loudness() ILoudnessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoudnessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoudnessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LoudnessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudnessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudnessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LoudnessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LoudnessParserRULE_expression)

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
		p.Loudness()
	}
	{
		p.SetState(5)
		p.Match(LoudnessParserEOF)
	}

	return localctx
}

// ILoudnessContext is an interface to support dynamic dispatch.
type ILoudnessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoudnessContext differentiates from other interfaces.
	IsLoudnessContext()
}

type LoudnessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoudnessContext() *LoudnessContext {
	var p = new(LoudnessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LoudnessParserRULE_loudness
	return p
}

func (*LoudnessContext) IsLoudnessContext() {}

func NewLoudnessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoudnessContext {
	var p = new(LoudnessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LoudnessParserRULE_loudness

	return p
}

func (s *LoudnessContext) GetParser() antlr.Parser { return s.parser }

func (s *LoudnessContext) LOUDNESS() antlr.TerminalNode {
	return s.GetToken(LoudnessParserLOUDNESS, 0)
}

func (s *LoudnessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoudnessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoudnessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudnessListener); ok {
		listenerT.EnterLoudness(s)
	}
}

func (s *LoudnessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LoudnessListener); ok {
		listenerT.ExitLoudness(s)
	}
}

func (p *LoudnessParser) Loudness() (localctx ILoudnessContext) {
	this := p
	_ = this

	localctx = NewLoudnessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LoudnessParserRULE_loudness)

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
		p.Match(LoudnessParserLOUDNESS)
	}

	return localctx
}
