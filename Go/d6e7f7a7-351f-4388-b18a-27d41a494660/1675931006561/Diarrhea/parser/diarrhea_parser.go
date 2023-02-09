// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931006561/Diarrhea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diarrhea

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

type DiarrheaParser struct {
	*antlr.BaseParser
}

var diarrheaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diarrheaParserInit() {
	staticData := &diarrheaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DIARRHEA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "diarrhea",
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

// DiarrheaParserInit initializes any static state used to implement DiarrheaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiarrheaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiarrheaParserInit() {
	staticData := &diarrheaParserStaticData
	staticData.once.Do(diarrheaParserInit)
}

// NewDiarrheaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiarrheaParser(input antlr.TokenStream) *DiarrheaParser {
	DiarrheaParserInit()
	this := new(DiarrheaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diarrheaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Diarrhea.g4"

	return this
}

// DiarrheaParser tokens.
const (
	DiarrheaParserEOF      = antlr.TokenEOF
	DiarrheaParserDIARRHEA = 1
	DiarrheaParserOWNKEY   = 2
	DiarrheaParserSPLITKEY = 3
	DiarrheaParserWS       = 4
)

// DiarrheaParser rules.
const (
	DiarrheaParserRULE_expression = 0
	DiarrheaParserRULE_diarrhea   = 1
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
	p.RuleIndex = DiarrheaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiarrheaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Diarrhea() IDiarrheaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiarrheaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiarrheaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiarrheaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiarrheaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiarrheaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiarrheaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiarrheaParserRULE_expression)

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
		p.Diarrhea()
	}
	{
		p.SetState(5)
		p.Match(DiarrheaParserEOF)
	}

	return localctx
}

// IDiarrheaContext is an interface to support dynamic dispatch.
type IDiarrheaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiarrheaContext differentiates from other interfaces.
	IsDiarrheaContext()
}

type DiarrheaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiarrheaContext() *DiarrheaContext {
	var p = new(DiarrheaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiarrheaParserRULE_diarrhea
	return p
}

func (*DiarrheaContext) IsDiarrheaContext() {}

func NewDiarrheaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiarrheaContext {
	var p = new(DiarrheaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiarrheaParserRULE_diarrhea

	return p
}

func (s *DiarrheaContext) GetParser() antlr.Parser { return s.parser }

func (s *DiarrheaContext) DIARRHEA() antlr.TerminalNode {
	return s.GetToken(DiarrheaParserDIARRHEA, 0)
}

func (s *DiarrheaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiarrheaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiarrheaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiarrheaListener); ok {
		listenerT.EnterDiarrhea(s)
	}
}

func (s *DiarrheaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiarrheaListener); ok {
		listenerT.ExitDiarrhea(s)
	}
}

func (p *DiarrheaParser) Diarrhea() (localctx IDiarrheaContext) {
	this := p
	_ = this

	localctx = NewDiarrheaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiarrheaParserRULE_diarrhea)

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
		p.Match(DiarrheaParserDIARRHEA)
	}

	return localctx
}
