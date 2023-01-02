// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672633185710/Z_Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Z_Score

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

type Z_ScoreParser struct {
	*antlr.BaseParser
}

var z_scoreParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func z_scoreParserInit() {
	staticData := &z_scoreParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "Z_SCORE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "z_score",
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

// Z_ScoreParserInit initializes any static state used to implement Z_ScoreParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewZ_ScoreParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Z_ScoreParserInit() {
	staticData := &z_scoreParserStaticData
	staticData.once.Do(z_scoreParserInit)
}

// NewZ_ScoreParser produces a new parser instance for the optional input antlr.TokenStream.
func NewZ_ScoreParser(input antlr.TokenStream) *Z_ScoreParser {
	Z_ScoreParserInit()
	this := new(Z_ScoreParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &z_scoreParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Z_Score.g4"

	return this
}

// Z_ScoreParser tokens.
const (
	Z_ScoreParserEOF      = antlr.TokenEOF
	Z_ScoreParserZ_SCORE  = 1
	Z_ScoreParserOWNKEY   = 2
	Z_ScoreParserSPLITKEY = 3
	Z_ScoreParserWS       = 4
)

// Z_ScoreParser rules.
const (
	Z_ScoreParserRULE_expression = 0
	Z_ScoreParserRULE_z_score    = 1
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
	p.RuleIndex = Z_ScoreParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Z_ScoreParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Z_score() IZ_scoreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZ_scoreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZ_scoreContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Z_ScoreParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Z_ScoreListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Z_ScoreListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Z_ScoreParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Z_ScoreParserRULE_expression)

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
		p.Z_score()
	}
	{
		p.SetState(5)
		p.Match(Z_ScoreParserEOF)
	}

	return localctx
}

// IZ_scoreContext is an interface to support dynamic dispatch.
type IZ_scoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZ_scoreContext differentiates from other interfaces.
	IsZ_scoreContext()
}

type Z_scoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZ_scoreContext() *Z_scoreContext {
	var p = new(Z_scoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Z_ScoreParserRULE_z_score
	return p
}

func (*Z_scoreContext) IsZ_scoreContext() {}

func NewZ_scoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Z_scoreContext {
	var p = new(Z_scoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Z_ScoreParserRULE_z_score

	return p
}

func (s *Z_scoreContext) GetParser() antlr.Parser { return s.parser }

func (s *Z_scoreContext) Z_SCORE() antlr.TerminalNode {
	return s.GetToken(Z_ScoreParserZ_SCORE, 0)
}

func (s *Z_scoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Z_scoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Z_scoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Z_ScoreListener); ok {
		listenerT.EnterZ_score(s)
	}
}

func (s *Z_scoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Z_ScoreListener); ok {
		listenerT.ExitZ_score(s)
	}
}

func (p *Z_ScoreParser) Z_score() (localctx IZ_scoreContext) {
	this := p
	_ = this

	localctx = NewZ_scoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Z_ScoreParserRULE_z_score)

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
		p.Match(Z_ScoreParserZ_SCORE)
	}

	return localctx
}
