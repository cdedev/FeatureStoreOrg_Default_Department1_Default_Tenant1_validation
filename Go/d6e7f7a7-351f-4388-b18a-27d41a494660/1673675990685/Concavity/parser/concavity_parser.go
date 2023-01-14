// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673675990685/Concavity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Concavity

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

type ConcavityParser struct {
	*antlr.BaseParser
}

var concavityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func concavityParserInit() {
	staticData := &concavityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONCAVITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "concavity",
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

// ConcavityParserInit initializes any static state used to implement ConcavityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConcavityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConcavityParserInit() {
	staticData := &concavityParserStaticData
	staticData.once.Do(concavityParserInit)
}

// NewConcavityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConcavityParser(input antlr.TokenStream) *ConcavityParser {
	ConcavityParserInit()
	this := new(ConcavityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &concavityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Concavity.g4"

	return this
}

// ConcavityParser tokens.
const (
	ConcavityParserEOF       = antlr.TokenEOF
	ConcavityParserCONCAVITY = 1
	ConcavityParserOWNKEY    = 2
	ConcavityParserSPLITKEY  = 3
	ConcavityParserWS        = 4
)

// ConcavityParser rules.
const (
	ConcavityParserRULE_expression = 0
	ConcavityParserRULE_concavity  = 1
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
	p.RuleIndex = ConcavityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcavityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Concavity() IConcavityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConcavityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConcavityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConcavityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcavityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcavityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConcavityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConcavityParserRULE_expression)

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
		p.Concavity()
	}
	{
		p.SetState(5)
		p.Match(ConcavityParserEOF)
	}

	return localctx
}

// IConcavityContext is an interface to support dynamic dispatch.
type IConcavityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcavityContext differentiates from other interfaces.
	IsConcavityContext()
}

type ConcavityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcavityContext() *ConcavityContext {
	var p = new(ConcavityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcavityParserRULE_concavity
	return p
}

func (*ConcavityContext) IsConcavityContext() {}

func NewConcavityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcavityContext {
	var p = new(ConcavityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcavityParserRULE_concavity

	return p
}

func (s *ConcavityContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcavityContext) CONCAVITY() antlr.TerminalNode {
	return s.GetToken(ConcavityParserCONCAVITY, 0)
}

func (s *ConcavityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcavityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcavityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcavityListener); ok {
		listenerT.EnterConcavity(s)
	}
}

func (s *ConcavityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcavityListener); ok {
		listenerT.ExitConcavity(s)
	}
}

func (p *ConcavityParser) Concavity() (localctx IConcavityContext) {
	this := p
	_ = this

	localctx = NewConcavityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConcavityParserRULE_concavity)

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
		p.Match(ConcavityParserCONCAVITY)
	}

	return localctx
}
