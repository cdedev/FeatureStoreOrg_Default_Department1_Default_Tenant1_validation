// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979442491/Tp1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tp1

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

type Tp1Parser struct {
	*antlr.BaseParser
}

var tp1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tp1ParserInit() {
	staticData := &tp1ParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TP1", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "tp1",
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

// Tp1ParserInit initializes any static state used to implement Tp1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTp1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Tp1ParserInit() {
	staticData := &tp1ParserStaticData
	staticData.once.Do(tp1ParserInit)
}

// NewTp1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewTp1Parser(input antlr.TokenStream) *Tp1Parser {
	Tp1ParserInit()
	this := new(Tp1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tp1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tp1.g4"

	return this
}

// Tp1Parser tokens.
const (
	Tp1ParserEOF      = antlr.TokenEOF
	Tp1ParserTP1      = 1
	Tp1ParserOWNKEY   = 2
	Tp1ParserSPLITKEY = 3
	Tp1ParserWS       = 4
)

// Tp1Parser rules.
const (
	Tp1ParserRULE_expression = 0
	Tp1ParserRULE_tp1        = 1
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
	p.RuleIndex = Tp1ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tp1ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tp1() ITp1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITp1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITp1Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Tp1ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tp1Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tp1Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Tp1Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Tp1ParserRULE_expression)

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
		p.Tp1()
	}
	{
		p.SetState(5)
		p.Match(Tp1ParserEOF)
	}

	return localctx
}

// ITp1Context is an interface to support dynamic dispatch.
type ITp1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTp1Context differentiates from other interfaces.
	IsTp1Context()
}

type Tp1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTp1Context() *Tp1Context {
	var p = new(Tp1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Tp1ParserRULE_tp1
	return p
}

func (*Tp1Context) IsTp1Context() {}

func NewTp1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tp1Context {
	var p = new(Tp1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tp1ParserRULE_tp1

	return p
}

func (s *Tp1Context) GetParser() antlr.Parser { return s.parser }

func (s *Tp1Context) TP1() antlr.TerminalNode {
	return s.GetToken(Tp1ParserTP1, 0)
}

func (s *Tp1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tp1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tp1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tp1Listener); ok {
		listenerT.EnterTp1(s)
	}
}

func (s *Tp1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tp1Listener); ok {
		listenerT.ExitTp1(s)
	}
}

func (p *Tp1Parser) Tp1() (localctx ITp1Context) {
	this := p
	_ = this

	localctx = NewTp1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Tp1ParserRULE_tp1)

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
		p.Match(Tp1ParserTP1)
	}

	return localctx
}
