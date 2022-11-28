// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628628354/Phosphorous.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phosphorous

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

type PhosphorousParser struct {
	*antlr.BaseParser
}

var phosphorousParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func phosphorousParserInit() {
	staticData := &phosphorousParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PHOSPHOROUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "phosphorous",
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

// PhosphorousParserInit initializes any static state used to implement PhosphorousParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPhosphorousParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PhosphorousParserInit() {
	staticData := &phosphorousParserStaticData
	staticData.once.Do(phosphorousParserInit)
}

// NewPhosphorousParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPhosphorousParser(input antlr.TokenStream) *PhosphorousParser {
	PhosphorousParserInit()
	this := new(PhosphorousParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &phosphorousParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Phosphorous.g4"

	return this
}

// PhosphorousParser tokens.
const (
	PhosphorousParserEOF         = antlr.TokenEOF
	PhosphorousParserPHOSPHOROUS = 1
	PhosphorousParserOWNKEY      = 2
	PhosphorousParserSPLITKEY    = 3
	PhosphorousParserWS          = 4
)

// PhosphorousParser rules.
const (
	PhosphorousParserRULE_expression  = 0
	PhosphorousParserRULE_phosphorous = 1
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
	p.RuleIndex = PhosphorousParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhosphorousParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Phosphorous() IPhosphorousContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhosphorousContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhosphorousContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PhosphorousParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhosphorousListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhosphorousListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PhosphorousParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PhosphorousParserRULE_expression)

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
		p.Phosphorous()
	}
	{
		p.SetState(5)
		p.Match(PhosphorousParserEOF)
	}

	return localctx
}

// IPhosphorousContext is an interface to support dynamic dispatch.
type IPhosphorousContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhosphorousContext differentiates from other interfaces.
	IsPhosphorousContext()
}

type PhosphorousContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhosphorousContext() *PhosphorousContext {
	var p = new(PhosphorousContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PhosphorousParserRULE_phosphorous
	return p
}

func (*PhosphorousContext) IsPhosphorousContext() {}

func NewPhosphorousContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhosphorousContext {
	var p = new(PhosphorousContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhosphorousParserRULE_phosphorous

	return p
}

func (s *PhosphorousContext) GetParser() antlr.Parser { return s.parser }

func (s *PhosphorousContext) PHOSPHOROUS() antlr.TerminalNode {
	return s.GetToken(PhosphorousParserPHOSPHOROUS, 0)
}

func (s *PhosphorousContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhosphorousContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhosphorousContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhosphorousListener); ok {
		listenerT.EnterPhosphorous(s)
	}
}

func (s *PhosphorousContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhosphorousListener); ok {
		listenerT.ExitPhosphorous(s)
	}
}

func (p *PhosphorousParser) Phosphorous() (localctx IPhosphorousContext) {
	this := p
	_ = this

	localctx = NewPhosphorousContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PhosphorousParserRULE_phosphorous)

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
		p.Match(PhosphorousParserPHOSPHOROUS)
	}

	return localctx
}
