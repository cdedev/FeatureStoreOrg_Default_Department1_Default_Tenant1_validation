// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778891384/Stream.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stream

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

type StreamParser struct {
	*antlr.BaseParser
}

var streamParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func streamParserInit() {
	staticData := &streamParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STREAM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "stream",
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

// StreamParserInit initializes any static state used to implement StreamParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStreamParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StreamParserInit() {
	staticData := &streamParserStaticData
	staticData.once.Do(streamParserInit)
}

// NewStreamParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStreamParser(input antlr.TokenStream) *StreamParser {
	StreamParserInit()
	this := new(StreamParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &streamParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Stream.g4"

	return this
}

// StreamParser tokens.
const (
	StreamParserEOF      = antlr.TokenEOF
	StreamParserSTREAM   = 1
	StreamParserOWNKEY   = 2
	StreamParserSPLITKEY = 3
	StreamParserWS       = 4
)

// StreamParser rules.
const (
	StreamParserRULE_expression = 0
	StreamParserRULE_stream     = 1
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
	p.RuleIndex = StreamParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StreamParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Stream() IStreamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStreamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStreamContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StreamParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreamListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreamListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StreamParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StreamParserRULE_expression)

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
		p.Stream()
	}
	{
		p.SetState(5)
		p.Match(StreamParserEOF)
	}

	return localctx
}

// IStreamContext is an interface to support dynamic dispatch.
type IStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreamContext differentiates from other interfaces.
	IsStreamContext()
}

type StreamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamContext() *StreamContext {
	var p = new(StreamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StreamParserRULE_stream
	return p
}

func (*StreamContext) IsStreamContext() {}

func NewStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamContext {
	var p = new(StreamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StreamParserRULE_stream

	return p
}

func (s *StreamContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamContext) STREAM() antlr.TerminalNode {
	return s.GetToken(StreamParserSTREAM, 0)
}

func (s *StreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreamListener); ok {
		listenerT.EnterStream(s)
	}
}

func (s *StreamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StreamListener); ok {
		listenerT.ExitStream(s)
	}
}

func (p *StreamParser) Stream() (localctx IStreamContext) {
	this := p
	_ = this

	localctx = NewStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StreamParserRULE_stream)

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
		p.Match(StreamParserSTREAM)
	}

	return localctx
}
