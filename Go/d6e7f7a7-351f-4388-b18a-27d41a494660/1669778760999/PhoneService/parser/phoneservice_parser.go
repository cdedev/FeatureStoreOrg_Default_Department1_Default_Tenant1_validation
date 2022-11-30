// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778760999/PhoneService.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PhoneService

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

type PhoneServiceParser struct {
	*antlr.BaseParser
}

var phoneserviceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func phoneserviceParserInit() {
	staticData := &phoneserviceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PHONESERVICE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "phoneservice",
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

// PhoneServiceParserInit initializes any static state used to implement PhoneServiceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPhoneServiceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PhoneServiceParserInit() {
	staticData := &phoneserviceParserStaticData
	staticData.once.Do(phoneserviceParserInit)
}

// NewPhoneServiceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPhoneServiceParser(input antlr.TokenStream) *PhoneServiceParser {
	PhoneServiceParserInit()
	this := new(PhoneServiceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &phoneserviceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PhoneService.g4"

	return this
}

// PhoneServiceParser tokens.
const (
	PhoneServiceParserEOF          = antlr.TokenEOF
	PhoneServiceParserPHONESERVICE = 1
	PhoneServiceParserOWNKEY       = 2
	PhoneServiceParserSPLITKEY     = 3
	PhoneServiceParserWS           = 4
)

// PhoneServiceParser rules.
const (
	PhoneServiceParserRULE_expression   = 0
	PhoneServiceParserRULE_phoneservice = 1
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
	p.RuleIndex = PhoneServiceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhoneServiceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Phoneservice() IPhoneserviceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhoneserviceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhoneserviceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PhoneServiceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneServiceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneServiceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PhoneServiceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PhoneServiceParserRULE_expression)

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
		p.Phoneservice()
	}
	{
		p.SetState(5)
		p.Match(PhoneServiceParserEOF)
	}

	return localctx
}

// IPhoneserviceContext is an interface to support dynamic dispatch.
type IPhoneserviceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhoneserviceContext differentiates from other interfaces.
	IsPhoneserviceContext()
}

type PhoneserviceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhoneserviceContext() *PhoneserviceContext {
	var p = new(PhoneserviceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PhoneServiceParserRULE_phoneservice
	return p
}

func (*PhoneserviceContext) IsPhoneserviceContext() {}

func NewPhoneserviceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhoneserviceContext {
	var p = new(PhoneserviceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhoneServiceParserRULE_phoneservice

	return p
}

func (s *PhoneserviceContext) GetParser() antlr.Parser { return s.parser }

func (s *PhoneserviceContext) PHONESERVICE() antlr.TerminalNode {
	return s.GetToken(PhoneServiceParserPHONESERVICE, 0)
}

func (s *PhoneserviceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhoneserviceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhoneserviceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneServiceListener); ok {
		listenerT.EnterPhoneservice(s)
	}
}

func (s *PhoneserviceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneServiceListener); ok {
		listenerT.ExitPhoneservice(s)
	}
}

func (p *PhoneServiceParser) Phoneservice() (localctx IPhoneserviceContext) {
	this := p
	_ = this

	localctx = NewPhoneserviceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PhoneServiceParserRULE_phoneservice)

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
		p.Match(PhoneServiceParserPHONESERVICE)
	}

	return localctx
}
