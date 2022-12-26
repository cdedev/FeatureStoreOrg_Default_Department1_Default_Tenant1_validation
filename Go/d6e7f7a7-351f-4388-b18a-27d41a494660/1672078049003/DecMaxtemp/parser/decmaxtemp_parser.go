// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078049003/DecMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DecMaxtemp

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

type DecMaxtempParser struct {
	*antlr.BaseParser
}

var decmaxtempParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func decmaxtempParserInit() {
	staticData := &decmaxtempParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "DECMAXTEMP", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "decmaxtemp",
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

// DecMaxtempParserInit initializes any static state used to implement DecMaxtempParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDecMaxtempParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DecMaxtempParserInit() {
	staticData := &decmaxtempParserStaticData
	staticData.once.Do(decmaxtempParserInit)
}

// NewDecMaxtempParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDecMaxtempParser(input antlr.TokenStream) *DecMaxtempParser {
	DecMaxtempParserInit()
	this := new(DecMaxtempParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &decmaxtempParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "DecMaxtemp.g4"

	return this
}

// DecMaxtempParser tokens.
const (
	DecMaxtempParserEOF        = antlr.TokenEOF
	DecMaxtempParserDECMAXTEMP = 1
	DecMaxtempParserOWNKEY     = 2
	DecMaxtempParserSPLITKEY   = 3
	DecMaxtempParserWS         = 4
)

// DecMaxtempParser rules.
const (
	DecMaxtempParserRULE_expression = 0
	DecMaxtempParserRULE_decmaxtemp = 1
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
	p.RuleIndex = DecMaxtempParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DecMaxtempParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Decmaxtemp() IDecmaxtempContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecmaxtempContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecmaxtempContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(DecMaxtempParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMaxtempListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMaxtempListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DecMaxtempParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DecMaxtempParserRULE_expression)

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
		p.Decmaxtemp()
	}
	{
		p.SetState(5)
		p.Match(DecMaxtempParserEOF)
	}

	return localctx
}

// IDecmaxtempContext is an interface to support dynamic dispatch.
type IDecmaxtempContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecmaxtempContext differentiates from other interfaces.
	IsDecmaxtempContext()
}

type DecmaxtempContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecmaxtempContext() *DecmaxtempContext {
	var p = new(DecmaxtempContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DecMaxtempParserRULE_decmaxtemp
	return p
}

func (*DecmaxtempContext) IsDecmaxtempContext() {}

func NewDecmaxtempContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecmaxtempContext {
	var p = new(DecmaxtempContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DecMaxtempParserRULE_decmaxtemp

	return p
}

func (s *DecmaxtempContext) GetParser() antlr.Parser { return s.parser }

func (s *DecmaxtempContext) DECMAXTEMP() antlr.TerminalNode {
	return s.GetToken(DecMaxtempParserDECMAXTEMP, 0)
}

func (s *DecmaxtempContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecmaxtempContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecmaxtempContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMaxtempListener); ok {
		listenerT.EnterDecmaxtemp(s)
	}
}

func (s *DecmaxtempContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DecMaxtempListener); ok {
		listenerT.ExitDecmaxtemp(s)
	}
}

func (p *DecMaxtempParser) Decmaxtemp() (localctx IDecmaxtempContext) {
	this := p
	_ = this

	localctx = NewDecmaxtempContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DecMaxtempParserRULE_decmaxtemp)

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
		p.Match(DecMaxtempParserDECMAXTEMP)
	}

	return localctx
}
