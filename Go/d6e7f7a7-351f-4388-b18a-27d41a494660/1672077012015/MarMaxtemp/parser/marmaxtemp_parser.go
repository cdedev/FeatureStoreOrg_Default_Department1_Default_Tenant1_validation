// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077012015/MarMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarMaxtemp

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

type MarMaxtempParser struct {
	*antlr.BaseParser
}

var marmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func marmaxtempParserInit() {
	staticData := &marmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MARMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "marmaxtemp",
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

// MarMaxtempParserInit initializes any static state used to implement MarMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMarMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MarMaxtempParserInit() {
	staticData := &marmaxtempParserStaticData
	staticData.once.Do(marmaxtempParserInit)
}

// NewMarMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMarMaxtempParser(input antlr.TokenStream) *MarMaxtempParser {
	MarMaxtempParserInit()
	this := new(MarMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &marmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MarMaxtemp.g4"

	return this
}

// MarMaxtempParser tokens.
const (
	MarMaxtempParserEOF        = antlr.TokenEOF
	MarMaxtempParserMARMAXTEMP = 1
	MarMaxtempParserOWNKEY     = 2
	MarMaxtempParserSPLITKEY   = 3
	MarMaxtempParserWS         = 4
)

// MarMaxtempParser rules.
const (
	MarMaxtempParserRULE_expression = 0
	MarMaxtempParserRULE_marmaxtemp = 1
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
	p.RuleIndex = MarMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Marmaxtemp() IMarmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMarmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MarMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MarMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MarMaxtempParserRULE_expression)

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
		p.Marmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(MarMaxtempParserEOF)
	}

	return localctx
}

// IMarmaxtempContext is an interface to support dynamic dispatch.
type IMarmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMarmaxtempContext differentiates from other interfaces.
	IsMarmaxtempContext()
}

type MarmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarmaxtempContext() *MarmaxtempContext {
	var p = new(MarmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MarMaxtempParserRULE_marmaxtemp
	return p
}

func (*MarmaxtempContext) IsMarmaxtempContext() {}

func NewMarmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarmaxtempContext {
	var p = new(MarmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MarMaxtempParserRULE_marmaxtemp

	return p
}

func (s *MarmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *MarmaxtempContext) MARMAXTEMP() antlr.TerminalNode {
	return s.GetToken(MarMaxtempParserMARMAXTEMP, 0)
}

func (s *MarmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarMaxtempListener); ok {
		listenerT.EnterMarmaxtemp(s)
	}
}

func (s *MarmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MarMaxtempListener); ok {
		listenerT.ExitMarmaxtemp(s)
	}
}

func (p *MarMaxtempParser) Marmaxtemp() (localctx IMarmaxtempContext) {
	this := p
	_ = this

	localctx = NewMarmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MarMaxtempParserRULE_marmaxtemp)

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
		p.Match(MarMaxtempParserMARMAXTEMP)
	}

	return localctx
}
