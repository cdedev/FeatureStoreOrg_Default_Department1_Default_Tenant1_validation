// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376894048/Airbags.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airbags

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

type AirbagsParser struct {
	*antlr.BaseParser
}

var airbagsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func airbagsParserInit() {
	staticData := &airbagsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AIRBAGS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "airbags",
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

// AirbagsParserInit initializes any static state used to implement AirbagsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAirbagsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AirbagsParserInit() {
	staticData := &airbagsParserStaticData
	staticData.once.Do(airbagsParserInit)
}

// NewAirbagsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAirbagsParser(input antlr.TokenStream) *AirbagsParser {
	AirbagsParserInit()
	this := new(AirbagsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &airbagsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Airbags.g4"

	return this
}

// AirbagsParser tokens.
const (
	AirbagsParserEOF      = antlr.TokenEOF
	AirbagsParserAIRBAGS  = 1
	AirbagsParserOWNKEY   = 2
	AirbagsParserSPLITKEY = 3
	AirbagsParserWS       = 4
)

// AirbagsParser rules.
const (
	AirbagsParserRULE_expression = 0
	AirbagsParserRULE_airbags    = 1
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
	p.RuleIndex = AirbagsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AirbagsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Airbags() IAirbagsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAirbagsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAirbagsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AirbagsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirbagsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirbagsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AirbagsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AirbagsParserRULE_expression)

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
		p.Airbags()
	}
	{
		p.SetState(5)
		p.Match(AirbagsParserEOF)
	}

	return localctx
}

// IAirbagsContext is an interface to support dynamic dispatch.
type IAirbagsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAirbagsContext differentiates from other interfaces.
	IsAirbagsContext()
}

type AirbagsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAirbagsContext() *AirbagsContext {
	var p = new(AirbagsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AirbagsParserRULE_airbags
	return p
}

func (*AirbagsContext) IsAirbagsContext() {}

func NewAirbagsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AirbagsContext {
	var p = new(AirbagsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AirbagsParserRULE_airbags

	return p
}

func (s *AirbagsContext) GetParser() antlr.Parser { return s.parser }

func (s *AirbagsContext) AIRBAGS() antlr.TerminalNode {
	return s.GetToken(AirbagsParserAIRBAGS, 0)
}

func (s *AirbagsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AirbagsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AirbagsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirbagsListener); ok {
		listenerT.EnterAirbags(s)
	}
}

func (s *AirbagsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AirbagsListener); ok {
		listenerT.ExitAirbags(s)
	}
}

func (p *AirbagsParser) Airbags() (localctx IAirbagsContext) {
	this := p
	_ = this

	localctx = NewAirbagsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AirbagsParserRULE_airbags)

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
		p.Match(AirbagsParserAIRBAGS)
	}

	return localctx
}
