// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377270134/IntMemory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntMemory

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

type IntMemoryParser struct {
	*antlr.BaseParser
}

var intmemoryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func intmemoryParserInit() {
	staticData := &intmemoryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INTMEMORY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "intmemory",
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

// IntMemoryParserInit initializes any static state used to implement IntMemoryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIntMemoryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IntMemoryParserInit() {
	staticData := &intmemoryParserStaticData
	staticData.once.Do(intmemoryParserInit)
}

// NewIntMemoryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIntMemoryParser(input antlr.TokenStream) *IntMemoryParser {
	IntMemoryParserInit()
	this := new(IntMemoryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &intmemoryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "IntMemory.g4"

	return this
}

// IntMemoryParser tokens.
const (
	IntMemoryParserEOF       = antlr.TokenEOF
	IntMemoryParserINTMEMORY = 1
	IntMemoryParserOWNKEY    = 2
	IntMemoryParserSPLITKEY  = 3
	IntMemoryParserWS        = 4
)

// IntMemoryParser rules.
const (
	IntMemoryParserRULE_expression = 0
	IntMemoryParserRULE_intmemory  = 1
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
	p.RuleIndex = IntMemoryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntMemoryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Intmemory() IIntmemoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntmemoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntmemoryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IntMemoryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntMemoryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntMemoryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IntMemoryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IntMemoryParserRULE_expression)

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
		p.Intmemory()
	}
	{
		p.SetState(5)
		p.Match(IntMemoryParserEOF)
	}

	return localctx
}

// IIntmemoryContext is an interface to support dynamic dispatch.
type IIntmemoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntmemoryContext differentiates from other interfaces.
	IsIntmemoryContext()
}

type IntmemoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntmemoryContext() *IntmemoryContext {
	var p = new(IntmemoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IntMemoryParserRULE_intmemory
	return p
}

func (*IntmemoryContext) IsIntmemoryContext() {}

func NewIntmemoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntmemoryContext {
	var p = new(IntmemoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntMemoryParserRULE_intmemory

	return p
}

func (s *IntmemoryContext) GetParser() antlr.Parser { return s.parser }

func (s *IntmemoryContext) INTMEMORY() antlr.TerminalNode {
	return s.GetToken(IntMemoryParserINTMEMORY, 0)
}

func (s *IntmemoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntmemoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntmemoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntMemoryListener); ok {
		listenerT.EnterIntmemory(s)
	}
}

func (s *IntmemoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntMemoryListener); ok {
		listenerT.ExitIntmemory(s)
	}
}

func (p *IntMemoryParser) Intmemory() (localctx IIntmemoryContext) {
	this := p
	_ = this

	localctx = NewIntmemoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IntMemoryParserRULE_intmemory)

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
		p.Match(IntMemoryParserINTMEMORY)
	}

	return localctx
}
