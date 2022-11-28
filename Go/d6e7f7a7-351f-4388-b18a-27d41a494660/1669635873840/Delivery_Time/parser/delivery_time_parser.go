// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669635873840/Delivery_Time.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Delivery_Time

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

type Delivery_TimeParser struct {
	*antlr.BaseParser
}

var delivery_timeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func delivery_timeParserInit() {
	staticData := &delivery_timeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DELIVERY_TIME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "delivery_time",
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

// Delivery_TimeParserInit initializes any static state used to implement Delivery_TimeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDelivery_TimeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Delivery_TimeParserInit() {
	staticData := &delivery_timeParserStaticData
	staticData.once.Do(delivery_timeParserInit)
}

// NewDelivery_TimeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDelivery_TimeParser(input antlr.TokenStream) *Delivery_TimeParser {
	Delivery_TimeParserInit()
	this := new(Delivery_TimeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &delivery_timeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Delivery_Time.g4"

	return this
}

// Delivery_TimeParser tokens.
const (
	Delivery_TimeParserEOF           = antlr.TokenEOF
	Delivery_TimeParserDELIVERY_TIME = 1
	Delivery_TimeParserOWNKEY        = 2
	Delivery_TimeParserSPLITKEY      = 3
	Delivery_TimeParserWS            = 4
)

// Delivery_TimeParser rules.
const (
	Delivery_TimeParserRULE_expression    = 0
	Delivery_TimeParserRULE_delivery_time = 1
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
	p.RuleIndex = Delivery_TimeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Delivery_TimeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Delivery_time() IDelivery_timeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDelivery_timeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDelivery_timeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Delivery_TimeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Delivery_TimeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Delivery_TimeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Delivery_TimeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Delivery_TimeParserRULE_expression)

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
		p.Delivery_time()
	}
	{
		p.SetState(5)
		p.Match(Delivery_TimeParserEOF)
	}

	return localctx
}

// IDelivery_timeContext is an interface to support dynamic dispatch.
type IDelivery_timeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDelivery_timeContext differentiates from other interfaces.
	IsDelivery_timeContext()
}

type Delivery_timeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDelivery_timeContext() *Delivery_timeContext {
	var p = new(Delivery_timeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Delivery_TimeParserRULE_delivery_time
	return p
}

func (*Delivery_timeContext) IsDelivery_timeContext() {}

func NewDelivery_timeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Delivery_timeContext {
	var p = new(Delivery_timeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Delivery_TimeParserRULE_delivery_time

	return p
}

func (s *Delivery_timeContext) GetParser() antlr.Parser { return s.parser }

func (s *Delivery_timeContext) DELIVERY_TIME() antlr.TerminalNode {
	return s.GetToken(Delivery_TimeParserDELIVERY_TIME, 0)
}

func (s *Delivery_timeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Delivery_timeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Delivery_timeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Delivery_TimeListener); ok {
		listenerT.EnterDelivery_time(s)
	}
}

func (s *Delivery_timeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Delivery_TimeListener); ok {
		listenerT.ExitDelivery_time(s)
	}
}

func (p *Delivery_TimeParser) Delivery_time() (localctx IDelivery_timeContext) {
	this := p
	_ = this

	localctx = NewDelivery_timeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Delivery_TimeParserRULE_delivery_time)

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
		p.Match(Delivery_TimeParserDELIVERY_TIME)
	}

	return localctx
}
