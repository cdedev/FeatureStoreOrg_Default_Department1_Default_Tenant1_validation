// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657111128/Wind.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wind

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

type WindParser struct {
	*antlr.BaseParser
}

var windParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func windParserInit() {
	staticData := &windParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WIND", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "wind",
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

// WindParserInit initializes any static state used to implement WindParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWindParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WindParserInit() {
	staticData := &windParserStaticData
	staticData.once.Do(windParserInit)
}

// NewWindParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWindParser(input antlr.TokenStream) *WindParser {
	WindParserInit()
	this := new(WindParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &windParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Wind.g4"

	return this
}

// WindParser tokens.
const (
	WindParserEOF      = antlr.TokenEOF
	WindParserWIND     = 1
	WindParserOWNKEY   = 2
	WindParserSPLITKEY = 3
	WindParserWS       = 4
)

// WindParser rules.
const (
	WindParserRULE_expression = 0
	WindParserRULE_wind       = 1
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
	p.RuleIndex = WindParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WindParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Wind() IWindContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WindParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WindParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WindParserRULE_expression)

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
		p.Wind()
	}
	{
		p.SetState(5)
		p.Match(WindParserEOF)
	}

	return localctx
}

// IWindContext is an interface to support dynamic dispatch.
type IWindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindContext differentiates from other interfaces.
	IsWindContext()
}

type WindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindContext() *WindContext {
	var p = new(WindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WindParserRULE_wind
	return p
}

func (*WindContext) IsWindContext() {}

func NewWindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindContext {
	var p = new(WindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WindParserRULE_wind

	return p
}

func (s *WindContext) GetParser() antlr.Parser { return s.parser }

func (s *WindContext) WIND() antlr.TerminalNode {
	return s.GetToken(WindParserWIND, 0)
}

func (s *WindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindListener); ok {
		listenerT.EnterWind(s)
	}
}

func (s *WindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindListener); ok {
		listenerT.ExitWind(s)
	}
}

func (p *WindParser) Wind() (localctx IWindContext) {
	this := p
	_ = this

	localctx = NewWindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WindParserRULE_wind)

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
		p.Match(WindParserWIND)
	}

	return localctx
}
