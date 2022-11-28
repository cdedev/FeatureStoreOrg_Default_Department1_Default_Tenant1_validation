// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669636205938/Rating.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rating

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

type RatingParser struct {
	*antlr.BaseParser
}

var ratingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ratingParserInit() {
	staticData := &ratingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RATING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rating",
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

// RatingParserInit initializes any static state used to implement RatingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRatingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RatingParserInit() {
	staticData := &ratingParserStaticData
	staticData.once.Do(ratingParserInit)
}

// NewRatingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRatingParser(input antlr.TokenStream) *RatingParser {
	RatingParserInit()
	this := new(RatingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ratingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rating.g4"

	return this
}

// RatingParser tokens.
const (
	RatingParserEOF      = antlr.TokenEOF
	RatingParserRATING   = 1
	RatingParserOWNKEY   = 2
	RatingParserSPLITKEY = 3
	RatingParserWS       = 4
)

// RatingParser rules.
const (
	RatingParserRULE_expression = 0
	RatingParserRULE_rating     = 1
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
	p.RuleIndex = RatingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RatingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rating() IRatingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRatingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRatingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RatingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RatingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RatingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RatingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RatingParserRULE_expression)

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
		p.Rating()
	}
	{
		p.SetState(5)
		p.Match(RatingParserEOF)
	}

	return localctx
}

// IRatingContext is an interface to support dynamic dispatch.
type IRatingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRatingContext differentiates from other interfaces.
	IsRatingContext()
}

type RatingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRatingContext() *RatingContext {
	var p = new(RatingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RatingParserRULE_rating
	return p
}

func (*RatingContext) IsRatingContext() {}

func NewRatingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RatingContext {
	var p = new(RatingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RatingParserRULE_rating

	return p
}

func (s *RatingContext) GetParser() antlr.Parser { return s.parser }

func (s *RatingContext) RATING() antlr.TerminalNode {
	return s.GetToken(RatingParserRATING, 0)
}

func (s *RatingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RatingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RatingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RatingListener); ok {
		listenerT.EnterRating(s)
	}
}

func (s *RatingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RatingListener); ok {
		listenerT.ExitRating(s)
	}
}

func (p *RatingParser) Rating() (localctx IRatingContext) {
	this := p
	_ = this

	localctx = NewRatingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RatingParserRULE_rating)

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
		p.Match(RatingParserRATING)
	}

	return localctx
}
