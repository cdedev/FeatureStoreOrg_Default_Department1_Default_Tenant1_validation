// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528758541/Manufacturer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Manufacturer

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

type ManufacturerParser struct {
	*antlr.BaseParser
}

var manufacturerParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func manufacturerParserInit() {
	staticData := &manufacturerParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MANUFACTURER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "manufacturer",
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

// ManufacturerParserInit initializes any static state used to implement ManufacturerParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewManufacturerParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ManufacturerParserInit() {
	staticData := &manufacturerParserStaticData
	staticData.once.Do(manufacturerParserInit)
}

// NewManufacturerParser produces a new parser instance for the optional input antlr.TokenStream.
func NewManufacturerParser(input antlr.TokenStream) *ManufacturerParser {
	ManufacturerParserInit()
	this := new(ManufacturerParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &manufacturerParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Manufacturer.g4"

	return this
}

// ManufacturerParser tokens.
const (
	ManufacturerParserEOF          = antlr.TokenEOF
	ManufacturerParserMANUFACTURER = 1
	ManufacturerParserOWNKEY       = 2
	ManufacturerParserSPLITKEY     = 3
	ManufacturerParserWS           = 4
)

// ManufacturerParser rules.
const (
	ManufacturerParserRULE_expression   = 0
	ManufacturerParserRULE_manufacturer = 1
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
	p.RuleIndex = ManufacturerParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManufacturerParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Manufacturer() IManufacturerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IManufacturerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IManufacturerContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ManufacturerParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManufacturerListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManufacturerListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ManufacturerParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ManufacturerParserRULE_expression)

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
		p.Manufacturer()
	}
	{
		p.SetState(5)
		p.Match(ManufacturerParserEOF)
	}

	return localctx
}

// IManufacturerContext is an interface to support dynamic dispatch.
type IManufacturerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsManufacturerContext differentiates from other interfaces.
	IsManufacturerContext()
}

type ManufacturerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyManufacturerContext() *ManufacturerContext {
	var p = new(ManufacturerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ManufacturerParserRULE_manufacturer
	return p
}

func (*ManufacturerContext) IsManufacturerContext() {}

func NewManufacturerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ManufacturerContext {
	var p = new(ManufacturerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ManufacturerParserRULE_manufacturer

	return p
}

func (s *ManufacturerContext) GetParser() antlr.Parser { return s.parser }

func (s *ManufacturerContext) MANUFACTURER() antlr.TerminalNode {
	return s.GetToken(ManufacturerParserMANUFACTURER, 0)
}

func (s *ManufacturerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ManufacturerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ManufacturerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManufacturerListener); ok {
		listenerT.EnterManufacturer(s)
	}
}

func (s *ManufacturerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ManufacturerListener); ok {
		listenerT.ExitManufacturer(s)
	}
}

func (p *ManufacturerParser) Manufacturer() (localctx IManufacturerContext) {
	this := p
	_ = this

	localctx = NewManufacturerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ManufacturerParserRULE_manufacturer)

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
		p.Match(ManufacturerParserMANUFACTURER)
	}

	return localctx
}
