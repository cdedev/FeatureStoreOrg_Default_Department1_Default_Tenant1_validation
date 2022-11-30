// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780323186/Oil.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Oil

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

type OilParser struct {
	*antlr.BaseParser
}

var oilParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func oilParserInit() {
	staticData := &oilParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OIL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "oil",
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

// OilParserInit initializes any static state used to implement OilParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOilParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OilParserInit() {
	staticData := &oilParserStaticData
	staticData.once.Do(oilParserInit)
}

// NewOilParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOilParser(input antlr.TokenStream) *OilParser {
	OilParserInit()
	this := new(OilParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &oilParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Oil.g4"

	return this
}

// OilParser tokens.
const (
	OilParserEOF      = antlr.TokenEOF
	OilParserOIL      = 1
	OilParserOWNKEY   = 2
	OilParserSPLITKEY = 3
	OilParserWS       = 4
)

// OilParser rules.
const (
	OilParserRULE_expression = 0
	OilParserRULE_oil        = 1
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
	p.RuleIndex = OilParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OilParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Oil() IOilContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOilContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOilContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OilParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OilParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OilParserRULE_expression)

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
		p.Oil()
	}
	{
		p.SetState(5)
		p.Match(OilParserEOF)
	}

	return localctx
}

// IOilContext is an interface to support dynamic dispatch.
type IOilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOilContext differentiates from other interfaces.
	IsOilContext()
}

type OilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOilContext() *OilContext {
	var p = new(OilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OilParserRULE_oil
	return p
}

func (*OilContext) IsOilContext() {}

func NewOilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OilContext {
	var p = new(OilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OilParserRULE_oil

	return p
}

func (s *OilContext) GetParser() antlr.Parser { return s.parser }

func (s *OilContext) OIL() antlr.TerminalNode {
	return s.GetToken(OilParserOIL, 0)
}

func (s *OilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilListener); ok {
		listenerT.EnterOil(s)
	}
}

func (s *OilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilListener); ok {
		listenerT.ExitOil(s)
	}
}

func (p *OilParser) Oil() (localctx IOilContext) {
	this := p
	_ = this

	localctx = NewOilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OilParserRULE_oil)

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
		p.Match(OilParserOIL)
	}

	return localctx
}
