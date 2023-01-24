// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674534485453/Sex.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sex

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

type SexParser struct {
	*antlr.BaseParser
}

var sexParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sexParserInit() {
	staticData := &sexParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SEX", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sex",
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

// SexParserInit initializes any static state used to implement SexParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSexParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SexParserInit() {
	staticData := &sexParserStaticData
	staticData.once.Do(sexParserInit)
}

// NewSexParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSexParser(input antlr.TokenStream) *SexParser {
	SexParserInit()
	this := new(SexParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sexParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Sex.g4"

	return this
}

// SexParser tokens.
const (
	SexParserEOF      = antlr.TokenEOF
	SexParserSEX      = 1
	SexParserOWNKEY   = 2
	SexParserSPLITKEY = 3
	SexParserWS       = 4
)

// SexParser rules.
const (
	SexParserRULE_expression = 0
	SexParserRULE_sex        = 1
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
	p.RuleIndex = SexParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SexParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sex() ISexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISexContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SexParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SexParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SexParserRULE_expression)

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
		p.Sex()
	}
	{
		p.SetState(5)
		p.Match(SexParserEOF)
	}

	return localctx
}

// ISexContext is an interface to support dynamic dispatch.
type ISexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSexContext differentiates from other interfaces.
	IsSexContext()
}

type SexContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexContext() *SexContext {
	var p = new(SexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SexParserRULE_sex
	return p
}

func (*SexContext) IsSexContext() {}

func NewSexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexContext {
	var p = new(SexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SexParserRULE_sex

	return p
}

func (s *SexContext) GetParser() antlr.Parser { return s.parser }

func (s *SexContext) SEX() antlr.TerminalNode {
	return s.GetToken(SexParserSEX, 0)
}

func (s *SexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexListener); ok {
		listenerT.EnterSex(s)
	}
}

func (s *SexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexListener); ok {
		listenerT.ExitSex(s)
	}
}

func (p *SexParser) Sex() (localctx ISexContext) {
	this := p
	_ = this

	localctx = NewSexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SexParserRULE_sex)

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
		p.Match(SexParserSEX)
	}

	return localctx
}
