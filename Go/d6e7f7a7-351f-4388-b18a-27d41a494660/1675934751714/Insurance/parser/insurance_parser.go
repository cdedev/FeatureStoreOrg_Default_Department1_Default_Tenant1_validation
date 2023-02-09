// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675934751714/Insurance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Insurance

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

type InsuranceParser struct {
	*antlr.BaseParser
}

var insuranceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func insuranceParserInit() {
	staticData := &insuranceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INSURANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "insurance",
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

// InsuranceParserInit initializes any static state used to implement InsuranceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInsuranceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InsuranceParserInit() {
	staticData := &insuranceParserStaticData
	staticData.once.Do(insuranceParserInit)
}

// NewInsuranceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInsuranceParser(input antlr.TokenStream) *InsuranceParser {
	InsuranceParserInit()
	this := new(InsuranceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &insuranceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Insurance.g4"

	return this
}

// InsuranceParser tokens.
const (
	InsuranceParserEOF       = antlr.TokenEOF
	InsuranceParserINSURANCE = 1
	InsuranceParserOWNKEY    = 2
	InsuranceParserSPLITKEY  = 3
	InsuranceParserWS        = 4
)

// InsuranceParser rules.
const (
	InsuranceParserRULE_expression = 0
	InsuranceParserRULE_insurance  = 1
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
	p.RuleIndex = InsuranceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InsuranceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Insurance() IInsuranceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsuranceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsuranceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(InsuranceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InsuranceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InsuranceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *InsuranceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InsuranceParserRULE_expression)

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
		p.Insurance()
	}
	{
		p.SetState(5)
		p.Match(InsuranceParserEOF)
	}

	return localctx
}

// IInsuranceContext is an interface to support dynamic dispatch.
type IInsuranceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsuranceContext differentiates from other interfaces.
	IsInsuranceContext()
}

type InsuranceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsuranceContext() *InsuranceContext {
	var p = new(InsuranceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InsuranceParserRULE_insurance
	return p
}

func (*InsuranceContext) IsInsuranceContext() {}

func NewInsuranceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsuranceContext {
	var p = new(InsuranceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InsuranceParserRULE_insurance

	return p
}

func (s *InsuranceContext) GetParser() antlr.Parser { return s.parser }

func (s *InsuranceContext) INSURANCE() antlr.TerminalNode {
	return s.GetToken(InsuranceParserINSURANCE, 0)
}

func (s *InsuranceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsuranceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsuranceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InsuranceListener); ok {
		listenerT.EnterInsurance(s)
	}
}

func (s *InsuranceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InsuranceListener); ok {
		listenerT.ExitInsurance(s)
	}
}

func (p *InsuranceParser) Insurance() (localctx IInsuranceContext) {
	this := p
	_ = this

	localctx = NewInsuranceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InsuranceParserRULE_insurance)

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
		p.Match(InsuranceParserINSURANCE)
	}

	return localctx
}
