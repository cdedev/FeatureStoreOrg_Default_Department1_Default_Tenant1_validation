// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284871047/SuggestiveThemes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SuggestiveThemes

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

type SuggestiveThemesParser struct {
	*antlr.BaseParser
}

var suggestivethemesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func suggestivethemesParserInit() {
	staticData := &suggestivethemesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SUGGESTIVETHEMES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "suggestivethemes",
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

// SuggestiveThemesParserInit initializes any static state used to implement SuggestiveThemesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSuggestiveThemesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SuggestiveThemesParserInit() {
	staticData := &suggestivethemesParserStaticData
	staticData.once.Do(suggestivethemesParserInit)
}

// NewSuggestiveThemesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSuggestiveThemesParser(input antlr.TokenStream) *SuggestiveThemesParser {
	SuggestiveThemesParserInit()
	this := new(SuggestiveThemesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &suggestivethemesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SuggestiveThemes.g4"

	return this
}

// SuggestiveThemesParser tokens.
const (
	SuggestiveThemesParserEOF              = antlr.TokenEOF
	SuggestiveThemesParserSUGGESTIVETHEMES = 1
	SuggestiveThemesParserOWNKEY           = 2
	SuggestiveThemesParserSPLITKEY         = 3
	SuggestiveThemesParserWS               = 4
)

// SuggestiveThemesParser rules.
const (
	SuggestiveThemesParserRULE_expression       = 0
	SuggestiveThemesParserRULE_suggestivethemes = 1
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
	p.RuleIndex = SuggestiveThemesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuggestiveThemesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Suggestivethemes() ISuggestivethemesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISuggestivethemesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISuggestivethemesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SuggestiveThemesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuggestiveThemesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuggestiveThemesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SuggestiveThemesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SuggestiveThemesParserRULE_expression)

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
		p.Suggestivethemes()
	}
	{
		p.SetState(5)
		p.Match(SuggestiveThemesParserEOF)
	}

	return localctx
}

// ISuggestivethemesContext is an interface to support dynamic dispatch.
type ISuggestivethemesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSuggestivethemesContext differentiates from other interfaces.
	IsSuggestivethemesContext()
}

type SuggestivethemesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySuggestivethemesContext() *SuggestivethemesContext {
	var p = new(SuggestivethemesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SuggestiveThemesParserRULE_suggestivethemes
	return p
}

func (*SuggestivethemesContext) IsSuggestivethemesContext() {}

func NewSuggestivethemesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SuggestivethemesContext {
	var p = new(SuggestivethemesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SuggestiveThemesParserRULE_suggestivethemes

	return p
}

func (s *SuggestivethemesContext) GetParser() antlr.Parser { return s.parser }

func (s *SuggestivethemesContext) SUGGESTIVETHEMES() antlr.TerminalNode {
	return s.GetToken(SuggestiveThemesParserSUGGESTIVETHEMES, 0)
}

func (s *SuggestivethemesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuggestivethemesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SuggestivethemesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuggestiveThemesListener); ok {
		listenerT.EnterSuggestivethemes(s)
	}
}

func (s *SuggestivethemesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SuggestiveThemesListener); ok {
		listenerT.ExitSuggestivethemes(s)
	}
}

func (p *SuggestiveThemesParser) Suggestivethemes() (localctx ISuggestivethemesContext) {
	this := p
	_ = this

	localctx = NewSuggestivethemesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SuggestiveThemesParserRULE_suggestivethemes)

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
		p.Match(SuggestiveThemesParserSUGGESTIVETHEMES)
	}

	return localctx
}
