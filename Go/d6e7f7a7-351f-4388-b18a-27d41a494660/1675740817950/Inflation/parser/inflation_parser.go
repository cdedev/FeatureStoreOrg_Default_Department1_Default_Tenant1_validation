// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675740817950/Inflation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Inflation

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

type InflationParser struct {
	*antlr.BaseParser
}

var inflationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func inflationParserInit() {
	staticData := &inflationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INFLATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "inflation",
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

// InflationParserInit initializes any static state used to implement InflationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInflationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InflationParserInit() {
	staticData := &inflationParserStaticData
	staticData.once.Do(inflationParserInit)
}

// NewInflationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInflationParser(input antlr.TokenStream) *InflationParser {
	InflationParserInit()
	this := new(InflationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &inflationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Inflation.g4"

	return this
}

// InflationParser tokens.
const (
	InflationParserEOF       = antlr.TokenEOF
	InflationParserINFLATION = 1
	InflationParserOWNKEY    = 2
	InflationParserSPLITKEY  = 3
	InflationParserWS        = 4
)

// InflationParser rules.
const (
	InflationParserRULE_expression = 0
	InflationParserRULE_inflation  = 1
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
	p.RuleIndex = InflationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InflationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Inflation() IInflationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInflationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInflationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(InflationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InflationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InflationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *InflationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InflationParserRULE_expression)

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
		p.Inflation()
	}
	{
		p.SetState(5)
		p.Match(InflationParserEOF)
	}

	return localctx
}

// IInflationContext is an interface to support dynamic dispatch.
type IInflationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInflationContext differentiates from other interfaces.
	IsInflationContext()
}

type InflationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInflationContext() *InflationContext {
	var p = new(InflationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InflationParserRULE_inflation
	return p
}

func (*InflationContext) IsInflationContext() {}

func NewInflationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InflationContext {
	var p = new(InflationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InflationParserRULE_inflation

	return p
}

func (s *InflationContext) GetParser() antlr.Parser { return s.parser }

func (s *InflationContext) INFLATION() antlr.TerminalNode {
	return s.GetToken(InflationParserINFLATION, 0)
}

func (s *InflationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InflationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InflationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InflationListener); ok {
		listenerT.EnterInflation(s)
	}
}

func (s *InflationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InflationListener); ok {
		listenerT.ExitInflation(s)
	}
}

func (p *InflationParser) Inflation() (localctx IInflationContext) {
	this := p
	_ = this

	localctx = NewInflationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InflationParserRULE_inflation)

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
		p.Match(InflationParserINFLATION)
	}

	return localctx
}
