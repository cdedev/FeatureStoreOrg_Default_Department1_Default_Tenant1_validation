// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672197373473/Furnishing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Furnishing

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

type FurnishingParser struct {
	*antlr.BaseParser
}

var furnishingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func furnishingParserInit() {
	staticData := &furnishingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FURNISHING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "furnishing",
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

// FurnishingParserInit initializes any static state used to implement FurnishingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFurnishingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FurnishingParserInit() {
	staticData := &furnishingParserStaticData
	staticData.once.Do(furnishingParserInit)
}

// NewFurnishingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFurnishingParser(input antlr.TokenStream) *FurnishingParser {
	FurnishingParserInit()
	this := new(FurnishingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &furnishingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Furnishing.g4"

	return this
}

// FurnishingParser tokens.
const (
	FurnishingParserEOF        = antlr.TokenEOF
	FurnishingParserFURNISHING = 1
	FurnishingParserOWNKEY     = 2
	FurnishingParserSPLITKEY   = 3
	FurnishingParserWS         = 4
)

// FurnishingParser rules.
const (
	FurnishingParserRULE_expression = 0
	FurnishingParserRULE_furnishing = 1
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
	p.RuleIndex = FurnishingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FurnishingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Furnishing() IFurnishingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFurnishingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFurnishingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FurnishingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FurnishingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FurnishingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FurnishingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FurnishingParserRULE_expression)

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
		p.Furnishing()
	}
	{
		p.SetState(5)
		p.Match(FurnishingParserEOF)
	}

	return localctx
}

// IFurnishingContext is an interface to support dynamic dispatch.
type IFurnishingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFurnishingContext differentiates from other interfaces.
	IsFurnishingContext()
}

type FurnishingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFurnishingContext() *FurnishingContext {
	var p = new(FurnishingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FurnishingParserRULE_furnishing
	return p
}

func (*FurnishingContext) IsFurnishingContext() {}

func NewFurnishingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FurnishingContext {
	var p = new(FurnishingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FurnishingParserRULE_furnishing

	return p
}

func (s *FurnishingContext) GetParser() antlr.Parser { return s.parser }

func (s *FurnishingContext) FURNISHING() antlr.TerminalNode {
	return s.GetToken(FurnishingParserFURNISHING, 0)
}

func (s *FurnishingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FurnishingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FurnishingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FurnishingListener); ok {
		listenerT.EnterFurnishing(s)
	}
}

func (s *FurnishingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FurnishingListener); ok {
		listenerT.ExitFurnishing(s)
	}
}

func (p *FurnishingParser) Furnishing() (localctx IFurnishingContext) {
	this := p
	_ = this

	localctx = NewFurnishingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FurnishingParserRULE_furnishing)

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
		p.Match(FurnishingParserFURNISHING)
	}

	return localctx
}
