// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203166456/Terrace.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Terrace

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

type TerraceParser struct {
	*antlr.BaseParser
}

var terraceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func terraceParserInit() {
	staticData := &terraceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TERRACE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "terrace",
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

// TerraceParserInit initializes any static state used to implement TerraceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTerraceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TerraceParserInit() {
	staticData := &terraceParserStaticData
	staticData.once.Do(terraceParserInit)
}

// NewTerraceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTerraceParser(input antlr.TokenStream) *TerraceParser {
	TerraceParserInit()
	this := new(TerraceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &terraceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Terrace.g4"

	return this
}

// TerraceParser tokens.
const (
	TerraceParserEOF      = antlr.TokenEOF
	TerraceParserTERRACE  = 1
	TerraceParserOWNKEY   = 2
	TerraceParserSPLITKEY = 3
	TerraceParserWS       = 4
)

// TerraceParser rules.
const (
	TerraceParserRULE_expression = 0
	TerraceParserRULE_terrace    = 1
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
	p.RuleIndex = TerraceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Terrace() ITerraceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITerraceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITerraceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TerraceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TerraceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TerraceParserRULE_expression)

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
		p.Terrace()
	}
	{
		p.SetState(5)
		p.Match(TerraceParserEOF)
	}

	return localctx
}

// ITerraceContext is an interface to support dynamic dispatch.
type ITerraceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerraceContext differentiates from other interfaces.
	IsTerraceContext()
}

type TerraceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerraceContext() *TerraceContext {
	var p = new(TerraceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TerraceParserRULE_terrace
	return p
}

func (*TerraceContext) IsTerraceContext() {}

func NewTerraceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerraceContext {
	var p = new(TerraceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TerraceParserRULE_terrace

	return p
}

func (s *TerraceContext) GetParser() antlr.Parser { return s.parser }

func (s *TerraceContext) TERRACE() antlr.TerminalNode {
	return s.GetToken(TerraceParserTERRACE, 0)
}

func (s *TerraceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerraceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerraceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraceListener); ok {
		listenerT.EnterTerrace(s)
	}
}

func (s *TerraceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TerraceListener); ok {
		listenerT.ExitTerrace(s)
	}
}

func (p *TerraceParser) Terrace() (localctx ITerraceContext) {
	this := p
	_ = this

	localctx = NewTerraceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TerraceParserRULE_terrace)

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
		p.Match(TerraceParserTERRACE)
	}

	return localctx
}
