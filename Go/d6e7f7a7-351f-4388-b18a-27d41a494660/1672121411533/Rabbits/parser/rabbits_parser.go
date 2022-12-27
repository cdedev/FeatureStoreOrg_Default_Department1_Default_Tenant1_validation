// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121411533/Rabbits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rabbits

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

type RabbitsParser struct {
	*antlr.BaseParser
}

var rabbitsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rabbitsParserInit() {
	staticData := &rabbitsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RABBITS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "rabbits",
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

// RabbitsParserInit initializes any static state used to implement RabbitsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRabbitsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RabbitsParserInit() {
	staticData := &rabbitsParserStaticData
	staticData.once.Do(rabbitsParserInit)
}

// NewRabbitsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRabbitsParser(input antlr.TokenStream) *RabbitsParser {
	RabbitsParserInit()
	this := new(RabbitsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rabbitsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Rabbits.g4"

	return this
}

// RabbitsParser tokens.
const (
	RabbitsParserEOF      = antlr.TokenEOF
	RabbitsParserRABBITS  = 1
	RabbitsParserOWNKEY   = 2
	RabbitsParserSPLITKEY = 3
	RabbitsParserWS       = 4
)

// RabbitsParser rules.
const (
	RabbitsParserRULE_expression = 0
	RabbitsParserRULE_rabbits    = 1
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
	p.RuleIndex = RabbitsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RabbitsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Rabbits() IRabbitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRabbitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRabbitsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RabbitsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RabbitsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RabbitsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RabbitsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RabbitsParserRULE_expression)

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
		p.Rabbits()
	}
	{
		p.SetState(5)
		p.Match(RabbitsParserEOF)
	}

	return localctx
}

// IRabbitsContext is an interface to support dynamic dispatch.
type IRabbitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRabbitsContext differentiates from other interfaces.
	IsRabbitsContext()
}

type RabbitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRabbitsContext() *RabbitsContext {
	var p = new(RabbitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RabbitsParserRULE_rabbits
	return p
}

func (*RabbitsContext) IsRabbitsContext() {}

func NewRabbitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RabbitsContext {
	var p = new(RabbitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RabbitsParserRULE_rabbits

	return p
}

func (s *RabbitsContext) GetParser() antlr.Parser { return s.parser }

func (s *RabbitsContext) RABBITS() antlr.TerminalNode {
	return s.GetToken(RabbitsParserRABBITS, 0)
}

func (s *RabbitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RabbitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RabbitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RabbitsListener); ok {
		listenerT.EnterRabbits(s)
	}
}

func (s *RabbitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RabbitsListener); ok {
		listenerT.ExitRabbits(s)
	}
}

func (p *RabbitsParser) Rabbits() (localctx IRabbitsContext) {
	this := p
	_ = this

	localctx = NewRabbitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RabbitsParserRULE_rabbits)

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
		p.Match(RabbitsParserRABBITS)
	}

	return localctx
}
