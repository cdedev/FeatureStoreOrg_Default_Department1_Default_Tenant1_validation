// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196098949/PM.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PM

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

type PMParser struct {
	*antlr.BaseParser
}

var pmParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pmParserInit() {
	staticData := &pmParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pm",
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

// PMParserInit initializes any static state used to implement PMParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPMParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PMParserInit() {
	staticData := &pmParserStaticData
	staticData.once.Do(pmParserInit)
}

// NewPMParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPMParser(input antlr.TokenStream) *PMParser {
	PMParserInit()
	this := new(PMParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pmParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PM.g4"

	return this
}

// PMParser tokens.
const (
	PMParserEOF      = antlr.TokenEOF
	PMParserPM       = 1
	PMParserOWNKEY   = 2
	PMParserSPLITKEY = 3
	PMParserWS       = 4
)

// PMParser rules.
const (
	PMParserRULE_expression = 0
	PMParserRULE_pm         = 1
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
	p.RuleIndex = PMParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pm() IPmContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPmContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPmContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PMParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PMParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PMParserRULE_expression)

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
		p.Pm()
	}
	{
		p.SetState(5)
		p.Match(PMParserEOF)
	}

	return localctx
}

// IPmContext is an interface to support dynamic dispatch.
type IPmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmContext differentiates from other interfaces.
	IsPmContext()
}

type PmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmContext() *PmContext {
	var p = new(PmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PMParserRULE_pm
	return p
}

func (*PmContext) IsPmContext() {}

func NewPmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmContext {
	var p = new(PmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PMParserRULE_pm

	return p
}

func (s *PmContext) GetParser() antlr.Parser { return s.parser }

func (s *PmContext) PM() antlr.TerminalNode {
	return s.GetToken(PMParserPM, 0)
}

func (s *PmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMListener); ok {
		listenerT.EnterPm(s)
	}
}

func (s *PmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PMListener); ok {
		listenerT.ExitPm(s)
	}
}

func (p *PMParser) Pm() (localctx IPmContext) {
	this := p
	_ = this

	localctx = NewPmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PMParserRULE_pm)

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
		p.Match(PMParserPM)
	}

	return localctx
}
