// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674790494867/Mass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mass

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

type MassParser struct {
	*antlr.BaseParser
}

var massParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func massParserInit() {
	staticData := &massParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MASS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mass",
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

// MassParserInit initializes any static state used to implement MassParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMassParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MassParserInit() {
	staticData := &massParserStaticData
	staticData.once.Do(massParserInit)
}

// NewMassParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMassParser(input antlr.TokenStream) *MassParser {
	MassParserInit()
	this := new(MassParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &massParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mass.g4"

	return this
}

// MassParser tokens.
const (
	MassParserEOF      = antlr.TokenEOF
	MassParserMASS     = 1
	MassParserOWNKEY   = 2
	MassParserSPLITKEY = 3
	MassParserWS       = 4
)

// MassParser rules.
const (
	MassParserRULE_expression = 0
	MassParserRULE_mass       = 1
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
	p.RuleIndex = MassParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MassParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mass() IMassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMassContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MassParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MassListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MassListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MassParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MassParserRULE_expression)

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
		p.Mass()
	}
	{
		p.SetState(5)
		p.Match(MassParserEOF)
	}

	return localctx
}

// IMassContext is an interface to support dynamic dispatch.
type IMassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMassContext differentiates from other interfaces.
	IsMassContext()
}

type MassContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMassContext() *MassContext {
	var p = new(MassContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MassParserRULE_mass
	return p
}

func (*MassContext) IsMassContext() {}

func NewMassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MassContext {
	var p = new(MassContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MassParserRULE_mass

	return p
}

func (s *MassContext) GetParser() antlr.Parser { return s.parser }

func (s *MassContext) MASS() antlr.TerminalNode {
	return s.GetToken(MassParserMASS, 0)
}

func (s *MassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MassListener); ok {
		listenerT.EnterMass(s)
	}
}

func (s *MassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MassListener); ok {
		listenerT.ExitMass(s)
	}
}

func (p *MassParser) Mass() (localctx IMassContext) {
	this := p
	_ = this

	localctx = NewMassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MassParserRULE_mass)

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
		p.Match(MassParserMASS)
	}

	return localctx
}
