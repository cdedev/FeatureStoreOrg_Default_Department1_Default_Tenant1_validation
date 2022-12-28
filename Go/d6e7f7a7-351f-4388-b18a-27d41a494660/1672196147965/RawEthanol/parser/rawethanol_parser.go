// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196147965/RawEthanol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RawEthanol

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

type RawEthanolParser struct {
	*antlr.BaseParser
}

var rawethanolParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rawethanolParserInit() {
	staticData := &rawethanolParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RAWETHANOL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rawethanol",
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

// RawEthanolParserInit initializes any static state used to implement RawEthanolParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRawEthanolParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RawEthanolParserInit() {
	staticData := &rawethanolParserStaticData
	staticData.once.Do(rawethanolParserInit)
}

// NewRawEthanolParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRawEthanolParser(input antlr.TokenStream) *RawEthanolParser {
	RawEthanolParserInit()
	this := new(RawEthanolParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rawethanolParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RawEthanol.g4"

	return this
}

// RawEthanolParser tokens.
const (
	RawEthanolParserEOF        = antlr.TokenEOF
	RawEthanolParserRAWETHANOL = 1
	RawEthanolParserOWNKEY     = 2
	RawEthanolParserSPLITKEY   = 3
	RawEthanolParserWS         = 4
)

// RawEthanolParser rules.
const (
	RawEthanolParserRULE_expression = 0
	RawEthanolParserRULE_rawethanol = 1
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
	p.RuleIndex = RawEthanolParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RawEthanolParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rawethanol() IRawethanolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRawethanolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRawethanolContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RawEthanolParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawEthanolListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawEthanolListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RawEthanolParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RawEthanolParserRULE_expression)

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
		p.Rawethanol()
	}
	{
		p.SetState(5)
		p.Match(RawEthanolParserEOF)
	}

	return localctx
}

// IRawethanolContext is an interface to support dynamic dispatch.
type IRawethanolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawethanolContext differentiates from other interfaces.
	IsRawethanolContext()
}

type RawethanolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawethanolContext() *RawethanolContext {
	var p = new(RawethanolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RawEthanolParserRULE_rawethanol
	return p
}

func (*RawethanolContext) IsRawethanolContext() {}

func NewRawethanolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawethanolContext {
	var p = new(RawethanolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RawEthanolParserRULE_rawethanol

	return p
}

func (s *RawethanolContext) GetParser() antlr.Parser { return s.parser }

func (s *RawethanolContext) RAWETHANOL() antlr.TerminalNode {
	return s.GetToken(RawEthanolParserRAWETHANOL, 0)
}

func (s *RawethanolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawethanolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawethanolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawEthanolListener); ok {
		listenerT.EnterRawethanol(s)
	}
}

func (s *RawethanolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RawEthanolListener); ok {
		listenerT.ExitRawethanol(s)
	}
}

func (p *RawEthanolParser) Rawethanol() (localctx IRawethanolContext) {
	this := p
	_ = this

	localctx = NewRawethanolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RawEthanolParserRULE_rawethanol)

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
		p.Match(RawEthanolParserRAWETHANOL)
	}

	return localctx
}
