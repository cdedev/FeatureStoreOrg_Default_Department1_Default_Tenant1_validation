// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626974586/Gender.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gender

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

type GenderParser struct {
	*antlr.BaseParser
}

var genderParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func genderParserInit() {
	staticData := &genderParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GENDER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gender",
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

// GenderParserInit initializes any static state used to implement GenderParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGenderParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GenderParserInit() {
	staticData := &genderParserStaticData
	staticData.once.Do(genderParserInit)
}

// NewGenderParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGenderParser(input antlr.TokenStream) *GenderParser {
	GenderParserInit()
	this := new(GenderParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &genderParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Gender.g4"

	return this
}

// GenderParser tokens.
const (
	GenderParserEOF      = antlr.TokenEOF
	GenderParserGENDER   = 1
	GenderParserOWNKEY   = 2
	GenderParserSPLITKEY = 3
	GenderParserWS       = 4
)

// GenderParser rules.
const (
	GenderParserRULE_expression = 0
	GenderParserRULE_gender     = 1
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
	p.RuleIndex = GenderParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GenderParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gender() IGenderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenderContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GenderParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenderListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenderListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GenderParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GenderParserRULE_expression)

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
		p.Gender()
	}
	{
		p.SetState(5)
		p.Match(GenderParserEOF)
	}

	return localctx
}

// IGenderContext is an interface to support dynamic dispatch.
type IGenderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenderContext differentiates from other interfaces.
	IsGenderContext()
}

type GenderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenderContext() *GenderContext {
	var p = new(GenderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GenderParserRULE_gender
	return p
}

func (*GenderContext) IsGenderContext() {}

func NewGenderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenderContext {
	var p = new(GenderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GenderParserRULE_gender

	return p
}

func (s *GenderContext) GetParser() antlr.Parser { return s.parser }

func (s *GenderContext) GENDER() antlr.TerminalNode {
	return s.GetToken(GenderParserGENDER, 0)
}

func (s *GenderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenderListener); ok {
		listenerT.EnterGender(s)
	}
}

func (s *GenderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenderListener); ok {
		listenerT.ExitGender(s)
	}
}

func (p *GenderParser) Gender() (localctx IGenderContext) {
	this := p
	_ = this

	localctx = NewGenderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GenderParserRULE_gender)

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
		p.Match(GenderParserGENDER)
	}

	return localctx
}
