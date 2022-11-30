// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789488857/Magnitude.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Magnitude

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

type MagnitudeParser struct {
	*antlr.BaseParser
}

var magnitudeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func magnitudeParserInit() {
	staticData := &magnitudeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MAGNITUDE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "magnitude",
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

// MagnitudeParserInit initializes any static state used to implement MagnitudeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMagnitudeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MagnitudeParserInit() {
	staticData := &magnitudeParserStaticData
	staticData.once.Do(magnitudeParserInit)
}

// NewMagnitudeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMagnitudeParser(input antlr.TokenStream) *MagnitudeParser {
	MagnitudeParserInit()
	this := new(MagnitudeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &magnitudeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Magnitude.g4"

	return this
}

// MagnitudeParser tokens.
const (
	MagnitudeParserEOF       = antlr.TokenEOF
	MagnitudeParserMAGNITUDE = 1
	MagnitudeParserOWNKEY    = 2
	MagnitudeParserSPLITKEY  = 3
	MagnitudeParserWS        = 4
)

// MagnitudeParser rules.
const (
	MagnitudeParserRULE_expression = 0
	MagnitudeParserRULE_magnitude  = 1
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
	p.RuleIndex = MagnitudeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagnitudeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Magnitude() IMagnitudeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMagnitudeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMagnitudeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MagnitudeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagnitudeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagnitudeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MagnitudeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MagnitudeParserRULE_expression)

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
		p.Magnitude()
	}
	{
		p.SetState(5)
		p.Match(MagnitudeParserEOF)
	}

	return localctx
}

// IMagnitudeContext is an interface to support dynamic dispatch.
type IMagnitudeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMagnitudeContext differentiates from other interfaces.
	IsMagnitudeContext()
}

type MagnitudeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMagnitudeContext() *MagnitudeContext {
	var p = new(MagnitudeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MagnitudeParserRULE_magnitude
	return p
}

func (*MagnitudeContext) IsMagnitudeContext() {}

func NewMagnitudeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MagnitudeContext {
	var p = new(MagnitudeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MagnitudeParserRULE_magnitude

	return p
}

func (s *MagnitudeContext) GetParser() antlr.Parser { return s.parser }

func (s *MagnitudeContext) MAGNITUDE() antlr.TerminalNode {
	return s.GetToken(MagnitudeParserMAGNITUDE, 0)
}

func (s *MagnitudeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MagnitudeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MagnitudeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagnitudeListener); ok {
		listenerT.EnterMagnitude(s)
	}
}

func (s *MagnitudeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MagnitudeListener); ok {
		listenerT.ExitMagnitude(s)
	}
}

func (p *MagnitudeParser) Magnitude() (localctx IMagnitudeContext) {
	this := p
	_ = this

	localctx = NewMagnitudeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MagnitudeParserRULE_magnitude)

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
		p.Match(MagnitudeParserMAGNITUDE)
	}

	return localctx
}
