// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674666543500/Gamma.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gamma

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

type GammaParser struct {
	*antlr.BaseParser
}

var gammaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gammaParserInit() {
	staticData := &gammaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GAMMA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gamma",
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

// GammaParserInit initializes any static state used to implement GammaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGammaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GammaParserInit() {
	staticData := &gammaParserStaticData
	staticData.once.Do(gammaParserInit)
}

// NewGammaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGammaParser(input antlr.TokenStream) *GammaParser {
	GammaParserInit()
	this := new(GammaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gammaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Gamma.g4"

	return this
}

// GammaParser tokens.
const (
	GammaParserEOF      = antlr.TokenEOF
	GammaParserGAMMA    = 1
	GammaParserOWNKEY   = 2
	GammaParserSPLITKEY = 3
	GammaParserWS       = 4
)

// GammaParser rules.
const (
	GammaParserRULE_expression = 0
	GammaParserRULE_gamma      = 1
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
	p.RuleIndex = GammaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GammaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gamma() IGammaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGammaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGammaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GammaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GammaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GammaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GammaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GammaParserRULE_expression)

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
		p.Gamma()
	}
	{
		p.SetState(5)
		p.Match(GammaParserEOF)
	}

	return localctx
}

// IGammaContext is an interface to support dynamic dispatch.
type IGammaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGammaContext differentiates from other interfaces.
	IsGammaContext()
}

type GammaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGammaContext() *GammaContext {
	var p = new(GammaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GammaParserRULE_gamma
	return p
}

func (*GammaContext) IsGammaContext() {}

func NewGammaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GammaContext {
	var p = new(GammaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GammaParserRULE_gamma

	return p
}

func (s *GammaContext) GetParser() antlr.Parser { return s.parser }

func (s *GammaContext) GAMMA() antlr.TerminalNode {
	return s.GetToken(GammaParserGAMMA, 0)
}

func (s *GammaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GammaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GammaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GammaListener); ok {
		listenerT.EnterGamma(s)
	}
}

func (s *GammaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GammaListener); ok {
		listenerT.ExitGamma(s)
	}
}

func (p *GammaParser) Gamma() (localctx IGammaContext) {
	this := p
	_ = this

	localctx = NewGammaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GammaParserRULE_gamma)

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
		p.Match(GammaParserGAMMA)
	}

	return localctx
}
