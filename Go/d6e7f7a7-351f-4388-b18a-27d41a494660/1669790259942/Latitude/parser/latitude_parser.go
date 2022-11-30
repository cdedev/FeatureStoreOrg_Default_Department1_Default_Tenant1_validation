// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669790259942/Latitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Latitude

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

type LatitudeParser struct {
	*antlr.BaseParser
}

var latitudeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func latitudeParserInit() {
	staticData := &latitudeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LATITUDE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "latitude",
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

// LatitudeParserInit initializes any static state used to implement LatitudeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLatitudeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LatitudeParserInit() {
	staticData := &latitudeParserStaticData
	staticData.once.Do(latitudeParserInit)
}

// NewLatitudeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLatitudeParser(input antlr.TokenStream) *LatitudeParser {
	LatitudeParserInit()
	this := new(LatitudeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &latitudeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Latitude.g4"

	return this
}

// LatitudeParser tokens.
const (
	LatitudeParserEOF      = antlr.TokenEOF
	LatitudeParserLATITUDE = 1
	LatitudeParserOWNKEY   = 2
	LatitudeParserSPLITKEY = 3
	LatitudeParserWS       = 4
)

// LatitudeParser rules.
const (
	LatitudeParserRULE_expression = 0
	LatitudeParserRULE_latitude   = 1
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
	p.RuleIndex = LatitudeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatitudeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Latitude() ILatitudeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILatitudeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILatitudeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LatitudeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatitudeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatitudeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LatitudeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LatitudeParserRULE_expression)

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
		p.Latitude()
	}
	{
		p.SetState(5)
		p.Match(LatitudeParserEOF)
	}

	return localctx
}

// ILatitudeContext is an interface to support dynamic dispatch.
type ILatitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLatitudeContext differentiates from other interfaces.
	IsLatitudeContext()
}

type LatitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLatitudeContext() *LatitudeContext {
	var p = new(LatitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LatitudeParserRULE_latitude
	return p
}

func (*LatitudeContext) IsLatitudeContext() {}

func NewLatitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LatitudeContext {
	var p = new(LatitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LatitudeParserRULE_latitude

	return p
}

func (s *LatitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *LatitudeContext) LATITUDE() antlr.TerminalNode {
	return s.GetToken(LatitudeParserLATITUDE, 0)
}

func (s *LatitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LatitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LatitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatitudeListener); ok {
		listenerT.EnterLatitude(s)
	}
}

func (s *LatitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LatitudeListener); ok {
		listenerT.ExitLatitude(s)
	}
}

func (p *LatitudeParser) Latitude() (localctx ILatitudeContext) {
	this := p
	_ = this

	localctx = NewLatitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LatitudeParserRULE_latitude)

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
		p.Match(LatitudeParserLATITUDE)
	}

	return localctx
}
