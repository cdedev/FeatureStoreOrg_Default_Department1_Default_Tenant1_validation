// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069318693/Course.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Course

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

type CourseParser struct {
	*antlr.BaseParser
}

var courseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func courseParserInit() {
	staticData := &courseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COURSE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "course",
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

// CourseParserInit initializes any static state used to implement CourseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCourseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CourseParserInit() {
	staticData := &courseParserStaticData
	staticData.once.Do(courseParserInit)
}

// NewCourseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCourseParser(input antlr.TokenStream) *CourseParser {
	CourseParserInit()
	this := new(CourseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &courseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Course.g4"

	return this
}

// CourseParser tokens.
const (
	CourseParserEOF      = antlr.TokenEOF
	CourseParserCOURSE   = 1
	CourseParserOWNKEY   = 2
	CourseParserSPLITKEY = 3
	CourseParserWS       = 4
)

// CourseParser rules.
const (
	CourseParserRULE_expression = 0
	CourseParserRULE_course     = 1
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
	p.RuleIndex = CourseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CourseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Course() ICourseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICourseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICourseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CourseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CourseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CourseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CourseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CourseParserRULE_expression)

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
		p.Course()
	}
	{
		p.SetState(5)
		p.Match(CourseParserEOF)
	}

	return localctx
}

// ICourseContext is an interface to support dynamic dispatch.
type ICourseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCourseContext differentiates from other interfaces.
	IsCourseContext()
}

type CourseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCourseContext() *CourseContext {
	var p = new(CourseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CourseParserRULE_course
	return p
}

func (*CourseContext) IsCourseContext() {}

func NewCourseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CourseContext {
	var p = new(CourseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CourseParserRULE_course

	return p
}

func (s *CourseContext) GetParser() antlr.Parser { return s.parser }

func (s *CourseContext) COURSE() antlr.TerminalNode {
	return s.GetToken(CourseParserCOURSE, 0)
}

func (s *CourseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CourseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CourseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CourseListener); ok {
		listenerT.EnterCourse(s)
	}
}

func (s *CourseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CourseListener); ok {
		listenerT.ExitCourse(s)
	}
}

func (p *CourseParser) Course() (localctx ICourseContext) {
	this := p
	_ = this

	localctx = NewCourseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CourseParserRULE_course)

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
		p.Match(CourseParserCOURSE)
	}

	return localctx
}
