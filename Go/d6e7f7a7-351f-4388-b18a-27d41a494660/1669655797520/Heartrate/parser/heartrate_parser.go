// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669655797520/Heartrate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Heartrate

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

type HeartrateParser struct {
	*antlr.BaseParser
}

var heartrateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func heartrateParserInit() {
	staticData := &heartrateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HEARTRATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "heartrate",
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

// HeartrateParserInit initializes any static state used to implement HeartrateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHeartrateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HeartrateParserInit() {
	staticData := &heartrateParserStaticData
	staticData.once.Do(heartrateParserInit)
}

// NewHeartrateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHeartrateParser(input antlr.TokenStream) *HeartrateParser {
	HeartrateParserInit()
	this := new(HeartrateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &heartrateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Heartrate.g4"

	return this
}

// HeartrateParser tokens.
const (
	HeartrateParserEOF       = antlr.TokenEOF
	HeartrateParserHEARTRATE = 1
	HeartrateParserOWNKEY    = 2
	HeartrateParserSPLITKEY  = 3
	HeartrateParserWS        = 4
)

// HeartrateParser rules.
const (
	HeartrateParserRULE_expression = 0
	HeartrateParserRULE_heartrate  = 1
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
	p.RuleIndex = HeartrateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HeartrateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Heartrate() IHeartrateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHeartrateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHeartrateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HeartrateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HeartrateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HeartrateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HeartrateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HeartrateParserRULE_expression)

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
		p.Heartrate()
	}
	{
		p.SetState(5)
		p.Match(HeartrateParserEOF)
	}

	return localctx
}

// IHeartrateContext is an interface to support dynamic dispatch.
type IHeartrateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeartrateContext differentiates from other interfaces.
	IsHeartrateContext()
}

type HeartrateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeartrateContext() *HeartrateContext {
	var p = new(HeartrateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HeartrateParserRULE_heartrate
	return p
}

func (*HeartrateContext) IsHeartrateContext() {}

func NewHeartrateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeartrateContext {
	var p = new(HeartrateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HeartrateParserRULE_heartrate

	return p
}

func (s *HeartrateContext) GetParser() antlr.Parser { return s.parser }

func (s *HeartrateContext) HEARTRATE() antlr.TerminalNode {
	return s.GetToken(HeartrateParserHEARTRATE, 0)
}

func (s *HeartrateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeartrateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeartrateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HeartrateListener); ok {
		listenerT.EnterHeartrate(s)
	}
}

func (s *HeartrateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HeartrateListener); ok {
		listenerT.ExitHeartrate(s)
	}
}

func (p *HeartrateParser) Heartrate() (localctx IHeartrateContext) {
	this := p
	_ = this

	localctx = NewHeartrateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HeartrateParserRULE_heartrate)

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
		p.Match(HeartrateParserHEARTRATE)
	}

	return localctx
}
