// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669698939196/ConvexArea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ConvexArea

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

type ConvexAreaParser struct {
	*antlr.BaseParser
}

var convexareaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func convexareaParserInit() {
	staticData := &convexareaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONVEXAREA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "convexarea",
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

// ConvexAreaParserInit initializes any static state used to implement ConvexAreaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewConvexAreaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ConvexAreaParserInit() {
	staticData := &convexareaParserStaticData
	staticData.once.Do(convexareaParserInit)
}

// NewConvexAreaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewConvexAreaParser(input antlr.TokenStream) *ConvexAreaParser {
	ConvexAreaParserInit()
	this := new(ConvexAreaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &convexareaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ConvexArea.g4"

	return this
}

// ConvexAreaParser tokens.
const (
	ConvexAreaParserEOF        = antlr.TokenEOF
	ConvexAreaParserCONVEXAREA = 1
	ConvexAreaParserOWNKEY     = 2
	ConvexAreaParserSPLITKEY   = 3
	ConvexAreaParserWS         = 4
)

// ConvexAreaParser rules.
const (
	ConvexAreaParserRULE_expression = 0
	ConvexAreaParserRULE_convexarea = 1
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
	p.RuleIndex = ConvexAreaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConvexAreaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Convexarea() IConvexareaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConvexareaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConvexareaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ConvexAreaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConvexAreaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConvexAreaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ConvexAreaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConvexAreaParserRULE_expression)

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
		p.Convexarea()
	}
	{
		p.SetState(5)
		p.Match(ConvexAreaParserEOF)
	}

	return localctx
}

// IConvexareaContext is an interface to support dynamic dispatch.
type IConvexareaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConvexareaContext differentiates from other interfaces.
	IsConvexareaContext()
}

type ConvexareaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConvexareaContext() *ConvexareaContext {
	var p = new(ConvexareaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConvexAreaParserRULE_convexarea
	return p
}

func (*ConvexareaContext) IsConvexareaContext() {}

func NewConvexareaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConvexareaContext {
	var p = new(ConvexareaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConvexAreaParserRULE_convexarea

	return p
}

func (s *ConvexareaContext) GetParser() antlr.Parser { return s.parser }

func (s *ConvexareaContext) CONVEXAREA() antlr.TerminalNode {
	return s.GetToken(ConvexAreaParserCONVEXAREA, 0)
}

func (s *ConvexareaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvexareaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConvexareaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConvexAreaListener); ok {
		listenerT.EnterConvexarea(s)
	}
}

func (s *ConvexareaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConvexAreaListener); ok {
		listenerT.ExitConvexarea(s)
	}
}

func (p *ConvexAreaParser) Convexarea() (localctx IConvexareaContext) {
	this := p
	_ = this

	localctx = NewConvexareaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConvexAreaParserRULE_convexarea)

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
		p.Match(ConvexAreaParserCONVEXAREA)
	}

	return localctx
}
