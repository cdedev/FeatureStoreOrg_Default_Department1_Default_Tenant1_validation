// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669981590412/Date_of_birth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Date_of_birth

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

type Date_of_birthParser struct {
	*antlr.BaseParser
}

var date_of_birthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func date_of_birthParserInit() {
	staticData := &date_of_birthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DATE_OF_BIRTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "date_of_birth",
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

// Date_of_birthParserInit initializes any static state used to implement Date_of_birthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDate_of_birthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Date_of_birthParserInit() {
	staticData := &date_of_birthParserStaticData
	staticData.once.Do(date_of_birthParserInit)
}

// NewDate_of_birthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDate_of_birthParser(input antlr.TokenStream) *Date_of_birthParser {
	Date_of_birthParserInit()
	this := new(Date_of_birthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &date_of_birthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Date_of_birth.g4"

	return this
}

// Date_of_birthParser tokens.
const (
	Date_of_birthParserEOF           = antlr.TokenEOF
	Date_of_birthParserDATE_OF_BIRTH = 1
	Date_of_birthParserOWNKEY        = 2
	Date_of_birthParserSPLITKEY      = 3
	Date_of_birthParserWS            = 4
)

// Date_of_birthParser rules.
const (
	Date_of_birthParserRULE_expression    = 0
	Date_of_birthParserRULE_date_of_birth = 1
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
	p.RuleIndex = Date_of_birthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Date_of_birthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Date_of_birth() IDate_of_birthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDate_of_birthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDate_of_birthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Date_of_birthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Date_of_birthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Date_of_birthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Date_of_birthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Date_of_birthParserRULE_expression)

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
		p.Date_of_birth()
	}
	{
		p.SetState(5)
		p.Match(Date_of_birthParserEOF)
	}

	return localctx
}

// IDate_of_birthContext is an interface to support dynamic dispatch.
type IDate_of_birthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_of_birthContext differentiates from other interfaces.
	IsDate_of_birthContext()
}

type Date_of_birthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_of_birthContext() *Date_of_birthContext {
	var p = new(Date_of_birthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Date_of_birthParserRULE_date_of_birth
	return p
}

func (*Date_of_birthContext) IsDate_of_birthContext() {}

func NewDate_of_birthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_of_birthContext {
	var p = new(Date_of_birthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Date_of_birthParserRULE_date_of_birth

	return p
}

func (s *Date_of_birthContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_of_birthContext) DATE_OF_BIRTH() antlr.TerminalNode {
	return s.GetToken(Date_of_birthParserDATE_OF_BIRTH, 0)
}

func (s *Date_of_birthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_of_birthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_of_birthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Date_of_birthListener); ok {
		listenerT.EnterDate_of_birth(s)
	}
}

func (s *Date_of_birthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Date_of_birthListener); ok {
		listenerT.ExitDate_of_birth(s)
	}
}

func (p *Date_of_birthParser) Date_of_birth() (localctx IDate_of_birthContext) {
	this := p
	_ = this

	localctx = NewDate_of_birthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Date_of_birthParserRULE_date_of_birth)

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
		p.Match(Date_of_birthParserDATE_OF_BIRTH)
	}

	return localctx
}
