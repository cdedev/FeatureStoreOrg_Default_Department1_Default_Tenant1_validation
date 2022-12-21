// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603739785/Arsenic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Arsenic

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

type ArsenicParser struct {
	*antlr.BaseParser
}

var arsenicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func arsenicParserInit() {
	staticData := &arsenicParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ARSENIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "arsenic",
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

// ArsenicParserInit initializes any static state used to implement ArsenicParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewArsenicParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ArsenicParserInit() {
	staticData := &arsenicParserStaticData
	staticData.once.Do(arsenicParserInit)
}

// NewArsenicParser produces a new parser instance for the optional input antlr.TokenStream.
func NewArsenicParser(input antlr.TokenStream) *ArsenicParser {
	ArsenicParserInit()
	this := new(ArsenicParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &arsenicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Arsenic.g4"

	return this
}

// ArsenicParser tokens.
const (
	ArsenicParserEOF      = antlr.TokenEOF
	ArsenicParserARSENIC  = 1
	ArsenicParserOWNKEY   = 2
	ArsenicParserSPLITKEY = 3
	ArsenicParserWS       = 4
)

// ArsenicParser rules.
const (
	ArsenicParserRULE_expression = 0
	ArsenicParserRULE_arsenic    = 1
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
	p.RuleIndex = ArsenicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArsenicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Arsenic() IArsenicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArsenicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArsenicContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ArsenicParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArsenicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArsenicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ArsenicParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ArsenicParserRULE_expression)

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
		p.Arsenic()
	}
	{
		p.SetState(5)
		p.Match(ArsenicParserEOF)
	}

	return localctx
}

// IArsenicContext is an interface to support dynamic dispatch.
type IArsenicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArsenicContext differentiates from other interfaces.
	IsArsenicContext()
}

type ArsenicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArsenicContext() *ArsenicContext {
	var p = new(ArsenicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArsenicParserRULE_arsenic
	return p
}

func (*ArsenicContext) IsArsenicContext() {}

func NewArsenicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArsenicContext {
	var p = new(ArsenicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArsenicParserRULE_arsenic

	return p
}

func (s *ArsenicContext) GetParser() antlr.Parser { return s.parser }

func (s *ArsenicContext) ARSENIC() antlr.TerminalNode {
	return s.GetToken(ArsenicParserARSENIC, 0)
}

func (s *ArsenicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArsenicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArsenicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArsenicListener); ok {
		listenerT.EnterArsenic(s)
	}
}

func (s *ArsenicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArsenicListener); ok {
		listenerT.ExitArsenic(s)
	}
}

func (p *ArsenicParser) Arsenic() (localctx IArsenicContext) {
	this := p
	_ = this

	localctx = NewArsenicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ArsenicParserRULE_arsenic)

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
		p.Match(ArsenicParserARSENIC)
	}

	return localctx
}
