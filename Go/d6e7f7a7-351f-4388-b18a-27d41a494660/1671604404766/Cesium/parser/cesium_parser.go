// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604404766/Cesium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cesium

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

type CesiumParser struct {
	*antlr.BaseParser
}

var cesiumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cesiumParserInit() {
	staticData := &cesiumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CESIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "cesium",
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

// CesiumParserInit initializes any static state used to implement CesiumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCesiumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CesiumParserInit() {
	staticData := &cesiumParserStaticData
	staticData.once.Do(cesiumParserInit)
}

// NewCesiumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCesiumParser(input antlr.TokenStream) *CesiumParser {
	CesiumParserInit()
	this := new(CesiumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &cesiumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Cesium.g4"

	return this
}

// CesiumParser tokens.
const (
	CesiumParserEOF      = antlr.TokenEOF
	CesiumParserCESIUM   = 1
	CesiumParserOWNKEY   = 2
	CesiumParserSPLITKEY = 3
	CesiumParserWS       = 4
)

// CesiumParser rules.
const (
	CesiumParserRULE_expression = 0
	CesiumParserRULE_cesium     = 1
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
	p.RuleIndex = CesiumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CesiumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Cesium() ICesiumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICesiumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICesiumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CesiumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CesiumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CesiumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CesiumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CesiumParserRULE_expression)

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
		p.Cesium()
	}
	{
		p.SetState(5)
		p.Match(CesiumParserEOF)
	}

	return localctx
}

// ICesiumContext is an interface to support dynamic dispatch.
type ICesiumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCesiumContext differentiates from other interfaces.
	IsCesiumContext()
}

type CesiumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCesiumContext() *CesiumContext {
	var p = new(CesiumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CesiumParserRULE_cesium
	return p
}

func (*CesiumContext) IsCesiumContext() {}

func NewCesiumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CesiumContext {
	var p = new(CesiumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CesiumParserRULE_cesium

	return p
}

func (s *CesiumContext) GetParser() antlr.Parser { return s.parser }

func (s *CesiumContext) CESIUM() antlr.TerminalNode {
	return s.GetToken(CesiumParserCESIUM, 0)
}

func (s *CesiumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CesiumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CesiumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CesiumListener); ok {
		listenerT.EnterCesium(s)
	}
}

func (s *CesiumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CesiumListener); ok {
		listenerT.ExitCesium(s)
	}
}

func (p *CesiumParser) Cesium() (localctx ICesiumContext) {
	this := p
	_ = this

	localctx = NewCesiumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CesiumParserRULE_cesium)

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
		p.Match(CesiumParserCESIUM)
	}

	return localctx
}
