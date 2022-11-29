// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669691961999/Thickness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Thickness

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

type ThicknessParser struct {
	*antlr.BaseParser
}

var thicknessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func thicknessParserInit() {
	staticData := &thicknessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "THICKNESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "thickness",
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

// ThicknessParserInit initializes any static state used to implement ThicknessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewThicknessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ThicknessParserInit() {
	staticData := &thicknessParserStaticData
	staticData.once.Do(thicknessParserInit)
}

// NewThicknessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewThicknessParser(input antlr.TokenStream) *ThicknessParser {
	ThicknessParserInit()
	this := new(ThicknessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &thicknessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Thickness.g4"

	return this
}

// ThicknessParser tokens.
const (
	ThicknessParserEOF       = antlr.TokenEOF
	ThicknessParserTHICKNESS = 1
	ThicknessParserOWNKEY    = 2
	ThicknessParserSPLITKEY  = 3
	ThicknessParserWS        = 4
)

// ThicknessParser rules.
const (
	ThicknessParserRULE_expression = 0
	ThicknessParserRULE_thickness  = 1
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
	p.RuleIndex = ThicknessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThicknessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Thickness() IThicknessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThicknessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThicknessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ThicknessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThicknessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThicknessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ThicknessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ThicknessParserRULE_expression)

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
		p.Thickness()
	}
	{
		p.SetState(5)
		p.Match(ThicknessParserEOF)
	}

	return localctx
}

// IThicknessContext is an interface to support dynamic dispatch.
type IThicknessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThicknessContext differentiates from other interfaces.
	IsThicknessContext()
}

type ThicknessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThicknessContext() *ThicknessContext {
	var p = new(ThicknessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThicknessParserRULE_thickness
	return p
}

func (*ThicknessContext) IsThicknessContext() {}

func NewThicknessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThicknessContext {
	var p = new(ThicknessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThicknessParserRULE_thickness

	return p
}

func (s *ThicknessContext) GetParser() antlr.Parser { return s.parser }

func (s *ThicknessContext) THICKNESS() antlr.TerminalNode {
	return s.GetToken(ThicknessParserTHICKNESS, 0)
}

func (s *ThicknessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThicknessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThicknessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThicknessListener); ok {
		listenerT.EnterThickness(s)
	}
}

func (s *ThicknessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThicknessListener); ok {
		listenerT.ExitThickness(s)
	}
}

func (p *ThicknessParser) Thickness() (localctx IThicknessContext) {
	this := p
	_ = this

	localctx = NewThicknessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ThicknessParserRULE_thickness)

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
		p.Match(ThicknessParserTHICKNESS)
	}

	return localctx
}
