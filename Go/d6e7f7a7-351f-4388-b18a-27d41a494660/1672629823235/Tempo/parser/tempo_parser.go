// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672629823235/Tempo.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tempo

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

type TempoParser struct {
	*antlr.BaseParser
}

var tempoParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tempoParserInit() {
	staticData := &tempoParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TEMPO", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "tempo",
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

// TempoParserInit initializes any static state used to implement TempoParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTempoParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TempoParserInit() {
	staticData := &tempoParserStaticData
	staticData.once.Do(tempoParserInit)
}

// NewTempoParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTempoParser(input antlr.TokenStream) *TempoParser {
	TempoParserInit()
	this := new(TempoParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tempoParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tempo.g4"

	return this
}

// TempoParser tokens.
const (
	TempoParserEOF      = antlr.TokenEOF
	TempoParserTEMPO    = 1
	TempoParserOWNKEY   = 2
	TempoParserSPLITKEY = 3
	TempoParserWS       = 4
)

// TempoParser rules.
const (
	TempoParserRULE_expression = 0
	TempoParserRULE_tempo      = 1
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
	p.RuleIndex = TempoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tempo() ITempoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITempoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITempoContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TempoParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TempoParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TempoParserRULE_expression)

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
		p.Tempo()
	}
	{
		p.SetState(5)
		p.Match(TempoParserEOF)
	}

	return localctx
}

// ITempoContext is an interface to support dynamic dispatch.
type ITempoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTempoContext differentiates from other interfaces.
	IsTempoContext()
}

type TempoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTempoContext() *TempoContext {
	var p = new(TempoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TempoParserRULE_tempo
	return p
}

func (*TempoContext) IsTempoContext() {}

func NewTempoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TempoContext {
	var p = new(TempoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TempoParserRULE_tempo

	return p
}

func (s *TempoContext) GetParser() antlr.Parser { return s.parser }

func (s *TempoContext) TEMPO() antlr.TerminalNode {
	return s.GetToken(TempoParserTEMPO, 0)
}

func (s *TempoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TempoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TempoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.EnterTempo(s)
	}
}

func (s *TempoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TempoListener); ok {
		listenerT.ExitTempo(s)
	}
}

func (p *TempoParser) Tempo() (localctx ITempoContext) {
	this := p
	_ = this

	localctx = NewTempoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TempoParserRULE_tempo)

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
		p.Match(TempoParserTEMPO)
	}

	return localctx
}
