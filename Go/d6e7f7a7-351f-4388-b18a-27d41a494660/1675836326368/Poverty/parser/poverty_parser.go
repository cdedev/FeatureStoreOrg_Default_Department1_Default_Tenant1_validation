// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836326368/Poverty.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Poverty

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

type PovertyParser struct {
	*antlr.BaseParser
}

var povertyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func povertyParserInit() {
	staticData := &povertyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "POVERTY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "poverty",
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

// PovertyParserInit initializes any static state used to implement PovertyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPovertyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PovertyParserInit() {
	staticData := &povertyParserStaticData
	staticData.once.Do(povertyParserInit)
}

// NewPovertyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPovertyParser(input antlr.TokenStream) *PovertyParser {
	PovertyParserInit()
	this := new(PovertyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &povertyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Poverty.g4"

	return this
}

// PovertyParser tokens.
const (
	PovertyParserEOF      = antlr.TokenEOF
	PovertyParserPOVERTY  = 1
	PovertyParserOWNKEY   = 2
	PovertyParserSPLITKEY = 3
	PovertyParserWS       = 4
)

// PovertyParser rules.
const (
	PovertyParserRULE_expression = 0
	PovertyParserRULE_poverty    = 1
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
	p.RuleIndex = PovertyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PovertyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Poverty() IPovertyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPovertyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPovertyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PovertyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PovertyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PovertyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PovertyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PovertyParserRULE_expression)

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
		p.Poverty()
	}
	{
		p.SetState(5)
		p.Match(PovertyParserEOF)
	}

	return localctx
}

// IPovertyContext is an interface to support dynamic dispatch.
type IPovertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPovertyContext differentiates from other interfaces.
	IsPovertyContext()
}

type PovertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPovertyContext() *PovertyContext {
	var p = new(PovertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PovertyParserRULE_poverty
	return p
}

func (*PovertyContext) IsPovertyContext() {}

func NewPovertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PovertyContext {
	var p = new(PovertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PovertyParserRULE_poverty

	return p
}

func (s *PovertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PovertyContext) POVERTY() antlr.TerminalNode {
	return s.GetToken(PovertyParserPOVERTY, 0)
}

func (s *PovertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PovertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PovertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PovertyListener); ok {
		listenerT.EnterPoverty(s)
	}
}

func (s *PovertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PovertyListener); ok {
		listenerT.ExitPoverty(s)
	}
}

func (p *PovertyParser) Poverty() (localctx IPovertyContext) {
	this := p
	_ = this

	localctx = NewPovertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PovertyParserRULE_poverty)

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
		p.Match(PovertyParserPOVERTY)
	}

	return localctx
}
