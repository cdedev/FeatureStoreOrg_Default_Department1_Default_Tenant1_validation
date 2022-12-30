// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377665474/Wifi.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wifi

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

type WifiParser struct {
	*antlr.BaseParser
}

var wifiParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func wifiParserInit() {
	staticData := &wifiParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WIFI", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "wifi",
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

// WifiParserInit initializes any static state used to implement WifiParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWifiParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WifiParserInit() {
	staticData := &wifiParserStaticData
	staticData.once.Do(wifiParserInit)
}

// NewWifiParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWifiParser(input antlr.TokenStream) *WifiParser {
	WifiParserInit()
	this := new(WifiParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &wifiParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Wifi.g4"

	return this
}

// WifiParser tokens.
const (
	WifiParserEOF      = antlr.TokenEOF
	WifiParserWIFI     = 1
	WifiParserOWNKEY   = 2
	WifiParserSPLITKEY = 3
	WifiParserWS       = 4
)

// WifiParser rules.
const (
	WifiParserRULE_expression = 0
	WifiParserRULE_wifi       = 1
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
	p.RuleIndex = WifiParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WifiParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Wifi() IWifiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWifiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWifiContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WifiParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WifiListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WifiListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WifiParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WifiParserRULE_expression)

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
		p.Wifi()
	}
	{
		p.SetState(5)
		p.Match(WifiParserEOF)
	}

	return localctx
}

// IWifiContext is an interface to support dynamic dispatch.
type IWifiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWifiContext differentiates from other interfaces.
	IsWifiContext()
}

type WifiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWifiContext() *WifiContext {
	var p = new(WifiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WifiParserRULE_wifi
	return p
}

func (*WifiContext) IsWifiContext() {}

func NewWifiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WifiContext {
	var p = new(WifiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WifiParserRULE_wifi

	return p
}

func (s *WifiContext) GetParser() antlr.Parser { return s.parser }

func (s *WifiContext) WIFI() antlr.TerminalNode {
	return s.GetToken(WifiParserWIFI, 0)
}

func (s *WifiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WifiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WifiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WifiListener); ok {
		listenerT.EnterWifi(s)
	}
}

func (s *WifiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WifiListener); ok {
		listenerT.ExitWifi(s)
	}
}

func (p *WifiParser) Wifi() (localctx IWifiContext) {
	this := p
	_ = this

	localctx = NewWifiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WifiParserRULE_wifi)

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
		p.Match(WifiParserWIFI)
	}

	return localctx
}
