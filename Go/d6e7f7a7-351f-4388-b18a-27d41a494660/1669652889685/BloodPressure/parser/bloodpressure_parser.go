// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652889685/BloodPressure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodPressure

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

type BloodPressureParser struct {
	*antlr.BaseParser
}

var bloodpressureParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bloodpressureParserInit() {
	staticData := &bloodpressureParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BLOODPRESSURE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bloodpressure",
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

// BloodPressureParserInit initializes any static state used to implement BloodPressureParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBloodPressureParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BloodPressureParserInit() {
	staticData := &bloodpressureParserStaticData
	staticData.once.Do(bloodpressureParserInit)
}

// NewBloodPressureParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBloodPressureParser(input antlr.TokenStream) *BloodPressureParser {
	BloodPressureParserInit()
	this := new(BloodPressureParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bloodpressureParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BloodPressure.g4"

	return this
}

// BloodPressureParser tokens.
const (
	BloodPressureParserEOF           = antlr.TokenEOF
	BloodPressureParserBLOODPRESSURE = 1
	BloodPressureParserOWNKEY        = 2
	BloodPressureParserSPLITKEY      = 3
	BloodPressureParserWS            = 4
)

// BloodPressureParser rules.
const (
	BloodPressureParserRULE_expression    = 0
	BloodPressureParserRULE_bloodpressure = 1
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
	p.RuleIndex = BloodPressureParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BloodPressureParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bloodpressure() IBloodpressureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloodpressureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloodpressureContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BloodPressureParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodPressureListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodPressureListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BloodPressureParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BloodPressureParserRULE_expression)

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
		p.Bloodpressure()
	}
	{
		p.SetState(5)
		p.Match(BloodPressureParserEOF)
	}

	return localctx
}

// IBloodpressureContext is an interface to support dynamic dispatch.
type IBloodpressureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloodpressureContext differentiates from other interfaces.
	IsBloodpressureContext()
}

type BloodpressureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloodpressureContext() *BloodpressureContext {
	var p = new(BloodpressureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BloodPressureParserRULE_bloodpressure
	return p
}

func (*BloodpressureContext) IsBloodpressureContext() {}

func NewBloodpressureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloodpressureContext {
	var p = new(BloodpressureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BloodPressureParserRULE_bloodpressure

	return p
}

func (s *BloodpressureContext) GetParser() antlr.Parser { return s.parser }

func (s *BloodpressureContext) BLOODPRESSURE() antlr.TerminalNode {
	return s.GetToken(BloodPressureParserBLOODPRESSURE, 0)
}

func (s *BloodpressureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloodpressureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloodpressureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodPressureListener); ok {
		listenerT.EnterBloodpressure(s)
	}
}

func (s *BloodpressureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodPressureListener); ok {
		listenerT.ExitBloodpressure(s)
	}
}

func (p *BloodPressureParser) Bloodpressure() (localctx IBloodpressureContext) {
	this := p
	_ = this

	localctx = NewBloodpressureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BloodPressureParserRULE_bloodpressure)

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
		p.Match(BloodPressureParserBLOODPRESSURE)
	}

	return localctx
}
