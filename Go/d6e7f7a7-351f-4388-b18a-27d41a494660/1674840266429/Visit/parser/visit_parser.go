// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674840266429/Visit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visit

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

type VisitParser struct {
	*antlr.BaseParser
}

var visitParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func visitParserInit() {
	staticData := &visitParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VISIT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "visit",
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

// VisitParserInit initializes any static state used to implement VisitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVisitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VisitParserInit() {
	staticData := &visitParserStaticData
	staticData.once.Do(visitParserInit)
}

// NewVisitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVisitParser(input antlr.TokenStream) *VisitParser {
	VisitParserInit()
	this := new(VisitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &visitParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Visit.g4"

	return this
}

// VisitParser tokens.
const (
	VisitParserEOF      = antlr.TokenEOF
	VisitParserVISIT    = 1
	VisitParserOWNKEY   = 2
	VisitParserSPLITKEY = 3
	VisitParserWS       = 4
)

// VisitParser rules.
const (
	VisitParserRULE_expression = 0
	VisitParserRULE_visit      = 1
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
	p.RuleIndex = VisitParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VisitParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Visit() IVisitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVisitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVisitContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VisitParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisitListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisitListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VisitParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VisitParserRULE_expression)

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
		p.Visit()
	}
	{
		p.SetState(5)
		p.Match(VisitParserEOF)
	}

	return localctx
}

// IVisitContext is an interface to support dynamic dispatch.
type IVisitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVisitContext differentiates from other interfaces.
	IsVisitContext()
}

type VisitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVisitContext() *VisitContext {
	var p = new(VisitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VisitParserRULE_visit
	return p
}

func (*VisitContext) IsVisitContext() {}

func NewVisitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VisitContext {
	var p = new(VisitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VisitParserRULE_visit

	return p
}

func (s *VisitContext) GetParser() antlr.Parser { return s.parser }

func (s *VisitContext) VISIT() antlr.TerminalNode {
	return s.GetToken(VisitParserVISIT, 0)
}

func (s *VisitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VisitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VisitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisitListener); ok {
		listenerT.EnterVisit(s)
	}
}

func (s *VisitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VisitListener); ok {
		listenerT.ExitVisit(s)
	}
}

func (p *VisitParser) Visit() (localctx IVisitContext) {
	this := p
	_ = this

	localctx = NewVisitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VisitParserRULE_visit)

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
		p.Match(VisitParserVISIT)
	}

	return localctx
}
