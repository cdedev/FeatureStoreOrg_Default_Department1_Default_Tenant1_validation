// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623938460/Street.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Street

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

type StreetParser struct {
	*antlr.BaseParser
}

var streetParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func streetParserInit() {
	staticData := &streetParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STREET", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "street",
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

// StreetParserInit initializes any static state used to implement StreetParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStreetParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StreetParserInit() {
	staticData := &streetParserStaticData
	staticData.once.Do(streetParserInit)
}

// NewStreetParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStreetParser(input antlr.TokenStream) *StreetParser {
	StreetParserInit()
	this := new(StreetParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &streetParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Street.g4"

	return this
}

// StreetParser tokens.
const (
	StreetParserEOF      = antlr.TokenEOF
	StreetParserSTREET   = 1
	StreetParserOWNKEY   = 2
	StreetParserSPLITKEY = 3
	StreetParserWS       = 4
)

// StreetParser rules.
const (
	StreetParserRULE_expression = 0
	StreetParserRULE_street     = 1
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
	p.RuleIndex = StreetParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StreetParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Street() IStreetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStreetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStreetContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StreetParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreetListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreetListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StreetParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StreetParserRULE_expression)

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
		p.Street()
	}
	{
		p.SetState(5)
		p.Match(StreetParserEOF)
	}

	return localctx
}

// IStreetContext is an interface to support dynamic dispatch.
type IStreetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreetContext differentiates from other interfaces.
	IsStreetContext()
}

type StreetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreetContext() *StreetContext {
	var p = new(StreetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StreetParserRULE_street
	return p
}

func (*StreetContext) IsStreetContext() {}

func NewStreetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreetContext {
	var p = new(StreetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StreetParserRULE_street

	return p
}

func (s *StreetContext) GetParser() antlr.Parser { return s.parser }

func (s *StreetContext) STREET() antlr.TerminalNode {
	return s.GetToken(StreetParserSTREET, 0)
}

func (s *StreetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreetListener); ok {
		listenerT.EnterStreet(s)
	}
}

func (s *StreetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreetListener); ok {
		listenerT.ExitStreet(s)
	}
}

func (p *StreetParser) Street() (localctx IStreetContext) {
	this := p
	_ = this

	localctx = NewStreetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StreetParserRULE_street)

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
		p.Match(StreetParserSTREET)
	}

	return localctx
}
