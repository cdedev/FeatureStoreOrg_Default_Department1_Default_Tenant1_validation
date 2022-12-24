// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671867422033/ExchUsd.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ExchUsd

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

type ExchUsdParser struct {
	*antlr.BaseParser
}

var exchusdParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func exchusdParserInit() {
	staticData := &exchusdParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EXCHUSD", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "exchusd",
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

// ExchUsdParserInit initializes any static state used to implement ExchUsdParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExchUsdParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExchUsdParserInit() {
	staticData := &exchusdParserStaticData
	staticData.once.Do(exchusdParserInit)
}

// NewExchUsdParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExchUsdParser(input antlr.TokenStream) *ExchUsdParser {
	ExchUsdParserInit()
	this := new(ExchUsdParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &exchusdParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ExchUsd.g4"

	return this
}

// ExchUsdParser tokens.
const (
	ExchUsdParserEOF      = antlr.TokenEOF
	ExchUsdParserEXCHUSD  = 1
	ExchUsdParserOWNKEY   = 2
	ExchUsdParserSPLITKEY = 3
	ExchUsdParserWS       = 4
)

// ExchUsdParser rules.
const (
	ExchUsdParserRULE_expression = 0
	ExchUsdParserRULE_exchusd    = 1
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
	p.RuleIndex = ExchUsdParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExchUsdParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Exchusd() IExchusdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExchusdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExchusdContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExchUsdParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExchUsdListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExchUsdListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExchUsdParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExchUsdParserRULE_expression)

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
		p.Exchusd()
	}
	{
		p.SetState(5)
		p.Match(ExchUsdParserEOF)
	}

	return localctx
}

// IExchusdContext is an interface to support dynamic dispatch.
type IExchusdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExchusdContext differentiates from other interfaces.
	IsExchusdContext()
}

type ExchusdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExchusdContext() *ExchusdContext {
	var p = new(ExchusdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExchUsdParserRULE_exchusd
	return p
}

func (*ExchusdContext) IsExchusdContext() {}

func NewExchusdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExchusdContext {
	var p = new(ExchusdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExchUsdParserRULE_exchusd

	return p
}

func (s *ExchusdContext) GetParser() antlr.Parser { return s.parser }

func (s *ExchusdContext) EXCHUSD() antlr.TerminalNode {
	return s.GetToken(ExchUsdParserEXCHUSD, 0)
}

func (s *ExchusdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExchusdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExchusdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExchUsdListener); ok {
		listenerT.EnterExchusd(s)
	}
}

func (s *ExchusdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExchUsdListener); ok {
		listenerT.ExitExchusd(s)
	}
}

func (p *ExchUsdParser) Exchusd() (localctx IExchusdContext) {
	this := p
	_ = this

	localctx = NewExchusdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExchUsdParserRULE_exchusd)

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
		p.Match(ExchUsdParserEXCHUSD)
	}

	return localctx
}
