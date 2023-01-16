// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878091669/Dysarthria.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dysarthria

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

type DysarthriaParser struct {
	*antlr.BaseParser
}

var dysarthriaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dysarthriaParserInit() {
	staticData := &dysarthriaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DYSARTHRIA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "dysarthria",
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

// DysarthriaParserInit initializes any static state used to implement DysarthriaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDysarthriaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DysarthriaParserInit() {
	staticData := &dysarthriaParserStaticData
	staticData.once.Do(dysarthriaParserInit)
}

// NewDysarthriaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDysarthriaParser(input antlr.TokenStream) *DysarthriaParser {
	DysarthriaParserInit()
	this := new(DysarthriaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &dysarthriaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Dysarthria.g4"

	return this
}

// DysarthriaParser tokens.
const (
	DysarthriaParserEOF        = antlr.TokenEOF
	DysarthriaParserDYSARTHRIA = 1
	DysarthriaParserOWNKEY     = 2
	DysarthriaParserSPLITKEY   = 3
	DysarthriaParserWS         = 4
)

// DysarthriaParser rules.
const (
	DysarthriaParserRULE_expression = 0
	DysarthriaParserRULE_dysarthria = 1
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
	p.RuleIndex = DysarthriaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DysarthriaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Dysarthria() IDysarthriaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDysarthriaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDysarthriaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DysarthriaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysarthriaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysarthriaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DysarthriaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DysarthriaParserRULE_expression)

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
		p.Dysarthria()
	}
	{
		p.SetState(5)
		p.Match(DysarthriaParserEOF)
	}

	return localctx
}

// IDysarthriaContext is an interface to support dynamic dispatch.
type IDysarthriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDysarthriaContext differentiates from other interfaces.
	IsDysarthriaContext()
}

type DysarthriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDysarthriaContext() *DysarthriaContext {
	var p = new(DysarthriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DysarthriaParserRULE_dysarthria
	return p
}

func (*DysarthriaContext) IsDysarthriaContext() {}

func NewDysarthriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DysarthriaContext {
	var p = new(DysarthriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DysarthriaParserRULE_dysarthria

	return p
}

func (s *DysarthriaContext) GetParser() antlr.Parser { return s.parser }

func (s *DysarthriaContext) DYSARTHRIA() antlr.TerminalNode {
	return s.GetToken(DysarthriaParserDYSARTHRIA, 0)
}

func (s *DysarthriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DysarthriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DysarthriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysarthriaListener); ok {
		listenerT.EnterDysarthria(s)
	}
}

func (s *DysarthriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DysarthriaListener); ok {
		listenerT.ExitDysarthria(s)
	}
}

func (p *DysarthriaParser) Dysarthria() (localctx IDysarthriaContext) {
	this := p
	_ = this

	localctx = NewDysarthriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DysarthriaParserRULE_dysarthria)

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
		p.Match(DysarthriaParserDYSARTHRIA)
	}

	return localctx
}
