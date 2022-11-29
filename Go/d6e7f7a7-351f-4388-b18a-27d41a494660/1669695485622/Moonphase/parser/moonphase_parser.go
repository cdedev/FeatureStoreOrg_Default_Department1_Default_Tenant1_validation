// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669695485622/Moonphase.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Moonphase

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

type MoonphaseParser struct {
	*antlr.BaseParser
}

var moonphaseParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func moonphaseParserInit() {
	staticData := &moonphaseParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MOONPHASE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "moonphase",
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

// MoonphaseParserInit initializes any static state used to implement MoonphaseParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMoonphaseParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MoonphaseParserInit() {
	staticData := &moonphaseParserStaticData
	staticData.once.Do(moonphaseParserInit)
}

// NewMoonphaseParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMoonphaseParser(input antlr.TokenStream) *MoonphaseParser {
	MoonphaseParserInit()
	this := new(MoonphaseParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &moonphaseParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Moonphase.g4"

	return this
}

// MoonphaseParser tokens.
const (
	MoonphaseParserEOF       = antlr.TokenEOF
	MoonphaseParserMOONPHASE = 1
	MoonphaseParserOWNKEY    = 2
	MoonphaseParserSPLITKEY  = 3
	MoonphaseParserWS        = 4
)

// MoonphaseParser rules.
const (
	MoonphaseParserRULE_expression = 0
	MoonphaseParserRULE_moonphase  = 1
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
	p.RuleIndex = MoonphaseParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MoonphaseParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Moonphase() IMoonphaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMoonphaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMoonphaseContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MoonphaseParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MoonphaseListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MoonphaseListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MoonphaseParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MoonphaseParserRULE_expression)

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
		p.Moonphase()
	}
	{
		p.SetState(5)
		p.Match(MoonphaseParserEOF)
	}

	return localctx
}

// IMoonphaseContext is an interface to support dynamic dispatch.
type IMoonphaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoonphaseContext differentiates from other interfaces.
	IsMoonphaseContext()
}

type MoonphaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoonphaseContext() *MoonphaseContext {
	var p = new(MoonphaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MoonphaseParserRULE_moonphase
	return p
}

func (*MoonphaseContext) IsMoonphaseContext() {}

func NewMoonphaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoonphaseContext {
	var p = new(MoonphaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MoonphaseParserRULE_moonphase

	return p
}

func (s *MoonphaseContext) GetParser() antlr.Parser { return s.parser }

func (s *MoonphaseContext) MOONPHASE() antlr.TerminalNode {
	return s.GetToken(MoonphaseParserMOONPHASE, 0)
}

func (s *MoonphaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoonphaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoonphaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MoonphaseListener); ok {
		listenerT.EnterMoonphase(s)
	}
}

func (s *MoonphaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MoonphaseListener); ok {
		listenerT.ExitMoonphase(s)
	}
}

func (p *MoonphaseParser) Moonphase() (localctx IMoonphaseContext) {
	this := p
	_ = this

	localctx = NewMoonphaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MoonphaseParserRULE_moonphase)

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
		p.Match(MoonphaseParserMOONPHASE)
	}

	return localctx
}
