// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284951145/IntenseViolence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntenseViolence

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

type IntenseViolenceParser struct {
	*antlr.BaseParser
}

var intenseviolenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func intenseviolenceParserInit() {
	staticData := &intenseviolenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INTENSEVIOLENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "intenseviolence",
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

// IntenseViolenceParserInit initializes any static state used to implement IntenseViolenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIntenseViolenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IntenseViolenceParserInit() {
	staticData := &intenseviolenceParserStaticData
	staticData.once.Do(intenseviolenceParserInit)
}

// NewIntenseViolenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIntenseViolenceParser(input antlr.TokenStream) *IntenseViolenceParser {
	IntenseViolenceParserInit()
	this := new(IntenseViolenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &intenseviolenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "IntenseViolence.g4"

	return this
}

// IntenseViolenceParser tokens.
const (
	IntenseViolenceParserEOF             = antlr.TokenEOF
	IntenseViolenceParserINTENSEVIOLENCE = 1
	IntenseViolenceParserOWNKEY          = 2
	IntenseViolenceParserSPLITKEY        = 3
	IntenseViolenceParserWS              = 4
)

// IntenseViolenceParser rules.
const (
	IntenseViolenceParserRULE_expression      = 0
	IntenseViolenceParserRULE_intenseviolence = 1
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
	p.RuleIndex = IntenseViolenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntenseViolenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Intenseviolence() IIntenseviolenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntenseviolenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntenseviolenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IntenseViolenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntenseViolenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntenseViolenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IntenseViolenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IntenseViolenceParserRULE_expression)

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
		p.Intenseviolence()
	}
	{
		p.SetState(5)
		p.Match(IntenseViolenceParserEOF)
	}

	return localctx
}

// IIntenseviolenceContext is an interface to support dynamic dispatch.
type IIntenseviolenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntenseviolenceContext differentiates from other interfaces.
	IsIntenseviolenceContext()
}

type IntenseviolenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntenseviolenceContext() *IntenseviolenceContext {
	var p = new(IntenseviolenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IntenseViolenceParserRULE_intenseviolence
	return p
}

func (*IntenseviolenceContext) IsIntenseviolenceContext() {}

func NewIntenseviolenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntenseviolenceContext {
	var p = new(IntenseviolenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IntenseViolenceParserRULE_intenseviolence

	return p
}

func (s *IntenseviolenceContext) GetParser() antlr.Parser { return s.parser }

func (s *IntenseviolenceContext) INTENSEVIOLENCE() antlr.TerminalNode {
	return s.GetToken(IntenseViolenceParserINTENSEVIOLENCE, 0)
}

func (s *IntenseviolenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntenseviolenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntenseviolenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntenseViolenceListener); ok {
		listenerT.EnterIntenseviolence(s)
	}
}

func (s *IntenseviolenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IntenseViolenceListener); ok {
		listenerT.ExitIntenseviolence(s)
	}
}

func (p *IntenseViolenceParser) Intenseviolence() (localctx IIntenseviolenceContext) {
	this := p
	_ = this

	localctx = NewIntenseviolenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IntenseViolenceParserRULE_intenseviolence)

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
		p.Match(IntenseViolenceParserINTENSEVIOLENCE)
	}

	return localctx
}
