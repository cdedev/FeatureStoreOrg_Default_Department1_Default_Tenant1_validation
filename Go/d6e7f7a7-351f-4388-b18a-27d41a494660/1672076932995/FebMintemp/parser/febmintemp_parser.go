// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076932995/FebMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FebMintemp

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

type FebMintempParser struct {
	*antlr.BaseParser
}

var febmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func febmintempParserInit() {
	staticData := &febmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FEBMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "febmintemp",
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

// FebMintempParserInit initializes any static state used to implement FebMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFebMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FebMintempParserInit() {
	staticData := &febmintempParserStaticData
	staticData.once.Do(febmintempParserInit)
}

// NewFebMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFebMintempParser(input antlr.TokenStream) *FebMintempParser {
	FebMintempParserInit()
	this := new(FebMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &febmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FebMintemp.g4"

	return this
}

// FebMintempParser tokens.
const (
	FebMintempParserEOF        = antlr.TokenEOF
	FebMintempParserFEBMINTEMP = 1
	FebMintempParserOWNKEY     = 2
	FebMintempParserSPLITKEY   = 3
	FebMintempParserWS         = 4
)

// FebMintempParser rules.
const (
	FebMintempParserRULE_expression = 0
	FebMintempParserRULE_febmintemp = 1
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
	p.RuleIndex = FebMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FebMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Febmintemp() IFebmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFebmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFebmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FebMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FebMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FebMintempParserRULE_expression)

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
		p.Febmintemp()
	}
	{
		p.SetState(5)
		p.Match(FebMintempParserEOF)
	}

	return localctx
}

// IFebmintempContext is an interface to support dynamic dispatch.
type IFebmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFebmintempContext differentiates from other interfaces.
	IsFebmintempContext()
}

type FebmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFebmintempContext() *FebmintempContext {
	var p = new(FebmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FebMintempParserRULE_febmintemp
	return p
}

func (*FebmintempContext) IsFebmintempContext() {}

func NewFebmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FebmintempContext {
	var p = new(FebmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FebMintempParserRULE_febmintemp

	return p
}

func (s *FebmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *FebmintempContext) FEBMINTEMP() antlr.TerminalNode {
	return s.GetToken(FebMintempParserFEBMINTEMP, 0)
}

func (s *FebmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FebmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FebmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMintempListener); ok {
		listenerT.EnterFebmintemp(s)
	}
}

func (s *FebmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMintempListener); ok {
		listenerT.ExitFebmintemp(s)
	}
}

func (p *FebMintempParser) Febmintemp() (localctx IFebmintempContext) {
	this := p
	_ = this

	localctx = NewFebmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FebMintempParserRULE_febmintemp)

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
		p.Match(FebMintempParserFEBMINTEMP)
	}

	return localctx
}
