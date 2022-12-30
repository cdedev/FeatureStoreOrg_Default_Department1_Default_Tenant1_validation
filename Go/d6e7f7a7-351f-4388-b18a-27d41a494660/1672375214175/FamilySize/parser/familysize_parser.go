// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672375214175/FamilySize.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FamilySize

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

type FamilySizeParser struct {
	*antlr.BaseParser
}

var familysizeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func familysizeParserInit() {
	staticData := &familysizeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FAMILYSIZE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "familysize",
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

// FamilySizeParserInit initializes any static state used to implement FamilySizeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFamilySizeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FamilySizeParserInit() {
	staticData := &familysizeParserStaticData
	staticData.once.Do(familysizeParserInit)
}

// NewFamilySizeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFamilySizeParser(input antlr.TokenStream) *FamilySizeParser {
	FamilySizeParserInit()
	this := new(FamilySizeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &familysizeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FamilySize.g4"

	return this
}

// FamilySizeParser tokens.
const (
	FamilySizeParserEOF        = antlr.TokenEOF
	FamilySizeParserFAMILYSIZE = 1
	FamilySizeParserOWNKEY     = 2
	FamilySizeParserSPLITKEY   = 3
	FamilySizeParserWS         = 4
)

// FamilySizeParser rules.
const (
	FamilySizeParserRULE_expression = 0
	FamilySizeParserRULE_familysize = 1
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
	p.RuleIndex = FamilySizeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FamilySizeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Familysize() IFamilysizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFamilysizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFamilysizeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FamilySizeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilySizeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilySizeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FamilySizeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FamilySizeParserRULE_expression)

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
		p.Familysize()
	}
	{
		p.SetState(5)
		p.Match(FamilySizeParserEOF)
	}

	return localctx
}

// IFamilysizeContext is an interface to support dynamic dispatch.
type IFamilysizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFamilysizeContext differentiates from other interfaces.
	IsFamilysizeContext()
}

type FamilysizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFamilysizeContext() *FamilysizeContext {
	var p = new(FamilysizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FamilySizeParserRULE_familysize
	return p
}

func (*FamilysizeContext) IsFamilysizeContext() {}

func NewFamilysizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FamilysizeContext {
	var p = new(FamilysizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FamilySizeParserRULE_familysize

	return p
}

func (s *FamilysizeContext) GetParser() antlr.Parser { return s.parser }

func (s *FamilysizeContext) FAMILYSIZE() antlr.TerminalNode {
	return s.GetToken(FamilySizeParserFAMILYSIZE, 0)
}

func (s *FamilysizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FamilysizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FamilysizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilySizeListener); ok {
		listenerT.EnterFamilysize(s)
	}
}

func (s *FamilysizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FamilySizeListener); ok {
		listenerT.ExitFamilysize(s)
	}
}

func (p *FamilySizeParser) Familysize() (localctx IFamilysizeContext) {
	this := p
	_ = this

	localctx = NewFamilysizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FamilySizeParserRULE_familysize)

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
		p.Match(FamilySizeParserFAMILYSIZE)
	}

	return localctx
}
