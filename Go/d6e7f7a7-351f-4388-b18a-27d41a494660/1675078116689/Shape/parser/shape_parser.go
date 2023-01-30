// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675078116689/Shape.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shape

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

type ShapeParser struct {
	*antlr.BaseParser
}

var shapeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func shapeParserInit() {
	staticData := &shapeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SHAPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "shape",
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

// ShapeParserInit initializes any static state used to implement ShapeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewShapeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ShapeParserInit() {
	staticData := &shapeParserStaticData
	staticData.once.Do(shapeParserInit)
}

// NewShapeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewShapeParser(input antlr.TokenStream) *ShapeParser {
	ShapeParserInit()
	this := new(ShapeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &shapeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Shape.g4"

	return this
}

// ShapeParser tokens.
const (
	ShapeParserEOF      = antlr.TokenEOF
	ShapeParserSHAPE    = 1
	ShapeParserOWNKEY   = 2
	ShapeParserSPLITKEY = 3
	ShapeParserWS       = 4
)

// ShapeParser rules.
const (
	ShapeParserRULE_expression = 0
	ShapeParserRULE_shape      = 1
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
	p.RuleIndex = ShapeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShapeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Shape() IShapeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IShapeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IShapeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ShapeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShapeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShapeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ShapeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ShapeParserRULE_expression)

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
		p.Shape()
	}
	{
		p.SetState(5)
		p.Match(ShapeParserEOF)
	}

	return localctx
}

// IShapeContext is an interface to support dynamic dispatch.
type IShapeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsShapeContext differentiates from other interfaces.
	IsShapeContext()
}

type ShapeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyShapeContext() *ShapeContext {
	var p = new(ShapeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ShapeParserRULE_shape
	return p
}

func (*ShapeContext) IsShapeContext() {}

func NewShapeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ShapeContext {
	var p = new(ShapeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ShapeParserRULE_shape

	return p
}

func (s *ShapeContext) GetParser() antlr.Parser { return s.parser }

func (s *ShapeContext) SHAPE() antlr.TerminalNode {
	return s.GetToken(ShapeParserSHAPE, 0)
}

func (s *ShapeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShapeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ShapeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShapeListener); ok {
		listenerT.EnterShape(s)
	}
}

func (s *ShapeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ShapeListener); ok {
		listenerT.ExitShape(s)
	}
}

func (p *ShapeParser) Shape() (localctx IShapeContext) {
	this := p
	_ = this

	localctx = NewShapeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ShapeParserRULE_shape)

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
		p.Match(ShapeParserSHAPE)
	}

	return localctx
}
