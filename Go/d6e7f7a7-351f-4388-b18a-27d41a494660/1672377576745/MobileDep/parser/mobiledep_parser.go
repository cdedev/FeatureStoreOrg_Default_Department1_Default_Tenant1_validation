// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377576745/MobileDep.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MobileDep

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

type MobileDepParser struct {
	*antlr.BaseParser
}

var mobiledepParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mobiledepParserInit() {
	staticData := &mobiledepParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MOBILEDEP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "mobiledep",
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

// MobileDepParserInit initializes any static state used to implement MobileDepParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMobileDepParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MobileDepParserInit() {
	staticData := &mobiledepParserStaticData
	staticData.once.Do(mobiledepParserInit)
}

// NewMobileDepParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMobileDepParser(input antlr.TokenStream) *MobileDepParser {
	MobileDepParserInit()
	this := new(MobileDepParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &mobiledepParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MobileDep.g4"

	return this
}

// MobileDepParser tokens.
const (
	MobileDepParserEOF       = antlr.TokenEOF
	MobileDepParserMOBILEDEP = 1
	MobileDepParserOWNKEY    = 2
	MobileDepParserSPLITKEY  = 3
	MobileDepParserWS        = 4
)

// MobileDepParser rules.
const (
	MobileDepParserRULE_expression = 0
	MobileDepParserRULE_mobiledep  = 1
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
	p.RuleIndex = MobileDepParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MobileDepParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Mobiledep() IMobiledepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMobiledepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMobiledepContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MobileDepParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileDepListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileDepListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MobileDepParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MobileDepParserRULE_expression)

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
		p.Mobiledep()
	}
	{
		p.SetState(5)
		p.Match(MobileDepParserEOF)
	}

	return localctx
}

// IMobiledepContext is an interface to support dynamic dispatch.
type IMobiledepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMobiledepContext differentiates from other interfaces.
	IsMobiledepContext()
}

type MobiledepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMobiledepContext() *MobiledepContext {
	var p = new(MobiledepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MobileDepParserRULE_mobiledep
	return p
}

func (*MobiledepContext) IsMobiledepContext() {}

func NewMobiledepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MobiledepContext {
	var p = new(MobiledepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MobileDepParserRULE_mobiledep

	return p
}

func (s *MobiledepContext) GetParser() antlr.Parser { return s.parser }

func (s *MobiledepContext) MOBILEDEP() antlr.TerminalNode {
	return s.GetToken(MobileDepParserMOBILEDEP, 0)
}

func (s *MobiledepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MobiledepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MobiledepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileDepListener); ok {
		listenerT.EnterMobiledep(s)
	}
}

func (s *MobiledepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MobileDepListener); ok {
		listenerT.ExitMobiledep(s)
	}
}

func (p *MobileDepParser) Mobiledep() (localctx IMobiledepContext) {
	this := p
	_ = this

	localctx = NewMobiledepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MobileDepParserRULE_mobiledep)

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
		p.Match(MobileDepParserMOBILEDEP)
	}

	return localctx
}
