// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675222431329/TaxType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TaxType

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

type TaxTypeParser struct {
	*antlr.BaseParser
}

var taxtypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func taxtypeParserInit() {
	staticData := &taxtypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TAXTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "taxtype",
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

// TaxTypeParserInit initializes any static state used to implement TaxTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTaxTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TaxTypeParserInit() {
	staticData := &taxtypeParserStaticData
	staticData.once.Do(taxtypeParserInit)
}

// NewTaxTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTaxTypeParser(input antlr.TokenStream) *TaxTypeParser {
	TaxTypeParserInit()
	this := new(TaxTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &taxtypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "TaxType.g4"

	return this
}

// TaxTypeParser tokens.
const (
	TaxTypeParserEOF      = antlr.TokenEOF
	TaxTypeParserTAXTYPE  = 1
	TaxTypeParserOWNKEY   = 2
	TaxTypeParserSPLITKEY = 3
	TaxTypeParserWS       = 4
)

// TaxTypeParser rules.
const (
	TaxTypeParserRULE_expression = 0
	TaxTypeParserRULE_taxtype    = 1
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
	p.RuleIndex = TaxTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TaxTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Taxtype() ITaxtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITaxtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITaxtypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TaxTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TaxTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TaxTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TaxTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TaxTypeParserRULE_expression)

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
		p.Taxtype()
	}
	{
		p.SetState(5)
		p.Match(TaxTypeParserEOF)
	}

	return localctx
}

// ITaxtypeContext is an interface to support dynamic dispatch.
type ITaxtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTaxtypeContext differentiates from other interfaces.
	IsTaxtypeContext()
}

type TaxtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTaxtypeContext() *TaxtypeContext {
	var p = new(TaxtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TaxTypeParserRULE_taxtype
	return p
}

func (*TaxtypeContext) IsTaxtypeContext() {}

func NewTaxtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TaxtypeContext {
	var p = new(TaxtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TaxTypeParserRULE_taxtype

	return p
}

func (s *TaxtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TaxtypeContext) TAXTYPE() antlr.TerminalNode {
	return s.GetToken(TaxTypeParserTAXTYPE, 0)
}

func (s *TaxtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TaxtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TaxtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TaxTypeListener); ok {
		listenerT.EnterTaxtype(s)
	}
}

func (s *TaxtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TaxTypeListener); ok {
		listenerT.ExitTaxtype(s)
	}
}

func (p *TaxTypeParser) Taxtype() (localctx ITaxtypeContext) {
	this := p
	_ = this

	localctx = NewTaxtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TaxTypeParserRULE_taxtype)

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
		p.Match(TaxTypeParserTAXTYPE)
	}

	return localctx
}
