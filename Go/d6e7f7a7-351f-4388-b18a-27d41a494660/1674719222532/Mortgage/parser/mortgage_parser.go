// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719222532/Mortgage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mortgage

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

type MortgageParser struct {
	*antlr.BaseParser
}

var mortgageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mortgageParserInit() {
	staticData := &mortgageParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MORTGAGE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mortgage",
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

// MortgageParserInit initializes any static state used to implement MortgageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMortgageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MortgageParserInit() {
	staticData := &mortgageParserStaticData
	staticData.once.Do(mortgageParserInit)
}

// NewMortgageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMortgageParser(input antlr.TokenStream) *MortgageParser {
	MortgageParserInit()
	this := new(MortgageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mortgageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mortgage.g4"

	return this
}

// MortgageParser tokens.
const (
	MortgageParserEOF      = antlr.TokenEOF
	MortgageParserMORTGAGE = 1
	MortgageParserOWNKEY   = 2
	MortgageParserSPLITKEY = 3
	MortgageParserWS       = 4
)

// MortgageParser rules.
const (
	MortgageParserRULE_expression = 0
	MortgageParserRULE_mortgage   = 1
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
	p.RuleIndex = MortgageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MortgageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mortgage() IMortgageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMortgageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMortgageContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MortgageParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortgageListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortgageListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MortgageParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MortgageParserRULE_expression)

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
		p.Mortgage()
	}
	{
		p.SetState(5)
		p.Match(MortgageParserEOF)
	}

	return localctx
}

// IMortgageContext is an interface to support dynamic dispatch.
type IMortgageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMortgageContext differentiates from other interfaces.
	IsMortgageContext()
}

type MortgageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMortgageContext() *MortgageContext {
	var p = new(MortgageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MortgageParserRULE_mortgage
	return p
}

func (*MortgageContext) IsMortgageContext() {}

func NewMortgageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MortgageContext {
	var p = new(MortgageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MortgageParserRULE_mortgage

	return p
}

func (s *MortgageContext) GetParser() antlr.Parser { return s.parser }

func (s *MortgageContext) MORTGAGE() antlr.TerminalNode {
	return s.GetToken(MortgageParserMORTGAGE, 0)
}

func (s *MortgageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MortgageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MortgageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortgageListener); ok {
		listenerT.EnterMortgage(s)
	}
}

func (s *MortgageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MortgageListener); ok {
		listenerT.ExitMortgage(s)
	}
}

func (p *MortgageParser) Mortgage() (localctx IMortgageContext) {
	this := p
	_ = this

	localctx = NewMortgageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MortgageParserRULE_mortgage)

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
		p.Match(MortgageParserMORTGAGE)
	}

	return localctx
}
