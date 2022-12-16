// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203051045/Parkingplaces.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Parkingplaces

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

type ParkingplacesParser struct {
	*antlr.BaseParser
}

var parkingplacesParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parkingplacesParserInit() {
	staticData := &parkingplacesParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PARKINGPLACES", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "parkingplaces",
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

// ParkingplacesParserInit initializes any static state used to implement ParkingplacesParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParkingplacesParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParkingplacesParserInit() {
	staticData := &parkingplacesParserStaticData
	staticData.once.Do(parkingplacesParserInit)
}

// NewParkingplacesParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParkingplacesParser(input antlr.TokenStream) *ParkingplacesParser {
	ParkingplacesParserInit()
	this := new(ParkingplacesParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &parkingplacesParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Parkingplaces.g4"

	return this
}

// ParkingplacesParser tokens.
const (
	ParkingplacesParserEOF           = antlr.TokenEOF
	ParkingplacesParserPARKINGPLACES = 1
	ParkingplacesParserOWNKEY        = 2
	ParkingplacesParserSPLITKEY      = 3
	ParkingplacesParserWS            = 4
)

// ParkingplacesParser rules.
const (
	ParkingplacesParserRULE_expression    = 0
	ParkingplacesParserRULE_parkingplaces = 1
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
	p.RuleIndex = ParkingplacesParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParkingplacesParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Parkingplaces() IParkingplacesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParkingplacesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParkingplacesContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParkingplacesParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParkingplacesListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParkingplacesListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ParkingplacesParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParkingplacesParserRULE_expression)

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
		p.Parkingplaces()
	}
	{
		p.SetState(5)
		p.Match(ParkingplacesParserEOF)
	}

	return localctx
}

// IParkingplacesContext is an interface to support dynamic dispatch.
type IParkingplacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParkingplacesContext differentiates from other interfaces.
	IsParkingplacesContext()
}

type ParkingplacesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParkingplacesContext() *ParkingplacesContext {
	var p = new(ParkingplacesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ParkingplacesParserRULE_parkingplaces
	return p
}

func (*ParkingplacesContext) IsParkingplacesContext() {}

func NewParkingplacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParkingplacesContext {
	var p = new(ParkingplacesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParkingplacesParserRULE_parkingplaces

	return p
}

func (s *ParkingplacesContext) GetParser() antlr.Parser { return s.parser }

func (s *ParkingplacesContext) PARKINGPLACES() antlr.TerminalNode {
	return s.GetToken(ParkingplacesParserPARKINGPLACES, 0)
}

func (s *ParkingplacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParkingplacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParkingplacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParkingplacesListener); ok {
		listenerT.EnterParkingplaces(s)
	}
}

func (s *ParkingplacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParkingplacesListener); ok {
		listenerT.ExitParkingplaces(s)
	}
}

func (p *ParkingplacesParser) Parkingplaces() (localctx IParkingplacesContext) {
	this := p
	_ = this

	localctx = NewParkingplacesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParkingplacesParserRULE_parkingplaces)

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
		p.Match(ParkingplacesParserPARKINGPLACES)
	}

	return localctx
}
