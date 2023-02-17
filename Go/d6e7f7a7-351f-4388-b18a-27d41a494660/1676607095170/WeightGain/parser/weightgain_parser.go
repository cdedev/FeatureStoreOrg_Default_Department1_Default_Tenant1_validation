// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676607095170/WeightGain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeightGain

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

type WeightGainParser struct {
	*antlr.BaseParser
}

var weightgainParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func weightgainParserInit() {
	staticData := &weightgainParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WEIGHTGAIN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "weightgain",
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

// WeightGainParserInit initializes any static state used to implement WeightGainParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWeightGainParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WeightGainParserInit() {
	staticData := &weightgainParserStaticData
	staticData.once.Do(weightgainParserInit)
}

// NewWeightGainParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWeightGainParser(input antlr.TokenStream) *WeightGainParser {
	WeightGainParserInit()
	this := new(WeightGainParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &weightgainParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "WeightGain.g4"

	return this
}

// WeightGainParser tokens.
const (
	WeightGainParserEOF        = antlr.TokenEOF
	WeightGainParserWEIGHTGAIN = 1
	WeightGainParserOWNKEY     = 2
	WeightGainParserSPLITKEY   = 3
	WeightGainParserWS         = 4
)

// WeightGainParser rules.
const (
	WeightGainParserRULE_expression = 0
	WeightGainParserRULE_weightgain = 1
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
	p.RuleIndex = WeightGainParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeightGainParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Weightgain() IWeightgainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWeightgainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWeightgainContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WeightGainParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeightGainListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeightGainListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WeightGainParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WeightGainParserRULE_expression)

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
		p.Weightgain()
	}
	{
		p.SetState(5)
		p.Match(WeightGainParserEOF)
	}

	return localctx
}

// IWeightgainContext is an interface to support dynamic dispatch.
type IWeightgainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeightgainContext differentiates from other interfaces.
	IsWeightgainContext()
}

type WeightgainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeightgainContext() *WeightgainContext {
	var p = new(WeightgainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WeightGainParserRULE_weightgain
	return p
}

func (*WeightgainContext) IsWeightgainContext() {}

func NewWeightgainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeightgainContext {
	var p = new(WeightgainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WeightGainParserRULE_weightgain

	return p
}

func (s *WeightgainContext) GetParser() antlr.Parser { return s.parser }

func (s *WeightgainContext) WEIGHTGAIN() antlr.TerminalNode {
	return s.GetToken(WeightGainParserWEIGHTGAIN, 0)
}

func (s *WeightgainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeightgainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeightgainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeightGainListener); ok {
		listenerT.EnterWeightgain(s)
	}
}

func (s *WeightgainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WeightGainListener); ok {
		listenerT.ExitWeightgain(s)
	}
}

func (p *WeightGainParser) Weightgain() (localctx IWeightgainContext) {
	this := p
	_ = this

	localctx = NewWeightgainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WeightGainParserRULE_weightgain)

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
		p.Match(WeightGainParserWEIGHTGAIN)
	}

	return localctx
}
