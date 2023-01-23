// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674454745948/Category.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Category

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

type CategoryParser struct {
	*antlr.BaseParser
}

var categoryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func categoryParserInit() {
	staticData := &categoryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CATEGORY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "category",
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

// CategoryParserInit initializes any static state used to implement CategoryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCategoryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CategoryParserInit() {
	staticData := &categoryParserStaticData
	staticData.once.Do(categoryParserInit)
}

// NewCategoryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCategoryParser(input antlr.TokenStream) *CategoryParser {
	CategoryParserInit()
	this := new(CategoryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &categoryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Category.g4"

	return this
}

// CategoryParser tokens.
const (
	CategoryParserEOF      = antlr.TokenEOF
	CategoryParserCATEGORY = 1
	CategoryParserOWNKEY   = 2
	CategoryParserSPLITKEY = 3
	CategoryParserWS       = 4
)

// CategoryParser rules.
const (
	CategoryParserRULE_expression = 0
	CategoryParserRULE_category   = 1
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
	p.RuleIndex = CategoryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CategoryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Category() ICategoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICategoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICategoryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CategoryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CategoryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CategoryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CategoryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CategoryParserRULE_expression)

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
		p.Category()
	}
	{
		p.SetState(5)
		p.Match(CategoryParserEOF)
	}

	return localctx
}

// ICategoryContext is an interface to support dynamic dispatch.
type ICategoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCategoryContext differentiates from other interfaces.
	IsCategoryContext()
}

type CategoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCategoryContext() *CategoryContext {
	var p = new(CategoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CategoryParserRULE_category
	return p
}

func (*CategoryContext) IsCategoryContext() {}

func NewCategoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CategoryContext {
	var p = new(CategoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CategoryParserRULE_category

	return p
}

func (s *CategoryContext) GetParser() antlr.Parser { return s.parser }

func (s *CategoryContext) CATEGORY() antlr.TerminalNode {
	return s.GetToken(CategoryParserCATEGORY, 0)
}

func (s *CategoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CategoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CategoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CategoryListener); ok {
		listenerT.EnterCategory(s)
	}
}

func (s *CategoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CategoryListener); ok {
		listenerT.ExitCategory(s)
	}
}

func (p *CategoryParser) Category() (localctx ICategoryContext) {
	this := p
	_ = this

	localctx = NewCategoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CategoryParserRULE_category)

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
		p.Match(CategoryParserCATEGORY)
	}

	return localctx
}
