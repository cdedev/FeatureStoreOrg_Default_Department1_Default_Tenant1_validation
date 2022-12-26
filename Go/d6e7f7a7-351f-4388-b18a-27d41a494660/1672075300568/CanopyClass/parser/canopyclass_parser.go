// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672075300568/CanopyClass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CanopyClass

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

type CanopyClassParser struct {
	*antlr.BaseParser
}

var canopyclassParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func canopyclassParserInit() {
	staticData := &canopyclassParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CANOPYCLASS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "canopyclass",
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

// CanopyClassParserInit initializes any static state used to implement CanopyClassParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCanopyClassParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CanopyClassParserInit() {
	staticData := &canopyclassParserStaticData
	staticData.once.Do(canopyclassParserInit)
}

// NewCanopyClassParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCanopyClassParser(input antlr.TokenStream) *CanopyClassParser {
	CanopyClassParserInit()
	this := new(CanopyClassParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &canopyclassParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CanopyClass.g4"

	return this
}

// CanopyClassParser tokens.
const (
	CanopyClassParserEOF         = antlr.TokenEOF
	CanopyClassParserCANOPYCLASS = 1
	CanopyClassParserOWNKEY      = 2
	CanopyClassParserSPLITKEY    = 3
	CanopyClassParserWS          = 4
)

// CanopyClassParser rules.
const (
	CanopyClassParserRULE_expression  = 0
	CanopyClassParserRULE_canopyclass = 1
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
	p.RuleIndex = CanopyClassParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CanopyClassParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Canopyclass() ICanopyclassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICanopyclassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICanopyclassContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CanopyClassParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CanopyClassListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CanopyClassListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CanopyClassParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CanopyClassParserRULE_expression)

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
		p.Canopyclass()
	}
	{
		p.SetState(5)
		p.Match(CanopyClassParserEOF)
	}

	return localctx
}

// ICanopyclassContext is an interface to support dynamic dispatch.
type ICanopyclassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCanopyclassContext differentiates from other interfaces.
	IsCanopyclassContext()
}

type CanopyclassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCanopyclassContext() *CanopyclassContext {
	var p = new(CanopyclassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CanopyClassParserRULE_canopyclass
	return p
}

func (*CanopyclassContext) IsCanopyclassContext() {}

func NewCanopyclassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CanopyclassContext {
	var p = new(CanopyclassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CanopyClassParserRULE_canopyclass

	return p
}

func (s *CanopyclassContext) GetParser() antlr.Parser { return s.parser }

func (s *CanopyclassContext) CANOPYCLASS() antlr.TerminalNode {
	return s.GetToken(CanopyClassParserCANOPYCLASS, 0)
}

func (s *CanopyclassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CanopyclassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CanopyclassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CanopyClassListener); ok {
		listenerT.EnterCanopyclass(s)
	}
}

func (s *CanopyclassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CanopyClassListener); ok {
		listenerT.ExitCanopyclass(s)
	}
}

func (p *CanopyClassParser) Canopyclass() (localctx ICanopyclassContext) {
	this := p
	_ = this

	localctx = NewCanopyclassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CanopyClassParserRULE_canopyclass)

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
		p.Match(CanopyClassParserCANOPYCLASS)
	}

	return localctx
}
