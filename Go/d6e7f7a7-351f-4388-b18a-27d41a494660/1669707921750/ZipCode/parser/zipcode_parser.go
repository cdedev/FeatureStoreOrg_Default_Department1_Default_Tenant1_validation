// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707921750/ZipCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ZipCode

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

type ZipCodeParser struct {
	*antlr.BaseParser
}

var zipcodeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zipcodeParserInit() {
	staticData := &zipcodeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ZIPCODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "zipcode",
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

// ZipCodeParserInit initializes any static state used to implement ZipCodeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewZipCodeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZipCodeParserInit() {
	staticData := &zipcodeParserStaticData
	staticData.once.Do(zipcodeParserInit)
}

// NewZipCodeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewZipCodeParser(input antlr.TokenStream) *ZipCodeParser {
	ZipCodeParserInit()
	this := new(ZipCodeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &zipcodeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ZipCode.g4"

	return this
}

// ZipCodeParser tokens.
const (
	ZipCodeParserEOF      = antlr.TokenEOF
	ZipCodeParserZIPCODE  = 1
	ZipCodeParserOWNKEY   = 2
	ZipCodeParserSPLITKEY = 3
	ZipCodeParserWS       = 4
)

// ZipCodeParser rules.
const (
	ZipCodeParserRULE_expression = 0
	ZipCodeParserRULE_zipcode    = 1
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
	p.RuleIndex = ZipCodeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZipCodeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Zipcode() IZipcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZipcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZipcodeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZipCodeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZipCodeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZipCodeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ZipCodeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZipCodeParserRULE_expression)

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
		p.Zipcode()
	}
	{
		p.SetState(5)
		p.Match(ZipCodeParserEOF)
	}

	return localctx
}

// IZipcodeContext is an interface to support dynamic dispatch.
type IZipcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZipcodeContext differentiates from other interfaces.
	IsZipcodeContext()
}

type ZipcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZipcodeContext() *ZipcodeContext {
	var p = new(ZipcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZipCodeParserRULE_zipcode
	return p
}

func (*ZipcodeContext) IsZipcodeContext() {}

func NewZipcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZipcodeContext {
	var p = new(ZipcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZipCodeParserRULE_zipcode

	return p
}

func (s *ZipcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *ZipcodeContext) ZIPCODE() antlr.TerminalNode {
	return s.GetToken(ZipCodeParserZIPCODE, 0)
}

func (s *ZipcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZipcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZipcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZipCodeListener); ok {
		listenerT.EnterZipcode(s)
	}
}

func (s *ZipcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZipCodeListener); ok {
		listenerT.ExitZipcode(s)
	}
}

func (p *ZipCodeParser) Zipcode() (localctx IZipcodeContext) {
	this := p
	_ = this

	localctx = NewZipcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZipCodeParserRULE_zipcode)

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
		p.Match(ZipCodeParserZIPCODE)
	}

	return localctx
}
