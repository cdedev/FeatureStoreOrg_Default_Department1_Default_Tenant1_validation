// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675223158255/Confidence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Confidence

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

type ConfidenceParser struct {
	*antlr.BaseParser
}

var confidenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func confidenceParserInit() {
	staticData := &confidenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONFIDENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "confidence",
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

// ConfidenceParserInit initializes any static state used to implement ConfidenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConfidenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConfidenceParserInit() {
	staticData := &confidenceParserStaticData
	staticData.once.Do(confidenceParserInit)
}

// NewConfidenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConfidenceParser(input antlr.TokenStream) *ConfidenceParser {
	ConfidenceParserInit()
	this := new(ConfidenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &confidenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Confidence.g4"

	return this
}

// ConfidenceParser tokens.
const (
	ConfidenceParserEOF        = antlr.TokenEOF
	ConfidenceParserCONFIDENCE = 1
	ConfidenceParserOWNKEY     = 2
	ConfidenceParserSPLITKEY   = 3
	ConfidenceParserWS         = 4
)

// ConfidenceParser rules.
const (
	ConfidenceParserRULE_expression = 0
	ConfidenceParserRULE_confidence = 1
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
	p.RuleIndex = ConfidenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfidenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Confidence() IConfidenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConfidenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConfidenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConfidenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfidenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfidenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConfidenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConfidenceParserRULE_expression)

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
		p.Confidence()
	}
	{
		p.SetState(5)
		p.Match(ConfidenceParserEOF)
	}

	return localctx
}

// IConfidenceContext is an interface to support dynamic dispatch.
type IConfidenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConfidenceContext differentiates from other interfaces.
	IsConfidenceContext()
}

type ConfidenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConfidenceContext() *ConfidenceContext {
	var p = new(ConfidenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConfidenceParserRULE_confidence
	return p
}

func (*ConfidenceContext) IsConfidenceContext() {}

func NewConfidenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConfidenceContext {
	var p = new(ConfidenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConfidenceParserRULE_confidence

	return p
}

func (s *ConfidenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ConfidenceContext) CONFIDENCE() antlr.TerminalNode {
	return s.GetToken(ConfidenceParserCONFIDENCE, 0)
}

func (s *ConfidenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConfidenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConfidenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfidenceListener); ok {
		listenerT.EnterConfidence(s)
	}
}

func (s *ConfidenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConfidenceListener); ok {
		listenerT.ExitConfidence(s)
	}
}

func (p *ConfidenceParser) Confidence() (localctx IConfidenceContext) {
	this := p
	_ = this

	localctx = NewConfidenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConfidenceParserRULE_confidence)

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
		p.Match(ConfidenceParserCONFIDENCE)
	}

	return localctx
}
