// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671607864497/Mileage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mileage

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

type MileageParser struct {
	*antlr.BaseParser
}

var mileageParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mileageParserInit() {
	staticData := &mileageParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MILEAGE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mileage",
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

// MileageParserInit initializes any static state used to implement MileageParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMileageParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MileageParserInit() {
	staticData := &mileageParserStaticData
	staticData.once.Do(mileageParserInit)
}

// NewMileageParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMileageParser(input antlr.TokenStream) *MileageParser {
	MileageParserInit()
	this := new(MileageParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mileageParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Mileage.g4"

	return this
}

// MileageParser tokens.
const (
	MileageParserEOF      = antlr.TokenEOF
	MileageParserMILEAGE  = 1
	MileageParserOWNKEY   = 2
	MileageParserSPLITKEY = 3
	MileageParserWS       = 4
)

// MileageParser rules.
const (
	MileageParserRULE_expression = 0
	MileageParserRULE_mileage    = 1
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
	p.RuleIndex = MileageParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MileageParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mileage() IMileageContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMileageContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMileageContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MileageParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MileageListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MileageListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MileageParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MileageParserRULE_expression)

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
		p.Mileage()
	}
	{
		p.SetState(5)
		p.Match(MileageParserEOF)
	}

	return localctx
}

// IMileageContext is an interface to support dynamic dispatch.
type IMileageContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMileageContext differentiates from other interfaces.
	IsMileageContext()
}

type MileageContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMileageContext() *MileageContext {
	var p = new(MileageContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MileageParserRULE_mileage
	return p
}

func (*MileageContext) IsMileageContext() {}

func NewMileageContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MileageContext {
	var p = new(MileageContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MileageParserRULE_mileage

	return p
}

func (s *MileageContext) GetParser() antlr.Parser { return s.parser }

func (s *MileageContext) MILEAGE() antlr.TerminalNode {
	return s.GetToken(MileageParserMILEAGE, 0)
}

func (s *MileageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MileageContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MileageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MileageListener); ok {
		listenerT.EnterMileage(s)
	}
}

func (s *MileageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MileageListener); ok {
		listenerT.ExitMileage(s)
	}
}

func (p *MileageParser) Mileage() (localctx IMileageContext) {
	this := p
	_ = this

	localctx = NewMileageContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MileageParserRULE_mileage)

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
		p.Match(MileageParserMILEAGE)
	}

	return localctx
}
