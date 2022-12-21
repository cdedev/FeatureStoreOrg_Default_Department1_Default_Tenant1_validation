// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601452315/SourceID.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SourceID

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

type SourceIDParser struct {
	*antlr.BaseParser
}

var sourceidParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sourceidParserInit() {
	staticData := &sourceidParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SOURCEID", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sourceid",
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

// SourceIDParserInit initializes any static state used to implement SourceIDParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSourceIDParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SourceIDParserInit() {
	staticData := &sourceidParserStaticData
	staticData.once.Do(sourceidParserInit)
}

// NewSourceIDParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSourceIDParser(input antlr.TokenStream) *SourceIDParser {
	SourceIDParserInit()
	this := new(SourceIDParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sourceidParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SourceID.g4"

	return this
}

// SourceIDParser tokens.
const (
	SourceIDParserEOF      = antlr.TokenEOF
	SourceIDParserSOURCEID = 1
	SourceIDParserOWNKEY   = 2
	SourceIDParserSPLITKEY = 3
	SourceIDParserWS       = 4
)

// SourceIDParser rules.
const (
	SourceIDParserRULE_expression = 0
	SourceIDParserRULE_sourceid   = 1
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
	p.RuleIndex = SourceIDParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceIDParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sourceid() ISourceidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceidContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SourceIDParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceIDListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceIDListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SourceIDParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SourceIDParserRULE_expression)

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
		p.Sourceid()
	}
	{
		p.SetState(5)
		p.Match(SourceIDParserEOF)
	}

	return localctx
}

// ISourceidContext is an interface to support dynamic dispatch.
type ISourceidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceidContext differentiates from other interfaces.
	IsSourceidContext()
}

type SourceidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceidContext() *SourceidContext {
	var p = new(SourceidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SourceIDParserRULE_sourceid
	return p
}

func (*SourceidContext) IsSourceidContext() {}

func NewSourceidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceidContext {
	var p = new(SourceidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SourceIDParserRULE_sourceid

	return p
}

func (s *SourceidContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceidContext) SOURCEID() antlr.TerminalNode {
	return s.GetToken(SourceIDParserSOURCEID, 0)
}

func (s *SourceidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceIDListener); ok {
		listenerT.EnterSourceid(s)
	}
}

func (s *SourceidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SourceIDListener); ok {
		listenerT.ExitSourceid(s)
	}
}

func (p *SourceIDParser) Sourceid() (localctx ISourceidContext) {
	this := p
	_ = this

	localctx = NewSourceidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SourceIDParserRULE_sourceid)

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
		p.Match(SourceIDParserSOURCEID)
	}

	return localctx
}
