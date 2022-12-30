// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672380142110/Bleeding.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bleeding

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

type BleedingParser struct {
	*antlr.BaseParser
}

var bleedingParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bleedingParserInit() {
	staticData := &bleedingParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BLEEDING", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bleeding",
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

// BleedingParserInit initializes any static state used to implement BleedingParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBleedingParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BleedingParserInit() {
	staticData := &bleedingParserStaticData
	staticData.once.Do(bleedingParserInit)
}

// NewBleedingParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBleedingParser(input antlr.TokenStream) *BleedingParser {
	BleedingParserInit()
	this := new(BleedingParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bleedingParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bleeding.g4"

	return this
}

// BleedingParser tokens.
const (
	BleedingParserEOF      = antlr.TokenEOF
	BleedingParserBLEEDING = 1
	BleedingParserOWNKEY   = 2
	BleedingParserSPLITKEY = 3
	BleedingParserWS       = 4
)

// BleedingParser rules.
const (
	BleedingParserRULE_expression = 0
	BleedingParserRULE_bleeding   = 1
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
	p.RuleIndex = BleedingParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BleedingParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bleeding() IBleedingContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBleedingContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBleedingContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BleedingParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BleedingListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BleedingListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BleedingParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BleedingParserRULE_expression)

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
		p.Bleeding()
	}
	{
		p.SetState(5)
		p.Match(BleedingParserEOF)
	}

	return localctx
}

// IBleedingContext is an interface to support dynamic dispatch.
type IBleedingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBleedingContext differentiates from other interfaces.
	IsBleedingContext()
}

type BleedingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBleedingContext() *BleedingContext {
	var p = new(BleedingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BleedingParserRULE_bleeding
	return p
}

func (*BleedingContext) IsBleedingContext() {}

func NewBleedingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BleedingContext {
	var p = new(BleedingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BleedingParserRULE_bleeding

	return p
}

func (s *BleedingContext) GetParser() antlr.Parser { return s.parser }

func (s *BleedingContext) BLEEDING() antlr.TerminalNode {
	return s.GetToken(BleedingParserBLEEDING, 0)
}

func (s *BleedingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BleedingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BleedingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BleedingListener); ok {
		listenerT.EnterBleeding(s)
	}
}

func (s *BleedingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BleedingListener); ok {
		listenerT.ExitBleeding(s)
	}
}

func (p *BleedingParser) Bleeding() (localctx IBleedingContext) {
	this := p
	_ = this

	localctx = NewBleedingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BleedingParserRULE_bleeding)

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
		p.Match(BleedingParserBLEEDING)
	}

	return localctx
}
