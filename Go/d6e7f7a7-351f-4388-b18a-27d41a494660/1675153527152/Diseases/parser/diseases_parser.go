// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675153527152/Diseases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diseases

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

type DiseasesParser struct {
	*antlr.BaseParser
}

var diseasesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diseasesParserInit() {
	staticData := &diseasesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DISEASES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "diseases",
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

// DiseasesParserInit initializes any static state used to implement DiseasesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiseasesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiseasesParserInit() {
	staticData := &diseasesParserStaticData
	staticData.once.Do(diseasesParserInit)
}

// NewDiseasesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiseasesParser(input antlr.TokenStream) *DiseasesParser {
	DiseasesParserInit()
	this := new(DiseasesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diseasesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Diseases.g4"

	return this
}

// DiseasesParser tokens.
const (
	DiseasesParserEOF      = antlr.TokenEOF
	DiseasesParserDISEASES = 1
	DiseasesParserOWNKEY   = 2
	DiseasesParserSPLITKEY = 3
	DiseasesParserWS       = 4
)

// DiseasesParser rules.
const (
	DiseasesParserRULE_expression = 0
	DiseasesParserRULE_diseases   = 1
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
	p.RuleIndex = DiseasesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiseasesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Diseases() IDiseasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiseasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiseasesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiseasesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseasesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseasesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiseasesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiseasesParserRULE_expression)

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
		p.Diseases()
	}
	{
		p.SetState(5)
		p.Match(DiseasesParserEOF)
	}

	return localctx
}

// IDiseasesContext is an interface to support dynamic dispatch.
type IDiseasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiseasesContext differentiates from other interfaces.
	IsDiseasesContext()
}

type DiseasesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiseasesContext() *DiseasesContext {
	var p = new(DiseasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiseasesParserRULE_diseases
	return p
}

func (*DiseasesContext) IsDiseasesContext() {}

func NewDiseasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiseasesContext {
	var p = new(DiseasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiseasesParserRULE_diseases

	return p
}

func (s *DiseasesContext) GetParser() antlr.Parser { return s.parser }

func (s *DiseasesContext) DISEASES() antlr.TerminalNode {
	return s.GetToken(DiseasesParserDISEASES, 0)
}

func (s *DiseasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiseasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiseasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseasesListener); ok {
		listenerT.EnterDiseases(s)
	}
}

func (s *DiseasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseasesListener); ok {
		listenerT.ExitDiseases(s)
	}
}

func (p *DiseasesParser) Diseases() (localctx IDiseasesContext) {
	this := p
	_ = this

	localctx = NewDiseasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiseasesParserRULE_diseases)

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
		p.Match(DiseasesParserDISEASES)
	}

	return localctx
}
