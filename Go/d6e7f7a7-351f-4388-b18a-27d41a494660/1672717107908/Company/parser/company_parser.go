// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672717107908/Company.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Company

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

type CompanyParser struct {
	*antlr.BaseParser
}

var companyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func companyParserInit() {
	staticData := &companyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COMPANY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "company",
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

// CompanyParserInit initializes any static state used to implement CompanyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCompanyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CompanyParserInit() {
	staticData := &companyParserStaticData
	staticData.once.Do(companyParserInit)
}

// NewCompanyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCompanyParser(input antlr.TokenStream) *CompanyParser {
	CompanyParserInit()
	this := new(CompanyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &companyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Company.g4"

	return this
}

// CompanyParser tokens.
const (
	CompanyParserEOF      = antlr.TokenEOF
	CompanyParserCOMPANY  = 1
	CompanyParserOWNKEY   = 2
	CompanyParserSPLITKEY = 3
	CompanyParserWS       = 4
)

// CompanyParser rules.
const (
	CompanyParserRULE_expression = 0
	CompanyParserRULE_company    = 1
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
	p.RuleIndex = CompanyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompanyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Company() ICompanyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompanyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompanyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CompanyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CompanyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CompanyParserRULE_expression)

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
		p.Company()
	}
	{
		p.SetState(5)
		p.Match(CompanyParserEOF)
	}

	return localctx
}

// ICompanyContext is an interface to support dynamic dispatch.
type ICompanyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompanyContext differentiates from other interfaces.
	IsCompanyContext()
}

type CompanyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompanyContext() *CompanyContext {
	var p = new(CompanyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CompanyParserRULE_company
	return p
}

func (*CompanyContext) IsCompanyContext() {}

func NewCompanyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompanyContext {
	var p = new(CompanyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompanyParserRULE_company

	return p
}

func (s *CompanyContext) GetParser() antlr.Parser { return s.parser }

func (s *CompanyContext) COMPANY() antlr.TerminalNode {
	return s.GetToken(CompanyParserCOMPANY, 0)
}

func (s *CompanyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompanyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompanyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyListener); ok {
		listenerT.EnterCompany(s)
	}
}

func (s *CompanyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyListener); ok {
		listenerT.ExitCompany(s)
	}
}

func (p *CompanyParser) Company() (localctx ICompanyContext) {
	this := p
	_ = this

	localctx = NewCompanyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CompanyParserRULE_company)

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
		p.Match(CompanyParserCOMPANY)
	}

	return localctx
}
