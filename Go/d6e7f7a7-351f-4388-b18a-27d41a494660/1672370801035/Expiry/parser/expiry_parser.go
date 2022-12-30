// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672370801035/Expiry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expiry

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

type ExpiryParser struct {
	*antlr.BaseParser
}

var expiryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expiryParserInit() {
	staticData := &expiryParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EXPIRY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "expiry",
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

// ExpiryParserInit initializes any static state used to implement ExpiryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpiryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpiryParserInit() {
	staticData := &expiryParserStaticData
	staticData.once.Do(expiryParserInit)
}

// NewExpiryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpiryParser(input antlr.TokenStream) *ExpiryParser {
	ExpiryParserInit()
	this := new(ExpiryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &expiryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Expiry.g4"

	return this
}

// ExpiryParser tokens.
const (
	ExpiryParserEOF      = antlr.TokenEOF
	ExpiryParserEXPIRY   = 1
	ExpiryParserOWNKEY   = 2
	ExpiryParserSPLITKEY = 3
	ExpiryParserWS       = 4
)

// ExpiryParser rules.
const (
	ExpiryParserRULE_expression = 0
	ExpiryParserRULE_expiry     = 1
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
	p.RuleIndex = ExpiryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpiryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Expiry() IExpiryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpiryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpiryContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpiryParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpiryListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpiryListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExpiryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpiryParserRULE_expression)

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
		p.Expiry()
	}
	{
		p.SetState(5)
		p.Match(ExpiryParserEOF)
	}

	return localctx
}

// IExpiryContext is an interface to support dynamic dispatch.
type IExpiryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpiryContext differentiates from other interfaces.
	IsExpiryContext()
}

type ExpiryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpiryContext() *ExpiryContext {
	var p = new(ExpiryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpiryParserRULE_expiry
	return p
}

func (*ExpiryContext) IsExpiryContext() {}

func NewExpiryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpiryContext {
	var p = new(ExpiryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpiryParserRULE_expiry

	return p
}

func (s *ExpiryContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpiryContext) EXPIRY() antlr.TerminalNode {
	return s.GetToken(ExpiryParserEXPIRY, 0)
}

func (s *ExpiryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpiryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpiryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpiryListener); ok {
		listenerT.EnterExpiry(s)
	}
}

func (s *ExpiryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpiryListener); ok {
		listenerT.ExitExpiry(s)
	}
}

func (p *ExpiryParser) Expiry() (localctx IExpiryContext) {
	this := p
	_ = this

	localctx = NewExpiryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpiryParserRULE_expiry)

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
		p.Match(ExpiryParserEXPIRY)
	}

	return localctx
}
