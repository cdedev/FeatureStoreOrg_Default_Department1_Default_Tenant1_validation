// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673924919318/Aquatic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Aquatic

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

type AquaticParser struct {
	*antlr.BaseParser
}

var aquaticParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aquaticParserInit() {
	staticData := &aquaticParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AQUATIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "aquatic",
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

// AquaticParserInit initializes any static state used to implement AquaticParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAquaticParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AquaticParserInit() {
	staticData := &aquaticParserStaticData
	staticData.once.Do(aquaticParserInit)
}

// NewAquaticParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAquaticParser(input antlr.TokenStream) *AquaticParser {
	AquaticParserInit()
	this := new(AquaticParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &aquaticParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Aquatic.g4"

	return this
}

// AquaticParser tokens.
const (
	AquaticParserEOF      = antlr.TokenEOF
	AquaticParserAQUATIC  = 1
	AquaticParserOWNKEY   = 2
	AquaticParserSPLITKEY = 3
	AquaticParserWS       = 4
)

// AquaticParser rules.
const (
	AquaticParserRULE_expression = 0
	AquaticParserRULE_aquatic    = 1
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
	p.RuleIndex = AquaticParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AquaticParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Aquatic() IAquaticContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAquaticContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAquaticContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AquaticParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AquaticListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AquaticListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AquaticParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AquaticParserRULE_expression)

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
		p.Aquatic()
	}
	{
		p.SetState(5)
		p.Match(AquaticParserEOF)
	}

	return localctx
}

// IAquaticContext is an interface to support dynamic dispatch.
type IAquaticContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAquaticContext differentiates from other interfaces.
	IsAquaticContext()
}

type AquaticContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAquaticContext() *AquaticContext {
	var p = new(AquaticContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AquaticParserRULE_aquatic
	return p
}

func (*AquaticContext) IsAquaticContext() {}

func NewAquaticContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AquaticContext {
	var p = new(AquaticContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AquaticParserRULE_aquatic

	return p
}

func (s *AquaticContext) GetParser() antlr.Parser { return s.parser }

func (s *AquaticContext) AQUATIC() antlr.TerminalNode {
	return s.GetToken(AquaticParserAQUATIC, 0)
}

func (s *AquaticContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AquaticContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AquaticContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AquaticListener); ok {
		listenerT.EnterAquatic(s)
	}
}

func (s *AquaticContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AquaticListener); ok {
		listenerT.ExitAquatic(s)
	}
}

func (p *AquaticParser) Aquatic() (localctx IAquaticContext) {
	this := p
	_ = this

	localctx = NewAquaticContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AquaticParserRULE_aquatic)

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
		p.Match(AquaticParserAQUATIC)
	}

	return localctx
}
