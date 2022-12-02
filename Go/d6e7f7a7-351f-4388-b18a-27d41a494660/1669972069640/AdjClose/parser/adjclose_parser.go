// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972069640/AdjClose.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AdjClose

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

type AdjCloseParser struct {
	*antlr.BaseParser
}

var adjcloseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func adjcloseParserInit() {
	staticData := &adjcloseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ADJCLOSE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "adjclose",
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

// AdjCloseParserInit initializes any static state used to implement AdjCloseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAdjCloseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AdjCloseParserInit() {
	staticData := &adjcloseParserStaticData
	staticData.once.Do(adjcloseParserInit)
}

// NewAdjCloseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAdjCloseParser(input antlr.TokenStream) *AdjCloseParser {
	AdjCloseParserInit()
	this := new(AdjCloseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &adjcloseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AdjClose.g4"

	return this
}

// AdjCloseParser tokens.
const (
	AdjCloseParserEOF      = antlr.TokenEOF
	AdjCloseParserADJCLOSE = 1
	AdjCloseParserOWNKEY   = 2
	AdjCloseParserSPLITKEY = 3
	AdjCloseParserWS       = 4
)

// AdjCloseParser rules.
const (
	AdjCloseParserRULE_expression = 0
	AdjCloseParserRULE_adjclose   = 1
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
	p.RuleIndex = AdjCloseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdjCloseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Adjclose() IAdjcloseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdjcloseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdjcloseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdjCloseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdjCloseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdjCloseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AdjCloseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AdjCloseParserRULE_expression)

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
		p.Adjclose()
	}
	{
		p.SetState(5)
		p.Match(AdjCloseParserEOF)
	}

	return localctx
}

// IAdjcloseContext is an interface to support dynamic dispatch.
type IAdjcloseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdjcloseContext differentiates from other interfaces.
	IsAdjcloseContext()
}

type AdjcloseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdjcloseContext() *AdjcloseContext {
	var p = new(AdjcloseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdjCloseParserRULE_adjclose
	return p
}

func (*AdjcloseContext) IsAdjcloseContext() {}

func NewAdjcloseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdjcloseContext {
	var p = new(AdjcloseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdjCloseParserRULE_adjclose

	return p
}

func (s *AdjcloseContext) GetParser() antlr.Parser { return s.parser }

func (s *AdjcloseContext) ADJCLOSE() antlr.TerminalNode {
	return s.GetToken(AdjCloseParserADJCLOSE, 0)
}

func (s *AdjcloseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdjcloseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdjcloseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdjCloseListener); ok {
		listenerT.EnterAdjclose(s)
	}
}

func (s *AdjcloseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdjCloseListener); ok {
		listenerT.ExitAdjclose(s)
	}
}

func (p *AdjCloseParser) Adjclose() (localctx IAdjcloseContext) {
	this := p
	_ = this

	localctx = NewAdjcloseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AdjCloseParserRULE_adjclose)

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
		p.Match(AdjCloseParserADJCLOSE)
	}

	return localctx
}
