// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675923077966/Variance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Variance

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

type VarianceParser struct {
	*antlr.BaseParser
}

var varianceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func varianceParserInit() {
	staticData := &varianceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VARIANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "variance",
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

// VarianceParserInit initializes any static state used to implement VarianceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVarianceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VarianceParserInit() {
	staticData := &varianceParserStaticData
	staticData.once.Do(varianceParserInit)
}

// NewVarianceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVarianceParser(input antlr.TokenStream) *VarianceParser {
	VarianceParserInit()
	this := new(VarianceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &varianceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Variance.g4"

	return this
}

// VarianceParser tokens.
const (
	VarianceParserEOF      = antlr.TokenEOF
	VarianceParserVARIANCE = 1
	VarianceParserOWNKEY   = 2
	VarianceParserSPLITKEY = 3
	VarianceParserWS       = 4
)

// VarianceParser rules.
const (
	VarianceParserRULE_expression = 0
	VarianceParserRULE_variance   = 1
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
	p.RuleIndex = VarianceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VarianceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Variance() IVarianceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarianceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarianceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VarianceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VarianceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VarianceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VarianceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VarianceParserRULE_expression)

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
		p.Variance()
	}
	{
		p.SetState(5)
		p.Match(VarianceParserEOF)
	}

	return localctx
}

// IVarianceContext is an interface to support dynamic dispatch.
type IVarianceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarianceContext differentiates from other interfaces.
	IsVarianceContext()
}

type VarianceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarianceContext() *VarianceContext {
	var p = new(VarianceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VarianceParserRULE_variance
	return p
}

func (*VarianceContext) IsVarianceContext() {}

func NewVarianceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarianceContext {
	var p = new(VarianceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VarianceParserRULE_variance

	return p
}

func (s *VarianceContext) GetParser() antlr.Parser { return s.parser }

func (s *VarianceContext) VARIANCE() antlr.TerminalNode {
	return s.GetToken(VarianceParserVARIANCE, 0)
}

func (s *VarianceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarianceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarianceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VarianceListener); ok {
		listenerT.EnterVariance(s)
	}
}

func (s *VarianceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VarianceListener); ok {
		listenerT.ExitVariance(s)
	}
}

func (p *VarianceParser) Variance() (localctx IVarianceContext) {
	this := p
	_ = this

	localctx = NewVarianceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VarianceParserRULE_variance)

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
		p.Match(VarianceParserVARIANCE)
	}

	return localctx
}
