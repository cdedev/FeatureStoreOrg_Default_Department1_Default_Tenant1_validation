// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078619149/OutdoorParticulateMatter.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OutdoorParticulateMatter

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

type OutdoorParticulateMatterParser struct {
	*antlr.BaseParser
}

var outdoorparticulatematterParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func outdoorparticulatematterParserInit() {
	staticData := &outdoorparticulatematterParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OUTDOORPARTICULATEMATTER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "outdoorparticulatematter",
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

// OutdoorParticulateMatterParserInit initializes any static state used to implement OutdoorParticulateMatterParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOutdoorParticulateMatterParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OutdoorParticulateMatterParserInit() {
	staticData := &outdoorparticulatematterParserStaticData
	staticData.once.Do(outdoorparticulatematterParserInit)
}

// NewOutdoorParticulateMatterParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOutdoorParticulateMatterParser(input antlr.TokenStream) *OutdoorParticulateMatterParser {
	OutdoorParticulateMatterParserInit()
	this := new(OutdoorParticulateMatterParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &outdoorparticulatematterParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "OutdoorParticulateMatter.g4"

	return this
}

// OutdoorParticulateMatterParser tokens.
const (
	OutdoorParticulateMatterParserEOF                      = antlr.TokenEOF
	OutdoorParticulateMatterParserOUTDOORPARTICULATEMATTER = 1
	OutdoorParticulateMatterParserOWNKEY                   = 2
	OutdoorParticulateMatterParserSPLITKEY                 = 3
	OutdoorParticulateMatterParserWS                       = 4
)

// OutdoorParticulateMatterParser rules.
const (
	OutdoorParticulateMatterParserRULE_expression               = 0
	OutdoorParticulateMatterParserRULE_outdoorparticulatematter = 1
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
	p.RuleIndex = OutdoorParticulateMatterParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OutdoorParticulateMatterParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Outdoorparticulatematter() IOutdoorparticulatematterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOutdoorparticulatematterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOutdoorparticulatematterContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OutdoorParticulateMatterParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OutdoorParticulateMatterListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OutdoorParticulateMatterListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OutdoorParticulateMatterParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OutdoorParticulateMatterParserRULE_expression)

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
		p.Outdoorparticulatematter()
	}
	{
		p.SetState(5)
		p.Match(OutdoorParticulateMatterParserEOF)
	}

	return localctx
}

// IOutdoorparticulatematterContext is an interface to support dynamic dispatch.
type IOutdoorparticulatematterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutdoorparticulatematterContext differentiates from other interfaces.
	IsOutdoorparticulatematterContext()
}

type OutdoorparticulatematterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutdoorparticulatematterContext() *OutdoorparticulatematterContext {
	var p = new(OutdoorparticulatematterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OutdoorParticulateMatterParserRULE_outdoorparticulatematter
	return p
}

func (*OutdoorparticulatematterContext) IsOutdoorparticulatematterContext() {}

func NewOutdoorparticulatematterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutdoorparticulatematterContext {
	var p = new(OutdoorparticulatematterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OutdoorParticulateMatterParserRULE_outdoorparticulatematter

	return p
}

func (s *OutdoorparticulatematterContext) GetParser() antlr.Parser { return s.parser }

func (s *OutdoorparticulatematterContext) OUTDOORPARTICULATEMATTER() antlr.TerminalNode {
	return s.GetToken(OutdoorParticulateMatterParserOUTDOORPARTICULATEMATTER, 0)
}

func (s *OutdoorparticulatematterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutdoorparticulatematterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OutdoorparticulatematterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OutdoorParticulateMatterListener); ok {
		listenerT.EnterOutdoorparticulatematter(s)
	}
}

func (s *OutdoorparticulatematterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OutdoorParticulateMatterListener); ok {
		listenerT.ExitOutdoorparticulatematter(s)
	}
}

func (p *OutdoorParticulateMatterParser) Outdoorparticulatematter() (localctx IOutdoorparticulatematterContext) {
	this := p
	_ = this

	localctx = NewOutdoorparticulatematterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OutdoorParticulateMatterParserRULE_outdoorparticulatematter)

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
		p.Match(OutdoorParticulateMatterParserOUTDOORPARTICULATEMATTER)
	}

	return localctx
}
