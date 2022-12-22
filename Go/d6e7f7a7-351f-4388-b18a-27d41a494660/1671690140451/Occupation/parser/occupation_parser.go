// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690140451/Occupation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occupation

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

type OccupationParser struct {
	*antlr.BaseParser
}

var occupationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func occupationParserInit() {
	staticData := &occupationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OCCUPATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "occupation",
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

// OccupationParserInit initializes any static state used to implement OccupationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOccupationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OccupationParserInit() {
	staticData := &occupationParserStaticData
	staticData.once.Do(occupationParserInit)
}

// NewOccupationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOccupationParser(input antlr.TokenStream) *OccupationParser {
	OccupationParserInit()
	this := new(OccupationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &occupationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Occupation.g4"

	return this
}

// OccupationParser tokens.
const (
	OccupationParserEOF        = antlr.TokenEOF
	OccupationParserOCCUPATION = 1
	OccupationParserOWNKEY     = 2
	OccupationParserSPLITKEY   = 3
	OccupationParserWS         = 4
)

// OccupationParser rules.
const (
	OccupationParserRULE_expression = 0
	OccupationParserRULE_occupation = 1
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
	p.RuleIndex = OccupationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OccupationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Occupation() IOccupationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOccupationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOccupationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OccupationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OccupationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OccupationParserRULE_expression)

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
		p.Occupation()
	}
	{
		p.SetState(5)
		p.Match(OccupationParserEOF)
	}

	return localctx
}

// IOccupationContext is an interface to support dynamic dispatch.
type IOccupationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOccupationContext differentiates from other interfaces.
	IsOccupationContext()
}

type OccupationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOccupationContext() *OccupationContext {
	var p = new(OccupationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OccupationParserRULE_occupation
	return p
}

func (*OccupationContext) IsOccupationContext() {}

func NewOccupationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OccupationContext {
	var p = new(OccupationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OccupationParserRULE_occupation

	return p
}

func (s *OccupationContext) GetParser() antlr.Parser { return s.parser }

func (s *OccupationContext) OCCUPATION() antlr.TerminalNode {
	return s.GetToken(OccupationParserOCCUPATION, 0)
}

func (s *OccupationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OccupationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OccupationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationListener); ok {
		listenerT.EnterOccupation(s)
	}
}

func (s *OccupationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationListener); ok {
		listenerT.ExitOccupation(s)
	}
}

func (p *OccupationParser) Occupation() (localctx IOccupationContext) {
	this := p
	_ = this

	localctx = NewOccupationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OccupationParserRULE_occupation)

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
		p.Match(OccupationParserOCCUPATION)
	}

	return localctx
}
