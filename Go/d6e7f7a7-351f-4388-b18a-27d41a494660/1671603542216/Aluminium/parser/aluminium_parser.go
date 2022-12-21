// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603542216/Aluminium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Aluminium

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

type AluminiumParser struct {
	*antlr.BaseParser
}

var aluminiumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aluminiumParserInit() {
	staticData := &aluminiumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ALUMINIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "aluminium",
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

// AluminiumParserInit initializes any static state used to implement AluminiumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAluminiumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AluminiumParserInit() {
	staticData := &aluminiumParserStaticData
	staticData.once.Do(aluminiumParserInit)
}

// NewAluminiumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAluminiumParser(input antlr.TokenStream) *AluminiumParser {
	AluminiumParserInit()
	this := new(AluminiumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &aluminiumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Aluminium.g4"

	return this
}

// AluminiumParser tokens.
const (
	AluminiumParserEOF       = antlr.TokenEOF
	AluminiumParserALUMINIUM = 1
	AluminiumParserOWNKEY    = 2
	AluminiumParserSPLITKEY  = 3
	AluminiumParserWS        = 4
)

// AluminiumParser rules.
const (
	AluminiumParserRULE_expression = 0
	AluminiumParserRULE_aluminium  = 1
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
	p.RuleIndex = AluminiumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AluminiumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Aluminium() IAluminiumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAluminiumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAluminiumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AluminiumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AluminiumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AluminiumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AluminiumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AluminiumParserRULE_expression)

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
		p.Aluminium()
	}
	{
		p.SetState(5)
		p.Match(AluminiumParserEOF)
	}

	return localctx
}

// IAluminiumContext is an interface to support dynamic dispatch.
type IAluminiumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAluminiumContext differentiates from other interfaces.
	IsAluminiumContext()
}

type AluminiumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAluminiumContext() *AluminiumContext {
	var p = new(AluminiumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AluminiumParserRULE_aluminium
	return p
}

func (*AluminiumContext) IsAluminiumContext() {}

func NewAluminiumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AluminiumContext {
	var p = new(AluminiumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AluminiumParserRULE_aluminium

	return p
}

func (s *AluminiumContext) GetParser() antlr.Parser { return s.parser }

func (s *AluminiumContext) ALUMINIUM() antlr.TerminalNode {
	return s.GetToken(AluminiumParserALUMINIUM, 0)
}

func (s *AluminiumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AluminiumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AluminiumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AluminiumListener); ok {
		listenerT.EnterAluminium(s)
	}
}

func (s *AluminiumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AluminiumListener); ok {
		listenerT.ExitAluminium(s)
	}
}

func (p *AluminiumParser) Aluminium() (localctx IAluminiumContext) {
	this := p
	_ = this

	localctx = NewAluminiumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AluminiumParserRULE_aluminium)

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
		p.Match(AluminiumParserALUMINIUM)
	}

	return localctx
}
