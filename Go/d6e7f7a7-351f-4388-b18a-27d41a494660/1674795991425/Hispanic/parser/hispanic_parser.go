// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674795991425/Hispanic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hispanic

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

type HispanicParser struct {
	*antlr.BaseParser
}

var hispanicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hispanicParserInit() {
	staticData := &hispanicParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HISPANIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hispanic",
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

// HispanicParserInit initializes any static state used to implement HispanicParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHispanicParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HispanicParserInit() {
	staticData := &hispanicParserStaticData
	staticData.once.Do(hispanicParserInit)
}

// NewHispanicParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHispanicParser(input antlr.TokenStream) *HispanicParser {
	HispanicParserInit()
	this := new(HispanicParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hispanicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Hispanic.g4"

	return this
}

// HispanicParser tokens.
const (
	HispanicParserEOF      = antlr.TokenEOF
	HispanicParserHISPANIC = 1
	HispanicParserOWNKEY   = 2
	HispanicParserSPLITKEY = 3
	HispanicParserWS       = 4
)

// HispanicParser rules.
const (
	HispanicParserRULE_expression = 0
	HispanicParserRULE_hispanic   = 1
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
	p.RuleIndex = HispanicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HispanicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hispanic() IHispanicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHispanicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHispanicContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HispanicParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HispanicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HispanicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HispanicParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HispanicParserRULE_expression)

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
		p.Hispanic()
	}
	{
		p.SetState(5)
		p.Match(HispanicParserEOF)
	}

	return localctx
}

// IHispanicContext is an interface to support dynamic dispatch.
type IHispanicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHispanicContext differentiates from other interfaces.
	IsHispanicContext()
}

type HispanicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHispanicContext() *HispanicContext {
	var p = new(HispanicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HispanicParserRULE_hispanic
	return p
}

func (*HispanicContext) IsHispanicContext() {}

func NewHispanicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HispanicContext {
	var p = new(HispanicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HispanicParserRULE_hispanic

	return p
}

func (s *HispanicContext) GetParser() antlr.Parser { return s.parser }

func (s *HispanicContext) HISPANIC() antlr.TerminalNode {
	return s.GetToken(HispanicParserHISPANIC, 0)
}

func (s *HispanicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HispanicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HispanicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HispanicListener); ok {
		listenerT.EnterHispanic(s)
	}
}

func (s *HispanicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HispanicListener); ok {
		listenerT.ExitHispanic(s)
	}
}

func (p *HispanicParser) Hispanic() (localctx IHispanicContext) {
	this := p
	_ = this

	localctx = NewHispanicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HispanicParserRULE_hispanic)

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
		p.Match(HispanicParserHISPANIC)
	}

	return localctx
}
