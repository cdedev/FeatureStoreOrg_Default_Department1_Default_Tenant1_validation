// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653223671/Slope.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Slope

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

type SlopeParser struct {
	*antlr.BaseParser
}

var slopeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func slopeParserInit() {
	staticData := &slopeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SLOPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "slope",
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

// SlopeParserInit initializes any static state used to implement SlopeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSlopeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SlopeParserInit() {
	staticData := &slopeParserStaticData
	staticData.once.Do(slopeParserInit)
}

// NewSlopeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSlopeParser(input antlr.TokenStream) *SlopeParser {
	SlopeParserInit()
	this := new(SlopeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &slopeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Slope.g4"

	return this
}

// SlopeParser tokens.
const (
	SlopeParserEOF      = antlr.TokenEOF
	SlopeParserSLOPE    = 1
	SlopeParserOWNKEY   = 2
	SlopeParserSPLITKEY = 3
	SlopeParserWS       = 4
)

// SlopeParser rules.
const (
	SlopeParserRULE_expression = 0
	SlopeParserRULE_slope      = 1
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
	p.RuleIndex = SlopeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SlopeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Slope() ISlopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISlopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISlopeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SlopeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SlopeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SlopeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SlopeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SlopeParserRULE_expression)

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
		p.Slope()
	}
	{
		p.SetState(5)
		p.Match(SlopeParserEOF)
	}

	return localctx
}

// ISlopeContext is an interface to support dynamic dispatch.
type ISlopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlopeContext differentiates from other interfaces.
	IsSlopeContext()
}

type SlopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlopeContext() *SlopeContext {
	var p = new(SlopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SlopeParserRULE_slope
	return p
}

func (*SlopeContext) IsSlopeContext() {}

func NewSlopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlopeContext {
	var p = new(SlopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SlopeParserRULE_slope

	return p
}

func (s *SlopeContext) GetParser() antlr.Parser { return s.parser }

func (s *SlopeContext) SLOPE() antlr.TerminalNode {
	return s.GetToken(SlopeParserSLOPE, 0)
}

func (s *SlopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SlopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SlopeListener); ok {
		listenerT.EnterSlope(s)
	}
}

func (s *SlopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SlopeListener); ok {
		listenerT.ExitSlope(s)
	}
}

func (p *SlopeParser) Slope() (localctx ISlopeContext) {
	this := p
	_ = this

	localctx = NewSlopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SlopeParserRULE_slope)

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
		p.Match(SlopeParserSLOPE)
	}

	return localctx
}
