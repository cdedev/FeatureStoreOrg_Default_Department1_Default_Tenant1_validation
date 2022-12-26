// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076891972/FebMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FebMaxtemp

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

type FebMaxtempParser struct {
	*antlr.BaseParser
}

var febmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func febmaxtempParserInit() {
	staticData := &febmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FEBMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "febmaxtemp",
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

// FebMaxtempParserInit initializes any static state used to implement FebMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFebMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FebMaxtempParserInit() {
	staticData := &febmaxtempParserStaticData
	staticData.once.Do(febmaxtempParserInit)
}

// NewFebMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFebMaxtempParser(input antlr.TokenStream) *FebMaxtempParser {
	FebMaxtempParserInit()
	this := new(FebMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &febmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FebMaxtemp.g4"

	return this
}

// FebMaxtempParser tokens.
const (
	FebMaxtempParserEOF        = antlr.TokenEOF
	FebMaxtempParserFEBMAXTEMP = 1
	FebMaxtempParserOWNKEY     = 2
	FebMaxtempParserSPLITKEY   = 3
	FebMaxtempParserWS         = 4
)

// FebMaxtempParser rules.
const (
	FebMaxtempParserRULE_expression = 0
	FebMaxtempParserRULE_febmaxtemp = 1
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
	p.RuleIndex = FebMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FebMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Febmaxtemp() IFebmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFebmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFebmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FebMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FebMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FebMaxtempParserRULE_expression)

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
		p.Febmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(FebMaxtempParserEOF)
	}

	return localctx
}

// IFebmaxtempContext is an interface to support dynamic dispatch.
type IFebmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFebmaxtempContext differentiates from other interfaces.
	IsFebmaxtempContext()
}

type FebmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFebmaxtempContext() *FebmaxtempContext {
	var p = new(FebmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FebMaxtempParserRULE_febmaxtemp
	return p
}

func (*FebmaxtempContext) IsFebmaxtempContext() {}

func NewFebmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FebmaxtempContext {
	var p = new(FebmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FebMaxtempParserRULE_febmaxtemp

	return p
}

func (s *FebmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *FebmaxtempContext) FEBMAXTEMP() antlr.TerminalNode {
	return s.GetToken(FebMaxtempParserFEBMAXTEMP, 0)
}

func (s *FebmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FebmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FebmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMaxtempListener); ok {
		listenerT.EnterFebmaxtemp(s)
	}
}

func (s *FebmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FebMaxtempListener); ok {
		listenerT.ExitFebmaxtemp(s)
	}
}

func (p *FebMaxtempParser) Febmaxtemp() (localctx IFebmaxtempContext) {
	this := p
	_ = this

	localctx = NewFebmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FebMaxtempParserRULE_febmaxtemp)

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
		p.Match(FebMaxtempParserFEBMAXTEMP)
	}

	return localctx
}
