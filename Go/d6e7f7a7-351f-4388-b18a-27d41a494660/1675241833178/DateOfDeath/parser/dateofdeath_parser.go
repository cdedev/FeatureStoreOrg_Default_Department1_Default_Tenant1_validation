// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241833178/DateOfDeath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DateOfDeath

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

type DateOfDeathParser struct {
	*antlr.BaseParser
}

var dateofdeathParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dateofdeathParserInit() {
	staticData := &dateofdeathParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DATEOFDEATH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dateofdeath",
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

// DateOfDeathParserInit initializes any static state used to implement DateOfDeathParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDateOfDeathParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DateOfDeathParserInit() {
	staticData := &dateofdeathParserStaticData
	staticData.once.Do(dateofdeathParserInit)
}

// NewDateOfDeathParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDateOfDeathParser(input antlr.TokenStream) *DateOfDeathParser {
	DateOfDeathParserInit()
	this := new(DateOfDeathParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dateofdeathParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DateOfDeath.g4"

	return this
}

// DateOfDeathParser tokens.
const (
	DateOfDeathParserEOF         = antlr.TokenEOF
	DateOfDeathParserDATEOFDEATH = 1
	DateOfDeathParserOWNKEY      = 2
	DateOfDeathParserSPLITKEY    = 3
	DateOfDeathParserWS          = 4
)

// DateOfDeathParser rules.
const (
	DateOfDeathParserRULE_expression  = 0
	DateOfDeathParserRULE_dateofdeath = 1
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
	p.RuleIndex = DateOfDeathParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DateOfDeathParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dateofdeath() IDateofdeathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateofdeathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateofdeathContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DateOfDeathParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfDeathListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfDeathListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DateOfDeathParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DateOfDeathParserRULE_expression)

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
		p.Dateofdeath()
	}
	{
		p.SetState(5)
		p.Match(DateOfDeathParserEOF)
	}

	return localctx
}

// IDateofdeathContext is an interface to support dynamic dispatch.
type IDateofdeathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateofdeathContext differentiates from other interfaces.
	IsDateofdeathContext()
}

type DateofdeathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateofdeathContext() *DateofdeathContext {
	var p = new(DateofdeathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DateOfDeathParserRULE_dateofdeath
	return p
}

func (*DateofdeathContext) IsDateofdeathContext() {}

func NewDateofdeathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateofdeathContext {
	var p = new(DateofdeathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DateOfDeathParserRULE_dateofdeath

	return p
}

func (s *DateofdeathContext) GetParser() antlr.Parser { return s.parser }

func (s *DateofdeathContext) DATEOFDEATH() antlr.TerminalNode {
	return s.GetToken(DateOfDeathParserDATEOFDEATH, 0)
}

func (s *DateofdeathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateofdeathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateofdeathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfDeathListener); ok {
		listenerT.EnterDateofdeath(s)
	}
}

func (s *DateofdeathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfDeathListener); ok {
		listenerT.ExitDateofdeath(s)
	}
}

func (p *DateOfDeathParser) Dateofdeath() (localctx IDateofdeathContext) {
	this := p
	_ = this

	localctx = NewDateofdeathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DateOfDeathParserRULE_dateofdeath)

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
		p.Match(DateOfDeathParserDATEOFDEATH)
	}

	return localctx
}
