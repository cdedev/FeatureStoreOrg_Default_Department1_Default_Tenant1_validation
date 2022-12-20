// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528680321/Residence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Residence

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

type ResidenceParser struct {
	*antlr.BaseParser
}

var residenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func residenceParserInit() {
	staticData := &residenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RESIDENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "residence",
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

// ResidenceParserInit initializes any static state used to implement ResidenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewResidenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ResidenceParserInit() {
	staticData := &residenceParserStaticData
	staticData.once.Do(residenceParserInit)
}

// NewResidenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewResidenceParser(input antlr.TokenStream) *ResidenceParser {
	ResidenceParserInit()
	this := new(ResidenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &residenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Residence.g4"

	return this
}

// ResidenceParser tokens.
const (
	ResidenceParserEOF       = antlr.TokenEOF
	ResidenceParserRESIDENCE = 1
	ResidenceParserOWNKEY    = 2
	ResidenceParserSPLITKEY  = 3
	ResidenceParserWS        = 4
)

// ResidenceParser rules.
const (
	ResidenceParserRULE_expression = 0
	ResidenceParserRULE_residence  = 1
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
	p.RuleIndex = ResidenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ResidenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Residence() IResidenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResidenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResidenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ResidenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResidenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResidenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ResidenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ResidenceParserRULE_expression)

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
		p.Residence()
	}
	{
		p.SetState(5)
		p.Match(ResidenceParserEOF)
	}

	return localctx
}

// IResidenceContext is an interface to support dynamic dispatch.
type IResidenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResidenceContext differentiates from other interfaces.
	IsResidenceContext()
}

type ResidenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResidenceContext() *ResidenceContext {
	var p = new(ResidenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ResidenceParserRULE_residence
	return p
}

func (*ResidenceContext) IsResidenceContext() {}

func NewResidenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResidenceContext {
	var p = new(ResidenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ResidenceParserRULE_residence

	return p
}

func (s *ResidenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ResidenceContext) RESIDENCE() antlr.TerminalNode {
	return s.GetToken(ResidenceParserRESIDENCE, 0)
}

func (s *ResidenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResidenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResidenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResidenceListener); ok {
		listenerT.EnterResidence(s)
	}
}

func (s *ResidenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResidenceListener); ok {
		listenerT.ExitResidence(s)
	}
}

func (p *ResidenceParser) Residence() (localctx IResidenceContext) {
	this := p
	_ = this

	localctx = NewResidenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ResidenceParserRULE_residence)

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
		p.Match(ResidenceParserRESIDENCE)
	}

	return localctx
}
