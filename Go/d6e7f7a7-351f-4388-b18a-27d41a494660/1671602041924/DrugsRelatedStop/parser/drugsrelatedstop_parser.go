// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671602041924/DrugsRelatedStop.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DrugsRelatedStop

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

type DrugsRelatedStopParser struct {
	*antlr.BaseParser
}

var drugsrelatedstopParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func drugsrelatedstopParserInit() {
	staticData := &drugsrelatedstopParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DRUGSRELATEDSTOP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "drugsrelatedstop",
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

// DrugsRelatedStopParserInit initializes any static state used to implement DrugsRelatedStopParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDrugsRelatedStopParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DrugsRelatedStopParserInit() {
	staticData := &drugsrelatedstopParserStaticData
	staticData.once.Do(drugsrelatedstopParserInit)
}

// NewDrugsRelatedStopParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDrugsRelatedStopParser(input antlr.TokenStream) *DrugsRelatedStopParser {
	DrugsRelatedStopParserInit()
	this := new(DrugsRelatedStopParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &drugsrelatedstopParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DrugsRelatedStop.g4"

	return this
}

// DrugsRelatedStopParser tokens.
const (
	DrugsRelatedStopParserEOF              = antlr.TokenEOF
	DrugsRelatedStopParserDRUGSRELATEDSTOP = 1
	DrugsRelatedStopParserOWNKEY           = 2
	DrugsRelatedStopParserSPLITKEY         = 3
	DrugsRelatedStopParserWS               = 4
)

// DrugsRelatedStopParser rules.
const (
	DrugsRelatedStopParserRULE_expression       = 0
	DrugsRelatedStopParserRULE_drugsrelatedstop = 1
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
	p.RuleIndex = DrugsRelatedStopParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrugsRelatedStopParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Drugsrelatedstop() IDrugsrelatedstopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDrugsrelatedstopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDrugsrelatedstopContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DrugsRelatedStopParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugsRelatedStopListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugsRelatedStopListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DrugsRelatedStopParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DrugsRelatedStopParserRULE_expression)

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
		p.Drugsrelatedstop()
	}
	{
		p.SetState(5)
		p.Match(DrugsRelatedStopParserEOF)
	}

	return localctx
}

// IDrugsrelatedstopContext is an interface to support dynamic dispatch.
type IDrugsrelatedstopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDrugsrelatedstopContext differentiates from other interfaces.
	IsDrugsrelatedstopContext()
}

type DrugsrelatedstopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrugsrelatedstopContext() *DrugsrelatedstopContext {
	var p = new(DrugsrelatedstopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DrugsRelatedStopParserRULE_drugsrelatedstop
	return p
}

func (*DrugsrelatedstopContext) IsDrugsrelatedstopContext() {}

func NewDrugsrelatedstopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DrugsrelatedstopContext {
	var p = new(DrugsrelatedstopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrugsRelatedStopParserRULE_drugsrelatedstop

	return p
}

func (s *DrugsrelatedstopContext) GetParser() antlr.Parser { return s.parser }

func (s *DrugsrelatedstopContext) DRUGSRELATEDSTOP() antlr.TerminalNode {
	return s.GetToken(DrugsRelatedStopParserDRUGSRELATEDSTOP, 0)
}

func (s *DrugsrelatedstopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DrugsrelatedstopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DrugsrelatedstopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugsRelatedStopListener); ok {
		listenerT.EnterDrugsrelatedstop(s)
	}
}

func (s *DrugsrelatedstopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugsRelatedStopListener); ok {
		listenerT.ExitDrugsrelatedstop(s)
	}
}

func (p *DrugsRelatedStopParser) Drugsrelatedstop() (localctx IDrugsrelatedstopContext) {
	this := p
	_ = this

	localctx = NewDrugsrelatedstopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DrugsRelatedStopParserRULE_drugsrelatedstop)

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
		p.Match(DrugsRelatedStopParserDRUGSRELATEDSTOP)
	}

	return localctx
}
