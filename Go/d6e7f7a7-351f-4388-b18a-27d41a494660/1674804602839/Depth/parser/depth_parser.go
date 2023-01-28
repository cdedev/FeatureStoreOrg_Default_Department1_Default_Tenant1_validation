// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674804602839/Depth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Depth

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

type DepthParser struct {
	*antlr.BaseParser
}

var depthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func depthParserInit() {
	staticData := &depthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DEPTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "depth",
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

// DepthParserInit initializes any static state used to implement DepthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDepthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DepthParserInit() {
	staticData := &depthParserStaticData
	staticData.once.Do(depthParserInit)
}

// NewDepthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDepthParser(input antlr.TokenStream) *DepthParser {
	DepthParserInit()
	this := new(DepthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &depthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Depth.g4"

	return this
}

// DepthParser tokens.
const (
	DepthParserEOF      = antlr.TokenEOF
	DepthParserDEPTH    = 1
	DepthParserOWNKEY   = 2
	DepthParserSPLITKEY = 3
	DepthParserWS       = 4
)

// DepthParser rules.
const (
	DepthParserRULE_expression = 0
	DepthParserRULE_depth      = 1
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
	p.RuleIndex = DepthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DepthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Depth() IDepthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDepthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDepthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DepthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DepthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DepthParserRULE_expression)

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
		p.Depth()
	}
	{
		p.SetState(5)
		p.Match(DepthParserEOF)
	}

	return localctx
}

// IDepthContext is an interface to support dynamic dispatch.
type IDepthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDepthContext differentiates from other interfaces.
	IsDepthContext()
}

type DepthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDepthContext() *DepthContext {
	var p = new(DepthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DepthParserRULE_depth
	return p
}

func (*DepthContext) IsDepthContext() {}

func NewDepthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DepthContext {
	var p = new(DepthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DepthParserRULE_depth

	return p
}

func (s *DepthContext) GetParser() antlr.Parser { return s.parser }

func (s *DepthContext) DEPTH() antlr.TerminalNode {
	return s.GetToken(DepthParserDEPTH, 0)
}

func (s *DepthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DepthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DepthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepthListener); ok {
		listenerT.EnterDepth(s)
	}
}

func (s *DepthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DepthListener); ok {
		listenerT.ExitDepth(s)
	}
}

func (p *DepthParser) Depth() (localctx IDepthContext) {
	this := p
	_ = this

	localctx = NewDepthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DepthParserRULE_depth)

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
		p.Match(DepthParserDEPTH)
	}

	return localctx
}