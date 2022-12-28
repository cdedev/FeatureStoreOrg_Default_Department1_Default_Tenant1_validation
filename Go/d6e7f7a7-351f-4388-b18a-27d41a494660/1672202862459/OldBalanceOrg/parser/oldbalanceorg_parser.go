// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202862459/OldBalanceOrg.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OldBalanceOrg

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

type OldBalanceOrgParser struct {
	*antlr.BaseParser
}

var oldbalanceorgParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func oldbalanceorgParserInit() {
	staticData := &oldbalanceorgParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OLDBALANCEORG", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "oldbalanceorg",
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

// OldBalanceOrgParserInit initializes any static state used to implement OldBalanceOrgParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOldBalanceOrgParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OldBalanceOrgParserInit() {
	staticData := &oldbalanceorgParserStaticData
	staticData.once.Do(oldbalanceorgParserInit)
}

// NewOldBalanceOrgParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOldBalanceOrgParser(input antlr.TokenStream) *OldBalanceOrgParser {
	OldBalanceOrgParserInit()
	this := new(OldBalanceOrgParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &oldbalanceorgParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "OldBalanceOrg.g4"

	return this
}

// OldBalanceOrgParser tokens.
const (
	OldBalanceOrgParserEOF           = antlr.TokenEOF
	OldBalanceOrgParserOLDBALANCEORG = 1
	OldBalanceOrgParserOWNKEY        = 2
	OldBalanceOrgParserSPLITKEY      = 3
	OldBalanceOrgParserWS            = 4
)

// OldBalanceOrgParser rules.
const (
	OldBalanceOrgParserRULE_expression    = 0
	OldBalanceOrgParserRULE_oldbalanceorg = 1
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
	p.RuleIndex = OldBalanceOrgParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OldBalanceOrgParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Oldbalanceorg() IOldbalanceorgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOldbalanceorgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOldbalanceorgContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OldBalanceOrgParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OldBalanceOrgListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OldBalanceOrgListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OldBalanceOrgParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OldBalanceOrgParserRULE_expression)

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
		p.Oldbalanceorg()
	}
	{
		p.SetState(5)
		p.Match(OldBalanceOrgParserEOF)
	}

	return localctx
}

// IOldbalanceorgContext is an interface to support dynamic dispatch.
type IOldbalanceorgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOldbalanceorgContext differentiates from other interfaces.
	IsOldbalanceorgContext()
}

type OldbalanceorgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOldbalanceorgContext() *OldbalanceorgContext {
	var p = new(OldbalanceorgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OldBalanceOrgParserRULE_oldbalanceorg
	return p
}

func (*OldbalanceorgContext) IsOldbalanceorgContext() {}

func NewOldbalanceorgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OldbalanceorgContext {
	var p = new(OldbalanceorgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OldBalanceOrgParserRULE_oldbalanceorg

	return p
}

func (s *OldbalanceorgContext) GetParser() antlr.Parser { return s.parser }

func (s *OldbalanceorgContext) OLDBALANCEORG() antlr.TerminalNode {
	return s.GetToken(OldBalanceOrgParserOLDBALANCEORG, 0)
}

func (s *OldbalanceorgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OldbalanceorgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OldbalanceorgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OldBalanceOrgListener); ok {
		listenerT.EnterOldbalanceorg(s)
	}
}

func (s *OldbalanceorgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OldBalanceOrgListener); ok {
		listenerT.ExitOldbalanceorg(s)
	}
}

func (p *OldBalanceOrgParser) Oldbalanceorg() (localctx IOldbalanceorgContext) {
	this := p
	_ = this

	localctx = NewOldbalanceorgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OldBalanceOrgParserRULE_oldbalanceorg)

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
		p.Match(OldBalanceOrgParserOLDBALANCEORG)
	}

	return localctx
}
