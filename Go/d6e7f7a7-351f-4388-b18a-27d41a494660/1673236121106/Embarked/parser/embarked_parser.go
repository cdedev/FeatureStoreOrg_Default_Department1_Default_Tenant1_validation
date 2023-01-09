// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236121106/Embarked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Embarked

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

type EmbarkedParser struct {
	*antlr.BaseParser
}

var embarkedParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func embarkedParserInit() {
	staticData := &embarkedParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EMBARKED", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "embarked",
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

// EmbarkedParserInit initializes any static state used to implement EmbarkedParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEmbarkedParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EmbarkedParserInit() {
	staticData := &embarkedParserStaticData
	staticData.once.Do(embarkedParserInit)
}

// NewEmbarkedParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEmbarkedParser(input antlr.TokenStream) *EmbarkedParser {
	EmbarkedParserInit()
	this := new(EmbarkedParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &embarkedParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Embarked.g4"

	return this
}

// EmbarkedParser tokens.
const (
	EmbarkedParserEOF      = antlr.TokenEOF
	EmbarkedParserEMBARKED = 1
	EmbarkedParserOWNKEY   = 2
	EmbarkedParserSPLITKEY = 3
	EmbarkedParserWS       = 4
)

// EmbarkedParser rules.
const (
	EmbarkedParserRULE_expression = 0
	EmbarkedParserRULE_embarked   = 1
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
	p.RuleIndex = EmbarkedParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmbarkedParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Embarked() IEmbarkedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmbarkedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmbarkedContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EmbarkedParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmbarkedListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmbarkedListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EmbarkedParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EmbarkedParserRULE_expression)

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
		p.Embarked()
	}
	{
		p.SetState(5)
		p.Match(EmbarkedParserEOF)
	}

	return localctx
}

// IEmbarkedContext is an interface to support dynamic dispatch.
type IEmbarkedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmbarkedContext differentiates from other interfaces.
	IsEmbarkedContext()
}

type EmbarkedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmbarkedContext() *EmbarkedContext {
	var p = new(EmbarkedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EmbarkedParserRULE_embarked
	return p
}

func (*EmbarkedContext) IsEmbarkedContext() {}

func NewEmbarkedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmbarkedContext {
	var p = new(EmbarkedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmbarkedParserRULE_embarked

	return p
}

func (s *EmbarkedContext) GetParser() antlr.Parser { return s.parser }

func (s *EmbarkedContext) EMBARKED() antlr.TerminalNode {
	return s.GetToken(EmbarkedParserEMBARKED, 0)
}

func (s *EmbarkedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmbarkedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmbarkedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmbarkedListener); ok {
		listenerT.EnterEmbarked(s)
	}
}

func (s *EmbarkedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmbarkedListener); ok {
		listenerT.ExitEmbarked(s)
	}
}

func (p *EmbarkedParser) Embarked() (localctx IEmbarkedContext) {
	this := p
	_ = this

	localctx = NewEmbarkedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EmbarkedParserRULE_embarked)

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
		p.Match(EmbarkedParserEMBARKED)
	}

	return localctx
}
