// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780393756/Valence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Valence

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

type ValenceParser struct {
	*antlr.BaseParser
}

var valenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func valenceParserInit() {
	staticData := &valenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VALENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "valence",
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

// ValenceParserInit initializes any static state used to implement ValenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewValenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ValenceParserInit() {
	staticData := &valenceParserStaticData
	staticData.once.Do(valenceParserInit)
}

// NewValenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewValenceParser(input antlr.TokenStream) *ValenceParser {
	ValenceParserInit()
	this := new(ValenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &valenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Valence.g4"

	return this
}

// ValenceParser tokens.
const (
	ValenceParserEOF      = antlr.TokenEOF
	ValenceParserVALENCE  = 1
	ValenceParserOWNKEY   = 2
	ValenceParserSPLITKEY = 3
	ValenceParserWS       = 4
)

// ValenceParser rules.
const (
	ValenceParserRULE_expression = 0
	ValenceParserRULE_valence    = 1
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
	p.RuleIndex = ValenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Valence() IValenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ValenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ValenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ValenceParserRULE_expression)

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
		p.Valence()
	}
	{
		p.SetState(5)
		p.Match(ValenceParserEOF)
	}

	return localctx
}

// IValenceContext is an interface to support dynamic dispatch.
type IValenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValenceContext differentiates from other interfaces.
	IsValenceContext()
}

type ValenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValenceContext() *ValenceContext {
	var p = new(ValenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ValenceParserRULE_valence
	return p
}

func (*ValenceContext) IsValenceContext() {}

func NewValenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValenceContext {
	var p = new(ValenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ValenceParserRULE_valence

	return p
}

func (s *ValenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ValenceContext) VALENCE() antlr.TerminalNode {
	return s.GetToken(ValenceParserVALENCE, 0)
}

func (s *ValenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValenceListener); ok {
		listenerT.EnterValence(s)
	}
}

func (s *ValenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ValenceListener); ok {
		listenerT.ExitValence(s)
	}
}

func (p *ValenceParser) Valence() (localctx IValenceContext) {
	this := p
	_ = this

	localctx = NewValenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ValenceParserRULE_valence)

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
		p.Match(ValenceParserVALENCE)
	}

	return localctx
}
