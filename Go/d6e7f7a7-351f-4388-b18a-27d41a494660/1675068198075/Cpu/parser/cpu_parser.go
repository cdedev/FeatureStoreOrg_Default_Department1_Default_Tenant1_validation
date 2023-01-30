// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675068198075/Cpu.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cpu

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

type CpuParser struct {
	*antlr.BaseParser
}

var cpuParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cpuParserInit() {
	staticData := &cpuParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CPU", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cpu",
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

// CpuParserInit initializes any static state used to implement CpuParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCpuParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CpuParserInit() {
	staticData := &cpuParserStaticData
	staticData.once.Do(cpuParserInit)
}

// NewCpuParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCpuParser(input antlr.TokenStream) *CpuParser {
	CpuParserInit()
	this := new(CpuParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cpuParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cpu.g4"

	return this
}

// CpuParser tokens.
const (
	CpuParserEOF      = antlr.TokenEOF
	CpuParserCPU      = 1
	CpuParserOWNKEY   = 2
	CpuParserSPLITKEY = 3
	CpuParserWS       = 4
)

// CpuParser rules.
const (
	CpuParserRULE_expression = 0
	CpuParserRULE_cpu        = 1
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
	p.RuleIndex = CpuParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CpuParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cpu() ICpuContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICpuContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICpuContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CpuParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CpuListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CpuListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CpuParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CpuParserRULE_expression)

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
		p.Cpu()
	}
	{
		p.SetState(5)
		p.Match(CpuParserEOF)
	}

	return localctx
}

// ICpuContext is an interface to support dynamic dispatch.
type ICpuContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCpuContext differentiates from other interfaces.
	IsCpuContext()
}

type CpuContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCpuContext() *CpuContext {
	var p = new(CpuContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CpuParserRULE_cpu
	return p
}

func (*CpuContext) IsCpuContext() {}

func NewCpuContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CpuContext {
	var p = new(CpuContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CpuParserRULE_cpu

	return p
}

func (s *CpuContext) GetParser() antlr.Parser { return s.parser }

func (s *CpuContext) CPU() antlr.TerminalNode {
	return s.GetToken(CpuParserCPU, 0)
}

func (s *CpuContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CpuContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CpuContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CpuListener); ok {
		listenerT.EnterCpu(s)
	}
}

func (s *CpuContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CpuListener); ok {
		listenerT.ExitCpu(s)
	}
}

func (p *CpuParser) Cpu() (localctx ICpuContext) {
	this := p
	_ = this

	localctx = NewCpuContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CpuParserRULE_cpu)

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
		p.Match(CpuParserCPU)
	}

	return localctx
}
