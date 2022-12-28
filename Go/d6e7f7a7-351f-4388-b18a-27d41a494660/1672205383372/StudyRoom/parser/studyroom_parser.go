// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205383372/StudyRoom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // StudyRoom

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

type StudyRoomParser struct {
	*antlr.BaseParser
}

var studyroomParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func studyroomParserInit() {
	staticData := &studyroomParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "STUDYROOM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "studyroom",
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

// StudyRoomParserInit initializes any static state used to implement StudyRoomParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewStudyRoomParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func StudyRoomParserInit() {
	staticData := &studyroomParserStaticData
	staticData.once.Do(studyroomParserInit)
}

// NewStudyRoomParser produces a new parser instance for the optional input antlr.TokenStream.
func NewStudyRoomParser(input antlr.TokenStream) *StudyRoomParser {
	StudyRoomParserInit()
	this := new(StudyRoomParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &studyroomParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "StudyRoom.g4"

	return this
}

// StudyRoomParser tokens.
const (
	StudyRoomParserEOF       = antlr.TokenEOF
	StudyRoomParserSTUDYROOM = 1
	StudyRoomParserOWNKEY    = 2
	StudyRoomParserSPLITKEY  = 3
	StudyRoomParserWS        = 4
)

// StudyRoomParser rules.
const (
	StudyRoomParserRULE_expression = 0
	StudyRoomParserRULE_studyroom  = 1
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
	p.RuleIndex = StudyRoomParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StudyRoomParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Studyroom() IStudyroomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStudyroomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStudyroomContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(StudyRoomParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StudyRoomListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StudyRoomListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *StudyRoomParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, StudyRoomParserRULE_expression)

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
		p.Studyroom()
	}
	{
		p.SetState(5)
		p.Match(StudyRoomParserEOF)
	}

	return localctx
}

// IStudyroomContext is an interface to support dynamic dispatch.
type IStudyroomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStudyroomContext differentiates from other interfaces.
	IsStudyroomContext()
}

type StudyroomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStudyroomContext() *StudyroomContext {
	var p = new(StudyroomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = StudyRoomParserRULE_studyroom
	return p
}

func (*StudyroomContext) IsStudyroomContext() {}

func NewStudyroomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StudyroomContext {
	var p = new(StudyroomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = StudyRoomParserRULE_studyroom

	return p
}

func (s *StudyroomContext) GetParser() antlr.Parser { return s.parser }

func (s *StudyroomContext) STUDYROOM() antlr.TerminalNode {
	return s.GetToken(StudyRoomParserSTUDYROOM, 0)
}

func (s *StudyroomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StudyroomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StudyroomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StudyRoomListener); ok {
		listenerT.EnterStudyroom(s)
	}
}

func (s *StudyroomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StudyRoomListener); ok {
		listenerT.ExitStudyroom(s)
	}
}

func (p *StudyRoomParser) Studyroom() (localctx IStudyroomContext) {
	this := p
	_ = this

	localctx = NewStudyroomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, StudyRoomParserRULE_studyroom)

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
		p.Match(StudyRoomParserSTUDYROOM)
	}

	return localctx
}
