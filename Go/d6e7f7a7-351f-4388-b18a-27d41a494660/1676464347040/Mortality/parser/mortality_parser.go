// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676464347040/Mortality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mortality

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

type MortalityParser struct {
	*antlr.BaseParser
}

var mortalityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mortalityParserInit() {
	staticData := &mortalityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MORTALITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mortality",
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

// MortalityParserInit initializes any static state used to implement MortalityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMortalityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MortalityParserInit() {
	staticData := &mortalityParserStaticData
	staticData.once.Do(mortalityParserInit)
}

// NewMortalityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMortalityParser(input antlr.TokenStream) *MortalityParser {
	MortalityParserInit()
	this := new(MortalityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mortalityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mortality.g4"

	return this
}

// MortalityParser tokens.
const (
	MortalityParserEOF       = antlr.TokenEOF
	MortalityParserMORTALITY = 1
	MortalityParserOWNKEY    = 2
	MortalityParserSPLITKEY  = 3
	MortalityParserWS        = 4
)

// MortalityParser rules.
const (
	MortalityParserRULE_expression = 0
	MortalityParserRULE_mortality  = 1
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
	p.RuleIndex = MortalityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MortalityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mortality() IMortalityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMortalityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMortalityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MortalityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortalityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortalityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MortalityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MortalityParserRULE_expression)

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
		p.Mortality()
	}
	{
		p.SetState(5)
		p.Match(MortalityParserEOF)
	}

	return localctx
}

// IMortalityContext is an interface to support dynamic dispatch.
type IMortalityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMortalityContext differentiates from other interfaces.
	IsMortalityContext()
}

type MortalityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMortalityContext() *MortalityContext {
	var p = new(MortalityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MortalityParserRULE_mortality
	return p
}

func (*MortalityContext) IsMortalityContext() {}

func NewMortalityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MortalityContext {
	var p = new(MortalityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MortalityParserRULE_mortality

	return p
}

func (s *MortalityContext) GetParser() antlr.Parser { return s.parser }

func (s *MortalityContext) MORTALITY() antlr.TerminalNode {
	return s.GetToken(MortalityParserMORTALITY, 0)
}

func (s *MortalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MortalityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MortalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortalityListener); ok {
		listenerT.EnterMortality(s)
	}
}

func (s *MortalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortalityListener); ok {
		listenerT.ExitMortality(s)
	}
}

func (p *MortalityParser) Mortality() (localctx IMortalityContext) {
	this := p
	_ = this

	localctx = NewMortalityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MortalityParserRULE_mortality)

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
		p.Match(MortalityParserMORTALITY)
	}

	return localctx
}
