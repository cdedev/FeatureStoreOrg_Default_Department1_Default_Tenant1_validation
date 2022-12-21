// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601050396/Agency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agency

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

type AgencyParser struct {
	*antlr.BaseParser
}

var agencyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func agencyParserInit() {
	staticData := &agencyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AGENCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "agency",
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

// AgencyParserInit initializes any static state used to implement AgencyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAgencyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AgencyParserInit() {
	staticData := &agencyParserStaticData
	staticData.once.Do(agencyParserInit)
}

// NewAgencyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAgencyParser(input antlr.TokenStream) *AgencyParser {
	AgencyParserInit()
	this := new(AgencyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &agencyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Agency.g4"

	return this
}

// AgencyParser tokens.
const (
	AgencyParserEOF      = antlr.TokenEOF
	AgencyParserAGENCY   = 1
	AgencyParserOWNKEY   = 2
	AgencyParserSPLITKEY = 3
	AgencyParserWS       = 4
)

// AgencyParser rules.
const (
	AgencyParserRULE_expression = 0
	AgencyParserRULE_agency     = 1
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
	p.RuleIndex = AgencyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgencyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Agency() IAgencyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAgencyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAgencyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AgencyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgencyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgencyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AgencyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AgencyParserRULE_expression)

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
		p.Agency()
	}
	{
		p.SetState(5)
		p.Match(AgencyParserEOF)
	}

	return localctx
}

// IAgencyContext is an interface to support dynamic dispatch.
type IAgencyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAgencyContext differentiates from other interfaces.
	IsAgencyContext()
}

type AgencyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAgencyContext() *AgencyContext {
	var p = new(AgencyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AgencyParserRULE_agency
	return p
}

func (*AgencyContext) IsAgencyContext() {}

func NewAgencyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AgencyContext {
	var p = new(AgencyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AgencyParserRULE_agency

	return p
}

func (s *AgencyContext) GetParser() antlr.Parser { return s.parser }

func (s *AgencyContext) AGENCY() antlr.TerminalNode {
	return s.GetToken(AgencyParserAGENCY, 0)
}

func (s *AgencyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AgencyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AgencyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgencyListener); ok {
		listenerT.EnterAgency(s)
	}
}

func (s *AgencyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AgencyListener); ok {
		listenerT.ExitAgency(s)
	}
}

func (p *AgencyParser) Agency() (localctx IAgencyContext) {
	this := p
	_ = this

	localctx = NewAgencyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AgencyParserRULE_agency)

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
		p.Match(AgencyParserAGENCY)
	}

	return localctx
}
