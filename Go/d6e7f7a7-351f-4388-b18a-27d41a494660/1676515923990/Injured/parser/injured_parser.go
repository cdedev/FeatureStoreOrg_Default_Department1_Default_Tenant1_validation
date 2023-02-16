// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676515923990/Injured.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Injured

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

type InjuredParser struct {
	*antlr.BaseParser
}

var injuredParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func injuredParserInit() {
	staticData := &injuredParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INJURED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "injured",
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

// InjuredParserInit initializes any static state used to implement InjuredParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInjuredParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InjuredParserInit() {
	staticData := &injuredParserStaticData
	staticData.once.Do(injuredParserInit)
}

// NewInjuredParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInjuredParser(input antlr.TokenStream) *InjuredParser {
	InjuredParserInit()
	this := new(InjuredParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &injuredParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Injured.g4"

	return this
}

// InjuredParser tokens.
const (
	InjuredParserEOF      = antlr.TokenEOF
	InjuredParserINJURED  = 1
	InjuredParserOWNKEY   = 2
	InjuredParserSPLITKEY = 3
	InjuredParserWS       = 4
)

// InjuredParser rules.
const (
	InjuredParserRULE_expression = 0
	InjuredParserRULE_injured    = 1
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
	p.RuleIndex = InjuredParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InjuredParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Injured() IInjuredContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInjuredContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInjuredContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(InjuredParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InjuredListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InjuredListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *InjuredParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InjuredParserRULE_expression)

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
		p.Injured()
	}
	{
		p.SetState(5)
		p.Match(InjuredParserEOF)
	}

	return localctx
}

// IInjuredContext is an interface to support dynamic dispatch.
type IInjuredContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInjuredContext differentiates from other interfaces.
	IsInjuredContext()
}

type InjuredContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInjuredContext() *InjuredContext {
	var p = new(InjuredContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InjuredParserRULE_injured
	return p
}

func (*InjuredContext) IsInjuredContext() {}

func NewInjuredContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InjuredContext {
	var p = new(InjuredContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InjuredParserRULE_injured

	return p
}

func (s *InjuredContext) GetParser() antlr.Parser { return s.parser }

func (s *InjuredContext) INJURED() antlr.TerminalNode {
	return s.GetToken(InjuredParserINJURED, 0)
}

func (s *InjuredContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InjuredContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InjuredContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InjuredListener); ok {
		listenerT.EnterInjured(s)
	}
}

func (s *InjuredContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InjuredListener); ok {
		listenerT.ExitInjured(s)
	}
}

func (p *InjuredParser) Injured() (localctx IInjuredContext) {
	this := p
	_ = this

	localctx = NewInjuredContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InjuredParserRULE_injured)

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
		p.Match(InjuredParserINJURED)
	}

	return localctx
}
