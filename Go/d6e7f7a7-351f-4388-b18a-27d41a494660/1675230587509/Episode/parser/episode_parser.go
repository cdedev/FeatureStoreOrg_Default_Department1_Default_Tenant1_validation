// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675230587509/Episode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Episode

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

type EpisodeParser struct {
	*antlr.BaseParser
}

var episodeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func episodeParserInit() {
	staticData := &episodeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EPISODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "episode",
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

// EpisodeParserInit initializes any static state used to implement EpisodeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEpisodeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EpisodeParserInit() {
	staticData := &episodeParserStaticData
	staticData.once.Do(episodeParserInit)
}

// NewEpisodeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEpisodeParser(input antlr.TokenStream) *EpisodeParser {
	EpisodeParserInit()
	this := new(EpisodeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &episodeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Episode.g4"

	return this
}

// EpisodeParser tokens.
const (
	EpisodeParserEOF      = antlr.TokenEOF
	EpisodeParserEPISODE  = 1
	EpisodeParserOWNKEY   = 2
	EpisodeParserSPLITKEY = 3
	EpisodeParserWS       = 4
)

// EpisodeParser rules.
const (
	EpisodeParserRULE_expression = 0
	EpisodeParserRULE_episode    = 1
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
	p.RuleIndex = EpisodeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EpisodeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Episode() IEpisodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpisodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpisodeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EpisodeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpisodeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpisodeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EpisodeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EpisodeParserRULE_expression)

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
		p.Episode()
	}
	{
		p.SetState(5)
		p.Match(EpisodeParserEOF)
	}

	return localctx
}

// IEpisodeContext is an interface to support dynamic dispatch.
type IEpisodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEpisodeContext differentiates from other interfaces.
	IsEpisodeContext()
}

type EpisodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpisodeContext() *EpisodeContext {
	var p = new(EpisodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EpisodeParserRULE_episode
	return p
}

func (*EpisodeContext) IsEpisodeContext() {}

func NewEpisodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpisodeContext {
	var p = new(EpisodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EpisodeParserRULE_episode

	return p
}

func (s *EpisodeContext) GetParser() antlr.Parser { return s.parser }

func (s *EpisodeContext) EPISODE() antlr.TerminalNode {
	return s.GetToken(EpisodeParserEPISODE, 0)
}

func (s *EpisodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpisodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EpisodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpisodeListener); ok {
		listenerT.EnterEpisode(s)
	}
}

func (s *EpisodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpisodeListener); ok {
		listenerT.ExitEpisode(s)
	}
}

func (p *EpisodeParser) Episode() (localctx IEpisodeContext) {
	this := p
	_ = this

	localctx = NewEpisodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EpisodeParserRULE_episode)

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
		p.Match(EpisodeParserEPISODE)
	}

	return localctx
}
