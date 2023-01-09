// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235280067/Eonal_stability.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eonal_stability

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

type Eonal_stabilityParser struct {
	*antlr.BaseParser
}

var eonal_stabilityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eonal_stabilityParserInit() {
	staticData := &eonal_stabilityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EONAL_STABILITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "eonal_stability",
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

// Eonal_stabilityParserInit initializes any static state used to implement Eonal_stabilityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEonal_stabilityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Eonal_stabilityParserInit() {
	staticData := &eonal_stabilityParserStaticData
	staticData.once.Do(eonal_stabilityParserInit)
}

// NewEonal_stabilityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEonal_stabilityParser(input antlr.TokenStream) *Eonal_stabilityParser {
	Eonal_stabilityParserInit()
	this := new(Eonal_stabilityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &eonal_stabilityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Eonal_stability.g4"

	return this
}

// Eonal_stabilityParser tokens.
const (
	Eonal_stabilityParserEOF             = antlr.TokenEOF
	Eonal_stabilityParserEONAL_STABILITY = 1
	Eonal_stabilityParserOWNKEY          = 2
	Eonal_stabilityParserSPLITKEY        = 3
	Eonal_stabilityParserWS              = 4
)

// Eonal_stabilityParser rules.
const (
	Eonal_stabilityParserRULE_expression      = 0
	Eonal_stabilityParserRULE_eonal_stability = 1
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
	p.RuleIndex = Eonal_stabilityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Eonal_stabilityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Eonal_stability() IEonal_stabilityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEonal_stabilityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEonal_stabilityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Eonal_stabilityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Eonal_stabilityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Eonal_stabilityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Eonal_stabilityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Eonal_stabilityParserRULE_expression)

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
		p.Eonal_stability()
	}
	{
		p.SetState(5)
		p.Match(Eonal_stabilityParserEOF)
	}

	return localctx
}

// IEonal_stabilityContext is an interface to support dynamic dispatch.
type IEonal_stabilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEonal_stabilityContext differentiates from other interfaces.
	IsEonal_stabilityContext()
}

type Eonal_stabilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEonal_stabilityContext() *Eonal_stabilityContext {
	var p = new(Eonal_stabilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Eonal_stabilityParserRULE_eonal_stability
	return p
}

func (*Eonal_stabilityContext) IsEonal_stabilityContext() {}

func NewEonal_stabilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Eonal_stabilityContext {
	var p = new(Eonal_stabilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Eonal_stabilityParserRULE_eonal_stability

	return p
}

func (s *Eonal_stabilityContext) GetParser() antlr.Parser { return s.parser }

func (s *Eonal_stabilityContext) EONAL_STABILITY() antlr.TerminalNode {
	return s.GetToken(Eonal_stabilityParserEONAL_STABILITY, 0)
}

func (s *Eonal_stabilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Eonal_stabilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Eonal_stabilityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Eonal_stabilityListener); ok {
		listenerT.EnterEonal_stability(s)
	}
}

func (s *Eonal_stabilityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Eonal_stabilityListener); ok {
		listenerT.ExitEonal_stability(s)
	}
}

func (p *Eonal_stabilityParser) Eonal_stability() (localctx IEonal_stabilityContext) {
	this := p
	_ = this

	localctx = NewEonal_stabilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Eonal_stabilityParserRULE_eonal_stability)

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
		p.Match(Eonal_stabilityParserEONAL_STABILITY)
	}

	return localctx
}
