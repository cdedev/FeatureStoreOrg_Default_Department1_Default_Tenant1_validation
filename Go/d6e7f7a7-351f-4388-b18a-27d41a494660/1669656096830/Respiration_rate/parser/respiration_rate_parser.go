// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656096830/Respiration_rate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Respiration_rate

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

type Respiration_rateParser struct {
	*antlr.BaseParser
}

var respiration_rateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func respiration_rateParserInit() {
	staticData := &respiration_rateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RESPIRATION_RATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "respiration_rate",
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

// Respiration_rateParserInit initializes any static state used to implement Respiration_rateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRespiration_rateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Respiration_rateParserInit() {
	staticData := &respiration_rateParserStaticData
	staticData.once.Do(respiration_rateParserInit)
}

// NewRespiration_rateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRespiration_rateParser(input antlr.TokenStream) *Respiration_rateParser {
	Respiration_rateParserInit()
	this := new(Respiration_rateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &respiration_rateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Respiration_rate.g4"

	return this
}

// Respiration_rateParser tokens.
const (
	Respiration_rateParserEOF              = antlr.TokenEOF
	Respiration_rateParserRESPIRATION_RATE = 1
	Respiration_rateParserOWNKEY           = 2
	Respiration_rateParserSPLITKEY         = 3
	Respiration_rateParserWS               = 4
)

// Respiration_rateParser rules.
const (
	Respiration_rateParserRULE_expression       = 0
	Respiration_rateParserRULE_respiration_rate = 1
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
	p.RuleIndex = Respiration_rateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Respiration_rateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Respiration_rate() IRespiration_rateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRespiration_rateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRespiration_rateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Respiration_rateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Respiration_rateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Respiration_rateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Respiration_rateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Respiration_rateParserRULE_expression)

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
		p.Respiration_rate()
	}
	{
		p.SetState(5)
		p.Match(Respiration_rateParserEOF)
	}

	return localctx
}

// IRespiration_rateContext is an interface to support dynamic dispatch.
type IRespiration_rateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRespiration_rateContext differentiates from other interfaces.
	IsRespiration_rateContext()
}

type Respiration_rateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRespiration_rateContext() *Respiration_rateContext {
	var p = new(Respiration_rateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Respiration_rateParserRULE_respiration_rate
	return p
}

func (*Respiration_rateContext) IsRespiration_rateContext() {}

func NewRespiration_rateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Respiration_rateContext {
	var p = new(Respiration_rateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Respiration_rateParserRULE_respiration_rate

	return p
}

func (s *Respiration_rateContext) GetParser() antlr.Parser { return s.parser }

func (s *Respiration_rateContext) RESPIRATION_RATE() antlr.TerminalNode {
	return s.GetToken(Respiration_rateParserRESPIRATION_RATE, 0)
}

func (s *Respiration_rateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Respiration_rateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Respiration_rateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Respiration_rateListener); ok {
		listenerT.EnterRespiration_rate(s)
	}
}

func (s *Respiration_rateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Respiration_rateListener); ok {
		listenerT.ExitRespiration_rate(s)
	}
}

func (p *Respiration_rateParser) Respiration_rate() (localctx IRespiration_rateContext) {
	this := p
	_ = this

	localctx = NewRespiration_rateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Respiration_rateParserRULE_respiration_rate)

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
		p.Match(Respiration_rateParserRESPIRATION_RATE)
	}

	return localctx
}
