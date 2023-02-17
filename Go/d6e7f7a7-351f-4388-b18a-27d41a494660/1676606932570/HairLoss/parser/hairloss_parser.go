// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606932570/HairLoss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HairLoss

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

type HairLossParser struct {
	*antlr.BaseParser
}

var hairlossParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hairlossParserInit() {
	staticData := &hairlossParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HAIRLOSS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hairloss",
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

// HairLossParserInit initializes any static state used to implement HairLossParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHairLossParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HairLossParserInit() {
	staticData := &hairlossParserStaticData
	staticData.once.Do(hairlossParserInit)
}

// NewHairLossParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHairLossParser(input antlr.TokenStream) *HairLossParser {
	HairLossParserInit()
	this := new(HairLossParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hairlossParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "HairLoss.g4"

	return this
}

// HairLossParser tokens.
const (
	HairLossParserEOF      = antlr.TokenEOF
	HairLossParserHAIRLOSS = 1
	HairLossParserOWNKEY   = 2
	HairLossParserSPLITKEY = 3
	HairLossParserWS       = 4
)

// HairLossParser rules.
const (
	HairLossParserRULE_expression = 0
	HairLossParserRULE_hairloss   = 1
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
	p.RuleIndex = HairLossParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HairLossParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hairloss() IHairlossContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHairlossContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHairlossContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HairLossParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HairLossListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HairLossListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HairLossParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HairLossParserRULE_expression)

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
		p.Hairloss()
	}
	{
		p.SetState(5)
		p.Match(HairLossParserEOF)
	}

	return localctx
}

// IHairlossContext is an interface to support dynamic dispatch.
type IHairlossContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHairlossContext differentiates from other interfaces.
	IsHairlossContext()
}

type HairlossContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHairlossContext() *HairlossContext {
	var p = new(HairlossContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HairLossParserRULE_hairloss
	return p
}

func (*HairlossContext) IsHairlossContext() {}

func NewHairlossContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HairlossContext {
	var p = new(HairlossContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HairLossParserRULE_hairloss

	return p
}

func (s *HairlossContext) GetParser() antlr.Parser { return s.parser }

func (s *HairlossContext) HAIRLOSS() antlr.TerminalNode {
	return s.GetToken(HairLossParserHAIRLOSS, 0)
}

func (s *HairlossContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HairlossContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HairlossContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HairLossListener); ok {
		listenerT.EnterHairloss(s)
	}
}

func (s *HairlossContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HairLossListener); ok {
		listenerT.ExitHairloss(s)
	}
}

func (p *HairLossParser) Hairloss() (localctx IHairlossContext) {
	this := p
	_ = this

	localctx = NewHairlossContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HairLossParserRULE_hairloss)

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
		p.Match(HairLossParserHAIRLOSS)
	}

	return localctx
}
