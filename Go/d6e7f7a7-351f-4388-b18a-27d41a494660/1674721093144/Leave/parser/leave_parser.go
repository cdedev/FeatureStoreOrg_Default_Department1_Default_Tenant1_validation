// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721093144/Leave.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Leave

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

type LeaveParser struct {
	*antlr.BaseParser
}

var leaveParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func leaveParserInit() {
	staticData := &leaveParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LEAVE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "leave",
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

// LeaveParserInit initializes any static state used to implement LeaveParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLeaveParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LeaveParserInit() {
	staticData := &leaveParserStaticData
	staticData.once.Do(leaveParserInit)
}

// NewLeaveParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLeaveParser(input antlr.TokenStream) *LeaveParser {
	LeaveParserInit()
	this := new(LeaveParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &leaveParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Leave.g4"

	return this
}

// LeaveParser tokens.
const (
	LeaveParserEOF      = antlr.TokenEOF
	LeaveParserLEAVE    = 1
	LeaveParserOWNKEY   = 2
	LeaveParserSPLITKEY = 3
	LeaveParserWS       = 4
)

// LeaveParser rules.
const (
	LeaveParserRULE_expression = 0
	LeaveParserRULE_leave      = 1
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
	p.RuleIndex = LeaveParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LeaveParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Leave() ILeaveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeaveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeaveContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LeaveParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeaveListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeaveListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LeaveParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LeaveParserRULE_expression)

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
		p.Leave()
	}
	{
		p.SetState(5)
		p.Match(LeaveParserEOF)
	}

	return localctx
}

// ILeaveContext is an interface to support dynamic dispatch.
type ILeaveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeaveContext differentiates from other interfaces.
	IsLeaveContext()
}

type LeaveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeaveContext() *LeaveContext {
	var p = new(LeaveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LeaveParserRULE_leave
	return p
}

func (*LeaveContext) IsLeaveContext() {}

func NewLeaveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeaveContext {
	var p = new(LeaveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LeaveParserRULE_leave

	return p
}

func (s *LeaveContext) GetParser() antlr.Parser { return s.parser }

func (s *LeaveContext) LEAVE() antlr.TerminalNode {
	return s.GetToken(LeaveParserLEAVE, 0)
}

func (s *LeaveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeaveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeaveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeaveListener); ok {
		listenerT.EnterLeave(s)
	}
}

func (s *LeaveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeaveListener); ok {
		listenerT.ExitLeave(s)
	}
}

func (p *LeaveParser) Leave() (localctx ILeaveContext) {
	this := p
	_ = this

	localctx = NewLeaveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LeaveParserRULE_leave)

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
		p.Match(LeaveParserLEAVE)
	}

	return localctx
}
