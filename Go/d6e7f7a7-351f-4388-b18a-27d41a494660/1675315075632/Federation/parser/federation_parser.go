// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675315075632/Federation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Federation

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

type FederationParser struct {
	*antlr.BaseParser
}

var federationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func federationParserInit() {
	staticData := &federationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FEDERATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "federation",
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

// FederationParserInit initializes any static state used to implement FederationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFederationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FederationParserInit() {
	staticData := &federationParserStaticData
	staticData.once.Do(federationParserInit)
}

// NewFederationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFederationParser(input antlr.TokenStream) *FederationParser {
	FederationParserInit()
	this := new(FederationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &federationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Federation.g4"

	return this
}

// FederationParser tokens.
const (
	FederationParserEOF        = antlr.TokenEOF
	FederationParserFEDERATION = 1
	FederationParserOWNKEY     = 2
	FederationParserSPLITKEY   = 3
	FederationParserWS         = 4
)

// FederationParser rules.
const (
	FederationParserRULE_expression = 0
	FederationParserRULE_federation = 1
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
	p.RuleIndex = FederationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FederationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Federation() IFederationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFederationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFederationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FederationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FederationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FederationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FederationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FederationParserRULE_expression)

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
		p.Federation()
	}
	{
		p.SetState(5)
		p.Match(FederationParserEOF)
	}

	return localctx
}

// IFederationContext is an interface to support dynamic dispatch.
type IFederationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFederationContext differentiates from other interfaces.
	IsFederationContext()
}

type FederationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFederationContext() *FederationContext {
	var p = new(FederationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FederationParserRULE_federation
	return p
}

func (*FederationContext) IsFederationContext() {}

func NewFederationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FederationContext {
	var p = new(FederationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FederationParserRULE_federation

	return p
}

func (s *FederationContext) GetParser() antlr.Parser { return s.parser }

func (s *FederationContext) FEDERATION() antlr.TerminalNode {
	return s.GetToken(FederationParserFEDERATION, 0)
}

func (s *FederationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FederationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FederationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FederationListener); ok {
		listenerT.EnterFederation(s)
	}
}

func (s *FederationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FederationListener); ok {
		listenerT.ExitFederation(s)
	}
}

func (p *FederationParser) Federation() (localctx IFederationContext) {
	this := p
	_ = this

	localctx = NewFederationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FederationParserRULE_federation)

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
		p.Match(FederationParserFEDERATION)
	}

	return localctx
}
