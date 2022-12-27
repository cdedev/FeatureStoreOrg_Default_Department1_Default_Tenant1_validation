// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115605743/NumConnections.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NumConnections

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

type NumConnectionsParser struct {
	*antlr.BaseParser
}

var numconnectionsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func numconnectionsParserInit() {
	staticData := &numconnectionsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NUMCONNECTIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "numconnections",
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

// NumConnectionsParserInit initializes any static state used to implement NumConnectionsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNumConnectionsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NumConnectionsParserInit() {
	staticData := &numconnectionsParserStaticData
	staticData.once.Do(numconnectionsParserInit)
}

// NewNumConnectionsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNumConnectionsParser(input antlr.TokenStream) *NumConnectionsParser {
	NumConnectionsParserInit()
	this := new(NumConnectionsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &numconnectionsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NumConnections.g4"

	return this
}

// NumConnectionsParser tokens.
const (
	NumConnectionsParserEOF            = antlr.TokenEOF
	NumConnectionsParserNUMCONNECTIONS = 1
	NumConnectionsParserOWNKEY         = 2
	NumConnectionsParserSPLITKEY       = 3
	NumConnectionsParserWS             = 4
)

// NumConnectionsParser rules.
const (
	NumConnectionsParserRULE_expression     = 0
	NumConnectionsParserRULE_numconnections = 1
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
	p.RuleIndex = NumConnectionsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumConnectionsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Numconnections() INumconnectionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumconnectionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumconnectionsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NumConnectionsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumConnectionsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumConnectionsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NumConnectionsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NumConnectionsParserRULE_expression)

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
		p.Numconnections()
	}
	{
		p.SetState(5)
		p.Match(NumConnectionsParserEOF)
	}

	return localctx
}

// INumconnectionsContext is an interface to support dynamic dispatch.
type INumconnectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumconnectionsContext differentiates from other interfaces.
	IsNumconnectionsContext()
}

type NumconnectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumconnectionsContext() *NumconnectionsContext {
	var p = new(NumconnectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NumConnectionsParserRULE_numconnections
	return p
}

func (*NumconnectionsContext) IsNumconnectionsContext() {}

func NewNumconnectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumconnectionsContext {
	var p = new(NumconnectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NumConnectionsParserRULE_numconnections

	return p
}

func (s *NumconnectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *NumconnectionsContext) NUMCONNECTIONS() antlr.TerminalNode {
	return s.GetToken(NumConnectionsParserNUMCONNECTIONS, 0)
}

func (s *NumconnectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumconnectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumconnectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumConnectionsListener); ok {
		listenerT.EnterNumconnections(s)
	}
}

func (s *NumconnectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumConnectionsListener); ok {
		listenerT.ExitNumconnections(s)
	}
}

func (p *NumConnectionsParser) Numconnections() (localctx INumconnectionsContext) {
	this := p
	_ = this

	localctx = NewNumconnectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NumConnectionsParserRULE_numconnections)

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
		p.Match(NumConnectionsParserNUMCONNECTIONS)
	}

	return localctx
}
