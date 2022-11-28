// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659170926/SquareFeet.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SquareFeet

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

type SquareFeetParser struct {
	*antlr.BaseParser
}

var squarefeetParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func squarefeetParserInit() {
	staticData := &squarefeetParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SQUAREFEET", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "squarefeet",
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

// SquareFeetParserInit initializes any static state used to implement SquareFeetParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSquareFeetParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SquareFeetParserInit() {
	staticData := &squarefeetParserStaticData
	staticData.once.Do(squarefeetParserInit)
}

// NewSquareFeetParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSquareFeetParser(input antlr.TokenStream) *SquareFeetParser {
	SquareFeetParserInit()
	this := new(SquareFeetParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &squarefeetParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SquareFeet.g4"

	return this
}

// SquareFeetParser tokens.
const (
	SquareFeetParserEOF        = antlr.TokenEOF
	SquareFeetParserSQUAREFEET = 1
	SquareFeetParserOWNKEY     = 2
	SquareFeetParserSPLITKEY   = 3
	SquareFeetParserWS         = 4
)

// SquareFeetParser rules.
const (
	SquareFeetParserRULE_expression = 0
	SquareFeetParserRULE_squarefeet = 1
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
	p.RuleIndex = SquareFeetParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SquareFeetParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Squarefeet() ISquarefeetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISquarefeetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISquarefeetContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SquareFeetParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SquareFeetListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SquareFeetListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SquareFeetParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SquareFeetParserRULE_expression)

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
		p.Squarefeet()
	}
	{
		p.SetState(5)
		p.Match(SquareFeetParserEOF)
	}

	return localctx
}

// ISquarefeetContext is an interface to support dynamic dispatch.
type ISquarefeetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSquarefeetContext differentiates from other interfaces.
	IsSquarefeetContext()
}

type SquarefeetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySquarefeetContext() *SquarefeetContext {
	var p = new(SquarefeetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SquareFeetParserRULE_squarefeet
	return p
}

func (*SquarefeetContext) IsSquarefeetContext() {}

func NewSquarefeetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SquarefeetContext {
	var p = new(SquarefeetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SquareFeetParserRULE_squarefeet

	return p
}

func (s *SquarefeetContext) GetParser() antlr.Parser { return s.parser }

func (s *SquarefeetContext) SQUAREFEET() antlr.TerminalNode {
	return s.GetToken(SquareFeetParserSQUAREFEET, 0)
}

func (s *SquarefeetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SquarefeetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SquarefeetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SquareFeetListener); ok {
		listenerT.EnterSquarefeet(s)
	}
}

func (s *SquarefeetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SquareFeetListener); ok {
		listenerT.ExitSquarefeet(s)
	}
}

func (p *SquareFeetParser) Squarefeet() (localctx ISquarefeetContext) {
	this := p
	_ = this

	localctx = NewSquarefeetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SquareFeetParserRULE_squarefeet)

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
		p.Match(SquareFeetParserSQUAREFEET)
	}

	return localctx
}
