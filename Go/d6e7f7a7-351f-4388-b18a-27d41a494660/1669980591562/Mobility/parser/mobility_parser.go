// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669980591562/Mobility.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mobility

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

type MobilityParser struct {
	*antlr.BaseParser
}

var mobilityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mobilityParserInit() {
	staticData := &mobilityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MOBILITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mobility",
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

// MobilityParserInit initializes any static state used to implement MobilityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMobilityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MobilityParserInit() {
	staticData := &mobilityParserStaticData
	staticData.once.Do(mobilityParserInit)
}

// NewMobilityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMobilityParser(input antlr.TokenStream) *MobilityParser {
	MobilityParserInit()
	this := new(MobilityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mobilityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mobility.g4"

	return this
}

// MobilityParser tokens.
const (
	MobilityParserEOF      = antlr.TokenEOF
	MobilityParserMOBILITY = 1
	MobilityParserOWNKEY   = 2
	MobilityParserSPLITKEY = 3
	MobilityParserWS       = 4
)

// MobilityParser rules.
const (
	MobilityParserRULE_expression = 0
	MobilityParserRULE_mobility   = 1
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
	p.RuleIndex = MobilityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MobilityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mobility() IMobilityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMobilityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMobilityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MobilityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobilityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobilityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MobilityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MobilityParserRULE_expression)

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
		p.Mobility()
	}
	{
		p.SetState(5)
		p.Match(MobilityParserEOF)
	}

	return localctx
}

// IMobilityContext is an interface to support dynamic dispatch.
type IMobilityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMobilityContext differentiates from other interfaces.
	IsMobilityContext()
}

type MobilityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMobilityContext() *MobilityContext {
	var p = new(MobilityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MobilityParserRULE_mobility
	return p
}

func (*MobilityContext) IsMobilityContext() {}

func NewMobilityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MobilityContext {
	var p = new(MobilityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MobilityParserRULE_mobility

	return p
}

func (s *MobilityContext) GetParser() antlr.Parser { return s.parser }

func (s *MobilityContext) MOBILITY() antlr.TerminalNode {
	return s.GetToken(MobilityParserMOBILITY, 0)
}

func (s *MobilityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MobilityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MobilityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobilityListener); ok {
		listenerT.EnterMobility(s)
	}
}

func (s *MobilityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobilityListener); ok {
		listenerT.ExitMobility(s)
	}
}

func (p *MobilityParser) Mobility() (localctx IMobilityContext) {
	this := p
	_ = this

	localctx = NewMobilityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MobilityParserRULE_mobility)

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
		p.Match(MobilityParserMOBILITY)
	}

	return localctx
}
