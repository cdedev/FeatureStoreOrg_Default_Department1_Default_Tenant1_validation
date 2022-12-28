// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672201533429/Salary.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Salary

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

type SalaryParser struct {
	*antlr.BaseParser
}

var salaryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func salaryParserInit() {
	staticData := &salaryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SALARY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "salary",
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

// SalaryParserInit initializes any static state used to implement SalaryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSalaryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SalaryParserInit() {
	staticData := &salaryParserStaticData
	staticData.once.Do(salaryParserInit)
}

// NewSalaryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSalaryParser(input antlr.TokenStream) *SalaryParser {
	SalaryParserInit()
	this := new(SalaryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &salaryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Salary.g4"

	return this
}

// SalaryParser tokens.
const (
	SalaryParserEOF      = antlr.TokenEOF
	SalaryParserSALARY   = 1
	SalaryParserOWNKEY   = 2
	SalaryParserSPLITKEY = 3
	SalaryParserWS       = 4
)

// SalaryParser rules.
const (
	SalaryParserRULE_expression = 0
	SalaryParserRULE_salary     = 1
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
	p.RuleIndex = SalaryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SalaryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Salary() ISalaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISalaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISalaryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SalaryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalaryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalaryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SalaryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SalaryParserRULE_expression)

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
		p.Salary()
	}
	{
		p.SetState(5)
		p.Match(SalaryParserEOF)
	}

	return localctx
}

// ISalaryContext is an interface to support dynamic dispatch.
type ISalaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalaryContext differentiates from other interfaces.
	IsSalaryContext()
}

type SalaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalaryContext() *SalaryContext {
	var p = new(SalaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SalaryParserRULE_salary
	return p
}

func (*SalaryContext) IsSalaryContext() {}

func NewSalaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalaryContext {
	var p = new(SalaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SalaryParserRULE_salary

	return p
}

func (s *SalaryContext) GetParser() antlr.Parser { return s.parser }

func (s *SalaryContext) SALARY() antlr.TerminalNode {
	return s.GetToken(SalaryParserSALARY, 0)
}

func (s *SalaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalaryListener); ok {
		listenerT.EnterSalary(s)
	}
}

func (s *SalaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SalaryListener); ok {
		listenerT.ExitSalary(s)
	}
}

func (p *SalaryParser) Salary() (localctx ISalaryContext) {
	this := p
	_ = this

	localctx = NewSalaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SalaryParserRULE_salary)

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
		p.Match(SalaryParserSALARY)
	}

	return localctx
}
