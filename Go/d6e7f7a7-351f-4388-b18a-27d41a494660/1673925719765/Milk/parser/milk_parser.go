// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925719765/Milk.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Milk

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

type MilkParser struct {
	*antlr.BaseParser
}

var milkParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func milkParserInit() {
	staticData := &milkParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MILK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "milk",
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

// MilkParserInit initializes any static state used to implement MilkParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMilkParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MilkParserInit() {
	staticData := &milkParserStaticData
	staticData.once.Do(milkParserInit)
}

// NewMilkParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMilkParser(input antlr.TokenStream) *MilkParser {
	MilkParserInit()
	this := new(MilkParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &milkParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Milk.g4"

	return this
}

// MilkParser tokens.
const (
	MilkParserEOF      = antlr.TokenEOF
	MilkParserMILK     = 1
	MilkParserOWNKEY   = 2
	MilkParserSPLITKEY = 3
	MilkParserWS       = 4
)

// MilkParser rules.
const (
	MilkParserRULE_expression = 0
	MilkParserRULE_milk       = 1
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
	p.RuleIndex = MilkParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MilkParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Milk() IMilkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMilkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMilkContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MilkParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MilkListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MilkListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MilkParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MilkParserRULE_expression)

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
		p.Milk()
	}
	{
		p.SetState(5)
		p.Match(MilkParserEOF)
	}

	return localctx
}

// IMilkContext is an interface to support dynamic dispatch.
type IMilkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMilkContext differentiates from other interfaces.
	IsMilkContext()
}

type MilkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMilkContext() *MilkContext {
	var p = new(MilkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MilkParserRULE_milk
	return p
}

func (*MilkContext) IsMilkContext() {}

func NewMilkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MilkContext {
	var p = new(MilkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MilkParserRULE_milk

	return p
}

func (s *MilkContext) GetParser() antlr.Parser { return s.parser }

func (s *MilkContext) MILK() antlr.TerminalNode {
	return s.GetToken(MilkParserMILK, 0)
}

func (s *MilkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MilkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MilkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MilkListener); ok {
		listenerT.EnterMilk(s)
	}
}

func (s *MilkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MilkListener); ok {
		listenerT.ExitMilk(s)
	}
}

func (p *MilkParser) Milk() (localctx IMilkContext) {
	this := p
	_ = this

	localctx = NewMilkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MilkParserRULE_milk)

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
		p.Match(MilkParserMILK)
	}

	return localctx
}
