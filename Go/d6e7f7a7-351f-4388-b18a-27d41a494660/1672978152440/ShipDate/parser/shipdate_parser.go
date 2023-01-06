// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672978152440/ShipDate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ShipDate

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

type ShipDateParser struct {
	*antlr.BaseParser
}

var shipdateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func shipdateParserInit() {
	staticData := &shipdateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SHIPDATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "shipdate",
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

// ShipDateParserInit initializes any static state used to implement ShipDateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewShipDateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ShipDateParserInit() {
	staticData := &shipdateParserStaticData
	staticData.once.Do(shipdateParserInit)
}

// NewShipDateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewShipDateParser(input antlr.TokenStream) *ShipDateParser {
	ShipDateParserInit()
	this := new(ShipDateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &shipdateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ShipDate.g4"

	return this
}

// ShipDateParser tokens.
const (
	ShipDateParserEOF      = antlr.TokenEOF
	ShipDateParserSHIPDATE = 1
	ShipDateParserOWNKEY   = 2
	ShipDateParserSPLITKEY = 3
	ShipDateParserWS       = 4
)

// ShipDateParser rules.
const (
	ShipDateParserRULE_expression = 0
	ShipDateParserRULE_shipdate   = 1
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
	p.RuleIndex = ShipDateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShipDateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Shipdate() IShipdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShipdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShipdateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ShipDateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShipDateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShipDateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ShipDateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ShipDateParserRULE_expression)

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
		p.Shipdate()
	}
	{
		p.SetState(5)
		p.Match(ShipDateParserEOF)
	}

	return localctx
}

// IShipdateContext is an interface to support dynamic dispatch.
type IShipdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShipdateContext differentiates from other interfaces.
	IsShipdateContext()
}

type ShipdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShipdateContext() *ShipdateContext {
	var p = new(ShipdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShipDateParserRULE_shipdate
	return p
}

func (*ShipdateContext) IsShipdateContext() {}

func NewShipdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShipdateContext {
	var p = new(ShipdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShipDateParserRULE_shipdate

	return p
}

func (s *ShipdateContext) GetParser() antlr.Parser { return s.parser }

func (s *ShipdateContext) SHIPDATE() antlr.TerminalNode {
	return s.GetToken(ShipDateParserSHIPDATE, 0)
}

func (s *ShipdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShipdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShipdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShipDateListener); ok {
		listenerT.EnterShipdate(s)
	}
}

func (s *ShipdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShipDateListener); ok {
		listenerT.ExitShipdate(s)
	}
}

func (p *ShipDateParser) Shipdate() (localctx IShipdateContext) {
	this := p
	_ = this

	localctx = NewShipdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ShipDateParserRULE_shipdate)

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
		p.Match(ShipDateParserSHIPDATE)
	}

	return localctx
}
