// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624525960/Measure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Measure

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

type MeasureParser struct {
	*antlr.BaseParser
}

var measureParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func measureParserInit() {
	staticData := &measureParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MEASURE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "measure",
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

// MeasureParserInit initializes any static state used to implement MeasureParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMeasureParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MeasureParserInit() {
	staticData := &measureParserStaticData
	staticData.once.Do(measureParserInit)
}

// NewMeasureParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMeasureParser(input antlr.TokenStream) *MeasureParser {
	MeasureParserInit()
	this := new(MeasureParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &measureParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Measure.g4"

	return this
}

// MeasureParser tokens.
const (
	MeasureParserEOF      = antlr.TokenEOF
	MeasureParserMEASURE  = 1
	MeasureParserOWNKEY   = 2
	MeasureParserSPLITKEY = 3
	MeasureParserWS       = 4
)

// MeasureParser rules.
const (
	MeasureParserRULE_expression = 0
	MeasureParserRULE_measure    = 1
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
	p.RuleIndex = MeasureParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MeasureParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Measure() IMeasureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeasureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeasureContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MeasureParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeasureListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeasureListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MeasureParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MeasureParserRULE_expression)

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
		p.Measure()
	}
	{
		p.SetState(5)
		p.Match(MeasureParserEOF)
	}

	return localctx
}

// IMeasureContext is an interface to support dynamic dispatch.
type IMeasureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeasureContext differentiates from other interfaces.
	IsMeasureContext()
}

type MeasureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasureContext() *MeasureContext {
	var p = new(MeasureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MeasureParserRULE_measure
	return p
}

func (*MeasureContext) IsMeasureContext() {}

func NewMeasureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasureContext {
	var p = new(MeasureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MeasureParserRULE_measure

	return p
}

func (s *MeasureContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasureContext) MEASURE() antlr.TerminalNode {
	return s.GetToken(MeasureParserMEASURE, 0)
}

func (s *MeasureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeasureListener); ok {
		listenerT.EnterMeasure(s)
	}
}

func (s *MeasureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeasureListener); ok {
		listenerT.ExitMeasure(s)
	}
}

func (p *MeasureParser) Measure() (localctx IMeasureContext) {
	this := p
	_ = this

	localctx = NewMeasureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MeasureParserRULE_measure)

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
		p.Match(MeasureParserMEASURE)
	}

	return localctx
}
