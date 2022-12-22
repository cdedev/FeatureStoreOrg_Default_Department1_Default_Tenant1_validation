// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698569487/YearsInCurrentRole.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // YearsInCurrentRole

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

type YearsInCurrentRoleParser struct {
	*antlr.BaseParser
}

var yearsincurrentroleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func yearsincurrentroleParserInit() {
	staticData := &yearsincurrentroleParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "YEARSINCURRENTROLE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "yearsincurrentrole",
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

// YearsInCurrentRoleParserInit initializes any static state used to implement YearsInCurrentRoleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewYearsInCurrentRoleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func YearsInCurrentRoleParserInit() {
	staticData := &yearsincurrentroleParserStaticData
	staticData.once.Do(yearsincurrentroleParserInit)
}

// NewYearsInCurrentRoleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewYearsInCurrentRoleParser(input antlr.TokenStream) *YearsInCurrentRoleParser {
	YearsInCurrentRoleParserInit()
	this := new(YearsInCurrentRoleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &yearsincurrentroleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "YearsInCurrentRole.g4"

	return this
}

// YearsInCurrentRoleParser tokens.
const (
	YearsInCurrentRoleParserEOF                = antlr.TokenEOF
	YearsInCurrentRoleParserYEARSINCURRENTROLE = 1
	YearsInCurrentRoleParserOWNKEY             = 2
	YearsInCurrentRoleParserSPLITKEY           = 3
	YearsInCurrentRoleParserWS                 = 4
)

// YearsInCurrentRoleParser rules.
const (
	YearsInCurrentRoleParserRULE_expression         = 0
	YearsInCurrentRoleParserRULE_yearsincurrentrole = 1
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
	p.RuleIndex = YearsInCurrentRoleParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YearsInCurrentRoleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Yearsincurrentrole() IYearsincurrentroleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IYearsincurrentroleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IYearsincurrentroleContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(YearsInCurrentRoleParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearsInCurrentRoleListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearsInCurrentRoleListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *YearsInCurrentRoleParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, YearsInCurrentRoleParserRULE_expression)

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
		p.Yearsincurrentrole()
	}
	{
		p.SetState(5)
		p.Match(YearsInCurrentRoleParserEOF)
	}

	return localctx
}

// IYearsincurrentroleContext is an interface to support dynamic dispatch.
type IYearsincurrentroleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearsincurrentroleContext differentiates from other interfaces.
	IsYearsincurrentroleContext()
}

type YearsincurrentroleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearsincurrentroleContext() *YearsincurrentroleContext {
	var p = new(YearsincurrentroleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = YearsInCurrentRoleParserRULE_yearsincurrentrole
	return p
}

func (*YearsincurrentroleContext) IsYearsincurrentroleContext() {}

func NewYearsincurrentroleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearsincurrentroleContext {
	var p = new(YearsincurrentroleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = YearsInCurrentRoleParserRULE_yearsincurrentrole

	return p
}

func (s *YearsincurrentroleContext) GetParser() antlr.Parser { return s.parser }

func (s *YearsincurrentroleContext) YEARSINCURRENTROLE() antlr.TerminalNode {
	return s.GetToken(YearsInCurrentRoleParserYEARSINCURRENTROLE, 0)
}

func (s *YearsincurrentroleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearsincurrentroleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *YearsincurrentroleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearsInCurrentRoleListener); ok {
		listenerT.EnterYearsincurrentrole(s)
	}
}

func (s *YearsincurrentroleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(YearsInCurrentRoleListener); ok {
		listenerT.ExitYearsincurrentrole(s)
	}
}

func (p *YearsInCurrentRoleParser) Yearsincurrentrole() (localctx IYearsincurrentroleContext) {
	this := p
	_ = this

	localctx = NewYearsincurrentroleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, YearsInCurrentRoleParserRULE_yearsincurrentrole)

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
		p.Match(YearsInCurrentRoleParserYEARSINCURRENTROLE)
	}

	return localctx
}
