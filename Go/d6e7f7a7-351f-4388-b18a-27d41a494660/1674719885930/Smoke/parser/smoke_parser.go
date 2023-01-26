// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719885930/Smoke.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Smoke

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

type SmokeParser struct {
	*antlr.BaseParser
}

var smokeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func smokeParserInit() {
	staticData := &smokeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SMOKE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "smoke",
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

// SmokeParserInit initializes any static state used to implement SmokeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSmokeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SmokeParserInit() {
	staticData := &smokeParserStaticData
	staticData.once.Do(smokeParserInit)
}

// NewSmokeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSmokeParser(input antlr.TokenStream) *SmokeParser {
	SmokeParserInit()
	this := new(SmokeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &smokeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Smoke.g4"

	return this
}

// SmokeParser tokens.
const (
	SmokeParserEOF      = antlr.TokenEOF
	SmokeParserSMOKE    = 1
	SmokeParserOWNKEY   = 2
	SmokeParserSPLITKEY = 3
	SmokeParserWS       = 4
)

// SmokeParser rules.
const (
	SmokeParserRULE_expression = 0
	SmokeParserRULE_smoke      = 1
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
	p.RuleIndex = SmokeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmokeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Smoke() ISmokeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISmokeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISmokeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SmokeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SmokeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SmokeParserRULE_expression)

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
		p.Smoke()
	}
	{
		p.SetState(5)
		p.Match(SmokeParserEOF)
	}

	return localctx
}

// ISmokeContext is an interface to support dynamic dispatch.
type ISmokeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSmokeContext differentiates from other interfaces.
	IsSmokeContext()
}

type SmokeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySmokeContext() *SmokeContext {
	var p = new(SmokeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmokeParserRULE_smoke
	return p
}

func (*SmokeContext) IsSmokeContext() {}

func NewSmokeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SmokeContext {
	var p = new(SmokeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmokeParserRULE_smoke

	return p
}

func (s *SmokeContext) GetParser() antlr.Parser { return s.parser }

func (s *SmokeContext) SMOKE() antlr.TerminalNode {
	return s.GetToken(SmokeParserSMOKE, 0)
}

func (s *SmokeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SmokeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SmokeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokeListener); ok {
		listenerT.EnterSmoke(s)
	}
}

func (s *SmokeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokeListener); ok {
		listenerT.ExitSmoke(s)
	}
}

func (p *SmokeParser) Smoke() (localctx ISmokeContext) {
	this := p
	_ = this

	localctx = NewSmokeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SmokeParserRULE_smoke)

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
		p.Match(SmokeParserSMOKE)
	}

	return localctx
}
