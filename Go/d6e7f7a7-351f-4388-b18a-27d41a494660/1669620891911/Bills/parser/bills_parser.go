// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669620891911/Bills.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bills

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

type BillsParser struct {
	*antlr.BaseParser
}

var billsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func billsParserInit() {
	staticData := &billsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BILLS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bills",
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

// BillsParserInit initializes any static state used to implement BillsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBillsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BillsParserInit() {
	staticData := &billsParserStaticData
	staticData.once.Do(billsParserInit)
}

// NewBillsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBillsParser(input antlr.TokenStream) *BillsParser {
	BillsParserInit()
	this := new(BillsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &billsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bills.g4"

	return this
}

// BillsParser tokens.
const (
	BillsParserEOF      = antlr.TokenEOF
	BillsParserBILLS    = 1
	BillsParserOWNKEY   = 2
	BillsParserSPLITKEY = 3
	BillsParserWS       = 4
)

// BillsParser rules.
const (
	BillsParserRULE_expression = 0
	BillsParserRULE_bills      = 1
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
	p.RuleIndex = BillsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BillsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bills() IBillsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBillsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBillsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BillsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BillsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BillsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BillsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BillsParserRULE_expression)

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
		p.Bills()
	}
	{
		p.SetState(5)
		p.Match(BillsParserEOF)
	}

	return localctx
}

// IBillsContext is an interface to support dynamic dispatch.
type IBillsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBillsContext differentiates from other interfaces.
	IsBillsContext()
}

type BillsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBillsContext() *BillsContext {
	var p = new(BillsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BillsParserRULE_bills
	return p
}

func (*BillsContext) IsBillsContext() {}

func NewBillsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BillsContext {
	var p = new(BillsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BillsParserRULE_bills

	return p
}

func (s *BillsContext) GetParser() antlr.Parser { return s.parser }

func (s *BillsContext) BILLS() antlr.TerminalNode {
	return s.GetToken(BillsParserBILLS, 0)
}

func (s *BillsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BillsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BillsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BillsListener); ok {
		listenerT.EnterBills(s)
	}
}

func (s *BillsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BillsListener); ok {
		listenerT.ExitBills(s)
	}
}

func (p *BillsParser) Bills() (localctx IBillsContext) {
	this := p
	_ = this

	localctx = NewBillsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BillsParserRULE_bills)

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
		p.Match(BillsParserBILLS)
	}

	return localctx
}
