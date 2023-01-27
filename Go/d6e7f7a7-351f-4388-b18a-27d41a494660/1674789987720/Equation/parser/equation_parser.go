// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789987720/Equation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Equation

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

type EquationParser struct {
	*antlr.BaseParser
}

var equationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func equationParserInit() {
	staticData := &equationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EQUATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "equation",
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

// EquationParserInit initializes any static state used to implement EquationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEquationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EquationParserInit() {
	staticData := &equationParserStaticData
	staticData.once.Do(equationParserInit)
}

// NewEquationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEquationParser(input antlr.TokenStream) *EquationParser {
	EquationParserInit()
	this := new(EquationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &equationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Equation.g4"

	return this
}

// EquationParser tokens.
const (
	EquationParserEOF      = antlr.TokenEOF
	EquationParserEQUATION = 1
	EquationParserOWNKEY   = 2
	EquationParserSPLITKEY = 3
	EquationParserWS       = 4
)

// EquationParser rules.
const (
	EquationParserRULE_expression = 0
	EquationParserRULE_equation   = 1
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
	p.RuleIndex = EquationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EquationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Equation() IEquationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEquationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEquationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EquationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EquationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EquationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EquationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EquationParserRULE_expression)

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
		p.Equation()
	}
	{
		p.SetState(5)
		p.Match(EquationParserEOF)
	}

	return localctx
}

// IEquationContext is an interface to support dynamic dispatch.
type IEquationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquationContext differentiates from other interfaces.
	IsEquationContext()
}

type EquationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquationContext() *EquationContext {
	var p = new(EquationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EquationParserRULE_equation
	return p
}

func (*EquationContext) IsEquationContext() {}

func NewEquationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EquationContext {
	var p = new(EquationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EquationParserRULE_equation

	return p
}

func (s *EquationContext) GetParser() antlr.Parser { return s.parser }

func (s *EquationContext) EQUATION() antlr.TerminalNode {
	return s.GetToken(EquationParserEQUATION, 0)
}

func (s *EquationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EquationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EquationListener); ok {
		listenerT.EnterEquation(s)
	}
}

func (s *EquationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EquationListener); ok {
		listenerT.ExitEquation(s)
	}
}

func (p *EquationParser) Equation() (localctx IEquationContext) {
	this := p
	_ = this

	localctx = NewEquationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EquationParserRULE_equation)

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
		p.Match(EquationParserEQUATION)
	}

	return localctx
}
