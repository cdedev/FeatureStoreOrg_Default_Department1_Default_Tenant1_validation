// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669694040966/Rank.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rank

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

type RankParser struct {
	*antlr.BaseParser
}

var rankParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rankParserInit() {
	staticData := &rankParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RANK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rank",
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

// RankParserInit initializes any static state used to implement RankParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRankParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RankParserInit() {
	staticData := &rankParserStaticData
	staticData.once.Do(rankParserInit)
}

// NewRankParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRankParser(input antlr.TokenStream) *RankParser {
	RankParserInit()
	this := new(RankParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rankParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rank.g4"

	return this
}

// RankParser tokens.
const (
	RankParserEOF      = antlr.TokenEOF
	RankParserRANK     = 1
	RankParserOWNKEY   = 2
	RankParserSPLITKEY = 3
	RankParserWS       = 4
)

// RankParser rules.
const (
	RankParserRULE_expression = 0
	RankParserRULE_rank       = 1
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
	p.RuleIndex = RankParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RankParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rank() IRankContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRankContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRankContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RankParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RankListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RankListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RankParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RankParserRULE_expression)

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
		p.Rank()
	}
	{
		p.SetState(5)
		p.Match(RankParserEOF)
	}

	return localctx
}

// IRankContext is an interface to support dynamic dispatch.
type IRankContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRankContext differentiates from other interfaces.
	IsRankContext()
}

type RankContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRankContext() *RankContext {
	var p = new(RankContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RankParserRULE_rank
	return p
}

func (*RankContext) IsRankContext() {}

func NewRankContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RankContext {
	var p = new(RankContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RankParserRULE_rank

	return p
}

func (s *RankContext) GetParser() antlr.Parser { return s.parser }

func (s *RankContext) RANK() antlr.TerminalNode {
	return s.GetToken(RankParserRANK, 0)
}

func (s *RankContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RankContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RankContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RankListener); ok {
		listenerT.EnterRank(s)
	}
}

func (s *RankContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RankListener); ok {
		listenerT.ExitRank(s)
	}
}

func (p *RankParser) Rank() (localctx IRankContext) {
	this := p
	_ = this

	localctx = NewRankContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RankParserRULE_rank)

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
		p.Match(RankParserRANK)
	}

	return localctx
}
