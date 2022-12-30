// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672374472885/CoOrdinates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CoOrdinates

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

type CoOrdinatesParser struct {
	*antlr.BaseParser
}

var coordinatesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func coordinatesParserInit() {
	staticData := &coordinatesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COORDINATES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "coordinates",
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

// CoOrdinatesParserInit initializes any static state used to implement CoOrdinatesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCoOrdinatesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CoOrdinatesParserInit() {
	staticData := &coordinatesParserStaticData
	staticData.once.Do(coordinatesParserInit)
}

// NewCoOrdinatesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCoOrdinatesParser(input antlr.TokenStream) *CoOrdinatesParser {
	CoOrdinatesParserInit()
	this := new(CoOrdinatesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &coordinatesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CoOrdinates.g4"

	return this
}

// CoOrdinatesParser tokens.
const (
	CoOrdinatesParserEOF         = antlr.TokenEOF
	CoOrdinatesParserCOORDINATES = 1
	CoOrdinatesParserOWNKEY      = 2
	CoOrdinatesParserSPLITKEY    = 3
	CoOrdinatesParserWS          = 4
)

// CoOrdinatesParser rules.
const (
	CoOrdinatesParserRULE_expression  = 0
	CoOrdinatesParserRULE_coordinates = 1
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
	p.RuleIndex = CoOrdinatesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CoOrdinatesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Coordinates() ICoordinatesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICoordinatesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICoordinatesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CoOrdinatesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoOrdinatesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoOrdinatesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CoOrdinatesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CoOrdinatesParserRULE_expression)

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
		p.Coordinates()
	}
	{
		p.SetState(5)
		p.Match(CoOrdinatesParserEOF)
	}

	return localctx
}

// ICoordinatesContext is an interface to support dynamic dispatch.
type ICoordinatesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordinatesContext differentiates from other interfaces.
	IsCoordinatesContext()
}

type CoordinatesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordinatesContext() *CoordinatesContext {
	var p = new(CoordinatesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CoOrdinatesParserRULE_coordinates
	return p
}

func (*CoordinatesContext) IsCoordinatesContext() {}

func NewCoordinatesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinatesContext {
	var p = new(CoordinatesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CoOrdinatesParserRULE_coordinates

	return p
}

func (s *CoordinatesContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinatesContext) COORDINATES() antlr.TerminalNode {
	return s.GetToken(CoOrdinatesParserCOORDINATES, 0)
}

func (s *CoordinatesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinatesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CoordinatesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoOrdinatesListener); ok {
		listenerT.EnterCoordinates(s)
	}
}

func (s *CoordinatesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoOrdinatesListener); ok {
		listenerT.ExitCoordinates(s)
	}
}

func (p *CoOrdinatesParser) Coordinates() (localctx ICoordinatesContext) {
	this := p
	_ = this

	localctx = NewCoordinatesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CoOrdinatesParserRULE_coordinates)

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
		p.Match(CoOrdinatesParserCOORDINATES)
	}

	return localctx
}
