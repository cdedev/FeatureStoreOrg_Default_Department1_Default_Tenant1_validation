// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780151851/Average.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Average

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

type AverageParser struct {
	*antlr.BaseParser
}

var averageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func averageParserInit() {
	staticData := &averageParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AVERAGE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "average",
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

// AverageParserInit initializes any static state used to implement AverageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAverageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AverageParserInit() {
	staticData := &averageParserStaticData
	staticData.once.Do(averageParserInit)
}

// NewAverageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAverageParser(input antlr.TokenStream) *AverageParser {
	AverageParserInit()
	this := new(AverageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &averageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Average.g4"

	return this
}

// AverageParser tokens.
const (
	AverageParserEOF      = antlr.TokenEOF
	AverageParserAVERAGE  = 1
	AverageParserOWNKEY   = 2
	AverageParserSPLITKEY = 3
	AverageParserWS       = 4
)

// AverageParser rules.
const (
	AverageParserRULE_expression = 0
	AverageParserRULE_average    = 1
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
	p.RuleIndex = AverageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AverageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Average() IAverageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAverageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAverageContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AverageParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AverageParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AverageParserRULE_expression)

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
		p.Average()
	}
	{
		p.SetState(5)
		p.Match(AverageParserEOF)
	}

	return localctx
}

// IAverageContext is an interface to support dynamic dispatch.
type IAverageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAverageContext differentiates from other interfaces.
	IsAverageContext()
}

type AverageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAverageContext() *AverageContext {
	var p = new(AverageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AverageParserRULE_average
	return p
}

func (*AverageContext) IsAverageContext() {}

func NewAverageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AverageContext {
	var p = new(AverageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AverageParserRULE_average

	return p
}

func (s *AverageContext) GetParser() antlr.Parser { return s.parser }

func (s *AverageContext) AVERAGE() antlr.TerminalNode {
	return s.GetToken(AverageParserAVERAGE, 0)
}

func (s *AverageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AverageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AverageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageListener); ok {
		listenerT.EnterAverage(s)
	}
}

func (s *AverageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageListener); ok {
		listenerT.ExitAverage(s)
	}
}

func (p *AverageParser) Average() (localctx IAverageContext) {
	this := p
	_ = this

	localctx = NewAverageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AverageParserRULE_average)

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
		p.Match(AverageParserAVERAGE)
	}

	return localctx
}
