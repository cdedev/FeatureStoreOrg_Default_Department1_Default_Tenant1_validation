// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604448956/Chromium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chromium

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

type ChromiumParser struct {
	*antlr.BaseParser
}

var chromiumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func chromiumParserInit() {
	staticData := &chromiumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHROMIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "chromium",
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

// ChromiumParserInit initializes any static state used to implement ChromiumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewChromiumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChromiumParserInit() {
	staticData := &chromiumParserStaticData
	staticData.once.Do(chromiumParserInit)
}

// NewChromiumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewChromiumParser(input antlr.TokenStream) *ChromiumParser {
	ChromiumParserInit()
	this := new(ChromiumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &chromiumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Chromium.g4"

	return this
}

// ChromiumParser tokens.
const (
	ChromiumParserEOF      = antlr.TokenEOF
	ChromiumParserCHROMIUM = 1
	ChromiumParserOWNKEY   = 2
	ChromiumParserSPLITKEY = 3
	ChromiumParserWS       = 4
)

// ChromiumParser rules.
const (
	ChromiumParserRULE_expression = 0
	ChromiumParserRULE_chromium   = 1
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
	p.RuleIndex = ChromiumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChromiumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Chromium() IChromiumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChromiumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChromiumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChromiumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChromiumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChromiumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ChromiumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChromiumParserRULE_expression)

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
		p.Chromium()
	}
	{
		p.SetState(5)
		p.Match(ChromiumParserEOF)
	}

	return localctx
}

// IChromiumContext is an interface to support dynamic dispatch.
type IChromiumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChromiumContext differentiates from other interfaces.
	IsChromiumContext()
}

type ChromiumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChromiumContext() *ChromiumContext {
	var p = new(ChromiumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChromiumParserRULE_chromium
	return p
}

func (*ChromiumContext) IsChromiumContext() {}

func NewChromiumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChromiumContext {
	var p = new(ChromiumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChromiumParserRULE_chromium

	return p
}

func (s *ChromiumContext) GetParser() antlr.Parser { return s.parser }

func (s *ChromiumContext) CHROMIUM() antlr.TerminalNode {
	return s.GetToken(ChromiumParserCHROMIUM, 0)
}

func (s *ChromiumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChromiumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChromiumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChromiumListener); ok {
		listenerT.EnterChromium(s)
	}
}

func (s *ChromiumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChromiumListener); ok {
		listenerT.ExitChromium(s)
	}
}

func (p *ChromiumParser) Chromium() (localctx IChromiumContext) {
	this := p
	_ = this

	localctx = NewChromiumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChromiumParserRULE_chromium)

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
		p.Match(ChromiumParserCHROMIUM)
	}

	return localctx
}
