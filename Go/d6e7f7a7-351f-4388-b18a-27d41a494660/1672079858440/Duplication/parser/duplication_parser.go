// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672079858440/Duplication.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Duplication

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

type DuplicationParser struct {
	*antlr.BaseParser
}

var duplicationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func duplicationParserInit() {
	staticData := &duplicationParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DUPLICATION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "duplication",
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

// DuplicationParserInit initializes any static state used to implement DuplicationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDuplicationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DuplicationParserInit() {
	staticData := &duplicationParserStaticData
	staticData.once.Do(duplicationParserInit)
}

// NewDuplicationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDuplicationParser(input antlr.TokenStream) *DuplicationParser {
	DuplicationParserInit()
	this := new(DuplicationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &duplicationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Duplication.g4"

	return this
}

// DuplicationParser tokens.
const (
	DuplicationParserEOF         = antlr.TokenEOF
	DuplicationParserDUPLICATION = 1
	DuplicationParserOWNKEY      = 2
	DuplicationParserSPLITKEY    = 3
	DuplicationParserWS          = 4
)

// DuplicationParser rules.
const (
	DuplicationParserRULE_expression  = 0
	DuplicationParserRULE_duplication = 1
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
	p.RuleIndex = DuplicationParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DuplicationParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Duplication() IDuplicationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDuplicationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDuplicationContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DuplicationParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DuplicationListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DuplicationListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DuplicationParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DuplicationParserRULE_expression)

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
		p.Duplication()
	}
	{
		p.SetState(5)
		p.Match(DuplicationParserEOF)
	}

	return localctx
}

// IDuplicationContext is an interface to support dynamic dispatch.
type IDuplicationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDuplicationContext differentiates from other interfaces.
	IsDuplicationContext()
}

type DuplicationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDuplicationContext() *DuplicationContext {
	var p = new(DuplicationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DuplicationParserRULE_duplication
	return p
}

func (*DuplicationContext) IsDuplicationContext() {}

func NewDuplicationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DuplicationContext {
	var p = new(DuplicationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DuplicationParserRULE_duplication

	return p
}

func (s *DuplicationContext) GetParser() antlr.Parser { return s.parser }

func (s *DuplicationContext) DUPLICATION() antlr.TerminalNode {
	return s.GetToken(DuplicationParserDUPLICATION, 0)
}

func (s *DuplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DuplicationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DuplicationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DuplicationListener); ok {
		listenerT.EnterDuplication(s)
	}
}

func (s *DuplicationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DuplicationListener); ok {
		listenerT.ExitDuplication(s)
	}
}

func (p *DuplicationParser) Duplication() (localctx IDuplicationContext) {
	this := p
	_ = this

	localctx = NewDuplicationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DuplicationParserRULE_duplication)

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
		p.Match(DuplicationParserDUPLICATION)
	}

	return localctx
}
