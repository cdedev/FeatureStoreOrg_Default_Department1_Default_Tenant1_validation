// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669627420501/Open.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Open

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

type OpenParser struct {
	*antlr.BaseParser
}

var openParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func openParserInit() {
	staticData := &openParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OPEN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "open",
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

// OpenParserInit initializes any static state used to implement OpenParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOpenParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenParserInit() {
	staticData := &openParserStaticData
	staticData.once.Do(openParserInit)
}

// NewOpenParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOpenParser(input antlr.TokenStream) *OpenParser {
	OpenParserInit()
	this := new(OpenParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &openParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Open.g4"

	return this
}

// OpenParser tokens.
const (
	OpenParserEOF      = antlr.TokenEOF
	OpenParserOPEN     = 1
	OpenParserOWNKEY   = 2
	OpenParserSPLITKEY = 3
	OpenParserWS       = 4
)

// OpenParser rules.
const (
	OpenParserRULE_expression = 0
	OpenParserRULE_open       = 1
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
	p.RuleIndex = OpenParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Open() IOpenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpenContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OpenParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OpenParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OpenParserRULE_expression)

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
		p.Open()
	}
	{
		p.SetState(5)
		p.Match(OpenParserEOF)
	}

	return localctx
}

// IOpenContext is an interface to support dynamic dispatch.
type IOpenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenContext differentiates from other interfaces.
	IsOpenContext()
}

type OpenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenContext() *OpenContext {
	var p = new(OpenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenParserRULE_open
	return p
}

func (*OpenContext) IsOpenContext() {}

func NewOpenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenContext {
	var p = new(OpenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenParserRULE_open

	return p
}

func (s *OpenContext) GetParser() antlr.Parser { return s.parser }

func (s *OpenContext) OPEN() antlr.TerminalNode {
	return s.GetToken(OpenParserOPEN, 0)
}

func (s *OpenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenListener); ok {
		listenerT.EnterOpen(s)
	}
}

func (s *OpenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenListener); ok {
		listenerT.ExitOpen(s)
	}
}

func (p *OpenParser) Open() (localctx IOpenContext) {
	this := p
	_ = this

	localctx = NewOpenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpenParserRULE_open)

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
		p.Match(OpenParserOPEN)
	}

	return localctx
}
