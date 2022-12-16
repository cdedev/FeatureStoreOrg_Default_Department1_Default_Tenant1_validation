// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203783527/RapeLegacy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RapeLegacy

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

type RapeLegacyParser struct {
	*antlr.BaseParser
}

var rapelegacyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rapelegacyParserInit() {
	staticData := &rapelegacyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RAPELEGACY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rapelegacy",
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

// RapeLegacyParserInit initializes any static state used to implement RapeLegacyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRapeLegacyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RapeLegacyParserInit() {
	staticData := &rapelegacyParserStaticData
	staticData.once.Do(rapelegacyParserInit)
}

// NewRapeLegacyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRapeLegacyParser(input antlr.TokenStream) *RapeLegacyParser {
	RapeLegacyParserInit()
	this := new(RapeLegacyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rapelegacyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RapeLegacy.g4"

	return this
}

// RapeLegacyParser tokens.
const (
	RapeLegacyParserEOF        = antlr.TokenEOF
	RapeLegacyParserRAPELEGACY = 1
	RapeLegacyParserOWNKEY     = 2
	RapeLegacyParserSPLITKEY   = 3
	RapeLegacyParserWS         = 4
)

// RapeLegacyParser rules.
const (
	RapeLegacyParserRULE_expression = 0
	RapeLegacyParserRULE_rapelegacy = 1
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
	p.RuleIndex = RapeLegacyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RapeLegacyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rapelegacy() IRapelegacyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRapelegacyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRapelegacyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RapeLegacyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RapeLegacyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RapeLegacyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RapeLegacyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RapeLegacyParserRULE_expression)

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
		p.Rapelegacy()
	}
	{
		p.SetState(5)
		p.Match(RapeLegacyParserEOF)
	}

	return localctx
}

// IRapelegacyContext is an interface to support dynamic dispatch.
type IRapelegacyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRapelegacyContext differentiates from other interfaces.
	IsRapelegacyContext()
}

type RapelegacyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRapelegacyContext() *RapelegacyContext {
	var p = new(RapelegacyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RapeLegacyParserRULE_rapelegacy
	return p
}

func (*RapelegacyContext) IsRapelegacyContext() {}

func NewRapelegacyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RapelegacyContext {
	var p = new(RapelegacyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RapeLegacyParserRULE_rapelegacy

	return p
}

func (s *RapelegacyContext) GetParser() antlr.Parser { return s.parser }

func (s *RapelegacyContext) RAPELEGACY() antlr.TerminalNode {
	return s.GetToken(RapeLegacyParserRAPELEGACY, 0)
}

func (s *RapelegacyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RapelegacyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RapelegacyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RapeLegacyListener); ok {
		listenerT.EnterRapelegacy(s)
	}
}

func (s *RapelegacyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RapeLegacyListener); ok {
		listenerT.ExitRapelegacy(s)
	}
}

func (p *RapeLegacyParser) Rapelegacy() (localctx IRapelegacyContext) {
	this := p
	_ = this

	localctx = NewRapelegacyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RapeLegacyParserRULE_rapelegacy)

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
		p.Match(RapeLegacyParserRAPELEGACY)
	}

	return localctx
}
