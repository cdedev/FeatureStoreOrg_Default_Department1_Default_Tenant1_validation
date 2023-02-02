// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675309170263/Cost.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cost

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

type CostParser struct {
	*antlr.BaseParser
}

var costParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func costParserInit() {
	staticData := &costParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COST", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cost",
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

// CostParserInit initializes any static state used to implement CostParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCostParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CostParserInit() {
	staticData := &costParserStaticData
	staticData.once.Do(costParserInit)
}

// NewCostParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCostParser(input antlr.TokenStream) *CostParser {
	CostParserInit()
	this := new(CostParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &costParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cost.g4"

	return this
}

// CostParser tokens.
const (
	CostParserEOF      = antlr.TokenEOF
	CostParserCOST     = 1
	CostParserOWNKEY   = 2
	CostParserSPLITKEY = 3
	CostParserWS       = 4
)

// CostParser rules.
const (
	CostParserRULE_expression = 0
	CostParserRULE_cost       = 1
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
	p.RuleIndex = CostParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CostParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cost() ICostContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICostContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICostContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CostParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CostListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CostListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CostParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CostParserRULE_expression)

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
		p.Cost()
	}
	{
		p.SetState(5)
		p.Match(CostParserEOF)
	}

	return localctx
}

// ICostContext is an interface to support dynamic dispatch.
type ICostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCostContext differentiates from other interfaces.
	IsCostContext()
}

type CostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCostContext() *CostContext {
	var p = new(CostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CostParserRULE_cost
	return p
}

func (*CostContext) IsCostContext() {}

func NewCostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CostContext {
	var p = new(CostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CostParserRULE_cost

	return p
}

func (s *CostContext) GetParser() antlr.Parser { return s.parser }

func (s *CostContext) COST() antlr.TerminalNode {
	return s.GetToken(CostParserCOST, 0)
}

func (s *CostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CostListener); ok {
		listenerT.EnterCost(s)
	}
}

func (s *CostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CostListener); ok {
		listenerT.ExitCost(s)
	}
}

func (p *CostParser) Cost() (localctx ICostContext) {
	this := p
	_ = this

	localctx = NewCostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CostParserRULE_cost)

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
		p.Match(CostParserCOST)
	}

	return localctx
}
