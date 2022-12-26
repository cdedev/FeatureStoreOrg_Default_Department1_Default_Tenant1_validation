// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077464293/JulMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JulMintemp

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

type JulMintempParser struct {
	*antlr.BaseParser
}

var julmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func julmintempParserInit() {
	staticData := &julmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JULMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "julmintemp",
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

// JulMintempParserInit initializes any static state used to implement JulMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJulMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JulMintempParserInit() {
	staticData := &julmintempParserStaticData
	staticData.once.Do(julmintempParserInit)
}

// NewJulMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJulMintempParser(input antlr.TokenStream) *JulMintempParser {
	JulMintempParserInit()
	this := new(JulMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &julmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JulMintemp.g4"

	return this
}

// JulMintempParser tokens.
const (
	JulMintempParserEOF        = antlr.TokenEOF
	JulMintempParserJULMINTEMP = 1
	JulMintempParserOWNKEY     = 2
	JulMintempParserSPLITKEY   = 3
	JulMintempParserWS         = 4
)

// JulMintempParser rules.
const (
	JulMintempParserRULE_expression = 0
	JulMintempParserRULE_julmintemp = 1
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
	p.RuleIndex = JulMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JulMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Julmintemp() IJulmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJulmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJulmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JulMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JulMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JulMintempParserRULE_expression)

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
		p.Julmintemp()
	}
	{
		p.SetState(5)
		p.Match(JulMintempParserEOF)
	}

	return localctx
}

// IJulmintempContext is an interface to support dynamic dispatch.
type IJulmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJulmintempContext differentiates from other interfaces.
	IsJulmintempContext()
}

type JulmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJulmintempContext() *JulmintempContext {
	var p = new(JulmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JulMintempParserRULE_julmintemp
	return p
}

func (*JulmintempContext) IsJulmintempContext() {}

func NewJulmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JulmintempContext {
	var p = new(JulmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JulMintempParserRULE_julmintemp

	return p
}

func (s *JulmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *JulmintempContext) JULMINTEMP() antlr.TerminalNode {
	return s.GetToken(JulMintempParserJULMINTEMP, 0)
}

func (s *JulmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JulmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JulmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMintempListener); ok {
		listenerT.EnterJulmintemp(s)
	}
}

func (s *JulmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMintempListener); ok {
		listenerT.ExitJulmintemp(s)
	}
}

func (p *JulMintempParser) Julmintemp() (localctx IJulmintempContext) {
	this := p
	_ = this

	localctx = NewJulmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JulMintempParserRULE_julmintemp)

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
		p.Match(JulMintempParserJULMINTEMP)
	}

	return localctx
}
