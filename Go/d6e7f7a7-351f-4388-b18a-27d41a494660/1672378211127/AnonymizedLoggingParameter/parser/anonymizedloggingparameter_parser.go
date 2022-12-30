// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672378211127/AnonymizedLoggingParameter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AnonymizedLoggingParameter

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

type AnonymizedLoggingParameterParser struct {
	*antlr.BaseParser
}

var anonymizedloggingparameterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func anonymizedloggingparameterParserInit() {
	staticData := &anonymizedloggingparameterParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ANONYMIZEDLOGGINGPARAMETER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "anonymizedloggingparameter",
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

// AnonymizedLoggingParameterParserInit initializes any static state used to implement AnonymizedLoggingParameterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAnonymizedLoggingParameterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AnonymizedLoggingParameterParserInit() {
	staticData := &anonymizedloggingparameterParserStaticData
	staticData.once.Do(anonymizedloggingparameterParserInit)
}

// NewAnonymizedLoggingParameterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAnonymizedLoggingParameterParser(input antlr.TokenStream) *AnonymizedLoggingParameterParser {
	AnonymizedLoggingParameterParserInit()
	this := new(AnonymizedLoggingParameterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &anonymizedloggingparameterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "AnonymizedLoggingParameter.g4"

	return this
}

// AnonymizedLoggingParameterParser tokens.
const (
	AnonymizedLoggingParameterParserEOF                        = antlr.TokenEOF
	AnonymizedLoggingParameterParserANONYMIZEDLOGGINGPARAMETER = 1
	AnonymizedLoggingParameterParserOWNKEY                     = 2
	AnonymizedLoggingParameterParserSPLITKEY                   = 3
	AnonymizedLoggingParameterParserWS                         = 4
)

// AnonymizedLoggingParameterParser rules.
const (
	AnonymizedLoggingParameterParserRULE_expression                 = 0
	AnonymizedLoggingParameterParserRULE_anonymizedloggingparameter = 1
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
	p.RuleIndex = AnonymizedLoggingParameterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnonymizedLoggingParameterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Anonymizedloggingparameter() IAnonymizedloggingparameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnonymizedloggingparameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnonymizedloggingparameterContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AnonymizedLoggingParameterParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnonymizedLoggingParameterListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnonymizedLoggingParameterListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AnonymizedLoggingParameterParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AnonymizedLoggingParameterParserRULE_expression)

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
		p.Anonymizedloggingparameter()
	}
	{
		p.SetState(5)
		p.Match(AnonymizedLoggingParameterParserEOF)
	}

	return localctx
}

// IAnonymizedloggingparameterContext is an interface to support dynamic dispatch.
type IAnonymizedloggingparameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnonymizedloggingparameterContext differentiates from other interfaces.
	IsAnonymizedloggingparameterContext()
}

type AnonymizedloggingparameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnonymizedloggingparameterContext() *AnonymizedloggingparameterContext {
	var p = new(AnonymizedloggingparameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnonymizedLoggingParameterParserRULE_anonymizedloggingparameter
	return p
}

func (*AnonymizedloggingparameterContext) IsAnonymizedloggingparameterContext() {}

func NewAnonymizedloggingparameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnonymizedloggingparameterContext {
	var p = new(AnonymizedloggingparameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnonymizedLoggingParameterParserRULE_anonymizedloggingparameter

	return p
}

func (s *AnonymizedloggingparameterContext) GetParser() antlr.Parser { return s.parser }

func (s *AnonymizedloggingparameterContext) ANONYMIZEDLOGGINGPARAMETER() antlr.TerminalNode {
	return s.GetToken(AnonymizedLoggingParameterParserANONYMIZEDLOGGINGPARAMETER, 0)
}

func (s *AnonymizedloggingparameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnonymizedloggingparameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnonymizedloggingparameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnonymizedLoggingParameterListener); ok {
		listenerT.EnterAnonymizedloggingparameter(s)
	}
}

func (s *AnonymizedloggingparameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnonymizedLoggingParameterListener); ok {
		listenerT.ExitAnonymizedloggingparameter(s)
	}
}

func (p *AnonymizedLoggingParameterParser) Anonymizedloggingparameter() (localctx IAnonymizedloggingparameterContext) {
	this := p
	_ = this

	localctx = NewAnonymizedloggingparameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AnonymizedLoggingParameterParserRULE_anonymizedloggingparameter)

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
		p.Match(AnonymizedLoggingParameterParserANONYMIZEDLOGGINGPARAMETER)
	}

	return localctx
}
