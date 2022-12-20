// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671513436611/Quarantine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Quarantine

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

type QuarantineParser struct {
	*antlr.BaseParser
}

var quarantineParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func quarantineParserInit() {
	staticData := &quarantineParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "QUARANTINE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "quarantine",
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

// QuarantineParserInit initializes any static state used to implement QuarantineParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewQuarantineParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func QuarantineParserInit() {
	staticData := &quarantineParserStaticData
	staticData.once.Do(quarantineParserInit)
}

// NewQuarantineParser produces a new parser instance for the optional input antlr.TokenStream.
func NewQuarantineParser(input antlr.TokenStream) *QuarantineParser {
	QuarantineParserInit()
	this := new(QuarantineParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &quarantineParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Quarantine.g4"

	return this
}

// QuarantineParser tokens.
const (
	QuarantineParserEOF        = antlr.TokenEOF
	QuarantineParserQUARANTINE = 1
	QuarantineParserOWNKEY     = 2
	QuarantineParserSPLITKEY   = 3
	QuarantineParserWS         = 4
)

// QuarantineParser rules.
const (
	QuarantineParserRULE_expression = 0
	QuarantineParserRULE_quarantine = 1
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
	p.RuleIndex = QuarantineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarantineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Quarantine() IQuarantineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuarantineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuarantineContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(QuarantineParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarantineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarantineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *QuarantineParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, QuarantineParserRULE_expression)

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
		p.Quarantine()
	}
	{
		p.SetState(5)
		p.Match(QuarantineParserEOF)
	}

	return localctx
}

// IQuarantineContext is an interface to support dynamic dispatch.
type IQuarantineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuarantineContext differentiates from other interfaces.
	IsQuarantineContext()
}

type QuarantineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuarantineContext() *QuarantineContext {
	var p = new(QuarantineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = QuarantineParserRULE_quarantine
	return p
}

func (*QuarantineContext) IsQuarantineContext() {}

func NewQuarantineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuarantineContext {
	var p = new(QuarantineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = QuarantineParserRULE_quarantine

	return p
}

func (s *QuarantineContext) GetParser() antlr.Parser { return s.parser }

func (s *QuarantineContext) QUARANTINE() antlr.TerminalNode {
	return s.GetToken(QuarantineParserQUARANTINE, 0)
}

func (s *QuarantineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuarantineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuarantineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarantineListener); ok {
		listenerT.EnterQuarantine(s)
	}
}

func (s *QuarantineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(QuarantineListener); ok {
		listenerT.ExitQuarantine(s)
	}
}

func (p *QuarantineParser) Quarantine() (localctx IQuarantineContext) {
	this := p
	_ = this

	localctx = NewQuarantineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, QuarantineParserRULE_quarantine)

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
		p.Match(QuarantineParserQUARANTINE)
	}

	return localctx
}
