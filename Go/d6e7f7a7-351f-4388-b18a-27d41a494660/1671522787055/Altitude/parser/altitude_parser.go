// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671522787055/Altitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Altitude

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

type AltitudeParser struct {
	*antlr.BaseParser
}

var altitudeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func altitudeParserInit() {
	staticData := &altitudeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ALTITUDE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "altitude",
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

// AltitudeParserInit initializes any static state used to implement AltitudeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAltitudeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AltitudeParserInit() {
	staticData := &altitudeParserStaticData
	staticData.once.Do(altitudeParserInit)
}

// NewAltitudeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAltitudeParser(input antlr.TokenStream) *AltitudeParser {
	AltitudeParserInit()
	this := new(AltitudeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &altitudeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Altitude.g4"

	return this
}

// AltitudeParser tokens.
const (
	AltitudeParserEOF      = antlr.TokenEOF
	AltitudeParserALTITUDE = 1
	AltitudeParserOWNKEY   = 2
	AltitudeParserSPLITKEY = 3
	AltitudeParserWS       = 4
)

// AltitudeParser rules.
const (
	AltitudeParserRULE_expression = 0
	AltitudeParserRULE_altitude   = 1
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
	p.RuleIndex = AltitudeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AltitudeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Altitude() IAltitudeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAltitudeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAltitudeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AltitudeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AltitudeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AltitudeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AltitudeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AltitudeParserRULE_expression)

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
		p.Altitude()
	}
	{
		p.SetState(5)
		p.Match(AltitudeParserEOF)
	}

	return localctx
}

// IAltitudeContext is an interface to support dynamic dispatch.
type IAltitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAltitudeContext differentiates from other interfaces.
	IsAltitudeContext()
}

type AltitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAltitudeContext() *AltitudeContext {
	var p = new(AltitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AltitudeParserRULE_altitude
	return p
}

func (*AltitudeContext) IsAltitudeContext() {}

func NewAltitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AltitudeContext {
	var p = new(AltitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AltitudeParserRULE_altitude

	return p
}

func (s *AltitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *AltitudeContext) ALTITUDE() antlr.TerminalNode {
	return s.GetToken(AltitudeParserALTITUDE, 0)
}

func (s *AltitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AltitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AltitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AltitudeListener); ok {
		listenerT.EnterAltitude(s)
	}
}

func (s *AltitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AltitudeListener); ok {
		listenerT.ExitAltitude(s)
	}
}

func (p *AltitudeParser) Altitude() (localctx IAltitudeContext) {
	this := p
	_ = this

	localctx = NewAltitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AltitudeParserRULE_altitude)

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
		p.Match(AltitudeParserALTITUDE)
	}

	return localctx
}
