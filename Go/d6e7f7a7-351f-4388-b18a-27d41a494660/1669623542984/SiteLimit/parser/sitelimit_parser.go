// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623542984/SiteLimit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SiteLimit

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

type SiteLimitParser struct {
	*antlr.BaseParser
}

var sitelimitParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sitelimitParserInit() {
	staticData := &sitelimitParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SITELIMIT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sitelimit",
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

// SiteLimitParserInit initializes any static state used to implement SiteLimitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSiteLimitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SiteLimitParserInit() {
	staticData := &sitelimitParserStaticData
	staticData.once.Do(sitelimitParserInit)
}

// NewSiteLimitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSiteLimitParser(input antlr.TokenStream) *SiteLimitParser {
	SiteLimitParserInit()
	this := new(SiteLimitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sitelimitParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SiteLimit.g4"

	return this
}

// SiteLimitParser tokens.
const (
	SiteLimitParserEOF       = antlr.TokenEOF
	SiteLimitParserSITELIMIT = 1
	SiteLimitParserOWNKEY    = 2
	SiteLimitParserSPLITKEY  = 3
	SiteLimitParserWS        = 4
)

// SiteLimitParser rules.
const (
	SiteLimitParserRULE_expression = 0
	SiteLimitParserRULE_sitelimit  = 1
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
	p.RuleIndex = SiteLimitParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SiteLimitParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sitelimit() ISitelimitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISitelimitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISitelimitContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SiteLimitParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SiteLimitListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SiteLimitListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SiteLimitParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SiteLimitParserRULE_expression)

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
		p.Sitelimit()
	}
	{
		p.SetState(5)
		p.Match(SiteLimitParserEOF)
	}

	return localctx
}

// ISitelimitContext is an interface to support dynamic dispatch.
type ISitelimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSitelimitContext differentiates from other interfaces.
	IsSitelimitContext()
}

type SitelimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySitelimitContext() *SitelimitContext {
	var p = new(SitelimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SiteLimitParserRULE_sitelimit
	return p
}

func (*SitelimitContext) IsSitelimitContext() {}

func NewSitelimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SitelimitContext {
	var p = new(SitelimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SiteLimitParserRULE_sitelimit

	return p
}

func (s *SitelimitContext) GetParser() antlr.Parser { return s.parser }

func (s *SitelimitContext) SITELIMIT() antlr.TerminalNode {
	return s.GetToken(SiteLimitParserSITELIMIT, 0)
}

func (s *SitelimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SitelimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SitelimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SiteLimitListener); ok {
		listenerT.EnterSitelimit(s)
	}
}

func (s *SitelimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SiteLimitListener); ok {
		listenerT.ExitSitelimit(s)
	}
}

func (p *SiteLimitParser) Sitelimit() (localctx ISitelimitContext) {
	this := p
	_ = this

	localctx = NewSitelimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SiteLimitParserRULE_sitelimit)

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
		p.Match(SiteLimitParserSITELIMIT)
	}

	return localctx
}
