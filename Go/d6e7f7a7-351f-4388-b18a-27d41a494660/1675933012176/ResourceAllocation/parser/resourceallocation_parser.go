// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675933012176/ResourceAllocation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ResourceAllocation

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

type ResourceAllocationParser struct {
	*antlr.BaseParser
}

var resourceallocationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func resourceallocationParserInit() {
	staticData := &resourceallocationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RESOURCEALLOCATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "resourceallocation",
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

// ResourceAllocationParserInit initializes any static state used to implement ResourceAllocationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewResourceAllocationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ResourceAllocationParserInit() {
	staticData := &resourceallocationParserStaticData
	staticData.once.Do(resourceallocationParserInit)
}

// NewResourceAllocationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewResourceAllocationParser(input antlr.TokenStream) *ResourceAllocationParser {
	ResourceAllocationParserInit()
	this := new(ResourceAllocationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &resourceallocationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ResourceAllocation.g4"

	return this
}

// ResourceAllocationParser tokens.
const (
	ResourceAllocationParserEOF                = antlr.TokenEOF
	ResourceAllocationParserRESOURCEALLOCATION = 1
	ResourceAllocationParserOWNKEY             = 2
	ResourceAllocationParserSPLITKEY           = 3
	ResourceAllocationParserWS                 = 4
)

// ResourceAllocationParser rules.
const (
	ResourceAllocationParserRULE_expression         = 0
	ResourceAllocationParserRULE_resourceallocation = 1
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
	p.RuleIndex = ResourceAllocationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ResourceAllocationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Resourceallocation() IResourceallocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResourceallocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResourceallocationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ResourceAllocationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResourceAllocationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResourceAllocationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ResourceAllocationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ResourceAllocationParserRULE_expression)

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
		p.Resourceallocation()
	}
	{
		p.SetState(5)
		p.Match(ResourceAllocationParserEOF)
	}

	return localctx
}

// IResourceallocationContext is an interface to support dynamic dispatch.
type IResourceallocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResourceallocationContext differentiates from other interfaces.
	IsResourceallocationContext()
}

type ResourceallocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResourceallocationContext() *ResourceallocationContext {
	var p = new(ResourceallocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ResourceAllocationParserRULE_resourceallocation
	return p
}

func (*ResourceallocationContext) IsResourceallocationContext() {}

func NewResourceallocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResourceallocationContext {
	var p = new(ResourceallocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ResourceAllocationParserRULE_resourceallocation

	return p
}

func (s *ResourceallocationContext) GetParser() antlr.Parser { return s.parser }

func (s *ResourceallocationContext) RESOURCEALLOCATION() antlr.TerminalNode {
	return s.GetToken(ResourceAllocationParserRESOURCEALLOCATION, 0)
}

func (s *ResourceallocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResourceallocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResourceallocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResourceAllocationListener); ok {
		listenerT.EnterResourceallocation(s)
	}
}

func (s *ResourceallocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResourceAllocationListener); ok {
		listenerT.ExitResourceallocation(s)
	}
}

func (p *ResourceAllocationParser) Resourceallocation() (localctx IResourceallocationContext) {
	this := p
	_ = this

	localctx = NewResourceallocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ResourceAllocationParserRULE_resourceallocation)

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
		p.Match(ResourceAllocationParserRESOURCEALLOCATION)
	}

	return localctx
}
