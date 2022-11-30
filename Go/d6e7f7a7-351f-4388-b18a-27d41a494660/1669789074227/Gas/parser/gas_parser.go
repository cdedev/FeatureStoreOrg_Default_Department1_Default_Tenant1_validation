// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789074227/Gas.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gas

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

type GasParser struct {
	*antlr.BaseParser
}

var gasParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gasParserInit() {
	staticData := &gasParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GAS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gas",
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

// GasParserInit initializes any static state used to implement GasParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGasParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GasParserInit() {
	staticData := &gasParserStaticData
	staticData.once.Do(gasParserInit)
}

// NewGasParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGasParser(input antlr.TokenStream) *GasParser {
	GasParserInit()
	this := new(GasParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gasParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Gas.g4"

	return this
}

// GasParser tokens.
const (
	GasParserEOF      = antlr.TokenEOF
	GasParserGAS      = 1
	GasParserOWNKEY   = 2
	GasParserSPLITKEY = 3
	GasParserWS       = 4
)

// GasParser rules.
const (
	GasParserRULE_expression = 0
	GasParserRULE_gas        = 1
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
	p.RuleIndex = GasParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GasParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gas() IGasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGasContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GasParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GasListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GasListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GasParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GasParserRULE_expression)

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
		p.Gas()
	}
	{
		p.SetState(5)
		p.Match(GasParserEOF)
	}

	return localctx
}

// IGasContext is an interface to support dynamic dispatch.
type IGasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGasContext differentiates from other interfaces.
	IsGasContext()
}

type GasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGasContext() *GasContext {
	var p = new(GasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GasParserRULE_gas
	return p
}

func (*GasContext) IsGasContext() {}

func NewGasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GasContext {
	var p = new(GasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GasParserRULE_gas

	return p
}

func (s *GasContext) GetParser() antlr.Parser { return s.parser }

func (s *GasContext) GAS() antlr.TerminalNode {
	return s.GetToken(GasParserGAS, 0)
}

func (s *GasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GasListener); ok {
		listenerT.EnterGas(s)
	}
}

func (s *GasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GasListener); ok {
		listenerT.ExitGas(s)
	}
}

func (p *GasParser) Gas() (localctx IGasContext) {
	this := p
	_ = this

	localctx = NewGasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GasParserRULE_gas)

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
		p.Match(GasParserGAS)
	}

	return localctx
}
