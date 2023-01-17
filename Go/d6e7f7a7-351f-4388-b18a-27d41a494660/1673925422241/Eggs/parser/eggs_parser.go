// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925422241/Eggs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eggs

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

type EggsParser struct {
	*antlr.BaseParser
}

var eggsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func eggsParserInit() {
	staticData := &eggsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EGGS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "eggs",
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

// EggsParserInit initializes any static state used to implement EggsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEggsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EggsParserInit() {
	staticData := &eggsParserStaticData
	staticData.once.Do(eggsParserInit)
}

// NewEggsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEggsParser(input antlr.TokenStream) *EggsParser {
	EggsParserInit()
	this := new(EggsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &eggsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Eggs.g4"

	return this
}

// EggsParser tokens.
const (
	EggsParserEOF      = antlr.TokenEOF
	EggsParserEGGS     = 1
	EggsParserOWNKEY   = 2
	EggsParserSPLITKEY = 3
	EggsParserWS       = 4
)

// EggsParser rules.
const (
	EggsParserRULE_expression = 0
	EggsParserRULE_eggs       = 1
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
	p.RuleIndex = EggsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EggsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Eggs() IEggsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEggsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEggsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EggsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EggsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EggsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EggsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EggsParserRULE_expression)

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
		p.Eggs()
	}
	{
		p.SetState(5)
		p.Match(EggsParserEOF)
	}

	return localctx
}

// IEggsContext is an interface to support dynamic dispatch.
type IEggsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEggsContext differentiates from other interfaces.
	IsEggsContext()
}

type EggsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEggsContext() *EggsContext {
	var p = new(EggsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EggsParserRULE_eggs
	return p
}

func (*EggsContext) IsEggsContext() {}

func NewEggsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EggsContext {
	var p = new(EggsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EggsParserRULE_eggs

	return p
}

func (s *EggsContext) GetParser() antlr.Parser { return s.parser }

func (s *EggsContext) EGGS() antlr.TerminalNode {
	return s.GetToken(EggsParserEGGS, 0)
}

func (s *EggsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EggsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EggsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EggsListener); ok {
		listenerT.EnterEggs(s)
	}
}

func (s *EggsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EggsListener); ok {
		listenerT.ExitEggs(s)
	}
}

func (p *EggsParser) Eggs() (localctx IEggsContext) {
	this := p
	_ = this

	localctx = NewEggsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EggsParserRULE_eggs)

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
		p.Match(EggsParserEGGS)
	}

	return localctx
}
