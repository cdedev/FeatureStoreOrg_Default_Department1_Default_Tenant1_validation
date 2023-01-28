// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674883414607/TotalLikes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalLikes

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

type TotalLikesParser struct {
	*antlr.BaseParser
}

var totallikesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func totallikesParserInit() {
	staticData := &totallikesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TOTALLIKES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "totallikes",
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

// TotalLikesParserInit initializes any static state used to implement TotalLikesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTotalLikesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TotalLikesParserInit() {
	staticData := &totallikesParserStaticData
	staticData.once.Do(totallikesParserInit)
}

// NewTotalLikesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTotalLikesParser(input antlr.TokenStream) *TotalLikesParser {
	TotalLikesParserInit()
	this := new(TotalLikesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &totallikesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "TotalLikes.g4"

	return this
}

// TotalLikesParser tokens.
const (
	TotalLikesParserEOF        = antlr.TokenEOF
	TotalLikesParserTOTALLIKES = 1
	TotalLikesParserOWNKEY     = 2
	TotalLikesParserSPLITKEY   = 3
	TotalLikesParserWS         = 4
)

// TotalLikesParser rules.
const (
	TotalLikesParserRULE_expression = 0
	TotalLikesParserRULE_totallikes = 1
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
	p.RuleIndex = TotalLikesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TotalLikesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Totallikes() ITotallikesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITotallikesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITotallikesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(TotalLikesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalLikesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalLikesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TotalLikesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TotalLikesParserRULE_expression)

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
		p.Totallikes()
	}
	{
		p.SetState(5)
		p.Match(TotalLikesParserEOF)
	}

	return localctx
}

// ITotallikesContext is an interface to support dynamic dispatch.
type ITotallikesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTotallikesContext differentiates from other interfaces.
	IsTotallikesContext()
}

type TotallikesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTotallikesContext() *TotallikesContext {
	var p = new(TotallikesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TotalLikesParserRULE_totallikes
	return p
}

func (*TotallikesContext) IsTotallikesContext() {}

func NewTotallikesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TotallikesContext {
	var p = new(TotallikesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TotalLikesParserRULE_totallikes

	return p
}

func (s *TotallikesContext) GetParser() antlr.Parser { return s.parser }

func (s *TotallikesContext) TOTALLIKES() antlr.TerminalNode {
	return s.GetToken(TotalLikesParserTOTALLIKES, 0)
}

func (s *TotallikesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TotallikesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TotallikesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalLikesListener); ok {
		listenerT.EnterTotallikes(s)
	}
}

func (s *TotallikesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TotalLikesListener); ok {
		listenerT.ExitTotallikes(s)
	}
}

func (p *TotalLikesParser) Totallikes() (localctx ITotallikesContext) {
	this := p
	_ = this

	localctx = NewTotallikesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TotalLikesParserRULE_totallikes)

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
		p.Match(TotalLikesParserTOTALLIKES)
	}

	return localctx
}
