// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676258983667/Acceleration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acceleration

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

type AccelerationParser struct {
	*antlr.BaseParser
}

var accelerationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func accelerationParserInit() {
	staticData := &accelerationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ACCELERATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "acceleration",
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

// AccelerationParserInit initializes any static state used to implement AccelerationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAccelerationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AccelerationParserInit() {
	staticData := &accelerationParserStaticData
	staticData.once.Do(accelerationParserInit)
}

// NewAccelerationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAccelerationParser(input antlr.TokenStream) *AccelerationParser {
	AccelerationParserInit()
	this := new(AccelerationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &accelerationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Acceleration.g4"

	return this
}

// AccelerationParser tokens.
const (
	AccelerationParserEOF          = antlr.TokenEOF
	AccelerationParserACCELERATION = 1
	AccelerationParserOWNKEY       = 2
	AccelerationParserSPLITKEY     = 3
	AccelerationParserWS           = 4
)

// AccelerationParser rules.
const (
	AccelerationParserRULE_expression   = 0
	AccelerationParserRULE_acceleration = 1
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
	p.RuleIndex = AccelerationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AccelerationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Acceleration() IAccelerationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccelerationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccelerationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AccelerationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AccelerationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AccelerationParserRULE_expression)

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
		p.Acceleration()
	}
	{
		p.SetState(5)
		p.Match(AccelerationParserEOF)
	}

	return localctx
}

// IAccelerationContext is an interface to support dynamic dispatch.
type IAccelerationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccelerationContext differentiates from other interfaces.
	IsAccelerationContext()
}

type AccelerationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccelerationContext() *AccelerationContext {
	var p = new(AccelerationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AccelerationParserRULE_acceleration
	return p
}

func (*AccelerationContext) IsAccelerationContext() {}

func NewAccelerationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccelerationContext {
	var p = new(AccelerationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AccelerationParserRULE_acceleration

	return p
}

func (s *AccelerationContext) GetParser() antlr.Parser { return s.parser }

func (s *AccelerationContext) ACCELERATION() antlr.TerminalNode {
	return s.GetToken(AccelerationParserACCELERATION, 0)
}

func (s *AccelerationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccelerationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccelerationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerationListener); ok {
		listenerT.EnterAcceleration(s)
	}
}

func (s *AccelerationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerationListener); ok {
		listenerT.ExitAcceleration(s)
	}
}

func (p *AccelerationParser) Acceleration() (localctx IAccelerationContext) {
	this := p
	_ = this

	localctx = NewAccelerationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AccelerationParserRULE_acceleration)

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
		p.Match(AccelerationParserACCELERATION)
	}

	return localctx
}
