// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659495314/Department.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Department

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

type DepartmentParser struct {
	*antlr.BaseParser
}

var departmentParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func departmentParserInit() {
	staticData := &departmentParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DEPARTMENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "department",
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

// DepartmentParserInit initializes any static state used to implement DepartmentParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDepartmentParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DepartmentParserInit() {
	staticData := &departmentParserStaticData
	staticData.once.Do(departmentParserInit)
}

// NewDepartmentParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDepartmentParser(input antlr.TokenStream) *DepartmentParser {
	DepartmentParserInit()
	this := new(DepartmentParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &departmentParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Department.g4"

	return this
}

// DepartmentParser tokens.
const (
	DepartmentParserEOF        = antlr.TokenEOF
	DepartmentParserDEPARTMENT = 1
	DepartmentParserOWNKEY     = 2
	DepartmentParserSPLITKEY   = 3
	DepartmentParserWS         = 4
)

// DepartmentParser rules.
const (
	DepartmentParserRULE_expression = 0
	DepartmentParserRULE_department = 1
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
	p.RuleIndex = DepartmentParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DepartmentParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Department() IDepartmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDepartmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDepartmentContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DepartmentParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepartmentListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepartmentListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DepartmentParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DepartmentParserRULE_expression)

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
		p.Department()
	}
	{
		p.SetState(5)
		p.Match(DepartmentParserEOF)
	}

	return localctx
}

// IDepartmentContext is an interface to support dynamic dispatch.
type IDepartmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDepartmentContext differentiates from other interfaces.
	IsDepartmentContext()
}

type DepartmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDepartmentContext() *DepartmentContext {
	var p = new(DepartmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DepartmentParserRULE_department
	return p
}

func (*DepartmentContext) IsDepartmentContext() {}

func NewDepartmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DepartmentContext {
	var p = new(DepartmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DepartmentParserRULE_department

	return p
}

func (s *DepartmentContext) GetParser() antlr.Parser { return s.parser }

func (s *DepartmentContext) DEPARTMENT() antlr.TerminalNode {
	return s.GetToken(DepartmentParserDEPARTMENT, 0)
}

func (s *DepartmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DepartmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DepartmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepartmentListener); ok {
		listenerT.EnterDepartment(s)
	}
}

func (s *DepartmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepartmentListener); ok {
		listenerT.ExitDepartment(s)
	}
}

func (p *DepartmentParser) Department() (localctx IDepartmentContext) {
	this := p
	_ = this

	localctx = NewDepartmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DepartmentParserRULE_department)

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
		p.Match(DepartmentParserDEPARTMENT)
	}

	return localctx
}
