// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669621912364/Education.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Education

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

type EducationParser struct {
	*antlr.BaseParser
}

var educationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func educationParserInit() {
	staticData := &educationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EDUCATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "education",
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

// EducationParserInit initializes any static state used to implement EducationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEducationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EducationParserInit() {
	staticData := &educationParserStaticData
	staticData.once.Do(educationParserInit)
}

// NewEducationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEducationParser(input antlr.TokenStream) *EducationParser {
	EducationParserInit()
	this := new(EducationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &educationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Education.g4"

	return this
}

// EducationParser tokens.
const (
	EducationParserEOF       = antlr.TokenEOF
	EducationParserEDUCATION = 1
	EducationParserOWNKEY    = 2
	EducationParserSPLITKEY  = 3
	EducationParserWS        = 4
)

// EducationParser rules.
const (
	EducationParserRULE_expression = 0
	EducationParserRULE_education  = 1
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
	p.RuleIndex = EducationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EducationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Education() IEducationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEducationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEducationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EducationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EducationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EducationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EducationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EducationParserRULE_expression)

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
		p.Education()
	}
	{
		p.SetState(5)
		p.Match(EducationParserEOF)
	}

	return localctx
}

// IEducationContext is an interface to support dynamic dispatch.
type IEducationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEducationContext differentiates from other interfaces.
	IsEducationContext()
}

type EducationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEducationContext() *EducationContext {
	var p = new(EducationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EducationParserRULE_education
	return p
}

func (*EducationContext) IsEducationContext() {}

func NewEducationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EducationContext {
	var p = new(EducationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EducationParserRULE_education

	return p
}

func (s *EducationContext) GetParser() antlr.Parser { return s.parser }

func (s *EducationContext) EDUCATION() antlr.TerminalNode {
	return s.GetToken(EducationParserEDUCATION, 0)
}

func (s *EducationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EducationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EducationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EducationListener); ok {
		listenerT.EnterEducation(s)
	}
}

func (s *EducationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EducationListener); ok {
		listenerT.ExitEducation(s)
	}
}

func (p *EducationParser) Education() (localctx IEducationContext) {
	this := p
	_ = this

	localctx = NewEducationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EducationParserRULE_education)

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
		p.Match(EducationParserEDUCATION)
	}

	return localctx
}
