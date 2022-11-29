// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690829280/Elongation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Elongation

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

type ElongationParser struct {
	*antlr.BaseParser
}

var elongationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func elongationParserInit() {
	staticData := &elongationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ELONGATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "elongation",
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

// ElongationParserInit initializes any static state used to implement ElongationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewElongationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ElongationParserInit() {
	staticData := &elongationParserStaticData
	staticData.once.Do(elongationParserInit)
}

// NewElongationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewElongationParser(input antlr.TokenStream) *ElongationParser {
	ElongationParserInit()
	this := new(ElongationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &elongationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Elongation.g4"

	return this
}

// ElongationParser tokens.
const (
	ElongationParserEOF        = antlr.TokenEOF
	ElongationParserELONGATION = 1
	ElongationParserOWNKEY     = 2
	ElongationParserSPLITKEY   = 3
	ElongationParserWS         = 4
)

// ElongationParser rules.
const (
	ElongationParserRULE_expression = 0
	ElongationParserRULE_elongation = 1
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
	p.RuleIndex = ElongationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElongationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Elongation() IElongationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElongationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElongationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ElongationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElongationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElongationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ElongationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ElongationParserRULE_expression)

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
		p.Elongation()
	}
	{
		p.SetState(5)
		p.Match(ElongationParserEOF)
	}

	return localctx
}

// IElongationContext is an interface to support dynamic dispatch.
type IElongationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElongationContext differentiates from other interfaces.
	IsElongationContext()
}

type ElongationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElongationContext() *ElongationContext {
	var p = new(ElongationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElongationParserRULE_elongation
	return p
}

func (*ElongationContext) IsElongationContext() {}

func NewElongationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElongationContext {
	var p = new(ElongationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElongationParserRULE_elongation

	return p
}

func (s *ElongationContext) GetParser() antlr.Parser { return s.parser }

func (s *ElongationContext) ELONGATION() antlr.TerminalNode {
	return s.GetToken(ElongationParserELONGATION, 0)
}

func (s *ElongationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElongationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElongationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElongationListener); ok {
		listenerT.EnterElongation(s)
	}
}

func (s *ElongationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElongationListener); ok {
		listenerT.ExitElongation(s)
	}
}

func (p *ElongationParser) Elongation() (localctx IElongationContext) {
	this := p
	_ = this

	localctx = NewElongationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ElongationParserRULE_elongation)

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
		p.Match(ElongationParserELONGATION)
	}

	return localctx
}
