// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145305020/Bytes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bytes

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

type BytesParser struct {
	*antlr.BaseParser
}

var bytesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bytesParserInit() {
	staticData := &bytesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BYTES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bytes",
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

// BytesParserInit initializes any static state used to implement BytesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBytesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BytesParserInit() {
	staticData := &bytesParserStaticData
	staticData.once.Do(bytesParserInit)
}

// NewBytesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBytesParser(input antlr.TokenStream) *BytesParser {
	BytesParserInit()
	this := new(BytesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bytesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bytes.g4"

	return this
}

// BytesParser tokens.
const (
	BytesParserEOF      = antlr.TokenEOF
	BytesParserBYTES    = 1
	BytesParserOWNKEY   = 2
	BytesParserSPLITKEY = 3
	BytesParserWS       = 4
)

// BytesParser rules.
const (
	BytesParserRULE_expression = 0
	BytesParserRULE_bytes      = 1
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
	p.RuleIndex = BytesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BytesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bytes() IBytesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBytesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBytesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BytesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BytesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BytesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BytesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BytesParserRULE_expression)

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
		p.Bytes()
	}
	{
		p.SetState(5)
		p.Match(BytesParserEOF)
	}

	return localctx
}

// IBytesContext is an interface to support dynamic dispatch.
type IBytesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBytesContext differentiates from other interfaces.
	IsBytesContext()
}

type BytesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBytesContext() *BytesContext {
	var p = new(BytesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BytesParserRULE_bytes
	return p
}

func (*BytesContext) IsBytesContext() {}

func NewBytesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BytesContext {
	var p = new(BytesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BytesParserRULE_bytes

	return p
}

func (s *BytesContext) GetParser() antlr.Parser { return s.parser }

func (s *BytesContext) BYTES() antlr.TerminalNode {
	return s.GetToken(BytesParserBYTES, 0)
}

func (s *BytesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BytesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BytesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BytesListener); ok {
		listenerT.EnterBytes(s)
	}
}

func (s *BytesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BytesListener); ok {
		listenerT.ExitBytes(s)
	}
}

func (p *BytesParser) Bytes() (localctx IBytesContext) {
	this := p
	_ = this

	localctx = NewBytesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BytesParserRULE_bytes)

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
		p.Match(BytesParserBYTES)
	}

	return localctx
}
