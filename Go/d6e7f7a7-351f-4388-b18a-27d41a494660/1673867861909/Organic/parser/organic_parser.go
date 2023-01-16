// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867861909/Organic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Organic

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

type OrganicParser struct {
	*antlr.BaseParser
}

var organicParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func organicParserInit() {
	staticData := &organicParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ORGANIC", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "organic",
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

// OrganicParserInit initializes any static state used to implement OrganicParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOrganicParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OrganicParserInit() {
	staticData := &organicParserStaticData
	staticData.once.Do(organicParserInit)
}

// NewOrganicParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOrganicParser(input antlr.TokenStream) *OrganicParser {
	OrganicParserInit()
	this := new(OrganicParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &organicParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Organic.g4"

	return this
}

// OrganicParser tokens.
const (
	OrganicParserEOF      = antlr.TokenEOF
	OrganicParserORGANIC  = 1
	OrganicParserOWNKEY   = 2
	OrganicParserSPLITKEY = 3
	OrganicParserWS       = 4
)

// OrganicParser rules.
const (
	OrganicParserRULE_expression = 0
	OrganicParserRULE_organic    = 1
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
	p.RuleIndex = OrganicParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OrganicParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Organic() IOrganicContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrganicContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrganicContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(OrganicParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrganicListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrganicListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *OrganicParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OrganicParserRULE_expression)

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
		p.Organic()
	}
	{
		p.SetState(5)
		p.Match(OrganicParserEOF)
	}

	return localctx
}

// IOrganicContext is an interface to support dynamic dispatch.
type IOrganicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrganicContext differentiates from other interfaces.
	IsOrganicContext()
}

type OrganicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrganicContext() *OrganicContext {
	var p = new(OrganicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OrganicParserRULE_organic
	return p
}

func (*OrganicContext) IsOrganicContext() {}

func NewOrganicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrganicContext {
	var p = new(OrganicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OrganicParserRULE_organic

	return p
}

func (s *OrganicContext) GetParser() antlr.Parser { return s.parser }

func (s *OrganicContext) ORGANIC() antlr.TerminalNode {
	return s.GetToken(OrganicParserORGANIC, 0)
}

func (s *OrganicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrganicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrganicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrganicListener); ok {
		listenerT.EnterOrganic(s)
	}
}

func (s *OrganicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OrganicListener); ok {
		listenerT.ExitOrganic(s)
	}
}

func (p *OrganicParser) Organic() (localctx IOrganicContext) {
	this := p
	_ = this

	localctx = NewOrganicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OrganicParserRULE_organic)

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
		p.Match(OrganicParserORGANIC)
	}

	return localctx
}
