// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675674332598/Disease.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Disease

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

type DiseaseParser struct {
	*antlr.BaseParser
}

var diseaseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diseaseParserInit() {
	staticData := &diseaseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DISEASE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "disease",
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

// DiseaseParserInit initializes any static state used to implement DiseaseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiseaseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiseaseParserInit() {
	staticData := &diseaseParserStaticData
	staticData.once.Do(diseaseParserInit)
}

// NewDiseaseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiseaseParser(input antlr.TokenStream) *DiseaseParser {
	DiseaseParserInit()
	this := new(DiseaseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diseaseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Disease.g4"

	return this
}

// DiseaseParser tokens.
const (
	DiseaseParserEOF      = antlr.TokenEOF
	DiseaseParserDISEASE  = 1
	DiseaseParserOWNKEY   = 2
	DiseaseParserSPLITKEY = 3
	DiseaseParserWS       = 4
)

// DiseaseParser rules.
const (
	DiseaseParserRULE_expression = 0
	DiseaseParserRULE_disease    = 1
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
	p.RuleIndex = DiseaseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiseaseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Disease() IDiseaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiseaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiseaseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiseaseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseaseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseaseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiseaseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiseaseParserRULE_expression)

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
		p.Disease()
	}
	{
		p.SetState(5)
		p.Match(DiseaseParserEOF)
	}

	return localctx
}

// IDiseaseContext is an interface to support dynamic dispatch.
type IDiseaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiseaseContext differentiates from other interfaces.
	IsDiseaseContext()
}

type DiseaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiseaseContext() *DiseaseContext {
	var p = new(DiseaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiseaseParserRULE_disease
	return p
}

func (*DiseaseContext) IsDiseaseContext() {}

func NewDiseaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiseaseContext {
	var p = new(DiseaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiseaseParserRULE_disease

	return p
}

func (s *DiseaseContext) GetParser() antlr.Parser { return s.parser }

func (s *DiseaseContext) DISEASE() antlr.TerminalNode {
	return s.GetToken(DiseaseParserDISEASE, 0)
}

func (s *DiseaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiseaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiseaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseaseListener); ok {
		listenerT.EnterDisease(s)
	}
}

func (s *DiseaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiseaseListener); ok {
		listenerT.ExitDisease(s)
	}
}

func (p *DiseaseParser) Disease() (localctx IDiseaseContext) {
	this := p
	_ = this

	localctx = NewDiseaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiseaseParserRULE_disease)

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
		p.Match(DiseaseParserDISEASE)
	}

	return localctx
}
