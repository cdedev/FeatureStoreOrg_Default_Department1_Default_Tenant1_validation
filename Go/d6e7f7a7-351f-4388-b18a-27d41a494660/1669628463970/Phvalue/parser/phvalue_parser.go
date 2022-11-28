// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628463970/Phvalue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phvalue

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

type PhvalueParser struct {
	*antlr.BaseParser
}

var phvalueParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func phvalueParserInit() {
	staticData := &phvalueParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PHVALUE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "phvalue",
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

// PhvalueParserInit initializes any static state used to implement PhvalueParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPhvalueParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PhvalueParserInit() {
	staticData := &phvalueParserStaticData
	staticData.once.Do(phvalueParserInit)
}

// NewPhvalueParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPhvalueParser(input antlr.TokenStream) *PhvalueParser {
	PhvalueParserInit()
	this := new(PhvalueParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &phvalueParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Phvalue.g4"

	return this
}

// PhvalueParser tokens.
const (
	PhvalueParserEOF      = antlr.TokenEOF
	PhvalueParserPHVALUE  = 1
	PhvalueParserOWNKEY   = 2
	PhvalueParserSPLITKEY = 3
	PhvalueParserWS       = 4
)

// PhvalueParser rules.
const (
	PhvalueParserRULE_expression = 0
	PhvalueParserRULE_phvalue    = 1
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
	p.RuleIndex = PhvalueParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhvalueParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Phvalue() IPhvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhvalueContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PhvalueParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhvalueListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhvalueListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PhvalueParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PhvalueParserRULE_expression)

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
		p.Phvalue()
	}
	{
		p.SetState(5)
		p.Match(PhvalueParserEOF)
	}

	return localctx
}

// IPhvalueContext is an interface to support dynamic dispatch.
type IPhvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhvalueContext differentiates from other interfaces.
	IsPhvalueContext()
}

type PhvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhvalueContext() *PhvalueContext {
	var p = new(PhvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PhvalueParserRULE_phvalue
	return p
}

func (*PhvalueContext) IsPhvalueContext() {}

func NewPhvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhvalueContext {
	var p = new(PhvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhvalueParserRULE_phvalue

	return p
}

func (s *PhvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *PhvalueContext) PHVALUE() antlr.TerminalNode {
	return s.GetToken(PhvalueParserPHVALUE, 0)
}

func (s *PhvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhvalueListener); ok {
		listenerT.EnterPhvalue(s)
	}
}

func (s *PhvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhvalueListener); ok {
		listenerT.ExitPhvalue(s)
	}
}

func (p *PhvalueParser) Phvalue() (localctx IPhvalueContext) {
	this := p
	_ = this

	localctx = NewPhvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PhvalueParserRULE_phvalue)

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
		p.Match(PhvalueParserPHVALUE)
	}

	return localctx
}
