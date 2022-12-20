// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515428619/Child.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Child

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

type ChildParser struct {
	*antlr.BaseParser
}

var childParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func childParserInit() {
	staticData := &childParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHILD", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "child",
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

// ChildParserInit initializes any static state used to implement ChildParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewChildParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChildParserInit() {
	staticData := &childParserStaticData
	staticData.once.Do(childParserInit)
}

// NewChildParser produces a new parser instance for the optional input antlr.TokenStream.
func NewChildParser(input antlr.TokenStream) *ChildParser {
	ChildParserInit()
	this := new(ChildParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &childParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Child.g4"

	return this
}

// ChildParser tokens.
const (
	ChildParserEOF      = antlr.TokenEOF
	ChildParserCHILD    = 1
	ChildParserOWNKEY   = 2
	ChildParserSPLITKEY = 3
	ChildParserWS       = 4
)

// ChildParser rules.
const (
	ChildParserRULE_expression = 0
	ChildParserRULE_child      = 1
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
	p.RuleIndex = ChildParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChildParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Child() IChildContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChildContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChildContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ChildParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChildListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChildListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ChildParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChildParserRULE_expression)

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
		p.Child()
	}
	{
		p.SetState(5)
		p.Match(ChildParserEOF)
	}

	return localctx
}

// IChildContext is an interface to support dynamic dispatch.
type IChildContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChildContext differentiates from other interfaces.
	IsChildContext()
}

type ChildContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChildContext() *ChildContext {
	var p = new(ChildContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChildParserRULE_child
	return p
}

func (*ChildContext) IsChildContext() {}

func NewChildContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChildContext {
	var p = new(ChildContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChildParserRULE_child

	return p
}

func (s *ChildContext) GetParser() antlr.Parser { return s.parser }

func (s *ChildContext) CHILD() antlr.TerminalNode {
	return s.GetToken(ChildParserCHILD, 0)
}

func (s *ChildContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChildContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChildContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChildListener); ok {
		listenerT.EnterChild(s)
	}
}

func (s *ChildContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChildListener); ok {
		listenerT.ExitChild(s)
	}
}

func (p *ChildParser) Child() (localctx IChildContext) {
	this := p
	_ = this

	localctx = NewChildContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChildParserRULE_child)

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
		p.Match(ChildParserCHILD)
	}

	return localctx
}
