// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690493750/Sodium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sodium

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

type SodiumParser struct {
	*antlr.BaseParser
}

var sodiumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sodiumParserInit() {
	staticData := &sodiumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SODIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sodium",
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

// SodiumParserInit initializes any static state used to implement SodiumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSodiumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SodiumParserInit() {
	staticData := &sodiumParserStaticData
	staticData.once.Do(sodiumParserInit)
}

// NewSodiumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSodiumParser(input antlr.TokenStream) *SodiumParser {
	SodiumParserInit()
	this := new(SodiumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sodiumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sodium.g4"

	return this
}

// SodiumParser tokens.
const (
	SodiumParserEOF      = antlr.TokenEOF
	SodiumParserSODIUM   = 1
	SodiumParserOWNKEY   = 2
	SodiumParserSPLITKEY = 3
	SodiumParserWS       = 4
)

// SodiumParser rules.
const (
	SodiumParserRULE_expression = 0
	SodiumParserRULE_sodium     = 1
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
	p.RuleIndex = SodiumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SodiumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sodium() ISodiumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISodiumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISodiumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SodiumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SodiumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SodiumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SodiumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SodiumParserRULE_expression)

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
		p.Sodium()
	}
	{
		p.SetState(5)
		p.Match(SodiumParserEOF)
	}

	return localctx
}

// ISodiumContext is an interface to support dynamic dispatch.
type ISodiumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSodiumContext differentiates from other interfaces.
	IsSodiumContext()
}

type SodiumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySodiumContext() *SodiumContext {
	var p = new(SodiumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SodiumParserRULE_sodium
	return p
}

func (*SodiumContext) IsSodiumContext() {}

func NewSodiumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SodiumContext {
	var p = new(SodiumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SodiumParserRULE_sodium

	return p
}

func (s *SodiumContext) GetParser() antlr.Parser { return s.parser }

func (s *SodiumContext) SODIUM() antlr.TerminalNode {
	return s.GetToken(SodiumParserSODIUM, 0)
}

func (s *SodiumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SodiumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SodiumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SodiumListener); ok {
		listenerT.EnterSodium(s)
	}
}

func (s *SodiumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SodiumListener); ok {
		listenerT.ExitSodium(s)
	}
}

func (p *SodiumParser) Sodium() (localctx ISodiumContext) {
	this := p
	_ = this

	localctx = NewSodiumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SodiumParserRULE_sodium)

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
		p.Match(SodiumParserSODIUM)
	}

	return localctx
}
