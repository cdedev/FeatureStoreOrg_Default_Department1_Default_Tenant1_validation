// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671206924431/Model.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Model

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

type ModelParser struct {
	*antlr.BaseParser
}

var modelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func modelParserInit() {
	staticData := &modelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MODEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "model",
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

// ModelParserInit initializes any static state used to implement ModelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewModelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ModelParserInit() {
	staticData := &modelParserStaticData
	staticData.once.Do(modelParserInit)
}

// NewModelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewModelParser(input antlr.TokenStream) *ModelParser {
	ModelParserInit()
	this := new(ModelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &modelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Model.g4"

	return this
}

// ModelParser tokens.
const (
	ModelParserEOF      = antlr.TokenEOF
	ModelParserMODEL    = 1
	ModelParserOWNKEY   = 2
	ModelParserSPLITKEY = 3
	ModelParserWS       = 4
)

// ModelParser rules.
const (
	ModelParserRULE_expression = 0
	ModelParserRULE_model      = 1
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
	p.RuleIndex = ModelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Model() IModelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ModelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ModelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ModelParserRULE_expression)

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
		p.Model()
	}
	{
		p.SetState(5)
		p.Match(ModelParserEOF)
	}

	return localctx
}

// IModelContext is an interface to support dynamic dispatch.
type IModelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelContext differentiates from other interfaces.
	IsModelContext()
}

type ModelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelContext() *ModelContext {
	var p = new(ModelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ModelParserRULE_model
	return p
}

func (*ModelContext) IsModelContext() {}

func NewModelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelContext {
	var p = new(ModelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ModelParserRULE_model

	return p
}

func (s *ModelContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelContext) MODEL() antlr.TerminalNode {
	return s.GetToken(ModelParserMODEL, 0)
}

func (s *ModelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelListener); ok {
		listenerT.EnterModel(s)
	}
}

func (s *ModelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModelListener); ok {
		listenerT.ExitModel(s)
	}
}

func (p *ModelParser) Model() (localctx IModelContext) {
	this := p
	_ = this

	localctx = NewModelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ModelParserRULE_model)

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
		p.Match(ModelParserMODEL)
	}

	return localctx
}
