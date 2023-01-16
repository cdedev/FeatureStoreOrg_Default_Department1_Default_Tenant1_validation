// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867958192/Plastic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Plastic

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

type PlasticParser struct {
	*antlr.BaseParser
}

var plasticParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func plasticParserInit() {
	staticData := &plasticParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PLASTIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "plastic",
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

// PlasticParserInit initializes any static state used to implement PlasticParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPlasticParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PlasticParserInit() {
	staticData := &plasticParserStaticData
	staticData.once.Do(plasticParserInit)
}

// NewPlasticParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPlasticParser(input antlr.TokenStream) *PlasticParser {
	PlasticParserInit()
	this := new(PlasticParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &plasticParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Plastic.g4"

	return this
}

// PlasticParser tokens.
const (
	PlasticParserEOF      = antlr.TokenEOF
	PlasticParserPLASTIC  = 1
	PlasticParserOWNKEY   = 2
	PlasticParserSPLITKEY = 3
	PlasticParserWS       = 4
)

// PlasticParser rules.
const (
	PlasticParserRULE_expression = 0
	PlasticParserRULE_plastic    = 1
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
	p.RuleIndex = PlasticParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlasticParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Plastic() IPlasticContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPlasticContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPlasticContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PlasticParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlasticListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlasticListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PlasticParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PlasticParserRULE_expression)

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
		p.Plastic()
	}
	{
		p.SetState(5)
		p.Match(PlasticParserEOF)
	}

	return localctx
}

// IPlasticContext is an interface to support dynamic dispatch.
type IPlasticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlasticContext differentiates from other interfaces.
	IsPlasticContext()
}

type PlasticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlasticContext() *PlasticContext {
	var p = new(PlasticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PlasticParserRULE_plastic
	return p
}

func (*PlasticContext) IsPlasticContext() {}

func NewPlasticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlasticContext {
	var p = new(PlasticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PlasticParserRULE_plastic

	return p
}

func (s *PlasticContext) GetParser() antlr.Parser { return s.parser }

func (s *PlasticContext) PLASTIC() antlr.TerminalNode {
	return s.GetToken(PlasticParserPLASTIC, 0)
}

func (s *PlasticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlasticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PlasticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlasticListener); ok {
		listenerT.EnterPlastic(s)
	}
}

func (s *PlasticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PlasticListener); ok {
		listenerT.ExitPlastic(s)
	}
}

func (p *PlasticParser) Plastic() (localctx IPlasticContext) {
	this := p
	_ = this

	localctx = NewPlasticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PlasticParserRULE_plastic)

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
		p.Match(PlasticParserPLASTIC)
	}

	return localctx
}
