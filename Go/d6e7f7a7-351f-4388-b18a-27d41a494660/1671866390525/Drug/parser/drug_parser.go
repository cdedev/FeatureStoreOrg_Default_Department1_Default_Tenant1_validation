// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671866390525/Drug.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Drug

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

type DrugParser struct {
	*antlr.BaseParser
}

var drugParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func drugParserInit() {
	staticData := &drugParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DRUG", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "drug",
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

// DrugParserInit initializes any static state used to implement DrugParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDrugParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DrugParserInit() {
	staticData := &drugParserStaticData
	staticData.once.Do(drugParserInit)
}

// NewDrugParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDrugParser(input antlr.TokenStream) *DrugParser {
	DrugParserInit()
	this := new(DrugParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &drugParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Drug.g4"

	return this
}

// DrugParser tokens.
const (
	DrugParserEOF      = antlr.TokenEOF
	DrugParserDRUG     = 1
	DrugParserOWNKEY   = 2
	DrugParserSPLITKEY = 3
	DrugParserWS       = 4
)

// DrugParser rules.
const (
	DrugParserRULE_expression = 0
	DrugParserRULE_drug       = 1
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
	p.RuleIndex = DrugParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrugParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Drug() IDrugContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDrugContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDrugContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DrugParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DrugParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DrugParserRULE_expression)

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
		p.Drug()
	}
	{
		p.SetState(5)
		p.Match(DrugParserEOF)
	}

	return localctx
}

// IDrugContext is an interface to support dynamic dispatch.
type IDrugContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDrugContext differentiates from other interfaces.
	IsDrugContext()
}

type DrugContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDrugContext() *DrugContext {
	var p = new(DrugContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DrugParserRULE_drug
	return p
}

func (*DrugContext) IsDrugContext() {}

func NewDrugContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DrugContext {
	var p = new(DrugContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DrugParserRULE_drug

	return p
}

func (s *DrugContext) GetParser() antlr.Parser { return s.parser }

func (s *DrugContext) DRUG() antlr.TerminalNode {
	return s.GetToken(DrugParserDRUG, 0)
}

func (s *DrugContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DrugContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DrugContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugListener); ok {
		listenerT.EnterDrug(s)
	}
}

func (s *DrugContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DrugListener); ok {
		listenerT.ExitDrug(s)
	}
}

func (p *DrugParser) Drug() (localctx IDrugContext) {
	this := p
	_ = this

	localctx = NewDrugContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DrugParserRULE_drug)

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
		p.Match(DrugParserDRUG)
	}

	return localctx
}
