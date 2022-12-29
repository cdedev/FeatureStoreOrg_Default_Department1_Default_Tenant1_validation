// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284768398/CartoonReference.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CartoonReference

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

type CartoonReferenceParser struct {
	*antlr.BaseParser
}

var cartoonreferenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cartoonreferenceParserInit() {
	staticData := &cartoonreferenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CARTOONREFERENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cartoonreference",
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

// CartoonReferenceParserInit initializes any static state used to implement CartoonReferenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCartoonReferenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CartoonReferenceParserInit() {
	staticData := &cartoonreferenceParserStaticData
	staticData.once.Do(cartoonreferenceParserInit)
}

// NewCartoonReferenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCartoonReferenceParser(input antlr.TokenStream) *CartoonReferenceParser {
	CartoonReferenceParserInit()
	this := new(CartoonReferenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cartoonreferenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CartoonReference.g4"

	return this
}

// CartoonReferenceParser tokens.
const (
	CartoonReferenceParserEOF              = antlr.TokenEOF
	CartoonReferenceParserCARTOONREFERENCE = 1
	CartoonReferenceParserOWNKEY           = 2
	CartoonReferenceParserSPLITKEY         = 3
	CartoonReferenceParserWS               = 4
)

// CartoonReferenceParser rules.
const (
	CartoonReferenceParserRULE_expression       = 0
	CartoonReferenceParserRULE_cartoonreference = 1
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
	p.RuleIndex = CartoonReferenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CartoonReferenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cartoonreference() ICartoonreferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICartoonreferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICartoonreferenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CartoonReferenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CartoonReferenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CartoonReferenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CartoonReferenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CartoonReferenceParserRULE_expression)

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
		p.Cartoonreference()
	}
	{
		p.SetState(5)
		p.Match(CartoonReferenceParserEOF)
	}

	return localctx
}

// ICartoonreferenceContext is an interface to support dynamic dispatch.
type ICartoonreferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCartoonreferenceContext differentiates from other interfaces.
	IsCartoonreferenceContext()
}

type CartoonreferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCartoonreferenceContext() *CartoonreferenceContext {
	var p = new(CartoonreferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CartoonReferenceParserRULE_cartoonreference
	return p
}

func (*CartoonreferenceContext) IsCartoonreferenceContext() {}

func NewCartoonreferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CartoonreferenceContext {
	var p = new(CartoonreferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CartoonReferenceParserRULE_cartoonreference

	return p
}

func (s *CartoonreferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *CartoonreferenceContext) CARTOONREFERENCE() antlr.TerminalNode {
	return s.GetToken(CartoonReferenceParserCARTOONREFERENCE, 0)
}

func (s *CartoonreferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CartoonreferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CartoonreferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CartoonReferenceListener); ok {
		listenerT.EnterCartoonreference(s)
	}
}

func (s *CartoonreferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CartoonReferenceListener); ok {
		listenerT.ExitCartoonreference(s)
	}
}

func (p *CartoonReferenceParser) Cartoonreference() (localctx ICartoonreferenceContext) {
	this := p
	_ = this

	localctx = NewCartoonreferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CartoonReferenceParserRULE_cartoonreference)

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
		p.Match(CartoonReferenceParserCARTOONREFERENCE)
	}

	return localctx
}
