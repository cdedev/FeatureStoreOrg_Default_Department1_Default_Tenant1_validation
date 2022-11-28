// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623152379/Distance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Distance

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

type DistanceParser struct {
	*antlr.BaseParser
}

var distanceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func distanceParserInit() {
	staticData := &distanceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DISTANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "distance",
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

// DistanceParserInit initializes any static state used to implement DistanceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDistanceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DistanceParserInit() {
	staticData := &distanceParserStaticData
	staticData.once.Do(distanceParserInit)
}

// NewDistanceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDistanceParser(input antlr.TokenStream) *DistanceParser {
	DistanceParserInit()
	this := new(DistanceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &distanceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Distance.g4"

	return this
}

// DistanceParser tokens.
const (
	DistanceParserEOF      = antlr.TokenEOF
	DistanceParserDISTANCE = 1
	DistanceParserOWNKEY   = 2
	DistanceParserSPLITKEY = 3
	DistanceParserWS       = 4
)

// DistanceParser rules.
const (
	DistanceParserRULE_expression = 0
	DistanceParserRULE_distance   = 1
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
	p.RuleIndex = DistanceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DistanceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Distance() IDistanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistanceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DistanceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DistanceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DistanceParserRULE_expression)

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
		p.Distance()
	}
	{
		p.SetState(5)
		p.Match(DistanceParserEOF)
	}

	return localctx
}

// IDistanceContext is an interface to support dynamic dispatch.
type IDistanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDistanceContext differentiates from other interfaces.
	IsDistanceContext()
}

type DistanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistanceContext() *DistanceContext {
	var p = new(DistanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DistanceParserRULE_distance
	return p
}

func (*DistanceContext) IsDistanceContext() {}

func NewDistanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistanceContext {
	var p = new(DistanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DistanceParserRULE_distance

	return p
}

func (s *DistanceContext) GetParser() antlr.Parser { return s.parser }

func (s *DistanceContext) DISTANCE() antlr.TerminalNode {
	return s.GetToken(DistanceParserDISTANCE, 0)
}

func (s *DistanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.EnterDistance(s)
	}
}

func (s *DistanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DistanceListener); ok {
		listenerT.ExitDistance(s)
	}
}

func (p *DistanceParser) Distance() (localctx IDistanceContext) {
	this := p
	_ = this

	localctx = NewDistanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DistanceParserRULE_distance)

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
		p.Match(DistanceParserDISTANCE)
	}

	return localctx
}
