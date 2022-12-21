// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671595329048/ModelYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ModelYear

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

type ModelYearParser struct {
	*antlr.BaseParser
}

var modelyearParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func modelyearParserInit() {
	staticData := &modelyearParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FIELD0", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "field0",
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

// ModelYearParserInit initializes any static state used to implement ModelYearParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewModelYearParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ModelYearParserInit() {
	staticData := &modelyearParserStaticData
	staticData.once.Do(modelyearParserInit)
}

// NewModelYearParser produces a new parser instance for the optional input antlr.TokenStream.
func NewModelYearParser(input antlr.TokenStream) *ModelYearParser {
	ModelYearParserInit()
	this := new(ModelYearParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &modelyearParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ModelYear.g4"

	return this
}

// ModelYearParser tokens.
const (
	ModelYearParserEOF      = antlr.TokenEOF
	ModelYearParserFIELD0   = 1
	ModelYearParserOWNKEY   = 2
	ModelYearParserSPLITKEY = 3
	ModelYearParserWS       = 4
)

// ModelYearParser rules.
const (
	ModelYearParserRULE_expression = 0
	ModelYearParserRULE_field0     = 1
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
	p.RuleIndex = ModelYearParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelYearParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Field0() IField0Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IField0Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IField0Context)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ModelYearParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelYearListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelYearListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ModelYearParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ModelYearParserRULE_expression)

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
		p.Field0()
	}
	{
		p.SetState(5)
		p.Match(ModelYearParserEOF)
	}

	return localctx
}

// IField0Context is an interface to support dynamic dispatch.
type IField0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField0Context differentiates from other interfaces.
	IsField0Context()
}

type Field0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField0Context() *Field0Context {
	var p = new(Field0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelYearParserRULE_field0
	return p
}

func (*Field0Context) IsField0Context() {}

func NewField0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field0Context {
	var p = new(Field0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelYearParserRULE_field0

	return p
}

func (s *Field0Context) GetParser() antlr.Parser { return s.parser }

func (s *Field0Context) FIELD0() antlr.TerminalNode {
	return s.GetToken(ModelYearParserFIELD0, 0)
}

func (s *Field0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelYearListener); ok {
		listenerT.EnterField0(s)
	}
}

func (s *Field0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelYearListener); ok {
		listenerT.ExitField0(s)
	}
}

func (p *ModelYearParser) Field0() (localctx IField0Context) {
	this := p
	_ = this

	localctx = NewField0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ModelYearParserRULE_field0)

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
		p.Match(ModelYearParserFIELD0)
	}

	return localctx
}
