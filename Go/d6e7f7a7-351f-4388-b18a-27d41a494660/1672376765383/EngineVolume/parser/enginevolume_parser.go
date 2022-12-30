// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376765383/EngineVolume.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EngineVolume

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

type EngineVolumeParser struct {
	*antlr.BaseParser
}

var enginevolumeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func enginevolumeParserInit() {
	staticData := &enginevolumeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ENGINEVOLUME", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "enginevolume",
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

// EngineVolumeParserInit initializes any static state used to implement EngineVolumeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEngineVolumeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EngineVolumeParserInit() {
	staticData := &enginevolumeParserStaticData
	staticData.once.Do(enginevolumeParserInit)
}

// NewEngineVolumeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEngineVolumeParser(input antlr.TokenStream) *EngineVolumeParser {
	EngineVolumeParserInit()
	this := new(EngineVolumeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &enginevolumeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "EngineVolume.g4"

	return this
}

// EngineVolumeParser tokens.
const (
	EngineVolumeParserEOF          = antlr.TokenEOF
	EngineVolumeParserENGINEVOLUME = 1
	EngineVolumeParserOWNKEY       = 2
	EngineVolumeParserSPLITKEY     = 3
	EngineVolumeParserWS           = 4
)

// EngineVolumeParser rules.
const (
	EngineVolumeParserRULE_expression   = 0
	EngineVolumeParserRULE_enginevolume = 1
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
	p.RuleIndex = EngineVolumeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EngineVolumeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Enginevolume() IEnginevolumeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnginevolumeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnginevolumeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EngineVolumeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineVolumeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineVolumeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EngineVolumeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EngineVolumeParserRULE_expression)

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
		p.Enginevolume()
	}
	{
		p.SetState(5)
		p.Match(EngineVolumeParserEOF)
	}

	return localctx
}

// IEnginevolumeContext is an interface to support dynamic dispatch.
type IEnginevolumeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnginevolumeContext differentiates from other interfaces.
	IsEnginevolumeContext()
}

type EnginevolumeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnginevolumeContext() *EnginevolumeContext {
	var p = new(EnginevolumeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EngineVolumeParserRULE_enginevolume
	return p
}

func (*EnginevolumeContext) IsEnginevolumeContext() {}

func NewEnginevolumeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnginevolumeContext {
	var p = new(EnginevolumeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EngineVolumeParserRULE_enginevolume

	return p
}

func (s *EnginevolumeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnginevolumeContext) ENGINEVOLUME() antlr.TerminalNode {
	return s.GetToken(EngineVolumeParserENGINEVOLUME, 0)
}

func (s *EnginevolumeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnginevolumeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnginevolumeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineVolumeListener); ok {
		listenerT.EnterEnginevolume(s)
	}
}

func (s *EnginevolumeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EngineVolumeListener); ok {
		listenerT.ExitEnginevolume(s)
	}
}

func (p *EngineVolumeParser) Enginevolume() (localctx IEnginevolumeContext) {
	this := p
	_ = this

	localctx = NewEnginevolumeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EngineVolumeParserRULE_enginevolume)

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
		p.Match(EngineVolumeParserENGINEVOLUME)
	}

	return localctx
}
