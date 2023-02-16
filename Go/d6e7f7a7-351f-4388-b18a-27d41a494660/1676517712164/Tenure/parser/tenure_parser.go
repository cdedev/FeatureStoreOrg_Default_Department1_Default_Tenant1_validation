// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676517712164/Tenure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tenure

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

type TenureParser struct {
	*antlr.BaseParser
}

var tenureParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tenureParserInit() {
	staticData := &tenureParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TENURE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "tenure",
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

// TenureParserInit initializes any static state used to implement TenureParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTenureParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TenureParserInit() {
	staticData := &tenureParserStaticData
	staticData.once.Do(tenureParserInit)
}

// NewTenureParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTenureParser(input antlr.TokenStream) *TenureParser {
	TenureParserInit()
	this := new(TenureParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tenureParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tenure.g4"

	return this
}

// TenureParser tokens.
const (
	TenureParserEOF      = antlr.TokenEOF
	TenureParserTENURE   = 1
	TenureParserOWNKEY   = 2
	TenureParserSPLITKEY = 3
	TenureParserWS       = 4
)

// TenureParser rules.
const (
	TenureParserRULE_expression = 0
	TenureParserRULE_tenure     = 1
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
	p.RuleIndex = TenureParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TenureParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tenure() ITenureContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITenureContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITenureContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TenureParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TenureListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TenureListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TenureParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TenureParserRULE_expression)

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
		p.Tenure()
	}
	{
		p.SetState(5)
		p.Match(TenureParserEOF)
	}

	return localctx
}

// ITenureContext is an interface to support dynamic dispatch.
type ITenureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTenureContext differentiates from other interfaces.
	IsTenureContext()
}

type TenureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTenureContext() *TenureContext {
	var p = new(TenureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TenureParserRULE_tenure
	return p
}

func (*TenureContext) IsTenureContext() {}

func NewTenureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TenureContext {
	var p = new(TenureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TenureParserRULE_tenure

	return p
}

func (s *TenureContext) GetParser() antlr.Parser { return s.parser }

func (s *TenureContext) TENURE() antlr.TerminalNode {
	return s.GetToken(TenureParserTENURE, 0)
}

func (s *TenureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TenureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TenureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TenureListener); ok {
		listenerT.EnterTenure(s)
	}
}

func (s *TenureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TenureListener); ok {
		listenerT.ExitTenure(s)
	}
}

func (p *TenureParser) Tenure() (localctx ITenureContext) {
	this := p
	_ = this

	localctx = NewTenureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TenureParserRULE_tenure)

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
		p.Match(TenureParserTENURE)
	}

	return localctx
}
