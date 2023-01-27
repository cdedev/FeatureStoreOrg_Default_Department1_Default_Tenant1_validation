// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674792977494/Community.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Community

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

type CommunityParser struct {
	*antlr.BaseParser
}

var communityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func communityParserInit() {
	staticData := &communityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COMMUNITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "community",
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

// CommunityParserInit initializes any static state used to implement CommunityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCommunityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CommunityParserInit() {
	staticData := &communityParserStaticData
	staticData.once.Do(communityParserInit)
}

// NewCommunityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCommunityParser(input antlr.TokenStream) *CommunityParser {
	CommunityParserInit()
	this := new(CommunityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &communityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Community.g4"

	return this
}

// CommunityParser tokens.
const (
	CommunityParserEOF       = antlr.TokenEOF
	CommunityParserCOMMUNITY = 1
	CommunityParserOWNKEY    = 2
	CommunityParserSPLITKEY  = 3
	CommunityParserWS        = 4
)

// CommunityParser rules.
const (
	CommunityParserRULE_expression = 0
	CommunityParserRULE_community  = 1
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
	p.RuleIndex = CommunityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommunityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Community() ICommunityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommunityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommunityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CommunityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommunityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommunityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CommunityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CommunityParserRULE_expression)

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
		p.Community()
	}
	{
		p.SetState(5)
		p.Match(CommunityParserEOF)
	}

	return localctx
}

// ICommunityContext is an interface to support dynamic dispatch.
type ICommunityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommunityContext differentiates from other interfaces.
	IsCommunityContext()
}

type CommunityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommunityContext() *CommunityContext {
	var p = new(CommunityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CommunityParserRULE_community
	return p
}

func (*CommunityContext) IsCommunityContext() {}

func NewCommunityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommunityContext {
	var p = new(CommunityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CommunityParserRULE_community

	return p
}

func (s *CommunityContext) GetParser() antlr.Parser { return s.parser }

func (s *CommunityContext) COMMUNITY() antlr.TerminalNode {
	return s.GetToken(CommunityParserCOMMUNITY, 0)
}

func (s *CommunityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommunityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommunityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommunityListener); ok {
		listenerT.EnterCommunity(s)
	}
}

func (s *CommunityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CommunityListener); ok {
		listenerT.ExitCommunity(s)
	}
}

func (p *CommunityParser) Community() (localctx ICommunityContext) {
	this := p
	_ = this

	localctx = NewCommunityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CommunityParserRULE_community)

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
		p.Match(CommunityParserCOMMUNITY)
	}

	return localctx
}
