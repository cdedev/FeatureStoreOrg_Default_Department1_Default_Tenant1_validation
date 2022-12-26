// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076311786/Resistivity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Resistivity

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

type ResistivityParser struct {
	*antlr.BaseParser
}

var resistivityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func resistivityParserInit() {
	staticData := &resistivityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RESISTIVITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "resistivity",
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

// ResistivityParserInit initializes any static state used to implement ResistivityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewResistivityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ResistivityParserInit() {
	staticData := &resistivityParserStaticData
	staticData.once.Do(resistivityParserInit)
}

// NewResistivityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewResistivityParser(input antlr.TokenStream) *ResistivityParser {
	ResistivityParserInit()
	this := new(ResistivityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &resistivityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Resistivity.g4"

	return this
}

// ResistivityParser tokens.
const (
	ResistivityParserEOF         = antlr.TokenEOF
	ResistivityParserRESISTIVITY = 1
	ResistivityParserOWNKEY      = 2
	ResistivityParserSPLITKEY    = 3
	ResistivityParserWS          = 4
)

// ResistivityParser rules.
const (
	ResistivityParserRULE_expression  = 0
	ResistivityParserRULE_resistivity = 1
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
	p.RuleIndex = ResistivityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ResistivityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Resistivity() IResistivityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResistivityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResistivityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ResistivityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResistivityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResistivityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ResistivityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ResistivityParserRULE_expression)

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
		p.Resistivity()
	}
	{
		p.SetState(5)
		p.Match(ResistivityParserEOF)
	}

	return localctx
}

// IResistivityContext is an interface to support dynamic dispatch.
type IResistivityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResistivityContext differentiates from other interfaces.
	IsResistivityContext()
}

type ResistivityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResistivityContext() *ResistivityContext {
	var p = new(ResistivityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ResistivityParserRULE_resistivity
	return p
}

func (*ResistivityContext) IsResistivityContext() {}

func NewResistivityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResistivityContext {
	var p = new(ResistivityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ResistivityParserRULE_resistivity

	return p
}

func (s *ResistivityContext) GetParser() antlr.Parser { return s.parser }

func (s *ResistivityContext) RESISTIVITY() antlr.TerminalNode {
	return s.GetToken(ResistivityParserRESISTIVITY, 0)
}

func (s *ResistivityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResistivityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResistivityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResistivityListener); ok {
		listenerT.EnterResistivity(s)
	}
}

func (s *ResistivityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ResistivityListener); ok {
		listenerT.ExitResistivity(s)
	}
}

func (p *ResistivityParser) Resistivity() (localctx IResistivityContext) {
	this := p
	_ = this

	localctx = NewResistivityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ResistivityParserRULE_resistivity)

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
		p.Match(ResistivityParserRESISTIVITY)
	}

	return localctx
}
