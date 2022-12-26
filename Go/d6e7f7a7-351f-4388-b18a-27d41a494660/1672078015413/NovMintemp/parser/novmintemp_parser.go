// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078015413/NovMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NovMintemp

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

type NovMintempParser struct {
	*antlr.BaseParser
}

var novmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func novmintempParserInit() {
	staticData := &novmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NOVMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "novmintemp",
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

// NovMintempParserInit initializes any static state used to implement NovMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNovMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NovMintempParserInit() {
	staticData := &novmintempParserStaticData
	staticData.once.Do(novmintempParserInit)
}

// NewNovMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNovMintempParser(input antlr.TokenStream) *NovMintempParser {
	NovMintempParserInit()
	this := new(NovMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &novmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "NovMintemp.g4"

	return this
}

// NovMintempParser tokens.
const (
	NovMintempParserEOF        = antlr.TokenEOF
	NovMintempParserNOVMINTEMP = 1
	NovMintempParserOWNKEY     = 2
	NovMintempParserSPLITKEY   = 3
	NovMintempParserWS         = 4
)

// NovMintempParser rules.
const (
	NovMintempParserRULE_expression = 0
	NovMintempParserRULE_novmintemp = 1
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
	p.RuleIndex = NovMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NovMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Novmintemp() INovmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INovmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INovmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NovMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NovMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NovMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NovMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NovMintempParserRULE_expression)

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
		p.Novmintemp()
	}
	{
		p.SetState(5)
		p.Match(NovMintempParserEOF)
	}

	return localctx
}

// INovmintempContext is an interface to support dynamic dispatch.
type INovmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNovmintempContext differentiates from other interfaces.
	IsNovmintempContext()
}

type NovmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNovmintempContext() *NovmintempContext {
	var p = new(NovmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NovMintempParserRULE_novmintemp
	return p
}

func (*NovmintempContext) IsNovmintempContext() {}

func NewNovmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NovmintempContext {
	var p = new(NovmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NovMintempParserRULE_novmintemp

	return p
}

func (s *NovmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *NovmintempContext) NOVMINTEMP() antlr.TerminalNode {
	return s.GetToken(NovMintempParserNOVMINTEMP, 0)
}

func (s *NovmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NovmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NovmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NovMintempListener); ok {
		listenerT.EnterNovmintemp(s)
	}
}

func (s *NovmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NovMintempListener); ok {
		listenerT.ExitNovmintemp(s)
	}
}

func (p *NovMintempParser) Novmintemp() (localctx INovmintempContext) {
	this := p
	_ = this

	localctx = NewNovmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NovMintempParserRULE_novmintemp)

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
		p.Match(NovMintempParserNOVMINTEMP)
	}

	return localctx
}
