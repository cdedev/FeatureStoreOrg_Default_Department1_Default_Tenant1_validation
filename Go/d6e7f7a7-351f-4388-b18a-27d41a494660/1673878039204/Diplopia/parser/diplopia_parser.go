// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878039204/Diplopia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diplopia

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

type DiplopiaParser struct {
	*antlr.BaseParser
}

var diplopiaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func diplopiaParserInit() {
	staticData := &diplopiaParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DIPLOPIA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "diplopia",
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

// DiplopiaParserInit initializes any static state used to implement DiplopiaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDiplopiaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DiplopiaParserInit() {
	staticData := &diplopiaParserStaticData
	staticData.once.Do(diplopiaParserInit)
}

// NewDiplopiaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDiplopiaParser(input antlr.TokenStream) *DiplopiaParser {
	DiplopiaParserInit()
	this := new(DiplopiaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &diplopiaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Diplopia.g4"

	return this
}

// DiplopiaParser tokens.
const (
	DiplopiaParserEOF      = antlr.TokenEOF
	DiplopiaParserDIPLOPIA = 1
	DiplopiaParserOWNKEY   = 2
	DiplopiaParserSPLITKEY = 3
	DiplopiaParserWS       = 4
)

// DiplopiaParser rules.
const (
	DiplopiaParserRULE_expression = 0
	DiplopiaParserRULE_diplopia   = 1
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
	p.RuleIndex = DiplopiaParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiplopiaParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Diplopia() IDiplopiaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDiplopiaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDiplopiaContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DiplopiaParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiplopiaListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiplopiaListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DiplopiaParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DiplopiaParserRULE_expression)

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
		p.Diplopia()
	}
	{
		p.SetState(5)
		p.Match(DiplopiaParserEOF)
	}

	return localctx
}

// IDiplopiaContext is an interface to support dynamic dispatch.
type IDiplopiaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiplopiaContext differentiates from other interfaces.
	IsDiplopiaContext()
}

type DiplopiaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiplopiaContext() *DiplopiaContext {
	var p = new(DiplopiaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DiplopiaParserRULE_diplopia
	return p
}

func (*DiplopiaContext) IsDiplopiaContext() {}

func NewDiplopiaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiplopiaContext {
	var p = new(DiplopiaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DiplopiaParserRULE_diplopia

	return p
}

func (s *DiplopiaContext) GetParser() antlr.Parser { return s.parser }

func (s *DiplopiaContext) DIPLOPIA() antlr.TerminalNode {
	return s.GetToken(DiplopiaParserDIPLOPIA, 0)
}

func (s *DiplopiaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiplopiaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiplopiaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiplopiaListener); ok {
		listenerT.EnterDiplopia(s)
	}
}

func (s *DiplopiaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DiplopiaListener); ok {
		listenerT.ExitDiplopia(s)
	}
}

func (p *DiplopiaParser) Diplopia() (localctx IDiplopiaContext) {
	this := p
	_ = this

	localctx = NewDiplopiaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DiplopiaParserRULE_diplopia)

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
		p.Match(DiplopiaParserDIPLOPIA)
	}

	return localctx
}
