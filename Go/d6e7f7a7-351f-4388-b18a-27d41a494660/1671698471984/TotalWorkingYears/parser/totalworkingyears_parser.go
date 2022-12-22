// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698471984/TotalWorkingYears.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalWorkingYears

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

type TotalWorkingYearsParser struct {
	*antlr.BaseParser
}

var totalworkingyearsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func totalworkingyearsParserInit() {
	staticData := &totalworkingyearsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TOTALWORKINGYEARS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "totalworkingyears",
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

// TotalWorkingYearsParserInit initializes any static state used to implement TotalWorkingYearsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTotalWorkingYearsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TotalWorkingYearsParserInit() {
	staticData := &totalworkingyearsParserStaticData
	staticData.once.Do(totalworkingyearsParserInit)
}

// NewTotalWorkingYearsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTotalWorkingYearsParser(input antlr.TokenStream) *TotalWorkingYearsParser {
	TotalWorkingYearsParserInit()
	this := new(TotalWorkingYearsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &totalworkingyearsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "TotalWorkingYears.g4"

	return this
}

// TotalWorkingYearsParser tokens.
const (
	TotalWorkingYearsParserEOF               = antlr.TokenEOF
	TotalWorkingYearsParserTOTALWORKINGYEARS = 1
	TotalWorkingYearsParserOWNKEY            = 2
	TotalWorkingYearsParserSPLITKEY          = 3
	TotalWorkingYearsParserWS                = 4
)

// TotalWorkingYearsParser rules.
const (
	TotalWorkingYearsParserRULE_expression        = 0
	TotalWorkingYearsParserRULE_totalworkingyears = 1
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
	p.RuleIndex = TotalWorkingYearsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TotalWorkingYearsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Totalworkingyears() ITotalworkingyearsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITotalworkingyearsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITotalworkingyearsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TotalWorkingYearsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalWorkingYearsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalWorkingYearsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TotalWorkingYearsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TotalWorkingYearsParserRULE_expression)

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
		p.Totalworkingyears()
	}
	{
		p.SetState(5)
		p.Match(TotalWorkingYearsParserEOF)
	}

	return localctx
}

// ITotalworkingyearsContext is an interface to support dynamic dispatch.
type ITotalworkingyearsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTotalworkingyearsContext differentiates from other interfaces.
	IsTotalworkingyearsContext()
}

type TotalworkingyearsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTotalworkingyearsContext() *TotalworkingyearsContext {
	var p = new(TotalworkingyearsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TotalWorkingYearsParserRULE_totalworkingyears
	return p
}

func (*TotalworkingyearsContext) IsTotalworkingyearsContext() {}

func NewTotalworkingyearsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotalworkingyearsContext {
	var p = new(TotalworkingyearsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TotalWorkingYearsParserRULE_totalworkingyears

	return p
}

func (s *TotalworkingyearsContext) GetParser() antlr.Parser { return s.parser }

func (s *TotalworkingyearsContext) TOTALWORKINGYEARS() antlr.TerminalNode {
	return s.GetToken(TotalWorkingYearsParserTOTALWORKINGYEARS, 0)
}

func (s *TotalworkingyearsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotalworkingyearsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TotalworkingyearsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalWorkingYearsListener); ok {
		listenerT.EnterTotalworkingyears(s)
	}
}

func (s *TotalworkingyearsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalWorkingYearsListener); ok {
		listenerT.ExitTotalworkingyears(s)
	}
}

func (p *TotalWorkingYearsParser) Totalworkingyears() (localctx ITotalworkingyearsContext) {
	this := p
	_ = this

	localctx = NewTotalworkingyearsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TotalWorkingYearsParserRULE_totalworkingyears)

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
		p.Match(TotalWorkingYearsParserTOTALWORKINGYEARS)
	}

	return localctx
}
