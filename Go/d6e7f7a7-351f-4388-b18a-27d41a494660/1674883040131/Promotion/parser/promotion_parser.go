// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674883040131/Promotion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Promotion

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

type PromotionParser struct {
	*antlr.BaseParser
}

var promotionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func promotionParserInit() {
	staticData := &promotionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PROMOTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "promotion",
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

// PromotionParserInit initializes any static state used to implement PromotionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPromotionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PromotionParserInit() {
	staticData := &promotionParserStaticData
	staticData.once.Do(promotionParserInit)
}

// NewPromotionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPromotionParser(input antlr.TokenStream) *PromotionParser {
	PromotionParserInit()
	this := new(PromotionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &promotionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Promotion.g4"

	return this
}

// PromotionParser tokens.
const (
	PromotionParserEOF       = antlr.TokenEOF
	PromotionParserPROMOTION = 1
	PromotionParserOWNKEY    = 2
	PromotionParserSPLITKEY  = 3
	PromotionParserWS        = 4
)

// PromotionParser rules.
const (
	PromotionParserRULE_expression = 0
	PromotionParserRULE_promotion  = 1
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
	p.RuleIndex = PromotionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromotionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Promotion() IPromotionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPromotionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPromotionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PromotionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromotionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromotionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PromotionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PromotionParserRULE_expression)

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
		p.Promotion()
	}
	{
		p.SetState(5)
		p.Match(PromotionParserEOF)
	}

	return localctx
}

// IPromotionContext is an interface to support dynamic dispatch.
type IPromotionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPromotionContext differentiates from other interfaces.
	IsPromotionContext()
}

type PromotionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPromotionContext() *PromotionContext {
	var p = new(PromotionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PromotionParserRULE_promotion
	return p
}

func (*PromotionContext) IsPromotionContext() {}

func NewPromotionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PromotionContext {
	var p = new(PromotionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PromotionParserRULE_promotion

	return p
}

func (s *PromotionContext) GetParser() antlr.Parser { return s.parser }

func (s *PromotionContext) PROMOTION() antlr.TerminalNode {
	return s.GetToken(PromotionParserPROMOTION, 0)
}

func (s *PromotionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PromotionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PromotionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromotionListener); ok {
		listenerT.EnterPromotion(s)
	}
}

func (s *PromotionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PromotionListener); ok {
		listenerT.ExitPromotion(s)
	}
}

func (p *PromotionParser) Promotion() (localctx IPromotionContext) {
	this := p
	_ = this

	localctx = NewPromotionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PromotionParserRULE_promotion)

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
		p.Match(PromotionParserPROMOTION)
	}

	return localctx
}
