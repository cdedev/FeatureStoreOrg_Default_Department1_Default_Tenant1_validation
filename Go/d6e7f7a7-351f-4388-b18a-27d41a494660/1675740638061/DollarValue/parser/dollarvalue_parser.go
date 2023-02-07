// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675740638061/DollarValue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DollarValue

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

type DollarValueParser struct {
	*antlr.BaseParser
}

var dollarvalueParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dollarvalueParserInit() {
	staticData := &dollarvalueParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DOLLARVALUE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dollarvalue",
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

// DollarValueParserInit initializes any static state used to implement DollarValueParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDollarValueParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DollarValueParserInit() {
	staticData := &dollarvalueParserStaticData
	staticData.once.Do(dollarvalueParserInit)
}

// NewDollarValueParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDollarValueParser(input antlr.TokenStream) *DollarValueParser {
	DollarValueParserInit()
	this := new(DollarValueParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dollarvalueParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DollarValue.g4"

	return this
}

// DollarValueParser tokens.
const (
	DollarValueParserEOF         = antlr.TokenEOF
	DollarValueParserDOLLARVALUE = 1
	DollarValueParserOWNKEY      = 2
	DollarValueParserSPLITKEY    = 3
	DollarValueParserWS          = 4
)

// DollarValueParser rules.
const (
	DollarValueParserRULE_expression  = 0
	DollarValueParserRULE_dollarvalue = 1
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
	p.RuleIndex = DollarValueParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DollarValueParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dollarvalue() IDollarvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDollarvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDollarvalueContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DollarValueParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DollarValueListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DollarValueListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DollarValueParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DollarValueParserRULE_expression)

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
		p.Dollarvalue()
	}
	{
		p.SetState(5)
		p.Match(DollarValueParserEOF)
	}

	return localctx
}

// IDollarvalueContext is an interface to support dynamic dispatch.
type IDollarvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDollarvalueContext differentiates from other interfaces.
	IsDollarvalueContext()
}

type DollarvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDollarvalueContext() *DollarvalueContext {
	var p = new(DollarvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DollarValueParserRULE_dollarvalue
	return p
}

func (*DollarvalueContext) IsDollarvalueContext() {}

func NewDollarvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DollarvalueContext {
	var p = new(DollarvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DollarValueParserRULE_dollarvalue

	return p
}

func (s *DollarvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *DollarvalueContext) DOLLARVALUE() antlr.TerminalNode {
	return s.GetToken(DollarValueParserDOLLARVALUE, 0)
}

func (s *DollarvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DollarvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DollarvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DollarValueListener); ok {
		listenerT.EnterDollarvalue(s)
	}
}

func (s *DollarvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DollarValueListener); ok {
		listenerT.ExitDollarvalue(s)
	}
}

func (p *DollarValueParser) Dollarvalue() (localctx IDollarvalueContext) {
	this := p
	_ = this

	localctx = NewDollarvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DollarValueParserRULE_dollarvalue)

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
		p.Match(DollarValueParserDOLLARVALUE)
	}

	return localctx
}
