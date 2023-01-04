// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672805602866/HyperTension.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HyperTension

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

type HyperTensionParser struct {
	*antlr.BaseParser
}

var hypertensionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hypertensionParserInit() {
	staticData := &hypertensionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HYPERTENSION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hypertension",
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

// HyperTensionParserInit initializes any static state used to implement HyperTensionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHyperTensionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HyperTensionParserInit() {
	staticData := &hypertensionParserStaticData
	staticData.once.Do(hypertensionParserInit)
}

// NewHyperTensionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHyperTensionParser(input antlr.TokenStream) *HyperTensionParser {
	HyperTensionParserInit()
	this := new(HyperTensionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hypertensionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "HyperTension.g4"

	return this
}

// HyperTensionParser tokens.
const (
	HyperTensionParserEOF          = antlr.TokenEOF
	HyperTensionParserHYPERTENSION = 1
	HyperTensionParserOWNKEY       = 2
	HyperTensionParserSPLITKEY     = 3
	HyperTensionParserWS           = 4
)

// HyperTensionParser rules.
const (
	HyperTensionParserRULE_expression   = 0
	HyperTensionParserRULE_hypertension = 1
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
	p.RuleIndex = HyperTensionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HyperTensionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hypertension() IHypertensionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHypertensionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHypertensionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HyperTensionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HyperTensionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HyperTensionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HyperTensionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HyperTensionParserRULE_expression)

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
		p.Hypertension()
	}
	{
		p.SetState(5)
		p.Match(HyperTensionParserEOF)
	}

	return localctx
}

// IHypertensionContext is an interface to support dynamic dispatch.
type IHypertensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHypertensionContext differentiates from other interfaces.
	IsHypertensionContext()
}

type HypertensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHypertensionContext() *HypertensionContext {
	var p = new(HypertensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HyperTensionParserRULE_hypertension
	return p
}

func (*HypertensionContext) IsHypertensionContext() {}

func NewHypertensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HypertensionContext {
	var p = new(HypertensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HyperTensionParserRULE_hypertension

	return p
}

func (s *HypertensionContext) GetParser() antlr.Parser { return s.parser }

func (s *HypertensionContext) HYPERTENSION() antlr.TerminalNode {
	return s.GetToken(HyperTensionParserHYPERTENSION, 0)
}

func (s *HypertensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HypertensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HypertensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HyperTensionListener); ok {
		listenerT.EnterHypertension(s)
	}
}

func (s *HypertensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HyperTensionListener); ok {
		listenerT.ExitHypertension(s)
	}
}

func (p *HyperTensionParser) Hypertension() (localctx IHypertensionContext) {
	this := p
	_ = this

	localctx = NewHypertensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HyperTensionParserRULE_hypertension)

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
		p.Match(HyperTensionParserHYPERTENSION)
	}

	return localctx
}
