// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925818010/Tail.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tail

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

type TailParser struct {
	*antlr.BaseParser
}

var tailParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tailParserInit() {
	staticData := &tailParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TAIL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "tail",
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

// TailParserInit initializes any static state used to implement TailParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTailParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TailParserInit() {
	staticData := &tailParserStaticData
	staticData.once.Do(tailParserInit)
}

// NewTailParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTailParser(input antlr.TokenStream) *TailParser {
	TailParserInit()
	this := new(TailParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tailParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tail.g4"

	return this
}

// TailParser tokens.
const (
	TailParserEOF      = antlr.TokenEOF
	TailParserTAIL     = 1
	TailParserOWNKEY   = 2
	TailParserSPLITKEY = 3
	TailParserWS       = 4
)

// TailParser rules.
const (
	TailParserRULE_expression = 0
	TailParserRULE_tail       = 1
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
	p.RuleIndex = TailParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TailParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tail() ITailContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITailContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITailContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TailParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TailListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TailListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TailParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TailParserRULE_expression)

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
		p.Tail()
	}
	{
		p.SetState(5)
		p.Match(TailParserEOF)
	}

	return localctx
}

// ITailContext is an interface to support dynamic dispatch.
type ITailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTailContext differentiates from other interfaces.
	IsTailContext()
}

type TailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTailContext() *TailContext {
	var p = new(TailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TailParserRULE_tail
	return p
}

func (*TailContext) IsTailContext() {}

func NewTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TailContext {
	var p = new(TailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TailParserRULE_tail

	return p
}

func (s *TailContext) GetParser() antlr.Parser { return s.parser }

func (s *TailContext) TAIL() antlr.TerminalNode {
	return s.GetToken(TailParserTAIL, 0)
}

func (s *TailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TailListener); ok {
		listenerT.EnterTail(s)
	}
}

func (s *TailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TailListener); ok {
		listenerT.ExitTail(s)
	}
}

func (p *TailParser) Tail() (localctx ITailContext) {
	this := p
	_ = this

	localctx = NewTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TailParserRULE_tail)

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
		p.Match(TailParserTAIL)
	}

	return localctx
}
