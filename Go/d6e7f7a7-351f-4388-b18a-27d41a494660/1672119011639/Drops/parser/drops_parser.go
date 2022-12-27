// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119011639/Drops.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Drops

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

type DropsParser struct {
	*antlr.BaseParser
}

var dropsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dropsParserInit() {
	staticData := &dropsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DROPS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "drops",
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

// DropsParserInit initializes any static state used to implement DropsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDropsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DropsParserInit() {
	staticData := &dropsParserStaticData
	staticData.once.Do(dropsParserInit)
}

// NewDropsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDropsParser(input antlr.TokenStream) *DropsParser {
	DropsParserInit()
	this := new(DropsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dropsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Drops.g4"

	return this
}

// DropsParser tokens.
const (
	DropsParserEOF      = antlr.TokenEOF
	DropsParserDROPS    = 1
	DropsParserOWNKEY   = 2
	DropsParserSPLITKEY = 3
	DropsParserWS       = 4
)

// DropsParser rules.
const (
	DropsParserRULE_expression = 0
	DropsParserRULE_drops      = 1
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
	p.RuleIndex = DropsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DropsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Drops() IDropsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DropsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DropsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DropsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DropsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DropsParserRULE_expression)

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
		p.Drops()
	}
	{
		p.SetState(5)
		p.Match(DropsParserEOF)
	}

	return localctx
}

// IDropsContext is an interface to support dynamic dispatch.
type IDropsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDropsContext differentiates from other interfaces.
	IsDropsContext()
}

type DropsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropsContext() *DropsContext {
	var p = new(DropsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DropsParserRULE_drops
	return p
}

func (*DropsContext) IsDropsContext() {}

func NewDropsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropsContext {
	var p = new(DropsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DropsParserRULE_drops

	return p
}

func (s *DropsContext) GetParser() antlr.Parser { return s.parser }

func (s *DropsContext) DROPS() antlr.TerminalNode {
	return s.GetToken(DropsParserDROPS, 0)
}

func (s *DropsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DropsListener); ok {
		listenerT.EnterDrops(s)
	}
}

func (s *DropsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DropsListener); ok {
		listenerT.ExitDrops(s)
	}
}

func (p *DropsParser) Drops() (localctx IDropsContext) {
	this := p
	_ = this

	localctx = NewDropsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DropsParserRULE_drops)

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
		p.Match(DropsParserDROPS)
	}

	return localctx
}
