// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675317863660/Memory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Memory

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

type MemoryParser struct {
	*antlr.BaseParser
}

var memoryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func memoryParserInit() {
	staticData := &memoryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MEMORY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "memory",
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

// MemoryParserInit initializes any static state used to implement MemoryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMemoryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MemoryParserInit() {
	staticData := &memoryParserStaticData
	staticData.once.Do(memoryParserInit)
}

// NewMemoryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMemoryParser(input antlr.TokenStream) *MemoryParser {
	MemoryParserInit()
	this := new(MemoryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &memoryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Memory.g4"

	return this
}

// MemoryParser tokens.
const (
	MemoryParserEOF      = antlr.TokenEOF
	MemoryParserMEMORY   = 1
	MemoryParserOWNKEY   = 2
	MemoryParserSPLITKEY = 3
	MemoryParserWS       = 4
)

// MemoryParser rules.
const (
	MemoryParserRULE_expression = 0
	MemoryParserRULE_memory     = 1
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
	p.RuleIndex = MemoryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MemoryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Memory() IMemoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMemoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMemoryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MemoryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MemoryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MemoryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MemoryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MemoryParserRULE_expression)

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
		p.Memory()
	}
	{
		p.SetState(5)
		p.Match(MemoryParserEOF)
	}

	return localctx
}

// IMemoryContext is an interface to support dynamic dispatch.
type IMemoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemoryContext differentiates from other interfaces.
	IsMemoryContext()
}

type MemoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemoryContext() *MemoryContext {
	var p = new(MemoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MemoryParserRULE_memory
	return p
}

func (*MemoryContext) IsMemoryContext() {}

func NewMemoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemoryContext {
	var p = new(MemoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MemoryParserRULE_memory

	return p
}

func (s *MemoryContext) GetParser() antlr.Parser { return s.parser }

func (s *MemoryContext) MEMORY() antlr.TerminalNode {
	return s.GetToken(MemoryParserMEMORY, 0)
}

func (s *MemoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MemoryListener); ok {
		listenerT.EnterMemory(s)
	}
}

func (s *MemoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MemoryListener); ok {
		listenerT.ExitMemory(s)
	}
}

func (p *MemoryParser) Memory() (localctx IMemoryContext) {
	this := p
	_ = this

	localctx = NewMemoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MemoryParserRULE_memory)

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
		p.Match(MemoryParserMEMORY)
	}

	return localctx
}
