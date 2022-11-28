// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652405435/Failure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Failure

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

type FailureParser struct {
	*antlr.BaseParser
}

var failureParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func failureParserInit() {
	staticData := &failureParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FAILURE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "failure",
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

// FailureParserInit initializes any static state used to implement FailureParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFailureParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FailureParserInit() {
	staticData := &failureParserStaticData
	staticData.once.Do(failureParserInit)
}

// NewFailureParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFailureParser(input antlr.TokenStream) *FailureParser {
	FailureParserInit()
	this := new(FailureParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &failureParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Failure.g4"

	return this
}

// FailureParser tokens.
const (
	FailureParserEOF      = antlr.TokenEOF
	FailureParserFAILURE  = 1
	FailureParserOWNKEY   = 2
	FailureParserSPLITKEY = 3
	FailureParserWS       = 4
)

// FailureParser rules.
const (
	FailureParserRULE_expression = 0
	FailureParserRULE_failure    = 1
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
	p.RuleIndex = FailureParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FailureParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Failure() IFailureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFailureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFailureContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FailureParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FailureListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FailureListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FailureParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FailureParserRULE_expression)

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
		p.Failure()
	}
	{
		p.SetState(5)
		p.Match(FailureParserEOF)
	}

	return localctx
}

// IFailureContext is an interface to support dynamic dispatch.
type IFailureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFailureContext differentiates from other interfaces.
	IsFailureContext()
}

type FailureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFailureContext() *FailureContext {
	var p = new(FailureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FailureParserRULE_failure
	return p
}

func (*FailureContext) IsFailureContext() {}

func NewFailureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FailureContext {
	var p = new(FailureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FailureParserRULE_failure

	return p
}

func (s *FailureContext) GetParser() antlr.Parser { return s.parser }

func (s *FailureContext) FAILURE() antlr.TerminalNode {
	return s.GetToken(FailureParserFAILURE, 0)
}

func (s *FailureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FailureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FailureListener); ok {
		listenerT.EnterFailure(s)
	}
}

func (s *FailureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FailureListener); ok {
		listenerT.ExitFailure(s)
	}
}

func (p *FailureParser) Failure() (localctx IFailureContext) {
	this := p
	_ = this

	localctx = NewFailureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FailureParserRULE_failure)

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
		p.Match(FailureParserFAILURE)
	}

	return localctx
}
