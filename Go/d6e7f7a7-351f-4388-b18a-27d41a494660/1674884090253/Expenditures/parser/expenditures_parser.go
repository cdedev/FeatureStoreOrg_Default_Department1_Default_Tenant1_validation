// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884090253/Expenditures.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expenditures

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

type ExpendituresParser struct {
	*antlr.BaseParser
}

var expendituresParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expendituresParserInit() {
	staticData := &expendituresParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EXPENDITURES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "expenditures",
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

// ExpendituresParserInit initializes any static state used to implement ExpendituresParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpendituresParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpendituresParserInit() {
	staticData := &expendituresParserStaticData
	staticData.once.Do(expendituresParserInit)
}

// NewExpendituresParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpendituresParser(input antlr.TokenStream) *ExpendituresParser {
	ExpendituresParserInit()
	this := new(ExpendituresParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &expendituresParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Expenditures.g4"

	return this
}

// ExpendituresParser tokens.
const (
	ExpendituresParserEOF          = antlr.TokenEOF
	ExpendituresParserEXPENDITURES = 1
	ExpendituresParserOWNKEY       = 2
	ExpendituresParserSPLITKEY     = 3
	ExpendituresParserWS           = 4
)

// ExpendituresParser rules.
const (
	ExpendituresParserRULE_expression   = 0
	ExpendituresParserRULE_expenditures = 1
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
	p.RuleIndex = ExpendituresParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpendituresParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Expenditures() IExpendituresContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpendituresContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpendituresContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpendituresParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpendituresListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpendituresListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExpendituresParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpendituresParserRULE_expression)

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
		p.Expenditures()
	}
	{
		p.SetState(5)
		p.Match(ExpendituresParserEOF)
	}

	return localctx
}

// IExpendituresContext is an interface to support dynamic dispatch.
type IExpendituresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpendituresContext differentiates from other interfaces.
	IsExpendituresContext()
}

type ExpendituresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpendituresContext() *ExpendituresContext {
	var p = new(ExpendituresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpendituresParserRULE_expenditures
	return p
}

func (*ExpendituresContext) IsExpendituresContext() {}

func NewExpendituresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpendituresContext {
	var p = new(ExpendituresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpendituresParserRULE_expenditures

	return p
}

func (s *ExpendituresContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpendituresContext) EXPENDITURES() antlr.TerminalNode {
	return s.GetToken(ExpendituresParserEXPENDITURES, 0)
}

func (s *ExpendituresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpendituresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpendituresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpendituresListener); ok {
		listenerT.EnterExpenditures(s)
	}
}

func (s *ExpendituresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpendituresListener); ok {
		listenerT.ExitExpenditures(s)
	}
}

func (p *ExpendituresParser) Expenditures() (localctx IExpendituresContext) {
	this := p
	_ = this

	localctx = NewExpendituresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpendituresParserRULE_expenditures)

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
		p.Match(ExpendituresParserEXPENDITURES)
	}

	return localctx
}
