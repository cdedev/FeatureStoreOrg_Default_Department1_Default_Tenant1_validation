// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671863650901/Taste.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Taste

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

type TasteParser struct {
	*antlr.BaseParser
}

var tasteParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tasteParserInit() {
	staticData := &tasteParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TASTE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "taste",
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

// TasteParserInit initializes any static state used to implement TasteParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTasteParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TasteParserInit() {
	staticData := &tasteParserStaticData
	staticData.once.Do(tasteParserInit)
}

// NewTasteParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTasteParser(input antlr.TokenStream) *TasteParser {
	TasteParserInit()
	this := new(TasteParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tasteParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Taste.g4"

	return this
}

// TasteParser tokens.
const (
	TasteParserEOF      = antlr.TokenEOF
	TasteParserTASTE    = 1
	TasteParserOWNKEY   = 2
	TasteParserSPLITKEY = 3
	TasteParserWS       = 4
)

// TasteParser rules.
const (
	TasteParserRULE_expression = 0
	TasteParserRULE_taste      = 1
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
	p.RuleIndex = TasteParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TasteParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Taste() ITasteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITasteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITasteContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TasteParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TasteListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TasteListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TasteParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TasteParserRULE_expression)

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
		p.Taste()
	}
	{
		p.SetState(5)
		p.Match(TasteParserEOF)
	}

	return localctx
}

// ITasteContext is an interface to support dynamic dispatch.
type ITasteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasteContext differentiates from other interfaces.
	IsTasteContext()
}

type TasteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasteContext() *TasteContext {
	var p = new(TasteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TasteParserRULE_taste
	return p
}

func (*TasteContext) IsTasteContext() {}

func NewTasteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TasteContext {
	var p = new(TasteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TasteParserRULE_taste

	return p
}

func (s *TasteContext) GetParser() antlr.Parser { return s.parser }

func (s *TasteContext) TASTE() antlr.TerminalNode {
	return s.GetToken(TasteParserTASTE, 0)
}

func (s *TasteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TasteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TasteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TasteListener); ok {
		listenerT.EnterTaste(s)
	}
}

func (s *TasteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TasteListener); ok {
		listenerT.ExitTaste(s)
	}
}

func (p *TasteParser) Taste() (localctx ITasteContext) {
	this := p
	_ = this

	localctx = NewTasteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TasteParserRULE_taste)

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
		p.Match(TasteParserTASTE)
	}

	return localctx
}
