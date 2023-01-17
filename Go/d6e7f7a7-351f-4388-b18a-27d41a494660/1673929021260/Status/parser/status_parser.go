// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673929021260/Status.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Status

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

type StatusParser struct {
	*antlr.BaseParser
}

var statusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func statusParserInit() {
	staticData := &statusParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STATUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "status",
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

// StatusParserInit initializes any static state used to implement StatusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStatusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StatusParserInit() {
	staticData := &statusParserStaticData
	staticData.once.Do(statusParserInit)
}

// NewStatusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStatusParser(input antlr.TokenStream) *StatusParser {
	StatusParserInit()
	this := new(StatusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &statusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Status.g4"

	return this
}

// StatusParser tokens.
const (
	StatusParserEOF      = antlr.TokenEOF
	StatusParserSTATUS   = 1
	StatusParserOWNKEY   = 2
	StatusParserSPLITKEY = 3
	StatusParserWS       = 4
)

// StatusParser rules.
const (
	StatusParserRULE_expression = 0
	StatusParserRULE_status     = 1
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
	p.RuleIndex = StatusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StatusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Status() IStatusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatusContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StatusParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StatusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StatusParserRULE_expression)

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
		p.Status()
	}
	{
		p.SetState(5)
		p.Match(StatusParserEOF)
	}

	return localctx
}

// IStatusContext is an interface to support dynamic dispatch.
type IStatusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatusContext differentiates from other interfaces.
	IsStatusContext()
}

type StatusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatusContext() *StatusContext {
	var p = new(StatusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StatusParserRULE_status
	return p
}

func (*StatusContext) IsStatusContext() {}

func NewStatusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatusContext {
	var p = new(StatusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StatusParserRULE_status

	return p
}

func (s *StatusContext) GetParser() antlr.Parser { return s.parser }

func (s *StatusContext) STATUS() antlr.TerminalNode {
	return s.GetToken(StatusParserSTATUS, 0)
}

func (s *StatusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusListener); ok {
		listenerT.EnterStatus(s)
	}
}

func (s *StatusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusListener); ok {
		listenerT.ExitStatus(s)
	}
}

func (p *StatusParser) Status() (localctx IStatusContext) {
	this := p
	_ = this

	localctx = NewStatusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StatusParserRULE_status)

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
		p.Match(StatusParserSTATUS)
	}

	return localctx
}
