// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116577334/NightHours.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NightHours

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

type NightHoursParser struct {
	*antlr.BaseParser
}

var nighthoursParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nighthoursParserInit() {
	staticData := &nighthoursParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NIGHTHOURS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "nighthours",
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

// NightHoursParserInit initializes any static state used to implement NightHoursParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNightHoursParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NightHoursParserInit() {
	staticData := &nighthoursParserStaticData
	staticData.once.Do(nighthoursParserInit)
}

// NewNightHoursParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNightHoursParser(input antlr.TokenStream) *NightHoursParser {
	NightHoursParserInit()
	this := new(NightHoursParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nighthoursParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NightHours.g4"

	return this
}

// NightHoursParser tokens.
const (
	NightHoursParserEOF        = antlr.TokenEOF
	NightHoursParserNIGHTHOURS = 1
	NightHoursParserOWNKEY     = 2
	NightHoursParserSPLITKEY   = 3
	NightHoursParserWS         = 4
)

// NightHoursParser rules.
const (
	NightHoursParserRULE_expression = 0
	NightHoursParserRULE_nighthours = 1
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
	p.RuleIndex = NightHoursParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NightHoursParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Nighthours() INighthoursContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INighthoursContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INighthoursContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NightHoursParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NightHoursListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NightHoursListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NightHoursParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NightHoursParserRULE_expression)

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
		p.Nighthours()
	}
	{
		p.SetState(5)
		p.Match(NightHoursParserEOF)
	}

	return localctx
}

// INighthoursContext is an interface to support dynamic dispatch.
type INighthoursContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNighthoursContext differentiates from other interfaces.
	IsNighthoursContext()
}

type NighthoursContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNighthoursContext() *NighthoursContext {
	var p = new(NighthoursContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NightHoursParserRULE_nighthours
	return p
}

func (*NighthoursContext) IsNighthoursContext() {}

func NewNighthoursContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NighthoursContext {
	var p = new(NighthoursContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NightHoursParserRULE_nighthours

	return p
}

func (s *NighthoursContext) GetParser() antlr.Parser { return s.parser }

func (s *NighthoursContext) NIGHTHOURS() antlr.TerminalNode {
	return s.GetToken(NightHoursParserNIGHTHOURS, 0)
}

func (s *NighthoursContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NighthoursContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NighthoursContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NightHoursListener); ok {
		listenerT.EnterNighthours(s)
	}
}

func (s *NighthoursContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NightHoursListener); ok {
		listenerT.ExitNighthours(s)
	}
}

func (p *NightHoursParser) Nighthours() (localctx INighthoursContext) {
	this := p
	_ = this

	localctx = NewNighthoursContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NightHoursParserRULE_nighthours)

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
		p.Match(NightHoursParserNIGHTHOURS)
	}

	return localctx
}
