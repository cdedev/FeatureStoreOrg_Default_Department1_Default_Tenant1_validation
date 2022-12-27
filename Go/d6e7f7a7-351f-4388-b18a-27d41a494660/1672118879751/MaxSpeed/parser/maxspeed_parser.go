// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118879751/MaxSpeed.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaxSpeed

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

type MaxSpeedParser struct {
	*antlr.BaseParser
}

var maxspeedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func maxspeedParserInit() {
	staticData := &maxspeedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MAXSPEED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "maxspeed",
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

// MaxSpeedParserInit initializes any static state used to implement MaxSpeedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMaxSpeedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MaxSpeedParserInit() {
	staticData := &maxspeedParserStaticData
	staticData.once.Do(maxspeedParserInit)
}

// NewMaxSpeedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMaxSpeedParser(input antlr.TokenStream) *MaxSpeedParser {
	MaxSpeedParserInit()
	this := new(MaxSpeedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &maxspeedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MaxSpeed.g4"

	return this
}

// MaxSpeedParser tokens.
const (
	MaxSpeedParserEOF      = antlr.TokenEOF
	MaxSpeedParserMAXSPEED = 1
	MaxSpeedParserOWNKEY   = 2
	MaxSpeedParserSPLITKEY = 3
	MaxSpeedParserWS       = 4
)

// MaxSpeedParser rules.
const (
	MaxSpeedParserRULE_expression = 0
	MaxSpeedParserRULE_maxspeed   = 1
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
	p.RuleIndex = MaxSpeedParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaxSpeedParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Maxspeed() IMaxspeedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMaxspeedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMaxspeedContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MaxSpeedParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaxSpeedListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaxSpeedListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MaxSpeedParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MaxSpeedParserRULE_expression)

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
		p.Maxspeed()
	}
	{
		p.SetState(5)
		p.Match(MaxSpeedParserEOF)
	}

	return localctx
}

// IMaxspeedContext is an interface to support dynamic dispatch.
type IMaxspeedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMaxspeedContext differentiates from other interfaces.
	IsMaxspeedContext()
}

type MaxspeedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMaxspeedContext() *MaxspeedContext {
	var p = new(MaxspeedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MaxSpeedParserRULE_maxspeed
	return p
}

func (*MaxspeedContext) IsMaxspeedContext() {}

func NewMaxspeedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MaxspeedContext {
	var p = new(MaxspeedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaxSpeedParserRULE_maxspeed

	return p
}

func (s *MaxspeedContext) GetParser() antlr.Parser { return s.parser }

func (s *MaxspeedContext) MAXSPEED() antlr.TerminalNode {
	return s.GetToken(MaxSpeedParserMAXSPEED, 0)
}

func (s *MaxspeedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxspeedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MaxspeedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaxSpeedListener); ok {
		listenerT.EnterMaxspeed(s)
	}
}

func (s *MaxspeedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaxSpeedListener); ok {
		listenerT.ExitMaxspeed(s)
	}
}

func (p *MaxSpeedParser) Maxspeed() (localctx IMaxspeedContext) {
	this := p
	_ = this

	localctx = NewMaxspeedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MaxSpeedParserRULE_maxspeed)

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
		p.Match(MaxSpeedParserMAXSPEED)
	}

	return localctx
}
