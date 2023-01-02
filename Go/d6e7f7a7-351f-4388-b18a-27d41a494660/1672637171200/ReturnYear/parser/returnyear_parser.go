// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637171200/ReturnYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReturnYear

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

type ReturnYearParser struct {
	*antlr.BaseParser
}

var returnyearParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func returnyearParserInit() {
	staticData := &returnyearParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RETURNYEAR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "returnyear",
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

// ReturnYearParserInit initializes any static state used to implement ReturnYearParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewReturnYearParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ReturnYearParserInit() {
	staticData := &returnyearParserStaticData
	staticData.once.Do(returnyearParserInit)
}

// NewReturnYearParser produces a new parser instance for the optional input antlr.TokenStream.
func NewReturnYearParser(input antlr.TokenStream) *ReturnYearParser {
	ReturnYearParserInit()
	this := new(ReturnYearParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &returnyearParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ReturnYear.g4"

	return this
}

// ReturnYearParser tokens.
const (
	ReturnYearParserEOF        = antlr.TokenEOF
	ReturnYearParserRETURNYEAR = 1
	ReturnYearParserOWNKEY     = 2
	ReturnYearParserSPLITKEY   = 3
	ReturnYearParserWS         = 4
)

// ReturnYearParser rules.
const (
	ReturnYearParserRULE_expression = 0
	ReturnYearParserRULE_returnyear = 1
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
	p.RuleIndex = ReturnYearParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReturnYearParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Returnyear() IReturnyearContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnyearContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnyearContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ReturnYearParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReturnYearListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReturnYearListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ReturnYearParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReturnYearParserRULE_expression)

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
		p.Returnyear()
	}
	{
		p.SetState(5)
		p.Match(ReturnYearParserEOF)
	}

	return localctx
}

// IReturnyearContext is an interface to support dynamic dispatch.
type IReturnyearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnyearContext differentiates from other interfaces.
	IsReturnyearContext()
}

type ReturnyearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnyearContext() *ReturnyearContext {
	var p = new(ReturnyearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReturnYearParserRULE_returnyear
	return p
}

func (*ReturnyearContext) IsReturnyearContext() {}

func NewReturnyearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnyearContext {
	var p = new(ReturnyearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReturnYearParserRULE_returnyear

	return p
}

func (s *ReturnyearContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnyearContext) RETURNYEAR() antlr.TerminalNode {
	return s.GetToken(ReturnYearParserRETURNYEAR, 0)
}

func (s *ReturnyearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnyearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnyearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReturnYearListener); ok {
		listenerT.EnterReturnyear(s)
	}
}

func (s *ReturnyearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReturnYearListener); ok {
		listenerT.ExitReturnyear(s)
	}
}

func (p *ReturnYearParser) Returnyear() (localctx IReturnyearContext) {
	this := p
	_ = this

	localctx = NewReturnyearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReturnYearParserRULE_returnyear)

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
		p.Match(ReturnYearParserRETURNYEAR)
	}

	return localctx
}
