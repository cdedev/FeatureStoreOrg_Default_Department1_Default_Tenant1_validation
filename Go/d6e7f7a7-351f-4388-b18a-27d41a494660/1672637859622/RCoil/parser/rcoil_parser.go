// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637859622/RCoil.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RCoil

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

type RCoilParser struct {
	*antlr.BaseParser
}

var rcoilParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rcoilParserInit() {
	staticData := &rcoilParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RCOIL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rcoil",
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

// RCoilParserInit initializes any static state used to implement RCoilParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRCoilParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RCoilParserInit() {
	staticData := &rcoilParserStaticData
	staticData.once.Do(rcoilParserInit)
}

// NewRCoilParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRCoilParser(input antlr.TokenStream) *RCoilParser {
	RCoilParserInit()
	this := new(RCoilParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rcoilParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RCoil.g4"

	return this
}

// RCoilParser tokens.
const (
	RCoilParserEOF      = antlr.TokenEOF
	RCoilParserRCOIL    = 1
	RCoilParserOWNKEY   = 2
	RCoilParserSPLITKEY = 3
	RCoilParserWS       = 4
)

// RCoilParser rules.
const (
	RCoilParserRULE_expression = 0
	RCoilParserRULE_rcoil      = 1
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
	p.RuleIndex = RCoilParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCoilParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rcoil() IRcoilContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRcoilContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRcoilContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RCoilParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCoilListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCoilListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RCoilParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RCoilParserRULE_expression)

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
		p.Rcoil()
	}
	{
		p.SetState(5)
		p.Match(RCoilParserEOF)
	}

	return localctx
}

// IRcoilContext is an interface to support dynamic dispatch.
type IRcoilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRcoilContext differentiates from other interfaces.
	IsRcoilContext()
}

type RcoilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRcoilContext() *RcoilContext {
	var p = new(RcoilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RCoilParserRULE_rcoil
	return p
}

func (*RcoilContext) IsRcoilContext() {}

func NewRcoilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RcoilContext {
	var p = new(RcoilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RCoilParserRULE_rcoil

	return p
}

func (s *RcoilContext) GetParser() antlr.Parser { return s.parser }

func (s *RcoilContext) RCOIL() antlr.TerminalNode {
	return s.GetToken(RCoilParserRCOIL, 0)
}

func (s *RcoilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RcoilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RcoilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCoilListener); ok {
		listenerT.EnterRcoil(s)
	}
}

func (s *RcoilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RCoilListener); ok {
		listenerT.ExitRcoil(s)
	}
}

func (p *RCoilParser) Rcoil() (localctx IRcoilContext) {
	this := p
	_ = this

	localctx = NewRcoilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RCoilParserRULE_rcoil)

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
		p.Match(RCoilParserRCOIL)
	}

	return localctx
}
