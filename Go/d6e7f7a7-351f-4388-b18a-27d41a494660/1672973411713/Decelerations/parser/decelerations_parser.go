// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672973411713/Decelerations.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Decelerations

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

type DecelerationsParser struct {
	*antlr.BaseParser
}

var decelerationsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func decelerationsParserInit() {
	staticData := &decelerationsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DECELERATIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "decelerations",
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

// DecelerationsParserInit initializes any static state used to implement DecelerationsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDecelerationsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DecelerationsParserInit() {
	staticData := &decelerationsParserStaticData
	staticData.once.Do(decelerationsParserInit)
}

// NewDecelerationsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDecelerationsParser(input antlr.TokenStream) *DecelerationsParser {
	DecelerationsParserInit()
	this := new(DecelerationsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &decelerationsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Decelerations.g4"

	return this
}

// DecelerationsParser tokens.
const (
	DecelerationsParserEOF           = antlr.TokenEOF
	DecelerationsParserDECELERATIONS = 1
	DecelerationsParserOWNKEY        = 2
	DecelerationsParserSPLITKEY      = 3
	DecelerationsParserWS            = 4
)

// DecelerationsParser rules.
const (
	DecelerationsParserRULE_expression    = 0
	DecelerationsParserRULE_decelerations = 1
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
	p.RuleIndex = DecelerationsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DecelerationsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Decelerations() IDecelerationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecelerationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecelerationsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DecelerationsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecelerationsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecelerationsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DecelerationsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DecelerationsParserRULE_expression)

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
		p.Decelerations()
	}
	{
		p.SetState(5)
		p.Match(DecelerationsParserEOF)
	}

	return localctx
}

// IDecelerationsContext is an interface to support dynamic dispatch.
type IDecelerationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecelerationsContext differentiates from other interfaces.
	IsDecelerationsContext()
}

type DecelerationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecelerationsContext() *DecelerationsContext {
	var p = new(DecelerationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DecelerationsParserRULE_decelerations
	return p
}

func (*DecelerationsContext) IsDecelerationsContext() {}

func NewDecelerationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecelerationsContext {
	var p = new(DecelerationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DecelerationsParserRULE_decelerations

	return p
}

func (s *DecelerationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DecelerationsContext) DECELERATIONS() antlr.TerminalNode {
	return s.GetToken(DecelerationsParserDECELERATIONS, 0)
}

func (s *DecelerationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecelerationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecelerationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecelerationsListener); ok {
		listenerT.EnterDecelerations(s)
	}
}

func (s *DecelerationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecelerationsListener); ok {
		listenerT.ExitDecelerations(s)
	}
}

func (p *DecelerationsParser) Decelerations() (localctx IDecelerationsContext) {
	this := p
	_ = this

	localctx = NewDecelerationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DecelerationsParserRULE_decelerations)

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
		p.Match(DecelerationsParserDECELERATIONS)
	}

	return localctx
}
