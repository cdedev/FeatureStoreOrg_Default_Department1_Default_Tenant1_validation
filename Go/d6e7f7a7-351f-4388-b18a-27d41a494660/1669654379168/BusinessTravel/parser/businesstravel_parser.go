// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669654379168/BusinessTravel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BusinessTravel

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

type BusinessTravelParser struct {
	*antlr.BaseParser
}

var businesstravelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func businesstravelParserInit() {
	staticData := &businesstravelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BUSINESSTRAVEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "businesstravel",
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

// BusinessTravelParserInit initializes any static state used to implement BusinessTravelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBusinessTravelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BusinessTravelParserInit() {
	staticData := &businesstravelParserStaticData
	staticData.once.Do(businesstravelParserInit)
}

// NewBusinessTravelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBusinessTravelParser(input antlr.TokenStream) *BusinessTravelParser {
	BusinessTravelParserInit()
	this := new(BusinessTravelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &businesstravelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BusinessTravel.g4"

	return this
}

// BusinessTravelParser tokens.
const (
	BusinessTravelParserEOF            = antlr.TokenEOF
	BusinessTravelParserBUSINESSTRAVEL = 1
	BusinessTravelParserOWNKEY         = 2
	BusinessTravelParserSPLITKEY       = 3
	BusinessTravelParserWS             = 4
)

// BusinessTravelParser rules.
const (
	BusinessTravelParserRULE_expression     = 0
	BusinessTravelParserRULE_businesstravel = 1
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
	p.RuleIndex = BusinessTravelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BusinessTravelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Businesstravel() IBusinesstravelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBusinesstravelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBusinesstravelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BusinessTravelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BusinessTravelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BusinessTravelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BusinessTravelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BusinessTravelParserRULE_expression)

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
		p.Businesstravel()
	}
	{
		p.SetState(5)
		p.Match(BusinessTravelParserEOF)
	}

	return localctx
}

// IBusinesstravelContext is an interface to support dynamic dispatch.
type IBusinesstravelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBusinesstravelContext differentiates from other interfaces.
	IsBusinesstravelContext()
}

type BusinesstravelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBusinesstravelContext() *BusinesstravelContext {
	var p = new(BusinesstravelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BusinessTravelParserRULE_businesstravel
	return p
}

func (*BusinesstravelContext) IsBusinesstravelContext() {}

func NewBusinesstravelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BusinesstravelContext {
	var p = new(BusinesstravelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BusinessTravelParserRULE_businesstravel

	return p
}

func (s *BusinesstravelContext) GetParser() antlr.Parser { return s.parser }

func (s *BusinesstravelContext) BUSINESSTRAVEL() antlr.TerminalNode {
	return s.GetToken(BusinessTravelParserBUSINESSTRAVEL, 0)
}

func (s *BusinesstravelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BusinesstravelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BusinesstravelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BusinessTravelListener); ok {
		listenerT.EnterBusinesstravel(s)
	}
}

func (s *BusinesstravelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BusinessTravelListener); ok {
		listenerT.ExitBusinesstravel(s)
	}
}

func (p *BusinessTravelParser) Businesstravel() (localctx IBusinesstravelContext) {
	this := p
	_ = this

	localctx = NewBusinesstravelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BusinessTravelParserRULE_businesstravel)

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
		p.Match(BusinessTravelParserBUSINESSTRAVEL)
	}

	return localctx
}
