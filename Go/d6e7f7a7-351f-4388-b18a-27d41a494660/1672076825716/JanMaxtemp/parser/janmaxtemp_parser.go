// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076825716/JanMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JanMaxtemp

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

type JanMaxtempParser struct {
	*antlr.BaseParser
}

var janmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func janmaxtempParserInit() {
	staticData := &janmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JANMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "janmaxtemp",
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

// JanMaxtempParserInit initializes any static state used to implement JanMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJanMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JanMaxtempParserInit() {
	staticData := &janmaxtempParserStaticData
	staticData.once.Do(janmaxtempParserInit)
}

// NewJanMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJanMaxtempParser(input antlr.TokenStream) *JanMaxtempParser {
	JanMaxtempParserInit()
	this := new(JanMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &janmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JanMaxtemp.g4"

	return this
}

// JanMaxtempParser tokens.
const (
	JanMaxtempParserEOF        = antlr.TokenEOF
	JanMaxtempParserJANMAXTEMP = 1
	JanMaxtempParserOWNKEY     = 2
	JanMaxtempParserSPLITKEY   = 3
	JanMaxtempParserWS         = 4
)

// JanMaxtempParser rules.
const (
	JanMaxtempParserRULE_expression = 0
	JanMaxtempParserRULE_janmaxtemp = 1
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
	p.RuleIndex = JanMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JanMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Janmaxtemp() IJanmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJanmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJanmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JanMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JanMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JanMaxtempParserRULE_expression)

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
		p.Janmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(JanMaxtempParserEOF)
	}

	return localctx
}

// IJanmaxtempContext is an interface to support dynamic dispatch.
type IJanmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJanmaxtempContext differentiates from other interfaces.
	IsJanmaxtempContext()
}

type JanmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJanmaxtempContext() *JanmaxtempContext {
	var p = new(JanmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JanMaxtempParserRULE_janmaxtemp
	return p
}

func (*JanmaxtempContext) IsJanmaxtempContext() {}

func NewJanmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JanmaxtempContext {
	var p = new(JanmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JanMaxtempParserRULE_janmaxtemp

	return p
}

func (s *JanmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *JanmaxtempContext) JANMAXTEMP() antlr.TerminalNode {
	return s.GetToken(JanMaxtempParserJANMAXTEMP, 0)
}

func (s *JanmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JanmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JanmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMaxtempListener); ok {
		listenerT.EnterJanmaxtemp(s)
	}
}

func (s *JanmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JanMaxtempListener); ok {
		listenerT.ExitJanmaxtemp(s)
	}
}

func (p *JanMaxtempParser) Janmaxtemp() (localctx IJanmaxtempContext) {
	this := p
	_ = this

	localctx = NewJanmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JanMaxtempParserRULE_janmaxtemp)

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
		p.Match(JanMaxtempParserJANMAXTEMP)
	}

	return localctx
}
