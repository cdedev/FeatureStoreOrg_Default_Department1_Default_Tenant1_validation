// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377433806/MobileWeight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MobileWeight

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

type MobileWeightParser struct {
	*antlr.BaseParser
}

var mobileweightParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mobileweightParserInit() {
	staticData := &mobileweightParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MOBILEWEIGHT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mobileweight",
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

// MobileWeightParserInit initializes any static state used to implement MobileWeightParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMobileWeightParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MobileWeightParserInit() {
	staticData := &mobileweightParserStaticData
	staticData.once.Do(mobileweightParserInit)
}

// NewMobileWeightParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMobileWeightParser(input antlr.TokenStream) *MobileWeightParser {
	MobileWeightParserInit()
	this := new(MobileWeightParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mobileweightParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MobileWeight.g4"

	return this
}

// MobileWeightParser tokens.
const (
	MobileWeightParserEOF          = antlr.TokenEOF
	MobileWeightParserMOBILEWEIGHT = 1
	MobileWeightParserOWNKEY       = 2
	MobileWeightParserSPLITKEY     = 3
	MobileWeightParserWS           = 4
)

// MobileWeightParser rules.
const (
	MobileWeightParserRULE_expression   = 0
	MobileWeightParserRULE_mobileweight = 1
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
	p.RuleIndex = MobileWeightParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MobileWeightParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mobileweight() IMobileweightContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMobileweightContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMobileweightContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MobileWeightParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileWeightListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileWeightListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MobileWeightParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MobileWeightParserRULE_expression)

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
		p.Mobileweight()
	}
	{
		p.SetState(5)
		p.Match(MobileWeightParserEOF)
	}

	return localctx
}

// IMobileweightContext is an interface to support dynamic dispatch.
type IMobileweightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMobileweightContext differentiates from other interfaces.
	IsMobileweightContext()
}

type MobileweightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMobileweightContext() *MobileweightContext {
	var p = new(MobileweightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MobileWeightParserRULE_mobileweight
	return p
}

func (*MobileweightContext) IsMobileweightContext() {}

func NewMobileweightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MobileweightContext {
	var p = new(MobileweightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MobileWeightParserRULE_mobileweight

	return p
}

func (s *MobileweightContext) GetParser() antlr.Parser { return s.parser }

func (s *MobileweightContext) MOBILEWEIGHT() antlr.TerminalNode {
	return s.GetToken(MobileWeightParserMOBILEWEIGHT, 0)
}

func (s *MobileweightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MobileweightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MobileweightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileWeightListener); ok {
		listenerT.EnterMobileweight(s)
	}
}

func (s *MobileweightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileWeightListener); ok {
		listenerT.ExitMobileweight(s)
	}
}

func (p *MobileWeightParser) Mobileweight() (localctx IMobileweightContext) {
	this := p
	_ = this

	localctx = NewMobileweightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MobileWeightParserRULE_mobileweight)

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
		p.Match(MobileWeightParserMOBILEWEIGHT)
	}

	return localctx
}
