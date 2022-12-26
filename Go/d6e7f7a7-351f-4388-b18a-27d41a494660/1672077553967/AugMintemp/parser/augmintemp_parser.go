// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077553967/AugMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AugMintemp

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

type AugMintempParser struct {
	*antlr.BaseParser
}

var augmintempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func augmintempParserInit() {
	staticData := &augmintempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AUGMINTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "augmintemp",
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

// AugMintempParserInit initializes any static state used to implement AugMintempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAugMintempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AugMintempParserInit() {
	staticData := &augmintempParserStaticData
	staticData.once.Do(augmintempParserInit)
}

// NewAugMintempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAugMintempParser(input antlr.TokenStream) *AugMintempParser {
	AugMintempParserInit()
	this := new(AugMintempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &augmintempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AugMintemp.g4"

	return this
}

// AugMintempParser tokens.
const (
	AugMintempParserEOF        = antlr.TokenEOF
	AugMintempParserAUGMINTEMP = 1
	AugMintempParserOWNKEY     = 2
	AugMintempParserSPLITKEY   = 3
	AugMintempParserWS         = 4
)

// AugMintempParser rules.
const (
	AugMintempParserRULE_expression = 0
	AugMintempParserRULE_augmintemp = 1
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
	p.RuleIndex = AugMintempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AugMintempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Augmintemp() IAugmintempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAugmintempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAugmintempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AugMintempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMintempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMintempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AugMintempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AugMintempParserRULE_expression)

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
		p.Augmintemp()
	}
	{
		p.SetState(5)
		p.Match(AugMintempParserEOF)
	}

	return localctx
}

// IAugmintempContext is an interface to support dynamic dispatch.
type IAugmintempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAugmintempContext differentiates from other interfaces.
	IsAugmintempContext()
}

type AugmintempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAugmintempContext() *AugmintempContext {
	var p = new(AugmintempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AugMintempParserRULE_augmintemp
	return p
}

func (*AugmintempContext) IsAugmintempContext() {}

func NewAugmintempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AugmintempContext {
	var p = new(AugmintempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AugMintempParserRULE_augmintemp

	return p
}

func (s *AugmintempContext) GetParser() antlr.Parser { return s.parser }

func (s *AugmintempContext) AUGMINTEMP() antlr.TerminalNode {
	return s.GetToken(AugMintempParserAUGMINTEMP, 0)
}

func (s *AugmintempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AugmintempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AugmintempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMintempListener); ok {
		listenerT.EnterAugmintemp(s)
	}
}

func (s *AugmintempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AugMintempListener); ok {
		listenerT.ExitAugmintemp(s)
	}
}

func (p *AugMintempParser) Augmintemp() (localctx IAugmintempContext) {
	this := p
	_ = this

	localctx = NewAugmintempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AugMintempParserRULE_augmintemp)

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
		p.Match(AugMintempParserAUGMINTEMP)
	}

	return localctx
}
