// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669692348781/Family.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Family

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

type FamilyParser struct {
	*antlr.BaseParser
}

var familyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func familyParserInit() {
	staticData := &familyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FAMILY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "family",
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

// FamilyParserInit initializes any static state used to implement FamilyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFamilyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FamilyParserInit() {
	staticData := &familyParserStaticData
	staticData.once.Do(familyParserInit)
}

// NewFamilyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFamilyParser(input antlr.TokenStream) *FamilyParser {
	FamilyParserInit()
	this := new(FamilyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &familyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Family.g4"

	return this
}

// FamilyParser tokens.
const (
	FamilyParserEOF      = antlr.TokenEOF
	FamilyParserFAMILY   = 1
	FamilyParserOWNKEY   = 2
	FamilyParserSPLITKEY = 3
	FamilyParserWS       = 4
)

// FamilyParser rules.
const (
	FamilyParserRULE_expression = 0
	FamilyParserRULE_family     = 1
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
	p.RuleIndex = FamilyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FamilyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Family() IFamilyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFamilyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFamilyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FamilyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FamilyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FamilyParserRULE_expression)

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
		p.Family()
	}
	{
		p.SetState(5)
		p.Match(FamilyParserEOF)
	}

	return localctx
}

// IFamilyContext is an interface to support dynamic dispatch.
type IFamilyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFamilyContext differentiates from other interfaces.
	IsFamilyContext()
}

type FamilyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFamilyContext() *FamilyContext {
	var p = new(FamilyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FamilyParserRULE_family
	return p
}

func (*FamilyContext) IsFamilyContext() {}

func NewFamilyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FamilyContext {
	var p = new(FamilyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FamilyParserRULE_family

	return p
}

func (s *FamilyContext) GetParser() antlr.Parser { return s.parser }

func (s *FamilyContext) FAMILY() antlr.TerminalNode {
	return s.GetToken(FamilyParserFAMILY, 0)
}

func (s *FamilyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FamilyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FamilyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilyListener); ok {
		listenerT.EnterFamily(s)
	}
}

func (s *FamilyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilyListener); ok {
		listenerT.ExitFamily(s)
	}
}

func (p *FamilyParser) Family() (localctx IFamilyContext) {
	this := p
	_ = this

	localctx = NewFamilyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FamilyParserRULE_family)

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
		p.Match(FamilyParserFAMILY)
	}

	return localctx
}
