// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202956618/NewBalanceOrig.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NewBalanceOrig

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

type NewBalanceOrigParser struct {
	*antlr.BaseParser
}

var newbalanceorigParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func newbalanceorigParserInit() {
	staticData := &newbalanceorigParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NEWBALANCEORIG", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "newbalanceorig",
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

// NewBalanceOrigParserInit initializes any static state used to implement NewBalanceOrigParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNewBalanceOrigParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NewBalanceOrigParserInit() {
	staticData := &newbalanceorigParserStaticData
	staticData.once.Do(newbalanceorigParserInit)
}

// NewNewBalanceOrigParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNewBalanceOrigParser(input antlr.TokenStream) *NewBalanceOrigParser {
	NewBalanceOrigParserInit()
	this := new(NewBalanceOrigParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &newbalanceorigParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NewBalanceOrig.g4"

	return this
}

// NewBalanceOrigParser tokens.
const (
	NewBalanceOrigParserEOF            = antlr.TokenEOF
	NewBalanceOrigParserNEWBALANCEORIG = 1
	NewBalanceOrigParserOWNKEY         = 2
	NewBalanceOrigParserSPLITKEY       = 3
	NewBalanceOrigParserWS             = 4
)

// NewBalanceOrigParser rules.
const (
	NewBalanceOrigParserRULE_expression     = 0
	NewBalanceOrigParserRULE_newbalanceorig = 1
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
	p.RuleIndex = NewBalanceOrigParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NewBalanceOrigParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Newbalanceorig() INewbalanceorigContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INewbalanceorigContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INewbalanceorigContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NewBalanceOrigParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewBalanceOrigListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewBalanceOrigListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NewBalanceOrigParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NewBalanceOrigParserRULE_expression)

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
		p.Newbalanceorig()
	}
	{
		p.SetState(5)
		p.Match(NewBalanceOrigParserEOF)
	}

	return localctx
}

// INewbalanceorigContext is an interface to support dynamic dispatch.
type INewbalanceorigContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewbalanceorigContext differentiates from other interfaces.
	IsNewbalanceorigContext()
}

type NewbalanceorigContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewbalanceorigContext() *NewbalanceorigContext {
	var p = new(NewbalanceorigContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NewBalanceOrigParserRULE_newbalanceorig
	return p
}

func (*NewbalanceorigContext) IsNewbalanceorigContext() {}

func NewNewbalanceorigContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewbalanceorigContext {
	var p = new(NewbalanceorigContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NewBalanceOrigParserRULE_newbalanceorig

	return p
}

func (s *NewbalanceorigContext) GetParser() antlr.Parser { return s.parser }

func (s *NewbalanceorigContext) NEWBALANCEORIG() antlr.TerminalNode {
	return s.GetToken(NewBalanceOrigParserNEWBALANCEORIG, 0)
}

func (s *NewbalanceorigContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewbalanceorigContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewbalanceorigContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewBalanceOrigListener); ok {
		listenerT.EnterNewbalanceorig(s)
	}
}

func (s *NewbalanceorigContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewBalanceOrigListener); ok {
		listenerT.ExitNewbalanceorig(s)
	}
}

func (p *NewBalanceOrigParser) Newbalanceorig() (localctx INewbalanceorigContext) {
	this := p
	_ = this

	localctx = NewNewbalanceorigContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NewBalanceOrigParserRULE_newbalanceorig)

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
		p.Match(NewBalanceOrigParserNEWBALANCEORIG)
	}

	return localctx
}
