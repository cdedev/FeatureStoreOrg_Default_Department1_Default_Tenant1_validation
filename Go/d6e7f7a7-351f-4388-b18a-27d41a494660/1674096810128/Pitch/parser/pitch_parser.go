// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096810128/Pitch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pitch

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

type PitchParser struct {
	*antlr.BaseParser
}

var pitchParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pitchParserInit() {
	staticData := &pitchParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PITCH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pitch",
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

// PitchParserInit initializes any static state used to implement PitchParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPitchParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PitchParserInit() {
	staticData := &pitchParserStaticData
	staticData.once.Do(pitchParserInit)
}

// NewPitchParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPitchParser(input antlr.TokenStream) *PitchParser {
	PitchParserInit()
	this := new(PitchParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pitchParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Pitch.g4"

	return this
}

// PitchParser tokens.
const (
	PitchParserEOF      = antlr.TokenEOF
	PitchParserPITCH    = 1
	PitchParserOWNKEY   = 2
	PitchParserSPLITKEY = 3
	PitchParserWS       = 4
)

// PitchParser rules.
const (
	PitchParserRULE_expression = 0
	PitchParserRULE_pitch      = 1
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
	p.RuleIndex = PitchParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PitchParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pitch() IPitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPitchContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PitchParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PitchListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PitchListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PitchParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PitchParserRULE_expression)

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
		p.Pitch()
	}
	{
		p.SetState(5)
		p.Match(PitchParserEOF)
	}

	return localctx
}

// IPitchContext is an interface to support dynamic dispatch.
type IPitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPitchContext differentiates from other interfaces.
	IsPitchContext()
}

type PitchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPitchContext() *PitchContext {
	var p = new(PitchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PitchParserRULE_pitch
	return p
}

func (*PitchContext) IsPitchContext() {}

func NewPitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PitchContext {
	var p = new(PitchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PitchParserRULE_pitch

	return p
}

func (s *PitchContext) GetParser() antlr.Parser { return s.parser }

func (s *PitchContext) PITCH() antlr.TerminalNode {
	return s.GetToken(PitchParserPITCH, 0)
}

func (s *PitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PitchListener); ok {
		listenerT.EnterPitch(s)
	}
}

func (s *PitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PitchListener); ok {
		listenerT.ExitPitch(s)
	}
}

func (p *PitchParser) Pitch() (localctx IPitchContext) {
	this := p
	_ = this

	localctx = NewPitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PitchParserRULE_pitch)

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
		p.Match(PitchParserPITCH)
	}

	return localctx
}
