// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675788016809/PressureLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PressureLevel

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

type PressureLevelParser struct {
	*antlr.BaseParser
}

var pressurelevelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pressurelevelParserInit() {
	staticData := &pressurelevelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PRESSURELEVEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pressurelevel",
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

// PressureLevelParserInit initializes any static state used to implement PressureLevelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPressureLevelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PressureLevelParserInit() {
	staticData := &pressurelevelParserStaticData
	staticData.once.Do(pressurelevelParserInit)
}

// NewPressureLevelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPressureLevelParser(input antlr.TokenStream) *PressureLevelParser {
	PressureLevelParserInit()
	this := new(PressureLevelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pressurelevelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PressureLevel.g4"

	return this
}

// PressureLevelParser tokens.
const (
	PressureLevelParserEOF           = antlr.TokenEOF
	PressureLevelParserPRESSURELEVEL = 1
	PressureLevelParserOWNKEY        = 2
	PressureLevelParserSPLITKEY      = 3
	PressureLevelParserWS            = 4
)

// PressureLevelParser rules.
const (
	PressureLevelParserRULE_expression    = 0
	PressureLevelParserRULE_pressurelevel = 1
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
	p.RuleIndex = PressureLevelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PressureLevelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pressurelevel() IPressurelevelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPressurelevelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPressurelevelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PressureLevelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureLevelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureLevelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PressureLevelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PressureLevelParserRULE_expression)

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
		p.Pressurelevel()
	}
	{
		p.SetState(5)
		p.Match(PressureLevelParserEOF)
	}

	return localctx
}

// IPressurelevelContext is an interface to support dynamic dispatch.
type IPressurelevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPressurelevelContext differentiates from other interfaces.
	IsPressurelevelContext()
}

type PressurelevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPressurelevelContext() *PressurelevelContext {
	var p = new(PressurelevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PressureLevelParserRULE_pressurelevel
	return p
}

func (*PressurelevelContext) IsPressurelevelContext() {}

func NewPressurelevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PressurelevelContext {
	var p = new(PressurelevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PressureLevelParserRULE_pressurelevel

	return p
}

func (s *PressurelevelContext) GetParser() antlr.Parser { return s.parser }

func (s *PressurelevelContext) PRESSURELEVEL() antlr.TerminalNode {
	return s.GetToken(PressureLevelParserPRESSURELEVEL, 0)
}

func (s *PressurelevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PressurelevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PressurelevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureLevelListener); ok {
		listenerT.EnterPressurelevel(s)
	}
}

func (s *PressurelevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PressureLevelListener); ok {
		listenerT.ExitPressurelevel(s)
	}
}

func (p *PressureLevelParser) Pressurelevel() (localctx IPressurelevelContext) {
	this := p
	_ = this

	localctx = NewPressurelevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PressureLevelParserRULE_pressurelevel)

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
		p.Match(PressureLevelParserPRESSURELEVEL)
	}

	return localctx
}
