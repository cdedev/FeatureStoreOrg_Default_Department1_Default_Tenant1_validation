// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789884692/Author.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Author

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

type AuthorParser struct {
	*antlr.BaseParser
}

var authorParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func authorParserInit() {
	staticData := &authorParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "AUTHOR", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "author",
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

// AuthorParserInit initializes any static state used to implement AuthorParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewAuthorParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func AuthorParserInit() {
	staticData := &authorParserStaticData
	staticData.once.Do(authorParserInit)
}

// NewAuthorParser produces a new parser instance for the optional input antlr.TokenStream.
func NewAuthorParser(input antlr.TokenStream) *AuthorParser {
	AuthorParserInit()
	this := new(AuthorParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &authorParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Author.g4"

	return this
}

// AuthorParser tokens.
const (
	AuthorParserEOF      = antlr.TokenEOF
	AuthorParserAUTHOR   = 1
	AuthorParserOWNKEY   = 2
	AuthorParserSPLITKEY = 3
	AuthorParserWS       = 4
)

// AuthorParser rules.
const (
	AuthorParserRULE_expression = 0
	AuthorParserRULE_author     = 1
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
	p.RuleIndex = AuthorParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuthorParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Author() IAuthorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAuthorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAuthorContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(AuthorParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuthorListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuthorListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *AuthorParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AuthorParserRULE_expression)

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
		p.Author()
	}
	{
		p.SetState(5)
		p.Match(AuthorParserEOF)
	}

	return localctx
}

// IAuthorContext is an interface to support dynamic dispatch.
type IAuthorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAuthorContext differentiates from other interfaces.
	IsAuthorContext()
}

type AuthorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAuthorContext() *AuthorContext {
	var p = new(AuthorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AuthorParserRULE_author
	return p
}

func (*AuthorContext) IsAuthorContext() {}

func NewAuthorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AuthorContext {
	var p = new(AuthorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AuthorParserRULE_author

	return p
}

func (s *AuthorContext) GetParser() antlr.Parser { return s.parser }

func (s *AuthorContext) AUTHOR() antlr.TerminalNode {
	return s.GetToken(AuthorParserAUTHOR, 0)
}

func (s *AuthorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AuthorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuthorListener); ok {
		listenerT.EnterAuthor(s)
	}
}

func (s *AuthorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AuthorListener); ok {
		listenerT.ExitAuthor(s)
	}
}

func (p *AuthorParser) Author() (localctx IAuthorContext) {
	this := p
	_ = this

	localctx = NewAuthorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AuthorParserRULE_author)

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
		p.Match(AuthorParserAUTHOR)
	}

	return localctx
}
