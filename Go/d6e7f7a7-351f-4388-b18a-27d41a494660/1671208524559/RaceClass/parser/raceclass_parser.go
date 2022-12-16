// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671208524559/RaceClass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RaceClass

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

type RaceClassParser struct {
	*antlr.BaseParser
}

var raceclassParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func raceclassParserInit() {
	staticData := &raceclassParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RACECLASS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "raceclass",
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

// RaceClassParserInit initializes any static state used to implement RaceClassParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRaceClassParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RaceClassParserInit() {
	staticData := &raceclassParserStaticData
	staticData.once.Do(raceclassParserInit)
}

// NewRaceClassParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRaceClassParser(input antlr.TokenStream) *RaceClassParser {
	RaceClassParserInit()
	this := new(RaceClassParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &raceclassParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RaceClass.g4"

	return this
}

// RaceClassParser tokens.
const (
	RaceClassParserEOF       = antlr.TokenEOF
	RaceClassParserRACECLASS = 1
	RaceClassParserOWNKEY    = 2
	RaceClassParserSPLITKEY  = 3
	RaceClassParserWS        = 4
)

// RaceClassParser rules.
const (
	RaceClassParserRULE_expression = 0
	RaceClassParserRULE_raceclass  = 1
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
	p.RuleIndex = RaceClassParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RaceClassParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Raceclass() IRaceclassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRaceclassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRaceclassContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RaceClassParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceClassListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceClassListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RaceClassParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RaceClassParserRULE_expression)

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
		p.Raceclass()
	}
	{
		p.SetState(5)
		p.Match(RaceClassParserEOF)
	}

	return localctx
}

// IRaceclassContext is an interface to support dynamic dispatch.
type IRaceclassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRaceclassContext differentiates from other interfaces.
	IsRaceclassContext()
}

type RaceclassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRaceclassContext() *RaceclassContext {
	var p = new(RaceclassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RaceClassParserRULE_raceclass
	return p
}

func (*RaceclassContext) IsRaceclassContext() {}

func NewRaceclassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RaceclassContext {
	var p = new(RaceclassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RaceClassParserRULE_raceclass

	return p
}

func (s *RaceclassContext) GetParser() antlr.Parser { return s.parser }

func (s *RaceclassContext) RACECLASS() antlr.TerminalNode {
	return s.GetToken(RaceClassParserRACECLASS, 0)
}

func (s *RaceclassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RaceclassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RaceclassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceClassListener); ok {
		listenerT.EnterRaceclass(s)
	}
}

func (s *RaceclassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RaceClassListener); ok {
		listenerT.ExitRaceclass(s)
	}
}

func (p *RaceClassParser) Raceclass() (localctx IRaceclassContext) {
	this := p
	_ = this

	localctx = NewRaceclassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RaceClassParserRULE_raceclass)

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
		p.Match(RaceClassParserRACECLASS)
	}

	return localctx
}
