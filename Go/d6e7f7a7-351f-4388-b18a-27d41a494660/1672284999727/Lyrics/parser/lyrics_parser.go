// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284999727/Lyrics.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lyrics

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

type LyricsParser struct {
	*antlr.BaseParser
}

var lyricsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lyricsParserInit() {
	staticData := &lyricsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LYRICS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "lyrics",
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

// LyricsParserInit initializes any static state used to implement LyricsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLyricsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LyricsParserInit() {
	staticData := &lyricsParserStaticData
	staticData.once.Do(lyricsParserInit)
}

// NewLyricsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLyricsParser(input antlr.TokenStream) *LyricsParser {
	LyricsParserInit()
	this := new(LyricsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lyricsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Lyrics.g4"

	return this
}

// LyricsParser tokens.
const (
	LyricsParserEOF      = antlr.TokenEOF
	LyricsParserLYRICS   = 1
	LyricsParserOWNKEY   = 2
	LyricsParserSPLITKEY = 3
	LyricsParserWS       = 4
)

// LyricsParser rules.
const (
	LyricsParserRULE_expression = 0
	LyricsParserRULE_lyrics     = 1
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
	p.RuleIndex = LyricsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LyricsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Lyrics() ILyricsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILyricsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILyricsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LyricsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LyricsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LyricsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LyricsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LyricsParserRULE_expression)

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
		p.Lyrics()
	}
	{
		p.SetState(5)
		p.Match(LyricsParserEOF)
	}

	return localctx
}

// ILyricsContext is an interface to support dynamic dispatch.
type ILyricsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLyricsContext differentiates from other interfaces.
	IsLyricsContext()
}

type LyricsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLyricsContext() *LyricsContext {
	var p = new(LyricsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LyricsParserRULE_lyrics
	return p
}

func (*LyricsContext) IsLyricsContext() {}

func NewLyricsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LyricsContext {
	var p = new(LyricsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LyricsParserRULE_lyrics

	return p
}

func (s *LyricsContext) GetParser() antlr.Parser { return s.parser }

func (s *LyricsContext) LYRICS() antlr.TerminalNode {
	return s.GetToken(LyricsParserLYRICS, 0)
}

func (s *LyricsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LyricsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LyricsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LyricsListener); ok {
		listenerT.EnterLyrics(s)
	}
}

func (s *LyricsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LyricsListener); ok {
		listenerT.ExitLyrics(s)
	}
}

func (p *LyricsParser) Lyrics() (localctx ILyricsContext) {
	this := p
	_ = this

	localctx = NewLyricsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LyricsParserRULE_lyrics)

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
		p.Match(LyricsParserLYRICS)
	}

	return localctx
}
