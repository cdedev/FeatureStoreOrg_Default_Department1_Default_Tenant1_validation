// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121312157/Sheep.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sheep

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

type SheepParser struct {
	*antlr.BaseParser
}

var sheepParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sheepParserInit() {
	staticData := &sheepParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SHEEP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sheep",
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

// SheepParserInit initializes any static state used to implement SheepParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSheepParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SheepParserInit() {
	staticData := &sheepParserStaticData
	staticData.once.Do(sheepParserInit)
}

// NewSheepParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSheepParser(input antlr.TokenStream) *SheepParser {
	SheepParserInit()
	this := new(SheepParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sheepParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sheep.g4"

	return this
}

// SheepParser tokens.
const (
	SheepParserEOF      = antlr.TokenEOF
	SheepParserSHEEP    = 1
	SheepParserOWNKEY   = 2
	SheepParserSPLITKEY = 3
	SheepParserWS       = 4
)

// SheepParser rules.
const (
	SheepParserRULE_expression = 0
	SheepParserRULE_sheep      = 1
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
	p.RuleIndex = SheepParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SheepParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sheep() ISheepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISheepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISheepContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SheepParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SheepListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SheepListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SheepParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SheepParserRULE_expression)

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
		p.Sheep()
	}
	{
		p.SetState(5)
		p.Match(SheepParserEOF)
	}

	return localctx
}

// ISheepContext is an interface to support dynamic dispatch.
type ISheepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSheepContext differentiates from other interfaces.
	IsSheepContext()
}

type SheepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySheepContext() *SheepContext {
	var p = new(SheepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SheepParserRULE_sheep
	return p
}

func (*SheepContext) IsSheepContext() {}

func NewSheepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SheepContext {
	var p = new(SheepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SheepParserRULE_sheep

	return p
}

func (s *SheepContext) GetParser() antlr.Parser { return s.parser }

func (s *SheepContext) SHEEP() antlr.TerminalNode {
	return s.GetToken(SheepParserSHEEP, 0)
}

func (s *SheepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SheepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SheepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SheepListener); ok {
		listenerT.EnterSheep(s)
	}
}

func (s *SheepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SheepListener); ok {
		listenerT.ExitSheep(s)
	}
}

func (p *SheepParser) Sheep() (localctx ISheepContext) {
	this := p
	_ = this

	localctx = NewSheepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SheepParserRULE_sheep)

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
		p.Match(SheepParserSHEEP)
	}

	return localctx
}
