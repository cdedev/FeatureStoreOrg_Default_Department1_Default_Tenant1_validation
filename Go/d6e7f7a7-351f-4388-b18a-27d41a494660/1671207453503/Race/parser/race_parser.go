// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671207453503/Race.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Race

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

type RaceParser struct {
	*antlr.BaseParser
}

var raceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func raceParserInit() {
	staticData := &raceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RACE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "race",
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

// RaceParserInit initializes any static state used to implement RaceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRaceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RaceParserInit() {
	staticData := &raceParserStaticData
	staticData.once.Do(raceParserInit)
}

// NewRaceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRaceParser(input antlr.TokenStream) *RaceParser {
	RaceParserInit()
	this := new(RaceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &raceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Race.g4"

	return this
}

// RaceParser tokens.
const (
	RaceParserEOF      = antlr.TokenEOF
	RaceParserRACE     = 1
	RaceParserOWNKEY   = 2
	RaceParserSPLITKEY = 3
	RaceParserWS       = 4
)

// RaceParser rules.
const (
	RaceParserRULE_expression = 0
	RaceParserRULE_race       = 1
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
	p.RuleIndex = RaceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RaceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Race() IRaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRaceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RaceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RaceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RaceParserRULE_expression)

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
		p.Race()
	}
	{
		p.SetState(5)
		p.Match(RaceParserEOF)
	}

	return localctx
}

// IRaceContext is an interface to support dynamic dispatch.
type IRaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRaceContext differentiates from other interfaces.
	IsRaceContext()
}

type RaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRaceContext() *RaceContext {
	var p = new(RaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RaceParserRULE_race
	return p
}

func (*RaceContext) IsRaceContext() {}

func NewRaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RaceContext {
	var p = new(RaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RaceParserRULE_race

	return p
}

func (s *RaceContext) GetParser() antlr.Parser { return s.parser }

func (s *RaceContext) RACE() antlr.TerminalNode {
	return s.GetToken(RaceParserRACE, 0)
}

func (s *RaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceListener); ok {
		listenerT.EnterRace(s)
	}
}

func (s *RaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceListener); ok {
		listenerT.ExitRace(s)
	}
}

func (p *RaceParser) Race() (localctx IRaceContext) {
	this := p
	_ = this

	localctx = NewRaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RaceParserRULE_race)

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
		p.Match(RaceParserRACE)
	}

	return localctx
}
