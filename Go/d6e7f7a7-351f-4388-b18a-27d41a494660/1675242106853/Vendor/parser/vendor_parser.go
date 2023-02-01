// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675242106853/Vendor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vendor

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

type VendorParser struct {
	*antlr.BaseParser
}

var vendorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vendorParserInit() {
	staticData := &vendorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VENDOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "vendor",
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

// VendorParserInit initializes any static state used to implement VendorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVendorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VendorParserInit() {
	staticData := &vendorParserStaticData
	staticData.once.Do(vendorParserInit)
}

// NewVendorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVendorParser(input antlr.TokenStream) *VendorParser {
	VendorParserInit()
	this := new(VendorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &vendorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Vendor.g4"

	return this
}

// VendorParser tokens.
const (
	VendorParserEOF      = antlr.TokenEOF
	VendorParserVENDOR   = 1
	VendorParserOWNKEY   = 2
	VendorParserSPLITKEY = 3
	VendorParserWS       = 4
)

// VendorParser rules.
const (
	VendorParserRULE_expression = 0
	VendorParserRULE_vendor     = 1
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
	p.RuleIndex = VendorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VendorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Vendor() IVendorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVendorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVendorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VendorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VendorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VendorParserRULE_expression)

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
		p.Vendor()
	}
	{
		p.SetState(5)
		p.Match(VendorParserEOF)
	}

	return localctx
}

// IVendorContext is an interface to support dynamic dispatch.
type IVendorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVendorContext differentiates from other interfaces.
	IsVendorContext()
}

type VendorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVendorContext() *VendorContext {
	var p = new(VendorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VendorParserRULE_vendor
	return p
}

func (*VendorContext) IsVendorContext() {}

func NewVendorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VendorContext {
	var p = new(VendorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VendorParserRULE_vendor

	return p
}

func (s *VendorContext) GetParser() antlr.Parser { return s.parser }

func (s *VendorContext) VENDOR() antlr.TerminalNode {
	return s.GetToken(VendorParserVENDOR, 0)
}

func (s *VendorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VendorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VendorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendorListener); ok {
		listenerT.EnterVendor(s)
	}
}

func (s *VendorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendorListener); ok {
		listenerT.ExitVendor(s)
	}
}

func (p *VendorParser) Vendor() (localctx IVendorContext) {
	this := p
	_ = this

	localctx = NewVendorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VendorParserRULE_vendor)

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
		p.Match(VendorParserVENDOR)
	}

	return localctx
}
