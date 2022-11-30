// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669779000828/MultipleLines.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MultipleLines

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

type MultipleLinesParser struct {
	*antlr.BaseParser
}

var multiplelinesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func multiplelinesParserInit() {
	staticData := &multiplelinesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MULTIPLELINES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "multiplelines",
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

// MultipleLinesParserInit initializes any static state used to implement MultipleLinesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMultipleLinesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MultipleLinesParserInit() {
	staticData := &multiplelinesParserStaticData
	staticData.once.Do(multiplelinesParserInit)
}

// NewMultipleLinesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMultipleLinesParser(input antlr.TokenStream) *MultipleLinesParser {
	MultipleLinesParserInit()
	this := new(MultipleLinesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &multiplelinesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MultipleLines.g4"

	return this
}

// MultipleLinesParser tokens.
const (
	MultipleLinesParserEOF           = antlr.TokenEOF
	MultipleLinesParserMULTIPLELINES = 1
	MultipleLinesParserOWNKEY        = 2
	MultipleLinesParserSPLITKEY      = 3
	MultipleLinesParserWS            = 4
)

// MultipleLinesParser rules.
const (
	MultipleLinesParserRULE_expression    = 0
	MultipleLinesParserRULE_multiplelines = 1
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
	p.RuleIndex = MultipleLinesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MultipleLinesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Multiplelines() IMultiplelinesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplelinesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplelinesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MultipleLinesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MultipleLinesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MultipleLinesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MultipleLinesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MultipleLinesParserRULE_expression)

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
		p.Multiplelines()
	}
	{
		p.SetState(5)
		p.Match(MultipleLinesParserEOF)
	}

	return localctx
}

// IMultiplelinesContext is an interface to support dynamic dispatch.
type IMultiplelinesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplelinesContext differentiates from other interfaces.
	IsMultiplelinesContext()
}

type MultiplelinesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplelinesContext() *MultiplelinesContext {
	var p = new(MultiplelinesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MultipleLinesParserRULE_multiplelines
	return p
}

func (*MultiplelinesContext) IsMultiplelinesContext() {}

func NewMultiplelinesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplelinesContext {
	var p = new(MultiplelinesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MultipleLinesParserRULE_multiplelines

	return p
}

func (s *MultiplelinesContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplelinesContext) MULTIPLELINES() antlr.TerminalNode {
	return s.GetToken(MultipleLinesParserMULTIPLELINES, 0)
}

func (s *MultiplelinesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplelinesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplelinesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MultipleLinesListener); ok {
		listenerT.EnterMultiplelines(s)
	}
}

func (s *MultiplelinesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MultipleLinesListener); ok {
		listenerT.ExitMultiplelines(s)
	}
}

func (p *MultipleLinesParser) Multiplelines() (localctx IMultiplelinesContext) {
	this := p
	_ = this

	localctx = NewMultiplelinesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MultipleLinesParserRULE_multiplelines)

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
		p.Match(MultipleLinesParserMULTIPLELINES)
	}

	return localctx
}
