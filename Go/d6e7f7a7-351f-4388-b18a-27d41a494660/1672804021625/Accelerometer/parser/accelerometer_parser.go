// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804021625/Accelerometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Accelerometer

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

type AccelerometerParser struct {
	*antlr.BaseParser
}

var accelerometerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func accelerometerParserInit() {
	staticData := &accelerometerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ACCELEROMETER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "accelerometer",
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

// AccelerometerParserInit initializes any static state used to implement AccelerometerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAccelerometerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AccelerometerParserInit() {
	staticData := &accelerometerParserStaticData
	staticData.once.Do(accelerometerParserInit)
}

// NewAccelerometerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAccelerometerParser(input antlr.TokenStream) *AccelerometerParser {
	AccelerometerParserInit()
	this := new(AccelerometerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &accelerometerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Accelerometer.g4"

	return this
}

// AccelerometerParser tokens.
const (
	AccelerometerParserEOF           = antlr.TokenEOF
	AccelerometerParserACCELEROMETER = 1
	AccelerometerParserOWNKEY        = 2
	AccelerometerParserSPLITKEY      = 3
	AccelerometerParserWS            = 4
)

// AccelerometerParser rules.
const (
	AccelerometerParserRULE_expression    = 0
	AccelerometerParserRULE_accelerometer = 1
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
	p.RuleIndex = AccelerometerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AccelerometerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Accelerometer() IAccelerometerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccelerometerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccelerometerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AccelerometerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerometerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerometerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AccelerometerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AccelerometerParserRULE_expression)

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
		p.Accelerometer()
	}
	{
		p.SetState(5)
		p.Match(AccelerometerParserEOF)
	}

	return localctx
}

// IAccelerometerContext is an interface to support dynamic dispatch.
type IAccelerometerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccelerometerContext differentiates from other interfaces.
	IsAccelerometerContext()
}

type AccelerometerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccelerometerContext() *AccelerometerContext {
	var p = new(AccelerometerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AccelerometerParserRULE_accelerometer
	return p
}

func (*AccelerometerContext) IsAccelerometerContext() {}

func NewAccelerometerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccelerometerContext {
	var p = new(AccelerometerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AccelerometerParserRULE_accelerometer

	return p
}

func (s *AccelerometerContext) GetParser() antlr.Parser { return s.parser }

func (s *AccelerometerContext) ACCELEROMETER() antlr.TerminalNode {
	return s.GetToken(AccelerometerParserACCELEROMETER, 0)
}

func (s *AccelerometerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccelerometerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccelerometerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerometerListener); ok {
		listenerT.EnterAccelerometer(s)
	}
}

func (s *AccelerometerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AccelerometerListener); ok {
		listenerT.ExitAccelerometer(s)
	}
}

func (p *AccelerometerParser) Accelerometer() (localctx IAccelerometerContext) {
	this := p
	_ = this

	localctx = NewAccelerometerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AccelerometerParserRULE_accelerometer)

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
		p.Match(AccelerometerParserACCELEROMETER)
	}

	return localctx
}
