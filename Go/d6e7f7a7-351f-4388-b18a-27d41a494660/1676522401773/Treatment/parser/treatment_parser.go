// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522401773/Treatment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Treatment

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

type TreatmentParser struct {
	*antlr.BaseParser
}

var treatmentParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func treatmentParserInit() {
	staticData := &treatmentParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TREATMENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "treatment",
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

// TreatmentParserInit initializes any static state used to implement TreatmentParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTreatmentParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TreatmentParserInit() {
	staticData := &treatmentParserStaticData
	staticData.once.Do(treatmentParserInit)
}

// NewTreatmentParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTreatmentParser(input antlr.TokenStream) *TreatmentParser {
	TreatmentParserInit()
	this := new(TreatmentParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &treatmentParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Treatment.g4"

	return this
}

// TreatmentParser tokens.
const (
	TreatmentParserEOF       = antlr.TokenEOF
	TreatmentParserTREATMENT = 1
	TreatmentParserOWNKEY    = 2
	TreatmentParserSPLITKEY  = 3
	TreatmentParserWS        = 4
)

// TreatmentParser rules.
const (
	TreatmentParserRULE_expression = 0
	TreatmentParserRULE_treatment  = 1
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
	p.RuleIndex = TreatmentParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreatmentParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Treatment() ITreatmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITreatmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITreatmentContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TreatmentParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreatmentListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreatmentListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TreatmentParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TreatmentParserRULE_expression)

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
		p.Treatment()
	}
	{
		p.SetState(5)
		p.Match(TreatmentParserEOF)
	}

	return localctx
}

// ITreatmentContext is an interface to support dynamic dispatch.
type ITreatmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTreatmentContext differentiates from other interfaces.
	IsTreatmentContext()
}

type TreatmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTreatmentContext() *TreatmentContext {
	var p = new(TreatmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TreatmentParserRULE_treatment
	return p
}

func (*TreatmentContext) IsTreatmentContext() {}

func NewTreatmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TreatmentContext {
	var p = new(TreatmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TreatmentParserRULE_treatment

	return p
}

func (s *TreatmentContext) GetParser() antlr.Parser { return s.parser }

func (s *TreatmentContext) TREATMENT() antlr.TerminalNode {
	return s.GetToken(TreatmentParserTREATMENT, 0)
}

func (s *TreatmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TreatmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TreatmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreatmentListener); ok {
		listenerT.EnterTreatment(s)
	}
}

func (s *TreatmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TreatmentListener); ok {
		listenerT.ExitTreatment(s)
	}
}

func (p *TreatmentParser) Treatment() (localctx ITreatmentContext) {
	this := p
	_ = this

	localctx = NewTreatmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TreatmentParserRULE_treatment)

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
		p.Match(TreatmentParserTREATMENT)
	}

	return localctx
}
