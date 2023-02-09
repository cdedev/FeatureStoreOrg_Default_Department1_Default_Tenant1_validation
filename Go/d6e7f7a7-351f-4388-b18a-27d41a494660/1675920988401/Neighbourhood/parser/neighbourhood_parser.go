// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675920988401/Neighbourhood.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Neighbourhood

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

type NeighbourhoodParser struct {
	*antlr.BaseParser
}

var neighbourhoodParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func neighbourhoodParserInit() {
	staticData := &neighbourhoodParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NEIGHBOURHOOD", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "neighbourhood",
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

// NeighbourhoodParserInit initializes any static state used to implement NeighbourhoodParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNeighbourhoodParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NeighbourhoodParserInit() {
	staticData := &neighbourhoodParserStaticData
	staticData.once.Do(neighbourhoodParserInit)
}

// NewNeighbourhoodParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNeighbourhoodParser(input antlr.TokenStream) *NeighbourhoodParser {
	NeighbourhoodParserInit()
	this := new(NeighbourhoodParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &neighbourhoodParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Neighbourhood.g4"

	return this
}

// NeighbourhoodParser tokens.
const (
	NeighbourhoodParserEOF           = antlr.TokenEOF
	NeighbourhoodParserNEIGHBOURHOOD = 1
	NeighbourhoodParserOWNKEY        = 2
	NeighbourhoodParserSPLITKEY      = 3
	NeighbourhoodParserWS            = 4
)

// NeighbourhoodParser rules.
const (
	NeighbourhoodParserRULE_expression    = 0
	NeighbourhoodParserRULE_neighbourhood = 1
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
	p.RuleIndex = NeighbourhoodParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeighbourhoodParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Neighbourhood() INeighbourhoodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INeighbourhoodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INeighbourhoodContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NeighbourhoodParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeighbourhoodListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeighbourhoodListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NeighbourhoodParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NeighbourhoodParserRULE_expression)

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
		p.Neighbourhood()
	}
	{
		p.SetState(5)
		p.Match(NeighbourhoodParserEOF)
	}

	return localctx
}

// INeighbourhoodContext is an interface to support dynamic dispatch.
type INeighbourhoodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeighbourhoodContext differentiates from other interfaces.
	IsNeighbourhoodContext()
}

type NeighbourhoodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeighbourhoodContext() *NeighbourhoodContext {
	var p = new(NeighbourhoodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NeighbourhoodParserRULE_neighbourhood
	return p
}

func (*NeighbourhoodContext) IsNeighbourhoodContext() {}

func NewNeighbourhoodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NeighbourhoodContext {
	var p = new(NeighbourhoodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NeighbourhoodParserRULE_neighbourhood

	return p
}

func (s *NeighbourhoodContext) GetParser() antlr.Parser { return s.parser }

func (s *NeighbourhoodContext) NEIGHBOURHOOD() antlr.TerminalNode {
	return s.GetToken(NeighbourhoodParserNEIGHBOURHOOD, 0)
}

func (s *NeighbourhoodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeighbourhoodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NeighbourhoodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeighbourhoodListener); ok {
		listenerT.EnterNeighbourhood(s)
	}
}

func (s *NeighbourhoodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NeighbourhoodListener); ok {
		listenerT.ExitNeighbourhood(s)
	}
}

func (p *NeighbourhoodParser) Neighbourhood() (localctx INeighbourhoodContext) {
	this := p
	_ = this

	localctx = NewNeighbourhoodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NeighbourhoodParserRULE_neighbourhood)

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
		p.Match(NeighbourhoodParserNEIGHBOURHOOD)
	}

	return localctx
}
