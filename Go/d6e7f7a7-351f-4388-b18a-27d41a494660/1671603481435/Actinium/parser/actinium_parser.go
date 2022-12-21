// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603481435/Actinium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Actinium

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

type ActiniumParser struct {
	*antlr.BaseParser
}

var actiniumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func actiniumParserInit() {
	staticData := &actiniumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ACTINIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "actinium",
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

// ActiniumParserInit initializes any static state used to implement ActiniumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewActiniumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ActiniumParserInit() {
	staticData := &actiniumParserStaticData
	staticData.once.Do(actiniumParserInit)
}

// NewActiniumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewActiniumParser(input antlr.TokenStream) *ActiniumParser {
	ActiniumParserInit()
	this := new(ActiniumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &actiniumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Actinium.g4"

	return this
}

// ActiniumParser tokens.
const (
	ActiniumParserEOF      = antlr.TokenEOF
	ActiniumParserACTINIUM = 1
	ActiniumParserOWNKEY   = 2
	ActiniumParserSPLITKEY = 3
	ActiniumParserWS       = 4
)

// ActiniumParser rules.
const (
	ActiniumParserRULE_expression = 0
	ActiniumParserRULE_actinium   = 1
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
	p.RuleIndex = ActiniumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActiniumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Actinium() IActiniumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActiniumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActiniumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ActiniumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ActiniumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ActiniumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ActiniumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ActiniumParserRULE_expression)

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
		p.Actinium()
	}
	{
		p.SetState(5)
		p.Match(ActiniumParserEOF)
	}

	return localctx
}

// IActiniumContext is an interface to support dynamic dispatch.
type IActiniumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsActiniumContext differentiates from other interfaces.
	IsActiniumContext()
}

type ActiniumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActiniumContext() *ActiniumContext {
	var p = new(ActiniumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ActiniumParserRULE_actinium
	return p
}

func (*ActiniumContext) IsActiniumContext() {}

func NewActiniumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActiniumContext {
	var p = new(ActiniumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ActiniumParserRULE_actinium

	return p
}

func (s *ActiniumContext) GetParser() antlr.Parser { return s.parser }

func (s *ActiniumContext) ACTINIUM() antlr.TerminalNode {
	return s.GetToken(ActiniumParserACTINIUM, 0)
}

func (s *ActiniumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActiniumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActiniumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ActiniumListener); ok {
		listenerT.EnterActinium(s)
	}
}

func (s *ActiniumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ActiniumListener); ok {
		listenerT.ExitActinium(s)
	}
}

func (p *ActiniumParser) Actinium() (localctx IActiniumContext) {
	this := p
	_ = this

	localctx = NewActiniumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ActiniumParserRULE_actinium)

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
		p.Match(ActiniumParserACTINIUM)
	}

	return localctx
}
