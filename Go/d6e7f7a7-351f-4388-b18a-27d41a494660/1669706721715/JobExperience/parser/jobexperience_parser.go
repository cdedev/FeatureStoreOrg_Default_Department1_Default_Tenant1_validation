// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669706721715/JobExperience.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JobExperience

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

type JobExperienceParser struct {
	*antlr.BaseParser
}

var jobexperienceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jobexperienceParserInit() {
	staticData := &jobexperienceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JOBEXPERIENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "jobexperience",
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

// JobExperienceParserInit initializes any static state used to implement JobExperienceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJobExperienceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JobExperienceParserInit() {
	staticData := &jobexperienceParserStaticData
	staticData.once.Do(jobexperienceParserInit)
}

// NewJobExperienceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJobExperienceParser(input antlr.TokenStream) *JobExperienceParser {
	JobExperienceParserInit()
	this := new(JobExperienceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jobexperienceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JobExperience.g4"

	return this
}

// JobExperienceParser tokens.
const (
	JobExperienceParserEOF           = antlr.TokenEOF
	JobExperienceParserJOBEXPERIENCE = 1
	JobExperienceParserOWNKEY        = 2
	JobExperienceParserSPLITKEY      = 3
	JobExperienceParserWS            = 4
)

// JobExperienceParser rules.
const (
	JobExperienceParserRULE_expression    = 0
	JobExperienceParserRULE_jobexperience = 1
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
	p.RuleIndex = JobExperienceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobExperienceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Jobexperience() IJobexperienceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJobexperienceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJobexperienceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JobExperienceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobExperienceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobExperienceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JobExperienceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JobExperienceParserRULE_expression)

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
		p.Jobexperience()
	}
	{
		p.SetState(5)
		p.Match(JobExperienceParserEOF)
	}

	return localctx
}

// IJobexperienceContext is an interface to support dynamic dispatch.
type IJobexperienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJobexperienceContext differentiates from other interfaces.
	IsJobexperienceContext()
}

type JobexperienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJobexperienceContext() *JobexperienceContext {
	var p = new(JobexperienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JobExperienceParserRULE_jobexperience
	return p
}

func (*JobexperienceContext) IsJobexperienceContext() {}

func NewJobexperienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JobexperienceContext {
	var p = new(JobexperienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobExperienceParserRULE_jobexperience

	return p
}

func (s *JobexperienceContext) GetParser() antlr.Parser { return s.parser }

func (s *JobexperienceContext) JOBEXPERIENCE() antlr.TerminalNode {
	return s.GetToken(JobExperienceParserJOBEXPERIENCE, 0)
}

func (s *JobexperienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JobexperienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JobexperienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobExperienceListener); ok {
		listenerT.EnterJobexperience(s)
	}
}

func (s *JobexperienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobExperienceListener); ok {
		listenerT.ExitJobexperience(s)
	}
}

func (p *JobExperienceParser) Jobexperience() (localctx IJobexperienceContext) {
	this := p
	_ = this

	localctx = NewJobexperienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JobExperienceParserRULE_jobexperience)

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
		p.Match(JobExperienceParserJOBEXPERIENCE)
	}

	return localctx
}
