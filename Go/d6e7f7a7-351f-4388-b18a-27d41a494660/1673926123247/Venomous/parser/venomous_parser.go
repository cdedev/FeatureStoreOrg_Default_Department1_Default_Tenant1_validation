// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673926123247/Venomous.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Venomous

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

type VenomousParser struct {
	*antlr.BaseParser
}

var venomousParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func venomousParserInit() {
	staticData := &venomousParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VENOMOUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "venomous",
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

// VenomousParserInit initializes any static state used to implement VenomousParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVenomousParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VenomousParserInit() {
	staticData := &venomousParserStaticData
	staticData.once.Do(venomousParserInit)
}

// NewVenomousParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVenomousParser(input antlr.TokenStream) *VenomousParser {
	VenomousParserInit()
	this := new(VenomousParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &venomousParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Venomous.g4"

	return this
}

// VenomousParser tokens.
const (
	VenomousParserEOF      = antlr.TokenEOF
	VenomousParserVENOMOUS = 1
	VenomousParserOWNKEY   = 2
	VenomousParserSPLITKEY = 3
	VenomousParserWS       = 4
)

// VenomousParser rules.
const (
	VenomousParserRULE_expression = 0
	VenomousParserRULE_venomous   = 1
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
	p.RuleIndex = VenomousParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VenomousParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Venomous() IVenomousContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVenomousContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVenomousContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VenomousParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VenomousListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VenomousListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VenomousParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VenomousParserRULE_expression)

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
		p.Venomous()
	}
	{
		p.SetState(5)
		p.Match(VenomousParserEOF)
	}

	return localctx
}

// IVenomousContext is an interface to support dynamic dispatch.
type IVenomousContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVenomousContext differentiates from other interfaces.
	IsVenomousContext()
}

type VenomousContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVenomousContext() *VenomousContext {
	var p = new(VenomousContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VenomousParserRULE_venomous
	return p
}

func (*VenomousContext) IsVenomousContext() {}

func NewVenomousContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VenomousContext {
	var p = new(VenomousContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VenomousParserRULE_venomous

	return p
}

func (s *VenomousContext) GetParser() antlr.Parser { return s.parser }

func (s *VenomousContext) VENOMOUS() antlr.TerminalNode {
	return s.GetToken(VenomousParserVENOMOUS, 0)
}

func (s *VenomousContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VenomousContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VenomousContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VenomousListener); ok {
		listenerT.EnterVenomous(s)
	}
}

func (s *VenomousContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VenomousListener); ok {
		listenerT.ExitVenomous(s)
	}
}

func (p *VenomousParser) Venomous() (localctx IVenomousContext) {
	this := p
	_ = this

	localctx = NewVenomousContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VenomousParserRULE_venomous)

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
		p.Match(VenomousParserVENOMOUS)
	}

	return localctx
}
