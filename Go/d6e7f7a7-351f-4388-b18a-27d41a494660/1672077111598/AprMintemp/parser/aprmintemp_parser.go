// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077111598/AprMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AprMintemp

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

type AprMintempParser struct {
	*antlr.BaseParser
}

var aprmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aprmintempParserInit() {
	staticData := &aprmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "APRMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "aprmintemp",
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

// AprMintempParserInit initializes any static state used to implement AprMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAprMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AprMintempParserInit() {
	staticData := &aprmintempParserStaticData
	staticData.once.Do(aprmintempParserInit)
}

// NewAprMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAprMintempParser(input antlr.TokenStream) *AprMintempParser {
	AprMintempParserInit()
	this := new(AprMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &aprmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AprMintemp.g4"

	return this
}

// AprMintempParser tokens.
const (
	AprMintempParserEOF        = antlr.TokenEOF
	AprMintempParserAPRMINTEMP = 1
	AprMintempParserOWNKEY     = 2
	AprMintempParserSPLITKEY   = 3
	AprMintempParserWS         = 4
)

// AprMintempParser rules.
const (
	AprMintempParserRULE_expression = 0
	AprMintempParserRULE_aprmintemp = 1
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
	p.RuleIndex = AprMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AprMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Aprmintemp() IAprmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAprmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAprmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AprMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AprMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AprMintempParserRULE_expression)

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
		p.Aprmintemp()
	}
	{
		p.SetState(5)
		p.Match(AprMintempParserEOF)
	}

	return localctx
}

// IAprmintempContext is an interface to support dynamic dispatch.
type IAprmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAprmintempContext differentiates from other interfaces.
	IsAprmintempContext()
}

type AprmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAprmintempContext() *AprmintempContext {
	var p = new(AprmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AprMintempParserRULE_aprmintemp
	return p
}

func (*AprmintempContext) IsAprmintempContext() {}

func NewAprmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AprmintempContext {
	var p = new(AprmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AprMintempParserRULE_aprmintemp

	return p
}

func (s *AprmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *AprmintempContext) APRMINTEMP() antlr.TerminalNode {
	return s.GetToken(AprMintempParserAPRMINTEMP, 0)
}

func (s *AprmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AprmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AprmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMintempListener); ok {
		listenerT.EnterAprmintemp(s)
	}
}

func (s *AprmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AprMintempListener); ok {
		listenerT.ExitAprmintemp(s)
	}
}

func (p *AprMintempParser) Aprmintemp() (localctx IAprmintempContext) {
	this := p
	_ = this

	localctx = NewAprmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AprMintempParserRULE_aprmintemp)

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
		p.Match(AprMintempParserAPRMINTEMP)
	}

	return localctx
}
