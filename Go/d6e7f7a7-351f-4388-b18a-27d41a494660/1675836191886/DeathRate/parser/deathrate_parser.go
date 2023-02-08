// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836191886/DeathRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DeathRate

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

type DeathRateParser struct {
	*antlr.BaseParser
}

var deathrateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func deathrateParserInit() {
	staticData := &deathrateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DEATHRATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "deathrate",
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

// DeathRateParserInit initializes any static state used to implement DeathRateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDeathRateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DeathRateParserInit() {
	staticData := &deathrateParserStaticData
	staticData.once.Do(deathrateParserInit)
}

// NewDeathRateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDeathRateParser(input antlr.TokenStream) *DeathRateParser {
	DeathRateParserInit()
	this := new(DeathRateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &deathrateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DeathRate.g4"

	return this
}

// DeathRateParser tokens.
const (
	DeathRateParserEOF       = antlr.TokenEOF
	DeathRateParserDEATHRATE = 1
	DeathRateParserOWNKEY    = 2
	DeathRateParserSPLITKEY  = 3
	DeathRateParserWS        = 4
)

// DeathRateParser rules.
const (
	DeathRateParserRULE_expression = 0
	DeathRateParserRULE_deathrate  = 1
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
	p.RuleIndex = DeathRateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeathRateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Deathrate() IDeathrateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeathrateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeathrateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DeathRateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeathRateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeathRateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DeathRateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DeathRateParserRULE_expression)

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
		p.Deathrate()
	}
	{
		p.SetState(5)
		p.Match(DeathRateParserEOF)
	}

	return localctx
}

// IDeathrateContext is an interface to support dynamic dispatch.
type IDeathrateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeathrateContext differentiates from other interfaces.
	IsDeathrateContext()
}

type DeathrateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeathrateContext() *DeathrateContext {
	var p = new(DeathrateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeathRateParserRULE_deathrate
	return p
}

func (*DeathrateContext) IsDeathrateContext() {}

func NewDeathrateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeathrateContext {
	var p = new(DeathrateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeathRateParserRULE_deathrate

	return p
}

func (s *DeathrateContext) GetParser() antlr.Parser { return s.parser }

func (s *DeathrateContext) DEATHRATE() antlr.TerminalNode {
	return s.GetToken(DeathRateParserDEATHRATE, 0)
}

func (s *DeathrateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeathrateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeathrateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeathRateListener); ok {
		listenerT.EnterDeathrate(s)
	}
}

func (s *DeathrateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeathRateListener); ok {
		listenerT.ExitDeathrate(s)
	}
}

func (p *DeathRateParser) Deathrate() (localctx IDeathrateContext) {
	this := p
	_ = this

	localctx = NewDeathrateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DeathRateParserRULE_deathrate)

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
		p.Match(DeathRateParserDEATHRATE)
	}

	return localctx
}
