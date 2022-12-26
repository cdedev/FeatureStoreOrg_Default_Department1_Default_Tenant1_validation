// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077422559/JulMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JulMaxtemp

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

type JulMaxtempParser struct {
	*antlr.BaseParser
}

var julmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func julmaxtempParserInit() {
	staticData := &julmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JULMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "julmaxtemp",
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

// JulMaxtempParserInit initializes any static state used to implement JulMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJulMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JulMaxtempParserInit() {
	staticData := &julmaxtempParserStaticData
	staticData.once.Do(julmaxtempParserInit)
}

// NewJulMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJulMaxtempParser(input antlr.TokenStream) *JulMaxtempParser {
	JulMaxtempParserInit()
	this := new(JulMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &julmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JulMaxtemp.g4"

	return this
}

// JulMaxtempParser tokens.
const (
	JulMaxtempParserEOF        = antlr.TokenEOF
	JulMaxtempParserJULMAXTEMP = 1
	JulMaxtempParserOWNKEY     = 2
	JulMaxtempParserSPLITKEY   = 3
	JulMaxtempParserWS         = 4
)

// JulMaxtempParser rules.
const (
	JulMaxtempParserRULE_expression = 0
	JulMaxtempParserRULE_julmaxtemp = 1
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
	p.RuleIndex = JulMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JulMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Julmaxtemp() IJulmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJulmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJulmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JulMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JulMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JulMaxtempParserRULE_expression)

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
		p.Julmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(JulMaxtempParserEOF)
	}

	return localctx
}

// IJulmaxtempContext is an interface to support dynamic dispatch.
type IJulmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJulmaxtempContext differentiates from other interfaces.
	IsJulmaxtempContext()
}

type JulmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJulmaxtempContext() *JulmaxtempContext {
	var p = new(JulmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JulMaxtempParserRULE_julmaxtemp
	return p
}

func (*JulmaxtempContext) IsJulmaxtempContext() {}

func NewJulmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JulmaxtempContext {
	var p = new(JulmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JulMaxtempParserRULE_julmaxtemp

	return p
}

func (s *JulmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *JulmaxtempContext) JULMAXTEMP() antlr.TerminalNode {
	return s.GetToken(JulMaxtempParserJULMAXTEMP, 0)
}

func (s *JulmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JulmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JulmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMaxtempListener); ok {
		listenerT.EnterJulmaxtemp(s)
	}
}

func (s *JulmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JulMaxtempListener); ok {
		listenerT.ExitJulmaxtemp(s)
	}
}

func (p *JulMaxtempParser) Julmaxtemp() (localctx IJulmaxtempContext) {
	this := p
	_ = this

	localctx = NewJulmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JulMaxtempParserRULE_julmaxtemp)

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
		p.Match(JulMaxtempParserJULMAXTEMP)
	}

	return localctx
}
