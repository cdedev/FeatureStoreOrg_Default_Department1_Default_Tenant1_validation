// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671769232472/MaleLifeExpectancy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaleLifeExpectancy

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

type MaleLifeExpectancyParser struct {
	*antlr.BaseParser
}

var malelifeexpectancyParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func malelifeexpectancyParserInit() {
	staticData := &malelifeexpectancyParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MALELIFEEXPECTANCY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "malelifeexpectancy",
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

// MaleLifeExpectancyParserInit initializes any static state used to implement MaleLifeExpectancyParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMaleLifeExpectancyParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MaleLifeExpectancyParserInit() {
	staticData := &malelifeexpectancyParserStaticData
	staticData.once.Do(malelifeexpectancyParserInit)
}

// NewMaleLifeExpectancyParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMaleLifeExpectancyParser(input antlr.TokenStream) *MaleLifeExpectancyParser {
	MaleLifeExpectancyParserInit()
	this := new(MaleLifeExpectancyParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &malelifeexpectancyParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MaleLifeExpectancy.g4"

	return this
}

// MaleLifeExpectancyParser tokens.
const (
	MaleLifeExpectancyParserEOF                = antlr.TokenEOF
	MaleLifeExpectancyParserMALELIFEEXPECTANCY = 1
	MaleLifeExpectancyParserOWNKEY             = 2
	MaleLifeExpectancyParserSPLITKEY           = 3
	MaleLifeExpectancyParserWS                 = 4
)

// MaleLifeExpectancyParser rules.
const (
	MaleLifeExpectancyParserRULE_expression         = 0
	MaleLifeExpectancyParserRULE_malelifeexpectancy = 1
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
	p.RuleIndex = MaleLifeExpectancyParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaleLifeExpectancyParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Malelifeexpectancy() IMalelifeexpectancyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMalelifeexpectancyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMalelifeexpectancyContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MaleLifeExpectancyParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaleLifeExpectancyListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaleLifeExpectancyListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MaleLifeExpectancyParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MaleLifeExpectancyParserRULE_expression)

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
		p.Malelifeexpectancy()
	}
	{
		p.SetState(5)
		p.Match(MaleLifeExpectancyParserEOF)
	}

	return localctx
}

// IMalelifeexpectancyContext is an interface to support dynamic dispatch.
type IMalelifeexpectancyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMalelifeexpectancyContext differentiates from other interfaces.
	IsMalelifeexpectancyContext()
}

type MalelifeexpectancyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMalelifeexpectancyContext() *MalelifeexpectancyContext {
	var p = new(MalelifeexpectancyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MaleLifeExpectancyParserRULE_malelifeexpectancy
	return p
}

func (*MalelifeexpectancyContext) IsMalelifeexpectancyContext() {}

func NewMalelifeexpectancyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MalelifeexpectancyContext {
	var p = new(MalelifeexpectancyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MaleLifeExpectancyParserRULE_malelifeexpectancy

	return p
}

func (s *MalelifeexpectancyContext) GetParser() antlr.Parser { return s.parser }

func (s *MalelifeexpectancyContext) MALELIFEEXPECTANCY() antlr.TerminalNode {
	return s.GetToken(MaleLifeExpectancyParserMALELIFEEXPECTANCY, 0)
}

func (s *MalelifeexpectancyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MalelifeexpectancyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MalelifeexpectancyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaleLifeExpectancyListener); ok {
		listenerT.EnterMalelifeexpectancy(s)
	}
}

func (s *MalelifeexpectancyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MaleLifeExpectancyListener); ok {
		listenerT.ExitMalelifeexpectancy(s)
	}
}

func (p *MaleLifeExpectancyParser) Malelifeexpectancy() (localctx IMalelifeexpectancyContext) {
	this := p
	_ = this

	localctx = NewMalelifeexpectancyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MaleLifeExpectancyParserRULE_malelifeexpectancy)

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
		p.Match(MaleLifeExpectancyParserMALELIFEEXPECTANCY)
	}

	return localctx
}
