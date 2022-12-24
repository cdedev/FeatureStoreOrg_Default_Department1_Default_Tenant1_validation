// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671862557285/StatementOfPurpose.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StatementOfPurpose

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

type StatementOfPurposeParser struct {
	*antlr.BaseParser
}

var statementofpurposeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func statementofpurposeParserInit() {
	staticData := &statementofpurposeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STATEMENTOFPURPOSE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "statementofpurpose",
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

// StatementOfPurposeParserInit initializes any static state used to implement StatementOfPurposeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStatementOfPurposeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StatementOfPurposeParserInit() {
	staticData := &statementofpurposeParserStaticData
	staticData.once.Do(statementofpurposeParserInit)
}

// NewStatementOfPurposeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStatementOfPurposeParser(input antlr.TokenStream) *StatementOfPurposeParser {
	StatementOfPurposeParserInit()
	this := new(StatementOfPurposeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &statementofpurposeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "StatementOfPurpose.g4"

	return this
}

// StatementOfPurposeParser tokens.
const (
	StatementOfPurposeParserEOF                = antlr.TokenEOF
	StatementOfPurposeParserSTATEMENTOFPURPOSE = 1
	StatementOfPurposeParserOWNKEY             = 2
	StatementOfPurposeParserSPLITKEY           = 3
	StatementOfPurposeParserWS                 = 4
)

// StatementOfPurposeParser rules.
const (
	StatementOfPurposeParserRULE_expression         = 0
	StatementOfPurposeParserRULE_statementofpurpose = 1
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
	p.RuleIndex = StatementOfPurposeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StatementOfPurposeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Statementofpurpose() IStatementofpurposeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementofpurposeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementofpurposeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StatementOfPurposeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatementOfPurposeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatementOfPurposeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StatementOfPurposeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StatementOfPurposeParserRULE_expression)

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
		p.Statementofpurpose()
	}
	{
		p.SetState(5)
		p.Match(StatementOfPurposeParserEOF)
	}

	return localctx
}

// IStatementofpurposeContext is an interface to support dynamic dispatch.
type IStatementofpurposeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementofpurposeContext differentiates from other interfaces.
	IsStatementofpurposeContext()
}

type StatementofpurposeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementofpurposeContext() *StatementofpurposeContext {
	var p = new(StatementofpurposeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StatementOfPurposeParserRULE_statementofpurpose
	return p
}

func (*StatementofpurposeContext) IsStatementofpurposeContext() {}

func NewStatementofpurposeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementofpurposeContext {
	var p = new(StatementofpurposeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StatementOfPurposeParserRULE_statementofpurpose

	return p
}

func (s *StatementofpurposeContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementofpurposeContext) STATEMENTOFPURPOSE() antlr.TerminalNode {
	return s.GetToken(StatementOfPurposeParserSTATEMENTOFPURPOSE, 0)
}

func (s *StatementofpurposeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementofpurposeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementofpurposeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatementOfPurposeListener); ok {
		listenerT.EnterStatementofpurpose(s)
	}
}

func (s *StatementofpurposeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StatementOfPurposeListener); ok {
		listenerT.ExitStatementofpurpose(s)
	}
}

func (p *StatementOfPurposeParser) Statementofpurpose() (localctx IStatementofpurposeContext) {
	this := p
	_ = this

	localctx = NewStatementofpurposeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StatementOfPurposeParserRULE_statementofpurpose)

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
		p.Match(StatementOfPurposeParserSTATEMENTOFPURPOSE)
	}

	return localctx
}
