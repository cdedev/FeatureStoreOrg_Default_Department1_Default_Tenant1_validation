// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693230058/Kilometer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Kilometer

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

type KilometerParser struct {
	*antlr.BaseParser
}

var kilometerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kilometerParserInit() {
	staticData := &kilometerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "KILOMETER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "kilometer",
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

// KilometerParserInit initializes any static state used to implement KilometerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKilometerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KilometerParserInit() {
	staticData := &kilometerParserStaticData
	staticData.once.Do(kilometerParserInit)
}

// NewKilometerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKilometerParser(input antlr.TokenStream) *KilometerParser {
	KilometerParserInit()
	this := new(KilometerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &kilometerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Kilometer.g4"

	return this
}

// KilometerParser tokens.
const (
	KilometerParserEOF       = antlr.TokenEOF
	KilometerParserKILOMETER = 1
	KilometerParserOWNKEY    = 2
	KilometerParserSPLITKEY  = 3
	KilometerParserWS        = 4
)

// KilometerParser rules.
const (
	KilometerParserRULE_expression = 0
	KilometerParserRULE_kilometer  = 1
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
	p.RuleIndex = KilometerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KilometerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Kilometer() IKilometerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKilometerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKilometerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(KilometerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilometerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilometerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *KilometerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KilometerParserRULE_expression)

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
		p.Kilometer()
	}
	{
		p.SetState(5)
		p.Match(KilometerParserEOF)
	}

	return localctx
}

// IKilometerContext is an interface to support dynamic dispatch.
type IKilometerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKilometerContext differentiates from other interfaces.
	IsKilometerContext()
}

type KilometerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKilometerContext() *KilometerContext {
	var p = new(KilometerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KilometerParserRULE_kilometer
	return p
}

func (*KilometerContext) IsKilometerContext() {}

func NewKilometerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KilometerContext {
	var p = new(KilometerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KilometerParserRULE_kilometer

	return p
}

func (s *KilometerContext) GetParser() antlr.Parser { return s.parser }

func (s *KilometerContext) KILOMETER() antlr.TerminalNode {
	return s.GetToken(KilometerParserKILOMETER, 0)
}

func (s *KilometerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KilometerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KilometerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilometerListener); ok {
		listenerT.EnterKilometer(s)
	}
}

func (s *KilometerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KilometerListener); ok {
		listenerT.ExitKilometer(s)
	}
}

func (p *KilometerParser) Kilometer() (localctx IKilometerContext) {
	this := p
	_ = this

	localctx = NewKilometerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KilometerParserRULE_kilometer)

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
		p.Match(KilometerParserKILOMETER)
	}

	return localctx
}
