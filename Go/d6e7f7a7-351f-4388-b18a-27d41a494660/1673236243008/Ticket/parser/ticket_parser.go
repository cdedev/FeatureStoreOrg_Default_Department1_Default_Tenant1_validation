// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236243008/Ticket.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ticket

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

type TicketParser struct {
	*antlr.BaseParser
}

var ticketParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ticketParserInit() {
	staticData := &ticketParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TICKET", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ticket",
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

// TicketParserInit initializes any static state used to implement TicketParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTicketParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TicketParserInit() {
	staticData := &ticketParserStaticData
	staticData.once.Do(ticketParserInit)
}

// NewTicketParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTicketParser(input antlr.TokenStream) *TicketParser {
	TicketParserInit()
	this := new(TicketParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ticketParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Ticket.g4"

	return this
}

// TicketParser tokens.
const (
	TicketParserEOF      = antlr.TokenEOF
	TicketParserTICKET   = 1
	TicketParserOWNKEY   = 2
	TicketParserSPLITKEY = 3
	TicketParserWS       = 4
)

// TicketParser rules.
const (
	TicketParserRULE_expression = 0
	TicketParserRULE_ticket     = 1
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
	p.RuleIndex = TicketParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TicketParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ticket() ITicketContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITicketContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITicketContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TicketParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TicketListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TicketListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TicketParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TicketParserRULE_expression)

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
		p.Ticket()
	}
	{
		p.SetState(5)
		p.Match(TicketParserEOF)
	}

	return localctx
}

// ITicketContext is an interface to support dynamic dispatch.
type ITicketContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTicketContext differentiates from other interfaces.
	IsTicketContext()
}

type TicketContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTicketContext() *TicketContext {
	var p = new(TicketContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TicketParserRULE_ticket
	return p
}

func (*TicketContext) IsTicketContext() {}

func NewTicketContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TicketContext {
	var p = new(TicketContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TicketParserRULE_ticket

	return p
}

func (s *TicketContext) GetParser() antlr.Parser { return s.parser }

func (s *TicketContext) TICKET() antlr.TerminalNode {
	return s.GetToken(TicketParserTICKET, 0)
}

func (s *TicketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TicketContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TicketContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TicketListener); ok {
		listenerT.EnterTicket(s)
	}
}

func (s *TicketContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TicketListener); ok {
		listenerT.ExitTicket(s)
	}
}

func (p *TicketParser) Ticket() (localctx ITicketContext) {
	this := p
	_ = this

	localctx = NewTicketContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TicketParserRULE_ticket)

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
		p.Match(TicketParserTICKET)
	}

	return localctx
}
