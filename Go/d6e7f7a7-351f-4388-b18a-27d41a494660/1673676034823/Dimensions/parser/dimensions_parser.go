// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676034823/Dimensions.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dimensions

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

type DimensionsParser struct {
	*antlr.BaseParser
}

var dimensionsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dimensionsParserInit() {
	staticData := &dimensionsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DIMENSIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dimensions",
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

// DimensionsParserInit initializes any static state used to implement DimensionsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDimensionsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DimensionsParserInit() {
	staticData := &dimensionsParserStaticData
	staticData.once.Do(dimensionsParserInit)
}

// NewDimensionsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDimensionsParser(input antlr.TokenStream) *DimensionsParser {
	DimensionsParserInit()
	this := new(DimensionsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dimensionsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Dimensions.g4"

	return this
}

// DimensionsParser tokens.
const (
	DimensionsParserEOF        = antlr.TokenEOF
	DimensionsParserDIMENSIONS = 1
	DimensionsParserOWNKEY     = 2
	DimensionsParserSPLITKEY   = 3
	DimensionsParserWS         = 4
)

// DimensionsParser rules.
const (
	DimensionsParserRULE_expression = 0
	DimensionsParserRULE_dimensions = 1
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
	p.RuleIndex = DimensionsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DimensionsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dimensions() IDimensionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimensionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimensionsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DimensionsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DimensionsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DimensionsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DimensionsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DimensionsParserRULE_expression)

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
		p.Dimensions()
	}
	{
		p.SetState(5)
		p.Match(DimensionsParserEOF)
	}

	return localctx
}

// IDimensionsContext is an interface to support dynamic dispatch.
type IDimensionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDimensionsContext differentiates from other interfaces.
	IsDimensionsContext()
}

type DimensionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimensionsContext() *DimensionsContext {
	var p = new(DimensionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DimensionsParserRULE_dimensions
	return p
}

func (*DimensionsContext) IsDimensionsContext() {}

func NewDimensionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensionsContext {
	var p = new(DimensionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DimensionsParserRULE_dimensions

	return p
}

func (s *DimensionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensionsContext) DIMENSIONS() antlr.TerminalNode {
	return s.GetToken(DimensionsParserDIMENSIONS, 0)
}

func (s *DimensionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DimensionsListener); ok {
		listenerT.EnterDimensions(s)
	}
}

func (s *DimensionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DimensionsListener); ok {
		listenerT.ExitDimensions(s)
	}
}

func (p *DimensionsParser) Dimensions() (localctx IDimensionsContext) {
	this := p
	_ = this

	localctx = NewDimensionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DimensionsParserRULE_dimensions)

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
		p.Match(DimensionsParserDIMENSIONS)
	}

	return localctx
}
