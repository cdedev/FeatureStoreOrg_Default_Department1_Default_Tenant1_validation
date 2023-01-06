// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979372969/Pe1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pe1

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

type Pe1Parser struct {
	*antlr.BaseParser
}

var pe1ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pe1ParserInit() {
	staticData := &pe1ParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PE1", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pe1",
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

// Pe1ParserInit initializes any static state used to implement Pe1Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPe1Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Pe1ParserInit() {
	staticData := &pe1ParserStaticData
	staticData.once.Do(pe1ParserInit)
}

// NewPe1Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewPe1Parser(input antlr.TokenStream) *Pe1Parser {
	Pe1ParserInit()
	this := new(Pe1Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pe1ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Pe1.g4"

	return this
}

// Pe1Parser tokens.
const (
	Pe1ParserEOF      = antlr.TokenEOF
	Pe1ParserPE1      = 1
	Pe1ParserOWNKEY   = 2
	Pe1ParserSPLITKEY = 3
	Pe1ParserWS       = 4
)

// Pe1Parser rules.
const (
	Pe1ParserRULE_expression = 0
	Pe1ParserRULE_pe1        = 1
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
	p.RuleIndex = Pe1ParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Pe1ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pe1() IPe1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPe1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPe1Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Pe1ParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pe1Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pe1Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Pe1Parser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Pe1ParserRULE_expression)

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
		p.Pe1()
	}
	{
		p.SetState(5)
		p.Match(Pe1ParserEOF)
	}

	return localctx
}

// IPe1Context is an interface to support dynamic dispatch.
type IPe1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPe1Context differentiates from other interfaces.
	IsPe1Context()
}

type Pe1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPe1Context() *Pe1Context {
	var p = new(Pe1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Pe1ParserRULE_pe1
	return p
}

func (*Pe1Context) IsPe1Context() {}

func NewPe1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pe1Context {
	var p = new(Pe1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Pe1ParserRULE_pe1

	return p
}

func (s *Pe1Context) GetParser() antlr.Parser { return s.parser }

func (s *Pe1Context) PE1() antlr.TerminalNode {
	return s.GetToken(Pe1ParserPE1, 0)
}

func (s *Pe1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pe1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pe1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pe1Listener); ok {
		listenerT.EnterPe1(s)
	}
}

func (s *Pe1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pe1Listener); ok {
		listenerT.ExitPe1(s)
	}
}

func (p *Pe1Parser) Pe1() (localctx IPe1Context) {
	this := p
	_ = this

	localctx = NewPe1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Pe1ParserRULE_pe1)

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
		p.Match(Pe1ParserPE1)
	}

	return localctx
}
