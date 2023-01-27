// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837087230/Project.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Project

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

type ProjectParser struct {
	*antlr.BaseParser
}

var projectParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func projectParserInit() {
	staticData := &projectParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PROJECT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "project",
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

// ProjectParserInit initializes any static state used to implement ProjectParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProjectParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ProjectParserInit() {
	staticData := &projectParserStaticData
	staticData.once.Do(projectParserInit)
}

// NewProjectParser produces a new parser instance for the optional input antlr.TokenStream.
func NewProjectParser(input antlr.TokenStream) *ProjectParser {
	ProjectParserInit()
	this := new(ProjectParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &projectParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Project.g4"

	return this
}

// ProjectParser tokens.
const (
	ProjectParserEOF      = antlr.TokenEOF
	ProjectParserPROJECT  = 1
	ProjectParserOWNKEY   = 2
	ProjectParserSPLITKEY = 3
	ProjectParserWS       = 4
)

// ProjectParser rules.
const (
	ProjectParserRULE_expression = 0
	ProjectParserRULE_project    = 1
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
	p.RuleIndex = ProjectParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Project() IProjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProjectContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ProjectParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ProjectParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ProjectParserRULE_expression)

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
		p.Project()
	}
	{
		p.SetState(5)
		p.Match(ProjectParserEOF)
	}

	return localctx
}

// IProjectContext is an interface to support dynamic dispatch.
type IProjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProjectContext differentiates from other interfaces.
	IsProjectContext()
}

type ProjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProjectContext() *ProjectContext {
	var p = new(ProjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ProjectParserRULE_project
	return p
}

func (*ProjectContext) IsProjectContext() {}

func NewProjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProjectContext {
	var p = new(ProjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ProjectParserRULE_project

	return p
}

func (s *ProjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ProjectContext) PROJECT() antlr.TerminalNode {
	return s.GetToken(ProjectParserPROJECT, 0)
}

func (s *ProjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectListener); ok {
		listenerT.EnterProject(s)
	}
}

func (s *ProjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ProjectListener); ok {
		listenerT.ExitProject(s)
	}
}

func (p *ProjectParser) Project() (localctx IProjectContext) {
	this := p
	_ = this

	localctx = NewProjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ProjectParserRULE_project)

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
		p.Match(ProjectParserPROJECT)
	}

	return localctx
}
