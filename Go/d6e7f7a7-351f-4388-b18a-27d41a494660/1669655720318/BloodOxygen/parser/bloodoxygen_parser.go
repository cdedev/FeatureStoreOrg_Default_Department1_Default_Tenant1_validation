// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669655720318/BloodOxygen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodOxygen

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

type BloodOxygenParser struct {
	*antlr.BaseParser
}

var bloodoxygenParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bloodoxygenParserInit() {
	staticData := &bloodoxygenParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BLOODOXYGEN", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bloodoxygen",
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

// BloodOxygenParserInit initializes any static state used to implement BloodOxygenParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBloodOxygenParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BloodOxygenParserInit() {
	staticData := &bloodoxygenParserStaticData
	staticData.once.Do(bloodoxygenParserInit)
}

// NewBloodOxygenParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBloodOxygenParser(input antlr.TokenStream) *BloodOxygenParser {
	BloodOxygenParserInit()
	this := new(BloodOxygenParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bloodoxygenParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BloodOxygen.g4"

	return this
}

// BloodOxygenParser tokens.
const (
	BloodOxygenParserEOF         = antlr.TokenEOF
	BloodOxygenParserBLOODOXYGEN = 1
	BloodOxygenParserOWNKEY      = 2
	BloodOxygenParserSPLITKEY    = 3
	BloodOxygenParserWS          = 4
)

// BloodOxygenParser rules.
const (
	BloodOxygenParserRULE_expression  = 0
	BloodOxygenParserRULE_bloodoxygen = 1
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
	p.RuleIndex = BloodOxygenParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BloodOxygenParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bloodoxygen() IBloodoxygenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloodoxygenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloodoxygenContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BloodOxygenParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodOxygenListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodOxygenListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BloodOxygenParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BloodOxygenParserRULE_expression)

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
		p.Bloodoxygen()
	}
	{
		p.SetState(5)
		p.Match(BloodOxygenParserEOF)
	}

	return localctx
}

// IBloodoxygenContext is an interface to support dynamic dispatch.
type IBloodoxygenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloodoxygenContext differentiates from other interfaces.
	IsBloodoxygenContext()
}

type BloodoxygenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloodoxygenContext() *BloodoxygenContext {
	var p = new(BloodoxygenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BloodOxygenParserRULE_bloodoxygen
	return p
}

func (*BloodoxygenContext) IsBloodoxygenContext() {}

func NewBloodoxygenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloodoxygenContext {
	var p = new(BloodoxygenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BloodOxygenParserRULE_bloodoxygen

	return p
}

func (s *BloodoxygenContext) GetParser() antlr.Parser { return s.parser }

func (s *BloodoxygenContext) BLOODOXYGEN() antlr.TerminalNode {
	return s.GetToken(BloodOxygenParserBLOODOXYGEN, 0)
}

func (s *BloodoxygenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloodoxygenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloodoxygenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodOxygenListener); ok {
		listenerT.EnterBloodoxygen(s)
	}
}

func (s *BloodoxygenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodOxygenListener); ok {
		listenerT.ExitBloodoxygen(s)
	}
}

func (p *BloodOxygenParser) Bloodoxygen() (localctx IBloodoxygenContext) {
	this := p
	_ = this

	localctx = NewBloodoxygenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BloodOxygenParserRULE_bloodoxygen)

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
		p.Match(BloodOxygenParserBLOODOXYGEN)
	}

	return localctx
}
