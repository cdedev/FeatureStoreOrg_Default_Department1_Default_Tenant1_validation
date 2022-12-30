// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377798987/ThreeG.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ThreeG

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

type ThreeGParser struct {
	*antlr.BaseParser
}

var threegParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func threegParserInit() {
	staticData := &threegParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "THREEG", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "threeg",
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

// ThreeGParserInit initializes any static state used to implement ThreeGParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewThreeGParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ThreeGParserInit() {
	staticData := &threegParserStaticData
	staticData.once.Do(threegParserInit)
}

// NewThreeGParser produces a new parser instance for the optional input antlr.TokenStream.
func NewThreeGParser(input antlr.TokenStream) *ThreeGParser {
	ThreeGParserInit()
	this := new(ThreeGParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &threegParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ThreeG.g4"

	return this
}

// ThreeGParser tokens.
const (
	ThreeGParserEOF      = antlr.TokenEOF
	ThreeGParserTHREEG   = 1
	ThreeGParserOWNKEY   = 2
	ThreeGParserSPLITKEY = 3
	ThreeGParserWS       = 4
)

// ThreeGParser rules.
const (
	ThreeGParserRULE_expression = 0
	ThreeGParserRULE_threeg     = 1
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
	p.RuleIndex = ThreeGParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreeGParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Threeg() IThreegContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IThreegContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IThreegContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ThreeGParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreeGListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreeGListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ThreeGParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ThreeGParserRULE_expression)

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
		p.Threeg()
	}
	{
		p.SetState(5)
		p.Match(ThreeGParserEOF)
	}

	return localctx
}

// IThreegContext is an interface to support dynamic dispatch.
type IThreegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThreegContext differentiates from other interfaces.
	IsThreegContext()
}

type ThreegContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThreegContext() *ThreegContext {
	var p = new(ThreegContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreeGParserRULE_threeg
	return p
}

func (*ThreegContext) IsThreegContext() {}

func NewThreegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThreegContext {
	var p = new(ThreegContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreeGParserRULE_threeg

	return p
}

func (s *ThreegContext) GetParser() antlr.Parser { return s.parser }

func (s *ThreegContext) THREEG() antlr.TerminalNode {
	return s.GetToken(ThreeGParserTHREEG, 0)
}

func (s *ThreegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThreegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreeGListener); ok {
		listenerT.EnterThreeg(s)
	}
}

func (s *ThreegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreeGListener); ok {
		listenerT.ExitThreeg(s)
	}
}

func (p *ThreeGParser) Threeg() (localctx IThreegContext) {
	this := p
	_ = this

	localctx = NewThreegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ThreeGParserRULE_threeg)

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
		p.Match(ThreeGParserTHREEG)
	}

	return localctx
}
