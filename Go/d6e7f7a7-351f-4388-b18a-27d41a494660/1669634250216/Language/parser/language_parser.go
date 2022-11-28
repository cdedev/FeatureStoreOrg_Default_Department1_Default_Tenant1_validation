// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669634250216/Language.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Language

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

type LanguageParser struct {
	*antlr.BaseParser
}

var languageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func languageParserInit() {
	staticData := &languageParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LANGUAGE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "language",
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

// LanguageParserInit initializes any static state used to implement LanguageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLanguageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LanguageParserInit() {
	staticData := &languageParserStaticData
	staticData.once.Do(languageParserInit)
}

// NewLanguageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLanguageParser(input antlr.TokenStream) *LanguageParser {
	LanguageParserInit()
	this := new(LanguageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &languageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Language.g4"

	return this
}

// LanguageParser tokens.
const (
	LanguageParserEOF      = antlr.TokenEOF
	LanguageParserLANGUAGE = 1
	LanguageParserOWNKEY   = 2
	LanguageParserSPLITKEY = 3
	LanguageParserWS       = 4
)

// LanguageParser rules.
const (
	LanguageParserRULE_expression = 0
	LanguageParserRULE_language   = 1
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
	p.RuleIndex = LanguageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Language() ILanguageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILanguageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILanguageContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LanguageParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LanguageParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LanguageParserRULE_expression)

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
		p.Language()
	}
	{
		p.SetState(5)
		p.Match(LanguageParserEOF)
	}

	return localctx
}

// ILanguageContext is an interface to support dynamic dispatch.
type ILanguageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLanguageContext differentiates from other interfaces.
	IsLanguageContext()
}

type LanguageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLanguageContext() *LanguageContext {
	var p = new(LanguageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LanguageParserRULE_language
	return p
}

func (*LanguageContext) IsLanguageContext() {}

func NewLanguageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LanguageContext {
	var p = new(LanguageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LanguageParserRULE_language

	return p
}

func (s *LanguageContext) GetParser() antlr.Parser { return s.parser }

func (s *LanguageContext) LANGUAGE() antlr.TerminalNode {
	return s.GetToken(LanguageParserLANGUAGE, 0)
}

func (s *LanguageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LanguageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LanguageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.EnterLanguage(s)
	}
}

func (s *LanguageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LanguageListener); ok {
		listenerT.ExitLanguage(s)
	}
}

func (p *LanguageParser) Language() (localctx ILanguageContext) {
	this := p
	_ = this

	localctx = NewLanguageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LanguageParserRULE_language)

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
		p.Match(LanguageParserLANGUAGE)
	}

	return localctx
}
