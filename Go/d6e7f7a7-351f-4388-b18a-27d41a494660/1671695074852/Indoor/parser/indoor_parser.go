// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671695074852/Indoor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Indoor

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

type IndoorParser struct {
	*antlr.BaseParser
}

var indoorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func indoorParserInit() {
	staticData := &indoorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INDOOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "indoor",
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

// IndoorParserInit initializes any static state used to implement IndoorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIndoorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IndoorParserInit() {
	staticData := &indoorParserStaticData
	staticData.once.Do(indoorParserInit)
}

// NewIndoorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIndoorParser(input antlr.TokenStream) *IndoorParser {
	IndoorParserInit()
	this := new(IndoorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &indoorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Indoor.g4"

	return this
}

// IndoorParser tokens.
const (
	IndoorParserEOF      = antlr.TokenEOF
	IndoorParserINDOOR   = 1
	IndoorParserOWNKEY   = 2
	IndoorParserSPLITKEY = 3
	IndoorParserWS       = 4
)

// IndoorParser rules.
const (
	IndoorParserRULE_expression = 0
	IndoorParserRULE_indoor     = 1
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
	p.RuleIndex = IndoorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IndoorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Indoor() IIndoorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndoorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndoorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IndoorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndoorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndoorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IndoorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IndoorParserRULE_expression)

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
		p.Indoor()
	}
	{
		p.SetState(5)
		p.Match(IndoorParserEOF)
	}

	return localctx
}

// IIndoorContext is an interface to support dynamic dispatch.
type IIndoorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIndoorContext differentiates from other interfaces.
	IsIndoorContext()
}

type IndoorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndoorContext() *IndoorContext {
	var p = new(IndoorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IndoorParserRULE_indoor
	return p
}

func (*IndoorContext) IsIndoorContext() {}

func NewIndoorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndoorContext {
	var p = new(IndoorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IndoorParserRULE_indoor

	return p
}

func (s *IndoorContext) GetParser() antlr.Parser { return s.parser }

func (s *IndoorContext) INDOOR() antlr.TerminalNode {
	return s.GetToken(IndoorParserINDOOR, 0)
}

func (s *IndoorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndoorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndoorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndoorListener); ok {
		listenerT.EnterIndoor(s)
	}
}

func (s *IndoorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IndoorListener); ok {
		listenerT.ExitIndoor(s)
	}
}

func (p *IndoorParser) Indoor() (localctx IIndoorContext) {
	this := p
	_ = this

	localctx = NewIndoorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IndoorParserRULE_indoor)

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
		p.Match(IndoorParserINDOOR)
	}

	return localctx
}
