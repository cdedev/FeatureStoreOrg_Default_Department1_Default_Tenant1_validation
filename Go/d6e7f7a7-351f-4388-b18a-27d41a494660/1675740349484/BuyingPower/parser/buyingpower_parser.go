// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675740349484/BuyingPower.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BuyingPower

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

type BuyingPowerParser struct {
	*antlr.BaseParser
}

var buyingpowerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func buyingpowerParserInit() {
	staticData := &buyingpowerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BUYINGPOWER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "buyingpower",
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

// BuyingPowerParserInit initializes any static state used to implement BuyingPowerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBuyingPowerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BuyingPowerParserInit() {
	staticData := &buyingpowerParserStaticData
	staticData.once.Do(buyingpowerParserInit)
}

// NewBuyingPowerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBuyingPowerParser(input antlr.TokenStream) *BuyingPowerParser {
	BuyingPowerParserInit()
	this := new(BuyingPowerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &buyingpowerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BuyingPower.g4"

	return this
}

// BuyingPowerParser tokens.
const (
	BuyingPowerParserEOF         = antlr.TokenEOF
	BuyingPowerParserBUYINGPOWER = 1
	BuyingPowerParserOWNKEY      = 2
	BuyingPowerParserSPLITKEY    = 3
	BuyingPowerParserWS          = 4
)

// BuyingPowerParser rules.
const (
	BuyingPowerParserRULE_expression  = 0
	BuyingPowerParserRULE_buyingpower = 1
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
	p.RuleIndex = BuyingPowerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BuyingPowerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Buyingpower() IBuyingpowerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuyingpowerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuyingpowerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BuyingPowerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BuyingPowerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BuyingPowerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BuyingPowerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BuyingPowerParserRULE_expression)

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
		p.Buyingpower()
	}
	{
		p.SetState(5)
		p.Match(BuyingPowerParserEOF)
	}

	return localctx
}

// IBuyingpowerContext is an interface to support dynamic dispatch.
type IBuyingpowerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuyingpowerContext differentiates from other interfaces.
	IsBuyingpowerContext()
}

type BuyingpowerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuyingpowerContext() *BuyingpowerContext {
	var p = new(BuyingpowerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BuyingPowerParserRULE_buyingpower
	return p
}

func (*BuyingpowerContext) IsBuyingpowerContext() {}

func NewBuyingpowerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuyingpowerContext {
	var p = new(BuyingpowerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BuyingPowerParserRULE_buyingpower

	return p
}

func (s *BuyingpowerContext) GetParser() antlr.Parser { return s.parser }

func (s *BuyingpowerContext) BUYINGPOWER() antlr.TerminalNode {
	return s.GetToken(BuyingPowerParserBUYINGPOWER, 0)
}

func (s *BuyingpowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuyingpowerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuyingpowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BuyingPowerListener); ok {
		listenerT.EnterBuyingpower(s)
	}
}

func (s *BuyingpowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BuyingPowerListener); ok {
		listenerT.ExitBuyingpower(s)
	}
}

func (p *BuyingPowerParser) Buyingpower() (localctx IBuyingpowerContext) {
	this := p
	_ = this

	localctx = NewBuyingpowerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BuyingPowerParserRULE_buyingpower)

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
		p.Match(BuyingPowerParserBUYINGPOWER)
	}

	return localctx
}
