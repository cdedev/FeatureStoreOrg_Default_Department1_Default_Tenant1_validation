// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675833307329/Cholesterol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cholesterol

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

type CholesterolParser struct {
	*antlr.BaseParser
}

var cholesterolParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cholesterolParserInit() {
	staticData := &cholesterolParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CHOLESTEROL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cholesterol",
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

// CholesterolParserInit initializes any static state used to implement CholesterolParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCholesterolParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CholesterolParserInit() {
	staticData := &cholesterolParserStaticData
	staticData.once.Do(cholesterolParserInit)
}

// NewCholesterolParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCholesterolParser(input antlr.TokenStream) *CholesterolParser {
	CholesterolParserInit()
	this := new(CholesterolParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cholesterolParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cholesterol.g4"

	return this
}

// CholesterolParser tokens.
const (
	CholesterolParserEOF         = antlr.TokenEOF
	CholesterolParserCHOLESTEROL = 1
	CholesterolParserOWNKEY      = 2
	CholesterolParserSPLITKEY    = 3
	CholesterolParserWS          = 4
)

// CholesterolParser rules.
const (
	CholesterolParserRULE_expression  = 0
	CholesterolParserRULE_cholesterol = 1
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
	p.RuleIndex = CholesterolParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CholesterolParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cholesterol() ICholesterolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICholesterolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICholesterolContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CholesterolParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholesterolListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholesterolListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CholesterolParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CholesterolParserRULE_expression)

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
		p.Cholesterol()
	}
	{
		p.SetState(5)
		p.Match(CholesterolParserEOF)
	}

	return localctx
}

// ICholesterolContext is an interface to support dynamic dispatch.
type ICholesterolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCholesterolContext differentiates from other interfaces.
	IsCholesterolContext()
}

type CholesterolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCholesterolContext() *CholesterolContext {
	var p = new(CholesterolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CholesterolParserRULE_cholesterol
	return p
}

func (*CholesterolContext) IsCholesterolContext() {}

func NewCholesterolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CholesterolContext {
	var p = new(CholesterolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CholesterolParserRULE_cholesterol

	return p
}

func (s *CholesterolContext) GetParser() antlr.Parser { return s.parser }

func (s *CholesterolContext) CHOLESTEROL() antlr.TerminalNode {
	return s.GetToken(CholesterolParserCHOLESTEROL, 0)
}

func (s *CholesterolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CholesterolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CholesterolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholesterolListener); ok {
		listenerT.EnterCholesterol(s)
	}
}

func (s *CholesterolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CholesterolListener); ok {
		listenerT.ExitCholesterol(s)
	}
}

func (p *CholesterolParser) Cholesterol() (localctx ICholesterolContext) {
	this := p
	_ = this

	localctx = NewCholesterolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CholesterolParserRULE_cholesterol)

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
		p.Match(CholesterolParserCHOLESTEROL)
	}

	return localctx
}
