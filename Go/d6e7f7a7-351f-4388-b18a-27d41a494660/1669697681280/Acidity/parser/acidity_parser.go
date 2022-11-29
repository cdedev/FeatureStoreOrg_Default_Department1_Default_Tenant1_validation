// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697681280/Acidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acidity

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

type AcidityParser struct {
	*antlr.BaseParser
}

var acidityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func acidityParserInit() {
	staticData := &acidityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ACIDITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "acidity",
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

// AcidityParserInit initializes any static state used to implement AcidityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAcidityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AcidityParserInit() {
	staticData := &acidityParserStaticData
	staticData.once.Do(acidityParserInit)
}

// NewAcidityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAcidityParser(input antlr.TokenStream) *AcidityParser {
	AcidityParserInit()
	this := new(AcidityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &acidityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Acidity.g4"

	return this
}

// AcidityParser tokens.
const (
	AcidityParserEOF      = antlr.TokenEOF
	AcidityParserACIDITY  = 1
	AcidityParserOWNKEY   = 2
	AcidityParserSPLITKEY = 3
	AcidityParserWS       = 4
)

// AcidityParser rules.
const (
	AcidityParserRULE_expression = 0
	AcidityParserRULE_acidity    = 1
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
	p.RuleIndex = AcidityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AcidityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Acidity() IAcidityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAcidityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAcidityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AcidityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcidityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcidityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AcidityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AcidityParserRULE_expression)

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
		p.Acidity()
	}
	{
		p.SetState(5)
		p.Match(AcidityParserEOF)
	}

	return localctx
}

// IAcidityContext is an interface to support dynamic dispatch.
type IAcidityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAcidityContext differentiates from other interfaces.
	IsAcidityContext()
}

type AcidityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAcidityContext() *AcidityContext {
	var p = new(AcidityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AcidityParserRULE_acidity
	return p
}

func (*AcidityContext) IsAcidityContext() {}

func NewAcidityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AcidityContext {
	var p = new(AcidityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AcidityParserRULE_acidity

	return p
}

func (s *AcidityContext) GetParser() antlr.Parser { return s.parser }

func (s *AcidityContext) ACIDITY() antlr.TerminalNode {
	return s.GetToken(AcidityParserACIDITY, 0)
}

func (s *AcidityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcidityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AcidityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcidityListener); ok {
		listenerT.EnterAcidity(s)
	}
}

func (s *AcidityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcidityListener); ok {
		listenerT.ExitAcidity(s)
	}
}

func (p *AcidityParser) Acidity() (localctx IAcidityContext) {
	this := p
	_ = this

	localctx = NewAcidityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AcidityParserRULE_acidity)

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
		p.Match(AcidityParserACIDITY)
	}

	return localctx
}
