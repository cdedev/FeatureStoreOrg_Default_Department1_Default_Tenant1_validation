// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669650939823/Pages.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pages

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

type PagesParser struct {
	*antlr.BaseParser
}

var pagesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pagesParserInit() {
	staticData := &pagesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PAGES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pages",
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

// PagesParserInit initializes any static state used to implement PagesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPagesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PagesParserInit() {
	staticData := &pagesParserStaticData
	staticData.once.Do(pagesParserInit)
}

// NewPagesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPagesParser(input antlr.TokenStream) *PagesParser {
	PagesParserInit()
	this := new(PagesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pagesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Pages.g4"

	return this
}

// PagesParser tokens.
const (
	PagesParserEOF      = antlr.TokenEOF
	PagesParserPAGES    = 1
	PagesParserOWNKEY   = 2
	PagesParserSPLITKEY = 3
	PagesParserWS       = 4
)

// PagesParser rules.
const (
	PagesParserRULE_expression = 0
	PagesParserRULE_pages      = 1
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
	p.RuleIndex = PagesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PagesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pages() IPagesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPagesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPagesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PagesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PagesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PagesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PagesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PagesParserRULE_expression)

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
		p.Pages()
	}
	{
		p.SetState(5)
		p.Match(PagesParserEOF)
	}

	return localctx
}

// IPagesContext is an interface to support dynamic dispatch.
type IPagesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPagesContext differentiates from other interfaces.
	IsPagesContext()
}

type PagesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPagesContext() *PagesContext {
	var p = new(PagesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PagesParserRULE_pages
	return p
}

func (*PagesContext) IsPagesContext() {}

func NewPagesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PagesContext {
	var p = new(PagesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PagesParserRULE_pages

	return p
}

func (s *PagesContext) GetParser() antlr.Parser { return s.parser }

func (s *PagesContext) PAGES() antlr.TerminalNode {
	return s.GetToken(PagesParserPAGES, 0)
}

func (s *PagesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PagesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PagesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PagesListener); ok {
		listenerT.EnterPages(s)
	}
}

func (s *PagesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PagesListener); ok {
		listenerT.ExitPages(s)
	}
}

func (p *PagesParser) Pages() (localctx IPagesContext) {
	this := p
	_ = this

	localctx = NewPagesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PagesParserRULE_pages)

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
		p.Match(PagesParserPAGES)
	}

	return localctx
}
