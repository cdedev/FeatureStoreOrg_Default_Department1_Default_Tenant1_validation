// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868350895/Wood.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wood

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

type WoodParser struct {
	*antlr.BaseParser
}

var woodParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func woodParserInit() {
	staticData := &woodParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WOOD", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "wood",
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

// WoodParserInit initializes any static state used to implement WoodParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWoodParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WoodParserInit() {
	staticData := &woodParserStaticData
	staticData.once.Do(woodParserInit)
}

// NewWoodParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWoodParser(input antlr.TokenStream) *WoodParser {
	WoodParserInit()
	this := new(WoodParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &woodParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Wood.g4"

	return this
}

// WoodParser tokens.
const (
	WoodParserEOF      = antlr.TokenEOF
	WoodParserWOOD     = 1
	WoodParserOWNKEY   = 2
	WoodParserSPLITKEY = 3
	WoodParserWS       = 4
)

// WoodParser rules.
const (
	WoodParserRULE_expression = 0
	WoodParserRULE_wood       = 1
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
	p.RuleIndex = WoodParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WoodParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Wood() IWoodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWoodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWoodContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WoodParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WoodListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WoodListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WoodParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WoodParserRULE_expression)

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
		p.Wood()
	}
	{
		p.SetState(5)
		p.Match(WoodParserEOF)
	}

	return localctx
}

// IWoodContext is an interface to support dynamic dispatch.
type IWoodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWoodContext differentiates from other interfaces.
	IsWoodContext()
}

type WoodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWoodContext() *WoodContext {
	var p = new(WoodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WoodParserRULE_wood
	return p
}

func (*WoodContext) IsWoodContext() {}

func NewWoodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WoodContext {
	var p = new(WoodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WoodParserRULE_wood

	return p
}

func (s *WoodContext) GetParser() antlr.Parser { return s.parser }

func (s *WoodContext) WOOD() antlr.TerminalNode {
	return s.GetToken(WoodParserWOOD, 0)
}

func (s *WoodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WoodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WoodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WoodListener); ok {
		listenerT.EnterWood(s)
	}
}

func (s *WoodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WoodListener); ok {
		listenerT.ExitWood(s)
	}
}

func (p *WoodParser) Wood() (localctx IWoodContext) {
	this := p
	_ = this

	localctx = NewWoodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WoodParserRULE_wood)

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
		p.Match(WoodParserWOOD)
	}

	return localctx
}
