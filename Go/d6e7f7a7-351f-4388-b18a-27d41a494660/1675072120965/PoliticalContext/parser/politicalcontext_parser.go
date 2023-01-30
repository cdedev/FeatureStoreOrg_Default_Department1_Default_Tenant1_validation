// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072120965/PoliticalContext.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PoliticalContext

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

type PoliticalContextParser struct {
	*antlr.BaseParser
}

var politicalcontextParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func politicalcontextParserInit() {
	staticData := &politicalcontextParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "POLITICALCONTEXT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "politicalcontext",
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

// PoliticalContextParserInit initializes any static state used to implement PoliticalContextParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPoliticalContextParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PoliticalContextParserInit() {
	staticData := &politicalcontextParserStaticData
	staticData.once.Do(politicalcontextParserInit)
}

// NewPoliticalContextParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPoliticalContextParser(input antlr.TokenStream) *PoliticalContextParser {
	PoliticalContextParserInit()
	this := new(PoliticalContextParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &politicalcontextParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PoliticalContext.g4"

	return this
}

// PoliticalContextParser tokens.
const (
	PoliticalContextParserEOF              = antlr.TokenEOF
	PoliticalContextParserPOLITICALCONTEXT = 1
	PoliticalContextParserOWNKEY           = 2
	PoliticalContextParserSPLITKEY         = 3
	PoliticalContextParserWS               = 4
)

// PoliticalContextParser rules.
const (
	PoliticalContextParserRULE_expression       = 0
	PoliticalContextParserRULE_politicalcontext = 1
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
	p.RuleIndex = PoliticalContextParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PoliticalContextParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Politicalcontext() IPoliticalcontextContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPoliticalcontextContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPoliticalcontextContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PoliticalContextParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoliticalContextListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoliticalContextListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PoliticalContextParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PoliticalContextParserRULE_expression)

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
		p.Politicalcontext()
	}
	{
		p.SetState(5)
		p.Match(PoliticalContextParserEOF)
	}

	return localctx
}

// IPoliticalcontextContext is an interface to support dynamic dispatch.
type IPoliticalcontextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPoliticalcontextContext differentiates from other interfaces.
	IsPoliticalcontextContext()
}

type PoliticalcontextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPoliticalcontextContext() *PoliticalcontextContext {
	var p = new(PoliticalcontextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PoliticalContextParserRULE_politicalcontext
	return p
}

func (*PoliticalcontextContext) IsPoliticalcontextContext() {}

func NewPoliticalcontextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PoliticalcontextContext {
	var p = new(PoliticalcontextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PoliticalContextParserRULE_politicalcontext

	return p
}

func (s *PoliticalcontextContext) GetParser() antlr.Parser { return s.parser }

func (s *PoliticalcontextContext) POLITICALCONTEXT() antlr.TerminalNode {
	return s.GetToken(PoliticalContextParserPOLITICALCONTEXT, 0)
}

func (s *PoliticalcontextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PoliticalcontextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PoliticalcontextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoliticalContextListener); ok {
		listenerT.EnterPoliticalcontext(s)
	}
}

func (s *PoliticalcontextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoliticalContextListener); ok {
		listenerT.ExitPoliticalcontext(s)
	}
}

func (p *PoliticalContextParser) Politicalcontext() (localctx IPoliticalcontextContext) {
	this := p
	_ = this

	localctx = NewPoliticalcontextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PoliticalContextParserRULE_politicalcontext)

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
		p.Match(PoliticalContextParserPOLITICALCONTEXT)
	}

	return localctx
}
