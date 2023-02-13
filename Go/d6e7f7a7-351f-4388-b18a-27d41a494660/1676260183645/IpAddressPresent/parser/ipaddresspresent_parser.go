// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676260183645/IpAddressPresent.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IpAddressPresent

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

type IpAddressPresentParser struct {
	*antlr.BaseParser
}

var ipaddresspresentParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ipaddresspresentParserInit() {
	staticData := &ipaddresspresentParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "IPADDRESSPRESENT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ipaddresspresent",
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

// IpAddressPresentParserInit initializes any static state used to implement IpAddressPresentParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIpAddressPresentParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IpAddressPresentParserInit() {
	staticData := &ipaddresspresentParserStaticData
	staticData.once.Do(ipaddresspresentParserInit)
}

// NewIpAddressPresentParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIpAddressPresentParser(input antlr.TokenStream) *IpAddressPresentParser {
	IpAddressPresentParserInit()
	this := new(IpAddressPresentParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ipaddresspresentParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "IpAddressPresent.g4"

	return this
}

// IpAddressPresentParser tokens.
const (
	IpAddressPresentParserEOF              = antlr.TokenEOF
	IpAddressPresentParserIPADDRESSPRESENT = 1
	IpAddressPresentParserOWNKEY           = 2
	IpAddressPresentParserSPLITKEY         = 3
	IpAddressPresentParserWS               = 4
)

// IpAddressPresentParser rules.
const (
	IpAddressPresentParserRULE_expression       = 0
	IpAddressPresentParserRULE_ipaddresspresent = 1
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
	p.RuleIndex = IpAddressPresentParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IpAddressPresentParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ipaddresspresent() IIpaddresspresentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIpaddresspresentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIpaddresspresentContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(IpAddressPresentParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IpAddressPresentListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IpAddressPresentListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *IpAddressPresentParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IpAddressPresentParserRULE_expression)

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
		p.Ipaddresspresent()
	}
	{
		p.SetState(5)
		p.Match(IpAddressPresentParserEOF)
	}

	return localctx
}

// IIpaddresspresentContext is an interface to support dynamic dispatch.
type IIpaddresspresentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpaddresspresentContext differentiates from other interfaces.
	IsIpaddresspresentContext()
}

type IpaddresspresentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpaddresspresentContext() *IpaddresspresentContext {
	var p = new(IpaddresspresentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = IpAddressPresentParserRULE_ipaddresspresent
	return p
}

func (*IpaddresspresentContext) IsIpaddresspresentContext() {}

func NewIpaddresspresentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IpaddresspresentContext {
	var p = new(IpaddresspresentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = IpAddressPresentParserRULE_ipaddresspresent

	return p
}

func (s *IpaddresspresentContext) GetParser() antlr.Parser { return s.parser }

func (s *IpaddresspresentContext) IPADDRESSPRESENT() antlr.TerminalNode {
	return s.GetToken(IpAddressPresentParserIPADDRESSPRESENT, 0)
}

func (s *IpaddresspresentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpaddresspresentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IpaddresspresentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IpAddressPresentListener); ok {
		listenerT.EnterIpaddresspresent(s)
	}
}

func (s *IpaddresspresentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IpAddressPresentListener); ok {
		listenerT.ExitIpaddresspresent(s)
	}
}

func (p *IpAddressPresentParser) Ipaddresspresent() (localctx IIpaddresspresentContext) {
	this := p
	_ = this

	localctx = NewIpaddresspresentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IpAddressPresentParserRULE_ipaddresspresent)

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
		p.Match(IpAddressPresentParserIPADDRESSPRESENT)
	}

	return localctx
}
