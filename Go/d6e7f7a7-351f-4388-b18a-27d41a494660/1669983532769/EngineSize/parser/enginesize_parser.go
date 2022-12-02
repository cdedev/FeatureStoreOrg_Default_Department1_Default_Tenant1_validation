// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983532769/EngineSize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EngineSize

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

type EngineSizeParser struct {
	*antlr.BaseParser
}

var enginesizeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func enginesizeParserInit() {
	staticData := &enginesizeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ENGINESIZE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "enginesize",
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

// EngineSizeParserInit initializes any static state used to implement EngineSizeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEngineSizeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EngineSizeParserInit() {
	staticData := &enginesizeParserStaticData
	staticData.once.Do(enginesizeParserInit)
}

// NewEngineSizeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEngineSizeParser(input antlr.TokenStream) *EngineSizeParser {
	EngineSizeParserInit()
	this := new(EngineSizeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &enginesizeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "EngineSize.g4"

	return this
}

// EngineSizeParser tokens.
const (
	EngineSizeParserEOF        = antlr.TokenEOF
	EngineSizeParserENGINESIZE = 1
	EngineSizeParserOWNKEY     = 2
	EngineSizeParserSPLITKEY   = 3
	EngineSizeParserWS         = 4
)

// EngineSizeParser rules.
const (
	EngineSizeParserRULE_expression = 0
	EngineSizeParserRULE_enginesize = 1
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
	p.RuleIndex = EngineSizeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EngineSizeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Enginesize() IEnginesizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnginesizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnginesizeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EngineSizeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineSizeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineSizeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EngineSizeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EngineSizeParserRULE_expression)

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
		p.Enginesize()
	}
	{
		p.SetState(5)
		p.Match(EngineSizeParserEOF)
	}

	return localctx
}

// IEnginesizeContext is an interface to support dynamic dispatch.
type IEnginesizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnginesizeContext differentiates from other interfaces.
	IsEnginesizeContext()
}

type EnginesizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnginesizeContext() *EnginesizeContext {
	var p = new(EnginesizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EngineSizeParserRULE_enginesize
	return p
}

func (*EnginesizeContext) IsEnginesizeContext() {}

func NewEnginesizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnginesizeContext {
	var p = new(EnginesizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EngineSizeParserRULE_enginesize

	return p
}

func (s *EnginesizeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnginesizeContext) ENGINESIZE() antlr.TerminalNode {
	return s.GetToken(EngineSizeParserENGINESIZE, 0)
}

func (s *EnginesizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnginesizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnginesizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineSizeListener); ok {
		listenerT.EnterEnginesize(s)
	}
}

func (s *EnginesizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineSizeListener); ok {
		listenerT.ExitEnginesize(s)
	}
}

func (p *EngineSizeParser) Enginesize() (localctx IEnginesizeContext) {
	this := p
	_ = this

	localctx = NewEnginesizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EngineSizeParserRULE_enginesize)

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
		p.Match(EngineSizeParserENGINESIZE)
	}

	return localctx
}
