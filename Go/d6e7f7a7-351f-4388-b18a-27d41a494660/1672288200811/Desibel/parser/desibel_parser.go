// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288200811/Desibel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Desibel

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

type DesibelParser struct {
	*antlr.BaseParser
}

var desibelParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func desibelParserInit() {
	staticData := &desibelParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DESIBEL", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "desibel",
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

// DesibelParserInit initializes any static state used to implement DesibelParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDesibelParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DesibelParserInit() {
	staticData := &desibelParserStaticData
	staticData.once.Do(desibelParserInit)
}

// NewDesibelParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDesibelParser(input antlr.TokenStream) *DesibelParser {
	DesibelParserInit()
	this := new(DesibelParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &desibelParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Desibel.g4"

	return this
}

// DesibelParser tokens.
const (
	DesibelParserEOF      = antlr.TokenEOF
	DesibelParserDESIBEL  = 1
	DesibelParserOWNKEY   = 2
	DesibelParserSPLITKEY = 3
	DesibelParserWS       = 4
)

// DesibelParser rules.
const (
	DesibelParserRULE_expression = 0
	DesibelParserRULE_desibel    = 1
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
	p.RuleIndex = DesibelParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesibelParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Desibel() IDesibelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDesibelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDesibelContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DesibelParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesibelListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesibelListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DesibelParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DesibelParserRULE_expression)

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
		p.Desibel()
	}
	{
		p.SetState(5)
		p.Match(DesibelParserEOF)
	}

	return localctx
}

// IDesibelContext is an interface to support dynamic dispatch.
type IDesibelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDesibelContext differentiates from other interfaces.
	IsDesibelContext()
}

type DesibelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDesibelContext() *DesibelContext {
	var p = new(DesibelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DesibelParserRULE_desibel
	return p
}

func (*DesibelContext) IsDesibelContext() {}

func NewDesibelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DesibelContext {
	var p = new(DesibelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DesibelParserRULE_desibel

	return p
}

func (s *DesibelContext) GetParser() antlr.Parser { return s.parser }

func (s *DesibelContext) DESIBEL() antlr.TerminalNode {
	return s.GetToken(DesibelParserDESIBEL, 0)
}

func (s *DesibelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DesibelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DesibelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesibelListener); ok {
		listenerT.EnterDesibel(s)
	}
}

func (s *DesibelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DesibelListener); ok {
		listenerT.ExitDesibel(s)
	}
}

func (p *DesibelParser) Desibel() (localctx IDesibelContext) {
	this := p
	_ = this

	localctx = NewDesibelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DesibelParserRULE_desibel)

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
		p.Match(DesibelParserDESIBEL)
	}

	return localctx
}
