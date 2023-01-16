// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867803590/Metal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Metal

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

type MetalParser struct {
	*antlr.BaseParser
}

var metalParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func metalParserInit() {
	staticData := &metalParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "METAL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "metal",
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

// MetalParserInit initializes any static state used to implement MetalParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMetalParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MetalParserInit() {
	staticData := &metalParserStaticData
	staticData.once.Do(metalParserInit)
}

// NewMetalParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMetalParser(input antlr.TokenStream) *MetalParser {
	MetalParserInit()
	this := new(MetalParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &metalParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Metal.g4"

	return this
}

// MetalParser tokens.
const (
	MetalParserEOF      = antlr.TokenEOF
	MetalParserMETAL    = 1
	MetalParserOWNKEY   = 2
	MetalParserSPLITKEY = 3
	MetalParserWS       = 4
)

// MetalParser rules.
const (
	MetalParserRULE_expression = 0
	MetalParserRULE_metal      = 1
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
	p.RuleIndex = MetalParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetalParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Metal() IMetalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetalContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MetalParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetalListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetalListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MetalParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MetalParserRULE_expression)

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
		p.Metal()
	}
	{
		p.SetState(5)
		p.Match(MetalParserEOF)
	}

	return localctx
}

// IMetalContext is an interface to support dynamic dispatch.
type IMetalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetalContext differentiates from other interfaces.
	IsMetalContext()
}

type MetalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetalContext() *MetalContext {
	var p = new(MetalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetalParserRULE_metal
	return p
}

func (*MetalContext) IsMetalContext() {}

func NewMetalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetalContext {
	var p = new(MetalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetalParserRULE_metal

	return p
}

func (s *MetalContext) GetParser() antlr.Parser { return s.parser }

func (s *MetalContext) METAL() antlr.TerminalNode {
	return s.GetToken(MetalParserMETAL, 0)
}

func (s *MetalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetalListener); ok {
		listenerT.EnterMetal(s)
	}
}

func (s *MetalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetalListener); ok {
		listenerT.ExitMetal(s)
	}
}

func (p *MetalParser) Metal() (localctx IMetalContext) {
	this := p
	_ = this

	localctx = NewMetalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MetalParserRULE_metal)

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
		p.Match(MetalParserMETAL)
	}

	return localctx
}
