// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671207762977/ConfirmedCases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ConfirmedCases

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

type ConfirmedCasesParser struct {
	*antlr.BaseParser
}

var confirmedcasesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func confirmedcasesParserInit() {
	staticData := &confirmedcasesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONFIRMEDCASES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "confirmedcases",
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

// ConfirmedCasesParserInit initializes any static state used to implement ConfirmedCasesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConfirmedCasesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfirmedCasesParserInit() {
	staticData := &confirmedcasesParserStaticData
	staticData.once.Do(confirmedcasesParserInit)
}

// NewConfirmedCasesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConfirmedCasesParser(input antlr.TokenStream) *ConfirmedCasesParser {
	ConfirmedCasesParserInit()
	this := new(ConfirmedCasesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &confirmedcasesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ConfirmedCases.g4"

	return this
}

// ConfirmedCasesParser tokens.
const (
	ConfirmedCasesParserEOF            = antlr.TokenEOF
	ConfirmedCasesParserCONFIRMEDCASES = 1
	ConfirmedCasesParserOWNKEY         = 2
	ConfirmedCasesParserSPLITKEY       = 3
	ConfirmedCasesParserWS             = 4
)

// ConfirmedCasesParser rules.
const (
	ConfirmedCasesParserRULE_expression     = 0
	ConfirmedCasesParserRULE_confirmedcases = 1
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
	p.RuleIndex = ConfirmedCasesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfirmedCasesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Confirmedcases() IConfirmedcasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfirmedcasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConfirmedcasesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfirmedCasesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfirmedCasesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfirmedCasesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConfirmedCasesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfirmedCasesParserRULE_expression)

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
		p.Confirmedcases()
	}
	{
		p.SetState(5)
		p.Match(ConfirmedCasesParserEOF)
	}

	return localctx
}

// IConfirmedcasesContext is an interface to support dynamic dispatch.
type IConfirmedcasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfirmedcasesContext differentiates from other interfaces.
	IsConfirmedcasesContext()
}

type ConfirmedcasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfirmedcasesContext() *ConfirmedcasesContext {
	var p = new(ConfirmedcasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfirmedCasesParserRULE_confirmedcases
	return p
}

func (*ConfirmedcasesContext) IsConfirmedcasesContext() {}

func NewConfirmedcasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfirmedcasesContext {
	var p = new(ConfirmedcasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfirmedCasesParserRULE_confirmedcases

	return p
}

func (s *ConfirmedcasesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfirmedcasesContext) CONFIRMEDCASES() antlr.TerminalNode {
	return s.GetToken(ConfirmedCasesParserCONFIRMEDCASES, 0)
}

func (s *ConfirmedcasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfirmedcasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfirmedcasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfirmedCasesListener); ok {
		listenerT.EnterConfirmedcases(s)
	}
}

func (s *ConfirmedcasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfirmedCasesListener); ok {
		listenerT.ExitConfirmedcases(s)
	}
}

func (p *ConfirmedCasesParser) Confirmedcases() (localctx IConfirmedcasesContext) {
	this := p
	_ = this

	localctx = NewConfirmedcasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfirmedCasesParserRULE_confirmedcases)

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
		p.Match(ConfirmedCasesParserCONFIRMEDCASES)
	}

	return localctx
}
