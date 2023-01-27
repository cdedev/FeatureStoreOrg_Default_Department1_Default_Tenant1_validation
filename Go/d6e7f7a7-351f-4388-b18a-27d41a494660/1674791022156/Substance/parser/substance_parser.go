// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674791022156/Substance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Substance

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

type SubstanceParser struct {
	*antlr.BaseParser
}

var substanceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func substanceParserInit() {
	staticData := &substanceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SUBSTANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "substance",
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

// SubstanceParserInit initializes any static state used to implement SubstanceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSubstanceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SubstanceParserInit() {
	staticData := &substanceParserStaticData
	staticData.once.Do(substanceParserInit)
}

// NewSubstanceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSubstanceParser(input antlr.TokenStream) *SubstanceParser {
	SubstanceParserInit()
	this := new(SubstanceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &substanceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Substance.g4"

	return this
}

// SubstanceParser tokens.
const (
	SubstanceParserEOF       = antlr.TokenEOF
	SubstanceParserSUBSTANCE = 1
	SubstanceParserOWNKEY    = 2
	SubstanceParserSPLITKEY  = 3
	SubstanceParserWS        = 4
)

// SubstanceParser rules.
const (
	SubstanceParserRULE_expression = 0
	SubstanceParserRULE_substance  = 1
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
	p.RuleIndex = SubstanceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstanceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Substance() ISubstanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubstanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubstanceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SubstanceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstanceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstanceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SubstanceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SubstanceParserRULE_expression)

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
		p.Substance()
	}
	{
		p.SetState(5)
		p.Match(SubstanceParserEOF)
	}

	return localctx
}

// ISubstanceContext is an interface to support dynamic dispatch.
type ISubstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubstanceContext differentiates from other interfaces.
	IsSubstanceContext()
}

type SubstanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubstanceContext() *SubstanceContext {
	var p = new(SubstanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SubstanceParserRULE_substance
	return p
}

func (*SubstanceContext) IsSubstanceContext() {}

func NewSubstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubstanceContext {
	var p = new(SubstanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubstanceParserRULE_substance

	return p
}

func (s *SubstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *SubstanceContext) SUBSTANCE() antlr.TerminalNode {
	return s.GetToken(SubstanceParserSUBSTANCE, 0)
}

func (s *SubstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstanceListener); ok {
		listenerT.EnterSubstance(s)
	}
}

func (s *SubstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubstanceListener); ok {
		listenerT.ExitSubstance(s)
	}
}

func (p *SubstanceParser) Substance() (localctx ISubstanceContext) {
	this := p
	_ = this

	localctx = NewSubstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SubstanceParserRULE_substance)

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
		p.Match(SubstanceParserSUBSTANCE)
	}

	return localctx
}
