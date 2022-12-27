// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672112566753/Games.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Games

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

type GamesParser struct {
	*antlr.BaseParser
}

var gamesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gamesParserInit() {
	staticData := &gamesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GAMES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "games",
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

// GamesParserInit initializes any static state used to implement GamesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGamesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GamesParserInit() {
	staticData := &gamesParserStaticData
	staticData.once.Do(gamesParserInit)
}

// NewGamesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGamesParser(input antlr.TokenStream) *GamesParser {
	GamesParserInit()
	this := new(GamesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gamesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Games.g4"

	return this
}

// GamesParser tokens.
const (
	GamesParserEOF      = antlr.TokenEOF
	GamesParserGAMES    = 1
	GamesParserOWNKEY   = 2
	GamesParserSPLITKEY = 3
	GamesParserWS       = 4
)

// GamesParser rules.
const (
	GamesParserRULE_expression = 0
	GamesParserRULE_games      = 1
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
	p.RuleIndex = GamesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GamesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Games() IGamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGamesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GamesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GamesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GamesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GamesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GamesParserRULE_expression)

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
		p.Games()
	}
	{
		p.SetState(5)
		p.Match(GamesParserEOF)
	}

	return localctx
}

// IGamesContext is an interface to support dynamic dispatch.
type IGamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGamesContext differentiates from other interfaces.
	IsGamesContext()
}

type GamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGamesContext() *GamesContext {
	var p = new(GamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GamesParserRULE_games
	return p
}

func (*GamesContext) IsGamesContext() {}

func NewGamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GamesContext {
	var p = new(GamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GamesParserRULE_games

	return p
}

func (s *GamesContext) GetParser() antlr.Parser { return s.parser }

func (s *GamesContext) GAMES() antlr.TerminalNode {
	return s.GetToken(GamesParserGAMES, 0)
}

func (s *GamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GamesListener); ok {
		listenerT.EnterGames(s)
	}
}

func (s *GamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GamesListener); ok {
		listenerT.ExitGames(s)
	}
}

func (p *GamesParser) Games() (localctx IGamesContext) {
	this := p
	_ = this

	localctx = NewGamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GamesParserRULE_games)

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
		p.Match(GamesParserGAMES)
	}

	return localctx
}
