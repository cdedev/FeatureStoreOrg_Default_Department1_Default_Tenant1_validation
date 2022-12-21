// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604495022/Cobalt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cobalt

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

type CobaltParser struct {
	*antlr.BaseParser
}

var cobaltParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cobaltParserInit() {
	staticData := &cobaltParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COBALT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cobalt",
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

// CobaltParserInit initializes any static state used to implement CobaltParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCobaltParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CobaltParserInit() {
	staticData := &cobaltParserStaticData
	staticData.once.Do(cobaltParserInit)
}

// NewCobaltParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCobaltParser(input antlr.TokenStream) *CobaltParser {
	CobaltParserInit()
	this := new(CobaltParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cobaltParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cobalt.g4"

	return this
}

// CobaltParser tokens.
const (
	CobaltParserEOF      = antlr.TokenEOF
	CobaltParserCOBALT   = 1
	CobaltParserOWNKEY   = 2
	CobaltParserSPLITKEY = 3
	CobaltParserWS       = 4
)

// CobaltParser rules.
const (
	CobaltParserRULE_expression = 0
	CobaltParserRULE_cobalt     = 1
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
	p.RuleIndex = CobaltParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CobaltParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cobalt() ICobaltContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICobaltContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICobaltContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CobaltParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CobaltListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CobaltListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CobaltParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CobaltParserRULE_expression)

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
		p.Cobalt()
	}
	{
		p.SetState(5)
		p.Match(CobaltParserEOF)
	}

	return localctx
}

// ICobaltContext is an interface to support dynamic dispatch.
type ICobaltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCobaltContext differentiates from other interfaces.
	IsCobaltContext()
}

type CobaltContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCobaltContext() *CobaltContext {
	var p = new(CobaltContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CobaltParserRULE_cobalt
	return p
}

func (*CobaltContext) IsCobaltContext() {}

func NewCobaltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CobaltContext {
	var p = new(CobaltContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CobaltParserRULE_cobalt

	return p
}

func (s *CobaltContext) GetParser() antlr.Parser { return s.parser }

func (s *CobaltContext) COBALT() antlr.TerminalNode {
	return s.GetToken(CobaltParserCOBALT, 0)
}

func (s *CobaltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CobaltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CobaltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CobaltListener); ok {
		listenerT.EnterCobalt(s)
	}
}

func (s *CobaltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CobaltListener); ok {
		listenerT.ExitCobalt(s)
	}
}

func (p *CobaltParser) Cobalt() (localctx ICobaltContext) {
	this := p
	_ = this

	localctx = NewCobaltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CobaltParserRULE_cobalt)

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
		p.Match(CobaltParserCOBALT)
	}

	return localctx
}
