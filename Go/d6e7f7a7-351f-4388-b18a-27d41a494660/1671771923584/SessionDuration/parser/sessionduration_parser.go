// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771923584/SessionDuration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SessionDuration

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

type SessionDurationParser struct {
	*antlr.BaseParser
}

var sessiondurationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sessiondurationParserInit() {
	staticData := &sessiondurationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SESSIONDURATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sessionduration",
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

// SessionDurationParserInit initializes any static state used to implement SessionDurationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSessionDurationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SessionDurationParserInit() {
	staticData := &sessiondurationParserStaticData
	staticData.once.Do(sessiondurationParserInit)
}

// NewSessionDurationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSessionDurationParser(input antlr.TokenStream) *SessionDurationParser {
	SessionDurationParserInit()
	this := new(SessionDurationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sessiondurationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SessionDuration.g4"

	return this
}

// SessionDurationParser tokens.
const (
	SessionDurationParserEOF             = antlr.TokenEOF
	SessionDurationParserSESSIONDURATION = 1
	SessionDurationParserOWNKEY          = 2
	SessionDurationParserSPLITKEY        = 3
	SessionDurationParserWS              = 4
)

// SessionDurationParser rules.
const (
	SessionDurationParserRULE_expression      = 0
	SessionDurationParserRULE_sessionduration = 1
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
	p.RuleIndex = SessionDurationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SessionDurationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sessionduration() ISessiondurationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISessiondurationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISessiondurationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SessionDurationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SessionDurationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SessionDurationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SessionDurationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SessionDurationParserRULE_expression)

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
		p.Sessionduration()
	}
	{
		p.SetState(5)
		p.Match(SessionDurationParserEOF)
	}

	return localctx
}

// ISessiondurationContext is an interface to support dynamic dispatch.
type ISessiondurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSessiondurationContext differentiates from other interfaces.
	IsSessiondurationContext()
}

type SessiondurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySessiondurationContext() *SessiondurationContext {
	var p = new(SessiondurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SessionDurationParserRULE_sessionduration
	return p
}

func (*SessiondurationContext) IsSessiondurationContext() {}

func NewSessiondurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SessiondurationContext {
	var p = new(SessiondurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SessionDurationParserRULE_sessionduration

	return p
}

func (s *SessiondurationContext) GetParser() antlr.Parser { return s.parser }

func (s *SessiondurationContext) SESSIONDURATION() antlr.TerminalNode {
	return s.GetToken(SessionDurationParserSESSIONDURATION, 0)
}

func (s *SessiondurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SessiondurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SessiondurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SessionDurationListener); ok {
		listenerT.EnterSessionduration(s)
	}
}

func (s *SessiondurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SessionDurationListener); ok {
		listenerT.ExitSessionduration(s)
	}
}

func (p *SessionDurationParser) Sessionduration() (localctx ISessiondurationContext) {
	this := p
	_ = this

	localctx = NewSessiondurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SessionDurationParserRULE_sessionduration)

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
		p.Match(SessionDurationParserSESSIONDURATION)
	}

	return localctx
}
