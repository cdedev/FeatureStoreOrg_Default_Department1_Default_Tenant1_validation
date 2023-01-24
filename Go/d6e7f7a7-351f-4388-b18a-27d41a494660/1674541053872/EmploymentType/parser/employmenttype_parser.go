// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674541053872/EmploymentType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EmploymentType

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

type EmploymentTypeParser struct {
	*antlr.BaseParser
}

var employmenttypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func employmenttypeParserInit() {
	staticData := &employmenttypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EMPLOYMENTTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "employmenttype",
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

// EmploymentTypeParserInit initializes any static state used to implement EmploymentTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEmploymentTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EmploymentTypeParserInit() {
	staticData := &employmenttypeParserStaticData
	staticData.once.Do(employmenttypeParserInit)
}

// NewEmploymentTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEmploymentTypeParser(input antlr.TokenStream) *EmploymentTypeParser {
	EmploymentTypeParserInit()
	this := new(EmploymentTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &employmenttypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "EmploymentType.g4"

	return this
}

// EmploymentTypeParser tokens.
const (
	EmploymentTypeParserEOF            = antlr.TokenEOF
	EmploymentTypeParserEMPLOYMENTTYPE = 1
	EmploymentTypeParserOWNKEY         = 2
	EmploymentTypeParserSPLITKEY       = 3
	EmploymentTypeParserWS             = 4
)

// EmploymentTypeParser rules.
const (
	EmploymentTypeParserRULE_expression     = 0
	EmploymentTypeParserRULE_employmenttype = 1
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
	p.RuleIndex = EmploymentTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmploymentTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Employmenttype() IEmploymenttypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmploymenttypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmploymenttypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EmploymentTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EmploymentTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EmploymentTypeParserRULE_expression)

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
		p.Employmenttype()
	}
	{
		p.SetState(5)
		p.Match(EmploymentTypeParserEOF)
	}

	return localctx
}

// IEmploymenttypeContext is an interface to support dynamic dispatch.
type IEmploymenttypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmploymenttypeContext differentiates from other interfaces.
	IsEmploymenttypeContext()
}

type EmploymenttypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmploymenttypeContext() *EmploymenttypeContext {
	var p = new(EmploymenttypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EmploymentTypeParserRULE_employmenttype
	return p
}

func (*EmploymenttypeContext) IsEmploymenttypeContext() {}

func NewEmploymenttypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmploymenttypeContext {
	var p = new(EmploymenttypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmploymentTypeParserRULE_employmenttype

	return p
}

func (s *EmploymenttypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EmploymenttypeContext) EMPLOYMENTTYPE() antlr.TerminalNode {
	return s.GetToken(EmploymentTypeParserEMPLOYMENTTYPE, 0)
}

func (s *EmploymenttypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmploymenttypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmploymenttypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentTypeListener); ok {
		listenerT.EnterEmploymenttype(s)
	}
}

func (s *EmploymenttypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmploymentTypeListener); ok {
		listenerT.ExitEmploymenttype(s)
	}
}

func (p *EmploymentTypeParser) Employmenttype() (localctx IEmploymenttypeContext) {
	this := p
	_ = this

	localctx = NewEmploymenttypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EmploymentTypeParserRULE_employmenttype)

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
		p.Match(EmploymentTypeParserEMPLOYMENTTYPE)
	}

	return localctx
}
