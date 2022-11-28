// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669651221577/Title.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Title

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

type TitleParser struct {
	*antlr.BaseParser
}

var titleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func titleParserInit() {
	staticData := &titleParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TITLE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "title",
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

// TitleParserInit initializes any static state used to implement TitleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTitleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TitleParserInit() {
	staticData := &titleParserStaticData
	staticData.once.Do(titleParserInit)
}

// NewTitleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTitleParser(input antlr.TokenStream) *TitleParser {
	TitleParserInit()
	this := new(TitleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &titleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Title.g4"

	return this
}

// TitleParser tokens.
const (
	TitleParserEOF      = antlr.TokenEOF
	TitleParserTITLE    = 1
	TitleParserOWNKEY   = 2
	TitleParserSPLITKEY = 3
	TitleParserWS       = 4
)

// TitleParser rules.
const (
	TitleParserRULE_expression = 0
	TitleParserRULE_title      = 1
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
	p.RuleIndex = TitleParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TitleParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Title() ITitleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITitleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITitleContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TitleParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TitleListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TitleListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TitleParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TitleParserRULE_expression)

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
		p.Title()
	}
	{
		p.SetState(5)
		p.Match(TitleParserEOF)
	}

	return localctx
}

// ITitleContext is an interface to support dynamic dispatch.
type ITitleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTitleContext differentiates from other interfaces.
	IsTitleContext()
}

type TitleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTitleContext() *TitleContext {
	var p = new(TitleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TitleParserRULE_title
	return p
}

func (*TitleContext) IsTitleContext() {}

func NewTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TitleContext {
	var p = new(TitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TitleParserRULE_title

	return p
}

func (s *TitleContext) GetParser() antlr.Parser { return s.parser }

func (s *TitleContext) TITLE() antlr.TerminalNode {
	return s.GetToken(TitleParserTITLE, 0)
}

func (s *TitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TitleListener); ok {
		listenerT.EnterTitle(s)
	}
}

func (s *TitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TitleListener); ok {
		listenerT.ExitTitle(s)
	}
}

func (p *TitleParser) Title() (localctx ITitleContext) {
	this := p
	_ = this

	localctx = NewTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TitleParserRULE_title)

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
		p.Match(TitleParserTITLE)
	}

	return localctx
}
