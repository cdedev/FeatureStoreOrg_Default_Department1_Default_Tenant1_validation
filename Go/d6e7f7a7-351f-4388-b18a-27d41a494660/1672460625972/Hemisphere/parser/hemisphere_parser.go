// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672460625972/Hemisphere.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hemisphere

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

type HemisphereParser struct {
	*antlr.BaseParser
}

var hemisphereParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hemisphereParserInit() {
	staticData := &hemisphereParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HEMISPHERE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hemisphere",
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

// HemisphereParserInit initializes any static state used to implement HemisphereParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHemisphereParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HemisphereParserInit() {
	staticData := &hemisphereParserStaticData
	staticData.once.Do(hemisphereParserInit)
}

// NewHemisphereParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHemisphereParser(input antlr.TokenStream) *HemisphereParser {
	HemisphereParserInit()
	this := new(HemisphereParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hemisphereParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Hemisphere.g4"

	return this
}

// HemisphereParser tokens.
const (
	HemisphereParserEOF        = antlr.TokenEOF
	HemisphereParserHEMISPHERE = 1
	HemisphereParserOWNKEY     = 2
	HemisphereParserSPLITKEY   = 3
	HemisphereParserWS         = 4
)

// HemisphereParser rules.
const (
	HemisphereParserRULE_expression = 0
	HemisphereParserRULE_hemisphere = 1
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
	p.RuleIndex = HemisphereParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HemisphereParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hemisphere() IHemisphereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHemisphereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHemisphereContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HemisphereParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HemisphereListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HemisphereListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HemisphereParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HemisphereParserRULE_expression)

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
		p.Hemisphere()
	}
	{
		p.SetState(5)
		p.Match(HemisphereParserEOF)
	}

	return localctx
}

// IHemisphereContext is an interface to support dynamic dispatch.
type IHemisphereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHemisphereContext differentiates from other interfaces.
	IsHemisphereContext()
}

type HemisphereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHemisphereContext() *HemisphereContext {
	var p = new(HemisphereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HemisphereParserRULE_hemisphere
	return p
}

func (*HemisphereContext) IsHemisphereContext() {}

func NewHemisphereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HemisphereContext {
	var p = new(HemisphereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HemisphereParserRULE_hemisphere

	return p
}

func (s *HemisphereContext) GetParser() antlr.Parser { return s.parser }

func (s *HemisphereContext) HEMISPHERE() antlr.TerminalNode {
	return s.GetToken(HemisphereParserHEMISPHERE, 0)
}

func (s *HemisphereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HemisphereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HemisphereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HemisphereListener); ok {
		listenerT.EnterHemisphere(s)
	}
}

func (s *HemisphereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HemisphereListener); ok {
		listenerT.ExitHemisphere(s)
	}
}

func (p *HemisphereParser) Hemisphere() (localctx IHemisphereContext) {
	this := p
	_ = this

	localctx = NewHemisphereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HemisphereParserRULE_hemisphere)

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
		p.Match(HemisphereParserHEMISPHERE)
	}

	return localctx
}
