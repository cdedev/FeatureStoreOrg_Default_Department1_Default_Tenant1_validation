// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671770834489/AbsoulteMagnitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AbsoulteMagnitude

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

type AbsoulteMagnitudeParser struct {
	*antlr.BaseParser
}

var absoultemagnitudeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func absoultemagnitudeParserInit() {
	staticData := &absoultemagnitudeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ABSOULTEMAGNITUDE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "absoultemagnitude",
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

// AbsoulteMagnitudeParserInit initializes any static state used to implement AbsoulteMagnitudeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAbsoulteMagnitudeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AbsoulteMagnitudeParserInit() {
	staticData := &absoultemagnitudeParserStaticData
	staticData.once.Do(absoultemagnitudeParserInit)
}

// NewAbsoulteMagnitudeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAbsoulteMagnitudeParser(input antlr.TokenStream) *AbsoulteMagnitudeParser {
	AbsoulteMagnitudeParserInit()
	this := new(AbsoulteMagnitudeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &absoultemagnitudeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AbsoulteMagnitude.g4"

	return this
}

// AbsoulteMagnitudeParser tokens.
const (
	AbsoulteMagnitudeParserEOF               = antlr.TokenEOF
	AbsoulteMagnitudeParserABSOULTEMAGNITUDE = 1
	AbsoulteMagnitudeParserOWNKEY            = 2
	AbsoulteMagnitudeParserSPLITKEY          = 3
	AbsoulteMagnitudeParserWS                = 4
)

// AbsoulteMagnitudeParser rules.
const (
	AbsoulteMagnitudeParserRULE_expression        = 0
	AbsoulteMagnitudeParserRULE_absoultemagnitude = 1
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
	p.RuleIndex = AbsoulteMagnitudeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbsoulteMagnitudeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Absoultemagnitude() IAbsoultemagnitudeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAbsoultemagnitudeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAbsoultemagnitudeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AbsoulteMagnitudeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbsoulteMagnitudeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbsoulteMagnitudeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AbsoulteMagnitudeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AbsoulteMagnitudeParserRULE_expression)

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
		p.Absoultemagnitude()
	}
	{
		p.SetState(5)
		p.Match(AbsoulteMagnitudeParserEOF)
	}

	return localctx
}

// IAbsoultemagnitudeContext is an interface to support dynamic dispatch.
type IAbsoultemagnitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAbsoultemagnitudeContext differentiates from other interfaces.
	IsAbsoultemagnitudeContext()
}

type AbsoultemagnitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAbsoultemagnitudeContext() *AbsoultemagnitudeContext {
	var p = new(AbsoultemagnitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AbsoulteMagnitudeParserRULE_absoultemagnitude
	return p
}

func (*AbsoultemagnitudeContext) IsAbsoultemagnitudeContext() {}

func NewAbsoultemagnitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AbsoultemagnitudeContext {
	var p = new(AbsoultemagnitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AbsoulteMagnitudeParserRULE_absoultemagnitude

	return p
}

func (s *AbsoultemagnitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *AbsoultemagnitudeContext) ABSOULTEMAGNITUDE() antlr.TerminalNode {
	return s.GetToken(AbsoulteMagnitudeParserABSOULTEMAGNITUDE, 0)
}

func (s *AbsoultemagnitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsoultemagnitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AbsoultemagnitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbsoulteMagnitudeListener); ok {
		listenerT.EnterAbsoultemagnitude(s)
	}
}

func (s *AbsoultemagnitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AbsoulteMagnitudeListener); ok {
		listenerT.ExitAbsoultemagnitude(s)
	}
}

func (p *AbsoulteMagnitudeParser) Absoultemagnitude() (localctx IAbsoultemagnitudeContext) {
	this := p
	_ = this

	localctx = NewAbsoultemagnitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AbsoulteMagnitudeParserRULE_absoultemagnitude)

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
		p.Match(AbsoulteMagnitudeParserABSOULTEMAGNITUDE)
	}

	return localctx
}
