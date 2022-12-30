// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376599168/Gear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gear

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

type GearParser struct {
	*antlr.BaseParser
}

var gearParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gearParserInit() {
	staticData := &gearParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GEAR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gear",
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

// GearParserInit initializes any static state used to implement GearParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGearParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GearParserInit() {
	staticData := &gearParserStaticData
	staticData.once.Do(gearParserInit)
}

// NewGearParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGearParser(input antlr.TokenStream) *GearParser {
	GearParserInit()
	this := new(GearParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gearParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Gear.g4"

	return this
}

// GearParser tokens.
const (
	GearParserEOF      = antlr.TokenEOF
	GearParserGEAR     = 1
	GearParserOWNKEY   = 2
	GearParserSPLITKEY = 3
	GearParserWS       = 4
)

// GearParser rules.
const (
	GearParserRULE_expression = 0
	GearParserRULE_gear       = 1
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
	p.RuleIndex = GearParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GearParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gear() IGearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGearContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GearParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GearParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GearParserRULE_expression)

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
		p.Gear()
	}
	{
		p.SetState(5)
		p.Match(GearParserEOF)
	}

	return localctx
}

// IGearContext is an interface to support dynamic dispatch.
type IGearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGearContext differentiates from other interfaces.
	IsGearContext()
}

type GearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGearContext() *GearContext {
	var p = new(GearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GearParserRULE_gear
	return p
}

func (*GearContext) IsGearContext() {}

func NewGearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GearContext {
	var p = new(GearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GearParserRULE_gear

	return p
}

func (s *GearContext) GetParser() antlr.Parser { return s.parser }

func (s *GearContext) GEAR() antlr.TerminalNode {
	return s.GetToken(GearParserGEAR, 0)
}

func (s *GearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearListener); ok {
		listenerT.EnterGear(s)
	}
}

func (s *GearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearListener); ok {
		listenerT.ExitGear(s)
	}
}

func (p *GearParser) Gear() (localctx IGearContext) {
	this := p
	_ = this

	localctx = NewGearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GearParserRULE_gear)

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
		p.Match(GearParserGEAR)
	}

	return localctx
}
