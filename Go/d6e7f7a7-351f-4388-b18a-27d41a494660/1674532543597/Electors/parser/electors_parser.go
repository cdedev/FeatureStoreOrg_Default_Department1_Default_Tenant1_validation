// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532543597/Electors.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Electors

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

type ElectorsParser struct {
	*antlr.BaseParser
}

var electorsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func electorsParserInit() {
	staticData := &electorsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ELECTORS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "electors",
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

// ElectorsParserInit initializes any static state used to implement ElectorsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewElectorsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ElectorsParserInit() {
	staticData := &electorsParserStaticData
	staticData.once.Do(electorsParserInit)
}

// NewElectorsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewElectorsParser(input antlr.TokenStream) *ElectorsParser {
	ElectorsParserInit()
	this := new(ElectorsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &electorsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Electors.g4"

	return this
}

// ElectorsParser tokens.
const (
	ElectorsParserEOF      = antlr.TokenEOF
	ElectorsParserELECTORS = 1
	ElectorsParserOWNKEY   = 2
	ElectorsParserSPLITKEY = 3
	ElectorsParserWS       = 4
)

// ElectorsParser rules.
const (
	ElectorsParserRULE_expression = 0
	ElectorsParserRULE_electors   = 1
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
	p.RuleIndex = ElectorsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElectorsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Electors() IElectorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElectorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElectorsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ElectorsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElectorsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElectorsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ElectorsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ElectorsParserRULE_expression)

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
		p.Electors()
	}
	{
		p.SetState(5)
		p.Match(ElectorsParserEOF)
	}

	return localctx
}

// IElectorsContext is an interface to support dynamic dispatch.
type IElectorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElectorsContext differentiates from other interfaces.
	IsElectorsContext()
}

type ElectorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElectorsContext() *ElectorsContext {
	var p = new(ElectorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ElectorsParserRULE_electors
	return p
}

func (*ElectorsContext) IsElectorsContext() {}

func NewElectorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElectorsContext {
	var p = new(ElectorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ElectorsParserRULE_electors

	return p
}

func (s *ElectorsContext) GetParser() antlr.Parser { return s.parser }

func (s *ElectorsContext) ELECTORS() antlr.TerminalNode {
	return s.GetToken(ElectorsParserELECTORS, 0)
}

func (s *ElectorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElectorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElectorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElectorsListener); ok {
		listenerT.EnterElectors(s)
	}
}

func (s *ElectorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ElectorsListener); ok {
		listenerT.ExitElectors(s)
	}
}

func (p *ElectorsParser) Electors() (localctx IElectorsContext) {
	this := p
	_ = this

	localctx = NewElectorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ElectorsParserRULE_electors)

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
		p.Match(ElectorsParserELECTORS)
	}

	return localctx
}
