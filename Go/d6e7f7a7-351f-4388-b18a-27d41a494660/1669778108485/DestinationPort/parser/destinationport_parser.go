// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778108485/DestinationPort.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DestinationPort

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

type DestinationPortParser struct {
	*antlr.BaseParser
}

var destinationportParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func destinationportParserInit() {
	staticData := &destinationportParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DESTINATIONPORT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "destinationport",
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

// DestinationPortParserInit initializes any static state used to implement DestinationPortParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDestinationPortParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DestinationPortParserInit() {
	staticData := &destinationportParserStaticData
	staticData.once.Do(destinationportParserInit)
}

// NewDestinationPortParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDestinationPortParser(input antlr.TokenStream) *DestinationPortParser {
	DestinationPortParserInit()
	this := new(DestinationPortParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &destinationportParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DestinationPort.g4"

	return this
}

// DestinationPortParser tokens.
const (
	DestinationPortParserEOF             = antlr.TokenEOF
	DestinationPortParserDESTINATIONPORT = 1
	DestinationPortParserOWNKEY          = 2
	DestinationPortParserSPLITKEY        = 3
	DestinationPortParserWS              = 4
)

// DestinationPortParser rules.
const (
	DestinationPortParserRULE_expression      = 0
	DestinationPortParserRULE_destinationport = 1
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
	p.RuleIndex = DestinationPortParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DestinationPortParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Destinationport() IDestinationportContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestinationportContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestinationportContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DestinationPortParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DestinationPortListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DestinationPortListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DestinationPortParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DestinationPortParserRULE_expression)

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
		p.Destinationport()
	}
	{
		p.SetState(5)
		p.Match(DestinationPortParserEOF)
	}

	return localctx
}

// IDestinationportContext is an interface to support dynamic dispatch.
type IDestinationportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestinationportContext differentiates from other interfaces.
	IsDestinationportContext()
}

type DestinationportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinationportContext() *DestinationportContext {
	var p = new(DestinationportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DestinationPortParserRULE_destinationport
	return p
}

func (*DestinationportContext) IsDestinationportContext() {}

func NewDestinationportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationportContext {
	var p = new(DestinationportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DestinationPortParserRULE_destinationport

	return p
}

func (s *DestinationportContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationportContext) DESTINATIONPORT() antlr.TerminalNode {
	return s.GetToken(DestinationPortParserDESTINATIONPORT, 0)
}

func (s *DestinationportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DestinationPortListener); ok {
		listenerT.EnterDestinationport(s)
	}
}

func (s *DestinationportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DestinationPortListener); ok {
		listenerT.ExitDestinationport(s)
	}
}

func (p *DestinationPortParser) Destinationport() (localctx IDestinationportContext) {
	this := p
	_ = this

	localctx = NewDestinationportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DestinationPortParserRULE_destinationport)

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
		p.Match(DestinationPortParserDESTINATIONPORT)
	}

	return localctx
}
