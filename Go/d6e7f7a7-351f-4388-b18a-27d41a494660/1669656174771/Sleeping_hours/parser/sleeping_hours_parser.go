// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656174771/Sleeping_hours.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sleeping_hours

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

type Sleeping_hoursParser struct {
	*antlr.BaseParser
}

var sleeping_hoursParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sleeping_hoursParserInit() {
	staticData := &sleeping_hoursParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SLEEPING_HOURS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sleeping_hours",
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

// Sleeping_hoursParserInit initializes any static state used to implement Sleeping_hoursParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSleeping_hoursParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Sleeping_hoursParserInit() {
	staticData := &sleeping_hoursParserStaticData
	staticData.once.Do(sleeping_hoursParserInit)
}

// NewSleeping_hoursParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSleeping_hoursParser(input antlr.TokenStream) *Sleeping_hoursParser {
	Sleeping_hoursParserInit()
	this := new(Sleeping_hoursParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sleeping_hoursParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sleeping_hours.g4"

	return this
}

// Sleeping_hoursParser tokens.
const (
	Sleeping_hoursParserEOF            = antlr.TokenEOF
	Sleeping_hoursParserSLEEPING_HOURS = 1
	Sleeping_hoursParserOWNKEY         = 2
	Sleeping_hoursParserSPLITKEY       = 3
	Sleeping_hoursParserWS             = 4
)

// Sleeping_hoursParser rules.
const (
	Sleeping_hoursParserRULE_expression     = 0
	Sleeping_hoursParserRULE_sleeping_hours = 1
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
	p.RuleIndex = Sleeping_hoursParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Sleeping_hoursParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sleeping_hours() ISleeping_hoursContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISleeping_hoursContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISleeping_hoursContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Sleeping_hoursParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Sleeping_hoursListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Sleeping_hoursListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Sleeping_hoursParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Sleeping_hoursParserRULE_expression)

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
		p.Sleeping_hours()
	}
	{
		p.SetState(5)
		p.Match(Sleeping_hoursParserEOF)
	}

	return localctx
}

// ISleeping_hoursContext is an interface to support dynamic dispatch.
type ISleeping_hoursContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSleeping_hoursContext differentiates from other interfaces.
	IsSleeping_hoursContext()
}

type Sleeping_hoursContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySleeping_hoursContext() *Sleeping_hoursContext {
	var p = new(Sleeping_hoursContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Sleeping_hoursParserRULE_sleeping_hours
	return p
}

func (*Sleeping_hoursContext) IsSleeping_hoursContext() {}

func NewSleeping_hoursContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sleeping_hoursContext {
	var p = new(Sleeping_hoursContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Sleeping_hoursParserRULE_sleeping_hours

	return p
}

func (s *Sleeping_hoursContext) GetParser() antlr.Parser { return s.parser }

func (s *Sleeping_hoursContext) SLEEPING_HOURS() antlr.TerminalNode {
	return s.GetToken(Sleeping_hoursParserSLEEPING_HOURS, 0)
}

func (s *Sleeping_hoursContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sleeping_hoursContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sleeping_hoursContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Sleeping_hoursListener); ok {
		listenerT.EnterSleeping_hours(s)
	}
}

func (s *Sleeping_hoursContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Sleeping_hoursListener); ok {
		listenerT.ExitSleeping_hours(s)
	}
}

func (p *Sleeping_hoursParser) Sleeping_hours() (localctx ISleeping_hoursContext) {
	this := p
	_ = this

	localctx = NewSleeping_hoursContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Sleeping_hoursParserRULE_sleeping_hours)

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
		p.Match(Sleeping_hoursParserSLEEPING_HOURS)
	}

	return localctx
}
