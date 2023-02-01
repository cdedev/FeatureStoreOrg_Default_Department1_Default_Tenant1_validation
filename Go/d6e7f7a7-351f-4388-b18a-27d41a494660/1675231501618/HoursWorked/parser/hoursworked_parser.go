// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675231501618/HoursWorked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HoursWorked

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

type HoursWorkedParser struct {
	*antlr.BaseParser
}

var hoursworkedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func hoursworkedParserInit() {
	staticData := &hoursworkedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "HOURSWORKED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "hoursworked",
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

// HoursWorkedParserInit initializes any static state used to implement HoursWorkedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewHoursWorkedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func HoursWorkedParserInit() {
	staticData := &hoursworkedParserStaticData
	staticData.once.Do(hoursworkedParserInit)
}

// NewHoursWorkedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewHoursWorkedParser(input antlr.TokenStream) *HoursWorkedParser {
	HoursWorkedParserInit()
	this := new(HoursWorkedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &hoursworkedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "HoursWorked.g4"

	return this
}

// HoursWorkedParser tokens.
const (
	HoursWorkedParserEOF         = antlr.TokenEOF
	HoursWorkedParserHOURSWORKED = 1
	HoursWorkedParserOWNKEY      = 2
	HoursWorkedParserSPLITKEY    = 3
	HoursWorkedParserWS          = 4
)

// HoursWorkedParser rules.
const (
	HoursWorkedParserRULE_expression  = 0
	HoursWorkedParserRULE_hoursworked = 1
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
	p.RuleIndex = HoursWorkedParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoursWorkedParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Hoursworked() IHoursworkedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHoursworkedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHoursworkedContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(HoursWorkedParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoursWorkedListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoursWorkedListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *HoursWorkedParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, HoursWorkedParserRULE_expression)

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
		p.Hoursworked()
	}
	{
		p.SetState(5)
		p.Match(HoursWorkedParserEOF)
	}

	return localctx
}

// IHoursworkedContext is an interface to support dynamic dispatch.
type IHoursworkedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHoursworkedContext differentiates from other interfaces.
	IsHoursworkedContext()
}

type HoursworkedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHoursworkedContext() *HoursworkedContext {
	var p = new(HoursworkedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = HoursWorkedParserRULE_hoursworked
	return p
}

func (*HoursworkedContext) IsHoursworkedContext() {}

func NewHoursworkedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HoursworkedContext {
	var p = new(HoursworkedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = HoursWorkedParserRULE_hoursworked

	return p
}

func (s *HoursworkedContext) GetParser() antlr.Parser { return s.parser }

func (s *HoursworkedContext) HOURSWORKED() antlr.TerminalNode {
	return s.GetToken(HoursWorkedParserHOURSWORKED, 0)
}

func (s *HoursworkedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HoursworkedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HoursworkedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoursWorkedListener); ok {
		listenerT.EnterHoursworked(s)
	}
}

func (s *HoursworkedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(HoursWorkedListener); ok {
		listenerT.ExitHoursworked(s)
	}
}

func (p *HoursWorkedParser) Hoursworked() (localctx IHoursworkedContext) {
	this := p
	_ = this

	localctx = NewHoursworkedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, HoursWorkedParserRULE_hoursworked)

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
		p.Match(HoursWorkedParserHOURSWORKED)
	}

	return localctx
}
