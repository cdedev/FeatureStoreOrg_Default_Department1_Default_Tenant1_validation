// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376854362/LeatherInterior.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LeatherInterior

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

type LeatherInteriorParser struct {
	*antlr.BaseParser
}

var leatherinteriorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func leatherinteriorParserInit() {
	staticData := &leatherinteriorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LEATHERINTERIOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "leatherinterior",
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

// LeatherInteriorParserInit initializes any static state used to implement LeatherInteriorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLeatherInteriorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LeatherInteriorParserInit() {
	staticData := &leatherinteriorParserStaticData
	staticData.once.Do(leatherinteriorParserInit)
}

// NewLeatherInteriorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLeatherInteriorParser(input antlr.TokenStream) *LeatherInteriorParser {
	LeatherInteriorParserInit()
	this := new(LeatherInteriorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &leatherinteriorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "LeatherInterior.g4"

	return this
}

// LeatherInteriorParser tokens.
const (
	LeatherInteriorParserEOF             = antlr.TokenEOF
	LeatherInteriorParserLEATHERINTERIOR = 1
	LeatherInteriorParserOWNKEY          = 2
	LeatherInteriorParserSPLITKEY        = 3
	LeatherInteriorParserWS              = 4
)

// LeatherInteriorParser rules.
const (
	LeatherInteriorParserRULE_expression      = 0
	LeatherInteriorParserRULE_leatherinterior = 1
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
	p.RuleIndex = LeatherInteriorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LeatherInteriorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Leatherinterior() ILeatherinteriorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeatherinteriorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeatherinteriorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LeatherInteriorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeatherInteriorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeatherInteriorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LeatherInteriorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LeatherInteriorParserRULE_expression)

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
		p.Leatherinterior()
	}
	{
		p.SetState(5)
		p.Match(LeatherInteriorParserEOF)
	}

	return localctx
}

// ILeatherinteriorContext is an interface to support dynamic dispatch.
type ILeatherinteriorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeatherinteriorContext differentiates from other interfaces.
	IsLeatherinteriorContext()
}

type LeatherinteriorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeatherinteriorContext() *LeatherinteriorContext {
	var p = new(LeatherinteriorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LeatherInteriorParserRULE_leatherinterior
	return p
}

func (*LeatherinteriorContext) IsLeatherinteriorContext() {}

func NewLeatherinteriorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeatherinteriorContext {
	var p = new(LeatherinteriorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LeatherInteriorParserRULE_leatherinterior

	return p
}

func (s *LeatherinteriorContext) GetParser() antlr.Parser { return s.parser }

func (s *LeatherinteriorContext) LEATHERINTERIOR() antlr.TerminalNode {
	return s.GetToken(LeatherInteriorParserLEATHERINTERIOR, 0)
}

func (s *LeatherinteriorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeatherinteriorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeatherinteriorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeatherInteriorListener); ok {
		listenerT.EnterLeatherinterior(s)
	}
}

func (s *LeatherinteriorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LeatherInteriorListener); ok {
		listenerT.ExitLeatherinterior(s)
	}
}

func (p *LeatherInteriorParser) Leatherinterior() (localctx ILeatherinteriorContext) {
	this := p
	_ = this

	localctx = NewLeatherinteriorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LeatherInteriorParserRULE_leatherinterior)

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
		p.Match(LeatherInteriorParserLEATHERINTERIOR)
	}

	return localctx
}
