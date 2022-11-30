// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789660215/Velocity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Velocity

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

type VelocityParser struct {
	*antlr.BaseParser
}

var velocityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func velocityParserInit() {
	staticData := &velocityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VELOCITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "velocity",
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

// VelocityParserInit initializes any static state used to implement VelocityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVelocityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VelocityParserInit() {
	staticData := &velocityParserStaticData
	staticData.once.Do(velocityParserInit)
}

// NewVelocityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVelocityParser(input antlr.TokenStream) *VelocityParser {
	VelocityParserInit()
	this := new(VelocityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &velocityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Velocity.g4"

	return this
}

// VelocityParser tokens.
const (
	VelocityParserEOF      = antlr.TokenEOF
	VelocityParserVELOCITY = 1
	VelocityParserOWNKEY   = 2
	VelocityParserSPLITKEY = 3
	VelocityParserWS       = 4
)

// VelocityParser rules.
const (
	VelocityParserRULE_expression = 0
	VelocityParserRULE_velocity   = 1
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
	p.RuleIndex = VelocityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VelocityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Velocity() IVelocityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVelocityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVelocityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VelocityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VelocityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VelocityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VelocityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VelocityParserRULE_expression)

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
		p.Velocity()
	}
	{
		p.SetState(5)
		p.Match(VelocityParserEOF)
	}

	return localctx
}

// IVelocityContext is an interface to support dynamic dispatch.
type IVelocityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVelocityContext differentiates from other interfaces.
	IsVelocityContext()
}

type VelocityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVelocityContext() *VelocityContext {
	var p = new(VelocityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VelocityParserRULE_velocity
	return p
}

func (*VelocityContext) IsVelocityContext() {}

func NewVelocityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VelocityContext {
	var p = new(VelocityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VelocityParserRULE_velocity

	return p
}

func (s *VelocityContext) GetParser() antlr.Parser { return s.parser }

func (s *VelocityContext) VELOCITY() antlr.TerminalNode {
	return s.GetToken(VelocityParserVELOCITY, 0)
}

func (s *VelocityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VelocityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VelocityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VelocityListener); ok {
		listenerT.EnterVelocity(s)
	}
}

func (s *VelocityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VelocityListener); ok {
		listenerT.ExitVelocity(s)
	}
}

func (p *VelocityParser) Velocity() (localctx IVelocityContext) {
	this := p
	_ = this

	localctx = NewVelocityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VelocityParserRULE_velocity)

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
		p.Match(VelocityParserVELOCITY)
	}

	return localctx
}
