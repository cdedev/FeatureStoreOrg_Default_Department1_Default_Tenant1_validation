// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672635660795/GameId.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GameId

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

type GameIdParser struct {
	*antlr.BaseParser
}

var gameidParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gameidParserInit() {
	staticData := &gameidParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GAMEID", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gameid",
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

// GameIdParserInit initializes any static state used to implement GameIdParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGameIdParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GameIdParserInit() {
	staticData := &gameidParserStaticData
	staticData.once.Do(gameidParserInit)
}

// NewGameIdParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGameIdParser(input antlr.TokenStream) *GameIdParser {
	GameIdParserInit()
	this := new(GameIdParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gameidParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "GameId.g4"

	return this
}

// GameIdParser tokens.
const (
	GameIdParserEOF      = antlr.TokenEOF
	GameIdParserGAMEID   = 1
	GameIdParserOWNKEY   = 2
	GameIdParserSPLITKEY = 3
	GameIdParserWS       = 4
)

// GameIdParser rules.
const (
	GameIdParserRULE_expression = 0
	GameIdParserRULE_gameid     = 1
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
	p.RuleIndex = GameIdParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GameIdParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gameid() IGameidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGameidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGameidContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GameIdParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GameIdListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GameIdListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GameIdParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GameIdParserRULE_expression)

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
		p.Gameid()
	}
	{
		p.SetState(5)
		p.Match(GameIdParserEOF)
	}

	return localctx
}

// IGameidContext is an interface to support dynamic dispatch.
type IGameidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGameidContext differentiates from other interfaces.
	IsGameidContext()
}

type GameidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGameidContext() *GameidContext {
	var p = new(GameidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GameIdParserRULE_gameid
	return p
}

func (*GameidContext) IsGameidContext() {}

func NewGameidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GameidContext {
	var p = new(GameidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GameIdParserRULE_gameid

	return p
}

func (s *GameidContext) GetParser() antlr.Parser { return s.parser }

func (s *GameidContext) GAMEID() antlr.TerminalNode {
	return s.GetToken(GameIdParserGAMEID, 0)
}

func (s *GameidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GameidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GameidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GameIdListener); ok {
		listenerT.EnterGameid(s)
	}
}

func (s *GameidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GameIdListener); ok {
		listenerT.ExitGameid(s)
	}
}

func (p *GameIdParser) Gameid() (localctx IGameidContext) {
	this := p
	_ = this

	localctx = NewGameidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GameIdParserRULE_gameid)

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
		p.Match(GameIdParserGAMEID)
	}

	return localctx
}
