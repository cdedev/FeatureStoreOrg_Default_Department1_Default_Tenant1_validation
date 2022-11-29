// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700897117/BMI.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BMI

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

type BMIParser struct {
	*antlr.BaseParser
}

var bmiParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bmiParserInit() {
	staticData := &bmiParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BMI", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bmi",
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

// BMIParserInit initializes any static state used to implement BMIParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBMIParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BMIParserInit() {
	staticData := &bmiParserStaticData
	staticData.once.Do(bmiParserInit)
}

// NewBMIParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBMIParser(input antlr.TokenStream) *BMIParser {
	BMIParserInit()
	this := new(BMIParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bmiParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BMI.g4"

	return this
}

// BMIParser tokens.
const (
	BMIParserEOF      = antlr.TokenEOF
	BMIParserBMI      = 1
	BMIParserOWNKEY   = 2
	BMIParserSPLITKEY = 3
	BMIParserWS       = 4
)

// BMIParser rules.
const (
	BMIParserRULE_expression = 0
	BMIParserRULE_bmi        = 1
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
	p.RuleIndex = BMIParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BMIParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bmi() IBmiContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBmiContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBmiContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BMIParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BMIListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BMIListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BMIParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BMIParserRULE_expression)

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
		p.Bmi()
	}
	{
		p.SetState(5)
		p.Match(BMIParserEOF)
	}

	return localctx
}

// IBmiContext is an interface to support dynamic dispatch.
type IBmiContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBmiContext differentiates from other interfaces.
	IsBmiContext()
}

type BmiContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBmiContext() *BmiContext {
	var p = new(BmiContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BMIParserRULE_bmi
	return p
}

func (*BmiContext) IsBmiContext() {}

func NewBmiContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BmiContext {
	var p = new(BmiContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BMIParserRULE_bmi

	return p
}

func (s *BmiContext) GetParser() antlr.Parser { return s.parser }

func (s *BmiContext) BMI() antlr.TerminalNode {
	return s.GetToken(BMIParserBMI, 0)
}

func (s *BmiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BmiContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BmiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BMIListener); ok {
		listenerT.EnterBmi(s)
	}
}

func (s *BmiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BMIListener); ok {
		listenerT.ExitBmi(s)
	}
}

func (p *BMIParser) Bmi() (localctx IBmiContext) {
	this := p
	_ = this

	localctx = NewBmiContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BMIParserRULE_bmi)

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
		p.Match(BMIParserBMI)
	}

	return localctx
}
