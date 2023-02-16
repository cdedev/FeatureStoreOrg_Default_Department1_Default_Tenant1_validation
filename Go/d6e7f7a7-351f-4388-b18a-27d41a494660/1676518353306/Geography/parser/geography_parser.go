// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676518353306/Geography.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Geography

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

type GeographyParser struct {
	*antlr.BaseParser
}

var geographyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func geographyParserInit() {
	staticData := &geographyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GEOGRAPHY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "geography",
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

// GeographyParserInit initializes any static state used to implement GeographyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGeographyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GeographyParserInit() {
	staticData := &geographyParserStaticData
	staticData.once.Do(geographyParserInit)
}

// NewGeographyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGeographyParser(input antlr.TokenStream) *GeographyParser {
	GeographyParserInit()
	this := new(GeographyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &geographyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Geography.g4"

	return this
}

// GeographyParser tokens.
const (
	GeographyParserEOF       = antlr.TokenEOF
	GeographyParserGEOGRAPHY = 1
	GeographyParserOWNKEY    = 2
	GeographyParserSPLITKEY  = 3
	GeographyParserWS        = 4
)

// GeographyParser rules.
const (
	GeographyParserRULE_expression = 0
	GeographyParserRULE_geography  = 1
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
	p.RuleIndex = GeographyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GeographyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Geography() IGeographyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeographyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeographyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GeographyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GeographyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GeographyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GeographyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GeographyParserRULE_expression)

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
		p.Geography()
	}
	{
		p.SetState(5)
		p.Match(GeographyParserEOF)
	}

	return localctx
}

// IGeographyContext is an interface to support dynamic dispatch.
type IGeographyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeographyContext differentiates from other interfaces.
	IsGeographyContext()
}

type GeographyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeographyContext() *GeographyContext {
	var p = new(GeographyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GeographyParserRULE_geography
	return p
}

func (*GeographyContext) IsGeographyContext() {}

func NewGeographyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeographyContext {
	var p = new(GeographyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GeographyParserRULE_geography

	return p
}

func (s *GeographyContext) GetParser() antlr.Parser { return s.parser }

func (s *GeographyContext) GEOGRAPHY() antlr.TerminalNode {
	return s.GetToken(GeographyParserGEOGRAPHY, 0)
}

func (s *GeographyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeographyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeographyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GeographyListener); ok {
		listenerT.EnterGeography(s)
	}
}

func (s *GeographyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GeographyListener); ok {
		listenerT.ExitGeography(s)
	}
}

func (p *GeographyParser) Geography() (localctx IGeographyContext) {
	this := p
	_ = this

	localctx = NewGeographyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GeographyParserRULE_geography)

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
		p.Match(GeographyParserGEOGRAPHY)
	}

	return localctx
}
