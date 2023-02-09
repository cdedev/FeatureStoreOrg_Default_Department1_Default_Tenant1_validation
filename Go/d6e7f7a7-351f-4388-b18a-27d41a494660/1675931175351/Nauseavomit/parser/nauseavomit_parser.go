// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931175351/Nauseavomit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nauseavomit

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

type NauseavomitParser struct {
	*antlr.BaseParser
}

var nauseavomitParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func nauseavomitParserInit() {
	staticData := &nauseavomitParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "NAUSEAVOMIT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "nauseavomit",
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

// NauseavomitParserInit initializes any static state used to implement NauseavomitParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewNauseavomitParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func NauseavomitParserInit() {
	staticData := &nauseavomitParserStaticData
	staticData.once.Do(nauseavomitParserInit)
}

// NewNauseavomitParser produces a new parser instance for the optional input antlr.TokenStream.
func NewNauseavomitParser(input antlr.TokenStream) *NauseavomitParser {
	NauseavomitParserInit()
	this := new(NauseavomitParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &nauseavomitParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Nauseavomit.g4"

	return this
}

// NauseavomitParser tokens.
const (
	NauseavomitParserEOF         = antlr.TokenEOF
	NauseavomitParserNAUSEAVOMIT = 1
	NauseavomitParserOWNKEY      = 2
	NauseavomitParserSPLITKEY    = 3
	NauseavomitParserWS          = 4
)

// NauseavomitParser rules.
const (
	NauseavomitParserRULE_expression  = 0
	NauseavomitParserRULE_nauseavomit = 1
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
	p.RuleIndex = NauseavomitParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NauseavomitParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Nauseavomit() INauseavomitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INauseavomitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INauseavomitContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(NauseavomitParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseavomitListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseavomitListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *NauseavomitParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NauseavomitParserRULE_expression)

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
		p.Nauseavomit()
	}
	{
		p.SetState(5)
		p.Match(NauseavomitParserEOF)
	}

	return localctx
}

// INauseavomitContext is an interface to support dynamic dispatch.
type INauseavomitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNauseavomitContext differentiates from other interfaces.
	IsNauseavomitContext()
}

type NauseavomitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNauseavomitContext() *NauseavomitContext {
	var p = new(NauseavomitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NauseavomitParserRULE_nauseavomit
	return p
}

func (*NauseavomitContext) IsNauseavomitContext() {}

func NewNauseavomitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NauseavomitContext {
	var p = new(NauseavomitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NauseavomitParserRULE_nauseavomit

	return p
}

func (s *NauseavomitContext) GetParser() antlr.Parser { return s.parser }

func (s *NauseavomitContext) NAUSEAVOMIT() antlr.TerminalNode {
	return s.GetToken(NauseavomitParserNAUSEAVOMIT, 0)
}

func (s *NauseavomitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NauseavomitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NauseavomitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseavomitListener); ok {
		listenerT.EnterNauseavomit(s)
	}
}

func (s *NauseavomitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NauseavomitListener); ok {
		listenerT.ExitNauseavomit(s)
	}
}

func (p *NauseavomitParser) Nauseavomit() (localctx INauseavomitContext) {
	this := p
	_ = this

	localctx = NewNauseavomitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NauseavomitParserRULE_nauseavomit)

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
		p.Match(NauseavomitParserNAUSEAVOMIT)
	}

	return localctx
}
