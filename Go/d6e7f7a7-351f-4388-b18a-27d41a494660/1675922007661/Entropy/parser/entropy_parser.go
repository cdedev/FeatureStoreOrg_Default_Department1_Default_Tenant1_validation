// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675922007661/Entropy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entropy

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

type EntropyParser struct {
	*antlr.BaseParser
}

var entropyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func entropyParserInit() {
	staticData := &entropyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ENTROPY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "entropy",
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

// EntropyParserInit initializes any static state used to implement EntropyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEntropyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EntropyParserInit() {
	staticData := &entropyParserStaticData
	staticData.once.Do(entropyParserInit)
}

// NewEntropyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEntropyParser(input antlr.TokenStream) *EntropyParser {
	EntropyParserInit()
	this := new(EntropyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &entropyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Entropy.g4"

	return this
}

// EntropyParser tokens.
const (
	EntropyParserEOF      = antlr.TokenEOF
	EntropyParserENTROPY  = 1
	EntropyParserOWNKEY   = 2
	EntropyParserSPLITKEY = 3
	EntropyParserWS       = 4
)

// EntropyParser rules.
const (
	EntropyParserRULE_expression = 0
	EntropyParserRULE_entropy    = 1
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
	p.RuleIndex = EntropyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EntropyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Entropy() IEntropyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntropyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntropyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EntropyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntropyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntropyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EntropyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EntropyParserRULE_expression)

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
		p.Entropy()
	}
	{
		p.SetState(5)
		p.Match(EntropyParserEOF)
	}

	return localctx
}

// IEntropyContext is an interface to support dynamic dispatch.
type IEntropyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntropyContext differentiates from other interfaces.
	IsEntropyContext()
}

type EntropyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntropyContext() *EntropyContext {
	var p = new(EntropyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EntropyParserRULE_entropy
	return p
}

func (*EntropyContext) IsEntropyContext() {}

func NewEntropyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntropyContext {
	var p = new(EntropyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EntropyParserRULE_entropy

	return p
}

func (s *EntropyContext) GetParser() antlr.Parser { return s.parser }

func (s *EntropyContext) ENTROPY() antlr.TerminalNode {
	return s.GetToken(EntropyParserENTROPY, 0)
}

func (s *EntropyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntropyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntropyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntropyListener); ok {
		listenerT.EnterEntropy(s)
	}
}

func (s *EntropyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntropyListener); ok {
		listenerT.ExitEntropy(s)
	}
}

func (p *EntropyParser) Entropy() (localctx IEntropyContext) {
	this := p
	_ = this

	localctx = NewEntropyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EntropyParserRULE_entropy)

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
		p.Match(EntropyParserENTROPY)
	}

	return localctx
}
