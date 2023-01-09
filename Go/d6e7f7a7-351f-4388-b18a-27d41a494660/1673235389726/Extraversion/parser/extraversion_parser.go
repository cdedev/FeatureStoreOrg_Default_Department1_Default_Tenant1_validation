// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235389726/Extraversion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Extraversion

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

type ExtraversionParser struct {
	*antlr.BaseParser
}

var extraversionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func extraversionParserInit() {
	staticData := &extraversionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EXTRAVERSION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "extraversion",
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

// ExtraversionParserInit initializes any static state used to implement ExtraversionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExtraversionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExtraversionParserInit() {
	staticData := &extraversionParserStaticData
	staticData.once.Do(extraversionParserInit)
}

// NewExtraversionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExtraversionParser(input antlr.TokenStream) *ExtraversionParser {
	ExtraversionParserInit()
	this := new(ExtraversionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &extraversionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Extraversion.g4"

	return this
}

// ExtraversionParser tokens.
const (
	ExtraversionParserEOF          = antlr.TokenEOF
	ExtraversionParserEXTRAVERSION = 1
	ExtraversionParserOWNKEY       = 2
	ExtraversionParserSPLITKEY     = 3
	ExtraversionParserWS           = 4
)

// ExtraversionParser rules.
const (
	ExtraversionParserRULE_expression   = 0
	ExtraversionParserRULE_extraversion = 1
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
	p.RuleIndex = ExtraversionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExtraversionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Extraversion() IExtraversionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtraversionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtraversionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExtraversionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExtraversionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExtraversionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExtraversionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExtraversionParserRULE_expression)

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
		p.Extraversion()
	}
	{
		p.SetState(5)
		p.Match(ExtraversionParserEOF)
	}

	return localctx
}

// IExtraversionContext is an interface to support dynamic dispatch.
type IExtraversionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtraversionContext differentiates from other interfaces.
	IsExtraversionContext()
}

type ExtraversionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtraversionContext() *ExtraversionContext {
	var p = new(ExtraversionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExtraversionParserRULE_extraversion
	return p
}

func (*ExtraversionContext) IsExtraversionContext() {}

func NewExtraversionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtraversionContext {
	var p = new(ExtraversionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExtraversionParserRULE_extraversion

	return p
}

func (s *ExtraversionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtraversionContext) EXTRAVERSION() antlr.TerminalNode {
	return s.GetToken(ExtraversionParserEXTRAVERSION, 0)
}

func (s *ExtraversionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtraversionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtraversionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExtraversionListener); ok {
		listenerT.EnterExtraversion(s)
	}
}

func (s *ExtraversionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExtraversionListener); ok {
		listenerT.ExitExtraversion(s)
	}
}

func (p *ExtraversionParser) Extraversion() (localctx IExtraversionContext) {
	this := p
	_ = this

	localctx = NewExtraversionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExtraversionParserRULE_extraversion)

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
		p.Match(ExtraversionParserEXTRAVERSION)
	}

	return localctx
}
