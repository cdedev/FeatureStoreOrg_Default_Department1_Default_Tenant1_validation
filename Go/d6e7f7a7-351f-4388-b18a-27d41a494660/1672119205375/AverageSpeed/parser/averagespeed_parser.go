// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119205375/AverageSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AverageSpeed

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

type AverageSpeedParser struct {
	*antlr.BaseParser
}

var averagespeedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func averagespeedParserInit() {
	staticData := &averagespeedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AVERAGESPEED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "averagespeed",
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

// AverageSpeedParserInit initializes any static state used to implement AverageSpeedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAverageSpeedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AverageSpeedParserInit() {
	staticData := &averagespeedParserStaticData
	staticData.once.Do(averagespeedParserInit)
}

// NewAverageSpeedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAverageSpeedParser(input antlr.TokenStream) *AverageSpeedParser {
	AverageSpeedParserInit()
	this := new(AverageSpeedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &averagespeedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AverageSpeed.g4"

	return this
}

// AverageSpeedParser tokens.
const (
	AverageSpeedParserEOF          = antlr.TokenEOF
	AverageSpeedParserAVERAGESPEED = 1
	AverageSpeedParserOWNKEY       = 2
	AverageSpeedParserSPLITKEY     = 3
	AverageSpeedParserWS           = 4
)

// AverageSpeedParser rules.
const (
	AverageSpeedParserRULE_expression   = 0
	AverageSpeedParserRULE_averagespeed = 1
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
	p.RuleIndex = AverageSpeedParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AverageSpeedParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Averagespeed() IAveragespeedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAveragespeedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAveragespeedContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AverageSpeedParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageSpeedListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageSpeedListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AverageSpeedParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AverageSpeedParserRULE_expression)

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
		p.Averagespeed()
	}
	{
		p.SetState(5)
		p.Match(AverageSpeedParserEOF)
	}

	return localctx
}

// IAveragespeedContext is an interface to support dynamic dispatch.
type IAveragespeedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAveragespeedContext differentiates from other interfaces.
	IsAveragespeedContext()
}

type AveragespeedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAveragespeedContext() *AveragespeedContext {
	var p = new(AveragespeedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AverageSpeedParserRULE_averagespeed
	return p
}

func (*AveragespeedContext) IsAveragespeedContext() {}

func NewAveragespeedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AveragespeedContext {
	var p = new(AveragespeedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AverageSpeedParserRULE_averagespeed

	return p
}

func (s *AveragespeedContext) GetParser() antlr.Parser { return s.parser }

func (s *AveragespeedContext) AVERAGESPEED() antlr.TerminalNode {
	return s.GetToken(AverageSpeedParserAVERAGESPEED, 0)
}

func (s *AveragespeedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AveragespeedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AveragespeedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageSpeedListener); ok {
		listenerT.EnterAveragespeed(s)
	}
}

func (s *AveragespeedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AverageSpeedListener); ok {
		listenerT.ExitAveragespeed(s)
	}
}

func (p *AverageSpeedParser) Averagespeed() (localctx IAveragespeedContext) {
	this := p
	_ = this

	localctx = NewAveragespeedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AverageSpeedParserRULE_averagespeed)

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
		p.Match(AverageSpeedParserAVERAGESPEED)
	}

	return localctx
}
