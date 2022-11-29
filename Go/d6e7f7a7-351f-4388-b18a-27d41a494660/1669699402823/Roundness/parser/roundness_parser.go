// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699402823/Roundness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Roundness

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

type RoundnessParser struct {
	*antlr.BaseParser
}

var roundnessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func roundnessParserInit() {
	staticData := &roundnessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ROUNDNESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "roundness",
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

// RoundnessParserInit initializes any static state used to implement RoundnessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRoundnessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RoundnessParserInit() {
	staticData := &roundnessParserStaticData
	staticData.once.Do(roundnessParserInit)
}

// NewRoundnessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRoundnessParser(input antlr.TokenStream) *RoundnessParser {
	RoundnessParserInit()
	this := new(RoundnessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &roundnessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Roundness.g4"

	return this
}

// RoundnessParser tokens.
const (
	RoundnessParserEOF       = antlr.TokenEOF
	RoundnessParserROUNDNESS = 1
	RoundnessParserOWNKEY    = 2
	RoundnessParserSPLITKEY  = 3
	RoundnessParserWS        = 4
)

// RoundnessParser rules.
const (
	RoundnessParserRULE_expression = 0
	RoundnessParserRULE_roundness  = 1
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
	p.RuleIndex = RoundnessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoundnessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Roundness() IRoundnessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoundnessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoundnessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RoundnessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoundnessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoundnessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RoundnessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RoundnessParserRULE_expression)

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
		p.Roundness()
	}
	{
		p.SetState(5)
		p.Match(RoundnessParserEOF)
	}

	return localctx
}

// IRoundnessContext is an interface to support dynamic dispatch.
type IRoundnessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoundnessContext differentiates from other interfaces.
	IsRoundnessContext()
}

type RoundnessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoundnessContext() *RoundnessContext {
	var p = new(RoundnessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoundnessParserRULE_roundness
	return p
}

func (*RoundnessContext) IsRoundnessContext() {}

func NewRoundnessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoundnessContext {
	var p = new(RoundnessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoundnessParserRULE_roundness

	return p
}

func (s *RoundnessContext) GetParser() antlr.Parser { return s.parser }

func (s *RoundnessContext) ROUNDNESS() antlr.TerminalNode {
	return s.GetToken(RoundnessParserROUNDNESS, 0)
}

func (s *RoundnessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundnessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoundnessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoundnessListener); ok {
		listenerT.EnterRoundness(s)
	}
}

func (s *RoundnessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoundnessListener); ok {
		listenerT.ExitRoundness(s)
	}
}

func (p *RoundnessParser) Roundness() (localctx IRoundnessContext) {
	this := p
	_ = this

	localctx = NewRoundnessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RoundnessParserRULE_roundness)

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
		p.Match(RoundnessParserROUNDNESS)
	}

	return localctx
}
