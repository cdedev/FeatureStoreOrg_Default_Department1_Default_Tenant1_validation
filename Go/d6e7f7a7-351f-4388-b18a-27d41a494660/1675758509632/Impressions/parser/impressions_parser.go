// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758509632/Impressions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Impressions

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

type ImpressionsParser struct {
	*antlr.BaseParser
}

var impressionsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func impressionsParserInit() {
	staticData := &impressionsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "IMPRESSIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "impressions",
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

// ImpressionsParserInit initializes any static state used to implement ImpressionsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewImpressionsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ImpressionsParserInit() {
	staticData := &impressionsParserStaticData
	staticData.once.Do(impressionsParserInit)
}

// NewImpressionsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewImpressionsParser(input antlr.TokenStream) *ImpressionsParser {
	ImpressionsParserInit()
	this := new(ImpressionsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &impressionsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Impressions.g4"

	return this
}

// ImpressionsParser tokens.
const (
	ImpressionsParserEOF         = antlr.TokenEOF
	ImpressionsParserIMPRESSIONS = 1
	ImpressionsParserOWNKEY      = 2
	ImpressionsParserSPLITKEY    = 3
	ImpressionsParserWS          = 4
)

// ImpressionsParser rules.
const (
	ImpressionsParserRULE_expression  = 0
	ImpressionsParserRULE_impressions = 1
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
	p.RuleIndex = ImpressionsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImpressionsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Impressions() IImpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImpressionsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ImpressionsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImpressionsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImpressionsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ImpressionsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ImpressionsParserRULE_expression)

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
		p.Impressions()
	}
	{
		p.SetState(5)
		p.Match(ImpressionsParserEOF)
	}

	return localctx
}

// IImpressionsContext is an interface to support dynamic dispatch.
type IImpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpressionsContext differentiates from other interfaces.
	IsImpressionsContext()
}

type ImpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpressionsContext() *ImpressionsContext {
	var p = new(ImpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ImpressionsParserRULE_impressions
	return p
}

func (*ImpressionsContext) IsImpressionsContext() {}

func NewImpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpressionsContext {
	var p = new(ImpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ImpressionsParserRULE_impressions

	return p
}

func (s *ImpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpressionsContext) IMPRESSIONS() antlr.TerminalNode {
	return s.GetToken(ImpressionsParserIMPRESSIONS, 0)
}

func (s *ImpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImpressionsListener); ok {
		listenerT.EnterImpressions(s)
	}
}

func (s *ImpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImpressionsListener); ok {
		listenerT.ExitImpressions(s)
	}
}

func (p *ImpressionsParser) Impressions() (localctx IImpressionsContext) {
	this := p
	_ = this

	localctx = NewImpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ImpressionsParserRULE_impressions)

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
		p.Match(ImpressionsParserIMPRESSIONS)
	}

	return localctx
}
