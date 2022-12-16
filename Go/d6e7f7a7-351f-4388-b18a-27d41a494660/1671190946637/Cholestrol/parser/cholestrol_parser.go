// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671190946637/Cholestrol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cholestrol

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

type CholestrolParser struct {
	*antlr.BaseParser
}

var cholestrolParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cholestrolParserInit() {
	staticData := &cholestrolParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHOLESTROL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cholestrol",
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

// CholestrolParserInit initializes any static state used to implement CholestrolParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCholestrolParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CholestrolParserInit() {
	staticData := &cholestrolParserStaticData
	staticData.once.Do(cholestrolParserInit)
}

// NewCholestrolParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCholestrolParser(input antlr.TokenStream) *CholestrolParser {
	CholestrolParserInit()
	this := new(CholestrolParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cholestrolParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cholestrol.g4"

	return this
}

// CholestrolParser tokens.
const (
	CholestrolParserEOF        = antlr.TokenEOF
	CholestrolParserCHOLESTROL = 1
	CholestrolParserOWNKEY     = 2
	CholestrolParserSPLITKEY   = 3
	CholestrolParserWS         = 4
)

// CholestrolParser rules.
const (
	CholestrolParserRULE_expression = 0
	CholestrolParserRULE_cholestrol = 1
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
	p.RuleIndex = CholestrolParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CholestrolParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cholestrol() ICholestrolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICholestrolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICholestrolContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CholestrolParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholestrolListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholestrolListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CholestrolParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CholestrolParserRULE_expression)

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
		p.Cholestrol()
	}
	{
		p.SetState(5)
		p.Match(CholestrolParserEOF)
	}

	return localctx
}

// ICholestrolContext is an interface to support dynamic dispatch.
type ICholestrolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCholestrolContext differentiates from other interfaces.
	IsCholestrolContext()
}

type CholestrolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCholestrolContext() *CholestrolContext {
	var p = new(CholestrolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CholestrolParserRULE_cholestrol
	return p
}

func (*CholestrolContext) IsCholestrolContext() {}

func NewCholestrolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CholestrolContext {
	var p = new(CholestrolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CholestrolParserRULE_cholestrol

	return p
}

func (s *CholestrolContext) GetParser() antlr.Parser { return s.parser }

func (s *CholestrolContext) CHOLESTROL() antlr.TerminalNode {
	return s.GetToken(CholestrolParserCHOLESTROL, 0)
}

func (s *CholestrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CholestrolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CholestrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholestrolListener); ok {
		listenerT.EnterCholestrol(s)
	}
}

func (s *CholestrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholestrolListener); ok {
		listenerT.ExitCholestrol(s)
	}
}

func (p *CholestrolParser) Cholestrol() (localctx ICholestrolContext) {
	this := p
	_ = this

	localctx = NewCholestrolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CholestrolParserRULE_cholestrol)

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
		p.Match(CholestrolParserCHOLESTROL)
	}

	return localctx
}
