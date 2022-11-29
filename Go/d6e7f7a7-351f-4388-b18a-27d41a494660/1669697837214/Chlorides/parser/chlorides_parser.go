// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697837214/Chlorides.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chlorides

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

type ChloridesParser struct {
	*antlr.BaseParser
}

var chloridesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func chloridesParserInit() {
	staticData := &chloridesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHLORIDES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "chlorides",
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

// ChloridesParserInit initializes any static state used to implement ChloridesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewChloridesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChloridesParserInit() {
	staticData := &chloridesParserStaticData
	staticData.once.Do(chloridesParserInit)
}

// NewChloridesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewChloridesParser(input antlr.TokenStream) *ChloridesParser {
	ChloridesParserInit()
	this := new(ChloridesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &chloridesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Chlorides.g4"

	return this
}

// ChloridesParser tokens.
const (
	ChloridesParserEOF       = antlr.TokenEOF
	ChloridesParserCHLORIDES = 1
	ChloridesParserOWNKEY    = 2
	ChloridesParserSPLITKEY  = 3
	ChloridesParserWS        = 4
)

// ChloridesParser rules.
const (
	ChloridesParserRULE_expression = 0
	ChloridesParserRULE_chlorides  = 1
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
	p.RuleIndex = ChloridesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChloridesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Chlorides() IChloridesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChloridesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChloridesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChloridesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChloridesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChloridesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ChloridesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChloridesParserRULE_expression)

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
		p.Chlorides()
	}
	{
		p.SetState(5)
		p.Match(ChloridesParserEOF)
	}

	return localctx
}

// IChloridesContext is an interface to support dynamic dispatch.
type IChloridesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChloridesContext differentiates from other interfaces.
	IsChloridesContext()
}

type ChloridesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChloridesContext() *ChloridesContext {
	var p = new(ChloridesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChloridesParserRULE_chlorides
	return p
}

func (*ChloridesContext) IsChloridesContext() {}

func NewChloridesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChloridesContext {
	var p = new(ChloridesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChloridesParserRULE_chlorides

	return p
}

func (s *ChloridesContext) GetParser() antlr.Parser { return s.parser }

func (s *ChloridesContext) CHLORIDES() antlr.TerminalNode {
	return s.GetToken(ChloridesParserCHLORIDES, 0)
}

func (s *ChloridesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChloridesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChloridesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChloridesListener); ok {
		listenerT.EnterChlorides(s)
	}
}

func (s *ChloridesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChloridesListener); ok {
		listenerT.ExitChlorides(s)
	}
}

func (p *ChloridesParser) Chlorides() (localctx IChloridesContext) {
	this := p
	_ = this

	localctx = NewChloridesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChloridesParserRULE_chlorides)

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
		p.Match(ChloridesParserCHLORIDES)
	}

	return localctx
}
