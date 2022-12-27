// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116531886/PaternityScore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PaternityScore

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

type PaternityScoreParser struct {
	*antlr.BaseParser
}

var paternityscoreParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paternityscoreParserInit() {
	staticData := &paternityscoreParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PATERNITYSCORE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "paternityscore",
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

// PaternityScoreParserInit initializes any static state used to implement PaternityScoreParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPaternityScoreParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PaternityScoreParserInit() {
	staticData := &paternityscoreParserStaticData
	staticData.once.Do(paternityscoreParserInit)
}

// NewPaternityScoreParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPaternityScoreParser(input antlr.TokenStream) *PaternityScoreParser {
	PaternityScoreParserInit()
	this := new(PaternityScoreParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &paternityscoreParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PaternityScore.g4"

	return this
}

// PaternityScoreParser tokens.
const (
	PaternityScoreParserEOF            = antlr.TokenEOF
	PaternityScoreParserPATERNITYSCORE = 1
	PaternityScoreParserOWNKEY         = 2
	PaternityScoreParserSPLITKEY       = 3
	PaternityScoreParserWS             = 4
)

// PaternityScoreParser rules.
const (
	PaternityScoreParserRULE_expression     = 0
	PaternityScoreParserRULE_paternityscore = 1
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
	p.RuleIndex = PaternityScoreParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PaternityScoreParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Paternityscore() IPaternityscoreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPaternityscoreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPaternityscoreContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PaternityScoreParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PaternityScoreListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PaternityScoreListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PaternityScoreParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PaternityScoreParserRULE_expression)

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
		p.Paternityscore()
	}
	{
		p.SetState(5)
		p.Match(PaternityScoreParserEOF)
	}

	return localctx
}

// IPaternityscoreContext is an interface to support dynamic dispatch.
type IPaternityscoreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPaternityscoreContext differentiates from other interfaces.
	IsPaternityscoreContext()
}

type PaternityscoreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPaternityscoreContext() *PaternityscoreContext {
	var p = new(PaternityscoreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PaternityScoreParserRULE_paternityscore
	return p
}

func (*PaternityscoreContext) IsPaternityscoreContext() {}

func NewPaternityscoreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PaternityscoreContext {
	var p = new(PaternityscoreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PaternityScoreParserRULE_paternityscore

	return p
}

func (s *PaternityscoreContext) GetParser() antlr.Parser { return s.parser }

func (s *PaternityscoreContext) PATERNITYSCORE() antlr.TerminalNode {
	return s.GetToken(PaternityScoreParserPATERNITYSCORE, 0)
}

func (s *PaternityscoreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PaternityscoreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PaternityscoreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PaternityScoreListener); ok {
		listenerT.EnterPaternityscore(s)
	}
}

func (s *PaternityscoreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PaternityScoreListener); ok {
		listenerT.ExitPaternityscore(s)
	}
}

func (p *PaternityScoreParser) Paternityscore() (localctx IPaternityscoreContext) {
	this := p
	_ = this

	localctx = NewPaternityscoreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PaternityScoreParserRULE_paternityscore)

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
		p.Match(PaternityScoreParserPATERNITYSCORE)
	}

	return localctx
}
