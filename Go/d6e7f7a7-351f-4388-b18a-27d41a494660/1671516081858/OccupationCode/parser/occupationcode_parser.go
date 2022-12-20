// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516081858/OccupationCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OccupationCode

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

type OccupationCodeParser struct {
	*antlr.BaseParser
}

var occupationcodeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func occupationcodeParserInit() {
	staticData := &occupationcodeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "OCCUPATIONCODE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "occupationcode",
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

// OccupationCodeParserInit initializes any static state used to implement OccupationCodeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOccupationCodeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OccupationCodeParserInit() {
	staticData := &occupationcodeParserStaticData
	staticData.once.Do(occupationcodeParserInit)
}

// NewOccupationCodeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOccupationCodeParser(input antlr.TokenStream) *OccupationCodeParser {
	OccupationCodeParserInit()
	this := new(OccupationCodeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &occupationcodeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "OccupationCode.g4"

	return this
}

// OccupationCodeParser tokens.
const (
	OccupationCodeParserEOF            = antlr.TokenEOF
	OccupationCodeParserOCCUPATIONCODE = 1
	OccupationCodeParserOWNKEY         = 2
	OccupationCodeParserSPLITKEY       = 3
	OccupationCodeParserWS             = 4
)

// OccupationCodeParser rules.
const (
	OccupationCodeParserRULE_expression     = 0
	OccupationCodeParserRULE_occupationcode = 1
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
	p.RuleIndex = OccupationCodeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OccupationCodeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Occupationcode() IOccupationcodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOccupationcodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOccupationcodeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OccupationCodeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationCodeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationCodeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OccupationCodeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OccupationCodeParserRULE_expression)

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
		p.Occupationcode()
	}
	{
		p.SetState(5)
		p.Match(OccupationCodeParserEOF)
	}

	return localctx
}

// IOccupationcodeContext is an interface to support dynamic dispatch.
type IOccupationcodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOccupationcodeContext differentiates from other interfaces.
	IsOccupationcodeContext()
}

type OccupationcodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOccupationcodeContext() *OccupationcodeContext {
	var p = new(OccupationcodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OccupationCodeParserRULE_occupationcode
	return p
}

func (*OccupationcodeContext) IsOccupationcodeContext() {}

func NewOccupationcodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OccupationcodeContext {
	var p = new(OccupationcodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OccupationCodeParserRULE_occupationcode

	return p
}

func (s *OccupationcodeContext) GetParser() antlr.Parser { return s.parser }

func (s *OccupationcodeContext) OCCUPATIONCODE() antlr.TerminalNode {
	return s.GetToken(OccupationCodeParserOCCUPATIONCODE, 0)
}

func (s *OccupationcodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OccupationcodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OccupationcodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationCodeListener); ok {
		listenerT.EnterOccupationcode(s)
	}
}

func (s *OccupationcodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OccupationCodeListener); ok {
		listenerT.ExitOccupationcode(s)
	}
}

func (p *OccupationCodeParser) Occupationcode() (localctx IOccupationcodeContext) {
	this := p
	_ = this

	localctx = NewOccupationcodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OccupationCodeParserRULE_occupationcode)

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
		p.Match(OccupationCodeParserOCCUPATIONCODE)
	}

	return localctx
}
