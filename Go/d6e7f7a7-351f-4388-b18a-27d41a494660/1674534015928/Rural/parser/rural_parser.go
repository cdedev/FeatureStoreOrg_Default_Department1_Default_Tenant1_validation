// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674534015928/Rural.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rural

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

type RuralParser struct {
	*antlr.BaseParser
}

var ruralParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ruralParserInit() {
	staticData := &ruralParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RURAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rural",
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

// RuralParserInit initializes any static state used to implement RuralParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRuralParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RuralParserInit() {
	staticData := &ruralParserStaticData
	staticData.once.Do(ruralParserInit)
}

// NewRuralParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRuralParser(input antlr.TokenStream) *RuralParser {
	RuralParserInit()
	this := new(RuralParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ruralParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rural.g4"

	return this
}

// RuralParser tokens.
const (
	RuralParserEOF      = antlr.TokenEOF
	RuralParserRURAL    = 1
	RuralParserOWNKEY   = 2
	RuralParserSPLITKEY = 3
	RuralParserWS       = 4
)

// RuralParser rules.
const (
	RuralParserRULE_expression = 0
	RuralParserRULE_rural      = 1
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
	p.RuleIndex = RuralParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuralParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rural() IRuralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRuralContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RuralParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuralListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuralListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RuralParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RuralParserRULE_expression)

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
		p.Rural()
	}
	{
		p.SetState(5)
		p.Match(RuralParserEOF)
	}

	return localctx
}

// IRuralContext is an interface to support dynamic dispatch.
type IRuralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuralContext differentiates from other interfaces.
	IsRuralContext()
}

type RuralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuralContext() *RuralContext {
	var p = new(RuralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RuralParserRULE_rural
	return p
}

func (*RuralContext) IsRuralContext() {}

func NewRuralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuralContext {
	var p = new(RuralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RuralParserRULE_rural

	return p
}

func (s *RuralContext) GetParser() antlr.Parser { return s.parser }

func (s *RuralContext) RURAL() antlr.TerminalNode {
	return s.GetToken(RuralParserRURAL, 0)
}

func (s *RuralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuralListener); ok {
		listenerT.EnterRural(s)
	}
}

func (s *RuralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RuralListener); ok {
		listenerT.ExitRural(s)
	}
}

func (p *RuralParser) Rural() (localctx IRuralContext) {
	this := p
	_ = this

	localctx = NewRuralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RuralParserRULE_rural)

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
		p.Match(RuralParserRURAL)
	}

	return localctx
}