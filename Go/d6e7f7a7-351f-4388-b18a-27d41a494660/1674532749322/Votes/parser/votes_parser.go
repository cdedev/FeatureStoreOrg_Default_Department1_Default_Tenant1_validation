// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532749322/Votes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Votes

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

type VotesParser struct {
	*antlr.BaseParser
}

var votesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func votesParserInit() {
	staticData := &votesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VOTES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "votes",
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

// VotesParserInit initializes any static state used to implement VotesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVotesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VotesParserInit() {
	staticData := &votesParserStaticData
	staticData.once.Do(votesParserInit)
}

// NewVotesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVotesParser(input antlr.TokenStream) *VotesParser {
	VotesParserInit()
	this := new(VotesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &votesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Votes.g4"

	return this
}

// VotesParser tokens.
const (
	VotesParserEOF      = antlr.TokenEOF
	VotesParserVOTES    = 1
	VotesParserOWNKEY   = 2
	VotesParserSPLITKEY = 3
	VotesParserWS       = 4
)

// VotesParser rules.
const (
	VotesParserRULE_expression = 0
	VotesParserRULE_votes      = 1
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
	p.RuleIndex = VotesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VotesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Votes() IVotesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVotesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVotesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VotesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VotesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VotesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VotesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VotesParserRULE_expression)

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
		p.Votes()
	}
	{
		p.SetState(5)
		p.Match(VotesParserEOF)
	}

	return localctx
}

// IVotesContext is an interface to support dynamic dispatch.
type IVotesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVotesContext differentiates from other interfaces.
	IsVotesContext()
}

type VotesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVotesContext() *VotesContext {
	var p = new(VotesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VotesParserRULE_votes
	return p
}

func (*VotesContext) IsVotesContext() {}

func NewVotesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VotesContext {
	var p = new(VotesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VotesParserRULE_votes

	return p
}

func (s *VotesContext) GetParser() antlr.Parser { return s.parser }

func (s *VotesContext) VOTES() antlr.TerminalNode {
	return s.GetToken(VotesParserVOTES, 0)
}

func (s *VotesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VotesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VotesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VotesListener); ok {
		listenerT.EnterVotes(s)
	}
}

func (s *VotesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VotesListener); ok {
		listenerT.ExitVotes(s)
	}
}

func (p *VotesParser) Votes() (localctx IVotesContext) {
	this := p
	_ = this

	localctx = NewVotesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VotesParserRULE_votes)

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
		p.Match(VotesParserVOTES)
	}

	return localctx
}
