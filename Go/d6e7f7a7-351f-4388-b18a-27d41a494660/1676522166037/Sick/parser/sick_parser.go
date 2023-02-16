// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522166037/Sick.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sick

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

type SickParser struct {
	*antlr.BaseParser
}

var sickParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sickParserInit() {
	staticData := &sickParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SICK", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sick",
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

// SickParserInit initializes any static state used to implement SickParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSickParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SickParserInit() {
	staticData := &sickParserStaticData
	staticData.once.Do(sickParserInit)
}

// NewSickParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSickParser(input antlr.TokenStream) *SickParser {
	SickParserInit()
	this := new(SickParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sickParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sick.g4"

	return this
}

// SickParser tokens.
const (
	SickParserEOF      = antlr.TokenEOF
	SickParserSICK     = 1
	SickParserOWNKEY   = 2
	SickParserSPLITKEY = 3
	SickParserWS       = 4
)

// SickParser rules.
const (
	SickParserRULE_expression = 0
	SickParserRULE_sick       = 1
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
	p.RuleIndex = SickParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SickParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sick() ISickContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISickContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISickContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SickParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SickListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SickListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SickParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SickParserRULE_expression)

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
		p.Sick()
	}
	{
		p.SetState(5)
		p.Match(SickParserEOF)
	}

	return localctx
}

// ISickContext is an interface to support dynamic dispatch.
type ISickContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSickContext differentiates from other interfaces.
	IsSickContext()
}

type SickContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySickContext() *SickContext {
	var p = new(SickContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SickParserRULE_sick
	return p
}

func (*SickContext) IsSickContext() {}

func NewSickContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SickContext {
	var p = new(SickContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SickParserRULE_sick

	return p
}

func (s *SickContext) GetParser() antlr.Parser { return s.parser }

func (s *SickContext) SICK() antlr.TerminalNode {
	return s.GetToken(SickParserSICK, 0)
}

func (s *SickContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SickContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SickContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SickListener); ok {
		listenerT.EnterSick(s)
	}
}

func (s *SickContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SickListener); ok {
		listenerT.ExitSick(s)
	}
}

func (p *SickParser) Sick() (localctx ISickContext) {
	this := p
	_ = this

	localctx = NewSickContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SickParserRULE_sick)

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
		p.Match(SickParserSICK)
	}

	return localctx
}
