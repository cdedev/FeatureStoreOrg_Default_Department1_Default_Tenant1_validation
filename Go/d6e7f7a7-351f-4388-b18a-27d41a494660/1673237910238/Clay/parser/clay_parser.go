// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237910238/Clay.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clay

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

type ClayParser struct {
	*antlr.BaseParser
}

var clayParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func clayParserInit() {
	staticData := &clayParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLAY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "clay",
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

// ClayParserInit initializes any static state used to implement ClayParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewClayParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ClayParserInit() {
	staticData := &clayParserStaticData
	staticData.once.Do(clayParserInit)
}

// NewClayParser produces a new parser instance for the optional input antlr.TokenStream.
func NewClayParser(input antlr.TokenStream) *ClayParser {
	ClayParserInit()
	this := new(ClayParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &clayParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Clay.g4"

	return this
}

// ClayParser tokens.
const (
	ClayParserEOF      = antlr.TokenEOF
	ClayParserCLAY     = 1
	ClayParserOWNKEY   = 2
	ClayParserSPLITKEY = 3
	ClayParserWS       = 4
)

// ClayParser rules.
const (
	ClayParserRULE_expression = 0
	ClayParserRULE_clay       = 1
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
	p.RuleIndex = ClayParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClayParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Clay() IClayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClayContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ClayParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClayListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClayListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ClayParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ClayParserRULE_expression)

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
		p.Clay()
	}
	{
		p.SetState(5)
		p.Match(ClayParserEOF)
	}

	return localctx
}

// IClayContext is an interface to support dynamic dispatch.
type IClayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClayContext differentiates from other interfaces.
	IsClayContext()
}

type ClayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClayContext() *ClayContext {
	var p = new(ClayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ClayParserRULE_clay
	return p
}

func (*ClayContext) IsClayContext() {}

func NewClayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClayContext {
	var p = new(ClayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ClayParserRULE_clay

	return p
}

func (s *ClayContext) GetParser() antlr.Parser { return s.parser }

func (s *ClayContext) CLAY() antlr.TerminalNode {
	return s.GetToken(ClayParserCLAY, 0)
}

func (s *ClayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClayListener); ok {
		listenerT.EnterClay(s)
	}
}

func (s *ClayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ClayListener); ok {
		listenerT.ExitClay(s)
	}
}

func (p *ClayParser) Clay() (localctx IClayContext) {
	this := p
	_ = this

	localctx = NewClayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ClayParserRULE_clay)

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
		p.Match(ClayParserCLAY)
	}

	return localctx
}
