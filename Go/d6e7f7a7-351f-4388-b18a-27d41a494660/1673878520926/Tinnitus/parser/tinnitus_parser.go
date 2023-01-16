// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878520926/Tinnitus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tinnitus

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

type TinnitusParser struct {
	*antlr.BaseParser
}

var tinnitusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinnitusParserInit() {
	staticData := &tinnitusParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TINNITUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "tinnitus",
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

// TinnitusParserInit initializes any static state used to implement TinnitusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTinnitusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinnitusParserInit() {
	staticData := &tinnitusParserStaticData
	staticData.once.Do(tinnitusParserInit)
}

// NewTinnitusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTinnitusParser(input antlr.TokenStream) *TinnitusParser {
	TinnitusParserInit()
	this := new(TinnitusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tinnitusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tinnitus.g4"

	return this
}

// TinnitusParser tokens.
const (
	TinnitusParserEOF      = antlr.TokenEOF
	TinnitusParserTINNITUS = 1
	TinnitusParserOWNKEY   = 2
	TinnitusParserSPLITKEY = 3
	TinnitusParserWS       = 4
)

// TinnitusParser rules.
const (
	TinnitusParserRULE_expression = 0
	TinnitusParserRULE_tinnitus   = 1
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
	p.RuleIndex = TinnitusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinnitusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tinnitus() ITinnitusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITinnitusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITinnitusContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TinnitusParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinnitusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinnitusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TinnitusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TinnitusParserRULE_expression)

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
		p.Tinnitus()
	}
	{
		p.SetState(5)
		p.Match(TinnitusParserEOF)
	}

	return localctx
}

// ITinnitusContext is an interface to support dynamic dispatch.
type ITinnitusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTinnitusContext differentiates from other interfaces.
	IsTinnitusContext()
}

type TinnitusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTinnitusContext() *TinnitusContext {
	var p = new(TinnitusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinnitusParserRULE_tinnitus
	return p
}

func (*TinnitusContext) IsTinnitusContext() {}

func NewTinnitusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TinnitusContext {
	var p = new(TinnitusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinnitusParserRULE_tinnitus

	return p
}

func (s *TinnitusContext) GetParser() antlr.Parser { return s.parser }

func (s *TinnitusContext) TINNITUS() antlr.TerminalNode {
	return s.GetToken(TinnitusParserTINNITUS, 0)
}

func (s *TinnitusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TinnitusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TinnitusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinnitusListener); ok {
		listenerT.EnterTinnitus(s)
	}
}

func (s *TinnitusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinnitusListener); ok {
		listenerT.ExitTinnitus(s)
	}
}

func (p *TinnitusParser) Tinnitus() (localctx ITinnitusContext) {
	this := p
	_ = this

	localctx = NewTinnitusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TinnitusParserRULE_tinnitus)

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
		p.Match(TinnitusParserTINNITUS)
	}

	return localctx
}
