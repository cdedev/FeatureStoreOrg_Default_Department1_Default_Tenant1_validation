// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758417557/AvgDailyPageviews.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AvgDailyPageviews

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

type AvgDailyPageviewsParser struct {
	*antlr.BaseParser
}

var avgdailypageviewsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func avgdailypageviewsParserInit() {
	staticData := &avgdailypageviewsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AVGDAILYPAGEVIEWS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "avgdailypageviews",
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

// AvgDailyPageviewsParserInit initializes any static state used to implement AvgDailyPageviewsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAvgDailyPageviewsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AvgDailyPageviewsParserInit() {
	staticData := &avgdailypageviewsParserStaticData
	staticData.once.Do(avgdailypageviewsParserInit)
}

// NewAvgDailyPageviewsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAvgDailyPageviewsParser(input antlr.TokenStream) *AvgDailyPageviewsParser {
	AvgDailyPageviewsParserInit()
	this := new(AvgDailyPageviewsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &avgdailypageviewsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AvgDailyPageviews.g4"

	return this
}

// AvgDailyPageviewsParser tokens.
const (
	AvgDailyPageviewsParserEOF               = antlr.TokenEOF
	AvgDailyPageviewsParserAVGDAILYPAGEVIEWS = 1
	AvgDailyPageviewsParserOWNKEY            = 2
	AvgDailyPageviewsParserSPLITKEY          = 3
	AvgDailyPageviewsParserWS                = 4
)

// AvgDailyPageviewsParser rules.
const (
	AvgDailyPageviewsParserRULE_expression        = 0
	AvgDailyPageviewsParserRULE_avgdailypageviews = 1
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
	p.RuleIndex = AvgDailyPageviewsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AvgDailyPageviewsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Avgdailypageviews() IAvgdailypageviewsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAvgdailypageviewsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAvgdailypageviewsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AvgDailyPageviewsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgDailyPageviewsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgDailyPageviewsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AvgDailyPageviewsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AvgDailyPageviewsParserRULE_expression)

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
		p.Avgdailypageviews()
	}
	{
		p.SetState(5)
		p.Match(AvgDailyPageviewsParserEOF)
	}

	return localctx
}

// IAvgdailypageviewsContext is an interface to support dynamic dispatch.
type IAvgdailypageviewsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAvgdailypageviewsContext differentiates from other interfaces.
	IsAvgdailypageviewsContext()
}

type AvgdailypageviewsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvgdailypageviewsContext() *AvgdailypageviewsContext {
	var p = new(AvgdailypageviewsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AvgDailyPageviewsParserRULE_avgdailypageviews
	return p
}

func (*AvgdailypageviewsContext) IsAvgdailypageviewsContext() {}

func NewAvgdailypageviewsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvgdailypageviewsContext {
	var p = new(AvgdailypageviewsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AvgDailyPageviewsParserRULE_avgdailypageviews

	return p
}

func (s *AvgdailypageviewsContext) GetParser() antlr.Parser { return s.parser }

func (s *AvgdailypageviewsContext) AVGDAILYPAGEVIEWS() antlr.TerminalNode {
	return s.GetToken(AvgDailyPageviewsParserAVGDAILYPAGEVIEWS, 0)
}

func (s *AvgdailypageviewsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvgdailypageviewsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AvgdailypageviewsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgDailyPageviewsListener); ok {
		listenerT.EnterAvgdailypageviews(s)
	}
}

func (s *AvgdailypageviewsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgDailyPageviewsListener); ok {
		listenerT.ExitAvgdailypageviews(s)
	}
}

func (p *AvgDailyPageviewsParser) Avgdailypageviews() (localctx IAvgdailypageviewsContext) {
	this := p
	_ = this

	localctx = NewAvgdailypageviewsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AvgDailyPageviewsParserRULE_avgdailypageviews)

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
		p.Match(AvgDailyPageviewsParserAVGDAILYPAGEVIEWS)
	}

	return localctx
}
