// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076364337/BulkDensity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BulkDensity

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

type BulkDensityParser struct {
	*antlr.BaseParser
}

var bulkdensityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bulkdensityParserInit() {
	staticData := &bulkdensityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BULKDENSITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bulkdensity",
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

// BulkDensityParserInit initializes any static state used to implement BulkDensityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBulkDensityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BulkDensityParserInit() {
	staticData := &bulkdensityParserStaticData
	staticData.once.Do(bulkdensityParserInit)
}

// NewBulkDensityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBulkDensityParser(input antlr.TokenStream) *BulkDensityParser {
	BulkDensityParserInit()
	this := new(BulkDensityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bulkdensityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BulkDensity.g4"

	return this
}

// BulkDensityParser tokens.
const (
	BulkDensityParserEOF         = antlr.TokenEOF
	BulkDensityParserBULKDENSITY = 1
	BulkDensityParserOWNKEY      = 2
	BulkDensityParserSPLITKEY    = 3
	BulkDensityParserWS          = 4
)

// BulkDensityParser rules.
const (
	BulkDensityParserRULE_expression  = 0
	BulkDensityParserRULE_bulkdensity = 1
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
	p.RuleIndex = BulkDensityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BulkDensityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bulkdensity() IBulkdensityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBulkdensityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBulkdensityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BulkDensityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BulkDensityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BulkDensityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BulkDensityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BulkDensityParserRULE_expression)

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
		p.Bulkdensity()
	}
	{
		p.SetState(5)
		p.Match(BulkDensityParserEOF)
	}

	return localctx
}

// IBulkdensityContext is an interface to support dynamic dispatch.
type IBulkdensityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBulkdensityContext differentiates from other interfaces.
	IsBulkdensityContext()
}

type BulkdensityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBulkdensityContext() *BulkdensityContext {
	var p = new(BulkdensityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BulkDensityParserRULE_bulkdensity
	return p
}

func (*BulkdensityContext) IsBulkdensityContext() {}

func NewBulkdensityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BulkdensityContext {
	var p = new(BulkdensityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BulkDensityParserRULE_bulkdensity

	return p
}

func (s *BulkdensityContext) GetParser() antlr.Parser { return s.parser }

func (s *BulkdensityContext) BULKDENSITY() antlr.TerminalNode {
	return s.GetToken(BulkDensityParserBULKDENSITY, 0)
}

func (s *BulkdensityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BulkdensityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BulkdensityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BulkDensityListener); ok {
		listenerT.EnterBulkdensity(s)
	}
}

func (s *BulkdensityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BulkDensityListener); ok {
		listenerT.ExitBulkdensity(s)
	}
}

func (p *BulkDensityParser) Bulkdensity() (localctx IBulkdensityContext) {
	this := p
	_ = this

	localctx = NewBulkdensityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BulkDensityParserRULE_bulkdensity)

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
		p.Match(BulkDensityParserBULKDENSITY)
	}

	return localctx
}
