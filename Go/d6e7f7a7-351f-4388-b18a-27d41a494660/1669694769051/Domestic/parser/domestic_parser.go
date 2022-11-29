// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669694769051/Domestic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Domestic

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

type DomesticParser struct {
	*antlr.BaseParser
}

var domesticParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func domesticParserInit() {
	staticData := &domesticParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DOMESTIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "domestic",
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

// DomesticParserInit initializes any static state used to implement DomesticParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDomesticParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DomesticParserInit() {
	staticData := &domesticParserStaticData
	staticData.once.Do(domesticParserInit)
}

// NewDomesticParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDomesticParser(input antlr.TokenStream) *DomesticParser {
	DomesticParserInit()
	this := new(DomesticParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &domesticParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Domestic.g4"

	return this
}

// DomesticParser tokens.
const (
	DomesticParserEOF      = antlr.TokenEOF
	DomesticParserDOMESTIC = 1
	DomesticParserOWNKEY   = 2
	DomesticParserSPLITKEY = 3
	DomesticParserWS       = 4
)

// DomesticParser rules.
const (
	DomesticParserRULE_expression = 0
	DomesticParserRULE_domestic   = 1
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
	p.RuleIndex = DomesticParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DomesticParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Domestic() IDomesticContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDomesticContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDomesticContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DomesticParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DomesticListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DomesticListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DomesticParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DomesticParserRULE_expression)

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
		p.Domestic()
	}
	{
		p.SetState(5)
		p.Match(DomesticParserEOF)
	}

	return localctx
}

// IDomesticContext is an interface to support dynamic dispatch.
type IDomesticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDomesticContext differentiates from other interfaces.
	IsDomesticContext()
}

type DomesticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDomesticContext() *DomesticContext {
	var p = new(DomesticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DomesticParserRULE_domestic
	return p
}

func (*DomesticContext) IsDomesticContext() {}

func NewDomesticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DomesticContext {
	var p = new(DomesticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DomesticParserRULE_domestic

	return p
}

func (s *DomesticContext) GetParser() antlr.Parser { return s.parser }

func (s *DomesticContext) DOMESTIC() antlr.TerminalNode {
	return s.GetToken(DomesticParserDOMESTIC, 0)
}

func (s *DomesticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DomesticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DomesticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DomesticListener); ok {
		listenerT.EnterDomestic(s)
	}
}

func (s *DomesticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DomesticListener); ok {
		listenerT.ExitDomestic(s)
	}
}

func (p *DomesticParser) Domestic() (localctx IDomesticContext) {
	this := p
	_ = this

	localctx = NewDomesticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DomesticParserRULE_domestic)

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
		p.Match(DomesticParserDOMESTIC)
	}

	return localctx
}
