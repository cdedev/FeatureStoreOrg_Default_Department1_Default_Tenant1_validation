// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675078615944/AmountOfMoney.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AmountOfMoney

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

type AmountOfMoneyParser struct {
	*antlr.BaseParser
}

var amountofmoneyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func amountofmoneyParserInit() {
	staticData := &amountofmoneyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AMOUNTOFMONEY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "amountofmoney",
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

// AmountOfMoneyParserInit initializes any static state used to implement AmountOfMoneyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAmountOfMoneyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AmountOfMoneyParserInit() {
	staticData := &amountofmoneyParserStaticData
	staticData.once.Do(amountofmoneyParserInit)
}

// NewAmountOfMoneyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAmountOfMoneyParser(input antlr.TokenStream) *AmountOfMoneyParser {
	AmountOfMoneyParserInit()
	this := new(AmountOfMoneyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &amountofmoneyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AmountOfMoney.g4"

	return this
}

// AmountOfMoneyParser tokens.
const (
	AmountOfMoneyParserEOF           = antlr.TokenEOF
	AmountOfMoneyParserAMOUNTOFMONEY = 1
	AmountOfMoneyParserOWNKEY        = 2
	AmountOfMoneyParserSPLITKEY      = 3
	AmountOfMoneyParserWS            = 4
)

// AmountOfMoneyParser rules.
const (
	AmountOfMoneyParserRULE_expression    = 0
	AmountOfMoneyParserRULE_amountofmoney = 1
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
	p.RuleIndex = AmountOfMoneyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmountOfMoneyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Amountofmoney() IAmountofmoneyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAmountofmoneyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAmountofmoneyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AmountOfMoneyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmountOfMoneyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmountOfMoneyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AmountOfMoneyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AmountOfMoneyParserRULE_expression)

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
		p.Amountofmoney()
	}
	{
		p.SetState(5)
		p.Match(AmountOfMoneyParserEOF)
	}

	return localctx
}

// IAmountofmoneyContext is an interface to support dynamic dispatch.
type IAmountofmoneyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmountofmoneyContext differentiates from other interfaces.
	IsAmountofmoneyContext()
}

type AmountofmoneyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmountofmoneyContext() *AmountofmoneyContext {
	var p = new(AmountofmoneyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AmountOfMoneyParserRULE_amountofmoney
	return p
}

func (*AmountofmoneyContext) IsAmountofmoneyContext() {}

func NewAmountofmoneyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmountofmoneyContext {
	var p = new(AmountofmoneyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmountOfMoneyParserRULE_amountofmoney

	return p
}

func (s *AmountofmoneyContext) GetParser() antlr.Parser { return s.parser }

func (s *AmountofmoneyContext) AMOUNTOFMONEY() antlr.TerminalNode {
	return s.GetToken(AmountOfMoneyParserAMOUNTOFMONEY, 0)
}

func (s *AmountofmoneyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmountofmoneyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AmountofmoneyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmountOfMoneyListener); ok {
		listenerT.EnterAmountofmoney(s)
	}
}

func (s *AmountofmoneyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmountOfMoneyListener); ok {
		listenerT.ExitAmountofmoney(s)
	}
}

func (p *AmountOfMoneyParser) Amountofmoney() (localctx IAmountofmoneyContext) {
	this := p
	_ = this

	localctx = NewAmountofmoneyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AmountOfMoneyParserRULE_amountofmoney)

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
		p.Match(AmountOfMoneyParserAMOUNTOFMONEY)
	}

	return localctx
}
