// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672985492567/Vendornoproduits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Vendornoproduits

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

type VendornoproduitsParser struct {
	*antlr.BaseParser
}

var vendornoproduitsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vendornoproduitsParserInit() {
	staticData := &vendornoproduitsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VENDORNOPRODUITS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "vendornoproduits",
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

// VendornoproduitsParserInit initializes any static state used to implement VendornoproduitsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVendornoproduitsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VendornoproduitsParserInit() {
	staticData := &vendornoproduitsParserStaticData
	staticData.once.Do(vendornoproduitsParserInit)
}

// NewVendornoproduitsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVendornoproduitsParser(input antlr.TokenStream) *VendornoproduitsParser {
	VendornoproduitsParserInit()
	this := new(VendornoproduitsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &vendornoproduitsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Vendornoproduits.g4"

	return this
}

// VendornoproduitsParser tokens.
const (
	VendornoproduitsParserEOF              = antlr.TokenEOF
	VendornoproduitsParserVENDORNOPRODUITS = 1
	VendornoproduitsParserOWNKEY           = 2
	VendornoproduitsParserSPLITKEY         = 3
	VendornoproduitsParserWS               = 4
)

// VendornoproduitsParser rules.
const (
	VendornoproduitsParserRULE_expression       = 0
	VendornoproduitsParserRULE_vendornoproduits = 1
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
	p.RuleIndex = VendornoproduitsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VendornoproduitsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Vendornoproduits() IVendornoproduitsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVendornoproduitsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVendornoproduitsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VendornoproduitsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendornoproduitsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendornoproduitsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VendornoproduitsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VendornoproduitsParserRULE_expression)

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
		p.Vendornoproduits()
	}
	{
		p.SetState(5)
		p.Match(VendornoproduitsParserEOF)
	}

	return localctx
}

// IVendornoproduitsContext is an interface to support dynamic dispatch.
type IVendornoproduitsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVendornoproduitsContext differentiates from other interfaces.
	IsVendornoproduitsContext()
}

type VendornoproduitsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVendornoproduitsContext() *VendornoproduitsContext {
	var p = new(VendornoproduitsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VendornoproduitsParserRULE_vendornoproduits
	return p
}

func (*VendornoproduitsContext) IsVendornoproduitsContext() {}

func NewVendornoproduitsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VendornoproduitsContext {
	var p = new(VendornoproduitsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VendornoproduitsParserRULE_vendornoproduits

	return p
}

func (s *VendornoproduitsContext) GetParser() antlr.Parser { return s.parser }

func (s *VendornoproduitsContext) VENDORNOPRODUITS() antlr.TerminalNode {
	return s.GetToken(VendornoproduitsParserVENDORNOPRODUITS, 0)
}

func (s *VendornoproduitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VendornoproduitsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VendornoproduitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendornoproduitsListener); ok {
		listenerT.EnterVendornoproduits(s)
	}
}

func (s *VendornoproduitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VendornoproduitsListener); ok {
		listenerT.ExitVendornoproduits(s)
	}
}

func (p *VendornoproduitsParser) Vendornoproduits() (localctx IVendornoproduitsContext) {
	this := p
	_ = this

	localctx = NewVendornoproduitsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VendornoproduitsParserRULE_vendornoproduits)

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
		p.Match(VendornoproduitsParserVENDORNOPRODUITS)
	}

	return localctx
}
