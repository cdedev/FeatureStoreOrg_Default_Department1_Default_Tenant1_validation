// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671863736515/Grade.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Grade

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

type GradeParser struct {
	*antlr.BaseParser
}

var gradeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gradeParserInit() {
	staticData := &gradeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GRADE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "grade",
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

// GradeParserInit initializes any static state used to implement GradeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGradeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GradeParserInit() {
	staticData := &gradeParserStaticData
	staticData.once.Do(gradeParserInit)
}

// NewGradeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGradeParser(input antlr.TokenStream) *GradeParser {
	GradeParserInit()
	this := new(GradeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gradeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Grade.g4"

	return this
}

// GradeParser tokens.
const (
	GradeParserEOF      = antlr.TokenEOF
	GradeParserGRADE    = 1
	GradeParserOWNKEY   = 2
	GradeParserSPLITKEY = 3
	GradeParserWS       = 4
)

// GradeParser rules.
const (
	GradeParserRULE_expression = 0
	GradeParserRULE_grade      = 1
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
	p.RuleIndex = GradeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GradeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Grade() IGradeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGradeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGradeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GradeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GradeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GradeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GradeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GradeParserRULE_expression)

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
		p.Grade()
	}
	{
		p.SetState(5)
		p.Match(GradeParserEOF)
	}

	return localctx
}

// IGradeContext is an interface to support dynamic dispatch.
type IGradeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGradeContext differentiates from other interfaces.
	IsGradeContext()
}

type GradeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGradeContext() *GradeContext {
	var p = new(GradeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GradeParserRULE_grade
	return p
}

func (*GradeContext) IsGradeContext() {}

func NewGradeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GradeContext {
	var p = new(GradeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GradeParserRULE_grade

	return p
}

func (s *GradeContext) GetParser() antlr.Parser { return s.parser }

func (s *GradeContext) GRADE() antlr.TerminalNode {
	return s.GetToken(GradeParserGRADE, 0)
}

func (s *GradeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GradeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GradeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GradeListener); ok {
		listenerT.EnterGrade(s)
	}
}

func (s *GradeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GradeListener); ok {
		listenerT.ExitGrade(s)
	}
}

func (p *GradeParser) Grade() (localctx IGradeContext) {
	this := p
	_ = this

	localctx = NewGradeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GradeParserRULE_grade)

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
		p.Match(GradeParserGRADE)
	}

	return localctx
}
