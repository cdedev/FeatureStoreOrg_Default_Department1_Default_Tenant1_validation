// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672120114803/Elevation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Elevation

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

type ElevationParser struct {
	*antlr.BaseParser
}

var elevationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func elevationParserInit() {
	staticData := &elevationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ELEVATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "elevation",
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

// ElevationParserInit initializes any static state used to implement ElevationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewElevationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ElevationParserInit() {
	staticData := &elevationParserStaticData
	staticData.once.Do(elevationParserInit)
}

// NewElevationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewElevationParser(input antlr.TokenStream) *ElevationParser {
	ElevationParserInit()
	this := new(ElevationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &elevationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Elevation.g4"

	return this
}

// ElevationParser tokens.
const (
	ElevationParserEOF       = antlr.TokenEOF
	ElevationParserELEVATION = 1
	ElevationParserOWNKEY    = 2
	ElevationParserSPLITKEY  = 3
	ElevationParserWS        = 4
)

// ElevationParser rules.
const (
	ElevationParserRULE_expression = 0
	ElevationParserRULE_elevation  = 1
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
	p.RuleIndex = ElevationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElevationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Elevation() IElevationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElevationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElevationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ElevationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElevationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElevationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ElevationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ElevationParserRULE_expression)

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
		p.Elevation()
	}
	{
		p.SetState(5)
		p.Match(ElevationParserEOF)
	}

	return localctx
}

// IElevationContext is an interface to support dynamic dispatch.
type IElevationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElevationContext differentiates from other interfaces.
	IsElevationContext()
}

type ElevationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElevationContext() *ElevationContext {
	var p = new(ElevationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElevationParserRULE_elevation
	return p
}

func (*ElevationContext) IsElevationContext() {}

func NewElevationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElevationContext {
	var p = new(ElevationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElevationParserRULE_elevation

	return p
}

func (s *ElevationContext) GetParser() antlr.Parser { return s.parser }

func (s *ElevationContext) ELEVATION() antlr.TerminalNode {
	return s.GetToken(ElevationParserELEVATION, 0)
}

func (s *ElevationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElevationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElevationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElevationListener); ok {
		listenerT.EnterElevation(s)
	}
}

func (s *ElevationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElevationListener); ok {
		listenerT.ExitElevation(s)
	}
}

func (p *ElevationParser) Elevation() (localctx IElevationContext) {
	this := p
	_ = this

	localctx = NewElevationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ElevationParserRULE_elevation)

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
		p.Match(ElevationParserELEVATION)
	}

	return localctx
}
