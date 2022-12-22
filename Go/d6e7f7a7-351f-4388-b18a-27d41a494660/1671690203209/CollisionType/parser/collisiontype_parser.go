// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690203209/CollisionType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CollisionType

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

type CollisionTypeParser struct {
	*antlr.BaseParser
}

var collisiontypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func collisiontypeParserInit() {
	staticData := &collisiontypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "COLLISIONTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "collisiontype",
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

// CollisionTypeParserInit initializes any static state used to implement CollisionTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCollisionTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CollisionTypeParserInit() {
	staticData := &collisiontypeParserStaticData
	staticData.once.Do(collisiontypeParserInit)
}

// NewCollisionTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCollisionTypeParser(input antlr.TokenStream) *CollisionTypeParser {
	CollisionTypeParserInit()
	this := new(CollisionTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &collisiontypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "CollisionType.g4"

	return this
}

// CollisionTypeParser tokens.
const (
	CollisionTypeParserEOF           = antlr.TokenEOF
	CollisionTypeParserCOLLISIONTYPE = 1
	CollisionTypeParserOWNKEY        = 2
	CollisionTypeParserSPLITKEY      = 3
	CollisionTypeParserWS            = 4
)

// CollisionTypeParser rules.
const (
	CollisionTypeParserRULE_expression    = 0
	CollisionTypeParserRULE_collisiontype = 1
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
	p.RuleIndex = CollisionTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CollisionTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Collisiontype() ICollisiontypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollisiontypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollisiontypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CollisionTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CollisionTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CollisionTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CollisionTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CollisionTypeParserRULE_expression)

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
		p.Collisiontype()
	}
	{
		p.SetState(5)
		p.Match(CollisionTypeParserEOF)
	}

	return localctx
}

// ICollisiontypeContext is an interface to support dynamic dispatch.
type ICollisiontypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollisiontypeContext differentiates from other interfaces.
	IsCollisiontypeContext()
}

type CollisiontypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollisiontypeContext() *CollisiontypeContext {
	var p = new(CollisiontypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CollisionTypeParserRULE_collisiontype
	return p
}

func (*CollisiontypeContext) IsCollisiontypeContext() {}

func NewCollisiontypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollisiontypeContext {
	var p = new(CollisiontypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CollisionTypeParserRULE_collisiontype

	return p
}

func (s *CollisiontypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CollisiontypeContext) COLLISIONTYPE() antlr.TerminalNode {
	return s.GetToken(CollisionTypeParserCOLLISIONTYPE, 0)
}

func (s *CollisiontypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollisiontypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollisiontypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CollisionTypeListener); ok {
		listenerT.EnterCollisiontype(s)
	}
}

func (s *CollisiontypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CollisionTypeListener); ok {
		listenerT.ExitCollisiontype(s)
	}
}

func (p *CollisionTypeParser) Collisiontype() (localctx ICollisiontypeContext) {
	this := p
	_ = this

	localctx = NewCollisiontypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CollisionTypeParserRULE_collisiontype)

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
		p.Match(CollisionTypeParserCOLLISIONTYPE)
	}

	return localctx
}
