// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675182928861/SmokingStatus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SmokingStatus

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

type SmokingStatusParser struct {
	*antlr.BaseParser
}

var smokingstatusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func smokingstatusParserInit() {
	staticData := &smokingstatusParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SMOKINGSTATUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "smokingstatus",
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

// SmokingStatusParserInit initializes any static state used to implement SmokingStatusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSmokingStatusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SmokingStatusParserInit() {
	staticData := &smokingstatusParserStaticData
	staticData.once.Do(smokingstatusParserInit)
}

// NewSmokingStatusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSmokingStatusParser(input antlr.TokenStream) *SmokingStatusParser {
	SmokingStatusParserInit()
	this := new(SmokingStatusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &smokingstatusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SmokingStatus.g4"

	return this
}

// SmokingStatusParser tokens.
const (
	SmokingStatusParserEOF           = antlr.TokenEOF
	SmokingStatusParserSMOKINGSTATUS = 1
	SmokingStatusParserOWNKEY        = 2
	SmokingStatusParserSPLITKEY      = 3
	SmokingStatusParserWS            = 4
)

// SmokingStatusParser rules.
const (
	SmokingStatusParserRULE_expression    = 0
	SmokingStatusParserRULE_smokingstatus = 1
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
	p.RuleIndex = SmokingStatusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmokingStatusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Smokingstatus() ISmokingstatusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISmokingstatusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISmokingstatusContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SmokingStatusParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokingStatusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokingStatusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SmokingStatusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SmokingStatusParserRULE_expression)

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
		p.Smokingstatus()
	}
	{
		p.SetState(5)
		p.Match(SmokingStatusParserEOF)
	}

	return localctx
}

// ISmokingstatusContext is an interface to support dynamic dispatch.
type ISmokingstatusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSmokingstatusContext differentiates from other interfaces.
	IsSmokingstatusContext()
}

type SmokingstatusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySmokingstatusContext() *SmokingstatusContext {
	var p = new(SmokingstatusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SmokingStatusParserRULE_smokingstatus
	return p
}

func (*SmokingstatusContext) IsSmokingstatusContext() {}

func NewSmokingstatusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SmokingstatusContext {
	var p = new(SmokingstatusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SmokingStatusParserRULE_smokingstatus

	return p
}

func (s *SmokingstatusContext) GetParser() antlr.Parser { return s.parser }

func (s *SmokingstatusContext) SMOKINGSTATUS() antlr.TerminalNode {
	return s.GetToken(SmokingStatusParserSMOKINGSTATUS, 0)
}

func (s *SmokingstatusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SmokingstatusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SmokingstatusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokingStatusListener); ok {
		listenerT.EnterSmokingstatus(s)
	}
}

func (s *SmokingstatusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SmokingStatusListener); ok {
		listenerT.ExitSmokingstatus(s)
	}
}

func (p *SmokingStatusParser) Smokingstatus() (localctx ISmokingstatusContext) {
	this := p
	_ = this

	localctx = NewSmokingstatusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SmokingStatusParserRULE_smokingstatus)

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
		p.Match(SmokingStatusParserSMOKINGSTATUS)
	}

	return localctx
}
