// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983008792/BasePrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BasePrice

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

type BasePriceParser struct {
	*antlr.BaseParser
}

var basepriceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func basepriceParserInit() {
	staticData := &basepriceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BASEPRICE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "baseprice",
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

// BasePriceParserInit initializes any static state used to implement BasePriceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBasePriceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BasePriceParserInit() {
	staticData := &basepriceParserStaticData
	staticData.once.Do(basepriceParserInit)
}

// NewBasePriceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBasePriceParser(input antlr.TokenStream) *BasePriceParser {
	BasePriceParserInit()
	this := new(BasePriceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &basepriceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BasePrice.g4"

	return this
}

// BasePriceParser tokens.
const (
	BasePriceParserEOF       = antlr.TokenEOF
	BasePriceParserBASEPRICE = 1
	BasePriceParserOWNKEY    = 2
	BasePriceParserSPLITKEY  = 3
	BasePriceParserWS        = 4
)

// BasePriceParser rules.
const (
	BasePriceParserRULE_expression = 0
	BasePriceParserRULE_baseprice  = 1
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
	p.RuleIndex = BasePriceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasePriceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Baseprice() IBasepriceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBasepriceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBasepriceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BasePriceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BasePriceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BasePriceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BasePriceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BasePriceParserRULE_expression)

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
		p.Baseprice()
	}
	{
		p.SetState(5)
		p.Match(BasePriceParserEOF)
	}

	return localctx
}

// IBasepriceContext is an interface to support dynamic dispatch.
type IBasepriceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBasepriceContext differentiates from other interfaces.
	IsBasepriceContext()
}

type BasepriceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBasepriceContext() *BasepriceContext {
	var p = new(BasepriceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BasePriceParserRULE_baseprice
	return p
}

func (*BasepriceContext) IsBasepriceContext() {}

func NewBasepriceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BasepriceContext {
	var p = new(BasepriceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BasePriceParserRULE_baseprice

	return p
}

func (s *BasepriceContext) GetParser() antlr.Parser { return s.parser }

func (s *BasepriceContext) BASEPRICE() antlr.TerminalNode {
	return s.GetToken(BasePriceParserBASEPRICE, 0)
}

func (s *BasepriceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BasepriceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BasepriceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BasePriceListener); ok {
		listenerT.EnterBaseprice(s)
	}
}

func (s *BasepriceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BasePriceListener); ok {
		listenerT.ExitBaseprice(s)
	}
}

func (p *BasePriceParser) Baseprice() (localctx IBasepriceContext) {
	this := p
	_ = this

	localctx = NewBasepriceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BasePriceParserRULE_baseprice)

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
		p.Match(BasePriceParserBASEPRICE)
	}

	return localctx
}
