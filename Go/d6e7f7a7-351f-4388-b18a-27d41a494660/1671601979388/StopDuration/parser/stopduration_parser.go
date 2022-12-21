// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601979388/StopDuration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StopDuration

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

type StopDurationParser struct {
	*antlr.BaseParser
}

var stopdurationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func stopdurationParserInit() {
	staticData := &stopdurationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STOPDURATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "stopduration",
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

// StopDurationParserInit initializes any static state used to implement StopDurationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStopDurationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StopDurationParserInit() {
	staticData := &stopdurationParserStaticData
	staticData.once.Do(stopdurationParserInit)
}

// NewStopDurationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStopDurationParser(input antlr.TokenStream) *StopDurationParser {
	StopDurationParserInit()
	this := new(StopDurationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &stopdurationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "StopDuration.g4"

	return this
}

// StopDurationParser tokens.
const (
	StopDurationParserEOF          = antlr.TokenEOF
	StopDurationParserSTOPDURATION = 1
	StopDurationParserOWNKEY       = 2
	StopDurationParserSPLITKEY     = 3
	StopDurationParserWS           = 4
)

// StopDurationParser rules.
const (
	StopDurationParserRULE_expression   = 0
	StopDurationParserRULE_stopduration = 1
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
	p.RuleIndex = StopDurationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StopDurationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Stopduration() IStopdurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStopdurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStopdurationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StopDurationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StopDurationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StopDurationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StopDurationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StopDurationParserRULE_expression)

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
		p.Stopduration()
	}
	{
		p.SetState(5)
		p.Match(StopDurationParserEOF)
	}

	return localctx
}

// IStopdurationContext is an interface to support dynamic dispatch.
type IStopdurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStopdurationContext differentiates from other interfaces.
	IsStopdurationContext()
}

type StopdurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStopdurationContext() *StopdurationContext {
	var p = new(StopdurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StopDurationParserRULE_stopduration
	return p
}

func (*StopdurationContext) IsStopdurationContext() {}

func NewStopdurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StopdurationContext {
	var p = new(StopdurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StopDurationParserRULE_stopduration

	return p
}

func (s *StopdurationContext) GetParser() antlr.Parser { return s.parser }

func (s *StopdurationContext) STOPDURATION() antlr.TerminalNode {
	return s.GetToken(StopDurationParserSTOPDURATION, 0)
}

func (s *StopdurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StopdurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StopdurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StopDurationListener); ok {
		listenerT.EnterStopduration(s)
	}
}

func (s *StopdurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StopDurationListener); ok {
		listenerT.ExitStopduration(s)
	}
}

func (p *StopDurationParser) Stopduration() (localctx IStopdurationContext) {
	this := p
	_ = this

	localctx = NewStopdurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StopDurationParserRULE_stopduration)

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
		p.Match(StopDurationParserSTOPDURATION)
	}

	return localctx
}
