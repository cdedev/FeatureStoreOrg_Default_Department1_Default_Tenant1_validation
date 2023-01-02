// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637967322/RPrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RPrice

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

type RPriceParser struct {
	*antlr.BaseParser
}

var rpriceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rpriceParserInit() {
	staticData := &rpriceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RPRICE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rprice",
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

// RPriceParserInit initializes any static state used to implement RPriceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRPriceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RPriceParserInit() {
	staticData := &rpriceParserStaticData
	staticData.once.Do(rpriceParserInit)
}

// NewRPriceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRPriceParser(input antlr.TokenStream) *RPriceParser {
	RPriceParserInit()
	this := new(RPriceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rpriceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RPrice.g4"

	return this
}

// RPriceParser tokens.
const (
	RPriceParserEOF      = antlr.TokenEOF
	RPriceParserRPRICE   = 1
	RPriceParserOWNKEY   = 2
	RPriceParserSPLITKEY = 3
	RPriceParserWS       = 4
)

// RPriceParser rules.
const (
	RPriceParserRULE_expression = 0
	RPriceParserRULE_rprice     = 1
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
	p.RuleIndex = RPriceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RPriceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rprice() IRpriceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpriceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpriceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RPriceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RPriceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RPriceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RPriceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RPriceParserRULE_expression)

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
		p.Rprice()
	}
	{
		p.SetState(5)
		p.Match(RPriceParserEOF)
	}

	return localctx
}

// IRpriceContext is an interface to support dynamic dispatch.
type IRpriceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRpriceContext differentiates from other interfaces.
	IsRpriceContext()
}

type RpriceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpriceContext() *RpriceContext {
	var p = new(RpriceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RPriceParserRULE_rprice
	return p
}

func (*RpriceContext) IsRpriceContext() {}

func NewRpriceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpriceContext {
	var p = new(RpriceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RPriceParserRULE_rprice

	return p
}

func (s *RpriceContext) GetParser() antlr.Parser { return s.parser }

func (s *RpriceContext) RPRICE() antlr.TerminalNode {
	return s.GetToken(RPriceParserRPRICE, 0)
}

func (s *RpriceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpriceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpriceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RPriceListener); ok {
		listenerT.EnterRprice(s)
	}
}

func (s *RpriceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RPriceListener); ok {
		listenerT.ExitRprice(s)
	}
}

func (p *RPriceParser) Rprice() (localctx IRpriceContext) {
	this := p
	_ = this

	localctx = NewRpriceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RPriceParserRULE_rprice)

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
		p.Match(RPriceParserRPRICE)
	}

	return localctx
}
