// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672981954579/InventoryCategory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InventoryCategory

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

type InventoryCategoryParser struct {
	*antlr.BaseParser
}

var inventorycategoryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func inventorycategoryParserInit() {
	staticData := &inventorycategoryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INVENTORYCATEGORY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "inventorycategory",
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

// InventoryCategoryParserInit initializes any static state used to implement InventoryCategoryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInventoryCategoryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InventoryCategoryParserInit() {
	staticData := &inventorycategoryParserStaticData
	staticData.once.Do(inventorycategoryParserInit)
}

// NewInventoryCategoryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInventoryCategoryParser(input antlr.TokenStream) *InventoryCategoryParser {
	InventoryCategoryParserInit()
	this := new(InventoryCategoryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &inventorycategoryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "InventoryCategory.g4"

	return this
}

// InventoryCategoryParser tokens.
const (
	InventoryCategoryParserEOF               = antlr.TokenEOF
	InventoryCategoryParserINVENTORYCATEGORY = 1
	InventoryCategoryParserOWNKEY            = 2
	InventoryCategoryParserSPLITKEY          = 3
	InventoryCategoryParserWS                = 4
)

// InventoryCategoryParser rules.
const (
	InventoryCategoryParserRULE_expression        = 0
	InventoryCategoryParserRULE_inventorycategory = 1
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
	p.RuleIndex = InventoryCategoryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InventoryCategoryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Inventorycategory() IInventorycategoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInventorycategoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInventorycategoryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(InventoryCategoryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InventoryCategoryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InventoryCategoryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *InventoryCategoryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InventoryCategoryParserRULE_expression)

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
		p.Inventorycategory()
	}
	{
		p.SetState(5)
		p.Match(InventoryCategoryParserEOF)
	}

	return localctx
}

// IInventorycategoryContext is an interface to support dynamic dispatch.
type IInventorycategoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInventorycategoryContext differentiates from other interfaces.
	IsInventorycategoryContext()
}

type InventorycategoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInventorycategoryContext() *InventorycategoryContext {
	var p = new(InventorycategoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InventoryCategoryParserRULE_inventorycategory
	return p
}

func (*InventorycategoryContext) IsInventorycategoryContext() {}

func NewInventorycategoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InventorycategoryContext {
	var p = new(InventorycategoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InventoryCategoryParserRULE_inventorycategory

	return p
}

func (s *InventorycategoryContext) GetParser() antlr.Parser { return s.parser }

func (s *InventorycategoryContext) INVENTORYCATEGORY() antlr.TerminalNode {
	return s.GetToken(InventoryCategoryParserINVENTORYCATEGORY, 0)
}

func (s *InventorycategoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InventorycategoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InventorycategoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InventoryCategoryListener); ok {
		listenerT.EnterInventorycategory(s)
	}
}

func (s *InventorycategoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InventoryCategoryListener); ok {
		listenerT.ExitInventorycategory(s)
	}
}

func (p *InventoryCategoryParser) Inventorycategory() (localctx IInventorycategoryContext) {
	this := p
	_ = this

	localctx = NewInventorycategoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InventoryCategoryParserRULE_inventorycategory)

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
		p.Match(InventoryCategoryParserINVENTORYCATEGORY)
	}

	return localctx
}
