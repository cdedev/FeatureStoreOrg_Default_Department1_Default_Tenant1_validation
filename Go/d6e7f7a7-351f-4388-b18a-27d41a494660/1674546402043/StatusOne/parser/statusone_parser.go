// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674546402043/StatusOne.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StatusOne

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

type StatusOneParser struct {
	*antlr.BaseParser
}

var statusoneParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func statusoneParserInit() {
	staticData := &statusoneParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STATUSONE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "statusone",
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

// StatusOneParserInit initializes any static state used to implement StatusOneParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStatusOneParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StatusOneParserInit() {
	staticData := &statusoneParserStaticData
	staticData.once.Do(statusoneParserInit)
}

// NewStatusOneParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStatusOneParser(input antlr.TokenStream) *StatusOneParser {
	StatusOneParserInit()
	this := new(StatusOneParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &statusoneParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "StatusOne.g4"

	return this
}

// StatusOneParser tokens.
const (
	StatusOneParserEOF       = antlr.TokenEOF
	StatusOneParserSTATUSONE = 1
	StatusOneParserOWNKEY    = 2
	StatusOneParserSPLITKEY  = 3
	StatusOneParserWS        = 4
)

// StatusOneParser rules.
const (
	StatusOneParserRULE_expression = 0
	StatusOneParserRULE_statusone  = 1
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
	p.RuleIndex = StatusOneParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StatusOneParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Statusone() IStatusoneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatusoneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatusoneContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StatusOneParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusOneListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusOneListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StatusOneParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StatusOneParserRULE_expression)

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
		p.Statusone()
	}
	{
		p.SetState(5)
		p.Match(StatusOneParserEOF)
	}

	return localctx
}

// IStatusoneContext is an interface to support dynamic dispatch.
type IStatusoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatusoneContext differentiates from other interfaces.
	IsStatusoneContext()
}

type StatusoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatusoneContext() *StatusoneContext {
	var p = new(StatusoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StatusOneParserRULE_statusone
	return p
}

func (*StatusoneContext) IsStatusoneContext() {}

func NewStatusoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatusoneContext {
	var p = new(StatusoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StatusOneParserRULE_statusone

	return p
}

func (s *StatusoneContext) GetParser() antlr.Parser { return s.parser }

func (s *StatusoneContext) STATUSONE() antlr.TerminalNode {
	return s.GetToken(StatusOneParserSTATUSONE, 0)
}

func (s *StatusoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatusoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatusoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusOneListener); ok {
		listenerT.EnterStatusone(s)
	}
}

func (s *StatusoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatusOneListener); ok {
		listenerT.ExitStatusone(s)
	}
}

func (p *StatusOneParser) Statusone() (localctx IStatusoneContext) {
	this := p
	_ = this

	localctx = NewStatusoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StatusOneParserRULE_statusone)

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
		p.Match(StatusOneParserSTATUSONE)
	}

	return localctx
}
