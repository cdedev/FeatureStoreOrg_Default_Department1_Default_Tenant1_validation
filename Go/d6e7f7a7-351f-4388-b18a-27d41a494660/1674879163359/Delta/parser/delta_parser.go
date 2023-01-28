// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674879163359/Delta.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Delta

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

type DeltaParser struct {
	*antlr.BaseParser
}

var deltaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func deltaParserInit() {
	staticData := &deltaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DELTA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "delta",
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

// DeltaParserInit initializes any static state used to implement DeltaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDeltaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DeltaParserInit() {
	staticData := &deltaParserStaticData
	staticData.once.Do(deltaParserInit)
}

// NewDeltaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDeltaParser(input antlr.TokenStream) *DeltaParser {
	DeltaParserInit()
	this := new(DeltaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &deltaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Delta.g4"

	return this
}

// DeltaParser tokens.
const (
	DeltaParserEOF      = antlr.TokenEOF
	DeltaParserDELTA    = 1
	DeltaParserOWNKEY   = 2
	DeltaParserSPLITKEY = 3
	DeltaParserWS       = 4
)

// DeltaParser rules.
const (
	DeltaParserRULE_expression = 0
	DeltaParserRULE_delta      = 1
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
	p.RuleIndex = DeltaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeltaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Delta() IDeltaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeltaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeltaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DeltaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeltaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeltaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DeltaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DeltaParserRULE_expression)

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
		p.Delta()
	}
	{
		p.SetState(5)
		p.Match(DeltaParserEOF)
	}

	return localctx
}

// IDeltaContext is an interface to support dynamic dispatch.
type IDeltaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeltaContext differentiates from other interfaces.
	IsDeltaContext()
}

type DeltaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeltaContext() *DeltaContext {
	var p = new(DeltaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DeltaParserRULE_delta
	return p
}

func (*DeltaContext) IsDeltaContext() {}

func NewDeltaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeltaContext {
	var p = new(DeltaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DeltaParserRULE_delta

	return p
}

func (s *DeltaContext) GetParser() antlr.Parser { return s.parser }

func (s *DeltaContext) DELTA() antlr.TerminalNode {
	return s.GetToken(DeltaParserDELTA, 0)
}

func (s *DeltaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeltaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeltaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeltaListener); ok {
		listenerT.EnterDelta(s)
	}
}

func (s *DeltaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeltaListener); ok {
		listenerT.ExitDelta(s)
	}
}

func (p *DeltaParser) Delta() (localctx IDeltaContext) {
	this := p
	_ = this

	localctx = NewDeltaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DeltaParserRULE_delta)

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
		p.Match(DeltaParserDELTA)
	}

	return localctx
}
