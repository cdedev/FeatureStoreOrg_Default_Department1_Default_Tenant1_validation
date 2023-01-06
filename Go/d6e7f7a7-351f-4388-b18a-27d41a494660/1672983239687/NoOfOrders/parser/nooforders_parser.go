// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983239687/NoOfOrders.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NoOfOrders

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

type NoOfOrdersParser struct {
	*antlr.BaseParser
}

var noofordersParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func noofordersParserInit() {
	staticData := &noofordersParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NOOFORDERS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "nooforders",
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

// NoOfOrdersParserInit initializes any static state used to implement NoOfOrdersParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNoOfOrdersParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NoOfOrdersParserInit() {
	staticData := &noofordersParserStaticData
	staticData.once.Do(noofordersParserInit)
}

// NewNoOfOrdersParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNoOfOrdersParser(input antlr.TokenStream) *NoOfOrdersParser {
	NoOfOrdersParserInit()
	this := new(NoOfOrdersParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &noofordersParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NoOfOrders.g4"

	return this
}

// NoOfOrdersParser tokens.
const (
	NoOfOrdersParserEOF        = antlr.TokenEOF
	NoOfOrdersParserNOOFORDERS = 1
	NoOfOrdersParserOWNKEY     = 2
	NoOfOrdersParserSPLITKEY   = 3
	NoOfOrdersParserWS         = 4
)

// NoOfOrdersParser rules.
const (
	NoOfOrdersParserRULE_expression = 0
	NoOfOrdersParserRULE_nooforders = 1
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
	p.RuleIndex = NoOfOrdersParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NoOfOrdersParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Nooforders() INoofordersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INoofordersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INoofordersContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NoOfOrdersParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoOfOrdersListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoOfOrdersListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NoOfOrdersParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NoOfOrdersParserRULE_expression)

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
		p.Nooforders()
	}
	{
		p.SetState(5)
		p.Match(NoOfOrdersParserEOF)
	}

	return localctx
}

// INoofordersContext is an interface to support dynamic dispatch.
type INoofordersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoofordersContext differentiates from other interfaces.
	IsNoofordersContext()
}

type NoofordersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoofordersContext() *NoofordersContext {
	var p = new(NoofordersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NoOfOrdersParserRULE_nooforders
	return p
}

func (*NoofordersContext) IsNoofordersContext() {}

func NewNoofordersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoofordersContext {
	var p = new(NoofordersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NoOfOrdersParserRULE_nooforders

	return p
}

func (s *NoofordersContext) GetParser() antlr.Parser { return s.parser }

func (s *NoofordersContext) NOOFORDERS() antlr.TerminalNode {
	return s.GetToken(NoOfOrdersParserNOOFORDERS, 0)
}

func (s *NoofordersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoofordersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NoofordersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoOfOrdersListener); ok {
		listenerT.EnterNooforders(s)
	}
}

func (s *NoofordersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NoOfOrdersListener); ok {
		listenerT.ExitNooforders(s)
	}
}

func (p *NoOfOrdersParser) Nooforders() (localctx INoofordersContext) {
	this := p
	_ = this

	localctx = NewNoofordersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NoOfOrdersParserRULE_nooforders)

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
		p.Match(NoOfOrdersParserNOOFORDERS)
	}

	return localctx
}
