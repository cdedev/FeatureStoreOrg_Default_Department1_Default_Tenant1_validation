// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669625336662/Strength.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Strength

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

type StrengthParser struct {
	*antlr.BaseParser
}

var strengthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func strengthParserInit() {
	staticData := &strengthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STRENGTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "strength",
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

// StrengthParserInit initializes any static state used to implement StrengthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStrengthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StrengthParserInit() {
	staticData := &strengthParserStaticData
	staticData.once.Do(strengthParserInit)
}

// NewStrengthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStrengthParser(input antlr.TokenStream) *StrengthParser {
	StrengthParserInit()
	this := new(StrengthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &strengthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Strength.g4"

	return this
}

// StrengthParser tokens.
const (
	StrengthParserEOF      = antlr.TokenEOF
	StrengthParserSTRENGTH = 1
	StrengthParserOWNKEY   = 2
	StrengthParserSPLITKEY = 3
	StrengthParserWS       = 4
)

// StrengthParser rules.
const (
	StrengthParserRULE_expression = 0
	StrengthParserRULE_strength   = 1
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
	p.RuleIndex = StrengthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrengthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Strength() IStrengthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrengthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrengthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StrengthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrengthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrengthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StrengthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StrengthParserRULE_expression)

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
		p.Strength()
	}
	{
		p.SetState(5)
		p.Match(StrengthParserEOF)
	}

	return localctx
}

// IStrengthContext is an interface to support dynamic dispatch.
type IStrengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrengthContext differentiates from other interfaces.
	IsStrengthContext()
}

type StrengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrengthContext() *StrengthContext {
	var p = new(StrengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StrengthParserRULE_strength
	return p
}

func (*StrengthContext) IsStrengthContext() {}

func NewStrengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrengthContext {
	var p = new(StrengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StrengthParserRULE_strength

	return p
}

func (s *StrengthContext) GetParser() antlr.Parser { return s.parser }

func (s *StrengthContext) STRENGTH() antlr.TerminalNode {
	return s.GetToken(StrengthParserSTRENGTH, 0)
}

func (s *StrengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrengthListener); ok {
		listenerT.EnterStrength(s)
	}
}

func (s *StrengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StrengthListener); ok {
		listenerT.ExitStrength(s)
	}
}

func (p *StrengthParser) Strength() (localctx IStrengthContext) {
	this := p
	_ = this

	localctx = NewStrengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StrengthParserRULE_strength)

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
		p.Match(StrengthParserSTRENGTH)
	}

	return localctx
}
