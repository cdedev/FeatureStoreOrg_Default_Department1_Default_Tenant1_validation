// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674722475152/Occurrence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occurrence

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

type OccurrenceParser struct {
	*antlr.BaseParser
}

var occurrenceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func occurrenceParserInit() {
	staticData := &occurrenceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OCCURRENCE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "occurrence",
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

// OccurrenceParserInit initializes any static state used to implement OccurrenceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOccurrenceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OccurrenceParserInit() {
	staticData := &occurrenceParserStaticData
	staticData.once.Do(occurrenceParserInit)
}

// NewOccurrenceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOccurrenceParser(input antlr.TokenStream) *OccurrenceParser {
	OccurrenceParserInit()
	this := new(OccurrenceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &occurrenceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Occurrence.g4"

	return this
}

// OccurrenceParser tokens.
const (
	OccurrenceParserEOF        = antlr.TokenEOF
	OccurrenceParserOCCURRENCE = 1
	OccurrenceParserOWNKEY     = 2
	OccurrenceParserSPLITKEY   = 3
	OccurrenceParserWS         = 4
)

// OccurrenceParser rules.
const (
	OccurrenceParserRULE_expression = 0
	OccurrenceParserRULE_occurrence = 1
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
	p.RuleIndex = OccurrenceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OccurrenceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Occurrence() IOccurrenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOccurrenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOccurrenceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OccurrenceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccurrenceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccurrenceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OccurrenceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OccurrenceParserRULE_expression)

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
		p.Occurrence()
	}
	{
		p.SetState(5)
		p.Match(OccurrenceParserEOF)
	}

	return localctx
}

// IOccurrenceContext is an interface to support dynamic dispatch.
type IOccurrenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOccurrenceContext differentiates from other interfaces.
	IsOccurrenceContext()
}

type OccurrenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOccurrenceContext() *OccurrenceContext {
	var p = new(OccurrenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OccurrenceParserRULE_occurrence
	return p
}

func (*OccurrenceContext) IsOccurrenceContext() {}

func NewOccurrenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OccurrenceContext {
	var p = new(OccurrenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OccurrenceParserRULE_occurrence

	return p
}

func (s *OccurrenceContext) GetParser() antlr.Parser { return s.parser }

func (s *OccurrenceContext) OCCURRENCE() antlr.TerminalNode {
	return s.GetToken(OccurrenceParserOCCURRENCE, 0)
}

func (s *OccurrenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OccurrenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OccurrenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccurrenceListener); ok {
		listenerT.EnterOccurrence(s)
	}
}

func (s *OccurrenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccurrenceListener); ok {
		listenerT.ExitOccurrence(s)
	}
}

func (p *OccurrenceParser) Occurrence() (localctx IOccurrenceContext) {
	this := p
	_ = this

	localctx = NewOccurrenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OccurrenceParserRULE_occurrence)

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
		p.Match(OccurrenceParserOCCURRENCE)
	}

	return localctx
}
