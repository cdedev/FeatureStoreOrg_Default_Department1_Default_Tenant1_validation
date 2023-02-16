// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522774644/Publisher.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Publisher

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

type PublisherParser struct {
	*antlr.BaseParser
}

var publisherParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func publisherParserInit() {
	staticData := &publisherParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PUBLISHER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "publisher",
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

// PublisherParserInit initializes any static state used to implement PublisherParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPublisherParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PublisherParserInit() {
	staticData := &publisherParserStaticData
	staticData.once.Do(publisherParserInit)
}

// NewPublisherParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPublisherParser(input antlr.TokenStream) *PublisherParser {
	PublisherParserInit()
	this := new(PublisherParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &publisherParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Publisher.g4"

	return this
}

// PublisherParser tokens.
const (
	PublisherParserEOF       = antlr.TokenEOF
	PublisherParserPUBLISHER = 1
	PublisherParserOWNKEY    = 2
	PublisherParserSPLITKEY  = 3
	PublisherParserWS        = 4
)

// PublisherParser rules.
const (
	PublisherParserRULE_expression = 0
	PublisherParserRULE_publisher  = 1
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
	p.RuleIndex = PublisherParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PublisherParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Publisher() IPublisherContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPublisherContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPublisherContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PublisherParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PublisherListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PublisherListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PublisherParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PublisherParserRULE_expression)

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
		p.Publisher()
	}
	{
		p.SetState(5)
		p.Match(PublisherParserEOF)
	}

	return localctx
}

// IPublisherContext is an interface to support dynamic dispatch.
type IPublisherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPublisherContext differentiates from other interfaces.
	IsPublisherContext()
}

type PublisherContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPublisherContext() *PublisherContext {
	var p = new(PublisherContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PublisherParserRULE_publisher
	return p
}

func (*PublisherContext) IsPublisherContext() {}

func NewPublisherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PublisherContext {
	var p = new(PublisherContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PublisherParserRULE_publisher

	return p
}

func (s *PublisherContext) GetParser() antlr.Parser { return s.parser }

func (s *PublisherContext) PUBLISHER() antlr.TerminalNode {
	return s.GetToken(PublisherParserPUBLISHER, 0)
}

func (s *PublisherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PublisherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PublisherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PublisherListener); ok {
		listenerT.EnterPublisher(s)
	}
}

func (s *PublisherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PublisherListener); ok {
		listenerT.ExitPublisher(s)
	}
}

func (p *PublisherParser) Publisher() (localctx IPublisherContext) {
	this := p
	_ = this

	localctx = NewPublisherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PublisherParserRULE_publisher)

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
		p.Match(PublisherParserPUBLISHER)
	}

	return localctx
}
