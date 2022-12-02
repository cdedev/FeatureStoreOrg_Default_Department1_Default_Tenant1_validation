// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669984432843/RecoveredCases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RecoveredCases

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

type RecoveredCasesParser struct {
	*antlr.BaseParser
}

var recoveredcasesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func recoveredcasesParserInit() {
	staticData := &recoveredcasesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RECOVEREDCASES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "recoveredcases",
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

// RecoveredCasesParserInit initializes any static state used to implement RecoveredCasesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRecoveredCasesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RecoveredCasesParserInit() {
	staticData := &recoveredcasesParserStaticData
	staticData.once.Do(recoveredcasesParserInit)
}

// NewRecoveredCasesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRecoveredCasesParser(input antlr.TokenStream) *RecoveredCasesParser {
	RecoveredCasesParserInit()
	this := new(RecoveredCasesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &recoveredcasesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RecoveredCases.g4"

	return this
}

// RecoveredCasesParser tokens.
const (
	RecoveredCasesParserEOF            = antlr.TokenEOF
	RecoveredCasesParserRECOVEREDCASES = 1
	RecoveredCasesParserOWNKEY         = 2
	RecoveredCasesParserSPLITKEY       = 3
	RecoveredCasesParserWS             = 4
)

// RecoveredCasesParser rules.
const (
	RecoveredCasesParserRULE_expression     = 0
	RecoveredCasesParserRULE_recoveredcases = 1
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
	p.RuleIndex = RecoveredCasesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RecoveredCasesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Recoveredcases() IRecoveredcasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecoveredcasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecoveredcasesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RecoveredCasesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RecoveredCasesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RecoveredCasesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RecoveredCasesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RecoveredCasesParserRULE_expression)

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
		p.Recoveredcases()
	}
	{
		p.SetState(5)
		p.Match(RecoveredCasesParserEOF)
	}

	return localctx
}

// IRecoveredcasesContext is an interface to support dynamic dispatch.
type IRecoveredcasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecoveredcasesContext differentiates from other interfaces.
	IsRecoveredcasesContext()
}

type RecoveredcasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecoveredcasesContext() *RecoveredcasesContext {
	var p = new(RecoveredcasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RecoveredCasesParserRULE_recoveredcases
	return p
}

func (*RecoveredcasesContext) IsRecoveredcasesContext() {}

func NewRecoveredcasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecoveredcasesContext {
	var p = new(RecoveredcasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RecoveredCasesParserRULE_recoveredcases

	return p
}

func (s *RecoveredcasesContext) GetParser() antlr.Parser { return s.parser }

func (s *RecoveredcasesContext) RECOVEREDCASES() antlr.TerminalNode {
	return s.GetToken(RecoveredCasesParserRECOVEREDCASES, 0)
}

func (s *RecoveredcasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecoveredcasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecoveredcasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RecoveredCasesListener); ok {
		listenerT.EnterRecoveredcases(s)
	}
}

func (s *RecoveredcasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RecoveredCasesListener); ok {
		listenerT.ExitRecoveredcases(s)
	}
}

func (p *RecoveredCasesParser) Recoveredcases() (localctx IRecoveredcasesContext) {
	this := p
	_ = this

	localctx = NewRecoveredcasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RecoveredCasesParserRULE_recoveredcases)

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
		p.Match(RecoveredCasesParserRECOVEREDCASES)
	}

	return localctx
}
