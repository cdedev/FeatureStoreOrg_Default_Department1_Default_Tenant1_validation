// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115646318/PerOfActiveConnections.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PerOfActiveConnections

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

type PerOfActiveConnectionsParser struct {
	*antlr.BaseParser
}

var perofactiveconnectionsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func perofactiveconnectionsParserInit() {
	staticData := &perofactiveconnectionsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PEROFACTIVECONNECTIONS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "perofactiveconnections",
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

// PerOfActiveConnectionsParserInit initializes any static state used to implement PerOfActiveConnectionsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPerOfActiveConnectionsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PerOfActiveConnectionsParserInit() {
	staticData := &perofactiveconnectionsParserStaticData
	staticData.once.Do(perofactiveconnectionsParserInit)
}

// NewPerOfActiveConnectionsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPerOfActiveConnectionsParser(input antlr.TokenStream) *PerOfActiveConnectionsParser {
	PerOfActiveConnectionsParserInit()
	this := new(PerOfActiveConnectionsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &perofactiveconnectionsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PerOfActiveConnections.g4"

	return this
}

// PerOfActiveConnectionsParser tokens.
const (
	PerOfActiveConnectionsParserEOF                    = antlr.TokenEOF
	PerOfActiveConnectionsParserPEROFACTIVECONNECTIONS = 1
	PerOfActiveConnectionsParserOWNKEY                 = 2
	PerOfActiveConnectionsParserSPLITKEY               = 3
	PerOfActiveConnectionsParserWS                     = 4
)

// PerOfActiveConnectionsParser rules.
const (
	PerOfActiveConnectionsParserRULE_expression             = 0
	PerOfActiveConnectionsParserRULE_perofactiveconnections = 1
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
	p.RuleIndex = PerOfActiveConnectionsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PerOfActiveConnectionsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Perofactiveconnections() IPerofactiveconnectionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPerofactiveconnectionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPerofactiveconnectionsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PerOfActiveConnectionsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerOfActiveConnectionsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerOfActiveConnectionsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PerOfActiveConnectionsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PerOfActiveConnectionsParserRULE_expression)

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
		p.Perofactiveconnections()
	}
	{
		p.SetState(5)
		p.Match(PerOfActiveConnectionsParserEOF)
	}

	return localctx
}

// IPerofactiveconnectionsContext is an interface to support dynamic dispatch.
type IPerofactiveconnectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPerofactiveconnectionsContext differentiates from other interfaces.
	IsPerofactiveconnectionsContext()
}

type PerofactiveconnectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPerofactiveconnectionsContext() *PerofactiveconnectionsContext {
	var p = new(PerofactiveconnectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PerOfActiveConnectionsParserRULE_perofactiveconnections
	return p
}

func (*PerofactiveconnectionsContext) IsPerofactiveconnectionsContext() {}

func NewPerofactiveconnectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PerofactiveconnectionsContext {
	var p = new(PerofactiveconnectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PerOfActiveConnectionsParserRULE_perofactiveconnections

	return p
}

func (s *PerofactiveconnectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *PerofactiveconnectionsContext) PEROFACTIVECONNECTIONS() antlr.TerminalNode {
	return s.GetToken(PerOfActiveConnectionsParserPEROFACTIVECONNECTIONS, 0)
}

func (s *PerofactiveconnectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PerofactiveconnectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PerofactiveconnectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerOfActiveConnectionsListener); ok {
		listenerT.EnterPerofactiveconnections(s)
	}
}

func (s *PerofactiveconnectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PerOfActiveConnectionsListener); ok {
		listenerT.ExitPerofactiveconnections(s)
	}
}

func (p *PerOfActiveConnectionsParser) Perofactiveconnections() (localctx IPerofactiveconnectionsContext) {
	this := p
	_ = this

	localctx = NewPerofactiveconnectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PerOfActiveConnectionsParserRULE_perofactiveconnections)

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
		p.Match(PerOfActiveConnectionsParserPEROFACTIVECONNECTIONS)
	}

	return localctx
}
