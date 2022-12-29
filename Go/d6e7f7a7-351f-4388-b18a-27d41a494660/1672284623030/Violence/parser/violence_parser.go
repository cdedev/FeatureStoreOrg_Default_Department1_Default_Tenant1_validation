// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284623030/Violence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Violence

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

type ViolenceParser struct {
	*antlr.BaseParser
}

var violenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func violenceParserInit() {
	staticData := &violenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VIOLENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "violence",
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

// ViolenceParserInit initializes any static state used to implement ViolenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewViolenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ViolenceParserInit() {
	staticData := &violenceParserStaticData
	staticData.once.Do(violenceParserInit)
}

// NewViolenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewViolenceParser(input antlr.TokenStream) *ViolenceParser {
	ViolenceParserInit()
	this := new(ViolenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &violenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Violence.g4"

	return this
}

// ViolenceParser tokens.
const (
	ViolenceParserEOF      = antlr.TokenEOF
	ViolenceParserVIOLENCE = 1
	ViolenceParserOWNKEY   = 2
	ViolenceParserSPLITKEY = 3
	ViolenceParserWS       = 4
)

// ViolenceParser rules.
const (
	ViolenceParserRULE_expression = 0
	ViolenceParserRULE_violence   = 1
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
	p.RuleIndex = ViolenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViolenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Violence() IViolenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IViolenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IViolenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ViolenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViolenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViolenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ViolenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ViolenceParserRULE_expression)

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
		p.Violence()
	}
	{
		p.SetState(5)
		p.Match(ViolenceParserEOF)
	}

	return localctx
}

// IViolenceContext is an interface to support dynamic dispatch.
type IViolenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsViolenceContext differentiates from other interfaces.
	IsViolenceContext()
}

type ViolenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyViolenceContext() *ViolenceContext {
	var p = new(ViolenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViolenceParserRULE_violence
	return p
}

func (*ViolenceContext) IsViolenceContext() {}

func NewViolenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ViolenceContext {
	var p = new(ViolenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViolenceParserRULE_violence

	return p
}

func (s *ViolenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ViolenceContext) VIOLENCE() antlr.TerminalNode {
	return s.GetToken(ViolenceParserVIOLENCE, 0)
}

func (s *ViolenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ViolenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ViolenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViolenceListener); ok {
		listenerT.EnterViolence(s)
	}
}

func (s *ViolenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViolenceListener); ok {
		listenerT.ExitViolence(s)
	}
}

func (p *ViolenceParser) Violence() (localctx IViolenceContext) {
	this := p
	_ = this

	localctx = NewViolenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ViolenceParserRULE_violence)

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
		p.Match(ViolenceParserVIOLENCE)
	}

	return localctx
}
