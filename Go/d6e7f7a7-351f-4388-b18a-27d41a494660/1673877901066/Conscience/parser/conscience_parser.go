// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673877901066/Conscience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Conscience

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

type ConscienceParser struct {
	*antlr.BaseParser
}

var conscienceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func conscienceParserInit() {
	staticData := &conscienceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONSCIENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "conscience",
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

// ConscienceParserInit initializes any static state used to implement ConscienceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConscienceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConscienceParserInit() {
	staticData := &conscienceParserStaticData
	staticData.once.Do(conscienceParserInit)
}

// NewConscienceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConscienceParser(input antlr.TokenStream) *ConscienceParser {
	ConscienceParserInit()
	this := new(ConscienceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &conscienceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Conscience.g4"

	return this
}

// ConscienceParser tokens.
const (
	ConscienceParserEOF        = antlr.TokenEOF
	ConscienceParserCONSCIENCE = 1
	ConscienceParserOWNKEY     = 2
	ConscienceParserSPLITKEY   = 3
	ConscienceParserWS         = 4
)

// ConscienceParser rules.
const (
	ConscienceParserRULE_expression = 0
	ConscienceParserRULE_conscience = 1
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
	p.RuleIndex = ConscienceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConscienceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Conscience() IConscienceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConscienceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConscienceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConscienceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConscienceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConscienceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConscienceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConscienceParserRULE_expression)

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
		p.Conscience()
	}
	{
		p.SetState(5)
		p.Match(ConscienceParserEOF)
	}

	return localctx
}

// IConscienceContext is an interface to support dynamic dispatch.
type IConscienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConscienceContext differentiates from other interfaces.
	IsConscienceContext()
}

type ConscienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConscienceContext() *ConscienceContext {
	var p = new(ConscienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConscienceParserRULE_conscience
	return p
}

func (*ConscienceContext) IsConscienceContext() {}

func NewConscienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConscienceContext {
	var p = new(ConscienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConscienceParserRULE_conscience

	return p
}

func (s *ConscienceContext) GetParser() antlr.Parser { return s.parser }

func (s *ConscienceContext) CONSCIENCE() antlr.TerminalNode {
	return s.GetToken(ConscienceParserCONSCIENCE, 0)
}

func (s *ConscienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConscienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConscienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConscienceListener); ok {
		listenerT.EnterConscience(s)
	}
}

func (s *ConscienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConscienceListener); ok {
		listenerT.ExitConscience(s)
	}
}

func (p *ConscienceParser) Conscience() (localctx IConscienceContext) {
	this := p
	_ = this

	localctx = NewConscienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConscienceParserRULE_conscience)

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
		p.Match(ConscienceParserCONSCIENCE)
	}

	return localctx
}
