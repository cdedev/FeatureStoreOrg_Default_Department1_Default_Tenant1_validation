// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674799725772/OperatingSystem.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OperatingSystem

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

type OperatingSystemParser struct {
	*antlr.BaseParser
}

var operatingsystemParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func operatingsystemParserInit() {
	staticData := &operatingsystemParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OPERATINGSYSTEM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "operatingsystem",
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

// OperatingSystemParserInit initializes any static state used to implement OperatingSystemParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOperatingSystemParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OperatingSystemParserInit() {
	staticData := &operatingsystemParserStaticData
	staticData.once.Do(operatingsystemParserInit)
}

// NewOperatingSystemParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOperatingSystemParser(input antlr.TokenStream) *OperatingSystemParser {
	OperatingSystemParserInit()
	this := new(OperatingSystemParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &operatingsystemParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "OperatingSystem.g4"

	return this
}

// OperatingSystemParser tokens.
const (
	OperatingSystemParserEOF             = antlr.TokenEOF
	OperatingSystemParserOPERATINGSYSTEM = 1
	OperatingSystemParserOWNKEY          = 2
	OperatingSystemParserSPLITKEY        = 3
	OperatingSystemParserWS              = 4
)

// OperatingSystemParser rules.
const (
	OperatingSystemParserRULE_expression      = 0
	OperatingSystemParserRULE_operatingsystem = 1
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
	p.RuleIndex = OperatingSystemParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OperatingSystemParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Operatingsystem() IOperatingsystemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatingsystemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatingsystemContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OperatingSystemParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OperatingSystemListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OperatingSystemListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OperatingSystemParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OperatingSystemParserRULE_expression)

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
		p.Operatingsystem()
	}
	{
		p.SetState(5)
		p.Match(OperatingSystemParserEOF)
	}

	return localctx
}

// IOperatingsystemContext is an interface to support dynamic dispatch.
type IOperatingsystemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperatingsystemContext differentiates from other interfaces.
	IsOperatingsystemContext()
}

type OperatingsystemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatingsystemContext() *OperatingsystemContext {
	var p = new(OperatingsystemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OperatingSystemParserRULE_operatingsystem
	return p
}

func (*OperatingsystemContext) IsOperatingsystemContext() {}

func NewOperatingsystemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatingsystemContext {
	var p = new(OperatingsystemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OperatingSystemParserRULE_operatingsystem

	return p
}

func (s *OperatingsystemContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatingsystemContext) OPERATINGSYSTEM() antlr.TerminalNode {
	return s.GetToken(OperatingSystemParserOPERATINGSYSTEM, 0)
}

func (s *OperatingsystemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatingsystemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatingsystemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OperatingSystemListener); ok {
		listenerT.EnterOperatingsystem(s)
	}
}

func (s *OperatingsystemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OperatingSystemListener); ok {
		listenerT.ExitOperatingsystem(s)
	}
}

func (p *OperatingSystemParser) Operatingsystem() (localctx IOperatingsystemContext) {
	this := p
	_ = this

	localctx = NewOperatingsystemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OperatingSystemParserRULE_operatingsystem)

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
		p.Match(OperatingSystemParserOPERATINGSYSTEM)
	}

	return localctx
}
