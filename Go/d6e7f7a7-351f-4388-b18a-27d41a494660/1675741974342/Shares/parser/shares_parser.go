// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675741974342/Shares.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shares

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

type SharesParser struct {
	*antlr.BaseParser
}

var sharesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sharesParserInit() {
	staticData := &sharesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SHARES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "shares",
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

// SharesParserInit initializes any static state used to implement SharesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSharesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SharesParserInit() {
	staticData := &sharesParserStaticData
	staticData.once.Do(sharesParserInit)
}

// NewSharesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSharesParser(input antlr.TokenStream) *SharesParser {
	SharesParserInit()
	this := new(SharesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sharesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Shares.g4"

	return this
}

// SharesParser tokens.
const (
	SharesParserEOF      = antlr.TokenEOF
	SharesParserSHARES   = 1
	SharesParserOWNKEY   = 2
	SharesParserSPLITKEY = 3
	SharesParserWS       = 4
)

// SharesParser rules.
const (
	SharesParserRULE_expression = 0
	SharesParserRULE_shares     = 1
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
	p.RuleIndex = SharesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SharesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Shares() ISharesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISharesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISharesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SharesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SharesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SharesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SharesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SharesParserRULE_expression)

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
		p.Shares()
	}
	{
		p.SetState(5)
		p.Match(SharesParserEOF)
	}

	return localctx
}

// ISharesContext is an interface to support dynamic dispatch.
type ISharesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSharesContext differentiates from other interfaces.
	IsSharesContext()
}

type SharesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySharesContext() *SharesContext {
	var p = new(SharesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SharesParserRULE_shares
	return p
}

func (*SharesContext) IsSharesContext() {}

func NewSharesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SharesContext {
	var p = new(SharesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SharesParserRULE_shares

	return p
}

func (s *SharesContext) GetParser() antlr.Parser { return s.parser }

func (s *SharesContext) SHARES() antlr.TerminalNode {
	return s.GetToken(SharesParserSHARES, 0)
}

func (s *SharesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SharesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SharesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SharesListener); ok {
		listenerT.EnterShares(s)
	}
}

func (s *SharesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SharesListener); ok {
		listenerT.ExitShares(s)
	}
}

func (p *SharesParser) Shares() (localctx ISharesContext) {
	this := p
	_ = this

	localctx = NewSharesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SharesParserRULE_shares)

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
		p.Match(SharesParserSHARES)
	}

	return localctx
}
