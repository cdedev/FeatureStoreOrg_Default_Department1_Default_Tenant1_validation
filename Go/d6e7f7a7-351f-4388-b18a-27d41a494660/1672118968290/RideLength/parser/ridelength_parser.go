// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118968290/RideLength.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RideLength

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

type RideLengthParser struct {
	*antlr.BaseParser
}

var ridelengthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ridelengthParserInit() {
	staticData := &ridelengthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RIDELENGTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ridelength",
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

// RideLengthParserInit initializes any static state used to implement RideLengthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRideLengthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RideLengthParserInit() {
	staticData := &ridelengthParserStaticData
	staticData.once.Do(ridelengthParserInit)
}

// NewRideLengthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRideLengthParser(input antlr.TokenStream) *RideLengthParser {
	RideLengthParserInit()
	this := new(RideLengthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ridelengthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "RideLength.g4"

	return this
}

// RideLengthParser tokens.
const (
	RideLengthParserEOF        = antlr.TokenEOF
	RideLengthParserRIDELENGTH = 1
	RideLengthParserOWNKEY     = 2
	RideLengthParserSPLITKEY   = 3
	RideLengthParserWS         = 4
)

// RideLengthParser rules.
const (
	RideLengthParserRULE_expression = 0
	RideLengthParserRULE_ridelength = 1
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
	p.RuleIndex = RideLengthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RideLengthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ridelength() IRidelengthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRidelengthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRidelengthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RideLengthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RideLengthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RideLengthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RideLengthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RideLengthParserRULE_expression)

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
		p.Ridelength()
	}
	{
		p.SetState(5)
		p.Match(RideLengthParserEOF)
	}

	return localctx
}

// IRidelengthContext is an interface to support dynamic dispatch.
type IRidelengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRidelengthContext differentiates from other interfaces.
	IsRidelengthContext()
}

type RidelengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRidelengthContext() *RidelengthContext {
	var p = new(RidelengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RideLengthParserRULE_ridelength
	return p
}

func (*RidelengthContext) IsRidelengthContext() {}

func NewRidelengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RidelengthContext {
	var p = new(RidelengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RideLengthParserRULE_ridelength

	return p
}

func (s *RidelengthContext) GetParser() antlr.Parser { return s.parser }

func (s *RidelengthContext) RIDELENGTH() antlr.TerminalNode {
	return s.GetToken(RideLengthParserRIDELENGTH, 0)
}

func (s *RidelengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RidelengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RidelengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RideLengthListener); ok {
		listenerT.EnterRidelength(s)
	}
}

func (s *RidelengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RideLengthListener); ok {
		listenerT.ExitRidelength(s)
	}
}

func (p *RideLengthParser) Ridelength() (localctx IRidelengthContext) {
	this := p
	_ = this

	localctx = NewRidelengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RideLengthParserRULE_ridelength)

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
		p.Match(RideLengthParserRIDELENGTH)
	}

	return localctx
}
