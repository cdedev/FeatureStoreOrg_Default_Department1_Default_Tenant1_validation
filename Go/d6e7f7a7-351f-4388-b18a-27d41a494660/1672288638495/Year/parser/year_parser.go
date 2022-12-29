// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288638495/Year.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Year

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

type YearParser struct {
	*antlr.BaseParser
}

var yearParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yearParserInit() {
	staticData := &yearParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "YEAR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "year",
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

// YearParserInit initializes any static state used to implement YearParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYearParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YearParserInit() {
	staticData := &yearParserStaticData
	staticData.once.Do(yearParserInit)
}

// NewYearParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYearParser(input antlr.TokenStream) *YearParser {
	YearParserInit()
	this := new(YearParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &yearParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Year.g4"

	return this
}

// YearParser tokens.
const (
	YearParserEOF      = antlr.TokenEOF
	YearParserYEAR     = 1
	YearParserOWNKEY   = 2
	YearParserSPLITKEY = 3
	YearParserWS       = 4
)

// YearParser rules.
const (
	YearParserRULE_expression = 0
	YearParserRULE_year       = 1
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
	p.RuleIndex = YearParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YearParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Year() IYearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(YearParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *YearParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YearParserRULE_expression)

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
		p.Year()
	}
	{
		p.SetState(5)
		p.Match(YearParserEOF)
	}

	return localctx
}

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YearParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YearParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) YEAR() antlr.TerminalNode {
	return s.GetToken(YearParserYEAR, 0)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearListener); ok {
		listenerT.EnterYear(s)
	}
}

func (s *YearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearListener); ok {
		listenerT.ExitYear(s)
	}
}

func (p *YearParser) Year() (localctx IYearContext) {
	this := p
	_ = this

	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YearParserRULE_year)

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
		p.Match(YearParserYEAR)
	}

	return localctx
}
