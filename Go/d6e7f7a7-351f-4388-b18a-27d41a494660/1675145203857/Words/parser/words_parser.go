// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145203857/Words.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Words

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

type WordsParser struct {
	*antlr.BaseParser
}

var wordsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func wordsParserInit() {
	staticData := &wordsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "WORDS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "words",
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

// WordsParserInit initializes any static state used to implement WordsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWordsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WordsParserInit() {
	staticData := &wordsParserStaticData
	staticData.once.Do(wordsParserInit)
}

// NewWordsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWordsParser(input antlr.TokenStream) *WordsParser {
	WordsParserInit()
	this := new(WordsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &wordsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Words.g4"

	return this
}

// WordsParser tokens.
const (
	WordsParserEOF      = antlr.TokenEOF
	WordsParserWORDS    = 1
	WordsParserOWNKEY   = 2
	WordsParserSPLITKEY = 3
	WordsParserWS       = 4
)

// WordsParser rules.
const (
	WordsParserRULE_expression = 0
	WordsParserRULE_words      = 1
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
	p.RuleIndex = WordsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WordsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Words() IWordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(WordsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WordsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WordsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *WordsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WordsParserRULE_expression)

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
		p.Words()
	}
	{
		p.SetState(5)
		p.Match(WordsParserEOF)
	}

	return localctx
}

// IWordsContext is an interface to support dynamic dispatch.
type IWordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWordsContext differentiates from other interfaces.
	IsWordsContext()
}

type WordsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordsContext() *WordsContext {
	var p = new(WordsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WordsParserRULE_words
	return p
}

func (*WordsContext) IsWordsContext() {}

func NewWordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordsContext {
	var p = new(WordsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WordsParserRULE_words

	return p
}

func (s *WordsContext) GetParser() antlr.Parser { return s.parser }

func (s *WordsContext) WORDS() antlr.TerminalNode {
	return s.GetToken(WordsParserWORDS, 0)
}

func (s *WordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WordsListener); ok {
		listenerT.EnterWords(s)
	}
}

func (s *WordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WordsListener); ok {
		listenerT.ExitWords(s)
	}
}

func (p *WordsParser) Words() (localctx IWordsContext) {
	this := p
	_ = this

	localctx = NewWordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WordsParserRULE_words)

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
		p.Match(WordsParserWORDS)
	}

	return localctx
}
