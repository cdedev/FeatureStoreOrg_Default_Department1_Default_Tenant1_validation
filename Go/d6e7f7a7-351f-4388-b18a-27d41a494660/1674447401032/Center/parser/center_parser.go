// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447401032/Center.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Center

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

type CenterParser struct {
	*antlr.BaseParser
}

var centerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func centerParserInit() {
	staticData := &centerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CENTER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "center",
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

// CenterParserInit initializes any static state used to implement CenterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCenterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CenterParserInit() {
	staticData := &centerParserStaticData
	staticData.once.Do(centerParserInit)
}

// NewCenterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCenterParser(input antlr.TokenStream) *CenterParser {
	CenterParserInit()
	this := new(CenterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &centerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Center.g4"

	return this
}

// CenterParser tokens.
const (
	CenterParserEOF      = antlr.TokenEOF
	CenterParserCENTER   = 1
	CenterParserOWNKEY   = 2
	CenterParserSPLITKEY = 3
	CenterParserWS       = 4
)

// CenterParser rules.
const (
	CenterParserRULE_expression = 0
	CenterParserRULE_center     = 1
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
	p.RuleIndex = CenterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CenterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Center() ICenterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICenterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICenterContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CenterParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CenterListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CenterListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CenterParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CenterParserRULE_expression)

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
		p.Center()
	}
	{
		p.SetState(5)
		p.Match(CenterParserEOF)
	}

	return localctx
}

// ICenterContext is an interface to support dynamic dispatch.
type ICenterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCenterContext differentiates from other interfaces.
	IsCenterContext()
}

type CenterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCenterContext() *CenterContext {
	var p = new(CenterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CenterParserRULE_center
	return p
}

func (*CenterContext) IsCenterContext() {}

func NewCenterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CenterContext {
	var p = new(CenterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CenterParserRULE_center

	return p
}

func (s *CenterContext) GetParser() antlr.Parser { return s.parser }

func (s *CenterContext) CENTER() antlr.TerminalNode {
	return s.GetToken(CenterParserCENTER, 0)
}

func (s *CenterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CenterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CenterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CenterListener); ok {
		listenerT.EnterCenter(s)
	}
}

func (s *CenterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CenterListener); ok {
		listenerT.ExitCenter(s)
	}
}

func (p *CenterParser) Center() (localctx ICenterContext) {
	this := p
	_ = this

	localctx = NewCenterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CenterParserRULE_center)

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
		p.Match(CenterParserCENTER)
	}

	return localctx
}
