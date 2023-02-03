// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675392826128/Concentration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Concentration

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

type ConcentrationParser struct {
	*antlr.BaseParser
}

var concentrationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func concentrationParserInit() {
	staticData := &concentrationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONCENTRATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "concentration",
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

// ConcentrationParserInit initializes any static state used to implement ConcentrationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConcentrationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConcentrationParserInit() {
	staticData := &concentrationParserStaticData
	staticData.once.Do(concentrationParserInit)
}

// NewConcentrationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConcentrationParser(input antlr.TokenStream) *ConcentrationParser {
	ConcentrationParserInit()
	this := new(ConcentrationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &concentrationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Concentration.g4"

	return this
}

// ConcentrationParser tokens.
const (
	ConcentrationParserEOF           = antlr.TokenEOF
	ConcentrationParserCONCENTRATION = 1
	ConcentrationParserOWNKEY        = 2
	ConcentrationParserSPLITKEY      = 3
	ConcentrationParserWS            = 4
)

// ConcentrationParser rules.
const (
	ConcentrationParserRULE_expression    = 0
	ConcentrationParserRULE_concentration = 1
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
	p.RuleIndex = ConcentrationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcentrationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Concentration() IConcentrationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcentrationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcentrationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConcentrationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcentrationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcentrationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConcentrationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConcentrationParserRULE_expression)

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
		p.Concentration()
	}
	{
		p.SetState(5)
		p.Match(ConcentrationParserEOF)
	}

	return localctx
}

// IConcentrationContext is an interface to support dynamic dispatch.
type IConcentrationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcentrationContext differentiates from other interfaces.
	IsConcentrationContext()
}

type ConcentrationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcentrationContext() *ConcentrationContext {
	var p = new(ConcentrationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcentrationParserRULE_concentration
	return p
}

func (*ConcentrationContext) IsConcentrationContext() {}

func NewConcentrationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcentrationContext {
	var p = new(ConcentrationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcentrationParserRULE_concentration

	return p
}

func (s *ConcentrationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcentrationContext) CONCENTRATION() antlr.TerminalNode {
	return s.GetToken(ConcentrationParserCONCENTRATION, 0)
}

func (s *ConcentrationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcentrationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcentrationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcentrationListener); ok {
		listenerT.EnterConcentration(s)
	}
}

func (s *ConcentrationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcentrationListener); ok {
		listenerT.ExitConcentration(s)
	}
}

func (p *ConcentrationParser) Concentration() (localctx IConcentrationContext) {
	this := p
	_ = this

	localctx = NewConcentrationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConcentrationParserRULE_concentration)

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
		p.Match(ConcentrationParserCONCENTRATION)
	}

	return localctx
}
