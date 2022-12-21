// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604371088/Cerium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cerium

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

type CeriumParser struct {
	*antlr.BaseParser
}

var ceriumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ceriumParserInit() {
	staticData := &ceriumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CERIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cerium",
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

// CeriumParserInit initializes any static state used to implement CeriumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCeriumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CeriumParserInit() {
	staticData := &ceriumParserStaticData
	staticData.once.Do(ceriumParserInit)
}

// NewCeriumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCeriumParser(input antlr.TokenStream) *CeriumParser {
	CeriumParserInit()
	this := new(CeriumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ceriumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cerium.g4"

	return this
}

// CeriumParser tokens.
const (
	CeriumParserEOF      = antlr.TokenEOF
	CeriumParserCERIUM   = 1
	CeriumParserOWNKEY   = 2
	CeriumParserSPLITKEY = 3
	CeriumParserWS       = 4
)

// CeriumParser rules.
const (
	CeriumParserRULE_expression = 0
	CeriumParserRULE_cerium     = 1
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
	p.RuleIndex = CeriumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CeriumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cerium() ICeriumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICeriumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICeriumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CeriumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CeriumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CeriumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CeriumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CeriumParserRULE_expression)

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
		p.Cerium()
	}
	{
		p.SetState(5)
		p.Match(CeriumParserEOF)
	}

	return localctx
}

// ICeriumContext is an interface to support dynamic dispatch.
type ICeriumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCeriumContext differentiates from other interfaces.
	IsCeriumContext()
}

type CeriumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCeriumContext() *CeriumContext {
	var p = new(CeriumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CeriumParserRULE_cerium
	return p
}

func (*CeriumContext) IsCeriumContext() {}

func NewCeriumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CeriumContext {
	var p = new(CeriumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CeriumParserRULE_cerium

	return p
}

func (s *CeriumContext) GetParser() antlr.Parser { return s.parser }

func (s *CeriumContext) CERIUM() antlr.TerminalNode {
	return s.GetToken(CeriumParserCERIUM, 0)
}

func (s *CeriumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeriumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CeriumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CeriumListener); ok {
		listenerT.EnterCerium(s)
	}
}

func (s *CeriumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CeriumListener); ok {
		listenerT.ExitCerium(s)
	}
}

func (p *CeriumParser) Cerium() (localctx ICeriumContext) {
	this := p
	_ = this

	localctx = NewCeriumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CeriumParserRULE_cerium)

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
		p.Match(CeriumParserCERIUM)
	}

	return localctx
}
