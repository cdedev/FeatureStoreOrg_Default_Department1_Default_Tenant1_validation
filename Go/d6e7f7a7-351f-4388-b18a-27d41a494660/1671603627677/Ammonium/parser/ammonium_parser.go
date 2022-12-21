// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603627677/Ammonium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ammonium

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

type AmmoniumParser struct {
	*antlr.BaseParser
}

var ammoniumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ammoniumParserInit() {
	staticData := &ammoniumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AMMONIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ammonium",
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

// AmmoniumParserInit initializes any static state used to implement AmmoniumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAmmoniumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AmmoniumParserInit() {
	staticData := &ammoniumParserStaticData
	staticData.once.Do(ammoniumParserInit)
}

// NewAmmoniumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAmmoniumParser(input antlr.TokenStream) *AmmoniumParser {
	AmmoniumParserInit()
	this := new(AmmoniumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ammoniumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Ammonium.g4"

	return this
}

// AmmoniumParser tokens.
const (
	AmmoniumParserEOF      = antlr.TokenEOF
	AmmoniumParserAMMONIUM = 1
	AmmoniumParserOWNKEY   = 2
	AmmoniumParserSPLITKEY = 3
	AmmoniumParserWS       = 4
)

// AmmoniumParser rules.
const (
	AmmoniumParserRULE_expression = 0
	AmmoniumParserRULE_ammonium   = 1
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
	p.RuleIndex = AmmoniumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmmoniumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ammonium() IAmmoniumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAmmoniumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAmmoniumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AmmoniumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmmoniumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmmoniumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AmmoniumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AmmoniumParserRULE_expression)

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
		p.Ammonium()
	}
	{
		p.SetState(5)
		p.Match(AmmoniumParserEOF)
	}

	return localctx
}

// IAmmoniumContext is an interface to support dynamic dispatch.
type IAmmoniumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmmoniumContext differentiates from other interfaces.
	IsAmmoniumContext()
}

type AmmoniumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmmoniumContext() *AmmoniumContext {
	var p = new(AmmoniumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AmmoniumParserRULE_ammonium
	return p
}

func (*AmmoniumContext) IsAmmoniumContext() {}

func NewAmmoniumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmmoniumContext {
	var p = new(AmmoniumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AmmoniumParserRULE_ammonium

	return p
}

func (s *AmmoniumContext) GetParser() antlr.Parser { return s.parser }

func (s *AmmoniumContext) AMMONIUM() antlr.TerminalNode {
	return s.GetToken(AmmoniumParserAMMONIUM, 0)
}

func (s *AmmoniumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmmoniumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AmmoniumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmmoniumListener); ok {
		listenerT.EnterAmmonium(s)
	}
}

func (s *AmmoniumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AmmoniumListener); ok {
		listenerT.ExitAmmonium(s)
	}
}

func (p *AmmoniumParser) Ammonium() (localctx IAmmoniumContext) {
	this := p
	_ = this

	localctx = NewAmmoniumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AmmoniumParserRULE_ammonium)

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
		p.Match(AmmoniumParserAMMONIUM)
	}

	return localctx
}
