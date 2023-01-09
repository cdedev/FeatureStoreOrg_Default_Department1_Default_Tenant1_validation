// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238830271/Oily.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Oily

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

type OilyParser struct {
	*antlr.BaseParser
}

var oilyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func oilyParserInit() {
	staticData := &oilyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OILY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "oily",
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

// OilyParserInit initializes any static state used to implement OilyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOilyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OilyParserInit() {
	staticData := &oilyParserStaticData
	staticData.once.Do(oilyParserInit)
}

// NewOilyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOilyParser(input antlr.TokenStream) *OilyParser {
	OilyParserInit()
	this := new(OilyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &oilyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Oily.g4"

	return this
}

// OilyParser tokens.
const (
	OilyParserEOF      = antlr.TokenEOF
	OilyParserOILY     = 1
	OilyParserOWNKEY   = 2
	OilyParserSPLITKEY = 3
	OilyParserWS       = 4
)

// OilyParser rules.
const (
	OilyParserRULE_expression = 0
	OilyParserRULE_oily       = 1
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
	p.RuleIndex = OilyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OilyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Oily() IOilyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOilyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOilyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OilyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OilyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OilyParserRULE_expression)

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
		p.Oily()
	}
	{
		p.SetState(5)
		p.Match(OilyParserEOF)
	}

	return localctx
}

// IOilyContext is an interface to support dynamic dispatch.
type IOilyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOilyContext differentiates from other interfaces.
	IsOilyContext()
}

type OilyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOilyContext() *OilyContext {
	var p = new(OilyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OilyParserRULE_oily
	return p
}

func (*OilyContext) IsOilyContext() {}

func NewOilyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OilyContext {
	var p = new(OilyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OilyParserRULE_oily

	return p
}

func (s *OilyContext) GetParser() antlr.Parser { return s.parser }

func (s *OilyContext) OILY() antlr.TerminalNode {
	return s.GetToken(OilyParserOILY, 0)
}

func (s *OilyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OilyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OilyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilyListener); ok {
		listenerT.EnterOily(s)
	}
}

func (s *OilyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OilyListener); ok {
		listenerT.ExitOily(s)
	}
}

func (p *OilyParser) Oily() (localctx IOilyContext) {
	this := p
	_ = this

	localctx = NewOilyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OilyParserRULE_oily)

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
		p.Match(OilyParserOILY)
	}

	return localctx
}
