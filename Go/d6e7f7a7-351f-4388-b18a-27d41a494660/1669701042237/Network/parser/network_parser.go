// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701042237/Network.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Network

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

type NetworkParser struct {
	*antlr.BaseParser
}

var networkParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func networkParserInit() {
	staticData := &networkParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NETWORK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "network",
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

// NetworkParserInit initializes any static state used to implement NetworkParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNetworkParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NetworkParserInit() {
	staticData := &networkParserStaticData
	staticData.once.Do(networkParserInit)
}

// NewNetworkParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNetworkParser(input antlr.TokenStream) *NetworkParser {
	NetworkParserInit()
	this := new(NetworkParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &networkParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Network.g4"

	return this
}

// NetworkParser tokens.
const (
	NetworkParserEOF      = antlr.TokenEOF
	NetworkParserNETWORK  = 1
	NetworkParserOWNKEY   = 2
	NetworkParserSPLITKEY = 3
	NetworkParserWS       = 4
)

// NetworkParser rules.
const (
	NetworkParserRULE_expression = 0
	NetworkParserRULE_network    = 1
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
	p.RuleIndex = NetworkParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NetworkParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Network() INetworkContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INetworkContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INetworkContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NetworkParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NetworkListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NetworkListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NetworkParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NetworkParserRULE_expression)

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
		p.Network()
	}
	{
		p.SetState(5)
		p.Match(NetworkParserEOF)
	}

	return localctx
}

// INetworkContext is an interface to support dynamic dispatch.
type INetworkContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNetworkContext differentiates from other interfaces.
	IsNetworkContext()
}

type NetworkContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNetworkContext() *NetworkContext {
	var p = new(NetworkContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NetworkParserRULE_network
	return p
}

func (*NetworkContext) IsNetworkContext() {}

func NewNetworkContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NetworkContext {
	var p = new(NetworkContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NetworkParserRULE_network

	return p
}

func (s *NetworkContext) GetParser() antlr.Parser { return s.parser }

func (s *NetworkContext) NETWORK() antlr.TerminalNode {
	return s.GetToken(NetworkParserNETWORK, 0)
}

func (s *NetworkContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NetworkContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NetworkContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NetworkListener); ok {
		listenerT.EnterNetwork(s)
	}
}

func (s *NetworkContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NetworkListener); ok {
		listenerT.ExitNetwork(s)
	}
}

func (p *NetworkParser) Network() (localctx INetworkContext) {
	this := p
	_ = this

	localctx = NewNetworkContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NetworkParserRULE_network)

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
		p.Match(NetworkParserNETWORK)
	}

	return localctx
}
