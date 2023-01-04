// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672809764289/Ram.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ram

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

type RamParser struct {
	*antlr.BaseParser
}

var ramParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ramParserInit() {
	staticData := &ramParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RAM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ram",
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

// RamParserInit initializes any static state used to implement RamParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRamParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RamParserInit() {
	staticData := &ramParserStaticData
	staticData.once.Do(ramParserInit)
}

// NewRamParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRamParser(input antlr.TokenStream) *RamParser {
	RamParserInit()
	this := new(RamParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ramParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Ram.g4"

	return this
}

// RamParser tokens.
const (
	RamParserEOF      = antlr.TokenEOF
	RamParserRAM      = 1
	RamParserOWNKEY   = 2
	RamParserSPLITKEY = 3
	RamParserWS       = 4
)

// RamParser rules.
const (
	RamParserRULE_expression = 0
	RamParserRULE_ram        = 1
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
	p.RuleIndex = RamParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RamParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ram() IRamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRamContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RamParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RamListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RamListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RamParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RamParserRULE_expression)

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
		p.Ram()
	}
	{
		p.SetState(5)
		p.Match(RamParserEOF)
	}

	return localctx
}

// IRamContext is an interface to support dynamic dispatch.
type IRamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRamContext differentiates from other interfaces.
	IsRamContext()
}

type RamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRamContext() *RamContext {
	var p = new(RamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RamParserRULE_ram
	return p
}

func (*RamContext) IsRamContext() {}

func NewRamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RamContext {
	var p = new(RamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RamParserRULE_ram

	return p
}

func (s *RamContext) GetParser() antlr.Parser { return s.parser }

func (s *RamContext) RAM() antlr.TerminalNode {
	return s.GetToken(RamParserRAM, 0)
}

func (s *RamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RamListener); ok {
		listenerT.EnterRam(s)
	}
}

func (s *RamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RamListener); ok {
		listenerT.ExitRam(s)
	}
}

func (p *RamParser) Ram() (localctx IRamContext) {
	this := p
	_ = this

	localctx = NewRamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RamParserRULE_ram)

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
		p.Match(RamParserRAM)
	}

	return localctx
}
