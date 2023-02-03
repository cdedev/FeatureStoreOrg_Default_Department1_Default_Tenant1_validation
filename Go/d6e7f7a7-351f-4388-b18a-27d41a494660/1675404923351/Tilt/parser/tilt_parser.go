// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675404923351/Tilt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tilt

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

type TiltParser struct {
	*antlr.BaseParser
}

var tiltParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tiltParserInit() {
	staticData := &tiltParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TILT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "tilt",
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

// TiltParserInit initializes any static state used to implement TiltParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTiltParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TiltParserInit() {
	staticData := &tiltParserStaticData
	staticData.once.Do(tiltParserInit)
}

// NewTiltParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTiltParser(input antlr.TokenStream) *TiltParser {
	TiltParserInit()
	this := new(TiltParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tiltParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tilt.g4"

	return this
}

// TiltParser tokens.
const (
	TiltParserEOF      = antlr.TokenEOF
	TiltParserTILT     = 1
	TiltParserOWNKEY   = 2
	TiltParserSPLITKEY = 3
	TiltParserWS       = 4
)

// TiltParser rules.
const (
	TiltParserRULE_expression = 0
	TiltParserRULE_tilt       = 1
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
	p.RuleIndex = TiltParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TiltParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Tilt() ITiltContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITiltContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITiltContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TiltParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TiltListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TiltListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TiltParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TiltParserRULE_expression)

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
		p.Tilt()
	}
	{
		p.SetState(5)
		p.Match(TiltParserEOF)
	}

	return localctx
}

// ITiltContext is an interface to support dynamic dispatch.
type ITiltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTiltContext differentiates from other interfaces.
	IsTiltContext()
}

type TiltContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTiltContext() *TiltContext {
	var p = new(TiltContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TiltParserRULE_tilt
	return p
}

func (*TiltContext) IsTiltContext() {}

func NewTiltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TiltContext {
	var p = new(TiltContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TiltParserRULE_tilt

	return p
}

func (s *TiltContext) GetParser() antlr.Parser { return s.parser }

func (s *TiltContext) TILT() antlr.TerminalNode {
	return s.GetToken(TiltParserTILT, 0)
}

func (s *TiltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TiltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TiltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TiltListener); ok {
		listenerT.EnterTilt(s)
	}
}

func (s *TiltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TiltListener); ok {
		listenerT.ExitTilt(s)
	}
}

func (p *TiltParser) Tilt() (localctx ITiltContext) {
	this := p
	_ = this

	localctx = NewTiltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TiltParserRULE_tilt)

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
		p.Match(TiltParserTILT)
	}

	return localctx
}
