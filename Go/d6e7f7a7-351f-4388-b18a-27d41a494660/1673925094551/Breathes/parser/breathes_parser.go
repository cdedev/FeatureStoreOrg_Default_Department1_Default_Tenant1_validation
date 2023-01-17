// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925094551/Breathes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Breathes

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

type BreathesParser struct {
	*antlr.BaseParser
}

var breathesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func breathesParserInit() {
	staticData := &breathesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BREATHES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "breathes",
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

// BreathesParserInit initializes any static state used to implement BreathesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBreathesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BreathesParserInit() {
	staticData := &breathesParserStaticData
	staticData.once.Do(breathesParserInit)
}

// NewBreathesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBreathesParser(input antlr.TokenStream) *BreathesParser {
	BreathesParserInit()
	this := new(BreathesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &breathesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Breathes.g4"

	return this
}

// BreathesParser tokens.
const (
	BreathesParserEOF      = antlr.TokenEOF
	BreathesParserBREATHES = 1
	BreathesParserOWNKEY   = 2
	BreathesParserSPLITKEY = 3
	BreathesParserWS       = 4
)

// BreathesParser rules.
const (
	BreathesParserRULE_expression = 0
	BreathesParserRULE_breathes   = 1
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
	p.RuleIndex = BreathesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BreathesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Breathes() IBreathesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreathesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreathesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BreathesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BreathesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BreathesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BreathesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BreathesParserRULE_expression)

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
		p.Breathes()
	}
	{
		p.SetState(5)
		p.Match(BreathesParserEOF)
	}

	return localctx
}

// IBreathesContext is an interface to support dynamic dispatch.
type IBreathesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreathesContext differentiates from other interfaces.
	IsBreathesContext()
}

type BreathesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreathesContext() *BreathesContext {
	var p = new(BreathesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BreathesParserRULE_breathes
	return p
}

func (*BreathesContext) IsBreathesContext() {}

func NewBreathesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreathesContext {
	var p = new(BreathesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BreathesParserRULE_breathes

	return p
}

func (s *BreathesContext) GetParser() antlr.Parser { return s.parser }

func (s *BreathesContext) BREATHES() antlr.TerminalNode {
	return s.GetToken(BreathesParserBREATHES, 0)
}

func (s *BreathesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreathesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreathesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BreathesListener); ok {
		listenerT.EnterBreathes(s)
	}
}

func (s *BreathesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BreathesListener); ok {
		listenerT.ExitBreathes(s)
	}
}

func (p *BreathesParser) Breathes() (localctx IBreathesContext) {
	this := p
	_ = this

	localctx = NewBreathesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BreathesParserRULE_breathes)

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
		p.Match(BreathesParserBREATHES)
	}

	return localctx
}
