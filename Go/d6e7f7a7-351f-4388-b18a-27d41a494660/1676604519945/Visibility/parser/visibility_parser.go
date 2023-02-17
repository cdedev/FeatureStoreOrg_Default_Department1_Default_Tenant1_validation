// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676604519945/Visibility.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visibility

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

type VisibilityParser struct {
	*antlr.BaseParser
}

var visibilityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func visibilityParserInit() {
	staticData := &visibilityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VISIBILITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "visibility",
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

// VisibilityParserInit initializes any static state used to implement VisibilityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVisibilityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VisibilityParserInit() {
	staticData := &visibilityParserStaticData
	staticData.once.Do(visibilityParserInit)
}

// NewVisibilityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVisibilityParser(input antlr.TokenStream) *VisibilityParser {
	VisibilityParserInit()
	this := new(VisibilityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &visibilityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Visibility.g4"

	return this
}

// VisibilityParser tokens.
const (
	VisibilityParserEOF        = antlr.TokenEOF
	VisibilityParserVISIBILITY = 1
	VisibilityParserOWNKEY     = 2
	VisibilityParserSPLITKEY   = 3
	VisibilityParserWS         = 4
)

// VisibilityParser rules.
const (
	VisibilityParserRULE_expression = 0
	VisibilityParserRULE_visibility = 1
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
	p.RuleIndex = VisibilityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VisibilityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Visibility() IVisibilityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVisibilityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVisibilityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VisibilityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisibilityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisibilityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VisibilityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VisibilityParserRULE_expression)

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
		p.Visibility()
	}
	{
		p.SetState(5)
		p.Match(VisibilityParserEOF)
	}

	return localctx
}

// IVisibilityContext is an interface to support dynamic dispatch.
type IVisibilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVisibilityContext differentiates from other interfaces.
	IsVisibilityContext()
}

type VisibilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVisibilityContext() *VisibilityContext {
	var p = new(VisibilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VisibilityParserRULE_visibility
	return p
}

func (*VisibilityContext) IsVisibilityContext() {}

func NewVisibilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VisibilityContext {
	var p = new(VisibilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VisibilityParserRULE_visibility

	return p
}

func (s *VisibilityContext) GetParser() antlr.Parser { return s.parser }

func (s *VisibilityContext) VISIBILITY() antlr.TerminalNode {
	return s.GetToken(VisibilityParserVISIBILITY, 0)
}

func (s *VisibilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisibilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VisibilityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisibilityListener); ok {
		listenerT.EnterVisibility(s)
	}
}

func (s *VisibilityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisibilityListener); ok {
		listenerT.ExitVisibility(s)
	}
}

func (p *VisibilityParser) Visibility() (localctx IVisibilityContext) {
	this := p
	_ = this

	localctx = NewVisibilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VisibilityParserRULE_visibility)

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
		p.Match(VisibilityParserVISIBILITY)
	}

	return localctx
}
