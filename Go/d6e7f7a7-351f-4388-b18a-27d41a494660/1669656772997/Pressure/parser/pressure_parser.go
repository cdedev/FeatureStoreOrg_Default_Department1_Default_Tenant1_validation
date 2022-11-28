// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656772997/Pressure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pressure

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

type PressureParser struct {
	*antlr.BaseParser
}

var pressureParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pressureParserInit() {
	staticData := &pressureParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PRESSURE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pressure",
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

// PressureParserInit initializes any static state used to implement PressureParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPressureParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PressureParserInit() {
	staticData := &pressureParserStaticData
	staticData.once.Do(pressureParserInit)
}

// NewPressureParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPressureParser(input antlr.TokenStream) *PressureParser {
	PressureParserInit()
	this := new(PressureParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pressureParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Pressure.g4"

	return this
}

// PressureParser tokens.
const (
	PressureParserEOF      = antlr.TokenEOF
	PressureParserPRESSURE = 1
	PressureParserOWNKEY   = 2
	PressureParserSPLITKEY = 3
	PressureParserWS       = 4
)

// PressureParser rules.
const (
	PressureParserRULE_expression = 0
	PressureParserRULE_pressure   = 1
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
	p.RuleIndex = PressureParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PressureParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pressure() IPressureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPressureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPressureContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PressureParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PressureParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PressureParserRULE_expression)

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
		p.Pressure()
	}
	{
		p.SetState(5)
		p.Match(PressureParserEOF)
	}

	return localctx
}

// IPressureContext is an interface to support dynamic dispatch.
type IPressureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPressureContext differentiates from other interfaces.
	IsPressureContext()
}

type PressureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPressureContext() *PressureContext {
	var p = new(PressureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PressureParserRULE_pressure
	return p
}

func (*PressureContext) IsPressureContext() {}

func NewPressureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PressureContext {
	var p = new(PressureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PressureParserRULE_pressure

	return p
}

func (s *PressureContext) GetParser() antlr.Parser { return s.parser }

func (s *PressureContext) PRESSURE() antlr.TerminalNode {
	return s.GetToken(PressureParserPRESSURE, 0)
}

func (s *PressureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PressureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PressureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureListener); ok {
		listenerT.EnterPressure(s)
	}
}

func (s *PressureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureListener); ok {
		listenerT.ExitPressure(s)
	}
}

func (p *PressureParser) Pressure() (localctx IPressureContext) {
	this := p
	_ = this

	localctx = NewPressureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PressureParserRULE_pressure)

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
		p.Match(PressureParserPRESSURE)
	}

	return localctx
}
