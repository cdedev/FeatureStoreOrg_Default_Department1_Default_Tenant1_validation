// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524614179/InstrumentMode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InstrumentMode

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

type InstrumentModeParser struct {
	*antlr.BaseParser
}

var instrumentmodeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func instrumentmodeParserInit() {
	staticData := &instrumentmodeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INSTRUMENTMODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "instrumentmode",
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

// InstrumentModeParserInit initializes any static state used to implement InstrumentModeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInstrumentModeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InstrumentModeParserInit() {
	staticData := &instrumentmodeParserStaticData
	staticData.once.Do(instrumentmodeParserInit)
}

// NewInstrumentModeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInstrumentModeParser(input antlr.TokenStream) *InstrumentModeParser {
	InstrumentModeParserInit()
	this := new(InstrumentModeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &instrumentmodeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "InstrumentMode.g4"

	return this
}

// InstrumentModeParser tokens.
const (
	InstrumentModeParserEOF            = antlr.TokenEOF
	InstrumentModeParserINSTRUMENTMODE = 1
	InstrumentModeParserOWNKEY         = 2
	InstrumentModeParserSPLITKEY       = 3
	InstrumentModeParserWS             = 4
)

// InstrumentModeParser rules.
const (
	InstrumentModeParserRULE_expression     = 0
	InstrumentModeParserRULE_instrumentmode = 1
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
	p.RuleIndex = InstrumentModeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InstrumentModeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Instrumentmode() IInstrumentmodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstrumentmodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstrumentmodeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(InstrumentModeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstrumentModeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstrumentModeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *InstrumentModeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InstrumentModeParserRULE_expression)

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
		p.Instrumentmode()
	}
	{
		p.SetState(5)
		p.Match(InstrumentModeParserEOF)
	}

	return localctx
}

// IInstrumentmodeContext is an interface to support dynamic dispatch.
type IInstrumentmodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstrumentmodeContext differentiates from other interfaces.
	IsInstrumentmodeContext()
}

type InstrumentmodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstrumentmodeContext() *InstrumentmodeContext {
	var p = new(InstrumentmodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InstrumentModeParserRULE_instrumentmode
	return p
}

func (*InstrumentmodeContext) IsInstrumentmodeContext() {}

func NewInstrumentmodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstrumentmodeContext {
	var p = new(InstrumentmodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InstrumentModeParserRULE_instrumentmode

	return p
}

func (s *InstrumentmodeContext) GetParser() antlr.Parser { return s.parser }

func (s *InstrumentmodeContext) INSTRUMENTMODE() antlr.TerminalNode {
	return s.GetToken(InstrumentModeParserINSTRUMENTMODE, 0)
}

func (s *InstrumentmodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrumentmodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstrumentmodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstrumentModeListener); ok {
		listenerT.EnterInstrumentmode(s)
	}
}

func (s *InstrumentmodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InstrumentModeListener); ok {
		listenerT.ExitInstrumentmode(s)
	}
}

func (p *InstrumentModeParser) Instrumentmode() (localctx IInstrumentmodeContext) {
	this := p
	_ = this

	localctx = NewInstrumentmodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InstrumentModeParserRULE_instrumentmode)

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
		p.Match(InstrumentModeParserINSTRUMENTMODE)
	}

	return localctx
}
