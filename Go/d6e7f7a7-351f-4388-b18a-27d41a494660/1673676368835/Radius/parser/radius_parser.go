// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676368835/Radius.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Radius

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

type RadiusParser struct {
	*antlr.BaseParser
}

var radiusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func radiusParserInit() {
	staticData := &radiusParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RADIUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "radius",
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

// RadiusParserInit initializes any static state used to implement RadiusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRadiusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RadiusParserInit() {
	staticData := &radiusParserStaticData
	staticData.once.Do(radiusParserInit)
}

// NewRadiusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRadiusParser(input antlr.TokenStream) *RadiusParser {
	RadiusParserInit()
	this := new(RadiusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &radiusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Radius.g4"

	return this
}

// RadiusParser tokens.
const (
	RadiusParserEOF      = antlr.TokenEOF
	RadiusParserRADIUS   = 1
	RadiusParserOWNKEY   = 2
	RadiusParserSPLITKEY = 3
	RadiusParserWS       = 4
)

// RadiusParser rules.
const (
	RadiusParserRULE_expression = 0
	RadiusParserRULE_radius     = 1
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
	p.RuleIndex = RadiusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RadiusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Radius() IRadiusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRadiusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRadiusContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RadiusParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RadiusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RadiusParserRULE_expression)

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
		p.Radius()
	}
	{
		p.SetState(5)
		p.Match(RadiusParserEOF)
	}

	return localctx
}

// IRadiusContext is an interface to support dynamic dispatch.
type IRadiusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRadiusContext differentiates from other interfaces.
	IsRadiusContext()
}

type RadiusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRadiusContext() *RadiusContext {
	var p = new(RadiusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RadiusParserRULE_radius
	return p
}

func (*RadiusContext) IsRadiusContext() {}

func NewRadiusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RadiusContext {
	var p = new(RadiusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RadiusParserRULE_radius

	return p
}

func (s *RadiusContext) GetParser() antlr.Parser { return s.parser }

func (s *RadiusContext) RADIUS() antlr.TerminalNode {
	return s.GetToken(RadiusParserRADIUS, 0)
}

func (s *RadiusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RadiusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RadiusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiusListener); ok {
		listenerT.EnterRadius(s)
	}
}

func (s *RadiusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiusListener); ok {
		listenerT.ExitRadius(s)
	}
}

func (p *RadiusParser) Radius() (localctx IRadiusContext) {
	this := p
	_ = this

	localctx = NewRadiusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RadiusParserRULE_radius)

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
		p.Match(RadiusParserRADIUS)
	}

	return localctx
}
