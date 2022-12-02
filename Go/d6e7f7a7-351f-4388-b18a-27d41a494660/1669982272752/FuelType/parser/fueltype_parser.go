// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982272752/FuelType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FuelType

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

type FuelTypeParser struct {
	*antlr.BaseParser
}

var fueltypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fueltypeParserInit() {
	staticData := &fueltypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FUELTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "fueltype",
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

// FuelTypeParserInit initializes any static state used to implement FuelTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFuelTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FuelTypeParserInit() {
	staticData := &fueltypeParserStaticData
	staticData.once.Do(fueltypeParserInit)
}

// NewFuelTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFuelTypeParser(input antlr.TokenStream) *FuelTypeParser {
	FuelTypeParserInit()
	this := new(FuelTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fueltypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FuelType.g4"

	return this
}

// FuelTypeParser tokens.
const (
	FuelTypeParserEOF      = antlr.TokenEOF
	FuelTypeParserFUELTYPE = 1
	FuelTypeParserOWNKEY   = 2
	FuelTypeParserSPLITKEY = 3
	FuelTypeParserWS       = 4
)

// FuelTypeParser rules.
const (
	FuelTypeParserRULE_expression = 0
	FuelTypeParserRULE_fueltype   = 1
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
	p.RuleIndex = FuelTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuelTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Fueltype() IFueltypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFueltypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFueltypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FuelTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FuelTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FuelTypeParserRULE_expression)

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
		p.Fueltype()
	}
	{
		p.SetState(5)
		p.Match(FuelTypeParserEOF)
	}

	return localctx
}

// IFueltypeContext is an interface to support dynamic dispatch.
type IFueltypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFueltypeContext differentiates from other interfaces.
	IsFueltypeContext()
}

type FueltypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFueltypeContext() *FueltypeContext {
	var p = new(FueltypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FuelTypeParserRULE_fueltype
	return p
}

func (*FueltypeContext) IsFueltypeContext() {}

func NewFueltypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FueltypeContext {
	var p = new(FueltypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FuelTypeParserRULE_fueltype

	return p
}

func (s *FueltypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FueltypeContext) FUELTYPE() antlr.TerminalNode {
	return s.GetToken(FuelTypeParserFUELTYPE, 0)
}

func (s *FueltypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FueltypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FueltypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelTypeListener); ok {
		listenerT.EnterFueltype(s)
	}
}

func (s *FueltypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FuelTypeListener); ok {
		listenerT.ExitFueltype(s)
	}
}

func (p *FuelTypeParser) Fueltype() (localctx IFueltypeContext) {
	this := p
	_ = this

	localctx = NewFueltypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FuelTypeParserRULE_fueltype)

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
		p.Match(FuelTypeParserFUELTYPE)
	}

	return localctx
}
