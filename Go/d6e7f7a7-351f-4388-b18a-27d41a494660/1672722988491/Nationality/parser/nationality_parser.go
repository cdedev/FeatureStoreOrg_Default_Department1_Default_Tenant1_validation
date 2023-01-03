// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672722988491/Nationality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nationality

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

type NationalityParser struct {
	*antlr.BaseParser
}

var nationalityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nationalityParserInit() {
	staticData := &nationalityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NATIONALITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "nationality",
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

// NationalityParserInit initializes any static state used to implement NationalityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNationalityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NationalityParserInit() {
	staticData := &nationalityParserStaticData
	staticData.once.Do(nationalityParserInit)
}

// NewNationalityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNationalityParser(input antlr.TokenStream) *NationalityParser {
	NationalityParserInit()
	this := new(NationalityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nationalityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Nationality.g4"

	return this
}

// NationalityParser tokens.
const (
	NationalityParserEOF         = antlr.TokenEOF
	NationalityParserNATIONALITY = 1
	NationalityParserOWNKEY      = 2
	NationalityParserSPLITKEY    = 3
	NationalityParserWS          = 4
)

// NationalityParser rules.
const (
	NationalityParserRULE_expression  = 0
	NationalityParserRULE_nationality = 1
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
	p.RuleIndex = NationalityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NationalityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Nationality() INationalityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INationalityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INationalityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NationalityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NationalityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NationalityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NationalityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NationalityParserRULE_expression)

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
		p.Nationality()
	}
	{
		p.SetState(5)
		p.Match(NationalityParserEOF)
	}

	return localctx
}

// INationalityContext is an interface to support dynamic dispatch.
type INationalityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNationalityContext differentiates from other interfaces.
	IsNationalityContext()
}

type NationalityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNationalityContext() *NationalityContext {
	var p = new(NationalityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NationalityParserRULE_nationality
	return p
}

func (*NationalityContext) IsNationalityContext() {}

func NewNationalityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NationalityContext {
	var p = new(NationalityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NationalityParserRULE_nationality

	return p
}

func (s *NationalityContext) GetParser() antlr.Parser { return s.parser }

func (s *NationalityContext) NATIONALITY() antlr.TerminalNode {
	return s.GetToken(NationalityParserNATIONALITY, 0)
}

func (s *NationalityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NationalityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NationalityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NationalityListener); ok {
		listenerT.EnterNationality(s)
	}
}

func (s *NationalityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NationalityListener); ok {
		listenerT.ExitNationality(s)
	}
}

func (p *NationalityParser) Nationality() (localctx INationalityContext) {
	this := p
	_ = this

	localctx = NewNationalityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NationalityParserRULE_nationality)

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
		p.Match(NationalityParserNATIONALITY)
	}

	return localctx
}
