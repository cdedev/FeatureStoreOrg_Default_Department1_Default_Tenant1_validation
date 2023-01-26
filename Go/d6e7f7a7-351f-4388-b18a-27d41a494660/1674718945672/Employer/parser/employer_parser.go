// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674718945672/Employer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Employer

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

type EmployerParser struct {
	*antlr.BaseParser
}

var employerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func employerParserInit() {
	staticData := &employerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EMPLOYER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "employer",
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

// EmployerParserInit initializes any static state used to implement EmployerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEmployerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EmployerParserInit() {
	staticData := &employerParserStaticData
	staticData.once.Do(employerParserInit)
}

// NewEmployerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEmployerParser(input antlr.TokenStream) *EmployerParser {
	EmployerParserInit()
	this := new(EmployerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &employerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Employer.g4"

	return this
}

// EmployerParser tokens.
const (
	EmployerParserEOF      = antlr.TokenEOF
	EmployerParserEMPLOYER = 1
	EmployerParserOWNKEY   = 2
	EmployerParserSPLITKEY = 3
	EmployerParserWS       = 4
)

// EmployerParser rules.
const (
	EmployerParserRULE_expression = 0
	EmployerParserRULE_employer   = 1
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
	p.RuleIndex = EmployerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmployerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Employer() IEmployerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmployerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmployerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EmployerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmployerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmployerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EmployerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EmployerParserRULE_expression)

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
		p.Employer()
	}
	{
		p.SetState(5)
		p.Match(EmployerParserEOF)
	}

	return localctx
}

// IEmployerContext is an interface to support dynamic dispatch.
type IEmployerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmployerContext differentiates from other interfaces.
	IsEmployerContext()
}

type EmployerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmployerContext() *EmployerContext {
	var p = new(EmployerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EmployerParserRULE_employer
	return p
}

func (*EmployerContext) IsEmployerContext() {}

func NewEmployerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmployerContext {
	var p = new(EmployerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmployerParserRULE_employer

	return p
}

func (s *EmployerContext) GetParser() antlr.Parser { return s.parser }

func (s *EmployerContext) EMPLOYER() antlr.TerminalNode {
	return s.GetToken(EmployerParserEMPLOYER, 0)
}

func (s *EmployerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmployerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmployerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmployerListener); ok {
		listenerT.EnterEmployer(s)
	}
}

func (s *EmployerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmployerListener); ok {
		listenerT.ExitEmployer(s)
	}
}

func (p *EmployerParser) Employer() (localctx IEmployerContext) {
	this := p
	_ = this

	localctx = NewEmployerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EmployerParserRULE_employer)

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
		p.Match(EmployerParserEMPLOYER)
	}

	return localctx
}
