// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721668865/Season.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Season

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

type SeasonParser struct {
	*antlr.BaseParser
}

var seasonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func seasonParserInit() {
	staticData := &seasonParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SEASON", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "season",
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

// SeasonParserInit initializes any static state used to implement SeasonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSeasonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SeasonParserInit() {
	staticData := &seasonParserStaticData
	staticData.once.Do(seasonParserInit)
}

// NewSeasonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSeasonParser(input antlr.TokenStream) *SeasonParser {
	SeasonParserInit()
	this := new(SeasonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &seasonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Season.g4"

	return this
}

// SeasonParser tokens.
const (
	SeasonParserEOF      = antlr.TokenEOF
	SeasonParserSEASON   = 1
	SeasonParserOWNKEY   = 2
	SeasonParserSPLITKEY = 3
	SeasonParserWS       = 4
)

// SeasonParser rules.
const (
	SeasonParserRULE_expression = 0
	SeasonParserRULE_season     = 1
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
	p.RuleIndex = SeasonParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SeasonParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Season() ISeasonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeasonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeasonContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SeasonParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SeasonListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SeasonListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SeasonParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SeasonParserRULE_expression)

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
		p.Season()
	}
	{
		p.SetState(5)
		p.Match(SeasonParserEOF)
	}

	return localctx
}

// ISeasonContext is an interface to support dynamic dispatch.
type ISeasonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSeasonContext differentiates from other interfaces.
	IsSeasonContext()
}

type SeasonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeasonContext() *SeasonContext {
	var p = new(SeasonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SeasonParserRULE_season
	return p
}

func (*SeasonContext) IsSeasonContext() {}

func NewSeasonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeasonContext {
	var p = new(SeasonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SeasonParserRULE_season

	return p
}

func (s *SeasonContext) GetParser() antlr.Parser { return s.parser }

func (s *SeasonContext) SEASON() antlr.TerminalNode {
	return s.GetToken(SeasonParserSEASON, 0)
}

func (s *SeasonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeasonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeasonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SeasonListener); ok {
		listenerT.EnterSeason(s)
	}
}

func (s *SeasonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SeasonListener); ok {
		listenerT.ExitSeason(s)
	}
}

func (p *SeasonParser) Season() (localctx ISeasonContext) {
	this := p
	_ = this

	localctx = NewSeasonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SeasonParserRULE_season)

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
		p.Match(SeasonParserSEASON)
	}

	return localctx
}