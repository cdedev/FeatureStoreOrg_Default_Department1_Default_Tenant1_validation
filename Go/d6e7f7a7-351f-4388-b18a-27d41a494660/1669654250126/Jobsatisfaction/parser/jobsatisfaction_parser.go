// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669654250126/Jobsatisfaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jobsatisfaction

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

type JobsatisfactionParser struct {
	*antlr.BaseParser
}

var jobsatisfactionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jobsatisfactionParserInit() {
	staticData := &jobsatisfactionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JOBSATISFACTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "jobsatisfaction",
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

// JobsatisfactionParserInit initializes any static state used to implement JobsatisfactionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJobsatisfactionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JobsatisfactionParserInit() {
	staticData := &jobsatisfactionParserStaticData
	staticData.once.Do(jobsatisfactionParserInit)
}

// NewJobsatisfactionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJobsatisfactionParser(input antlr.TokenStream) *JobsatisfactionParser {
	JobsatisfactionParserInit()
	this := new(JobsatisfactionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jobsatisfactionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Jobsatisfaction.g4"

	return this
}

// JobsatisfactionParser tokens.
const (
	JobsatisfactionParserEOF             = antlr.TokenEOF
	JobsatisfactionParserJOBSATISFACTION = 1
	JobsatisfactionParserOWNKEY          = 2
	JobsatisfactionParserSPLITKEY        = 3
	JobsatisfactionParserWS              = 4
)

// JobsatisfactionParser rules.
const (
	JobsatisfactionParserRULE_expression      = 0
	JobsatisfactionParserRULE_jobsatisfaction = 1
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
	p.RuleIndex = JobsatisfactionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobsatisfactionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Jobsatisfaction() IJobsatisfactionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJobsatisfactionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJobsatisfactionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JobsatisfactionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobsatisfactionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobsatisfactionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JobsatisfactionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JobsatisfactionParserRULE_expression)

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
		p.Jobsatisfaction()
	}
	{
		p.SetState(5)
		p.Match(JobsatisfactionParserEOF)
	}

	return localctx
}

// IJobsatisfactionContext is an interface to support dynamic dispatch.
type IJobsatisfactionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJobsatisfactionContext differentiates from other interfaces.
	IsJobsatisfactionContext()
}

type JobsatisfactionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJobsatisfactionContext() *JobsatisfactionContext {
	var p = new(JobsatisfactionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JobsatisfactionParserRULE_jobsatisfaction
	return p
}

func (*JobsatisfactionContext) IsJobsatisfactionContext() {}

func NewJobsatisfactionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JobsatisfactionContext {
	var p = new(JobsatisfactionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobsatisfactionParserRULE_jobsatisfaction

	return p
}

func (s *JobsatisfactionContext) GetParser() antlr.Parser { return s.parser }

func (s *JobsatisfactionContext) JOBSATISFACTION() antlr.TerminalNode {
	return s.GetToken(JobsatisfactionParserJOBSATISFACTION, 0)
}

func (s *JobsatisfactionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JobsatisfactionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JobsatisfactionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobsatisfactionListener); ok {
		listenerT.EnterJobsatisfaction(s)
	}
}

func (s *JobsatisfactionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobsatisfactionListener); ok {
		listenerT.ExitJobsatisfaction(s)
	}
}

func (p *JobsatisfactionParser) Jobsatisfaction() (localctx IJobsatisfactionContext) {
	this := p
	_ = this

	localctx = NewJobsatisfactionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JobsatisfactionParserRULE_jobsatisfaction)

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
		p.Match(JobsatisfactionParserJOBSATISFACTION)
	}

	return localctx
}
