// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675246511321/Offence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Offence

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

type OffenceParser struct {
	*antlr.BaseParser
}

var offenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func offenceParserInit() {
	staticData := &offenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OFFENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "offence",
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

// OffenceParserInit initializes any static state used to implement OffenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOffenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OffenceParserInit() {
	staticData := &offenceParserStaticData
	staticData.once.Do(offenceParserInit)
}

// NewOffenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOffenceParser(input antlr.TokenStream) *OffenceParser {
	OffenceParserInit()
	this := new(OffenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &offenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Offence.g4"

	return this
}

// OffenceParser tokens.
const (
	OffenceParserEOF      = antlr.TokenEOF
	OffenceParserOFFENCE  = 1
	OffenceParserOWNKEY   = 2
	OffenceParserSPLITKEY = 3
	OffenceParserWS       = 4
)

// OffenceParser rules.
const (
	OffenceParserRULE_expression = 0
	OffenceParserRULE_offence    = 1
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
	p.RuleIndex = OffenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OffenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Offence() IOffenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OffenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OffenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OffenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OffenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OffenceParserRULE_expression)

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
		p.Offence()
	}
	{
		p.SetState(5)
		p.Match(OffenceParserEOF)
	}

	return localctx
}

// IOffenceContext is an interface to support dynamic dispatch.
type IOffenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffenceContext differentiates from other interfaces.
	IsOffenceContext()
}

type OffenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffenceContext() *OffenceContext {
	var p = new(OffenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OffenceParserRULE_offence
	return p
}

func (*OffenceContext) IsOffenceContext() {}

func NewOffenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffenceContext {
	var p = new(OffenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OffenceParserRULE_offence

	return p
}

func (s *OffenceContext) GetParser() antlr.Parser { return s.parser }

func (s *OffenceContext) OFFENCE() antlr.TerminalNode {
	return s.GetToken(OffenceParserOFFENCE, 0)
}

func (s *OffenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OffenceListener); ok {
		listenerT.EnterOffence(s)
	}
}

func (s *OffenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OffenceListener); ok {
		listenerT.ExitOffence(s)
	}
}

func (p *OffenceParser) Offence() (localctx IOffenceContext) {
	this := p
	_ = this

	localctx = NewOffenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OffenceParserRULE_offence)

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
		p.Match(OffenceParserOFFENCE)
	}

	return localctx
}
