// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672810323495/Circumference.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Circumference

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

type CircumferenceParser struct {
	*antlr.BaseParser
}

var circumferenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func circumferenceParserInit() {
	staticData := &circumferenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CIRCUMFERENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "circumference",
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

// CircumferenceParserInit initializes any static state used to implement CircumferenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCircumferenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CircumferenceParserInit() {
	staticData := &circumferenceParserStaticData
	staticData.once.Do(circumferenceParserInit)
}

// NewCircumferenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCircumferenceParser(input antlr.TokenStream) *CircumferenceParser {
	CircumferenceParserInit()
	this := new(CircumferenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &circumferenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Circumference.g4"

	return this
}

// CircumferenceParser tokens.
const (
	CircumferenceParserEOF           = antlr.TokenEOF
	CircumferenceParserCIRCUMFERENCE = 1
	CircumferenceParserOWNKEY        = 2
	CircumferenceParserSPLITKEY      = 3
	CircumferenceParserWS            = 4
)

// CircumferenceParser rules.
const (
	CircumferenceParserRULE_expression    = 0
	CircumferenceParserRULE_circumference = 1
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
	p.RuleIndex = CircumferenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CircumferenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Circumference() ICircumferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICircumferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICircumferenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CircumferenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CircumferenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CircumferenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CircumferenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CircumferenceParserRULE_expression)

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
		p.Circumference()
	}
	{
		p.SetState(5)
		p.Match(CircumferenceParserEOF)
	}

	return localctx
}

// ICircumferenceContext is an interface to support dynamic dispatch.
type ICircumferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCircumferenceContext differentiates from other interfaces.
	IsCircumferenceContext()
}

type CircumferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCircumferenceContext() *CircumferenceContext {
	var p = new(CircumferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CircumferenceParserRULE_circumference
	return p
}

func (*CircumferenceContext) IsCircumferenceContext() {}

func NewCircumferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CircumferenceContext {
	var p = new(CircumferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CircumferenceParserRULE_circumference

	return p
}

func (s *CircumferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *CircumferenceContext) CIRCUMFERENCE() antlr.TerminalNode {
	return s.GetToken(CircumferenceParserCIRCUMFERENCE, 0)
}

func (s *CircumferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircumferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CircumferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CircumferenceListener); ok {
		listenerT.EnterCircumference(s)
	}
}

func (s *CircumferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CircumferenceListener); ok {
		listenerT.ExitCircumference(s)
	}
}

func (p *CircumferenceParser) Circumference() (localctx ICircumferenceContext) {
	this := p
	_ = this

	localctx = NewCircumferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CircumferenceParserRULE_circumference)

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
		p.Match(CircumferenceParserCIRCUMFERENCE)
	}

	return localctx
}
