// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284909159/Nudity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nudity

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

type NudityParser struct {
	*antlr.BaseParser
}

var nudityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nudityParserInit() {
	staticData := &nudityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NUDITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "nudity",
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

// NudityParserInit initializes any static state used to implement NudityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNudityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NudityParserInit() {
	staticData := &nudityParserStaticData
	staticData.once.Do(nudityParserInit)
}

// NewNudityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNudityParser(input antlr.TokenStream) *NudityParser {
	NudityParserInit()
	this := new(NudityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nudityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Nudity.g4"

	return this
}

// NudityParser tokens.
const (
	NudityParserEOF      = antlr.TokenEOF
	NudityParserNUDITY   = 1
	NudityParserOWNKEY   = 2
	NudityParserSPLITKEY = 3
	NudityParserWS       = 4
)

// NudityParser rules.
const (
	NudityParserRULE_expression = 0
	NudityParserRULE_nudity     = 1
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
	p.RuleIndex = NudityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NudityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Nudity() INudityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INudityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INudityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NudityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NudityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NudityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NudityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NudityParserRULE_expression)

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
		p.Nudity()
	}
	{
		p.SetState(5)
		p.Match(NudityParserEOF)
	}

	return localctx
}

// INudityContext is an interface to support dynamic dispatch.
type INudityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNudityContext differentiates from other interfaces.
	IsNudityContext()
}

type NudityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNudityContext() *NudityContext {
	var p = new(NudityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NudityParserRULE_nudity
	return p
}

func (*NudityContext) IsNudityContext() {}

func NewNudityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NudityContext {
	var p = new(NudityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NudityParserRULE_nudity

	return p
}

func (s *NudityContext) GetParser() antlr.Parser { return s.parser }

func (s *NudityContext) NUDITY() antlr.TerminalNode {
	return s.GetToken(NudityParserNUDITY, 0)
}

func (s *NudityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NudityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NudityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NudityListener); ok {
		listenerT.EnterNudity(s)
	}
}

func (s *NudityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NudityListener); ok {
		listenerT.ExitNudity(s)
	}
}

func (p *NudityParser) Nudity() (localctx INudityContext) {
	this := p
	_ = this

	localctx = NewNudityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NudityParserRULE_nudity)

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
		p.Match(NudityParserNUDITY)
	}

	return localctx
}
