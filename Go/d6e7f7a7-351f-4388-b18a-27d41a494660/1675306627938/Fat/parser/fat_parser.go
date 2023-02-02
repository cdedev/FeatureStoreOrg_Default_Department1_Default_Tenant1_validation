// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675306627938/Fat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fat

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

type FatParser struct {
	*antlr.BaseParser
}

var fatParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func fatParserInit() {
	staticData := &fatParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FAT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "fat",
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

// FatParserInit initializes any static state used to implement FatParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFatParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FatParserInit() {
	staticData := &fatParserStaticData
	staticData.once.Do(fatParserInit)
}

// NewFatParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFatParser(input antlr.TokenStream) *FatParser {
	FatParserInit()
	this := new(FatParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &fatParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Fat.g4"

	return this
}

// FatParser tokens.
const (
	FatParserEOF      = antlr.TokenEOF
	FatParserFAT      = 1
	FatParserOWNKEY   = 2
	FatParserSPLITKEY = 3
	FatParserWS       = 4
)

// FatParser rules.
const (
	FatParserRULE_expression = 0
	FatParserRULE_fat        = 1
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
	p.RuleIndex = FatParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FatParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Fat() IFatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFatContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FatParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FatListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FatListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FatParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FatParserRULE_expression)

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
		p.Fat()
	}
	{
		p.SetState(5)
		p.Match(FatParserEOF)
	}

	return localctx
}

// IFatContext is an interface to support dynamic dispatch.
type IFatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFatContext differentiates from other interfaces.
	IsFatContext()
}

type FatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFatContext() *FatContext {
	var p = new(FatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FatParserRULE_fat
	return p
}

func (*FatContext) IsFatContext() {}

func NewFatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FatContext {
	var p = new(FatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FatParserRULE_fat

	return p
}

func (s *FatContext) GetParser() antlr.Parser { return s.parser }

func (s *FatContext) FAT() antlr.TerminalNode {
	return s.GetToken(FatParserFAT, 0)
}

func (s *FatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FatListener); ok {
		listenerT.EnterFat(s)
	}
}

func (s *FatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FatListener); ok {
		listenerT.ExitFat(s)
	}
}

func (p *FatParser) Fat() (localctx IFatContext) {
	this := p
	_ = this

	localctx = NewFatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FatParserRULE_fat)

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
		p.Match(FatParserFAT)
	}

	return localctx
}
