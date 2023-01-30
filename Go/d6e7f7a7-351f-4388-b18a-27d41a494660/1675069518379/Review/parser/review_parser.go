// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069518379/Review.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Review

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

type ReviewParser struct {
	*antlr.BaseParser
}

var reviewParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func reviewParserInit() {
	staticData := &reviewParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "REVIEW", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "review",
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

// ReviewParserInit initializes any static state used to implement ReviewParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewReviewParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ReviewParserInit() {
	staticData := &reviewParserStaticData
	staticData.once.Do(reviewParserInit)
}

// NewReviewParser produces a new parser instance for the optional input antlr.TokenStream.
func NewReviewParser(input antlr.TokenStream) *ReviewParser {
	ReviewParserInit()
	this := new(ReviewParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &reviewParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Review.g4"

	return this
}

// ReviewParser tokens.
const (
	ReviewParserEOF      = antlr.TokenEOF
	ReviewParserREVIEW   = 1
	ReviewParserOWNKEY   = 2
	ReviewParserSPLITKEY = 3
	ReviewParserWS       = 4
)

// ReviewParser rules.
const (
	ReviewParserRULE_expression = 0
	ReviewParserRULE_review     = 1
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
	p.RuleIndex = ReviewParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReviewParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Review() IReviewContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReviewContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReviewContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ReviewParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ReviewParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReviewParserRULE_expression)

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
		p.Review()
	}
	{
		p.SetState(5)
		p.Match(ReviewParserEOF)
	}

	return localctx
}

// IReviewContext is an interface to support dynamic dispatch.
type IReviewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReviewContext differentiates from other interfaces.
	IsReviewContext()
}

type ReviewContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReviewContext() *ReviewContext {
	var p = new(ReviewContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReviewParserRULE_review
	return p
}

func (*ReviewContext) IsReviewContext() {}

func NewReviewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReviewContext {
	var p = new(ReviewContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReviewParserRULE_review

	return p
}

func (s *ReviewContext) GetParser() antlr.Parser { return s.parser }

func (s *ReviewContext) REVIEW() antlr.TerminalNode {
	return s.GetToken(ReviewParserREVIEW, 0)
}

func (s *ReviewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReviewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReviewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewListener); ok {
		listenerT.EnterReview(s)
	}
}

func (s *ReviewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReviewListener); ok {
		listenerT.ExitReview(s)
	}
}

func (p *ReviewParser) Review() (localctx IReviewContext) {
	this := p
	_ = this

	localctx = NewReviewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReviewParserRULE_review)

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
		p.Match(ReviewParserREVIEW)
	}

	return localctx
}
