// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603781202/Barium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barium

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

type BariumParser struct {
	*antlr.BaseParser
}

var bariumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bariumParserInit() {
	staticData := &bariumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BARIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "barium",
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

// BariumParserInit initializes any static state used to implement BariumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBariumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BariumParserInit() {
	staticData := &bariumParserStaticData
	staticData.once.Do(bariumParserInit)
}

// NewBariumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBariumParser(input antlr.TokenStream) *BariumParser {
	BariumParserInit()
	this := new(BariumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bariumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Barium.g4"

	return this
}

// BariumParser tokens.
const (
	BariumParserEOF      = antlr.TokenEOF
	BariumParserBARIUM   = 1
	BariumParserOWNKEY   = 2
	BariumParserSPLITKEY = 3
	BariumParserWS       = 4
)

// BariumParser rules.
const (
	BariumParserRULE_expression = 0
	BariumParserRULE_barium     = 1
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
	p.RuleIndex = BariumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BariumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Barium() IBariumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBariumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBariumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BariumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BariumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BariumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BariumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BariumParserRULE_expression)

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
		p.Barium()
	}
	{
		p.SetState(5)
		p.Match(BariumParserEOF)
	}

	return localctx
}

// IBariumContext is an interface to support dynamic dispatch.
type IBariumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBariumContext differentiates from other interfaces.
	IsBariumContext()
}

type BariumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBariumContext() *BariumContext {
	var p = new(BariumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BariumParserRULE_barium
	return p
}

func (*BariumContext) IsBariumContext() {}

func NewBariumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BariumContext {
	var p = new(BariumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BariumParserRULE_barium

	return p
}

func (s *BariumContext) GetParser() antlr.Parser { return s.parser }

func (s *BariumContext) BARIUM() antlr.TerminalNode {
	return s.GetToken(BariumParserBARIUM, 0)
}

func (s *BariumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BariumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BariumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BariumListener); ok {
		listenerT.EnterBarium(s)
	}
}

func (s *BariumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BariumListener); ok {
		listenerT.ExitBarium(s)
	}
}

func (p *BariumParser) Barium() (localctx IBariumContext) {
	this := p
	_ = this

	localctx = NewBariumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BariumParserRULE_barium)

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
		p.Match(BariumParserBARIUM)
	}

	return localctx
}
