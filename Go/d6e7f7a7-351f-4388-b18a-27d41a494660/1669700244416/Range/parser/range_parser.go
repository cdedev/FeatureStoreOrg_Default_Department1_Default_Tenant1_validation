// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700244416/Range.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Range

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

type RangeParser struct {
	*antlr.BaseParser
}

var rangeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rangeParserInit() {
	staticData := &rangeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RANGE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "range",
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

// RangeParserInit initializes any static state used to implement RangeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRangeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RangeParserInit() {
	staticData := &rangeParserStaticData
	staticData.once.Do(rangeParserInit)
}

// NewRangeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRangeParser(input antlr.TokenStream) *RangeParser {
	RangeParserInit()
	this := new(RangeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rangeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Range.g4"

	return this
}

// RangeParser tokens.
const (
	RangeParserEOF      = antlr.TokenEOF
	RangeParserRANGE    = 1
	RangeParserOWNKEY   = 2
	RangeParserSPLITKEY = 3
	RangeParserWS       = 4
)

// RangeParser rules.
const (
	RangeParserRULE_expression = 0
	RangeParserRULE_range      = 1
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
	p.RuleIndex = RangeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RangeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Range() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RangeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RangeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RangeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RangeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RangeParserRULE_expression)

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
		p.Range()
	}
	{
		p.SetState(5)
		p.Match(RangeParserEOF)
	}

	return localctx
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RangeParserRULE_range
	return p
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RangeParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) RANGE() antlr.TerminalNode {
	return s.GetToken(RangeParserRANGE, 0)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RangeListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RangeListener); ok {
		listenerT.ExitRange(s)
	}
}

func (p *RangeParser) Range() (localctx IRangeContext) {
	this := p
	_ = this

	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RangeParserRULE_range)

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
		p.Match(RangeParserRANGE)
	}

	return localctx
}
