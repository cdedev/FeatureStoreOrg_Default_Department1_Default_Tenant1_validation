// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701421426/Brandname.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Brandname

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

type BrandnameParser struct {
	*antlr.BaseParser
}

var brandnameParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func brandnameParserInit() {
	staticData := &brandnameParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BRANDNAME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "brandname",
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

// BrandnameParserInit initializes any static state used to implement BrandnameParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBrandnameParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BrandnameParserInit() {
	staticData := &brandnameParserStaticData
	staticData.once.Do(brandnameParserInit)
}

// NewBrandnameParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBrandnameParser(input antlr.TokenStream) *BrandnameParser {
	BrandnameParserInit()
	this := new(BrandnameParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &brandnameParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Brandname.g4"

	return this
}

// BrandnameParser tokens.
const (
	BrandnameParserEOF       = antlr.TokenEOF
	BrandnameParserBRANDNAME = 1
	BrandnameParserOWNKEY    = 2
	BrandnameParserSPLITKEY  = 3
	BrandnameParserWS        = 4
)

// BrandnameParser rules.
const (
	BrandnameParserRULE_expression = 0
	BrandnameParserRULE_brandname  = 1
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
	p.RuleIndex = BrandnameParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrandnameParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Brandname() IBrandnameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrandnameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrandnameContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BrandnameParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrandnameListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrandnameListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BrandnameParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BrandnameParserRULE_expression)

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
		p.Brandname()
	}
	{
		p.SetState(5)
		p.Match(BrandnameParserEOF)
	}

	return localctx
}

// IBrandnameContext is an interface to support dynamic dispatch.
type IBrandnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBrandnameContext differentiates from other interfaces.
	IsBrandnameContext()
}

type BrandnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrandnameContext() *BrandnameContext {
	var p = new(BrandnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrandnameParserRULE_brandname
	return p
}

func (*BrandnameContext) IsBrandnameContext() {}

func NewBrandnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrandnameContext {
	var p = new(BrandnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrandnameParserRULE_brandname

	return p
}

func (s *BrandnameContext) GetParser() antlr.Parser { return s.parser }

func (s *BrandnameContext) BRANDNAME() antlr.TerminalNode {
	return s.GetToken(BrandnameParserBRANDNAME, 0)
}

func (s *BrandnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrandnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrandnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrandnameListener); ok {
		listenerT.EnterBrandname(s)
	}
}

func (s *BrandnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrandnameListener); ok {
		listenerT.ExitBrandname(s)
	}
}

func (p *BrandnameParser) Brandname() (localctx IBrandnameContext) {
	this := p
	_ = this

	localctx = NewBrandnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BrandnameParserRULE_brandname)

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
		p.Match(BrandnameParserBRANDNAME)
	}

	return localctx
}
