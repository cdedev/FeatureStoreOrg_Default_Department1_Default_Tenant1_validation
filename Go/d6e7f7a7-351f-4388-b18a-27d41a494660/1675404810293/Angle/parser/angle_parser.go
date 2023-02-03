// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675404810293/Angle.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Angle

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

type AngleParser struct {
	*antlr.BaseParser
}

var angleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func angleParserInit() {
	staticData := &angleParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ANGLE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "angle",
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

// AngleParserInit initializes any static state used to implement AngleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAngleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AngleParserInit() {
	staticData := &angleParserStaticData
	staticData.once.Do(angleParserInit)
}

// NewAngleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAngleParser(input antlr.TokenStream) *AngleParser {
	AngleParserInit()
	this := new(AngleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &angleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Angle.g4"

	return this
}

// AngleParser tokens.
const (
	AngleParserEOF      = antlr.TokenEOF
	AngleParserANGLE    = 1
	AngleParserOWNKEY   = 2
	AngleParserSPLITKEY = 3
	AngleParserWS       = 4
)

// AngleParser rules.
const (
	AngleParserRULE_expression = 0
	AngleParserRULE_angle      = 1
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
	p.RuleIndex = AngleParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AngleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Angle() IAngleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAngleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAngleContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AngleParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AngleListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AngleListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AngleParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AngleParserRULE_expression)

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
		p.Angle()
	}
	{
		p.SetState(5)
		p.Match(AngleParserEOF)
	}

	return localctx
}

// IAngleContext is an interface to support dynamic dispatch.
type IAngleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAngleContext differentiates from other interfaces.
	IsAngleContext()
}

type AngleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAngleContext() *AngleContext {
	var p = new(AngleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AngleParserRULE_angle
	return p
}

func (*AngleContext) IsAngleContext() {}

func NewAngleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AngleContext {
	var p = new(AngleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AngleParserRULE_angle

	return p
}

func (s *AngleContext) GetParser() antlr.Parser { return s.parser }

func (s *AngleContext) ANGLE() antlr.TerminalNode {
	return s.GetToken(AngleParserANGLE, 0)
}

func (s *AngleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AngleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AngleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AngleListener); ok {
		listenerT.EnterAngle(s)
	}
}

func (s *AngleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AngleListener); ok {
		listenerT.ExitAngle(s)
	}
}

func (p *AngleParser) Angle() (localctx IAngleContext) {
	this := p
	_ = this

	localctx = NewAngleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AngleParserRULE_angle)

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
		p.Match(AngleParserANGLE)
	}

	return localctx
}
