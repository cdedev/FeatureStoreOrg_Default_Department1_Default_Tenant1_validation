// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671597104664/Experience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Experience

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

type ExperienceParser struct {
	*antlr.BaseParser
}

var experienceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func experienceParserInit() {
	staticData := &experienceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EXPERIENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "experience",
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

// ExperienceParserInit initializes any static state used to implement ExperienceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExperienceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExperienceParserInit() {
	staticData := &experienceParserStaticData
	staticData.once.Do(experienceParserInit)
}

// NewExperienceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExperienceParser(input antlr.TokenStream) *ExperienceParser {
	ExperienceParserInit()
	this := new(ExperienceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &experienceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Experience.g4"

	return this
}

// ExperienceParser tokens.
const (
	ExperienceParserEOF        = antlr.TokenEOF
	ExperienceParserEXPERIENCE = 1
	ExperienceParserOWNKEY     = 2
	ExperienceParserSPLITKEY   = 3
	ExperienceParserWS         = 4
)

// ExperienceParser rules.
const (
	ExperienceParserRULE_expression = 0
	ExperienceParserRULE_experience = 1
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
	p.RuleIndex = ExperienceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExperienceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Experience() IExperienceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExperienceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExperienceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExperienceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExperienceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExperienceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExperienceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExperienceParserRULE_expression)

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
		p.Experience()
	}
	{
		p.SetState(5)
		p.Match(ExperienceParserEOF)
	}

	return localctx
}

// IExperienceContext is an interface to support dynamic dispatch.
type IExperienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExperienceContext differentiates from other interfaces.
	IsExperienceContext()
}

type ExperienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExperienceContext() *ExperienceContext {
	var p = new(ExperienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExperienceParserRULE_experience
	return p
}

func (*ExperienceContext) IsExperienceContext() {}

func NewExperienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExperienceContext {
	var p = new(ExperienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExperienceParserRULE_experience

	return p
}

func (s *ExperienceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExperienceContext) EXPERIENCE() antlr.TerminalNode {
	return s.GetToken(ExperienceParserEXPERIENCE, 0)
}

func (s *ExperienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExperienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExperienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExperienceListener); ok {
		listenerT.EnterExperience(s)
	}
}

func (s *ExperienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExperienceListener); ok {
		listenerT.ExitExperience(s)
	}
}

func (p *ExperienceParser) Experience() (localctx IExperienceContext) {
	this := p
	_ = this

	localctx = NewExperienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExperienceParserRULE_experience)

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
		p.Match(ExperienceParserEXPERIENCE)
	}

	return localctx
}
