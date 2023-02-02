// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675307081052/Systolic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Systolic

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

type SystolicParser struct {
	*antlr.BaseParser
}

var systolicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func systolicParserInit() {
	staticData := &systolicParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SYSTOLIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "systolic",
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

// SystolicParserInit initializes any static state used to implement SystolicParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSystolicParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SystolicParserInit() {
	staticData := &systolicParserStaticData
	staticData.once.Do(systolicParserInit)
}

// NewSystolicParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSystolicParser(input antlr.TokenStream) *SystolicParser {
	SystolicParserInit()
	this := new(SystolicParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &systolicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Systolic.g4"

	return this
}

// SystolicParser tokens.
const (
	SystolicParserEOF      = antlr.TokenEOF
	SystolicParserSYSTOLIC = 1
	SystolicParserOWNKEY   = 2
	SystolicParserSPLITKEY = 3
	SystolicParserWS       = 4
)

// SystolicParser rules.
const (
	SystolicParserRULE_expression = 0
	SystolicParserRULE_systolic   = 1
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
	p.RuleIndex = SystolicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SystolicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Systolic() ISystolicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISystolicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISystolicContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SystolicParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SystolicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SystolicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SystolicParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SystolicParserRULE_expression)

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
		p.Systolic()
	}
	{
		p.SetState(5)
		p.Match(SystolicParserEOF)
	}

	return localctx
}

// ISystolicContext is an interface to support dynamic dispatch.
type ISystolicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSystolicContext differentiates from other interfaces.
	IsSystolicContext()
}

type SystolicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystolicContext() *SystolicContext {
	var p = new(SystolicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SystolicParserRULE_systolic
	return p
}

func (*SystolicContext) IsSystolicContext() {}

func NewSystolicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystolicContext {
	var p = new(SystolicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SystolicParserRULE_systolic

	return p
}

func (s *SystolicContext) GetParser() antlr.Parser { return s.parser }

func (s *SystolicContext) SYSTOLIC() antlr.TerminalNode {
	return s.GetToken(SystolicParserSYSTOLIC, 0)
}

func (s *SystolicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystolicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystolicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SystolicListener); ok {
		listenerT.EnterSystolic(s)
	}
}

func (s *SystolicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SystolicListener); ok {
		listenerT.ExitSystolic(s)
	}
}

func (p *SystolicParser) Systolic() (localctx ISystolicContext) {
	this := p
	_ = this

	localctx = NewSystolicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SystolicParserRULE_systolic)

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
		p.Match(SystolicParserSYSTOLIC)
	}

	return localctx
}
