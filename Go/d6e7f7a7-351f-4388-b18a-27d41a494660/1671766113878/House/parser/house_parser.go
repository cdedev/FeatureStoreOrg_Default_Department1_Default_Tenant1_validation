// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671766113878/House.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // House

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

type HouseParser struct {
	*antlr.BaseParser
}

var houseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func houseParserInit() {
	staticData := &houseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HOUSE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "house",
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

// HouseParserInit initializes any static state used to implement HouseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHouseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HouseParserInit() {
	staticData := &houseParserStaticData
	staticData.once.Do(houseParserInit)
}

// NewHouseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHouseParser(input antlr.TokenStream) *HouseParser {
	HouseParserInit()
	this := new(HouseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &houseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "House.g4"

	return this
}

// HouseParser tokens.
const (
	HouseParserEOF      = antlr.TokenEOF
	HouseParserHOUSE    = 1
	HouseParserOWNKEY   = 2
	HouseParserSPLITKEY = 3
	HouseParserWS       = 4
)

// HouseParser rules.
const (
	HouseParserRULE_expression = 0
	HouseParserRULE_house      = 1
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
	p.RuleIndex = HouseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HouseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) House() IHouseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHouseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHouseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HouseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HouseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HouseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HouseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HouseParserRULE_expression)

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
		p.House()
	}
	{
		p.SetState(5)
		p.Match(HouseParserEOF)
	}

	return localctx
}

// IHouseContext is an interface to support dynamic dispatch.
type IHouseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHouseContext differentiates from other interfaces.
	IsHouseContext()
}

type HouseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHouseContext() *HouseContext {
	var p = new(HouseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HouseParserRULE_house
	return p
}

func (*HouseContext) IsHouseContext() {}

func NewHouseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HouseContext {
	var p = new(HouseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HouseParserRULE_house

	return p
}

func (s *HouseContext) GetParser() antlr.Parser { return s.parser }

func (s *HouseContext) HOUSE() antlr.TerminalNode {
	return s.GetToken(HouseParserHOUSE, 0)
}

func (s *HouseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HouseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HouseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HouseListener); ok {
		listenerT.EnterHouse(s)
	}
}

func (s *HouseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HouseListener); ok {
		listenerT.ExitHouse(s)
	}
}

func (p *HouseParser) House() (localctx IHouseContext) {
	this := p
	_ = this

	localctx = NewHouseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HouseParserRULE_house)

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
		p.Match(HouseParserHOUSE)
	}

	return localctx
}
