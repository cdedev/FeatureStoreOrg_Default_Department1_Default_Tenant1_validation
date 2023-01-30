// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675074867543/Ethnicity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ethnicity

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

type EthnicityParser struct {
	*antlr.BaseParser
}

var ethnicityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ethnicityParserInit() {
	staticData := &ethnicityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ETHNICITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "ethnicity",
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

// EthnicityParserInit initializes any static state used to implement EthnicityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEthnicityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EthnicityParserInit() {
	staticData := &ethnicityParserStaticData
	staticData.once.Do(ethnicityParserInit)
}

// NewEthnicityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEthnicityParser(input antlr.TokenStream) *EthnicityParser {
	EthnicityParserInit()
	this := new(EthnicityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ethnicityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Ethnicity.g4"

	return this
}

// EthnicityParser tokens.
const (
	EthnicityParserEOF       = antlr.TokenEOF
	EthnicityParserETHNICITY = 1
	EthnicityParserOWNKEY    = 2
	EthnicityParserSPLITKEY  = 3
	EthnicityParserWS        = 4
)

// EthnicityParser rules.
const (
	EthnicityParserRULE_expression = 0
	EthnicityParserRULE_ethnicity  = 1
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
	p.RuleIndex = EthnicityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EthnicityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Ethnicity() IEthnicityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEthnicityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEthnicityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EthnicityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EthnicityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EthnicityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EthnicityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EthnicityParserRULE_expression)

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
		p.Ethnicity()
	}
	{
		p.SetState(5)
		p.Match(EthnicityParserEOF)
	}

	return localctx
}

// IEthnicityContext is an interface to support dynamic dispatch.
type IEthnicityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEthnicityContext differentiates from other interfaces.
	IsEthnicityContext()
}

type EthnicityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEthnicityContext() *EthnicityContext {
	var p = new(EthnicityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EthnicityParserRULE_ethnicity
	return p
}

func (*EthnicityContext) IsEthnicityContext() {}

func NewEthnicityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EthnicityContext {
	var p = new(EthnicityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EthnicityParserRULE_ethnicity

	return p
}

func (s *EthnicityContext) GetParser() antlr.Parser { return s.parser }

func (s *EthnicityContext) ETHNICITY() antlr.TerminalNode {
	return s.GetToken(EthnicityParserETHNICITY, 0)
}

func (s *EthnicityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EthnicityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EthnicityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EthnicityListener); ok {
		listenerT.EnterEthnicity(s)
	}
}

func (s *EthnicityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EthnicityListener); ok {
		listenerT.ExitEthnicity(s)
	}
}

func (p *EthnicityParser) Ethnicity() (localctx IEthnicityContext) {
	this := p
	_ = this

	localctx = NewEthnicityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EthnicityParserRULE_ethnicity)

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
		p.Match(EthnicityParserETHNICITY)
	}

	return localctx
}
