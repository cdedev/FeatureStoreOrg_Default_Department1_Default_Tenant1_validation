// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672461844535/Battery.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Battery

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

type BatteryParser struct {
	*antlr.BaseParser
}

var batteryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func batteryParserInit() {
	staticData := &batteryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BATTERY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "battery",
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

// BatteryParserInit initializes any static state used to implement BatteryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBatteryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BatteryParserInit() {
	staticData := &batteryParserStaticData
	staticData.once.Do(batteryParserInit)
}

// NewBatteryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBatteryParser(input antlr.TokenStream) *BatteryParser {
	BatteryParserInit()
	this := new(BatteryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &batteryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Battery.g4"

	return this
}

// BatteryParser tokens.
const (
	BatteryParserEOF      = antlr.TokenEOF
	BatteryParserBATTERY  = 1
	BatteryParserOWNKEY   = 2
	BatteryParserSPLITKEY = 3
	BatteryParserWS       = 4
)

// BatteryParser rules.
const (
	BatteryParserRULE_expression = 0
	BatteryParserRULE_battery    = 1
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
	p.RuleIndex = BatteryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BatteryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Battery() IBatteryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBatteryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBatteryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BatteryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BatteryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BatteryParserRULE_expression)

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
		p.Battery()
	}
	{
		p.SetState(5)
		p.Match(BatteryParserEOF)
	}

	return localctx
}

// IBatteryContext is an interface to support dynamic dispatch.
type IBatteryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatteryContext differentiates from other interfaces.
	IsBatteryContext()
}

type BatteryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatteryContext() *BatteryContext {
	var p = new(BatteryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BatteryParserRULE_battery
	return p
}

func (*BatteryContext) IsBatteryContext() {}

func NewBatteryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BatteryContext {
	var p = new(BatteryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BatteryParserRULE_battery

	return p
}

func (s *BatteryContext) GetParser() antlr.Parser { return s.parser }

func (s *BatteryContext) BATTERY() antlr.TerminalNode {
	return s.GetToken(BatteryParserBATTERY, 0)
}

func (s *BatteryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BatteryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BatteryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryListener); ok {
		listenerT.EnterBattery(s)
	}
}

func (s *BatteryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryListener); ok {
		listenerT.ExitBattery(s)
	}
}

func (p *BatteryParser) Battery() (localctx IBatteryContext) {
	this := p
	_ = this

	localctx = NewBatteryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BatteryParserRULE_battery)

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
		p.Match(BatteryParserBATTERY)
	}

	return localctx
}
