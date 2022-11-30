// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795091007/Carbon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbon

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

type CarbonParser struct {
	*antlr.BaseParser
}

var carbonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func carbonParserInit() {
	staticData := &carbonParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CARBON", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "carbon",
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

// CarbonParserInit initializes any static state used to implement CarbonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCarbonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CarbonParserInit() {
	staticData := &carbonParserStaticData
	staticData.once.Do(carbonParserInit)
}

// NewCarbonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCarbonParser(input antlr.TokenStream) *CarbonParser {
	CarbonParserInit()
	this := new(CarbonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &carbonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Carbon.g4"

	return this
}

// CarbonParser tokens.
const (
	CarbonParserEOF      = antlr.TokenEOF
	CarbonParserCARBON   = 1
	CarbonParserOWNKEY   = 2
	CarbonParserSPLITKEY = 3
	CarbonParserWS       = 4
)

// CarbonParser rules.
const (
	CarbonParserRULE_expression = 0
	CarbonParserRULE_carbon     = 1
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
	p.RuleIndex = CarbonParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CarbonParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Carbon() ICarbonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICarbonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICarbonContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CarbonParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbonListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbonListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CarbonParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CarbonParserRULE_expression)

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
		p.Carbon()
	}
	{
		p.SetState(5)
		p.Match(CarbonParserEOF)
	}

	return localctx
}

// ICarbonContext is an interface to support dynamic dispatch.
type ICarbonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCarbonContext differentiates from other interfaces.
	IsCarbonContext()
}

type CarbonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCarbonContext() *CarbonContext {
	var p = new(CarbonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CarbonParserRULE_carbon
	return p
}

func (*CarbonContext) IsCarbonContext() {}

func NewCarbonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CarbonContext {
	var p = new(CarbonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CarbonParserRULE_carbon

	return p
}

func (s *CarbonContext) GetParser() antlr.Parser { return s.parser }

func (s *CarbonContext) CARBON() antlr.TerminalNode {
	return s.GetToken(CarbonParserCARBON, 0)
}

func (s *CarbonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CarbonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CarbonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbonListener); ok {
		listenerT.EnterCarbon(s)
	}
}

func (s *CarbonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbonListener); ok {
		listenerT.ExitCarbon(s)
	}
}

func (p *CarbonParser) Carbon() (localctx ICarbonContext) {
	this := p
	_ = this

	localctx = NewCarbonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CarbonParserRULE_carbon)

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
		p.Match(CarbonParserCARBON)
	}

	return localctx
}
