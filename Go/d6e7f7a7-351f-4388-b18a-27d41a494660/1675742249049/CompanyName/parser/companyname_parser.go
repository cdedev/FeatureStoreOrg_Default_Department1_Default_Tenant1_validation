// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742249049/CompanyName.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CompanyName

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

type CompanyNameParser struct {
	*antlr.BaseParser
}

var companynameParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func companynameParserInit() {
	staticData := &companynameParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COMPANYNAME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "companyname",
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

// CompanyNameParserInit initializes any static state used to implement CompanyNameParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCompanyNameParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CompanyNameParserInit() {
	staticData := &companynameParserStaticData
	staticData.once.Do(companynameParserInit)
}

// NewCompanyNameParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCompanyNameParser(input antlr.TokenStream) *CompanyNameParser {
	CompanyNameParserInit()
	this := new(CompanyNameParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &companynameParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CompanyName.g4"

	return this
}

// CompanyNameParser tokens.
const (
	CompanyNameParserEOF         = antlr.TokenEOF
	CompanyNameParserCOMPANYNAME = 1
	CompanyNameParserOWNKEY      = 2
	CompanyNameParserSPLITKEY    = 3
	CompanyNameParserWS          = 4
)

// CompanyNameParser rules.
const (
	CompanyNameParserRULE_expression  = 0
	CompanyNameParserRULE_companyname = 1
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
	p.RuleIndex = CompanyNameParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompanyNameParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Companyname() ICompanynameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompanynameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompanynameContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CompanyNameParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyNameListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyNameListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CompanyNameParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CompanyNameParserRULE_expression)

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
		p.Companyname()
	}
	{
		p.SetState(5)
		p.Match(CompanyNameParserEOF)
	}

	return localctx
}

// ICompanynameContext is an interface to support dynamic dispatch.
type ICompanynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompanynameContext differentiates from other interfaces.
	IsCompanynameContext()
}

type CompanynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompanynameContext() *CompanynameContext {
	var p = new(CompanynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CompanyNameParserRULE_companyname
	return p
}

func (*CompanynameContext) IsCompanynameContext() {}

func NewCompanynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompanynameContext {
	var p = new(CompanynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CompanyNameParserRULE_companyname

	return p
}

func (s *CompanynameContext) GetParser() antlr.Parser { return s.parser }

func (s *CompanynameContext) COMPANYNAME() antlr.TerminalNode {
	return s.GetToken(CompanyNameParserCOMPANYNAME, 0)
}

func (s *CompanynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompanynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompanynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyNameListener); ok {
		listenerT.EnterCompanyname(s)
	}
}

func (s *CompanynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CompanyNameListener); ok {
		listenerT.ExitCompanyname(s)
	}
}

func (p *CompanyNameParser) Companyname() (localctx ICompanynameContext) {
	this := p
	_ = this

	localctx = NewCompanynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CompanyNameParserRULE_companyname)

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
		p.Match(CompanyNameParserCOMPANYNAME)
	}

	return localctx
}
