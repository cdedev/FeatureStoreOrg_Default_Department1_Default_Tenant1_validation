// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603693191/Argon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Argon

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

type ArgonParser struct {
	*antlr.BaseParser
}

var argonParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func argonParserInit() {
	staticData := &argonParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ARGON", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "argon",
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

// ArgonParserInit initializes any static state used to implement ArgonParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewArgonParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ArgonParserInit() {
	staticData := &argonParserStaticData
	staticData.once.Do(argonParserInit)
}

// NewArgonParser produces a new parser instance for the optional input antlr.TokenStream.
func NewArgonParser(input antlr.TokenStream) *ArgonParser {
	ArgonParserInit()
	this := new(ArgonParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &argonParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Argon.g4"

	return this
}

// ArgonParser tokens.
const (
	ArgonParserEOF      = antlr.TokenEOF
	ArgonParserARGON    = 1
	ArgonParserOWNKEY   = 2
	ArgonParserSPLITKEY = 3
	ArgonParserWS       = 4
)

// ArgonParser rules.
const (
	ArgonParserRULE_expression = 0
	ArgonParserRULE_argon      = 1
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
	p.RuleIndex = ArgonParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgonParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Argon() IArgonContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgonContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgonContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ArgonParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgonListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgonListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ArgonParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArgonParserRULE_expression)

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
		p.Argon()
	}
	{
		p.SetState(5)
		p.Match(ArgonParserEOF)
	}

	return localctx
}

// IArgonContext is an interface to support dynamic dispatch.
type IArgonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgonContext differentiates from other interfaces.
	IsArgonContext()
}

type ArgonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgonContext() *ArgonContext {
	var p = new(ArgonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArgonParserRULE_argon
	return p
}

func (*ArgonContext) IsArgonContext() {}

func NewArgonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgonContext {
	var p = new(ArgonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArgonParserRULE_argon

	return p
}

func (s *ArgonContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgonContext) ARGON() antlr.TerminalNode {
	return s.GetToken(ArgonParserARGON, 0)
}

func (s *ArgonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgonListener); ok {
		listenerT.EnterArgon(s)
	}
}

func (s *ArgonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArgonListener); ok {
		listenerT.ExitArgon(s)
	}
}

func (p *ArgonParser) Argon() (localctx IArgonContext) {
	this := p
	_ = this

	localctx = NewArgonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArgonParserRULE_argon)

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
		p.Match(ArgonParserARGON)
	}

	return localctx
}
