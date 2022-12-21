// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671599159281/Entity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entity

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

type EntityParser struct {
	*antlr.BaseParser
}

var entityParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func entityParserInit() {
	staticData := &entityParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ENTITY", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "entity",
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

// EntityParserInit initializes any static state used to implement EntityParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEntityParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EntityParserInit() {
	staticData := &entityParserStaticData
	staticData.once.Do(entityParserInit)
}

// NewEntityParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEntityParser(input antlr.TokenStream) *EntityParser {
	EntityParserInit()
	this := new(EntityParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &entityParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Entity.g4"

	return this
}

// EntityParser tokens.
const (
	EntityParserEOF      = antlr.TokenEOF
	EntityParserENTITY   = 1
	EntityParserOWNKEY   = 2
	EntityParserSPLITKEY = 3
	EntityParserWS       = 4
)

// EntityParser rules.
const (
	EntityParserRULE_expression = 0
	EntityParserRULE_entity     = 1
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
	p.RuleIndex = EntityParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EntityParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Entity() IEntityContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEntityContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEntityContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EntityParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntityListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntityListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EntityParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EntityParserRULE_expression)

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
		p.Entity()
	}
	{
		p.SetState(5)
		p.Match(EntityParserEOF)
	}

	return localctx
}

// IEntityContext is an interface to support dynamic dispatch.
type IEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityContext differentiates from other interfaces.
	IsEntityContext()
}

type EntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityContext() *EntityContext {
	var p = new(EntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EntityParserRULE_entity
	return p
}

func (*EntityContext) IsEntityContext() {}

func NewEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityContext {
	var p = new(EntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EntityParserRULE_entity

	return p
}

func (s *EntityContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(EntityParserENTITY, 0)
}

func (s *EntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntityListener); ok {
		listenerT.EnterEntity(s)
	}
}

func (s *EntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EntityListener); ok {
		listenerT.ExitEntity(s)
	}
}

func (p *EntityParser) Entity() (localctx IEntityContext) {
	this := p
	_ = this

	localctx = NewEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EntityParserRULE_entity)

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
		p.Match(EntityParserENTITY)
	}

	return localctx
}
