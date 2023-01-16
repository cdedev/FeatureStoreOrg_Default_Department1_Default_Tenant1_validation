// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868669341/Clarity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clarity

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

type ClarityParser struct {
	*antlr.BaseParser
}

var clarityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clarityParserInit() {
	staticData := &clarityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLARITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "clarity",
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

// ClarityParserInit initializes any static state used to implement ClarityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClarityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClarityParserInit() {
	staticData := &clarityParserStaticData
	staticData.once.Do(clarityParserInit)
}

// NewClarityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClarityParser(input antlr.TokenStream) *ClarityParser {
	ClarityParserInit()
	this := new(ClarityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clarityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Clarity.g4"

	return this
}

// ClarityParser tokens.
const (
	ClarityParserEOF      = antlr.TokenEOF
	ClarityParserCLARITY  = 1
	ClarityParserOWNKEY   = 2
	ClarityParserSPLITKEY = 3
	ClarityParserWS       = 4
)

// ClarityParser rules.
const (
	ClarityParserRULE_expression = 0
	ClarityParserRULE_clarity    = 1
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
	p.RuleIndex = ClarityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Clarity() IClarityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClarityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClarityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClarityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ClarityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClarityParserRULE_expression)

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
		p.Clarity()
	}
	{
		p.SetState(5)
		p.Match(ClarityParserEOF)
	}

	return localctx
}

// IClarityContext is an interface to support dynamic dispatch.
type IClarityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClarityContext differentiates from other interfaces.
	IsClarityContext()
}

type ClarityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClarityContext() *ClarityContext {
	var p = new(ClarityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClarityParserRULE_clarity
	return p
}

func (*ClarityContext) IsClarityContext() {}

func NewClarityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClarityContext {
	var p = new(ClarityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClarityParserRULE_clarity

	return p
}

func (s *ClarityContext) GetParser() antlr.Parser { return s.parser }

func (s *ClarityContext) CLARITY() antlr.TerminalNode {
	return s.GetToken(ClarityParserCLARITY, 0)
}

func (s *ClarityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClarityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClarityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.EnterClarity(s)
	}
}

func (s *ClarityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClarityListener); ok {
		listenerT.ExitClarity(s)
	}
}

func (p *ClarityParser) Clarity() (localctx IClarityContext) {
	this := p
	_ = this

	localctx = NewClarityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClarityParserRULE_clarity)

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
		p.Match(ClarityParserCLARITY)
	}

	return localctx
}
