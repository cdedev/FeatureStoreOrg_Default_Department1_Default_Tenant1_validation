// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673249142899/Genus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Genus

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

type GenusParser struct {
	*antlr.BaseParser
}

var genusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func genusParserInit() {
	staticData := &genusParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GENUS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "genus",
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

// GenusParserInit initializes any static state used to implement GenusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGenusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GenusParserInit() {
	staticData := &genusParserStaticData
	staticData.once.Do(genusParserInit)
}

// NewGenusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGenusParser(input antlr.TokenStream) *GenusParser {
	GenusParserInit()
	this := new(GenusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &genusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Genus.g4"

	return this
}

// GenusParser tokens.
const (
	GenusParserEOF      = antlr.TokenEOF
	GenusParserGENUS    = 1
	GenusParserOWNKEY   = 2
	GenusParserSPLITKEY = 3
	GenusParserWS       = 4
)

// GenusParser rules.
const (
	GenusParserRULE_expression = 0
	GenusParserRULE_genus      = 1
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
	p.RuleIndex = GenusParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GenusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Genus() IGenusContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGenusContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGenusContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GenusParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenusListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenusListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GenusParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GenusParserRULE_expression)

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
		p.Genus()
	}
	{
		p.SetState(5)
		p.Match(GenusParserEOF)
	}

	return localctx
}

// IGenusContext is an interface to support dynamic dispatch.
type IGenusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGenusContext differentiates from other interfaces.
	IsGenusContext()
}

type GenusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGenusContext() *GenusContext {
	var p = new(GenusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GenusParserRULE_genus
	return p
}

func (*GenusContext) IsGenusContext() {}

func NewGenusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GenusContext {
	var p = new(GenusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GenusParserRULE_genus

	return p
}

func (s *GenusContext) GetParser() antlr.Parser { return s.parser }

func (s *GenusContext) GENUS() antlr.TerminalNode {
	return s.GetToken(GenusParserGENUS, 0)
}

func (s *GenusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GenusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GenusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenusListener); ok {
		listenerT.EnterGenus(s)
	}
}

func (s *GenusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GenusListener); ok {
		listenerT.ExitGenus(s)
	}
}

func (p *GenusParser) Genus() (localctx IGenusContext) {
	this := p
	_ = this

	localctx = NewGenusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GenusParserRULE_genus)

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
		p.Match(GenusParserGENUS)
	}

	return localctx
}
