// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232685841/Kurtosis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Kurtosis

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

type KurtosisParser struct {
	*antlr.BaseParser
}

var kurtosisParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func kurtosisParserInit() {
	staticData := &kurtosisParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "KURTOSIS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "kurtosis",
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

// KurtosisParserInit initializes any static state used to implement KurtosisParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewKurtosisParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func KurtosisParserInit() {
	staticData := &kurtosisParserStaticData
	staticData.once.Do(kurtosisParserInit)
}

// NewKurtosisParser produces a new parser instance for the optional input antlr.TokenStream.
func NewKurtosisParser(input antlr.TokenStream) *KurtosisParser {
	KurtosisParserInit()
	this := new(KurtosisParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &kurtosisParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Kurtosis.g4"

	return this
}

// KurtosisParser tokens.
const (
	KurtosisParserEOF      = antlr.TokenEOF
	KurtosisParserKURTOSIS = 1
	KurtosisParserOWNKEY   = 2
	KurtosisParserSPLITKEY = 3
	KurtosisParserWS       = 4
)

// KurtosisParser rules.
const (
	KurtosisParserRULE_expression = 0
	KurtosisParserRULE_kurtosis   = 1
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
	p.RuleIndex = KurtosisParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KurtosisParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Kurtosis() IKurtosisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKurtosisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKurtosisContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(KurtosisParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KurtosisListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KurtosisListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *KurtosisParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, KurtosisParserRULE_expression)

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
		p.Kurtosis()
	}
	{
		p.SetState(5)
		p.Match(KurtosisParserEOF)
	}

	return localctx
}

// IKurtosisContext is an interface to support dynamic dispatch.
type IKurtosisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKurtosisContext differentiates from other interfaces.
	IsKurtosisContext()
}

type KurtosisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKurtosisContext() *KurtosisContext {
	var p = new(KurtosisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = KurtosisParserRULE_kurtosis
	return p
}

func (*KurtosisContext) IsKurtosisContext() {}

func NewKurtosisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KurtosisContext {
	var p = new(KurtosisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = KurtosisParserRULE_kurtosis

	return p
}

func (s *KurtosisContext) GetParser() antlr.Parser { return s.parser }

func (s *KurtosisContext) KURTOSIS() antlr.TerminalNode {
	return s.GetToken(KurtosisParserKURTOSIS, 0)
}

func (s *KurtosisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KurtosisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KurtosisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KurtosisListener); ok {
		listenerT.EnterKurtosis(s)
	}
}

func (s *KurtosisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(KurtosisListener); ok {
		listenerT.ExitKurtosis(s)
	}
}

func (p *KurtosisParser) Kurtosis() (localctx IKurtosisContext) {
	this := p
	_ = this

	localctx = NewKurtosisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, KurtosisParserRULE_kurtosis)

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
		p.Match(KurtosisParserKURTOSIS)
	}

	return localctx
}
