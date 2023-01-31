// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675182829473/AvgGlucoseLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AvgGlucoseLevel

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

type AvgGlucoseLevelParser struct {
	*antlr.BaseParser
}

var avgglucoselevelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func avgglucoselevelParserInit() {
	staticData := &avgglucoselevelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AVGGLUCOSELEVEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "avgglucoselevel",
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

// AvgGlucoseLevelParserInit initializes any static state used to implement AvgGlucoseLevelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAvgGlucoseLevelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AvgGlucoseLevelParserInit() {
	staticData := &avgglucoselevelParserStaticData
	staticData.once.Do(avgglucoselevelParserInit)
}

// NewAvgGlucoseLevelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAvgGlucoseLevelParser(input antlr.TokenStream) *AvgGlucoseLevelParser {
	AvgGlucoseLevelParserInit()
	this := new(AvgGlucoseLevelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &avgglucoselevelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AvgGlucoseLevel.g4"

	return this
}

// AvgGlucoseLevelParser tokens.
const (
	AvgGlucoseLevelParserEOF             = antlr.TokenEOF
	AvgGlucoseLevelParserAVGGLUCOSELEVEL = 1
	AvgGlucoseLevelParserOWNKEY          = 2
	AvgGlucoseLevelParserSPLITKEY        = 3
	AvgGlucoseLevelParserWS              = 4
)

// AvgGlucoseLevelParser rules.
const (
	AvgGlucoseLevelParserRULE_expression      = 0
	AvgGlucoseLevelParserRULE_avgglucoselevel = 1
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
	p.RuleIndex = AvgGlucoseLevelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AvgGlucoseLevelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Avgglucoselevel() IAvgglucoselevelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAvgglucoselevelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAvgglucoselevelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AvgGlucoseLevelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgGlucoseLevelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgGlucoseLevelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AvgGlucoseLevelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AvgGlucoseLevelParserRULE_expression)

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
		p.Avgglucoselevel()
	}
	{
		p.SetState(5)
		p.Match(AvgGlucoseLevelParserEOF)
	}

	return localctx
}

// IAvgglucoselevelContext is an interface to support dynamic dispatch.
type IAvgglucoselevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAvgglucoselevelContext differentiates from other interfaces.
	IsAvgglucoselevelContext()
}

type AvgglucoselevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvgglucoselevelContext() *AvgglucoselevelContext {
	var p = new(AvgglucoselevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AvgGlucoseLevelParserRULE_avgglucoselevel
	return p
}

func (*AvgglucoselevelContext) IsAvgglucoselevelContext() {}

func NewAvgglucoselevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvgglucoselevelContext {
	var p = new(AvgglucoselevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AvgGlucoseLevelParserRULE_avgglucoselevel

	return p
}

func (s *AvgglucoselevelContext) GetParser() antlr.Parser { return s.parser }

func (s *AvgglucoselevelContext) AVGGLUCOSELEVEL() antlr.TerminalNode {
	return s.GetToken(AvgGlucoseLevelParserAVGGLUCOSELEVEL, 0)
}

func (s *AvgglucoselevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvgglucoselevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AvgglucoselevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgGlucoseLevelListener); ok {
		listenerT.EnterAvgglucoselevel(s)
	}
}

func (s *AvgglucoselevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AvgGlucoseLevelListener); ok {
		listenerT.ExitAvgglucoselevel(s)
	}
}

func (p *AvgGlucoseLevelParser) Avgglucoselevel() (localctx IAvgglucoselevelContext) {
	this := p
	_ = this

	localctx = NewAvgglucoselevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AvgGlucoseLevelParserRULE_avgglucoselevel)

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
		p.Match(AvgGlucoseLevelParserAVGGLUCOSELEVEL)
	}

	return localctx
}
