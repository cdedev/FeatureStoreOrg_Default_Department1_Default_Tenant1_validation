// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675672391602/Underweight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Underweight

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

type UnderweightParser struct {
	*antlr.BaseParser
}

var underweightParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func underweightParserInit() {
	staticData := &underweightParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "UNDERWEIGHT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "underweight",
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

// UnderweightParserInit initializes any static state used to implement UnderweightParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewUnderweightParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func UnderweightParserInit() {
	staticData := &underweightParserStaticData
	staticData.once.Do(underweightParserInit)
}

// NewUnderweightParser produces a new parser instance for the optional input antlr.TokenStream.
func NewUnderweightParser(input antlr.TokenStream) *UnderweightParser {
	UnderweightParserInit()
	this := new(UnderweightParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &underweightParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Underweight.g4"

	return this
}

// UnderweightParser tokens.
const (
	UnderweightParserEOF         = antlr.TokenEOF
	UnderweightParserUNDERWEIGHT = 1
	UnderweightParserOWNKEY      = 2
	UnderweightParserSPLITKEY    = 3
	UnderweightParserWS          = 4
)

// UnderweightParser rules.
const (
	UnderweightParserRULE_expression  = 0
	UnderweightParserRULE_underweight = 1
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
	p.RuleIndex = UnderweightParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnderweightParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Underweight() IUnderweightContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnderweightContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnderweightContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(UnderweightParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnderweightListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnderweightListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *UnderweightParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UnderweightParserRULE_expression)

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
		p.Underweight()
	}
	{
		p.SetState(5)
		p.Match(UnderweightParserEOF)
	}

	return localctx
}

// IUnderweightContext is an interface to support dynamic dispatch.
type IUnderweightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnderweightContext differentiates from other interfaces.
	IsUnderweightContext()
}

type UnderweightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnderweightContext() *UnderweightContext {
	var p = new(UnderweightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UnderweightParserRULE_underweight
	return p
}

func (*UnderweightContext) IsUnderweightContext() {}

func NewUnderweightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnderweightContext {
	var p = new(UnderweightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UnderweightParserRULE_underweight

	return p
}

func (s *UnderweightContext) GetParser() antlr.Parser { return s.parser }

func (s *UnderweightContext) UNDERWEIGHT() antlr.TerminalNode {
	return s.GetToken(UnderweightParserUNDERWEIGHT, 0)
}

func (s *UnderweightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnderweightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnderweightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnderweightListener); ok {
		listenerT.EnterUnderweight(s)
	}
}

func (s *UnderweightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnderweightListener); ok {
		listenerT.ExitUnderweight(s)
	}
}

func (p *UnderweightParser) Underweight() (localctx IUnderweightContext) {
	this := p
	_ = this

	localctx = NewUnderweightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UnderweightParserRULE_underweight)

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
		p.Match(UnderweightParserUNDERWEIGHT)
	}

	return localctx
}
