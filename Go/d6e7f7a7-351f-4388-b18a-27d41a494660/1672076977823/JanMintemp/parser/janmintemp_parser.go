// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076977823/JanMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JanMintemp

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

type JanMintempParser struct {
	*antlr.BaseParser
}

var janmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func janmintempParserInit() {
	staticData := &janmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JANMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "janmintemp",
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

// JanMintempParserInit initializes any static state used to implement JanMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJanMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JanMintempParserInit() {
	staticData := &janmintempParserStaticData
	staticData.once.Do(janmintempParserInit)
}

// NewJanMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJanMintempParser(input antlr.TokenStream) *JanMintempParser {
	JanMintempParserInit()
	this := new(JanMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &janmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JanMintemp.g4"

	return this
}

// JanMintempParser tokens.
const (
	JanMintempParserEOF        = antlr.TokenEOF
	JanMintempParserJANMINTEMP = 1
	JanMintempParserOWNKEY     = 2
	JanMintempParserSPLITKEY   = 3
	JanMintempParserWS         = 4
)

// JanMintempParser rules.
const (
	JanMintempParserRULE_expression = 0
	JanMintempParserRULE_janmintemp = 1
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
	p.RuleIndex = JanMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JanMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Janmintemp() IJanmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJanmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJanmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JanMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JanMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JanMintempParserRULE_expression)

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
		p.Janmintemp()
	}
	{
		p.SetState(5)
		p.Match(JanMintempParserEOF)
	}

	return localctx
}

// IJanmintempContext is an interface to support dynamic dispatch.
type IJanmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJanmintempContext differentiates from other interfaces.
	IsJanmintempContext()
}

type JanmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJanmintempContext() *JanmintempContext {
	var p = new(JanmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JanMintempParserRULE_janmintemp
	return p
}

func (*JanmintempContext) IsJanmintempContext() {}

func NewJanmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JanmintempContext {
	var p = new(JanmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JanMintempParserRULE_janmintemp

	return p
}

func (s *JanmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *JanmintempContext) JANMINTEMP() antlr.TerminalNode {
	return s.GetToken(JanMintempParserJANMINTEMP, 0)
}

func (s *JanmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JanmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JanmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMintempListener); ok {
		listenerT.EnterJanmintemp(s)
	}
}

func (s *JanmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMintempListener); ok {
		listenerT.ExitJanmintemp(s)
	}
}

func (p *JanMintempParser) Janmintemp() (localctx IJanmintempContext) {
	this := p
	_ = this

	localctx = NewJanmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JanMintempParserRULE_janmintemp)

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
		p.Match(JanMintempParserJANMINTEMP)
	}

	return localctx
}
