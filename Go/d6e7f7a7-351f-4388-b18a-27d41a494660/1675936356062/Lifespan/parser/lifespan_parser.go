// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675936356062/Lifespan.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lifespan

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

type LifespanParser struct {
	*antlr.BaseParser
}

var lifespanParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lifespanParserInit() {
	staticData := &lifespanParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LIFESPAN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "lifespan",
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

// LifespanParserInit initializes any static state used to implement LifespanParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLifespanParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LifespanParserInit() {
	staticData := &lifespanParserStaticData
	staticData.once.Do(lifespanParserInit)
}

// NewLifespanParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLifespanParser(input antlr.TokenStream) *LifespanParser {
	LifespanParserInit()
	this := new(LifespanParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lifespanParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Lifespan.g4"

	return this
}

// LifespanParser tokens.
const (
	LifespanParserEOF      = antlr.TokenEOF
	LifespanParserLIFESPAN = 1
	LifespanParserOWNKEY   = 2
	LifespanParserSPLITKEY = 3
	LifespanParserWS       = 4
)

// LifespanParser rules.
const (
	LifespanParserRULE_expression = 0
	LifespanParserRULE_lifespan   = 1
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
	p.RuleIndex = LifespanParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LifespanParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Lifespan() ILifespanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILifespanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILifespanContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LifespanParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LifespanListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LifespanListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LifespanParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LifespanParserRULE_expression)

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
		p.Lifespan()
	}
	{
		p.SetState(5)
		p.Match(LifespanParserEOF)
	}

	return localctx
}

// ILifespanContext is an interface to support dynamic dispatch.
type ILifespanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLifespanContext differentiates from other interfaces.
	IsLifespanContext()
}

type LifespanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLifespanContext() *LifespanContext {
	var p = new(LifespanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LifespanParserRULE_lifespan
	return p
}

func (*LifespanContext) IsLifespanContext() {}

func NewLifespanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LifespanContext {
	var p = new(LifespanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LifespanParserRULE_lifespan

	return p
}

func (s *LifespanContext) GetParser() antlr.Parser { return s.parser }

func (s *LifespanContext) LIFESPAN() antlr.TerminalNode {
	return s.GetToken(LifespanParserLIFESPAN, 0)
}

func (s *LifespanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LifespanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LifespanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LifespanListener); ok {
		listenerT.EnterLifespan(s)
	}
}

func (s *LifespanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LifespanListener); ok {
		listenerT.ExitLifespan(s)
	}
}

func (p *LifespanParser) Lifespan() (localctx ILifespanContext) {
	this := p
	_ = this

	localctx = NewLifespanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LifespanParserRULE_lifespan)

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
		p.Match(LifespanParserLIFESPAN)
	}

	return localctx
}
