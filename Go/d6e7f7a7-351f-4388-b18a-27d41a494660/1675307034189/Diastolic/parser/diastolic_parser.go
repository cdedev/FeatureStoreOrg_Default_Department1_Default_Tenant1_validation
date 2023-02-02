// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675307034189/Diastolic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diastolic

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

type DiastolicParser struct {
	*antlr.BaseParser
}

var diastolicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diastolicParserInit() {
	staticData := &diastolicParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DIASTOLIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "diastolic",
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

// DiastolicParserInit initializes any static state used to implement DiastolicParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiastolicParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiastolicParserInit() {
	staticData := &diastolicParserStaticData
	staticData.once.Do(diastolicParserInit)
}

// NewDiastolicParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiastolicParser(input antlr.TokenStream) *DiastolicParser {
	DiastolicParserInit()
	this := new(DiastolicParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diastolicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Diastolic.g4"

	return this
}

// DiastolicParser tokens.
const (
	DiastolicParserEOF       = antlr.TokenEOF
	DiastolicParserDIASTOLIC = 1
	DiastolicParserOWNKEY    = 2
	DiastolicParserSPLITKEY  = 3
	DiastolicParserWS        = 4
)

// DiastolicParser rules.
const (
	DiastolicParserRULE_expression = 0
	DiastolicParserRULE_diastolic  = 1
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
	p.RuleIndex = DiastolicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiastolicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Diastolic() IDiastolicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiastolicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiastolicContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiastolicParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiastolicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiastolicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiastolicParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiastolicParserRULE_expression)

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
		p.Diastolic()
	}
	{
		p.SetState(5)
		p.Match(DiastolicParserEOF)
	}

	return localctx
}

// IDiastolicContext is an interface to support dynamic dispatch.
type IDiastolicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiastolicContext differentiates from other interfaces.
	IsDiastolicContext()
}

type DiastolicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiastolicContext() *DiastolicContext {
	var p = new(DiastolicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiastolicParserRULE_diastolic
	return p
}

func (*DiastolicContext) IsDiastolicContext() {}

func NewDiastolicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiastolicContext {
	var p = new(DiastolicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiastolicParserRULE_diastolic

	return p
}

func (s *DiastolicContext) GetParser() antlr.Parser { return s.parser }

func (s *DiastolicContext) DIASTOLIC() antlr.TerminalNode {
	return s.GetToken(DiastolicParserDIASTOLIC, 0)
}

func (s *DiastolicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiastolicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiastolicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiastolicListener); ok {
		listenerT.EnterDiastolic(s)
	}
}

func (s *DiastolicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiastolicListener); ok {
		listenerT.ExitDiastolic(s)
	}
}

func (p *DiastolicParser) Diastolic() (localctx IDiastolicContext) {
	this := p
	_ = this

	localctx = NewDiastolicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiastolicParserRULE_diastolic)

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
		p.Match(DiastolicParserDIASTOLIC)
	}

	return localctx
}
