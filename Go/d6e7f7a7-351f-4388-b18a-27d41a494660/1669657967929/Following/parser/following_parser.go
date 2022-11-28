// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657967929/Following.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Following

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

type FollowingParser struct {
	*antlr.BaseParser
}

var followingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func followingParserInit() {
	staticData := &followingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FOLLOWING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "following",
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

// FollowingParserInit initializes any static state used to implement FollowingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFollowingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FollowingParserInit() {
	staticData := &followingParserStaticData
	staticData.once.Do(followingParserInit)
}

// NewFollowingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFollowingParser(input antlr.TokenStream) *FollowingParser {
	FollowingParserInit()
	this := new(FollowingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &followingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Following.g4"

	return this
}

// FollowingParser tokens.
const (
	FollowingParserEOF       = antlr.TokenEOF
	FollowingParserFOLLOWING = 1
	FollowingParserOWNKEY    = 2
	FollowingParserSPLITKEY  = 3
	FollowingParserWS        = 4
)

// FollowingParser rules.
const (
	FollowingParserRULE_expression = 0
	FollowingParserRULE_following  = 1
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
	p.RuleIndex = FollowingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FollowingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Following() IFollowingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFollowingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFollowingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FollowingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FollowingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FollowingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FollowingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FollowingParserRULE_expression)

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
		p.Following()
	}
	{
		p.SetState(5)
		p.Match(FollowingParserEOF)
	}

	return localctx
}

// IFollowingContext is an interface to support dynamic dispatch.
type IFollowingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFollowingContext differentiates from other interfaces.
	IsFollowingContext()
}

type FollowingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFollowingContext() *FollowingContext {
	var p = new(FollowingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FollowingParserRULE_following
	return p
}

func (*FollowingContext) IsFollowingContext() {}

func NewFollowingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FollowingContext {
	var p = new(FollowingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FollowingParserRULE_following

	return p
}

func (s *FollowingContext) GetParser() antlr.Parser { return s.parser }

func (s *FollowingContext) FOLLOWING() antlr.TerminalNode {
	return s.GetToken(FollowingParserFOLLOWING, 0)
}

func (s *FollowingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FollowingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FollowingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FollowingListener); ok {
		listenerT.EnterFollowing(s)
	}
}

func (s *FollowingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FollowingListener); ok {
		listenerT.ExitFollowing(s)
	}
}

func (p *FollowingParser) Following() (localctx IFollowingContext) {
	this := p
	_ = this

	localctx = NewFollowingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FollowingParserRULE_following)

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
		p.Match(FollowingParserFOLLOWING)
	}

	return localctx
}
