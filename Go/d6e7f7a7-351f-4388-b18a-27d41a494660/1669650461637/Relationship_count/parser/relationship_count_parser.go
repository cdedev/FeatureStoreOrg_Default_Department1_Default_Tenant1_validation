// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669650461637/Relationship_count.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Relationship_count

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

type Relationship_countParser struct {
	*antlr.BaseParser
}

var relationship_countParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func relationship_countParserInit() {
	staticData := &relationship_countParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RELATIONSHIP_COUNT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "relationship_count",
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

// Relationship_countParserInit initializes any static state used to implement Relationship_countParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRelationship_countParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Relationship_countParserInit() {
	staticData := &relationship_countParserStaticData
	staticData.once.Do(relationship_countParserInit)
}

// NewRelationship_countParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRelationship_countParser(input antlr.TokenStream) *Relationship_countParser {
	Relationship_countParserInit()
	this := new(Relationship_countParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &relationship_countParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Relationship_count.g4"

	return this
}

// Relationship_countParser tokens.
const (
	Relationship_countParserEOF                = antlr.TokenEOF
	Relationship_countParserRELATIONSHIP_COUNT = 1
	Relationship_countParserOWNKEY             = 2
	Relationship_countParserSPLITKEY           = 3
	Relationship_countParserWS                 = 4
)

// Relationship_countParser rules.
const (
	Relationship_countParserRULE_expression         = 0
	Relationship_countParserRULE_relationship_count = 1
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
	p.RuleIndex = Relationship_countParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Relationship_countParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Relationship_count() IRelationship_countContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationship_countContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationship_countContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Relationship_countParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Relationship_countListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Relationship_countListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Relationship_countParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Relationship_countParserRULE_expression)

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
		p.Relationship_count()
	}
	{
		p.SetState(5)
		p.Match(Relationship_countParserEOF)
	}

	return localctx
}

// IRelationship_countContext is an interface to support dynamic dispatch.
type IRelationship_countContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationship_countContext differentiates from other interfaces.
	IsRelationship_countContext()
}

type Relationship_countContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationship_countContext() *Relationship_countContext {
	var p = new(Relationship_countContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Relationship_countParserRULE_relationship_count
	return p
}

func (*Relationship_countContext) IsRelationship_countContext() {}

func NewRelationship_countContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relationship_countContext {
	var p = new(Relationship_countContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Relationship_countParserRULE_relationship_count

	return p
}

func (s *Relationship_countContext) GetParser() antlr.Parser { return s.parser }

func (s *Relationship_countContext) RELATIONSHIP_COUNT() antlr.TerminalNode {
	return s.GetToken(Relationship_countParserRELATIONSHIP_COUNT, 0)
}

func (s *Relationship_countContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relationship_countContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relationship_countContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Relationship_countListener); ok {
		listenerT.EnterRelationship_count(s)
	}
}

func (s *Relationship_countContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Relationship_countListener); ok {
		listenerT.ExitRelationship_count(s)
	}
}

func (p *Relationship_countParser) Relationship_count() (localctx IRelationship_countContext) {
	this := p
	_ = this

	localctx = NewRelationship_countContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Relationship_countParserRULE_relationship_count)

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
		p.Match(Relationship_countParserRELATIONSHIP_COUNT)
	}

	return localctx
}
