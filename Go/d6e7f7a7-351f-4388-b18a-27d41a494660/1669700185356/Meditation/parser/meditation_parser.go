// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700185356/Meditation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Meditation

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

type MeditationParser struct {
	*antlr.BaseParser
}

var meditationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func meditationParserInit() {
	staticData := &meditationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MEDITATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "meditation",
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

// MeditationParserInit initializes any static state used to implement MeditationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMeditationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MeditationParserInit() {
	staticData := &meditationParserStaticData
	staticData.once.Do(meditationParserInit)
}

// NewMeditationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMeditationParser(input antlr.TokenStream) *MeditationParser {
	MeditationParserInit()
	this := new(MeditationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &meditationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Meditation.g4"

	return this
}

// MeditationParser tokens.
const (
	MeditationParserEOF        = antlr.TokenEOF
	MeditationParserMEDITATION = 1
	MeditationParserOWNKEY     = 2
	MeditationParserSPLITKEY   = 3
	MeditationParserWS         = 4
)

// MeditationParser rules.
const (
	MeditationParserRULE_expression = 0
	MeditationParserRULE_meditation = 1
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
	p.RuleIndex = MeditationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MeditationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Meditation() IMeditationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMeditationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMeditationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MeditationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeditationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeditationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MeditationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MeditationParserRULE_expression)

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
		p.Meditation()
	}
	{
		p.SetState(5)
		p.Match(MeditationParserEOF)
	}

	return localctx
}

// IMeditationContext is an interface to support dynamic dispatch.
type IMeditationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeditationContext differentiates from other interfaces.
	IsMeditationContext()
}

type MeditationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeditationContext() *MeditationContext {
	var p = new(MeditationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MeditationParserRULE_meditation
	return p
}

func (*MeditationContext) IsMeditationContext() {}

func NewMeditationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeditationContext {
	var p = new(MeditationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MeditationParserRULE_meditation

	return p
}

func (s *MeditationContext) GetParser() antlr.Parser { return s.parser }

func (s *MeditationContext) MEDITATION() antlr.TerminalNode {
	return s.GetToken(MeditationParserMEDITATION, 0)
}

func (s *MeditationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeditationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeditationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeditationListener); ok {
		listenerT.EnterMeditation(s)
	}
}

func (s *MeditationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MeditationListener); ok {
		listenerT.ExitMeditation(s)
	}
}

func (p *MeditationParser) Meditation() (localctx IMeditationContext) {
	this := p
	_ = this

	localctx = NewMeditationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MeditationParserRULE_meditation)

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
		p.Match(MeditationParserMEDITATION)
	}

	return localctx
}
