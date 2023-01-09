// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235164422/Openness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Openness

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

type OpennessParser struct {
	*antlr.BaseParser
}

var opennessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func opennessParserInit() {
	staticData := &opennessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OPENNESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "openness",
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

// OpennessParserInit initializes any static state used to implement OpennessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOpennessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpennessParserInit() {
	staticData := &opennessParserStaticData
	staticData.once.Do(opennessParserInit)
}

// NewOpennessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOpennessParser(input antlr.TokenStream) *OpennessParser {
	OpennessParserInit()
	this := new(OpennessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &opennessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Openness.g4"

	return this
}

// OpennessParser tokens.
const (
	OpennessParserEOF      = antlr.TokenEOF
	OpennessParserOPENNESS = 1
	OpennessParserOWNKEY   = 2
	OpennessParserSPLITKEY = 3
	OpennessParserWS       = 4
)

// OpennessParser rules.
const (
	OpennessParserRULE_expression = 0
	OpennessParserRULE_openness   = 1
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
	p.RuleIndex = OpennessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpennessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Openness() IOpennessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpennessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpennessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OpennessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpennessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpennessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OpennessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OpennessParserRULE_expression)

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
		p.Openness()
	}
	{
		p.SetState(5)
		p.Match(OpennessParserEOF)
	}

	return localctx
}

// IOpennessContext is an interface to support dynamic dispatch.
type IOpennessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpennessContext differentiates from other interfaces.
	IsOpennessContext()
}

type OpennessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpennessContext() *OpennessContext {
	var p = new(OpennessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpennessParserRULE_openness
	return p
}

func (*OpennessContext) IsOpennessContext() {}

func NewOpennessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpennessContext {
	var p = new(OpennessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpennessParserRULE_openness

	return p
}

func (s *OpennessContext) GetParser() antlr.Parser { return s.parser }

func (s *OpennessContext) OPENNESS() antlr.TerminalNode {
	return s.GetToken(OpennessParserOPENNESS, 0)
}

func (s *OpennessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpennessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpennessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpennessListener); ok {
		listenerT.EnterOpenness(s)
	}
}

func (s *OpennessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpennessListener); ok {
		listenerT.ExitOpenness(s)
	}
}

func (p *OpennessParser) Openness() (localctx IOpennessContext) {
	this := p
	_ = this

	localctx = NewOpennessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpennessParserRULE_openness)

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
		p.Match(OpennessParserOPENNESS)
	}

	return localctx
}
