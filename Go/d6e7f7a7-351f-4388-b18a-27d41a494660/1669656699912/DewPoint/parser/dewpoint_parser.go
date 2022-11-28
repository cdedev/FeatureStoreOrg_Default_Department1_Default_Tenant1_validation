// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656699912/DewPoint.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DewPoint

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

type DewPointParser struct {
	*antlr.BaseParser
}

var dewpointParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dewpointParserInit() {
	staticData := &dewpointParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DEWPOINT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dewpoint",
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

// DewPointParserInit initializes any static state used to implement DewPointParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDewPointParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DewPointParserInit() {
	staticData := &dewpointParserStaticData
	staticData.once.Do(dewpointParserInit)
}

// NewDewPointParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDewPointParser(input antlr.TokenStream) *DewPointParser {
	DewPointParserInit()
	this := new(DewPointParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dewpointParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DewPoint.g4"

	return this
}

// DewPointParser tokens.
const (
	DewPointParserEOF      = antlr.TokenEOF
	DewPointParserDEWPOINT = 1
	DewPointParserOWNKEY   = 2
	DewPointParserSPLITKEY = 3
	DewPointParserWS       = 4
)

// DewPointParser rules.
const (
	DewPointParserRULE_expression = 0
	DewPointParserRULE_dewpoint   = 1
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
	p.RuleIndex = DewPointParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DewPointParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dewpoint() IDewpointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDewpointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDewpointContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DewPointParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DewPointListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DewPointListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DewPointParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DewPointParserRULE_expression)

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
		p.Dewpoint()
	}
	{
		p.SetState(5)
		p.Match(DewPointParserEOF)
	}

	return localctx
}

// IDewpointContext is an interface to support dynamic dispatch.
type IDewpointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDewpointContext differentiates from other interfaces.
	IsDewpointContext()
}

type DewpointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDewpointContext() *DewpointContext {
	var p = new(DewpointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DewPointParserRULE_dewpoint
	return p
}

func (*DewpointContext) IsDewpointContext() {}

func NewDewpointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DewpointContext {
	var p = new(DewpointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DewPointParserRULE_dewpoint

	return p
}

func (s *DewpointContext) GetParser() antlr.Parser { return s.parser }

func (s *DewpointContext) DEWPOINT() antlr.TerminalNode {
	return s.GetToken(DewPointParserDEWPOINT, 0)
}

func (s *DewpointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DewpointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DewpointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DewPointListener); ok {
		listenerT.EnterDewpoint(s)
	}
}

func (s *DewpointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DewPointListener); ok {
		listenerT.ExitDewpoint(s)
	}
}

func (p *DewPointParser) Dewpoint() (localctx IDewpointContext) {
	this := p
	_ = this

	localctx = NewDewpointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DewPointParserRULE_dewpoint)

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
		p.Match(DewPointParserDEWPOINT)
	}

	return localctx
}
