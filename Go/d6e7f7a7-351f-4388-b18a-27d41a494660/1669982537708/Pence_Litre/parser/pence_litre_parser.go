// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669982537708/Pence_Litre.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pence_Litre

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

type Pence_LitreParser struct {
	*antlr.BaseParser
}

var pence_litreParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pence_litreParserInit() {
	staticData := &pence_litreParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PENCE_LITRE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "pence_litre",
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

// Pence_LitreParserInit initializes any static state used to implement Pence_LitreParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPence_LitreParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Pence_LitreParserInit() {
	staticData := &pence_litreParserStaticData
	staticData.once.Do(pence_litreParserInit)
}

// NewPence_LitreParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPence_LitreParser(input antlr.TokenStream) *Pence_LitreParser {
	Pence_LitreParserInit()
	this := new(Pence_LitreParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &pence_litreParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Pence_Litre.g4"

	return this
}

// Pence_LitreParser tokens.
const (
	Pence_LitreParserEOF         = antlr.TokenEOF
	Pence_LitreParserPENCE_LITRE = 1
	Pence_LitreParserOWNKEY      = 2
	Pence_LitreParserSPLITKEY    = 3
	Pence_LitreParserWS          = 4
)

// Pence_LitreParser rules.
const (
	Pence_LitreParserRULE_expression  = 0
	Pence_LitreParserRULE_pence_litre = 1
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
	p.RuleIndex = Pence_LitreParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Pence_LitreParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Pence_litre() IPence_litreContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPence_litreContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPence_litreContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Pence_LitreParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pence_LitreListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pence_LitreListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Pence_LitreParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Pence_LitreParserRULE_expression)

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
		p.Pence_litre()
	}
	{
		p.SetState(5)
		p.Match(Pence_LitreParserEOF)
	}

	return localctx
}

// IPence_litreContext is an interface to support dynamic dispatch.
type IPence_litreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPence_litreContext differentiates from other interfaces.
	IsPence_litreContext()
}

type Pence_litreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPence_litreContext() *Pence_litreContext {
	var p = new(Pence_litreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Pence_LitreParserRULE_pence_litre
	return p
}

func (*Pence_litreContext) IsPence_litreContext() {}

func NewPence_litreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pence_litreContext {
	var p = new(Pence_litreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Pence_LitreParserRULE_pence_litre

	return p
}

func (s *Pence_litreContext) GetParser() antlr.Parser { return s.parser }

func (s *Pence_litreContext) PENCE_LITRE() antlr.TerminalNode {
	return s.GetToken(Pence_LitreParserPENCE_LITRE, 0)
}

func (s *Pence_litreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pence_litreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pence_litreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pence_LitreListener); ok {
		listenerT.EnterPence_litre(s)
	}
}

func (s *Pence_litreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Pence_LitreListener); ok {
		listenerT.ExitPence_litre(s)
	}
}

func (p *Pence_LitreParser) Pence_litre() (localctx IPence_litreContext) {
	this := p
	_ = this

	localctx = NewPence_litreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Pence_LitreParserRULE_pence_litre)

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
		p.Match(Pence_LitreParserPENCE_LITRE)
	}

	return localctx
}
