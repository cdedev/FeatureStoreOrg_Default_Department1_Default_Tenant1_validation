// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193181705/Cluster.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cluster

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

type ClusterParser struct {
	*antlr.BaseParser
}

var clusterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clusterParserInit() {
	staticData := &clusterParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLUSTER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cluster",
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

// ClusterParserInit initializes any static state used to implement ClusterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClusterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClusterParserInit() {
	staticData := &clusterParserStaticData
	staticData.once.Do(clusterParserInit)
}

// NewClusterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClusterParser(input antlr.TokenStream) *ClusterParser {
	ClusterParserInit()
	this := new(ClusterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clusterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cluster.g4"

	return this
}

// ClusterParser tokens.
const (
	ClusterParserEOF      = antlr.TokenEOF
	ClusterParserCLUSTER  = 1
	ClusterParserOWNKEY   = 2
	ClusterParserSPLITKEY = 3
	ClusterParserWS       = 4
)

// ClusterParser rules.
const (
	ClusterParserRULE_expression = 0
	ClusterParserRULE_cluster    = 1
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
	p.RuleIndex = ClusterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClusterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cluster() IClusterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClusterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClusterContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClusterParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClusterListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClusterListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ClusterParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClusterParserRULE_expression)

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
		p.Cluster()
	}
	{
		p.SetState(5)
		p.Match(ClusterParserEOF)
	}

	return localctx
}

// IClusterContext is an interface to support dynamic dispatch.
type IClusterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClusterContext differentiates from other interfaces.
	IsClusterContext()
}

type ClusterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClusterContext() *ClusterContext {
	var p = new(ClusterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClusterParserRULE_cluster
	return p
}

func (*ClusterContext) IsClusterContext() {}

func NewClusterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClusterContext {
	var p = new(ClusterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClusterParserRULE_cluster

	return p
}

func (s *ClusterContext) GetParser() antlr.Parser { return s.parser }

func (s *ClusterContext) CLUSTER() antlr.TerminalNode {
	return s.GetToken(ClusterParserCLUSTER, 0)
}

func (s *ClusterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClusterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClusterListener); ok {
		listenerT.EnterCluster(s)
	}
}

func (s *ClusterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClusterListener); ok {
		listenerT.ExitCluster(s)
	}
}

func (p *ClusterParser) Cluster() (localctx IClusterContext) {
	this := p
	_ = this

	localctx = NewClusterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClusterParserRULE_cluster)

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
		p.Match(ClusterParserCLUSTER)
	}

	return localctx
}
