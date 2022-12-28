// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672203911702/JobLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JobLevel

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

type JobLevelParser struct {
	*antlr.BaseParser
}

var joblevelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func joblevelParserInit() {
	staticData := &joblevelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "JOBLEVEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "joblevel",
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

// JobLevelParserInit initializes any static state used to implement JobLevelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJobLevelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JobLevelParserInit() {
	staticData := &joblevelParserStaticData
	staticData.once.Do(joblevelParserInit)
}

// NewJobLevelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJobLevelParser(input antlr.TokenStream) *JobLevelParser {
	JobLevelParserInit()
	this := new(JobLevelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &joblevelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "JobLevel.g4"

	return this
}

// JobLevelParser tokens.
const (
	JobLevelParserEOF      = antlr.TokenEOF
	JobLevelParserJOBLEVEL = 1
	JobLevelParserOWNKEY   = 2
	JobLevelParserSPLITKEY = 3
	JobLevelParserWS       = 4
)

// JobLevelParser rules.
const (
	JobLevelParserRULE_expression = 0
	JobLevelParserRULE_joblevel   = 1
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
	p.RuleIndex = JobLevelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobLevelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Joblevel() IJoblevelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoblevelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoblevelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(JobLevelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobLevelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobLevelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *JobLevelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JobLevelParserRULE_expression)

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
		p.Joblevel()
	}
	{
		p.SetState(5)
		p.Match(JobLevelParserEOF)
	}

	return localctx
}

// IJoblevelContext is an interface to support dynamic dispatch.
type IJoblevelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoblevelContext differentiates from other interfaces.
	IsJoblevelContext()
}

type JoblevelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoblevelContext() *JoblevelContext {
	var p = new(JoblevelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = JobLevelParserRULE_joblevel
	return p
}

func (*JoblevelContext) IsJoblevelContext() {}

func NewJoblevelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoblevelContext {
	var p = new(JoblevelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = JobLevelParserRULE_joblevel

	return p
}

func (s *JoblevelContext) GetParser() antlr.Parser { return s.parser }

func (s *JoblevelContext) JOBLEVEL() antlr.TerminalNode {
	return s.GetToken(JobLevelParserJOBLEVEL, 0)
}

func (s *JoblevelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoblevelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoblevelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobLevelListener); ok {
		listenerT.EnterJoblevel(s)
	}
}

func (s *JoblevelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JobLevelListener); ok {
		listenerT.ExitJoblevel(s)
	}
}

func (p *JobLevelParser) Joblevel() (localctx IJoblevelContext) {
	this := p
	_ = this

	localctx = NewJoblevelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JobLevelParserRULE_joblevel)

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
		p.Match(JobLevelParserJOBLEVEL)
	}

	return localctx
}
