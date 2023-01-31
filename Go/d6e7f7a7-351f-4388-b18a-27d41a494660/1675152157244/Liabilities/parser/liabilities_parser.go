// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675152157244/Liabilities.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liabilities

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

type LiabilitiesParser struct {
	*antlr.BaseParser
}

var liabilitiesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func liabilitiesParserInit() {
	staticData := &liabilitiesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LIABILITIES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "liabilities",
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

// LiabilitiesParserInit initializes any static state used to implement LiabilitiesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLiabilitiesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LiabilitiesParserInit() {
	staticData := &liabilitiesParserStaticData
	staticData.once.Do(liabilitiesParserInit)
}

// NewLiabilitiesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLiabilitiesParser(input antlr.TokenStream) *LiabilitiesParser {
	LiabilitiesParserInit()
	this := new(LiabilitiesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &liabilitiesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Liabilities.g4"

	return this
}

// LiabilitiesParser tokens.
const (
	LiabilitiesParserEOF         = antlr.TokenEOF
	LiabilitiesParserLIABILITIES = 1
	LiabilitiesParserOWNKEY      = 2
	LiabilitiesParserSPLITKEY    = 3
	LiabilitiesParserWS          = 4
)

// LiabilitiesParser rules.
const (
	LiabilitiesParserRULE_expression  = 0
	LiabilitiesParserRULE_liabilities = 1
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
	p.RuleIndex = LiabilitiesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiabilitiesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Liabilities() ILiabilitiesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiabilitiesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiabilitiesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LiabilitiesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiabilitiesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiabilitiesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LiabilitiesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LiabilitiesParserRULE_expression)

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
		p.Liabilities()
	}
	{
		p.SetState(5)
		p.Match(LiabilitiesParserEOF)
	}

	return localctx
}

// ILiabilitiesContext is an interface to support dynamic dispatch.
type ILiabilitiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiabilitiesContext differentiates from other interfaces.
	IsLiabilitiesContext()
}

type LiabilitiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiabilitiesContext() *LiabilitiesContext {
	var p = new(LiabilitiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LiabilitiesParserRULE_liabilities
	return p
}

func (*LiabilitiesContext) IsLiabilitiesContext() {}

func NewLiabilitiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiabilitiesContext {
	var p = new(LiabilitiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LiabilitiesParserRULE_liabilities

	return p
}

func (s *LiabilitiesContext) GetParser() antlr.Parser { return s.parser }

func (s *LiabilitiesContext) LIABILITIES() antlr.TerminalNode {
	return s.GetToken(LiabilitiesParserLIABILITIES, 0)
}

func (s *LiabilitiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiabilitiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiabilitiesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiabilitiesListener); ok {
		listenerT.EnterLiabilities(s)
	}
}

func (s *LiabilitiesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LiabilitiesListener); ok {
		listenerT.ExitLiabilities(s)
	}
}

func (p *LiabilitiesParser) Liabilities() (localctx ILiabilitiesContext) {
	this := p
	_ = this

	localctx = NewLiabilitiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LiabilitiesParserRULE_liabilities)

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
		p.Match(LiabilitiesParserLIABILITIES)
	}

	return localctx
}
