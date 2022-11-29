// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699123294/Eccentricity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eccentricity

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

type EccentricityParser struct {
	*antlr.BaseParser
}

var eccentricityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eccentricityParserInit() {
	staticData := &eccentricityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ECCENTRICITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "eccentricity",
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

// EccentricityParserInit initializes any static state used to implement EccentricityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEccentricityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EccentricityParserInit() {
	staticData := &eccentricityParserStaticData
	staticData.once.Do(eccentricityParserInit)
}

// NewEccentricityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEccentricityParser(input antlr.TokenStream) *EccentricityParser {
	EccentricityParserInit()
	this := new(EccentricityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &eccentricityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Eccentricity.g4"

	return this
}

// EccentricityParser tokens.
const (
	EccentricityParserEOF          = antlr.TokenEOF
	EccentricityParserECCENTRICITY = 1
	EccentricityParserOWNKEY       = 2
	EccentricityParserSPLITKEY     = 3
	EccentricityParserWS           = 4
)

// EccentricityParser rules.
const (
	EccentricityParserRULE_expression   = 0
	EccentricityParserRULE_eccentricity = 1
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
	p.RuleIndex = EccentricityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EccentricityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Eccentricity() IEccentricityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEccentricityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEccentricityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EccentricityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EccentricityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EccentricityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EccentricityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EccentricityParserRULE_expression)

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
		p.Eccentricity()
	}
	{
		p.SetState(5)
		p.Match(EccentricityParserEOF)
	}

	return localctx
}

// IEccentricityContext is an interface to support dynamic dispatch.
type IEccentricityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEccentricityContext differentiates from other interfaces.
	IsEccentricityContext()
}

type EccentricityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEccentricityContext() *EccentricityContext {
	var p = new(EccentricityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EccentricityParserRULE_eccentricity
	return p
}

func (*EccentricityContext) IsEccentricityContext() {}

func NewEccentricityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EccentricityContext {
	var p = new(EccentricityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EccentricityParserRULE_eccentricity

	return p
}

func (s *EccentricityContext) GetParser() antlr.Parser { return s.parser }

func (s *EccentricityContext) ECCENTRICITY() antlr.TerminalNode {
	return s.GetToken(EccentricityParserECCENTRICITY, 0)
}

func (s *EccentricityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EccentricityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EccentricityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EccentricityListener); ok {
		listenerT.EnterEccentricity(s)
	}
}

func (s *EccentricityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EccentricityListener); ok {
		listenerT.ExitEccentricity(s)
	}
}

func (p *EccentricityParser) Eccentricity() (localctx IEccentricityContext) {
	this := p
	_ = this

	localctx = NewEccentricityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EccentricityParserRULE_eccentricity)

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
		p.Match(EccentricityParserECCENTRICITY)
	}

	return localctx
}
