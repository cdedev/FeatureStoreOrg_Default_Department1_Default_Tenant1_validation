// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377318029/DualSim.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DualSim

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

type DualSimParser struct {
	*antlr.BaseParser
}

var dualsimParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dualsimParserInit() {
	staticData := &dualsimParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DUALSIM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dualsim",
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

// DualSimParserInit initializes any static state used to implement DualSimParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDualSimParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DualSimParserInit() {
	staticData := &dualsimParserStaticData
	staticData.once.Do(dualsimParserInit)
}

// NewDualSimParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDualSimParser(input antlr.TokenStream) *DualSimParser {
	DualSimParserInit()
	this := new(DualSimParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dualsimParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DualSim.g4"

	return this
}

// DualSimParser tokens.
const (
	DualSimParserEOF      = antlr.TokenEOF
	DualSimParserDUALSIM  = 1
	DualSimParserOWNKEY   = 2
	DualSimParserSPLITKEY = 3
	DualSimParserWS       = 4
)

// DualSimParser rules.
const (
	DualSimParserRULE_expression = 0
	DualSimParserRULE_dualsim    = 1
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
	p.RuleIndex = DualSimParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DualSimParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dualsim() IDualsimContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDualsimContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDualsimContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DualSimParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DualSimListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DualSimListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DualSimParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DualSimParserRULE_expression)

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
		p.Dualsim()
	}
	{
		p.SetState(5)
		p.Match(DualSimParserEOF)
	}

	return localctx
}

// IDualsimContext is an interface to support dynamic dispatch.
type IDualsimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDualsimContext differentiates from other interfaces.
	IsDualsimContext()
}

type DualsimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDualsimContext() *DualsimContext {
	var p = new(DualsimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DualSimParserRULE_dualsim
	return p
}

func (*DualsimContext) IsDualsimContext() {}

func NewDualsimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DualsimContext {
	var p = new(DualsimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DualSimParserRULE_dualsim

	return p
}

func (s *DualsimContext) GetParser() antlr.Parser { return s.parser }

func (s *DualsimContext) DUALSIM() antlr.TerminalNode {
	return s.GetToken(DualSimParserDUALSIM, 0)
}

func (s *DualsimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DualsimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DualsimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DualSimListener); ok {
		listenerT.EnterDualsim(s)
	}
}

func (s *DualsimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DualSimListener); ok {
		listenerT.ExitDualsim(s)
	}
}

func (p *DualSimParser) Dualsim() (localctx IDualsimContext) {
	this := p
	_ = this

	localctx = NewDualsimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DualSimParserRULE_dualsim)

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
		p.Match(DualSimParserDUALSIM)
	}

	return localctx
}
