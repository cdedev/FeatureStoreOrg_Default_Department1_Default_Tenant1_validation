// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673856910412/Efficiency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Efficiency

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

type EfficiencyParser struct {
	*antlr.BaseParser
}

var efficiencyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func efficiencyParserInit() {
	staticData := &efficiencyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EFFICIENCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "efficiency",
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

// EfficiencyParserInit initializes any static state used to implement EfficiencyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEfficiencyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EfficiencyParserInit() {
	staticData := &efficiencyParserStaticData
	staticData.once.Do(efficiencyParserInit)
}

// NewEfficiencyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEfficiencyParser(input antlr.TokenStream) *EfficiencyParser {
	EfficiencyParserInit()
	this := new(EfficiencyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &efficiencyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Efficiency.g4"

	return this
}

// EfficiencyParser tokens.
const (
	EfficiencyParserEOF        = antlr.TokenEOF
	EfficiencyParserEFFICIENCY = 1
	EfficiencyParserOWNKEY     = 2
	EfficiencyParserSPLITKEY   = 3
	EfficiencyParserWS         = 4
)

// EfficiencyParser rules.
const (
	EfficiencyParserRULE_expression = 0
	EfficiencyParserRULE_efficiency = 1
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
	p.RuleIndex = EfficiencyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EfficiencyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Efficiency() IEfficiencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEfficiencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEfficiencyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EfficiencyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EfficiencyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EfficiencyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EfficiencyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EfficiencyParserRULE_expression)

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
		p.Efficiency()
	}
	{
		p.SetState(5)
		p.Match(EfficiencyParserEOF)
	}

	return localctx
}

// IEfficiencyContext is an interface to support dynamic dispatch.
type IEfficiencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEfficiencyContext differentiates from other interfaces.
	IsEfficiencyContext()
}

type EfficiencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEfficiencyContext() *EfficiencyContext {
	var p = new(EfficiencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EfficiencyParserRULE_efficiency
	return p
}

func (*EfficiencyContext) IsEfficiencyContext() {}

func NewEfficiencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EfficiencyContext {
	var p = new(EfficiencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EfficiencyParserRULE_efficiency

	return p
}

func (s *EfficiencyContext) GetParser() antlr.Parser { return s.parser }

func (s *EfficiencyContext) EFFICIENCY() antlr.TerminalNode {
	return s.GetToken(EfficiencyParserEFFICIENCY, 0)
}

func (s *EfficiencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EfficiencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EfficiencyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EfficiencyListener); ok {
		listenerT.EnterEfficiency(s)
	}
}

func (s *EfficiencyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EfficiencyListener); ok {
		listenerT.ExitEfficiency(s)
	}
}

func (p *EfficiencyParser) Efficiency() (localctx IEfficiencyContext) {
	this := p
	_ = this

	localctx = NewEfficiencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EfficiencyParserRULE_efficiency)

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
		p.Match(EfficiencyParserEFFICIENCY)
	}

	return localctx
}
