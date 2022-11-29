// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699051579/Diameter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diameter

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

type DiameterParser struct {
	*antlr.BaseParser
}

var diameterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diameterParserInit() {
	staticData := &diameterParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DIAMETER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "diameter",
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

// DiameterParserInit initializes any static state used to implement DiameterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiameterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiameterParserInit() {
	staticData := &diameterParserStaticData
	staticData.once.Do(diameterParserInit)
}

// NewDiameterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiameterParser(input antlr.TokenStream) *DiameterParser {
	DiameterParserInit()
	this := new(DiameterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diameterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Diameter.g4"

	return this
}

// DiameterParser tokens.
const (
	DiameterParserEOF      = antlr.TokenEOF
	DiameterParserDIAMETER = 1
	DiameterParserOWNKEY   = 2
	DiameterParserSPLITKEY = 3
	DiameterParserWS       = 4
)

// DiameterParser rules.
const (
	DiameterParserRULE_expression = 0
	DiameterParserRULE_diameter   = 1
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
	p.RuleIndex = DiameterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiameterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Diameter() IDiameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiameterContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiameterParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiameterListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiameterListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiameterParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiameterParserRULE_expression)

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
		p.Diameter()
	}
	{
		p.SetState(5)
		p.Match(DiameterParserEOF)
	}

	return localctx
}

// IDiameterContext is an interface to support dynamic dispatch.
type IDiameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiameterContext differentiates from other interfaces.
	IsDiameterContext()
}

type DiameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiameterContext() *DiameterContext {
	var p = new(DiameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiameterParserRULE_diameter
	return p
}

func (*DiameterContext) IsDiameterContext() {}

func NewDiameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiameterContext {
	var p = new(DiameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiameterParserRULE_diameter

	return p
}

func (s *DiameterContext) GetParser() antlr.Parser { return s.parser }

func (s *DiameterContext) DIAMETER() antlr.TerminalNode {
	return s.GetToken(DiameterParserDIAMETER, 0)
}

func (s *DiameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiameterListener); ok {
		listenerT.EnterDiameter(s)
	}
}

func (s *DiameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiameterListener); ok {
		listenerT.ExitDiameter(s)
	}
}

func (p *DiameterParser) Diameter() (localctx IDiameterContext) {
	this := p
	_ = this

	localctx = NewDiameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiameterParserRULE_diameter)

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
		p.Match(DiameterParserDIAMETER)
	}

	return localctx
}
