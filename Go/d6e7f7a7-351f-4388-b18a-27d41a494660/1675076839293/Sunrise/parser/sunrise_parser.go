// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675076839293/Sunrise.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sunrise

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

type SunriseParser struct {
	*antlr.BaseParser
}

var sunriseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sunriseParserInit() {
	staticData := &sunriseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SUNRISE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sunrise",
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

// SunriseParserInit initializes any static state used to implement SunriseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSunriseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SunriseParserInit() {
	staticData := &sunriseParserStaticData
	staticData.once.Do(sunriseParserInit)
}

// NewSunriseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSunriseParser(input antlr.TokenStream) *SunriseParser {
	SunriseParserInit()
	this := new(SunriseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sunriseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sunrise.g4"

	return this
}

// SunriseParser tokens.
const (
	SunriseParserEOF      = antlr.TokenEOF
	SunriseParserSUNRISE  = 1
	SunriseParserOWNKEY   = 2
	SunriseParserSPLITKEY = 3
	SunriseParserWS       = 4
)

// SunriseParser rules.
const (
	SunriseParserRULE_expression = 0
	SunriseParserRULE_sunrise    = 1
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
	p.RuleIndex = SunriseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SunriseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sunrise() ISunriseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISunriseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISunriseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SunriseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunriseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunriseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SunriseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SunriseParserRULE_expression)

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
		p.Sunrise()
	}
	{
		p.SetState(5)
		p.Match(SunriseParserEOF)
	}

	return localctx
}

// ISunriseContext is an interface to support dynamic dispatch.
type ISunriseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSunriseContext differentiates from other interfaces.
	IsSunriseContext()
}

type SunriseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySunriseContext() *SunriseContext {
	var p = new(SunriseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SunriseParserRULE_sunrise
	return p
}

func (*SunriseContext) IsSunriseContext() {}

func NewSunriseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SunriseContext {
	var p = new(SunriseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SunriseParserRULE_sunrise

	return p
}

func (s *SunriseContext) GetParser() antlr.Parser { return s.parser }

func (s *SunriseContext) SUNRISE() antlr.TerminalNode {
	return s.GetToken(SunriseParserSUNRISE, 0)
}

func (s *SunriseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SunriseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SunriseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunriseListener); ok {
		listenerT.EnterSunrise(s)
	}
}

func (s *SunriseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SunriseListener); ok {
		listenerT.ExitSunrise(s)
	}
}

func (p *SunriseParser) Sunrise() (localctx ISunriseContext) {
	this := p
	_ = this

	localctx = NewSunriseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SunriseParserRULE_sunrise)

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
		p.Match(SunriseParserSUNRISE)
	}

	return localctx
}
