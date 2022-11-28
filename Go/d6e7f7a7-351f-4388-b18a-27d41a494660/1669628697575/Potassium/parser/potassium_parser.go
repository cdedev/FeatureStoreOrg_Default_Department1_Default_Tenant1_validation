// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628697575/Potassium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Potassium

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

type PotassiumParser struct {
	*antlr.BaseParser
}

var potassiumParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func potassiumParserInit() {
	staticData := &potassiumParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "POTASSIUM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "potassium",
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

// PotassiumParserInit initializes any static state used to implement PotassiumParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPotassiumParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PotassiumParserInit() {
	staticData := &potassiumParserStaticData
	staticData.once.Do(potassiumParserInit)
}

// NewPotassiumParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPotassiumParser(input antlr.TokenStream) *PotassiumParser {
	PotassiumParserInit()
	this := new(PotassiumParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &potassiumParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Potassium.g4"

	return this
}

// PotassiumParser tokens.
const (
	PotassiumParserEOF       = antlr.TokenEOF
	PotassiumParserPOTASSIUM = 1
	PotassiumParserOWNKEY    = 2
	PotassiumParserSPLITKEY  = 3
	PotassiumParserWS        = 4
)

// PotassiumParser rules.
const (
	PotassiumParserRULE_expression = 0
	PotassiumParserRULE_potassium  = 1
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
	p.RuleIndex = PotassiumParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PotassiumParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Potassium() IPotassiumContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPotassiumContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPotassiumContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PotassiumParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PotassiumListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PotassiumListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PotassiumParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PotassiumParserRULE_expression)

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
		p.Potassium()
	}
	{
		p.SetState(5)
		p.Match(PotassiumParserEOF)
	}

	return localctx
}

// IPotassiumContext is an interface to support dynamic dispatch.
type IPotassiumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPotassiumContext differentiates from other interfaces.
	IsPotassiumContext()
}

type PotassiumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPotassiumContext() *PotassiumContext {
	var p = new(PotassiumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PotassiumParserRULE_potassium
	return p
}

func (*PotassiumContext) IsPotassiumContext() {}

func NewPotassiumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PotassiumContext {
	var p = new(PotassiumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PotassiumParserRULE_potassium

	return p
}

func (s *PotassiumContext) GetParser() antlr.Parser { return s.parser }

func (s *PotassiumContext) POTASSIUM() antlr.TerminalNode {
	return s.GetToken(PotassiumParserPOTASSIUM, 0)
}

func (s *PotassiumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PotassiumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PotassiumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PotassiumListener); ok {
		listenerT.EnterPotassium(s)
	}
}

func (s *PotassiumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PotassiumListener); ok {
		listenerT.ExitPotassium(s)
	}
}

func (p *PotassiumParser) Potassium() (localctx IPotassiumContext) {
	this := p
	_ = this

	localctx = NewPotassiumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PotassiumParserRULE_potassium)

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
		p.Match(PotassiumParserPOTASSIUM)
	}

	return localctx
}
