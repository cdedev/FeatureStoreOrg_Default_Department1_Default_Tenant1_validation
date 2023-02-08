// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834537820/Employment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Employment

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

type EmploymentParser struct {
	*antlr.BaseParser
}

var employmentParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func employmentParserInit() {
	staticData := &employmentParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EMPLOYMENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "employment",
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

// EmploymentParserInit initializes any static state used to implement EmploymentParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEmploymentParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EmploymentParserInit() {
	staticData := &employmentParserStaticData
	staticData.once.Do(employmentParserInit)
}

// NewEmploymentParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEmploymentParser(input antlr.TokenStream) *EmploymentParser {
	EmploymentParserInit()
	this := new(EmploymentParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &employmentParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Employment.g4"

	return this
}

// EmploymentParser tokens.
const (
	EmploymentParserEOF        = antlr.TokenEOF
	EmploymentParserEMPLOYMENT = 1
	EmploymentParserOWNKEY     = 2
	EmploymentParserSPLITKEY   = 3
	EmploymentParserWS         = 4
)

// EmploymentParser rules.
const (
	EmploymentParserRULE_expression = 0
	EmploymentParserRULE_employment = 1
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
	p.RuleIndex = EmploymentParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmploymentParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Employment() IEmploymentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmploymentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmploymentContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EmploymentParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EmploymentParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EmploymentParserRULE_expression)

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
		p.Employment()
	}
	{
		p.SetState(5)
		p.Match(EmploymentParserEOF)
	}

	return localctx
}

// IEmploymentContext is an interface to support dynamic dispatch.
type IEmploymentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmploymentContext differentiates from other interfaces.
	IsEmploymentContext()
}

type EmploymentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmploymentContext() *EmploymentContext {
	var p = new(EmploymentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EmploymentParserRULE_employment
	return p
}

func (*EmploymentContext) IsEmploymentContext() {}

func NewEmploymentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmploymentContext {
	var p = new(EmploymentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmploymentParserRULE_employment

	return p
}

func (s *EmploymentContext) GetParser() antlr.Parser { return s.parser }

func (s *EmploymentContext) EMPLOYMENT() antlr.TerminalNode {
	return s.GetToken(EmploymentParserEMPLOYMENT, 0)
}

func (s *EmploymentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmploymentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmploymentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentListener); ok {
		listenerT.EnterEmployment(s)
	}
}

func (s *EmploymentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentListener); ok {
		listenerT.ExitEmployment(s)
	}
}

func (p *EmploymentParser) Employment() (localctx IEmploymentContext) {
	this := p
	_ = this

	localctx = NewEmploymentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EmploymentParserRULE_employment)

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
		p.Match(EmploymentParserEMPLOYMENT)
	}

	return localctx
}
