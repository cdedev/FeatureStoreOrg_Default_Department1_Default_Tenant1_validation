// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690427235/MonthAsCustomer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MonthAsCustomer

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

type MonthAsCustomerParser struct {
	*antlr.BaseParser
}

var monthascustomerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func monthascustomerParserInit() {
	staticData := &monthascustomerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MONTHASCUSTOMER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "monthascustomer",
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

// MonthAsCustomerParserInit initializes any static state used to implement MonthAsCustomerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMonthAsCustomerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MonthAsCustomerParserInit() {
	staticData := &monthascustomerParserStaticData
	staticData.once.Do(monthascustomerParserInit)
}

// NewMonthAsCustomerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMonthAsCustomerParser(input antlr.TokenStream) *MonthAsCustomerParser {
	MonthAsCustomerParserInit()
	this := new(MonthAsCustomerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &monthascustomerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MonthAsCustomer.g4"

	return this
}

// MonthAsCustomerParser tokens.
const (
	MonthAsCustomerParserEOF             = antlr.TokenEOF
	MonthAsCustomerParserMONTHASCUSTOMER = 1
	MonthAsCustomerParserOWNKEY          = 2
	MonthAsCustomerParserSPLITKEY        = 3
	MonthAsCustomerParserWS              = 4
)

// MonthAsCustomerParser rules.
const (
	MonthAsCustomerParserRULE_expression      = 0
	MonthAsCustomerParserRULE_monthascustomer = 1
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
	p.RuleIndex = MonthAsCustomerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonthAsCustomerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Monthascustomer() IMonthascustomerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMonthascustomerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMonthascustomerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MonthAsCustomerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthAsCustomerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthAsCustomerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MonthAsCustomerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MonthAsCustomerParserRULE_expression)

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
		p.Monthascustomer()
	}
	{
		p.SetState(5)
		p.Match(MonthAsCustomerParserEOF)
	}

	return localctx
}

// IMonthascustomerContext is an interface to support dynamic dispatch.
type IMonthascustomerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthascustomerContext differentiates from other interfaces.
	IsMonthascustomerContext()
}

type MonthascustomerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthascustomerContext() *MonthascustomerContext {
	var p = new(MonthascustomerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MonthAsCustomerParserRULE_monthascustomer
	return p
}

func (*MonthascustomerContext) IsMonthascustomerContext() {}

func NewMonthascustomerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthascustomerContext {
	var p = new(MonthascustomerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MonthAsCustomerParserRULE_monthascustomer

	return p
}

func (s *MonthascustomerContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthascustomerContext) MONTHASCUSTOMER() antlr.TerminalNode {
	return s.GetToken(MonthAsCustomerParserMONTHASCUSTOMER, 0)
}

func (s *MonthascustomerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthascustomerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthascustomerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthAsCustomerListener); ok {
		listenerT.EnterMonthascustomer(s)
	}
}

func (s *MonthascustomerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MonthAsCustomerListener); ok {
		listenerT.ExitMonthascustomer(s)
	}
}

func (p *MonthAsCustomerParser) Monthascustomer() (localctx IMonthascustomerContext) {
	this := p
	_ = this

	localctx = NewMonthascustomerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MonthAsCustomerParserRULE_monthascustomer)

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
		p.Match(MonthAsCustomerParserMONTHASCUSTOMER)
	}

	return localctx
}
