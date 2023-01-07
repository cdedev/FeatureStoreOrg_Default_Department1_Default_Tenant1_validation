// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083863888/Subregion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Subregion

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

type SubregionParser struct {
	*antlr.BaseParser
}

var subregionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func subregionParserInit() {
	staticData := &subregionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SUBREGION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "subregion",
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

// SubregionParserInit initializes any static state used to implement SubregionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSubregionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SubregionParserInit() {
	staticData := &subregionParserStaticData
	staticData.once.Do(subregionParserInit)
}

// NewSubregionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSubregionParser(input antlr.TokenStream) *SubregionParser {
	SubregionParserInit()
	this := new(SubregionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &subregionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Subregion.g4"

	return this
}

// SubregionParser tokens.
const (
	SubregionParserEOF       = antlr.TokenEOF
	SubregionParserSUBREGION = 1
	SubregionParserOWNKEY    = 2
	SubregionParserSPLITKEY  = 3
	SubregionParserWS        = 4
)

// SubregionParser rules.
const (
	SubregionParserRULE_expression = 0
	SubregionParserRULE_subregion  = 1
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
	p.RuleIndex = SubregionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubregionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Subregion() ISubregionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubregionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubregionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SubregionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubregionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubregionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SubregionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SubregionParserRULE_expression)

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
		p.Subregion()
	}
	{
		p.SetState(5)
		p.Match(SubregionParserEOF)
	}

	return localctx
}

// ISubregionContext is an interface to support dynamic dispatch.
type ISubregionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubregionContext differentiates from other interfaces.
	IsSubregionContext()
}

type SubregionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubregionContext() *SubregionContext {
	var p = new(SubregionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SubregionParserRULE_subregion
	return p
}

func (*SubregionContext) IsSubregionContext() {}

func NewSubregionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubregionContext {
	var p = new(SubregionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SubregionParserRULE_subregion

	return p
}

func (s *SubregionContext) GetParser() antlr.Parser { return s.parser }

func (s *SubregionContext) SUBREGION() antlr.TerminalNode {
	return s.GetToken(SubregionParserSUBREGION, 0)
}

func (s *SubregionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubregionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubregionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubregionListener); ok {
		listenerT.EnterSubregion(s)
	}
}

func (s *SubregionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SubregionListener); ok {
		listenerT.ExitSubregion(s)
	}
}

func (p *SubregionParser) Subregion() (localctx ISubregionContext) {
	this := p
	_ = this

	localctx = NewSubregionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SubregionParserRULE_subregion)

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
		p.Match(SubregionParserSUBREGION)
	}

	return localctx
}
