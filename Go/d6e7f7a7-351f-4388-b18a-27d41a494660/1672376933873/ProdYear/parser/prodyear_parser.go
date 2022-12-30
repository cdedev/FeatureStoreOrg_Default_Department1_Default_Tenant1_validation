// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376933873/ProdYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ProdYear

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

type ProdYearParser struct {
	*antlr.BaseParser
}

var prodyearParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func prodyearParserInit() {
	staticData := &prodyearParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PRODYEAR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "prodyear",
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

// ProdYearParserInit initializes any static state used to implement ProdYearParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProdYearParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProdYearParserInit() {
	staticData := &prodyearParserStaticData
	staticData.once.Do(prodyearParserInit)
}

// NewProdYearParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProdYearParser(input antlr.TokenStream) *ProdYearParser {
	ProdYearParserInit()
	this := new(ProdYearParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &prodyearParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ProdYear.g4"

	return this
}

// ProdYearParser tokens.
const (
	ProdYearParserEOF      = antlr.TokenEOF
	ProdYearParserPRODYEAR = 1
	ProdYearParserOWNKEY   = 2
	ProdYearParserSPLITKEY = 3
	ProdYearParserWS       = 4
)

// ProdYearParser rules.
const (
	ProdYearParserRULE_expression = 0
	ProdYearParserRULE_prodyear   = 1
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
	p.RuleIndex = ProdYearParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProdYearParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Prodyear() IProdyearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProdyearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProdyearContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProdYearParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProdYearListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProdYearListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ProdYearParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProdYearParserRULE_expression)

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
		p.Prodyear()
	}
	{
		p.SetState(5)
		p.Match(ProdYearParserEOF)
	}

	return localctx
}

// IProdyearContext is an interface to support dynamic dispatch.
type IProdyearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProdyearContext differentiates from other interfaces.
	IsProdyearContext()
}

type ProdyearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProdyearContext() *ProdyearContext {
	var p = new(ProdyearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProdYearParserRULE_prodyear
	return p
}

func (*ProdyearContext) IsProdyearContext() {}

func NewProdyearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProdyearContext {
	var p = new(ProdyearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProdYearParserRULE_prodyear

	return p
}

func (s *ProdyearContext) GetParser() antlr.Parser { return s.parser }

func (s *ProdyearContext) PRODYEAR() antlr.TerminalNode {
	return s.GetToken(ProdYearParserPRODYEAR, 0)
}

func (s *ProdyearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProdyearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProdyearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProdYearListener); ok {
		listenerT.EnterProdyear(s)
	}
}

func (s *ProdyearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProdYearListener); ok {
		listenerT.ExitProdyear(s)
	}
}

func (p *ProdYearParser) Prodyear() (localctx IProdyearContext) {
	this := p
	_ = this

	localctx = NewProdyearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProdYearParserRULE_prodyear)

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
		p.Match(ProdYearParserPRODYEAR)
	}

	return localctx
}
