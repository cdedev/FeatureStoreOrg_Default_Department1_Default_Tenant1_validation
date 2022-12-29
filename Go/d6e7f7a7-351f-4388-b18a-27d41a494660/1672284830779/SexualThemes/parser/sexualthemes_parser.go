// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284830779/SexualThemes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SexualThemes

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

type SexualThemesParser struct {
	*antlr.BaseParser
}

var sexualthemesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sexualthemesParserInit() {
	staticData := &sexualthemesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SEXUALTHEMES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sexualthemes",
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

// SexualThemesParserInit initializes any static state used to implement SexualThemesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSexualThemesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SexualThemesParserInit() {
	staticData := &sexualthemesParserStaticData
	staticData.once.Do(sexualthemesParserInit)
}

// NewSexualThemesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSexualThemesParser(input antlr.TokenStream) *SexualThemesParser {
	SexualThemesParserInit()
	this := new(SexualThemesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sexualthemesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SexualThemes.g4"

	return this
}

// SexualThemesParser tokens.
const (
	SexualThemesParserEOF          = antlr.TokenEOF
	SexualThemesParserSEXUALTHEMES = 1
	SexualThemesParserOWNKEY       = 2
	SexualThemesParserSPLITKEY     = 3
	SexualThemesParserWS           = 4
)

// SexualThemesParser rules.
const (
	SexualThemesParserRULE_expression   = 0
	SexualThemesParserRULE_sexualthemes = 1
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
	p.RuleIndex = SexualThemesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SexualThemesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sexualthemes() ISexualthemesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISexualthemesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISexualthemesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SexualThemesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexualThemesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexualThemesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SexualThemesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SexualThemesParserRULE_expression)

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
		p.Sexualthemes()
	}
	{
		p.SetState(5)
		p.Match(SexualThemesParserEOF)
	}

	return localctx
}

// ISexualthemesContext is an interface to support dynamic dispatch.
type ISexualthemesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSexualthemesContext differentiates from other interfaces.
	IsSexualthemesContext()
}

type SexualthemesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexualthemesContext() *SexualthemesContext {
	var p = new(SexualthemesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SexualThemesParserRULE_sexualthemes
	return p
}

func (*SexualthemesContext) IsSexualthemesContext() {}

func NewSexualthemesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexualthemesContext {
	var p = new(SexualthemesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SexualThemesParserRULE_sexualthemes

	return p
}

func (s *SexualthemesContext) GetParser() antlr.Parser { return s.parser }

func (s *SexualthemesContext) SEXUALTHEMES() antlr.TerminalNode {
	return s.GetToken(SexualThemesParserSEXUALTHEMES, 0)
}

func (s *SexualthemesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexualthemesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SexualthemesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexualThemesListener); ok {
		listenerT.EnterSexualthemes(s)
	}
}

func (s *SexualthemesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SexualThemesListener); ok {
		listenerT.ExitSexualthemes(s)
	}
}

func (p *SexualThemesParser) Sexualthemes() (localctx ISexualthemesContext) {
	this := p
	_ = this

	localctx = NewSexualthemesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SexualThemesParserRULE_sexualthemes)

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
		p.Match(SexualThemesParserSEXUALTHEMES)
	}

	return localctx
}
