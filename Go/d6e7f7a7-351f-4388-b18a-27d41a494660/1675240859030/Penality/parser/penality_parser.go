// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675240859030/Penality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Penality

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

type PenalityParser struct {
	*antlr.BaseParser
}

var penalityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func penalityParserInit() {
	staticData := &penalityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PENALITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "penality",
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

// PenalityParserInit initializes any static state used to implement PenalityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPenalityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PenalityParserInit() {
	staticData := &penalityParserStaticData
	staticData.once.Do(penalityParserInit)
}

// NewPenalityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPenalityParser(input antlr.TokenStream) *PenalityParser {
	PenalityParserInit()
	this := new(PenalityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &penalityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Penality.g4"

	return this
}

// PenalityParser tokens.
const (
	PenalityParserEOF      = antlr.TokenEOF
	PenalityParserPENALITY = 1
	PenalityParserOWNKEY   = 2
	PenalityParserSPLITKEY = 3
	PenalityParserWS       = 4
)

// PenalityParser rules.
const (
	PenalityParserRULE_expression = 0
	PenalityParserRULE_penality   = 1
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
	p.RuleIndex = PenalityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PenalityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Penality() IPenalityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPenalityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPenalityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PenalityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PenalityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PenalityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PenalityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PenalityParserRULE_expression)

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
		p.Penality()
	}
	{
		p.SetState(5)
		p.Match(PenalityParserEOF)
	}

	return localctx
}

// IPenalityContext is an interface to support dynamic dispatch.
type IPenalityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPenalityContext differentiates from other interfaces.
	IsPenalityContext()
}

type PenalityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPenalityContext() *PenalityContext {
	var p = new(PenalityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PenalityParserRULE_penality
	return p
}

func (*PenalityContext) IsPenalityContext() {}

func NewPenalityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PenalityContext {
	var p = new(PenalityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PenalityParserRULE_penality

	return p
}

func (s *PenalityContext) GetParser() antlr.Parser { return s.parser }

func (s *PenalityContext) PENALITY() antlr.TerminalNode {
	return s.GetToken(PenalityParserPENALITY, 0)
}

func (s *PenalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PenalityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PenalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PenalityListener); ok {
		listenerT.EnterPenality(s)
	}
}

func (s *PenalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PenalityListener); ok {
		listenerT.ExitPenality(s)
	}
}

func (p *PenalityParser) Penality() (localctx IPenalityContext) {
	this := p
	_ = this

	localctx = NewPenalityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PenalityParserRULE_penality)

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
		p.Match(PenalityParserPENALITY)
	}

	return localctx
}
