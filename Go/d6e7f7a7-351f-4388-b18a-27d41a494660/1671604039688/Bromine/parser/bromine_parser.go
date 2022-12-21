// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604039688/Bromine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bromine

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

type BromineParser struct {
	*antlr.BaseParser
}

var bromineParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bromineParserInit() {
	staticData := &bromineParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BROMINE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bromine",
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

// BromineParserInit initializes any static state used to implement BromineParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBromineParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BromineParserInit() {
	staticData := &bromineParserStaticData
	staticData.once.Do(bromineParserInit)
}

// NewBromineParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBromineParser(input antlr.TokenStream) *BromineParser {
	BromineParserInit()
	this := new(BromineParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bromineParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bromine.g4"

	return this
}

// BromineParser tokens.
const (
	BromineParserEOF      = antlr.TokenEOF
	BromineParserBROMINE  = 1
	BromineParserOWNKEY   = 2
	BromineParserSPLITKEY = 3
	BromineParserWS       = 4
)

// BromineParser rules.
const (
	BromineParserRULE_expression = 0
	BromineParserRULE_bromine    = 1
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
	p.RuleIndex = BromineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BromineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bromine() IBromineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBromineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBromineContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BromineParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BromineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BromineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BromineParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BromineParserRULE_expression)

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
		p.Bromine()
	}
	{
		p.SetState(5)
		p.Match(BromineParserEOF)
	}

	return localctx
}

// IBromineContext is an interface to support dynamic dispatch.
type IBromineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBromineContext differentiates from other interfaces.
	IsBromineContext()
}

type BromineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBromineContext() *BromineContext {
	var p = new(BromineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BromineParserRULE_bromine
	return p
}

func (*BromineContext) IsBromineContext() {}

func NewBromineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BromineContext {
	var p = new(BromineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BromineParserRULE_bromine

	return p
}

func (s *BromineContext) GetParser() antlr.Parser { return s.parser }

func (s *BromineContext) BROMINE() antlr.TerminalNode {
	return s.GetToken(BromineParserBROMINE, 0)
}

func (s *BromineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BromineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BromineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BromineListener); ok {
		listenerT.EnterBromine(s)
	}
}

func (s *BromineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BromineListener); ok {
		listenerT.ExitBromine(s)
	}
}

func (p *BromineParser) Bromine() (localctx IBromineContext) {
	this := p
	_ = this

	localctx = NewBromineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BromineParserRULE_bromine)

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
		p.Match(BromineParserBROMINE)
	}

	return localctx
}
