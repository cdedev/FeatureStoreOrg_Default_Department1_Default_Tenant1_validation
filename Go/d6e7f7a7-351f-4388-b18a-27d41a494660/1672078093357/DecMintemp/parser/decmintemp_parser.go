// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078093357/DecMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DecMintemp

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

type DecMintempParser struct {
	*antlr.BaseParser
}

var decmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func decmintempParserInit() {
	staticData := &decmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DECMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "decmintemp",
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

// DecMintempParserInit initializes any static state used to implement DecMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDecMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DecMintempParserInit() {
	staticData := &decmintempParserStaticData
	staticData.once.Do(decmintempParserInit)
}

// NewDecMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDecMintempParser(input antlr.TokenStream) *DecMintempParser {
	DecMintempParserInit()
	this := new(DecMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &decmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DecMintemp.g4"

	return this
}

// DecMintempParser tokens.
const (
	DecMintempParserEOF        = antlr.TokenEOF
	DecMintempParserDECMINTEMP = 1
	DecMintempParserOWNKEY     = 2
	DecMintempParserSPLITKEY   = 3
	DecMintempParserWS         = 4
)

// DecMintempParser rules.
const (
	DecMintempParserRULE_expression = 0
	DecMintempParserRULE_decmintemp = 1
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
	p.RuleIndex = DecMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DecMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Decmintemp() IDecmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DecMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DecMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DecMintempParserRULE_expression)

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
		p.Decmintemp()
	}
	{
		p.SetState(5)
		p.Match(DecMintempParserEOF)
	}

	return localctx
}

// IDecmintempContext is an interface to support dynamic dispatch.
type IDecmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecmintempContext differentiates from other interfaces.
	IsDecmintempContext()
}

type DecmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecmintempContext() *DecmintempContext {
	var p = new(DecmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DecMintempParserRULE_decmintemp
	return p
}

func (*DecmintempContext) IsDecmintempContext() {}

func NewDecmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecmintempContext {
	var p = new(DecmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DecMintempParserRULE_decmintemp

	return p
}

func (s *DecmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *DecmintempContext) DECMINTEMP() antlr.TerminalNode {
	return s.GetToken(DecMintempParserDECMINTEMP, 0)
}

func (s *DecmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMintempListener); ok {
		listenerT.EnterDecmintemp(s)
	}
}

func (s *DecmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMintempListener); ok {
		listenerT.ExitDecmintemp(s)
	}
}

func (p *DecMintempParser) Decmintemp() (localctx IDecmintempContext) {
	this := p
	_ = this

	localctx = NewDecmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DecMintempParserRULE_decmintemp)

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
		p.Match(DecMintempParserDECMINTEMP)
	}

	return localctx
}
