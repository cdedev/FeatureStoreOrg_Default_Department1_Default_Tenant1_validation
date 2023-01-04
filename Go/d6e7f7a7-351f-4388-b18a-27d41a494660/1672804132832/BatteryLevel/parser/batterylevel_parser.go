// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672804132832/BatteryLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BatteryLevel

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

type BatteryLevelParser struct {
	*antlr.BaseParser
}

var batterylevelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func batterylevelParserInit() {
	staticData := &batterylevelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BATTERYLEVEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "batterylevel",
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

// BatteryLevelParserInit initializes any static state used to implement BatteryLevelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBatteryLevelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BatteryLevelParserInit() {
	staticData := &batterylevelParserStaticData
	staticData.once.Do(batterylevelParserInit)
}

// NewBatteryLevelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBatteryLevelParser(input antlr.TokenStream) *BatteryLevelParser {
	BatteryLevelParserInit()
	this := new(BatteryLevelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &batterylevelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BatteryLevel.g4"

	return this
}

// BatteryLevelParser tokens.
const (
	BatteryLevelParserEOF          = antlr.TokenEOF
	BatteryLevelParserBATTERYLEVEL = 1
	BatteryLevelParserOWNKEY       = 2
	BatteryLevelParserSPLITKEY     = 3
	BatteryLevelParserWS           = 4
)

// BatteryLevelParser rules.
const (
	BatteryLevelParserRULE_expression   = 0
	BatteryLevelParserRULE_batterylevel = 1
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
	p.RuleIndex = BatteryLevelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BatteryLevelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Batterylevel() IBatterylevelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBatterylevelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBatterylevelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BatteryLevelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryLevelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryLevelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BatteryLevelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BatteryLevelParserRULE_expression)

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
		p.Batterylevel()
	}
	{
		p.SetState(5)
		p.Match(BatteryLevelParserEOF)
	}

	return localctx
}

// IBatterylevelContext is an interface to support dynamic dispatch.
type IBatterylevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBatterylevelContext differentiates from other interfaces.
	IsBatterylevelContext()
}

type BatterylevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBatterylevelContext() *BatterylevelContext {
	var p = new(BatterylevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BatteryLevelParserRULE_batterylevel
	return p
}

func (*BatterylevelContext) IsBatterylevelContext() {}

func NewBatterylevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BatterylevelContext {
	var p = new(BatterylevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BatteryLevelParserRULE_batterylevel

	return p
}

func (s *BatterylevelContext) GetParser() antlr.Parser { return s.parser }

func (s *BatterylevelContext) BATTERYLEVEL() antlr.TerminalNode {
	return s.GetToken(BatteryLevelParserBATTERYLEVEL, 0)
}

func (s *BatterylevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BatterylevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BatterylevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryLevelListener); ok {
		listenerT.EnterBatterylevel(s)
	}
}

func (s *BatterylevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BatteryLevelListener); ok {
		listenerT.ExitBatterylevel(s)
	}
}

func (p *BatteryLevelParser) Batterylevel() (localctx IBatterylevelContext) {
	this := p
	_ = this

	localctx = NewBatterylevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BatteryLevelParserRULE_batterylevel)

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
		p.Match(BatteryLevelParserBATTERYLEVEL)
	}

	return localctx
}
