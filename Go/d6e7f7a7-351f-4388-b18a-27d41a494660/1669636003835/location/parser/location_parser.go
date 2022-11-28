// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669636003835/location.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // location

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

type locationParser struct {
	*antlr.BaseParser
}

var locationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func locationParserInit() {
	staticData := &locationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LOCATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "location",
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

// locationParserInit initializes any static state used to implement locationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewlocationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LocationParserInit() {
	staticData := &locationParserStaticData
	staticData.once.Do(locationParserInit)
}

// NewlocationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewlocationParser(input antlr.TokenStream) *locationParser {
	LocationParserInit()
	this := new(locationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &locationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "location.g4"

	return this
}

// locationParser tokens.
const (
	locationParserEOF      = antlr.TokenEOF
	locationParserLOCATION = 1
	locationParserOWNKEY   = 2
	locationParserSPLITKEY = 3
	locationParserWS       = 4
)

// locationParser rules.
const (
	locationParserRULE_expression = 0
	locationParserRULE_location   = 1
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
	p.RuleIndex = locationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = locationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Location() ILocationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILocationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILocationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(locationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(locationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(locationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *locationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, locationParserRULE_expression)

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
		p.Location()
	}
	{
		p.SetState(5)
		p.Match(locationParserEOF)
	}

	return localctx
}

// ILocationContext is an interface to support dynamic dispatch.
type ILocationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocationContext differentiates from other interfaces.
	IsLocationContext()
}

type LocationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocationContext() *LocationContext {
	var p = new(LocationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = locationParserRULE_location
	return p
}

func (*LocationContext) IsLocationContext() {}

func NewLocationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocationContext {
	var p = new(LocationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = locationParserRULE_location

	return p
}

func (s *LocationContext) GetParser() antlr.Parser { return s.parser }

func (s *LocationContext) LOCATION() antlr.TerminalNode {
	return s.GetToken(locationParserLOCATION, 0)
}

func (s *LocationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(locationListener); ok {
		listenerT.EnterLocation(s)
	}
}

func (s *LocationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(locationListener); ok {
		listenerT.ExitLocation(s)
	}
}

func (p *locationParser) Location() (localctx ILocationContext) {
	this := p
	_ = this

	localctx = NewLocationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, locationParserRULE_location)

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
		p.Match(locationParserLOCATION)
	}

	return localctx
}
