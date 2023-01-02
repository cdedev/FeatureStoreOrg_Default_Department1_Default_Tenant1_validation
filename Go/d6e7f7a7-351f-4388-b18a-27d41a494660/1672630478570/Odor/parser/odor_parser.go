// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672630478570/Odor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Odor

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

type OdorParser struct {
	*antlr.BaseParser
}

var odorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func odorParserInit() {
	staticData := &odorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ODOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "odor",
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

// OdorParserInit initializes any static state used to implement OdorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOdorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OdorParserInit() {
	staticData := &odorParserStaticData
	staticData.once.Do(odorParserInit)
}

// NewOdorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOdorParser(input antlr.TokenStream) *OdorParser {
	OdorParserInit()
	this := new(OdorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &odorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Odor.g4"

	return this
}

// OdorParser tokens.
const (
	OdorParserEOF      = antlr.TokenEOF
	OdorParserODOR     = 1
	OdorParserOWNKEY   = 2
	OdorParserSPLITKEY = 3
	OdorParserWS       = 4
)

// OdorParser rules.
const (
	OdorParserRULE_expression = 0
	OdorParserRULE_odor       = 1
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
	p.RuleIndex = OdorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Odor() IOdorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOdorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOdorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OdorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OdorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OdorParserRULE_expression)

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
		p.Odor()
	}
	{
		p.SetState(5)
		p.Match(OdorParserEOF)
	}

	return localctx
}

// IOdorContext is an interface to support dynamic dispatch.
type IOdorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOdorContext differentiates from other interfaces.
	IsOdorContext()
}

type OdorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOdorContext() *OdorContext {
	var p = new(OdorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OdorParserRULE_odor
	return p
}

func (*OdorContext) IsOdorContext() {}

func NewOdorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OdorContext {
	var p = new(OdorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OdorParserRULE_odor

	return p
}

func (s *OdorContext) GetParser() antlr.Parser { return s.parser }

func (s *OdorContext) ODOR() antlr.TerminalNode {
	return s.GetToken(OdorParserODOR, 0)
}

func (s *OdorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OdorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OdorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdorListener); ok {
		listenerT.EnterOdor(s)
	}
}

func (s *OdorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OdorListener); ok {
		listenerT.ExitOdor(s)
	}
}

func (p *OdorParser) Odor() (localctx IOdorContext) {
	this := p
	_ = this

	localctx = NewOdorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OdorParserRULE_odor)

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
		p.Match(OdorParserODOR)
	}

	return localctx
}
