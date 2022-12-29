// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672286988541/Job.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Job

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

type JobParser struct {
	*antlr.BaseParser
}

var jobParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jobParserInit() {
	staticData := &jobParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JOB", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "job",
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

// JobParserInit initializes any static state used to implement JobParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJobParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JobParserInit() {
	staticData := &jobParserStaticData
	staticData.once.Do(jobParserInit)
}

// NewJobParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJobParser(input antlr.TokenStream) *JobParser {
	JobParserInit()
	this := new(JobParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &jobParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Job.g4"

	return this
}

// JobParser tokens.
const (
	JobParserEOF      = antlr.TokenEOF
	JobParserJOB      = 1
	JobParserOWNKEY   = 2
	JobParserSPLITKEY = 3
	JobParserWS       = 4
)

// JobParser rules.
const (
	JobParserRULE_expression = 0
	JobParserRULE_job        = 1
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
	p.RuleIndex = JobParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Job() IJobContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJobContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJobContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JobParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JobParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JobParserRULE_expression)

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
		p.Job()
	}
	{
		p.SetState(5)
		p.Match(JobParserEOF)
	}

	return localctx
}

// IJobContext is an interface to support dynamic dispatch.
type IJobContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJobContext differentiates from other interfaces.
	IsJobContext()
}

type JobContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJobContext() *JobContext {
	var p = new(JobContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JobParserRULE_job
	return p
}

func (*JobContext) IsJobContext() {}

func NewJobContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JobContext {
	var p = new(JobContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobParserRULE_job

	return p
}

func (s *JobContext) GetParser() antlr.Parser { return s.parser }

func (s *JobContext) JOB() antlr.TerminalNode {
	return s.GetToken(JobParserJOB, 0)
}

func (s *JobContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JobContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JobContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobListener); ok {
		listenerT.EnterJob(s)
	}
}

func (s *JobContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobListener); ok {
		listenerT.ExitJob(s)
	}
}

func (p *JobParser) Job() (localctx IJobContext) {
	this := p
	_ = this

	localctx = NewJobContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JobParserRULE_job)

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
		p.Match(JobParserJOB)
	}

	return localctx
}
