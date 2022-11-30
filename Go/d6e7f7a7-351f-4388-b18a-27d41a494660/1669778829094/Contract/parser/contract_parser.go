// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778829094/Contract.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contract

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

type ContractParser struct {
	*antlr.BaseParser
}

var contractParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func contractParserInit() {
	staticData := &contractParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONTRACT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "contract",
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

// ContractParserInit initializes any static state used to implement ContractParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewContractParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ContractParserInit() {
	staticData := &contractParserStaticData
	staticData.once.Do(contractParserInit)
}

// NewContractParser produces a new parser instance for the optional input antlr.TokenStream.
func NewContractParser(input antlr.TokenStream) *ContractParser {
	ContractParserInit()
	this := new(ContractParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &contractParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Contract.g4"

	return this
}

// ContractParser tokens.
const (
	ContractParserEOF      = antlr.TokenEOF
	ContractParserCONTRACT = 1
	ContractParserOWNKEY   = 2
	ContractParserSPLITKEY = 3
	ContractParserWS       = 4
)

// ContractParser rules.
const (
	ContractParserRULE_expression = 0
	ContractParserRULE_contract   = 1
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
	p.RuleIndex = ContractParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ContractParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Contract() IContractContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContractContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContractContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ContractParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContractListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContractListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ContractParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ContractParserRULE_expression)

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
		p.Contract()
	}
	{
		p.SetState(5)
		p.Match(ContractParserEOF)
	}

	return localctx
}

// IContractContext is an interface to support dynamic dispatch.
type IContractContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContractContext differentiates from other interfaces.
	IsContractContext()
}

type ContractContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContractContext() *ContractContext {
	var p = new(ContractContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ContractParserRULE_contract
	return p
}

func (*ContractContext) IsContractContext() {}

func NewContractContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContractContext {
	var p = new(ContractContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ContractParserRULE_contract

	return p
}

func (s *ContractContext) GetParser() antlr.Parser { return s.parser }

func (s *ContractContext) CONTRACT() antlr.TerminalNode {
	return s.GetToken(ContractParserCONTRACT, 0)
}

func (s *ContractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContractContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContractListener); ok {
		listenerT.EnterContract(s)
	}
}

func (s *ContractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContractListener); ok {
		listenerT.ExitContract(s)
	}
}

func (p *ContractParser) Contract() (localctx IContractContext) {
	this := p
	_ = this

	localctx = NewContractContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ContractParserRULE_contract)

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
		p.Match(ContractParserCONTRACT)
	}

	return localctx
}
