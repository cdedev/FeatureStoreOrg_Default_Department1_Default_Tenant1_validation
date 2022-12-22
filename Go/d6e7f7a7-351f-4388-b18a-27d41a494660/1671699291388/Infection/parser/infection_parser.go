// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671699291388/Infection.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Infection

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

type InfectionParser struct {
	*antlr.BaseParser
}

var infectionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func infectionParserInit() {
	staticData := &infectionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "INFECTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "infection",
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

// InfectionParserInit initializes any static state used to implement InfectionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewInfectionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func InfectionParserInit() {
	staticData := &infectionParserStaticData
	staticData.once.Do(infectionParserInit)
}

// NewInfectionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewInfectionParser(input antlr.TokenStream) *InfectionParser {
	InfectionParserInit()
	this := new(InfectionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &infectionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Infection.g4"

	return this
}

// InfectionParser tokens.
const (
	InfectionParserEOF       = antlr.TokenEOF
	InfectionParserINFECTION = 1
	InfectionParserOWNKEY    = 2
	InfectionParserSPLITKEY  = 3
	InfectionParserWS        = 4
)

// InfectionParser rules.
const (
	InfectionParserRULE_expression = 0
	InfectionParserRULE_infection  = 1
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
	p.RuleIndex = InfectionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InfectionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Infection() IInfectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInfectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInfectionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(InfectionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InfectionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InfectionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *InfectionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, InfectionParserRULE_expression)

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
		p.Infection()
	}
	{
		p.SetState(5)
		p.Match(InfectionParserEOF)
	}

	return localctx
}

// IInfectionContext is an interface to support dynamic dispatch.
type IInfectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfectionContext differentiates from other interfaces.
	IsInfectionContext()
}

type InfectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfectionContext() *InfectionContext {
	var p = new(InfectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = InfectionParserRULE_infection
	return p
}

func (*InfectionContext) IsInfectionContext() {}

func NewInfectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfectionContext {
	var p = new(InfectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = InfectionParserRULE_infection

	return p
}

func (s *InfectionContext) GetParser() antlr.Parser { return s.parser }

func (s *InfectionContext) INFECTION() antlr.TerminalNode {
	return s.GetToken(InfectionParserINFECTION, 0)
}

func (s *InfectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InfectionListener); ok {
		listenerT.EnterInfection(s)
	}
}

func (s *InfectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(InfectionListener); ok {
		listenerT.ExitInfection(s)
	}
}

func (p *InfectionParser) Infection() (localctx IInfectionContext) {
	this := p
	_ = this

	localctx = NewInfectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, InfectionParserRULE_infection)

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
		p.Match(InfectionParserINFECTION)
	}

	return localctx
}
