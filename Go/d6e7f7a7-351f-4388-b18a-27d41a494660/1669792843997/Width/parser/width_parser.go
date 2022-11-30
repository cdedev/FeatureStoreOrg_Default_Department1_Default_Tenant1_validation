// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669792843997/Width.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Width

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

type WidthParser struct {
	*antlr.BaseParser
}

var widthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func widthParserInit() {
	staticData := &widthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WIDTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "width",
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

// WidthParserInit initializes any static state used to implement WidthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWidthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WidthParserInit() {
	staticData := &widthParserStaticData
	staticData.once.Do(widthParserInit)
}

// NewWidthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWidthParser(input antlr.TokenStream) *WidthParser {
	WidthParserInit()
	this := new(WidthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &widthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Width.g4"

	return this
}

// WidthParser tokens.
const (
	WidthParserEOF      = antlr.TokenEOF
	WidthParserWIDTH    = 1
	WidthParserOWNKEY   = 2
	WidthParserSPLITKEY = 3
	WidthParserWS       = 4
)

// WidthParser rules.
const (
	WidthParserRULE_expression = 0
	WidthParserRULE_width      = 1
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
	p.RuleIndex = WidthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WidthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Width() IWidthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWidthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWidthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WidthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WidthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WidthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WidthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WidthParserRULE_expression)

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
		p.Width()
	}
	{
		p.SetState(5)
		p.Match(WidthParserEOF)
	}

	return localctx
}

// IWidthContext is an interface to support dynamic dispatch.
type IWidthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWidthContext differentiates from other interfaces.
	IsWidthContext()
}

type WidthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWidthContext() *WidthContext {
	var p = new(WidthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WidthParserRULE_width
	return p
}

func (*WidthContext) IsWidthContext() {}

func NewWidthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WidthContext {
	var p = new(WidthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WidthParserRULE_width

	return p
}

func (s *WidthContext) GetParser() antlr.Parser { return s.parser }

func (s *WidthContext) WIDTH() antlr.TerminalNode {
	return s.GetToken(WidthParserWIDTH, 0)
}

func (s *WidthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WidthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WidthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WidthListener); ok {
		listenerT.EnterWidth(s)
	}
}

func (s *WidthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WidthListener); ok {
		listenerT.ExitWidth(s)
	}
}

func (p *WidthParser) Width() (localctx IWidthContext) {
	this := p
	_ = this

	localctx = NewWidthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WidthParserRULE_width)

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
		p.Match(WidthParserWIDTH)
	}

	return localctx
}
