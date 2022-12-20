// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515680829/Patients.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Patients

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

type PatientsParser struct {
	*antlr.BaseParser
}

var patientsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func patientsParserInit() {
	staticData := &patientsParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PATIENTS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "patients",
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

// PatientsParserInit initializes any static state used to implement PatientsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPatientsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PatientsParserInit() {
	staticData := &patientsParserStaticData
	staticData.once.Do(patientsParserInit)
}

// NewPatientsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPatientsParser(input antlr.TokenStream) *PatientsParser {
	PatientsParserInit()
	this := new(PatientsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &patientsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Patients.g4"

	return this
}

// PatientsParser tokens.
const (
	PatientsParserEOF      = antlr.TokenEOF
	PatientsParserPATIENTS = 1
	PatientsParserOWNKEY   = 2
	PatientsParserSPLITKEY = 3
	PatientsParserWS       = 4
)

// PatientsParser rules.
const (
	PatientsParserRULE_expression = 0
	PatientsParserRULE_patients   = 1
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
	p.RuleIndex = PatientsParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatientsParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Patients() IPatientsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPatientsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPatientsContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PatientsParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatientsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatientsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PatientsParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PatientsParserRULE_expression)

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
		p.Patients()
	}
	{
		p.SetState(5)
		p.Match(PatientsParserEOF)
	}

	return localctx
}

// IPatientsContext is an interface to support dynamic dispatch.
type IPatientsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatientsContext differentiates from other interfaces.
	IsPatientsContext()
}

type PatientsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatientsContext() *PatientsContext {
	var p = new(PatientsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PatientsParserRULE_patients
	return p
}

func (*PatientsContext) IsPatientsContext() {}

func NewPatientsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatientsContext {
	var p = new(PatientsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PatientsParserRULE_patients

	return p
}

func (s *PatientsContext) GetParser() antlr.Parser { return s.parser }

func (s *PatientsContext) PATIENTS() antlr.TerminalNode {
	return s.GetToken(PatientsParserPATIENTS, 0)
}

func (s *PatientsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatientsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatientsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatientsListener); ok {
		listenerT.EnterPatients(s)
	}
}

func (s *PatientsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PatientsListener); ok {
		listenerT.ExitPatients(s)
	}
}

func (p *PatientsParser) Patients() (localctx IPatientsContext) {
	this := p
	_ = this

	localctx = NewPatientsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PatientsParserRULE_patients)

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
		p.Match(PatientsParserPATIENTS)
	}

	return localctx
}
