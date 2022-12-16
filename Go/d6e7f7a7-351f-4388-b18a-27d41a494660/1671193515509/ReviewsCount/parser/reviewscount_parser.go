// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193515509/ReviewsCount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReviewsCount

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

type ReviewsCountParser struct {
	*antlr.BaseParser
}

var reviewscountParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func reviewscountParserInit() {
	staticData := &reviewscountParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "REVIEWSCOUNT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "reviewscount",
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

// ReviewsCountParserInit initializes any static state used to implement ReviewsCountParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewReviewsCountParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ReviewsCountParserInit() {
	staticData := &reviewscountParserStaticData
	staticData.once.Do(reviewscountParserInit)
}

// NewReviewsCountParser produces a new parser instance for the optional input antlr.TokenStream.
func NewReviewsCountParser(input antlr.TokenStream) *ReviewsCountParser {
	ReviewsCountParserInit()
	this := new(ReviewsCountParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &reviewscountParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ReviewsCount.g4"

	return this
}

// ReviewsCountParser tokens.
const (
	ReviewsCountParserEOF          = antlr.TokenEOF
	ReviewsCountParserREVIEWSCOUNT = 1
	ReviewsCountParserOWNKEY       = 2
	ReviewsCountParserSPLITKEY     = 3
	ReviewsCountParserWS           = 4
)

// ReviewsCountParser rules.
const (
	ReviewsCountParserRULE_expression   = 0
	ReviewsCountParserRULE_reviewscount = 1
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
	p.RuleIndex = ReviewsCountParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReviewsCountParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Reviewscount() IReviewscountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReviewscountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReviewscountContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ReviewsCountParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewsCountListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewsCountListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ReviewsCountParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReviewsCountParserRULE_expression)

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
		p.Reviewscount()
	}
	{
		p.SetState(5)
		p.Match(ReviewsCountParserEOF)
	}

	return localctx
}

// IReviewscountContext is an interface to support dynamic dispatch.
type IReviewscountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReviewscountContext differentiates from other interfaces.
	IsReviewscountContext()
}

type ReviewscountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReviewscountContext() *ReviewscountContext {
	var p = new(ReviewscountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReviewsCountParserRULE_reviewscount
	return p
}

func (*ReviewscountContext) IsReviewscountContext() {}

func NewReviewscountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReviewscountContext {
	var p = new(ReviewscountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReviewsCountParserRULE_reviewscount

	return p
}

func (s *ReviewscountContext) GetParser() antlr.Parser { return s.parser }

func (s *ReviewscountContext) REVIEWSCOUNT() antlr.TerminalNode {
	return s.GetToken(ReviewsCountParserREVIEWSCOUNT, 0)
}

func (s *ReviewscountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReviewscountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReviewscountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewsCountListener); ok {
		listenerT.EnterReviewscount(s)
	}
}

func (s *ReviewscountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewsCountListener); ok {
		listenerT.ExitReviewscount(s)
	}
}

func (p *ReviewsCountParser) Reviewscount() (localctx IReviewscountContext) {
	this := p
	_ = this

	localctx = NewReviewscountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReviewsCountParserRULE_reviewscount)

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
		p.Match(ReviewsCountParserREVIEWSCOUNT)
	}

	return localctx
}
