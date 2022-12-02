// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972122105/Close.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Close

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

type CloseParser struct {
	*antlr.BaseParser
}

var closeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func closeParserInit() {
	staticData := &closeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CLOSE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "close",
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

// CloseParserInit initializes any static state used to implement CloseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCloseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CloseParserInit() {
	staticData := &closeParserStaticData
	staticData.once.Do(closeParserInit)
}

// NewCloseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCloseParser(input antlr.TokenStream) *CloseParser {
	CloseParserInit()
	this := new(CloseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &closeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Close.g4"

	return this
}

// CloseParser tokens.
const (
	CloseParserEOF      = antlr.TokenEOF
	CloseParserCLOSE    = 1
	CloseParserOWNKEY   = 2
	CloseParserSPLITKEY = 3
	CloseParserWS       = 4
)

// CloseParser rules.
const (
	CloseParserRULE_expression = 0
	CloseParserRULE_close      = 1
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
	p.RuleIndex = CloseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CloseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Close() ICloseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICloseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICloseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CloseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CloseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CloseParserRULE_expression)

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
		p.Close()
	}
	{
		p.SetState(5)
		p.Match(CloseParserEOF)
	}

	return localctx
}

// ICloseContext is an interface to support dynamic dispatch.
type ICloseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCloseContext differentiates from other interfaces.
	IsCloseContext()
}

type CloseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloseContext() *CloseContext {
	var p = new(CloseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CloseParserRULE_close
	return p
}

func (*CloseContext) IsCloseContext() {}

func NewCloseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseContext {
	var p = new(CloseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CloseParserRULE_close

	return p
}

func (s *CloseContext) GetParser() antlr.Parser { return s.parser }

func (s *CloseContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(CloseParserCLOSE, 0)
}

func (s *CloseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CloseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloseListener); ok {
		listenerT.EnterClose(s)
	}
}

func (s *CloseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CloseListener); ok {
		listenerT.ExitClose(s)
	}
}

func (p *CloseParser) Close() (localctx ICloseContext) {
	this := p
	_ = this

	localctx = NewCloseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CloseParserRULE_close)

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
		p.Match(CloseParserCLOSE)
	}

	return localctx
}
