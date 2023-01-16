// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878269550/Paresthesia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Paresthesia

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

type ParesthesiaParser struct {
	*antlr.BaseParser
}

var paresthesiaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paresthesiaParserInit() {
	staticData := &paresthesiaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PARESTHESIA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "paresthesia",
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

// ParesthesiaParserInit initializes any static state used to implement ParesthesiaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParesthesiaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParesthesiaParserInit() {
	staticData := &paresthesiaParserStaticData
	staticData.once.Do(paresthesiaParserInit)
}

// NewParesthesiaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParesthesiaParser(input antlr.TokenStream) *ParesthesiaParser {
	ParesthesiaParserInit()
	this := new(ParesthesiaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &paresthesiaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Paresthesia.g4"

	return this
}

// ParesthesiaParser tokens.
const (
	ParesthesiaParserEOF         = antlr.TokenEOF
	ParesthesiaParserPARESTHESIA = 1
	ParesthesiaParserOWNKEY      = 2
	ParesthesiaParserSPLITKEY    = 3
	ParesthesiaParserWS          = 4
)

// ParesthesiaParser rules.
const (
	ParesthesiaParserRULE_expression  = 0
	ParesthesiaParserRULE_paresthesia = 1
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
	p.RuleIndex = ParesthesiaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParesthesiaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Paresthesia() IParesthesiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParesthesiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParesthesiaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParesthesiaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParesthesiaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParesthesiaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ParesthesiaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParesthesiaParserRULE_expression)

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
		p.Paresthesia()
	}
	{
		p.SetState(5)
		p.Match(ParesthesiaParserEOF)
	}

	return localctx
}

// IParesthesiaContext is an interface to support dynamic dispatch.
type IParesthesiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParesthesiaContext differentiates from other interfaces.
	IsParesthesiaContext()
}

type ParesthesiaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParesthesiaContext() *ParesthesiaContext {
	var p = new(ParesthesiaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParesthesiaParserRULE_paresthesia
	return p
}

func (*ParesthesiaContext) IsParesthesiaContext() {}

func NewParesthesiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParesthesiaContext {
	var p = new(ParesthesiaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParesthesiaParserRULE_paresthesia

	return p
}

func (s *ParesthesiaContext) GetParser() antlr.Parser { return s.parser }

func (s *ParesthesiaContext) PARESTHESIA() antlr.TerminalNode {
	return s.GetToken(ParesthesiaParserPARESTHESIA, 0)
}

func (s *ParesthesiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParesthesiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParesthesiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParesthesiaListener); ok {
		listenerT.EnterParesthesia(s)
	}
}

func (s *ParesthesiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParesthesiaListener); ok {
		listenerT.ExitParesthesia(s)
	}
}

func (p *ParesthesiaParser) Paresthesia() (localctx IParesthesiaContext) {
	this := p
	_ = this

	localctx = NewParesthesiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParesthesiaParserRULE_paresthesia)

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
		p.Match(ParesthesiaParserPARESTHESIA)
	}

	return localctx
}
