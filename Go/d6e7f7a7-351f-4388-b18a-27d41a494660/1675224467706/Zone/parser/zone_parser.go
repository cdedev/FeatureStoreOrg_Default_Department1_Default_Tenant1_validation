// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675224467706/Zone.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Zone

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

type ZoneParser struct {
	*antlr.BaseParser
}

var zoneParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zoneParserInit() {
	staticData := &zoneParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ZONE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "zone",
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

// ZoneParserInit initializes any static state used to implement ZoneParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewZoneParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZoneParserInit() {
	staticData := &zoneParserStaticData
	staticData.once.Do(zoneParserInit)
}

// NewZoneParser produces a new parser instance for the optional input antlr.TokenStream.
func NewZoneParser(input antlr.TokenStream) *ZoneParser {
	ZoneParserInit()
	this := new(ZoneParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &zoneParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Zone.g4"

	return this
}

// ZoneParser tokens.
const (
	ZoneParserEOF      = antlr.TokenEOF
	ZoneParserZONE     = 1
	ZoneParserOWNKEY   = 2
	ZoneParserSPLITKEY = 3
	ZoneParserWS       = 4
)

// ZoneParser rules.
const (
	ZoneParserRULE_expression = 0
	ZoneParserRULE_zone       = 1
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
	p.RuleIndex = ZoneParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZoneParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Zone() IZoneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZoneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZoneContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZoneParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZoneListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZoneListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ZoneParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZoneParserRULE_expression)

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
		p.Zone()
	}
	{
		p.SetState(5)
		p.Match(ZoneParserEOF)
	}

	return localctx
}

// IZoneContext is an interface to support dynamic dispatch.
type IZoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsZoneContext differentiates from other interfaces.
	IsZoneContext()
}

type ZoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZoneContext() *ZoneContext {
	var p = new(ZoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ZoneParserRULE_zone
	return p
}

func (*ZoneContext) IsZoneContext() {}

func NewZoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZoneContext {
	var p = new(ZoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZoneParserRULE_zone

	return p
}

func (s *ZoneContext) GetParser() antlr.Parser { return s.parser }

func (s *ZoneContext) ZONE() antlr.TerminalNode {
	return s.GetToken(ZoneParserZONE, 0)
}

func (s *ZoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZoneListener); ok {
		listenerT.EnterZone(s)
	}
}

func (s *ZoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZoneListener); ok {
		listenerT.ExitZone(s)
	}
}

func (p *ZoneParser) Zone() (localctx IZoneContext) {
	this := p
	_ = this

	localctx = NewZoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZoneParserRULE_zone)

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
		p.Match(ZoneParserZONE)
	}

	return localctx
}
