// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077078876/AprMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AprMaxtemp

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

type AprMaxtempParser struct {
	*antlr.BaseParser
}

var aprmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aprmaxtempParserInit() {
	staticData := &aprmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "APRMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "aprmaxtemp",
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

// AprMaxtempParserInit initializes any static state used to implement AprMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAprMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AprMaxtempParserInit() {
	staticData := &aprmaxtempParserStaticData
	staticData.once.Do(aprmaxtempParserInit)
}

// NewAprMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAprMaxtempParser(input antlr.TokenStream) *AprMaxtempParser {
	AprMaxtempParserInit()
	this := new(AprMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &aprmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AprMaxtemp.g4"

	return this
}

// AprMaxtempParser tokens.
const (
	AprMaxtempParserEOF        = antlr.TokenEOF
	AprMaxtempParserAPRMAXTEMP = 1
	AprMaxtempParserOWNKEY     = 2
	AprMaxtempParserSPLITKEY   = 3
	AprMaxtempParserWS         = 4
)

// AprMaxtempParser rules.
const (
	AprMaxtempParserRULE_expression = 0
	AprMaxtempParserRULE_aprmaxtemp = 1
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
	p.RuleIndex = AprMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AprMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Aprmaxtemp() IAprmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAprmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAprmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AprMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AprMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AprMaxtempParserRULE_expression)

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
		p.Aprmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(AprMaxtempParserEOF)
	}

	return localctx
}

// IAprmaxtempContext is an interface to support dynamic dispatch.
type IAprmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAprmaxtempContext differentiates from other interfaces.
	IsAprmaxtempContext()
}

type AprmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAprmaxtempContext() *AprmaxtempContext {
	var p = new(AprmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AprMaxtempParserRULE_aprmaxtemp
	return p
}

func (*AprmaxtempContext) IsAprmaxtempContext() {}

func NewAprmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AprmaxtempContext {
	var p = new(AprmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AprMaxtempParserRULE_aprmaxtemp

	return p
}

func (s *AprmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *AprmaxtempContext) APRMAXTEMP() antlr.TerminalNode {
	return s.GetToken(AprMaxtempParserAPRMAXTEMP, 0)
}

func (s *AprmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AprmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AprmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMaxtempListener); ok {
		listenerT.EnterAprmaxtemp(s)
	}
}

func (s *AprmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMaxtempListener); ok {
		listenerT.ExitAprmaxtemp(s)
	}
}

func (p *AprMaxtempParser) Aprmaxtemp() (localctx IAprmaxtempContext) {
	this := p
	_ = this

	localctx = NewAprmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AprMaxtempParserRULE_aprmaxtemp)

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
		p.Match(AprMaxtempParserAPRMAXTEMP)
	}

	return localctx
}
