// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205224922/CornePoint.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CornePoint

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

type CornePointParser struct {
	*antlr.BaseParser
}

var cornepointParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cornepointParserInit() {
	staticData := &cornepointParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CORNEPOINT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cornepoint",
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

// CornePointParserInit initializes any static state used to implement CornePointParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCornePointParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CornePointParserInit() {
	staticData := &cornepointParserStaticData
	staticData.once.Do(cornepointParserInit)
}

// NewCornePointParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCornePointParser(input antlr.TokenStream) *CornePointParser {
	CornePointParserInit()
	this := new(CornePointParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cornepointParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CornePoint.g4"

	return this
}

// CornePointParser tokens.
const (
	CornePointParserEOF        = antlr.TokenEOF
	CornePointParserCORNEPOINT = 1
	CornePointParserOWNKEY     = 2
	CornePointParserSPLITKEY   = 3
	CornePointParserWS         = 4
)

// CornePointParser rules.
const (
	CornePointParserRULE_expression = 0
	CornePointParserRULE_cornepoint = 1
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
	p.RuleIndex = CornePointParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CornePointParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cornepoint() ICornepointContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICornepointContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICornepointContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CornePointParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CornePointListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CornePointListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CornePointParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CornePointParserRULE_expression)

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
		p.Cornepoint()
	}
	{
		p.SetState(5)
		p.Match(CornePointParserEOF)
	}

	return localctx
}

// ICornepointContext is an interface to support dynamic dispatch.
type ICornepointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCornepointContext differentiates from other interfaces.
	IsCornepointContext()
}

type CornepointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCornepointContext() *CornepointContext {
	var p = new(CornepointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CornePointParserRULE_cornepoint
	return p
}

func (*CornepointContext) IsCornepointContext() {}

func NewCornepointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CornepointContext {
	var p = new(CornepointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CornePointParserRULE_cornepoint

	return p
}

func (s *CornepointContext) GetParser() antlr.Parser { return s.parser }

func (s *CornepointContext) CORNEPOINT() antlr.TerminalNode {
	return s.GetToken(CornePointParserCORNEPOINT, 0)
}

func (s *CornepointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CornepointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CornepointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CornePointListener); ok {
		listenerT.EnterCornepoint(s)
	}
}

func (s *CornepointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CornePointListener); ok {
		listenerT.ExitCornepoint(s)
	}
}

func (p *CornePointParser) Cornepoint() (localctx ICornepointContext) {
	this := p
	_ = this

	localctx = NewCornepointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CornePointParserRULE_cornepoint)

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
		p.Match(CornePointParserCORNEPOINT)
	}

	return localctx
}
