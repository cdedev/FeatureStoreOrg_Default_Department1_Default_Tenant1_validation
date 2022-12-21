// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604083198/Cadmium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cadmium

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

type CadmiumParser struct {
	*antlr.BaseParser
}

var cadmiumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cadmiumParserInit() {
	staticData := &cadmiumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CADMIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cadmium",
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

// CadmiumParserInit initializes any static state used to implement CadmiumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCadmiumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CadmiumParserInit() {
	staticData := &cadmiumParserStaticData
	staticData.once.Do(cadmiumParserInit)
}

// NewCadmiumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCadmiumParser(input antlr.TokenStream) *CadmiumParser {
	CadmiumParserInit()
	this := new(CadmiumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cadmiumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cadmium.g4"

	return this
}

// CadmiumParser tokens.
const (
	CadmiumParserEOF      = antlr.TokenEOF
	CadmiumParserCADMIUM  = 1
	CadmiumParserOWNKEY   = 2
	CadmiumParserSPLITKEY = 3
	CadmiumParserWS       = 4
)

// CadmiumParser rules.
const (
	CadmiumParserRULE_expression = 0
	CadmiumParserRULE_cadmium    = 1
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
	p.RuleIndex = CadmiumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadmiumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cadmium() ICadmiumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICadmiumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICadmiumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CadmiumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadmiumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadmiumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CadmiumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CadmiumParserRULE_expression)

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
		p.Cadmium()
	}
	{
		p.SetState(5)
		p.Match(CadmiumParserEOF)
	}

	return localctx
}

// ICadmiumContext is an interface to support dynamic dispatch.
type ICadmiumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCadmiumContext differentiates from other interfaces.
	IsCadmiumContext()
}

type CadmiumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCadmiumContext() *CadmiumContext {
	var p = new(CadmiumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CadmiumParserRULE_cadmium
	return p
}

func (*CadmiumContext) IsCadmiumContext() {}

func NewCadmiumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CadmiumContext {
	var p = new(CadmiumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CadmiumParserRULE_cadmium

	return p
}

func (s *CadmiumContext) GetParser() antlr.Parser { return s.parser }

func (s *CadmiumContext) CADMIUM() antlr.TerminalNode {
	return s.GetToken(CadmiumParserCADMIUM, 0)
}

func (s *CadmiumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CadmiumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CadmiumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadmiumListener); ok {
		listenerT.EnterCadmium(s)
	}
}

func (s *CadmiumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CadmiumListener); ok {
		listenerT.ExitCadmium(s)
	}
}

func (p *CadmiumParser) Cadmium() (localctx ICadmiumContext) {
	this := p
	_ = this

	localctx = NewCadmiumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CadmiumParserRULE_cadmium)

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
		p.Match(CadmiumParserCADMIUM)
	}

	return localctx
}
