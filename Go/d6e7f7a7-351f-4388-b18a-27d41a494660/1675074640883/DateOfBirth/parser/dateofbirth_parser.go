// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675074640883/DateOfBirth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DateOfBirth

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

type DateOfBirthParser struct {
	*antlr.BaseParser
}

var dateofbirthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dateofbirthParserInit() {
	staticData := &dateofbirthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DATEOFBIRTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dateofbirth",
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

// DateOfBirthParserInit initializes any static state used to implement DateOfBirthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDateOfBirthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DateOfBirthParserInit() {
	staticData := &dateofbirthParserStaticData
	staticData.once.Do(dateofbirthParserInit)
}

// NewDateOfBirthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDateOfBirthParser(input antlr.TokenStream) *DateOfBirthParser {
	DateOfBirthParserInit()
	this := new(DateOfBirthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dateofbirthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DateOfBirth.g4"

	return this
}

// DateOfBirthParser tokens.
const (
	DateOfBirthParserEOF         = antlr.TokenEOF
	DateOfBirthParserDATEOFBIRTH = 1
	DateOfBirthParserOWNKEY      = 2
	DateOfBirthParserSPLITKEY    = 3
	DateOfBirthParserWS          = 4
)

// DateOfBirthParser rules.
const (
	DateOfBirthParserRULE_expression  = 0
	DateOfBirthParserRULE_dateofbirth = 1
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
	p.RuleIndex = DateOfBirthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DateOfBirthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dateofbirth() IDateofbirthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateofbirthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateofbirthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DateOfBirthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfBirthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfBirthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DateOfBirthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DateOfBirthParserRULE_expression)

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
		p.Dateofbirth()
	}
	{
		p.SetState(5)
		p.Match(DateOfBirthParserEOF)
	}

	return localctx
}

// IDateofbirthContext is an interface to support dynamic dispatch.
type IDateofbirthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateofbirthContext differentiates from other interfaces.
	IsDateofbirthContext()
}

type DateofbirthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateofbirthContext() *DateofbirthContext {
	var p = new(DateofbirthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DateOfBirthParserRULE_dateofbirth
	return p
}

func (*DateofbirthContext) IsDateofbirthContext() {}

func NewDateofbirthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateofbirthContext {
	var p = new(DateofbirthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DateOfBirthParserRULE_dateofbirth

	return p
}

func (s *DateofbirthContext) GetParser() antlr.Parser { return s.parser }

func (s *DateofbirthContext) DATEOFBIRTH() antlr.TerminalNode {
	return s.GetToken(DateOfBirthParserDATEOFBIRTH, 0)
}

func (s *DateofbirthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateofbirthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateofbirthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfBirthListener); ok {
		listenerT.EnterDateofbirth(s)
	}
}

func (s *DateofbirthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DateOfBirthListener); ok {
		listenerT.ExitDateofbirth(s)
	}
}

func (p *DateOfBirthParser) Dateofbirth() (localctx IDateofbirthContext) {
	this := p
	_ = this

	localctx = NewDateofbirthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DateOfBirthParserRULE_dateofbirth)

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
		p.Match(DateOfBirthParserDATEOFBIRTH)
	}

	return localctx
}
