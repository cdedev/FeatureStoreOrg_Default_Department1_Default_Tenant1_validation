// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673243876019/Radiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Radiation

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

type RadiationParser struct {
	*antlr.BaseParser
}

var radiationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func radiationParserInit() {
	staticData := &radiationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RADIATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "radiation",
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

// RadiationParserInit initializes any static state used to implement RadiationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRadiationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RadiationParserInit() {
	staticData := &radiationParserStaticData
	staticData.once.Do(radiationParserInit)
}

// NewRadiationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRadiationParser(input antlr.TokenStream) *RadiationParser {
	RadiationParserInit()
	this := new(RadiationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &radiationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Radiation.g4"

	return this
}

// RadiationParser tokens.
const (
	RadiationParserEOF       = antlr.TokenEOF
	RadiationParserRADIATION = 1
	RadiationParserOWNKEY    = 2
	RadiationParserSPLITKEY  = 3
	RadiationParserWS        = 4
)

// RadiationParser rules.
const (
	RadiationParserRULE_expression = 0
	RadiationParserRULE_radiation  = 1
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
	p.RuleIndex = RadiationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RadiationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Radiation() IRadiationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRadiationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRadiationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RadiationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RadiationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RadiationParserRULE_expression)

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
		p.Radiation()
	}
	{
		p.SetState(5)
		p.Match(RadiationParserEOF)
	}

	return localctx
}

// IRadiationContext is an interface to support dynamic dispatch.
type IRadiationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRadiationContext differentiates from other interfaces.
	IsRadiationContext()
}

type RadiationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRadiationContext() *RadiationContext {
	var p = new(RadiationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RadiationParserRULE_radiation
	return p
}

func (*RadiationContext) IsRadiationContext() {}

func NewRadiationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RadiationContext {
	var p = new(RadiationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RadiationParserRULE_radiation

	return p
}

func (s *RadiationContext) GetParser() antlr.Parser { return s.parser }

func (s *RadiationContext) RADIATION() antlr.TerminalNode {
	return s.GetToken(RadiationParserRADIATION, 0)
}

func (s *RadiationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RadiationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RadiationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiationListener); ok {
		listenerT.EnterRadiation(s)
	}
}

func (s *RadiationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiationListener); ok {
		listenerT.ExitRadiation(s)
	}
}

func (p *RadiationParser) Radiation() (localctx IRadiationContext) {
	this := p
	_ = this

	localctx = NewRadiationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RadiationParserRULE_radiation)

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
		p.Match(RadiationParserRADIATION)
	}

	return localctx
}
