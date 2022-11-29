// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669692269411/Diabetes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diabetes

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

type DiabetesParser struct {
	*antlr.BaseParser
}

var diabetesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diabetesParserInit() {
	staticData := &diabetesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DIABETES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "diabetes",
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

// DiabetesParserInit initializes any static state used to implement DiabetesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiabetesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiabetesParserInit() {
	staticData := &diabetesParserStaticData
	staticData.once.Do(diabetesParserInit)
}

// NewDiabetesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiabetesParser(input antlr.TokenStream) *DiabetesParser {
	DiabetesParserInit()
	this := new(DiabetesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diabetesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Diabetes.g4"

	return this
}

// DiabetesParser tokens.
const (
	DiabetesParserEOF      = antlr.TokenEOF
	DiabetesParserDIABETES = 1
	DiabetesParserOWNKEY   = 2
	DiabetesParserSPLITKEY = 3
	DiabetesParserWS       = 4
)

// DiabetesParser rules.
const (
	DiabetesParserRULE_expression = 0
	DiabetesParserRULE_diabetes   = 1
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
	p.RuleIndex = DiabetesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiabetesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Diabetes() IDiabetesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiabetesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiabetesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiabetesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiabetesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiabetesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiabetesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiabetesParserRULE_expression)

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
		p.Diabetes()
	}
	{
		p.SetState(5)
		p.Match(DiabetesParserEOF)
	}

	return localctx
}

// IDiabetesContext is an interface to support dynamic dispatch.
type IDiabetesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiabetesContext differentiates from other interfaces.
	IsDiabetesContext()
}

type DiabetesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiabetesContext() *DiabetesContext {
	var p = new(DiabetesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiabetesParserRULE_diabetes
	return p
}

func (*DiabetesContext) IsDiabetesContext() {}

func NewDiabetesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiabetesContext {
	var p = new(DiabetesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiabetesParserRULE_diabetes

	return p
}

func (s *DiabetesContext) GetParser() antlr.Parser { return s.parser }

func (s *DiabetesContext) DIABETES() antlr.TerminalNode {
	return s.GetToken(DiabetesParserDIABETES, 0)
}

func (s *DiabetesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiabetesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiabetesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiabetesListener); ok {
		listenerT.EnterDiabetes(s)
	}
}

func (s *DiabetesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiabetesListener); ok {
		listenerT.ExitDiabetes(s)
	}
}

func (p *DiabetesParser) Diabetes() (localctx IDiabetesContext) {
	this := p
	_ = this

	localctx = NewDiabetesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiabetesParserRULE_diabetes)

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
		p.Match(DiabetesParserDIABETES)
	}

	return localctx
}
