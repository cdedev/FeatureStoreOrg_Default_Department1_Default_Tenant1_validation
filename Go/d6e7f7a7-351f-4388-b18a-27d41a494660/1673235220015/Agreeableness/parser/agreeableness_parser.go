// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235220015/Agreeableness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agreeableness

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

type AgreeablenessParser struct {
	*antlr.BaseParser
}

var agreeablenessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agreeablenessParserInit() {
	staticData := &agreeablenessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AGREEABLENESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "agreeableness",
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

// AgreeablenessParserInit initializes any static state used to implement AgreeablenessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAgreeablenessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgreeablenessParserInit() {
	staticData := &agreeablenessParserStaticData
	staticData.once.Do(agreeablenessParserInit)
}

// NewAgreeablenessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAgreeablenessParser(input antlr.TokenStream) *AgreeablenessParser {
	AgreeablenessParserInit()
	this := new(AgreeablenessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &agreeablenessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Agreeableness.g4"

	return this
}

// AgreeablenessParser tokens.
const (
	AgreeablenessParserEOF           = antlr.TokenEOF
	AgreeablenessParserAGREEABLENESS = 1
	AgreeablenessParserOWNKEY        = 2
	AgreeablenessParserSPLITKEY      = 3
	AgreeablenessParserWS            = 4
)

// AgreeablenessParser rules.
const (
	AgreeablenessParserRULE_expression    = 0
	AgreeablenessParserRULE_agreeableness = 1
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
	p.RuleIndex = AgreeablenessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgreeablenessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Agreeableness() IAgreeablenessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAgreeablenessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAgreeablenessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AgreeablenessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreeablenessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreeablenessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AgreeablenessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgreeablenessParserRULE_expression)

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
		p.Agreeableness()
	}
	{
		p.SetState(5)
		p.Match(AgreeablenessParserEOF)
	}

	return localctx
}

// IAgreeablenessContext is an interface to support dynamic dispatch.
type IAgreeablenessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgreeablenessContext differentiates from other interfaces.
	IsAgreeablenessContext()
}

type AgreeablenessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgreeablenessContext() *AgreeablenessContext {
	var p = new(AgreeablenessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgreeablenessParserRULE_agreeableness
	return p
}

func (*AgreeablenessContext) IsAgreeablenessContext() {}

func NewAgreeablenessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgreeablenessContext {
	var p = new(AgreeablenessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgreeablenessParserRULE_agreeableness

	return p
}

func (s *AgreeablenessContext) GetParser() antlr.Parser { return s.parser }

func (s *AgreeablenessContext) AGREEABLENESS() antlr.TerminalNode {
	return s.GetToken(AgreeablenessParserAGREEABLENESS, 0)
}

func (s *AgreeablenessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgreeablenessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgreeablenessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreeablenessListener); ok {
		listenerT.EnterAgreeableness(s)
	}
}

func (s *AgreeablenessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgreeablenessListener); ok {
		listenerT.ExitAgreeableness(s)
	}
}

func (p *AgreeablenessParser) Agreeableness() (localctx IAgreeablenessContext) {
	this := p
	_ = this

	localctx = NewAgreeablenessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgreeablenessParserRULE_agreeableness)

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
		p.Match(AgreeablenessParserAGREEABLENESS)
	}

	return localctx
}
