// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672073289495/Level.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Level

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

type LevelParser struct {
	*antlr.BaseParser
}

var levelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func levelParserInit() {
	staticData := &levelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LEVEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "level",
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

// LevelParserInit initializes any static state used to implement LevelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLevelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LevelParserInit() {
	staticData := &levelParserStaticData
	staticData.once.Do(levelParserInit)
}

// NewLevelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLevelParser(input antlr.TokenStream) *LevelParser {
	LevelParserInit()
	this := new(LevelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &levelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Level.g4"

	return this
}

// LevelParser tokens.
const (
	LevelParserEOF      = antlr.TokenEOF
	LevelParserLEVEL    = 1
	LevelParserOWNKEY   = 2
	LevelParserSPLITKEY = 3
	LevelParserWS       = 4
)

// LevelParser rules.
const (
	LevelParserRULE_expression = 0
	LevelParserRULE_level      = 1
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
	p.RuleIndex = LevelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LevelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Level() ILevelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILevelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILevelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LevelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LevelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LevelParserRULE_expression)

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
		p.Level()
	}
	{
		p.SetState(5)
		p.Match(LevelParserEOF)
	}

	return localctx
}

// ILevelContext is an interface to support dynamic dispatch.
type ILevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLevelContext differentiates from other interfaces.
	IsLevelContext()
}

type LevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLevelContext() *LevelContext {
	var p = new(LevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LevelParserRULE_level
	return p
}

func (*LevelContext) IsLevelContext() {}

func NewLevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LevelContext {
	var p = new(LevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LevelParserRULE_level

	return p
}

func (s *LevelContext) GetParser() antlr.Parser { return s.parser }

func (s *LevelContext) LEVEL() antlr.TerminalNode {
	return s.GetToken(LevelParserLEVEL, 0)
}

func (s *LevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevelListener); ok {
		listenerT.EnterLevel(s)
	}
}

func (s *LevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LevelListener); ok {
		listenerT.ExitLevel(s)
	}
}

func (p *LevelParser) Level() (localctx ILevelContext) {
	this := p
	_ = this

	localctx = NewLevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LevelParserRULE_level)

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
		p.Match(LevelParserLEVEL)
	}

	return localctx
}
