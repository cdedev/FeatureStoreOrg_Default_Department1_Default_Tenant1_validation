// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676462327753/WorkingCapacity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WorkingCapacity

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

type WorkingCapacityParser struct {
	*antlr.BaseParser
}

var workingcapacityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func workingcapacityParserInit() {
	staticData := &workingcapacityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WORKINGCAPACITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "workingcapacity",
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

// WorkingCapacityParserInit initializes any static state used to implement WorkingCapacityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWorkingCapacityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WorkingCapacityParserInit() {
	staticData := &workingcapacityParserStaticData
	staticData.once.Do(workingcapacityParserInit)
}

// NewWorkingCapacityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWorkingCapacityParser(input antlr.TokenStream) *WorkingCapacityParser {
	WorkingCapacityParserInit()
	this := new(WorkingCapacityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &workingcapacityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "WorkingCapacity.g4"

	return this
}

// WorkingCapacityParser tokens.
const (
	WorkingCapacityParserEOF             = antlr.TokenEOF
	WorkingCapacityParserWORKINGCAPACITY = 1
	WorkingCapacityParserOWNKEY          = 2
	WorkingCapacityParserSPLITKEY        = 3
	WorkingCapacityParserWS              = 4
)

// WorkingCapacityParser rules.
const (
	WorkingCapacityParserRULE_expression      = 0
	WorkingCapacityParserRULE_workingcapacity = 1
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
	p.RuleIndex = WorkingCapacityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkingCapacityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Workingcapacity() IWorkingcapacityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWorkingcapacityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWorkingcapacityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WorkingCapacityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkingCapacityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkingCapacityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WorkingCapacityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WorkingCapacityParserRULE_expression)

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
		p.Workingcapacity()
	}
	{
		p.SetState(5)
		p.Match(WorkingCapacityParserEOF)
	}

	return localctx
}

// IWorkingcapacityContext is an interface to support dynamic dispatch.
type IWorkingcapacityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWorkingcapacityContext differentiates from other interfaces.
	IsWorkingcapacityContext()
}

type WorkingcapacityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWorkingcapacityContext() *WorkingcapacityContext {
	var p = new(WorkingcapacityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WorkingCapacityParserRULE_workingcapacity
	return p
}

func (*WorkingcapacityContext) IsWorkingcapacityContext() {}

func NewWorkingcapacityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WorkingcapacityContext {
	var p = new(WorkingcapacityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WorkingCapacityParserRULE_workingcapacity

	return p
}

func (s *WorkingcapacityContext) GetParser() antlr.Parser { return s.parser }

func (s *WorkingcapacityContext) WORKINGCAPACITY() antlr.TerminalNode {
	return s.GetToken(WorkingCapacityParserWORKINGCAPACITY, 0)
}

func (s *WorkingcapacityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WorkingcapacityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WorkingcapacityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkingCapacityListener); ok {
		listenerT.EnterWorkingcapacity(s)
	}
}

func (s *WorkingcapacityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WorkingCapacityListener); ok {
		listenerT.ExitWorkingcapacity(s)
	}
}

func (p *WorkingCapacityParser) Workingcapacity() (localctx IWorkingcapacityContext) {
	this := p
	_ = this

	localctx = NewWorkingcapacityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WorkingCapacityParserRULE_workingcapacity)

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
		p.Match(WorkingCapacityParserWORKINGCAPACITY)
	}

	return localctx
}
