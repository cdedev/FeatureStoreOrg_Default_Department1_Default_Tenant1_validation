// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118775991/Nausea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nausea

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

type NauseaParser struct {
	*antlr.BaseParser
}

var nauseaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nauseaParserInit() {
	staticData := &nauseaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NAUSEA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "nausea",
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

// NauseaParserInit initializes any static state used to implement NauseaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNauseaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NauseaParserInit() {
	staticData := &nauseaParserStaticData
	staticData.once.Do(nauseaParserInit)
}

// NewNauseaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNauseaParser(input antlr.TokenStream) *NauseaParser {
	NauseaParserInit()
	this := new(NauseaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nauseaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Nausea.g4"

	return this
}

// NauseaParser tokens.
const (
	NauseaParserEOF      = antlr.TokenEOF
	NauseaParserNAUSEA   = 1
	NauseaParserOWNKEY   = 2
	NauseaParserSPLITKEY = 3
	NauseaParserWS       = 4
)

// NauseaParser rules.
const (
	NauseaParserRULE_expression = 0
	NauseaParserRULE_nausea     = 1
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
	p.RuleIndex = NauseaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NauseaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Nausea() INauseaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INauseaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INauseaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NauseaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NauseaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NauseaParserRULE_expression)

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
		p.Nausea()
	}
	{
		p.SetState(5)
		p.Match(NauseaParserEOF)
	}

	return localctx
}

// INauseaContext is an interface to support dynamic dispatch.
type INauseaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNauseaContext differentiates from other interfaces.
	IsNauseaContext()
}

type NauseaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNauseaContext() *NauseaContext {
	var p = new(NauseaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NauseaParserRULE_nausea
	return p
}

func (*NauseaContext) IsNauseaContext() {}

func NewNauseaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NauseaContext {
	var p = new(NauseaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NauseaParserRULE_nausea

	return p
}

func (s *NauseaContext) GetParser() antlr.Parser { return s.parser }

func (s *NauseaContext) NAUSEA() antlr.TerminalNode {
	return s.GetToken(NauseaParserNAUSEA, 0)
}

func (s *NauseaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NauseaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NauseaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseaListener); ok {
		listenerT.EnterNausea(s)
	}
}

func (s *NauseaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseaListener); ok {
		listenerT.ExitNausea(s)
	}
}

func (p *NauseaParser) Nausea() (localctx INauseaContext) {
	this := p
	_ = this

	localctx = NewNauseaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NauseaParserRULE_nausea)

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
		p.Match(NauseaParserNAUSEA)
	}

	return localctx
}
