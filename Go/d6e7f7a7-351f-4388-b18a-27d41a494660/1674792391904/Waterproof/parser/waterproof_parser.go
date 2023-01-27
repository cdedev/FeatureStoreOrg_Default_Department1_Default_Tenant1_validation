// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674792391904/Waterproof.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Waterproof

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

type WaterproofParser struct {
	*antlr.BaseParser
}

var waterproofParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func waterproofParserInit() {
	staticData := &waterproofParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WATERPROOF", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "waterproof",
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

// WaterproofParserInit initializes any static state used to implement WaterproofParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWaterproofParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WaterproofParserInit() {
	staticData := &waterproofParserStaticData
	staticData.once.Do(waterproofParserInit)
}

// NewWaterproofParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWaterproofParser(input antlr.TokenStream) *WaterproofParser {
	WaterproofParserInit()
	this := new(WaterproofParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &waterproofParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Waterproof.g4"

	return this
}

// WaterproofParser tokens.
const (
	WaterproofParserEOF        = antlr.TokenEOF
	WaterproofParserWATERPROOF = 1
	WaterproofParserOWNKEY     = 2
	WaterproofParserSPLITKEY   = 3
	WaterproofParserWS         = 4
)

// WaterproofParser rules.
const (
	WaterproofParserRULE_expression = 0
	WaterproofParserRULE_waterproof = 1
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
	p.RuleIndex = WaterproofParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WaterproofParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Waterproof() IWaterproofContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWaterproofContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWaterproofContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WaterproofParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WaterproofListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WaterproofListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WaterproofParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WaterproofParserRULE_expression)

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
		p.Waterproof()
	}
	{
		p.SetState(5)
		p.Match(WaterproofParserEOF)
	}

	return localctx
}

// IWaterproofContext is an interface to support dynamic dispatch.
type IWaterproofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWaterproofContext differentiates from other interfaces.
	IsWaterproofContext()
}

type WaterproofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWaterproofContext() *WaterproofContext {
	var p = new(WaterproofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WaterproofParserRULE_waterproof
	return p
}

func (*WaterproofContext) IsWaterproofContext() {}

func NewWaterproofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WaterproofContext {
	var p = new(WaterproofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WaterproofParserRULE_waterproof

	return p
}

func (s *WaterproofContext) GetParser() antlr.Parser { return s.parser }

func (s *WaterproofContext) WATERPROOF() antlr.TerminalNode {
	return s.GetToken(WaterproofParserWATERPROOF, 0)
}

func (s *WaterproofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WaterproofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WaterproofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WaterproofListener); ok {
		listenerT.EnterWaterproof(s)
	}
}

func (s *WaterproofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WaterproofListener); ok {
		listenerT.ExitWaterproof(s)
	}
}

func (p *WaterproofParser) Waterproof() (localctx IWaterproofContext) {
	this := p
	_ = this

	localctx = NewWaterproofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WaterproofParserRULE_waterproof)

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
		p.Match(WaterproofParserWATERPROOF)
	}

	return localctx
}
