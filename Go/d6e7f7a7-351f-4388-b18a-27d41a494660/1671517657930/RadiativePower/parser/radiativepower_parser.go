// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517657930/RadiativePower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RadiativePower

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

type RadiativePowerParser struct {
	*antlr.BaseParser
}

var radiativepowerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func radiativepowerParserInit() {
	staticData := &radiativepowerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RADIATIVEPOWER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "radiativepower",
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

// RadiativePowerParserInit initializes any static state used to implement RadiativePowerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRadiativePowerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RadiativePowerParserInit() {
	staticData := &radiativepowerParserStaticData
	staticData.once.Do(radiativepowerParserInit)
}

// NewRadiativePowerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRadiativePowerParser(input antlr.TokenStream) *RadiativePowerParser {
	RadiativePowerParserInit()
	this := new(RadiativePowerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &radiativepowerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RadiativePower.g4"

	return this
}

// RadiativePowerParser tokens.
const (
	RadiativePowerParserEOF            = antlr.TokenEOF
	RadiativePowerParserRADIATIVEPOWER = 1
	RadiativePowerParserOWNKEY         = 2
	RadiativePowerParserSPLITKEY       = 3
	RadiativePowerParserWS             = 4
)

// RadiativePowerParser rules.
const (
	RadiativePowerParserRULE_expression     = 0
	RadiativePowerParserRULE_radiativepower = 1
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
	p.RuleIndex = RadiativePowerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RadiativePowerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Radiativepower() IRadiativepowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRadiativepowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRadiativepowerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RadiativePowerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiativePowerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiativePowerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RadiativePowerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RadiativePowerParserRULE_expression)

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
		p.Radiativepower()
	}
	{
		p.SetState(5)
		p.Match(RadiativePowerParserEOF)
	}

	return localctx
}

// IRadiativepowerContext is an interface to support dynamic dispatch.
type IRadiativepowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRadiativepowerContext differentiates from other interfaces.
	IsRadiativepowerContext()
}

type RadiativepowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRadiativepowerContext() *RadiativepowerContext {
	var p = new(RadiativepowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RadiativePowerParserRULE_radiativepower
	return p
}

func (*RadiativepowerContext) IsRadiativepowerContext() {}

func NewRadiativepowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RadiativepowerContext {
	var p = new(RadiativepowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RadiativePowerParserRULE_radiativepower

	return p
}

func (s *RadiativepowerContext) GetParser() antlr.Parser { return s.parser }

func (s *RadiativepowerContext) RADIATIVEPOWER() antlr.TerminalNode {
	return s.GetToken(RadiativePowerParserRADIATIVEPOWER, 0)
}

func (s *RadiativepowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RadiativepowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RadiativepowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiativePowerListener); ok {
		listenerT.EnterRadiativepower(s)
	}
}

func (s *RadiativepowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RadiativePowerListener); ok {
		listenerT.ExitRadiativepower(s)
	}
}

func (p *RadiativePowerParser) Radiativepower() (localctx IRadiativepowerContext) {
	this := p
	_ = this

	localctx = NewRadiativepowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RadiativePowerParserRULE_radiativepower)

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
		p.Match(RadiativePowerParserRADIATIVEPOWER)
	}

	return localctx
}
