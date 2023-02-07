// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675739456793/Agreement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agreement

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

type AgreementParser struct {
	*antlr.BaseParser
}

var agreementParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agreementParserInit() {
	staticData := &agreementParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AGREEMENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "agreement",
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

// AgreementParserInit initializes any static state used to implement AgreementParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAgreementParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgreementParserInit() {
	staticData := &agreementParserStaticData
	staticData.once.Do(agreementParserInit)
}

// NewAgreementParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAgreementParser(input antlr.TokenStream) *AgreementParser {
	AgreementParserInit()
	this := new(AgreementParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &agreementParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Agreement.g4"

	return this
}

// AgreementParser tokens.
const (
	AgreementParserEOF       = antlr.TokenEOF
	AgreementParserAGREEMENT = 1
	AgreementParserOWNKEY    = 2
	AgreementParserSPLITKEY  = 3
	AgreementParserWS        = 4
)

// AgreementParser rules.
const (
	AgreementParserRULE_expression = 0
	AgreementParserRULE_agreement  = 1
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
	p.RuleIndex = AgreementParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgreementParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Agreement() IAgreementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAgreementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAgreementContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AgreementParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreementListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreementListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AgreementParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgreementParserRULE_expression)

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
		p.Agreement()
	}
	{
		p.SetState(5)
		p.Match(AgreementParserEOF)
	}

	return localctx
}

// IAgreementContext is an interface to support dynamic dispatch.
type IAgreementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgreementContext differentiates from other interfaces.
	IsAgreementContext()
}

type AgreementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgreementContext() *AgreementContext {
	var p = new(AgreementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgreementParserRULE_agreement
	return p
}

func (*AgreementContext) IsAgreementContext() {}

func NewAgreementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgreementContext {
	var p = new(AgreementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgreementParserRULE_agreement

	return p
}

func (s *AgreementContext) GetParser() antlr.Parser { return s.parser }

func (s *AgreementContext) AGREEMENT() antlr.TerminalNode {
	return s.GetToken(AgreementParserAGREEMENT, 0)
}

func (s *AgreementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgreementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgreementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreementListener); ok {
		listenerT.EnterAgreement(s)
	}
}

func (s *AgreementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreementListener); ok {
		listenerT.ExitAgreement(s)
	}
}

func (p *AgreementParser) Agreement() (localctx IAgreementContext) {
	this := p
	_ = this

	localctx = NewAgreementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgreementParserRULE_agreement)

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
		p.Match(AgreementParserAGREEMENT)
	}

	return localctx
}
