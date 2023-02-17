// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676608227546/Satisfaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Satisfaction

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

type SatisfactionParser struct {
	*antlr.BaseParser
}

var satisfactionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func satisfactionParserInit() {
	staticData := &satisfactionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SATISFACTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "satisfaction",
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

// SatisfactionParserInit initializes any static state used to implement SatisfactionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSatisfactionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SatisfactionParserInit() {
	staticData := &satisfactionParserStaticData
	staticData.once.Do(satisfactionParserInit)
}

// NewSatisfactionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSatisfactionParser(input antlr.TokenStream) *SatisfactionParser {
	SatisfactionParserInit()
	this := new(SatisfactionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &satisfactionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Satisfaction.g4"

	return this
}

// SatisfactionParser tokens.
const (
	SatisfactionParserEOF          = antlr.TokenEOF
	SatisfactionParserSATISFACTION = 1
	SatisfactionParserOWNKEY       = 2
	SatisfactionParserSPLITKEY     = 3
	SatisfactionParserWS           = 4
)

// SatisfactionParser rules.
const (
	SatisfactionParserRULE_expression   = 0
	SatisfactionParserRULE_satisfaction = 1
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
	p.RuleIndex = SatisfactionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SatisfactionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Satisfaction() ISatisfactionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISatisfactionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISatisfactionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SatisfactionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SatisfactionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SatisfactionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SatisfactionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SatisfactionParserRULE_expression)

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
		p.Satisfaction()
	}
	{
		p.SetState(5)
		p.Match(SatisfactionParserEOF)
	}

	return localctx
}

// ISatisfactionContext is an interface to support dynamic dispatch.
type ISatisfactionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSatisfactionContext differentiates from other interfaces.
	IsSatisfactionContext()
}

type SatisfactionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySatisfactionContext() *SatisfactionContext {
	var p = new(SatisfactionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SatisfactionParserRULE_satisfaction
	return p
}

func (*SatisfactionContext) IsSatisfactionContext() {}

func NewSatisfactionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SatisfactionContext {
	var p = new(SatisfactionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SatisfactionParserRULE_satisfaction

	return p
}

func (s *SatisfactionContext) GetParser() antlr.Parser { return s.parser }

func (s *SatisfactionContext) SATISFACTION() antlr.TerminalNode {
	return s.GetToken(SatisfactionParserSATISFACTION, 0)
}

func (s *SatisfactionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SatisfactionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SatisfactionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SatisfactionListener); ok {
		listenerT.EnterSatisfaction(s)
	}
}

func (s *SatisfactionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SatisfactionListener); ok {
		listenerT.ExitSatisfaction(s)
	}
}

func (p *SatisfactionParser) Satisfaction() (localctx ISatisfactionContext) {
	this := p
	_ = this

	localctx = NewSatisfactionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SatisfactionParserRULE_satisfaction)

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
		p.Match(SatisfactionParserSATISFACTION)
	}

	return localctx
}
