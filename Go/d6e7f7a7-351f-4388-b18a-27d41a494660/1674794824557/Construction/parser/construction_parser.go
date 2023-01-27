// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674794824557/Construction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Construction

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

type ConstructionParser struct {
	*antlr.BaseParser
}

var constructionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func constructionParserInit() {
	staticData := &constructionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONSTRUCTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "construction",
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

// ConstructionParserInit initializes any static state used to implement ConstructionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConstructionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConstructionParserInit() {
	staticData := &constructionParserStaticData
	staticData.once.Do(constructionParserInit)
}

// NewConstructionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConstructionParser(input antlr.TokenStream) *ConstructionParser {
	ConstructionParserInit()
	this := new(ConstructionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &constructionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Construction.g4"

	return this
}

// ConstructionParser tokens.
const (
	ConstructionParserEOF          = antlr.TokenEOF
	ConstructionParserCONSTRUCTION = 1
	ConstructionParserOWNKEY       = 2
	ConstructionParserSPLITKEY     = 3
	ConstructionParserWS           = 4
)

// ConstructionParser rules.
const (
	ConstructionParserRULE_expression   = 0
	ConstructionParserRULE_construction = 1
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
	p.RuleIndex = ConstructionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstructionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Construction() IConstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConstructionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstructionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstructionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConstructionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConstructionParserRULE_expression)

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
		p.Construction()
	}
	{
		p.SetState(5)
		p.Match(ConstructionParserEOF)
	}

	return localctx
}

// IConstructionContext is an interface to support dynamic dispatch.
type IConstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructionContext differentiates from other interfaces.
	IsConstructionContext()
}

type ConstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructionContext() *ConstructionContext {
	var p = new(ConstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConstructionParserRULE_construction
	return p
}

func (*ConstructionContext) IsConstructionContext() {}

func NewConstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructionContext {
	var p = new(ConstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConstructionParserRULE_construction

	return p
}

func (s *ConstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructionContext) CONSTRUCTION() antlr.TerminalNode {
	return s.GetToken(ConstructionParserCONSTRUCTION, 0)
}

func (s *ConstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstructionListener); ok {
		listenerT.EnterConstruction(s)
	}
}

func (s *ConstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConstructionListener); ok {
		listenerT.ExitConstruction(s)
	}
}

func (p *ConstructionParser) Construction() (localctx IConstructionContext) {
	this := p
	_ = this

	localctx = NewConstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConstructionParserRULE_construction)

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
		p.Match(ConstructionParserCONSTRUCTION)
	}

	return localctx
}
