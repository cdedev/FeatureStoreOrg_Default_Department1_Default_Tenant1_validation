// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674541115725/Married.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Married

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

type MarriedParser struct {
	*antlr.BaseParser
}

var marriedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func marriedParserInit() {
	staticData := &marriedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MARRIED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "married",
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

// MarriedParserInit initializes any static state used to implement MarriedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMarriedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarriedParserInit() {
	staticData := &marriedParserStaticData
	staticData.once.Do(marriedParserInit)
}

// NewMarriedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMarriedParser(input antlr.TokenStream) *MarriedParser {
	MarriedParserInit()
	this := new(MarriedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &marriedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Married.g4"

	return this
}

// MarriedParser tokens.
const (
	MarriedParserEOF      = antlr.TokenEOF
	MarriedParserMARRIED  = 1
	MarriedParserOWNKEY   = 2
	MarriedParserSPLITKEY = 3
	MarriedParserWS       = 4
)

// MarriedParser rules.
const (
	MarriedParserRULE_expression = 0
	MarriedParserRULE_married    = 1
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
	p.RuleIndex = MarriedParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarriedParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Married() IMarriedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarriedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarriedContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MarriedParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarriedListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarriedListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MarriedParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarriedParserRULE_expression)

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
		p.Married()
	}
	{
		p.SetState(5)
		p.Match(MarriedParserEOF)
	}

	return localctx
}

// IMarriedContext is an interface to support dynamic dispatch.
type IMarriedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMarriedContext differentiates from other interfaces.
	IsMarriedContext()
}

type MarriedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarriedContext() *MarriedContext {
	var p = new(MarriedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MarriedParserRULE_married
	return p
}

func (*MarriedContext) IsMarriedContext() {}

func NewMarriedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarriedContext {
	var p = new(MarriedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarriedParserRULE_married

	return p
}

func (s *MarriedContext) GetParser() antlr.Parser { return s.parser }

func (s *MarriedContext) MARRIED() antlr.TerminalNode {
	return s.GetToken(MarriedParserMARRIED, 0)
}

func (s *MarriedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarriedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarriedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarriedListener); ok {
		listenerT.EnterMarried(s)
	}
}

func (s *MarriedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarriedListener); ok {
		listenerT.ExitMarried(s)
	}
}

func (p *MarriedParser) Married() (localctx IMarriedContext) {
	this := p
	_ = this

	localctx = NewMarriedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarriedParserRULE_married)

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
		p.Match(MarriedParserMARRIED)
	}

	return localctx
}
