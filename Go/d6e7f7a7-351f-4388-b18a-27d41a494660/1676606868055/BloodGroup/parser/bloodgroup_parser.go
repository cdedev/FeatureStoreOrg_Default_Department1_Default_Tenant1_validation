// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606868055/BloodGroup.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodGroup

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

type BloodGroupParser struct {
	*antlr.BaseParser
}

var bloodgroupParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bloodgroupParserInit() {
	staticData := &bloodgroupParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BLOODGROUP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bloodgroup",
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

// BloodGroupParserInit initializes any static state used to implement BloodGroupParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBloodGroupParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BloodGroupParserInit() {
	staticData := &bloodgroupParserStaticData
	staticData.once.Do(bloodgroupParserInit)
}

// NewBloodGroupParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBloodGroupParser(input antlr.TokenStream) *BloodGroupParser {
	BloodGroupParserInit()
	this := new(BloodGroupParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bloodgroupParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BloodGroup.g4"

	return this
}

// BloodGroupParser tokens.
const (
	BloodGroupParserEOF        = antlr.TokenEOF
	BloodGroupParserBLOODGROUP = 1
	BloodGroupParserOWNKEY     = 2
	BloodGroupParserSPLITKEY   = 3
	BloodGroupParserWS         = 4
)

// BloodGroupParser rules.
const (
	BloodGroupParserRULE_expression = 0
	BloodGroupParserRULE_bloodgroup = 1
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
	p.RuleIndex = BloodGroupParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BloodGroupParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bloodgroup() IBloodgroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloodgroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloodgroupContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BloodGroupParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodGroupListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodGroupListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BloodGroupParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BloodGroupParserRULE_expression)

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
		p.Bloodgroup()
	}
	{
		p.SetState(5)
		p.Match(BloodGroupParserEOF)
	}

	return localctx
}

// IBloodgroupContext is an interface to support dynamic dispatch.
type IBloodgroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloodgroupContext differentiates from other interfaces.
	IsBloodgroupContext()
}

type BloodgroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloodgroupContext() *BloodgroupContext {
	var p = new(BloodgroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BloodGroupParserRULE_bloodgroup
	return p
}

func (*BloodgroupContext) IsBloodgroupContext() {}

func NewBloodgroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloodgroupContext {
	var p = new(BloodgroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BloodGroupParserRULE_bloodgroup

	return p
}

func (s *BloodgroupContext) GetParser() antlr.Parser { return s.parser }

func (s *BloodgroupContext) BLOODGROUP() antlr.TerminalNode {
	return s.GetToken(BloodGroupParserBLOODGROUP, 0)
}

func (s *BloodgroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloodgroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloodgroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodGroupListener); ok {
		listenerT.EnterBloodgroup(s)
	}
}

func (s *BloodgroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BloodGroupListener); ok {
		listenerT.ExitBloodgroup(s)
	}
}

func (p *BloodGroupParser) Bloodgroup() (localctx IBloodgroupContext) {
	this := p
	_ = this

	localctx = NewBloodgroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BloodGroupParserRULE_bloodgroup)

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
		p.Match(BloodGroupParserBLOODGROUP)
	}

	return localctx
}
