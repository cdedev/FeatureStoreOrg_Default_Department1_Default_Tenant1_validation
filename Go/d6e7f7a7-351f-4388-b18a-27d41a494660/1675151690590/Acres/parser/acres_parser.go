// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675151690590/Acres.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acres

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

type AcresParser struct {
	*antlr.BaseParser
}

var acresParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func acresParserInit() {
	staticData := &acresParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ACRES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "acres",
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

// AcresParserInit initializes any static state used to implement AcresParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAcresParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AcresParserInit() {
	staticData := &acresParserStaticData
	staticData.once.Do(acresParserInit)
}

// NewAcresParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAcresParser(input antlr.TokenStream) *AcresParser {
	AcresParserInit()
	this := new(AcresParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &acresParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Acres.g4"

	return this
}

// AcresParser tokens.
const (
	AcresParserEOF      = antlr.TokenEOF
	AcresParserACRES    = 1
	AcresParserOWNKEY   = 2
	AcresParserSPLITKEY = 3
	AcresParserWS       = 4
)

// AcresParser rules.
const (
	AcresParserRULE_expression = 0
	AcresParserRULE_acres      = 1
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
	p.RuleIndex = AcresParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AcresParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Acres() IAcresContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAcresContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAcresContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AcresParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcresListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcresListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AcresParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AcresParserRULE_expression)

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
		p.Acres()
	}
	{
		p.SetState(5)
		p.Match(AcresParserEOF)
	}

	return localctx
}

// IAcresContext is an interface to support dynamic dispatch.
type IAcresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAcresContext differentiates from other interfaces.
	IsAcresContext()
}

type AcresContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAcresContext() *AcresContext {
	var p = new(AcresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AcresParserRULE_acres
	return p
}

func (*AcresContext) IsAcresContext() {}

func NewAcresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AcresContext {
	var p = new(AcresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AcresParserRULE_acres

	return p
}

func (s *AcresContext) GetParser() antlr.Parser { return s.parser }

func (s *AcresContext) ACRES() antlr.TerminalNode {
	return s.GetToken(AcresParserACRES, 0)
}

func (s *AcresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AcresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcresListener); ok {
		listenerT.EnterAcres(s)
	}
}

func (s *AcresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AcresListener); ok {
		listenerT.ExitAcres(s)
	}
}

func (p *AcresParser) Acres() (localctx IAcresContext) {
	this := p
	_ = this

	localctx = NewAcresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AcresParserRULE_acres)

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
		p.Match(AcresParserACRES)
	}

	return localctx
}
