// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672286056324/Haul.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Haul

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

type HaulParser struct {
	*antlr.BaseParser
}

var haulParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func haulParserInit() {
	staticData := &haulParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HAUL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "haul",
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

// HaulParserInit initializes any static state used to implement HaulParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHaulParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HaulParserInit() {
	staticData := &haulParserStaticData
	staticData.once.Do(haulParserInit)
}

// NewHaulParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHaulParser(input antlr.TokenStream) *HaulParser {
	HaulParserInit()
	this := new(HaulParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &haulParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Haul.g4"

	return this
}

// HaulParser tokens.
const (
	HaulParserEOF      = antlr.TokenEOF
	HaulParserHAUL     = 1
	HaulParserOWNKEY   = 2
	HaulParserSPLITKEY = 3
	HaulParserWS       = 4
)

// HaulParser rules.
const (
	HaulParserRULE_expression = 0
	HaulParserRULE_haul       = 1
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
	p.RuleIndex = HaulParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HaulParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Haul() IHaulContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHaulContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHaulContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HaulParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HaulListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HaulListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HaulParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HaulParserRULE_expression)

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
		p.Haul()
	}
	{
		p.SetState(5)
		p.Match(HaulParserEOF)
	}

	return localctx
}

// IHaulContext is an interface to support dynamic dispatch.
type IHaulContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHaulContext differentiates from other interfaces.
	IsHaulContext()
}

type HaulContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHaulContext() *HaulContext {
	var p = new(HaulContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HaulParserRULE_haul
	return p
}

func (*HaulContext) IsHaulContext() {}

func NewHaulContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HaulContext {
	var p = new(HaulContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HaulParserRULE_haul

	return p
}

func (s *HaulContext) GetParser() antlr.Parser { return s.parser }

func (s *HaulContext) HAUL() antlr.TerminalNode {
	return s.GetToken(HaulParserHAUL, 0)
}

func (s *HaulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HaulContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HaulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HaulListener); ok {
		listenerT.EnterHaul(s)
	}
}

func (s *HaulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HaulListener); ok {
		listenerT.ExitHaul(s)
	}
}

func (p *HaulParser) Haul() (localctx IHaulContext) {
	this := p
	_ = this

	localctx = NewHaulContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HaulParserRULE_haul)

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
		p.Match(HaulParserHAUL)
	}

	return localctx
}
