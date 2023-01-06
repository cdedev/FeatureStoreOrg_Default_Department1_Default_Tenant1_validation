// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979299463/Fm1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fm1

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

type Fm1Parser struct {
	*antlr.BaseParser
}

var fm1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fm1ParserInit() {
	staticData := &fm1ParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FM1", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "fm1",
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

// Fm1ParserInit initializes any static state used to implement Fm1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFm1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Fm1ParserInit() {
	staticData := &fm1ParserStaticData
	staticData.once.Do(fm1ParserInit)
}

// NewFm1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewFm1Parser(input antlr.TokenStream) *Fm1Parser {
	Fm1ParserInit()
	this := new(Fm1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fm1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Fm1.g4"

	return this
}

// Fm1Parser tokens.
const (
	Fm1ParserEOF      = antlr.TokenEOF
	Fm1ParserFM1      = 1
	Fm1ParserOWNKEY   = 2
	Fm1ParserSPLITKEY = 3
	Fm1ParserWS       = 4
)

// Fm1Parser rules.
const (
	Fm1ParserRULE_expression = 0
	Fm1ParserRULE_fm1        = 1
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
	p.RuleIndex = Fm1ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Fm1ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Fm1() IFm1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFm1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFm1Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Fm1ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Fm1Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Fm1Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Fm1Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Fm1ParserRULE_expression)

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
		p.Fm1()
	}
	{
		p.SetState(5)
		p.Match(Fm1ParserEOF)
	}

	return localctx
}

// IFm1Context is an interface to support dynamic dispatch.
type IFm1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFm1Context differentiates from other interfaces.
	IsFm1Context()
}

type Fm1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFm1Context() *Fm1Context {
	var p = new(Fm1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Fm1ParserRULE_fm1
	return p
}

func (*Fm1Context) IsFm1Context() {}

func NewFm1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fm1Context {
	var p = new(Fm1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Fm1ParserRULE_fm1

	return p
}

func (s *Fm1Context) GetParser() antlr.Parser { return s.parser }

func (s *Fm1Context) FM1() antlr.TerminalNode {
	return s.GetToken(Fm1ParserFM1, 0)
}

func (s *Fm1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fm1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fm1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Fm1Listener); ok {
		listenerT.EnterFm1(s)
	}
}

func (s *Fm1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Fm1Listener); ok {
		listenerT.ExitFm1(s)
	}
}

func (p *Fm1Parser) Fm1() (localctx IFm1Context) {
	this := p
	_ = this

	localctx = NewFm1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Fm1ParserRULE_fm1)

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
		p.Match(Fm1ParserFM1)
	}

	return localctx
}
