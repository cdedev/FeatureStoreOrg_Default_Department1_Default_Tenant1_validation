// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674787464947/Barcode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Barcode

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

type BarcodeParser struct {
	*antlr.BaseParser
}

var barcodeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func barcodeParserInit() {
	staticData := &barcodeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BARCODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "barcode",
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

// BarcodeParserInit initializes any static state used to implement BarcodeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBarcodeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BarcodeParserInit() {
	staticData := &barcodeParserStaticData
	staticData.once.Do(barcodeParserInit)
}

// NewBarcodeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBarcodeParser(input antlr.TokenStream) *BarcodeParser {
	BarcodeParserInit()
	this := new(BarcodeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &barcodeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Barcode.g4"

	return this
}

// BarcodeParser tokens.
const (
	BarcodeParserEOF      = antlr.TokenEOF
	BarcodeParserBARCODE  = 1
	BarcodeParserOWNKEY   = 2
	BarcodeParserSPLITKEY = 3
	BarcodeParserWS       = 4
)

// BarcodeParser rules.
const (
	BarcodeParserRULE_expression = 0
	BarcodeParserRULE_barcode    = 1
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
	p.RuleIndex = BarcodeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BarcodeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Barcode() IBarcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBarcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBarcodeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BarcodeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BarcodeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BarcodeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BarcodeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BarcodeParserRULE_expression)

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
		p.Barcode()
	}
	{
		p.SetState(5)
		p.Match(BarcodeParserEOF)
	}

	return localctx
}

// IBarcodeContext is an interface to support dynamic dispatch.
type IBarcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBarcodeContext differentiates from other interfaces.
	IsBarcodeContext()
}

type BarcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBarcodeContext() *BarcodeContext {
	var p = new(BarcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BarcodeParserRULE_barcode
	return p
}

func (*BarcodeContext) IsBarcodeContext() {}

func NewBarcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BarcodeContext {
	var p = new(BarcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BarcodeParserRULE_barcode

	return p
}

func (s *BarcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BarcodeContext) BARCODE() antlr.TerminalNode {
	return s.GetToken(BarcodeParserBARCODE, 0)
}

func (s *BarcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BarcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BarcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BarcodeListener); ok {
		listenerT.EnterBarcode(s)
	}
}

func (s *BarcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BarcodeListener); ok {
		listenerT.ExitBarcode(s)
	}
}

func (p *BarcodeParser) Barcode() (localctx IBarcodeContext) {
	this := p
	_ = this

	localctx = NewBarcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BarcodeParserRULE_barcode)

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
		p.Match(BarcodeParserBARCODE)
	}

	return localctx
}
