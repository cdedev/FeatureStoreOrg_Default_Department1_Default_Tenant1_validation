// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878144865/Dysphasia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dysphasia

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

type DysphasiaParser struct {
	*antlr.BaseParser
}

var dysphasiaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dysphasiaParserInit() {
	staticData := &dysphasiaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DYSPHASIA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dysphasia",
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

// DysphasiaParserInit initializes any static state used to implement DysphasiaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDysphasiaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DysphasiaParserInit() {
	staticData := &dysphasiaParserStaticData
	staticData.once.Do(dysphasiaParserInit)
}

// NewDysphasiaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDysphasiaParser(input antlr.TokenStream) *DysphasiaParser {
	DysphasiaParserInit()
	this := new(DysphasiaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dysphasiaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Dysphasia.g4"

	return this
}

// DysphasiaParser tokens.
const (
	DysphasiaParserEOF       = antlr.TokenEOF
	DysphasiaParserDYSPHASIA = 1
	DysphasiaParserOWNKEY    = 2
	DysphasiaParserSPLITKEY  = 3
	DysphasiaParserWS        = 4
)

// DysphasiaParser rules.
const (
	DysphasiaParserRULE_expression = 0
	DysphasiaParserRULE_dysphasia  = 1
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
	p.RuleIndex = DysphasiaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DysphasiaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dysphasia() IDysphasiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDysphasiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDysphasiaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DysphasiaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysphasiaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysphasiaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DysphasiaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DysphasiaParserRULE_expression)

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
		p.Dysphasia()
	}
	{
		p.SetState(5)
		p.Match(DysphasiaParserEOF)
	}

	return localctx
}

// IDysphasiaContext is an interface to support dynamic dispatch.
type IDysphasiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDysphasiaContext differentiates from other interfaces.
	IsDysphasiaContext()
}

type DysphasiaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDysphasiaContext() *DysphasiaContext {
	var p = new(DysphasiaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DysphasiaParserRULE_dysphasia
	return p
}

func (*DysphasiaContext) IsDysphasiaContext() {}

func NewDysphasiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DysphasiaContext {
	var p = new(DysphasiaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DysphasiaParserRULE_dysphasia

	return p
}

func (s *DysphasiaContext) GetParser() antlr.Parser { return s.parser }

func (s *DysphasiaContext) DYSPHASIA() antlr.TerminalNode {
	return s.GetToken(DysphasiaParserDYSPHASIA, 0)
}

func (s *DysphasiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DysphasiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DysphasiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysphasiaListener); ok {
		listenerT.EnterDysphasia(s)
	}
}

func (s *DysphasiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysphasiaListener); ok {
		listenerT.ExitDysphasia(s)
	}
}

func (p *DysphasiaParser) Dysphasia() (localctx IDysphasiaContext) {
	this := p
	_ = this

	localctx = NewDysphasiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DysphasiaParserRULE_dysphasia)

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
		p.Match(DysphasiaParserDYSPHASIA)
	}

	return localctx
}
