// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516172268/MaritalStatus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaritalStatus

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

type MaritalStatusParser struct {
	*antlr.BaseParser
}

var maritalstatusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func maritalstatusParserInit() {
	staticData := &maritalstatusParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MARITALSTATUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "maritalstatus",
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

// MaritalStatusParserInit initializes any static state used to implement MaritalStatusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMaritalStatusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MaritalStatusParserInit() {
	staticData := &maritalstatusParserStaticData
	staticData.once.Do(maritalstatusParserInit)
}

// NewMaritalStatusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMaritalStatusParser(input antlr.TokenStream) *MaritalStatusParser {
	MaritalStatusParserInit()
	this := new(MaritalStatusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &maritalstatusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MaritalStatus.g4"

	return this
}

// MaritalStatusParser tokens.
const (
	MaritalStatusParserEOF           = antlr.TokenEOF
	MaritalStatusParserMARITALSTATUS = 1
	MaritalStatusParserOWNKEY        = 2
	MaritalStatusParserSPLITKEY      = 3
	MaritalStatusParserWS            = 4
)

// MaritalStatusParser rules.
const (
	MaritalStatusParserRULE_expression    = 0
	MaritalStatusParserRULE_maritalstatus = 1
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
	p.RuleIndex = MaritalStatusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaritalStatusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Maritalstatus() IMaritalstatusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaritalstatusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaritalstatusContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MaritalStatusParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaritalStatusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaritalStatusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MaritalStatusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MaritalStatusParserRULE_expression)

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
		p.Maritalstatus()
	}
	{
		p.SetState(5)
		p.Match(MaritalStatusParserEOF)
	}

	return localctx
}

// IMaritalstatusContext is an interface to support dynamic dispatch.
type IMaritalstatusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaritalstatusContext differentiates from other interfaces.
	IsMaritalstatusContext()
}

type MaritalstatusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaritalstatusContext() *MaritalstatusContext {
	var p = new(MaritalstatusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MaritalStatusParserRULE_maritalstatus
	return p
}

func (*MaritalstatusContext) IsMaritalstatusContext() {}

func NewMaritalstatusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaritalstatusContext {
	var p = new(MaritalstatusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaritalStatusParserRULE_maritalstatus

	return p
}

func (s *MaritalstatusContext) GetParser() antlr.Parser { return s.parser }

func (s *MaritalstatusContext) MARITALSTATUS() antlr.TerminalNode {
	return s.GetToken(MaritalStatusParserMARITALSTATUS, 0)
}

func (s *MaritalstatusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaritalstatusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaritalstatusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaritalStatusListener); ok {
		listenerT.EnterMaritalstatus(s)
	}
}

func (s *MaritalstatusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaritalStatusListener); ok {
		listenerT.ExitMaritalstatus(s)
	}
}

func (p *MaritalStatusParser) Maritalstatus() (localctx IMaritalstatusContext) {
	this := p
	_ = this

	localctx = NewMaritalstatusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MaritalStatusParserRULE_maritalstatus)

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
		p.Match(MaritalStatusParserMARITALSTATUS)
	}

	return localctx
}
