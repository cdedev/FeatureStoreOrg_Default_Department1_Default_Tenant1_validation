// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671689783635/Skew.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Skew

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

type SkewParser struct {
	*antlr.BaseParser
}

var skewParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func skewParserInit() {
	staticData := &skewParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SKEW", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "skew",
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

// SkewParserInit initializes any static state used to implement SkewParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSkewParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SkewParserInit() {
	staticData := &skewParserStaticData
	staticData.once.Do(skewParserInit)
}

// NewSkewParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSkewParser(input antlr.TokenStream) *SkewParser {
	SkewParserInit()
	this := new(SkewParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &skewParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Skew.g4"

	return this
}

// SkewParser tokens.
const (
	SkewParserEOF      = antlr.TokenEOF
	SkewParserSKEW     = 1
	SkewParserOWNKEY   = 2
	SkewParserSPLITKEY = 3
	SkewParserWS       = 4
)

// SkewParser rules.
const (
	SkewParserRULE_expression = 0
	SkewParserRULE_skew       = 1
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
	p.RuleIndex = SkewParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkewParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Skew() ISkewContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISkewContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISkewContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SkewParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkewListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkewListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SkewParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SkewParserRULE_expression)

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
		p.Skew()
	}
	{
		p.SetState(5)
		p.Match(SkewParserEOF)
	}

	return localctx
}

// ISkewContext is an interface to support dynamic dispatch.
type ISkewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSkewContext differentiates from other interfaces.
	IsSkewContext()
}

type SkewContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySkewContext() *SkewContext {
	var p = new(SkewContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SkewParserRULE_skew
	return p
}

func (*SkewContext) IsSkewContext() {}

func NewSkewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SkewContext {
	var p = new(SkewContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SkewParserRULE_skew

	return p
}

func (s *SkewContext) GetParser() antlr.Parser { return s.parser }

func (s *SkewContext) SKEW() antlr.TerminalNode {
	return s.GetToken(SkewParserSKEW, 0)
}

func (s *SkewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SkewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkewListener); ok {
		listenerT.EnterSkew(s)
	}
}

func (s *SkewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SkewListener); ok {
		listenerT.ExitSkew(s)
	}
}

func (p *SkewParser) Skew() (localctx ISkewContext) {
	this := p
	_ = this

	localctx = NewSkewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SkewParserRULE_skew)

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
		p.Match(SkewParserSKEW)
	}

	return localctx
}
