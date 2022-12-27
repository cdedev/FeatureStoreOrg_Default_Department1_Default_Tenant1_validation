// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119101300/TotalAirTime.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalAirTime

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

type TotalAirTimeParser struct {
	*antlr.BaseParser
}

var totalairtimeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func totalairtimeParserInit() {
	staticData := &totalairtimeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TOTALAIRTIME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "totalairtime",
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

// TotalAirTimeParserInit initializes any static state used to implement TotalAirTimeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTotalAirTimeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TotalAirTimeParserInit() {
	staticData := &totalairtimeParserStaticData
	staticData.once.Do(totalairtimeParserInit)
}

// NewTotalAirTimeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTotalAirTimeParser(input antlr.TokenStream) *TotalAirTimeParser {
	TotalAirTimeParserInit()
	this := new(TotalAirTimeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &totalairtimeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "TotalAirTime.g4"

	return this
}

// TotalAirTimeParser tokens.
const (
	TotalAirTimeParserEOF          = antlr.TokenEOF
	TotalAirTimeParserTOTALAIRTIME = 1
	TotalAirTimeParserOWNKEY       = 2
	TotalAirTimeParserSPLITKEY     = 3
	TotalAirTimeParserWS           = 4
)

// TotalAirTimeParser rules.
const (
	TotalAirTimeParserRULE_expression   = 0
	TotalAirTimeParserRULE_totalairtime = 1
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
	p.RuleIndex = TotalAirTimeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TotalAirTimeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Totalairtime() ITotalairtimeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITotalairtimeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITotalairtimeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TotalAirTimeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalAirTimeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalAirTimeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TotalAirTimeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TotalAirTimeParserRULE_expression)

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
		p.Totalairtime()
	}
	{
		p.SetState(5)
		p.Match(TotalAirTimeParserEOF)
	}

	return localctx
}

// ITotalairtimeContext is an interface to support dynamic dispatch.
type ITotalairtimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTotalairtimeContext differentiates from other interfaces.
	IsTotalairtimeContext()
}

type TotalairtimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTotalairtimeContext() *TotalairtimeContext {
	var p = new(TotalairtimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TotalAirTimeParserRULE_totalairtime
	return p
}

func (*TotalairtimeContext) IsTotalairtimeContext() {}

func NewTotalairtimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotalairtimeContext {
	var p = new(TotalairtimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TotalAirTimeParserRULE_totalairtime

	return p
}

func (s *TotalairtimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TotalairtimeContext) TOTALAIRTIME() antlr.TerminalNode {
	return s.GetToken(TotalAirTimeParserTOTALAIRTIME, 0)
}

func (s *TotalairtimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalairtimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TotalairtimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalAirTimeListener); ok {
		listenerT.EnterTotalairtime(s)
	}
}

func (s *TotalairtimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalAirTimeListener); ok {
		listenerT.ExitTotalairtime(s)
	}
}

func (p *TotalAirTimeParser) Totalairtime() (localctx ITotalairtimeContext) {
	this := p
	_ = this

	localctx = NewTotalairtimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TotalAirTimeParserRULE_totalairtime)

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
		p.Match(TotalAirTimeParserTOTALAIRTIME)
	}

	return localctx
}
