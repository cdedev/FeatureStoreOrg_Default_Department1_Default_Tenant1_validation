// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675076910173/Sunset.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sunset

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

type SunsetParser struct {
	*antlr.BaseParser
}

var sunsetParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sunsetParserInit() {
	staticData := &sunsetParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SUNSET", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sunset",
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

// SunsetParserInit initializes any static state used to implement SunsetParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSunsetParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SunsetParserInit() {
	staticData := &sunsetParserStaticData
	staticData.once.Do(sunsetParserInit)
}

// NewSunsetParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSunsetParser(input antlr.TokenStream) *SunsetParser {
	SunsetParserInit()
	this := new(SunsetParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sunsetParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sunset.g4"

	return this
}

// SunsetParser tokens.
const (
	SunsetParserEOF      = antlr.TokenEOF
	SunsetParserSUNSET   = 1
	SunsetParserOWNKEY   = 2
	SunsetParserSPLITKEY = 3
	SunsetParserWS       = 4
)

// SunsetParser rules.
const (
	SunsetParserRULE_expression = 0
	SunsetParserRULE_sunset     = 1
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
	p.RuleIndex = SunsetParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SunsetParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sunset() ISunsetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISunsetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISunsetContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SunsetParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunsetListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunsetListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SunsetParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SunsetParserRULE_expression)

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
		p.Sunset()
	}
	{
		p.SetState(5)
		p.Match(SunsetParserEOF)
	}

	return localctx
}

// ISunsetContext is an interface to support dynamic dispatch.
type ISunsetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSunsetContext differentiates from other interfaces.
	IsSunsetContext()
}

type SunsetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySunsetContext() *SunsetContext {
	var p = new(SunsetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SunsetParserRULE_sunset
	return p
}

func (*SunsetContext) IsSunsetContext() {}

func NewSunsetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SunsetContext {
	var p = new(SunsetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SunsetParserRULE_sunset

	return p
}

func (s *SunsetContext) GetParser() antlr.Parser { return s.parser }

func (s *SunsetContext) SUNSET() antlr.TerminalNode {
	return s.GetToken(SunsetParserSUNSET, 0)
}

func (s *SunsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SunsetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SunsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunsetListener); ok {
		listenerT.EnterSunset(s)
	}
}

func (s *SunsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunsetListener); ok {
		listenerT.ExitSunset(s)
	}
}

func (p *SunsetParser) Sunset() (localctx ISunsetContext) {
	this := p
	_ = this

	localctx = NewSunsetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SunsetParserRULE_sunset)

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
		p.Match(SunsetParserSUNSET)
	}

	return localctx
}
