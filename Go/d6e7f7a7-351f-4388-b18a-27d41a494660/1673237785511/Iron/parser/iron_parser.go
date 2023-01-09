// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237785511/Iron.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Iron

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

type IronParser struct {
	*antlr.BaseParser
}

var ironParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ironParserInit() {
	staticData := &ironParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "IRON", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "iron",
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

// IronParserInit initializes any static state used to implement IronParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIronParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IronParserInit() {
	staticData := &ironParserStaticData
	staticData.once.Do(ironParserInit)
}

// NewIronParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIronParser(input antlr.TokenStream) *IronParser {
	IronParserInit()
	this := new(IronParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ironParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Iron.g4"

	return this
}

// IronParser tokens.
const (
	IronParserEOF      = antlr.TokenEOF
	IronParserIRON     = 1
	IronParserOWNKEY   = 2
	IronParserSPLITKEY = 3
	IronParserWS       = 4
)

// IronParser rules.
const (
	IronParserRULE_expression = 0
	IronParserRULE_iron       = 1
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
	p.RuleIndex = IronParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Iron() IIronContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIronContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIronContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IronParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IronParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IronParserRULE_expression)

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
		p.Iron()
	}
	{
		p.SetState(5)
		p.Match(IronParserEOF)
	}

	return localctx
}

// IIronContext is an interface to support dynamic dispatch.
type IIronContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIronContext differentiates from other interfaces.
	IsIronContext()
}

type IronContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIronContext() *IronContext {
	var p = new(IronContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IronParserRULE_iron
	return p
}

func (*IronContext) IsIronContext() {}

func NewIronContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IronContext {
	var p = new(IronContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IronParserRULE_iron

	return p
}

func (s *IronContext) GetParser() antlr.Parser { return s.parser }

func (s *IronContext) IRON() antlr.TerminalNode {
	return s.GetToken(IronParserIRON, 0)
}

func (s *IronContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IronContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IronContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronListener); ok {
		listenerT.EnterIron(s)
	}
}

func (s *IronContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IronListener); ok {
		listenerT.ExitIron(s)
	}
}

func (p *IronParser) Iron() (localctx IIronContext) {
	this := p
	_ = this

	localctx = NewIronContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IronParserRULE_iron)

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
		p.Match(IronParserIRON)
	}

	return localctx
}
