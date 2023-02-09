// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675932748157/CompanyType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CompanyType

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

type CompanyTypeParser struct {
	*antlr.BaseParser
}

var companytypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func companytypeParserInit() {
	staticData := &companytypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COMPANYTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "companytype",
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

// CompanyTypeParserInit initializes any static state used to implement CompanyTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCompanyTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CompanyTypeParserInit() {
	staticData := &companytypeParserStaticData
	staticData.once.Do(companytypeParserInit)
}

// NewCompanyTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCompanyTypeParser(input antlr.TokenStream) *CompanyTypeParser {
	CompanyTypeParserInit()
	this := new(CompanyTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &companytypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CompanyType.g4"

	return this
}

// CompanyTypeParser tokens.
const (
	CompanyTypeParserEOF         = antlr.TokenEOF
	CompanyTypeParserCOMPANYTYPE = 1
	CompanyTypeParserOWNKEY      = 2
	CompanyTypeParserSPLITKEY    = 3
	CompanyTypeParserWS          = 4
)

// CompanyTypeParser rules.
const (
	CompanyTypeParserRULE_expression  = 0
	CompanyTypeParserRULE_companytype = 1
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
	p.RuleIndex = CompanyTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompanyTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Companytype() ICompanytypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompanytypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompanytypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CompanyTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CompanyTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CompanyTypeParserRULE_expression)

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
		p.Companytype()
	}
	{
		p.SetState(5)
		p.Match(CompanyTypeParserEOF)
	}

	return localctx
}

// ICompanytypeContext is an interface to support dynamic dispatch.
type ICompanytypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompanytypeContext differentiates from other interfaces.
	IsCompanytypeContext()
}

type CompanytypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompanytypeContext() *CompanytypeContext {
	var p = new(CompanytypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CompanyTypeParserRULE_companytype
	return p
}

func (*CompanytypeContext) IsCompanytypeContext() {}

func NewCompanytypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompanytypeContext {
	var p = new(CompanytypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompanyTypeParserRULE_companytype

	return p
}

func (s *CompanytypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CompanytypeContext) COMPANYTYPE() antlr.TerminalNode {
	return s.GetToken(CompanyTypeParserCOMPANYTYPE, 0)
}

func (s *CompanytypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompanytypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompanytypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyTypeListener); ok {
		listenerT.EnterCompanytype(s)
	}
}

func (s *CompanytypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyTypeListener); ok {
		listenerT.ExitCompanytype(s)
	}
}

func (p *CompanyTypeParser) Companytype() (localctx ICompanytypeContext) {
	this := p
	_ = this

	localctx = NewCompanytypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CompanyTypeParserRULE_companytype)

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
		p.Match(CompanyTypeParserCOMPANYTYPE)
	}

	return localctx
}
