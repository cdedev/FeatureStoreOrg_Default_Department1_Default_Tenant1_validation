// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235590013/Axons.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Axons

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

type AxonsParser struct {
	*antlr.BaseParser
}

var axonsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func axonsParserInit() {
	staticData := &axonsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AXONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "axons",
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

// AxonsParserInit initializes any static state used to implement AxonsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAxonsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AxonsParserInit() {
	staticData := &axonsParserStaticData
	staticData.once.Do(axonsParserInit)
}

// NewAxonsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAxonsParser(input antlr.TokenStream) *AxonsParser {
	AxonsParserInit()
	this := new(AxonsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &axonsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Axons.g4"

	return this
}

// AxonsParser tokens.
const (
	AxonsParserEOF      = antlr.TokenEOF
	AxonsParserAXONS    = 1
	AxonsParserOWNKEY   = 2
	AxonsParserSPLITKEY = 3
	AxonsParserWS       = 4
)

// AxonsParser rules.
const (
	AxonsParserRULE_expression = 0
	AxonsParserRULE_axons      = 1
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
	p.RuleIndex = AxonsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AxonsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Axons() IAxonsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAxonsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAxonsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AxonsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AxonsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AxonsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AxonsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AxonsParserRULE_expression)

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
		p.Axons()
	}
	{
		p.SetState(5)
		p.Match(AxonsParserEOF)
	}

	return localctx
}

// IAxonsContext is an interface to support dynamic dispatch.
type IAxonsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxonsContext differentiates from other interfaces.
	IsAxonsContext()
}

type AxonsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxonsContext() *AxonsContext {
	var p = new(AxonsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AxonsParserRULE_axons
	return p
}

func (*AxonsContext) IsAxonsContext() {}

func NewAxonsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AxonsContext {
	var p = new(AxonsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AxonsParserRULE_axons

	return p
}

func (s *AxonsContext) GetParser() antlr.Parser { return s.parser }

func (s *AxonsContext) AXONS() antlr.TerminalNode {
	return s.GetToken(AxonsParserAXONS, 0)
}

func (s *AxonsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AxonsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AxonsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AxonsListener); ok {
		listenerT.EnterAxons(s)
	}
}

func (s *AxonsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AxonsListener); ok {
		listenerT.ExitAxons(s)
	}
}

func (p *AxonsParser) Axons() (localctx IAxonsContext) {
	this := p
	_ = this

	localctx = NewAxonsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AxonsParserRULE_axons)

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
		p.Match(AxonsParserAXONS)
	}

	return localctx
}
