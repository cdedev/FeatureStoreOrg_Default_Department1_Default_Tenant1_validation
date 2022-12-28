// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202276107/Charisma.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Charisma

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

type CharismaParser struct {
	*antlr.BaseParser
}

var charismaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func charismaParserInit() {
	staticData := &charismaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHARISMA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "charisma",
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

// CharismaParserInit initializes any static state used to implement CharismaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCharismaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CharismaParserInit() {
	staticData := &charismaParserStaticData
	staticData.once.Do(charismaParserInit)
}

// NewCharismaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCharismaParser(input antlr.TokenStream) *CharismaParser {
	CharismaParserInit()
	this := new(CharismaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &charismaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Charisma.g4"

	return this
}

// CharismaParser tokens.
const (
	CharismaParserEOF      = antlr.TokenEOF
	CharismaParserCHARISMA = 1
	CharismaParserOWNKEY   = 2
	CharismaParserSPLITKEY = 3
	CharismaParserWS       = 4
)

// CharismaParser rules.
const (
	CharismaParserRULE_expression = 0
	CharismaParserRULE_charisma   = 1
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
	p.RuleIndex = CharismaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CharismaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Charisma() ICharismaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharismaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharismaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CharismaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CharismaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CharismaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CharismaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CharismaParserRULE_expression)

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
		p.Charisma()
	}
	{
		p.SetState(5)
		p.Match(CharismaParserEOF)
	}

	return localctx
}

// ICharismaContext is an interface to support dynamic dispatch.
type ICharismaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharismaContext differentiates from other interfaces.
	IsCharismaContext()
}

type CharismaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharismaContext() *CharismaContext {
	var p = new(CharismaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CharismaParserRULE_charisma
	return p
}

func (*CharismaContext) IsCharismaContext() {}

func NewCharismaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharismaContext {
	var p = new(CharismaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CharismaParserRULE_charisma

	return p
}

func (s *CharismaContext) GetParser() antlr.Parser { return s.parser }

func (s *CharismaContext) CHARISMA() antlr.TerminalNode {
	return s.GetToken(CharismaParserCHARISMA, 0)
}

func (s *CharismaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharismaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharismaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CharismaListener); ok {
		listenerT.EnterCharisma(s)
	}
}

func (s *CharismaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CharismaListener); ok {
		listenerT.ExitCharisma(s)
	}
}

func (p *CharismaParser) Charisma() (localctx ICharismaContext) {
	this := p
	_ = this

	localctx = NewCharismaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CharismaParserRULE_charisma)

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
		p.Match(CharismaParserCHARISMA)
	}

	return localctx
}
