// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884881303/AvgNoiseLvl.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AvgNoiseLvl

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

type AvgNoiseLvlParser struct {
	*antlr.BaseParser
}

var avgnoiselvlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func avgnoiselvlParserInit() {
	staticData := &avgnoiselvlParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AVGNOISELVL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "avgnoiselvl",
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

// AvgNoiseLvlParserInit initializes any static state used to implement AvgNoiseLvlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAvgNoiseLvlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AvgNoiseLvlParserInit() {
	staticData := &avgnoiselvlParserStaticData
	staticData.once.Do(avgnoiselvlParserInit)
}

// NewAvgNoiseLvlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAvgNoiseLvlParser(input antlr.TokenStream) *AvgNoiseLvlParser {
	AvgNoiseLvlParserInit()
	this := new(AvgNoiseLvlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &avgnoiselvlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AvgNoiseLvl.g4"

	return this
}

// AvgNoiseLvlParser tokens.
const (
	AvgNoiseLvlParserEOF         = antlr.TokenEOF
	AvgNoiseLvlParserAVGNOISELVL = 1
	AvgNoiseLvlParserOWNKEY      = 2
	AvgNoiseLvlParserSPLITKEY    = 3
	AvgNoiseLvlParserWS          = 4
)

// AvgNoiseLvlParser rules.
const (
	AvgNoiseLvlParserRULE_expression  = 0
	AvgNoiseLvlParserRULE_avgnoiselvl = 1
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
	p.RuleIndex = AvgNoiseLvlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AvgNoiseLvlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Avgnoiselvl() IAvgnoiselvlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAvgnoiselvlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAvgnoiselvlContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AvgNoiseLvlParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgNoiseLvlListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgNoiseLvlListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AvgNoiseLvlParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AvgNoiseLvlParserRULE_expression)

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
		p.Avgnoiselvl()
	}
	{
		p.SetState(5)
		p.Match(AvgNoiseLvlParserEOF)
	}

	return localctx
}

// IAvgnoiselvlContext is an interface to support dynamic dispatch.
type IAvgnoiselvlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAvgnoiselvlContext differentiates from other interfaces.
	IsAvgnoiselvlContext()
}

type AvgnoiselvlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvgnoiselvlContext() *AvgnoiselvlContext {
	var p = new(AvgnoiselvlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AvgNoiseLvlParserRULE_avgnoiselvl
	return p
}

func (*AvgnoiselvlContext) IsAvgnoiselvlContext() {}

func NewAvgnoiselvlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvgnoiselvlContext {
	var p = new(AvgnoiselvlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AvgNoiseLvlParserRULE_avgnoiselvl

	return p
}

func (s *AvgnoiselvlContext) GetParser() antlr.Parser { return s.parser }

func (s *AvgnoiselvlContext) AVGNOISELVL() antlr.TerminalNode {
	return s.GetToken(AvgNoiseLvlParserAVGNOISELVL, 0)
}

func (s *AvgnoiselvlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvgnoiselvlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AvgnoiselvlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgNoiseLvlListener); ok {
		listenerT.EnterAvgnoiselvl(s)
	}
}

func (s *AvgnoiselvlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgNoiseLvlListener); ok {
		listenerT.ExitAvgnoiselvl(s)
	}
}

func (p *AvgNoiseLvlParser) Avgnoiselvl() (localctx IAvgnoiselvlContext) {
	this := p
	_ = this

	localctx = NewAvgnoiselvlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AvgNoiseLvlParserRULE_avgnoiselvl)

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
		p.Match(AvgNoiseLvlParserAVGNOISELVL)
	}

	return localctx
}
