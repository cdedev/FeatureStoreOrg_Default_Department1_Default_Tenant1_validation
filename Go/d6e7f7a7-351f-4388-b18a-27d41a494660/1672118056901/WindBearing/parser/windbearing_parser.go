// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118056901/WindBearing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WindBearing

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

type WindBearingParser struct {
	*antlr.BaseParser
}

var windbearingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func windbearingParserInit() {
	staticData := &windbearingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WINDBEARING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "windbearing",
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

// WindBearingParserInit initializes any static state used to implement WindBearingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWindBearingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WindBearingParserInit() {
	staticData := &windbearingParserStaticData
	staticData.once.Do(windbearingParserInit)
}

// NewWindBearingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWindBearingParser(input antlr.TokenStream) *WindBearingParser {
	WindBearingParserInit()
	this := new(WindBearingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &windbearingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "WindBearing.g4"

	return this
}

// WindBearingParser tokens.
const (
	WindBearingParserEOF         = antlr.TokenEOF
	WindBearingParserWINDBEARING = 1
	WindBearingParserOWNKEY      = 2
	WindBearingParserSPLITKEY    = 3
	WindBearingParserWS          = 4
)

// WindBearingParser rules.
const (
	WindBearingParserRULE_expression  = 0
	WindBearingParserRULE_windbearing = 1
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
	p.RuleIndex = WindBearingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WindBearingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Windbearing() IWindbearingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWindbearingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWindbearingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WindBearingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindBearingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindBearingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WindBearingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WindBearingParserRULE_expression)

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
		p.Windbearing()
	}
	{
		p.SetState(5)
		p.Match(WindBearingParserEOF)
	}

	return localctx
}

// IWindbearingContext is an interface to support dynamic dispatch.
type IWindbearingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWindbearingContext differentiates from other interfaces.
	IsWindbearingContext()
}

type WindbearingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWindbearingContext() *WindbearingContext {
	var p = new(WindbearingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WindBearingParserRULE_windbearing
	return p
}

func (*WindbearingContext) IsWindbearingContext() {}

func NewWindbearingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WindbearingContext {
	var p = new(WindbearingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WindBearingParserRULE_windbearing

	return p
}

func (s *WindbearingContext) GetParser() antlr.Parser { return s.parser }

func (s *WindbearingContext) WINDBEARING() antlr.TerminalNode {
	return s.GetToken(WindBearingParserWINDBEARING, 0)
}

func (s *WindbearingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WindbearingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WindbearingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindBearingListener); ok {
		listenerT.EnterWindbearing(s)
	}
}

func (s *WindbearingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WindBearingListener); ok {
		listenerT.ExitWindbearing(s)
	}
}

func (p *WindBearingParser) Windbearing() (localctx IWindbearingContext) {
	this := p
	_ = this

	localctx = NewWindbearingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WindBearingParserRULE_windbearing)

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
		p.Match(WindBearingParserWINDBEARING)
	}

	return localctx
}
