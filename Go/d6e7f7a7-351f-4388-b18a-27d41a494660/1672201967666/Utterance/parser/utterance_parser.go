// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672201967666/Utterance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Utterance

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

type UtteranceParser struct {
	*antlr.BaseParser
}

var utteranceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func utteranceParserInit() {
	staticData := &utteranceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "UTTERANCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "utterance",
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

// UtteranceParserInit initializes any static state used to implement UtteranceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewUtteranceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func UtteranceParserInit() {
	staticData := &utteranceParserStaticData
	staticData.once.Do(utteranceParserInit)
}

// NewUtteranceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewUtteranceParser(input antlr.TokenStream) *UtteranceParser {
	UtteranceParserInit()
	this := new(UtteranceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &utteranceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Utterance.g4"

	return this
}

// UtteranceParser tokens.
const (
	UtteranceParserEOF       = antlr.TokenEOF
	UtteranceParserUTTERANCE = 1
	UtteranceParserOWNKEY    = 2
	UtteranceParserSPLITKEY  = 3
	UtteranceParserWS        = 4
)

// UtteranceParser rules.
const (
	UtteranceParserRULE_expression = 0
	UtteranceParserRULE_utterance  = 1
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
	p.RuleIndex = UtteranceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UtteranceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Utterance() IUtteranceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUtteranceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUtteranceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(UtteranceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UtteranceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UtteranceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *UtteranceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UtteranceParserRULE_expression)

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
		p.Utterance()
	}
	{
		p.SetState(5)
		p.Match(UtteranceParserEOF)
	}

	return localctx
}

// IUtteranceContext is an interface to support dynamic dispatch.
type IUtteranceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUtteranceContext differentiates from other interfaces.
	IsUtteranceContext()
}

type UtteranceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUtteranceContext() *UtteranceContext {
	var p = new(UtteranceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UtteranceParserRULE_utterance
	return p
}

func (*UtteranceContext) IsUtteranceContext() {}

func NewUtteranceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UtteranceContext {
	var p = new(UtteranceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UtteranceParserRULE_utterance

	return p
}

func (s *UtteranceContext) GetParser() antlr.Parser { return s.parser }

func (s *UtteranceContext) UTTERANCE() antlr.TerminalNode {
	return s.GetToken(UtteranceParserUTTERANCE, 0)
}

func (s *UtteranceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UtteranceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UtteranceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UtteranceListener); ok {
		listenerT.EnterUtterance(s)
	}
}

func (s *UtteranceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UtteranceListener); ok {
		listenerT.ExitUtterance(s)
	}
}

func (p *UtteranceParser) Utterance() (localctx IUtteranceContext) {
	this := p
	_ = this

	localctx = NewUtteranceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UtteranceParserRULE_utterance)

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
		p.Match(UtteranceParserUTTERANCE)
	}

	return localctx
}
