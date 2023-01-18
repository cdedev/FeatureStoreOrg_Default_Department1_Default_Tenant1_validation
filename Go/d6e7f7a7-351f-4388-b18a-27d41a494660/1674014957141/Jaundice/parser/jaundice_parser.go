// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674014957141/Jaundice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jaundice

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

type JaundiceParser struct {
	*antlr.BaseParser
}

var jaundiceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jaundiceParserInit() {
	staticData := &jaundiceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JAUNDICE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "jaundice",
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

// JaundiceParserInit initializes any static state used to implement JaundiceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJaundiceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JaundiceParserInit() {
	staticData := &jaundiceParserStaticData
	staticData.once.Do(jaundiceParserInit)
}

// NewJaundiceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJaundiceParser(input antlr.TokenStream) *JaundiceParser {
	JaundiceParserInit()
	this := new(JaundiceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jaundiceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Jaundice.g4"

	return this
}

// JaundiceParser tokens.
const (
	JaundiceParserEOF      = antlr.TokenEOF
	JaundiceParserJAUNDICE = 1
	JaundiceParserOWNKEY   = 2
	JaundiceParserSPLITKEY = 3
	JaundiceParserWS       = 4
)

// JaundiceParser rules.
const (
	JaundiceParserRULE_expression = 0
	JaundiceParserRULE_jaundice   = 1
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
	p.RuleIndex = JaundiceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaundiceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Jaundice() IJaundiceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJaundiceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJaundiceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JaundiceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaundiceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaundiceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JaundiceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JaundiceParserRULE_expression)

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
		p.Jaundice()
	}
	{
		p.SetState(5)
		p.Match(JaundiceParserEOF)
	}

	return localctx
}

// IJaundiceContext is an interface to support dynamic dispatch.
type IJaundiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJaundiceContext differentiates from other interfaces.
	IsJaundiceContext()
}

type JaundiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJaundiceContext() *JaundiceContext {
	var p = new(JaundiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JaundiceParserRULE_jaundice
	return p
}

func (*JaundiceContext) IsJaundiceContext() {}

func NewJaundiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JaundiceContext {
	var p = new(JaundiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JaundiceParserRULE_jaundice

	return p
}

func (s *JaundiceContext) GetParser() antlr.Parser { return s.parser }

func (s *JaundiceContext) JAUNDICE() antlr.TerminalNode {
	return s.GetToken(JaundiceParserJAUNDICE, 0)
}

func (s *JaundiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JaundiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JaundiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaundiceListener); ok {
		listenerT.EnterJaundice(s)
	}
}

func (s *JaundiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JaundiceListener); ok {
		listenerT.ExitJaundice(s)
	}
}

func (p *JaundiceParser) Jaundice() (localctx IJaundiceContext) {
	this := p
	_ = this

	localctx = NewJaundiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JaundiceParserRULE_jaundice)

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
		p.Match(JaundiceParserJAUNDICE)
	}

	return localctx
}
