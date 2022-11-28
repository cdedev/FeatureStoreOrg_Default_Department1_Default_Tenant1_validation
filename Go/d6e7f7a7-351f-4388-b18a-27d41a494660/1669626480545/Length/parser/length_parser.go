// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626480545/Length.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Length

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

type LengthParser struct {
	*antlr.BaseParser
}

var lengthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func lengthParserInit() {
	staticData := &lengthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "LENGTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "length",
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

// LengthParserInit initializes any static state used to implement LengthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLengthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LengthParserInit() {
	staticData := &lengthParserStaticData
	staticData.once.Do(lengthParserInit)
}

// NewLengthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLengthParser(input antlr.TokenStream) *LengthParser {
	LengthParserInit()
	this := new(LengthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &lengthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Length.g4"

	return this
}

// LengthParser tokens.
const (
	LengthParserEOF      = antlr.TokenEOF
	LengthParserLENGTH   = 1
	LengthParserOWNKEY   = 2
	LengthParserSPLITKEY = 3
	LengthParserWS       = 4
)

// LengthParser rules.
const (
	LengthParserRULE_expression = 0
	LengthParserRULE_length     = 1
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
	p.RuleIndex = LengthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LengthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Length() ILengthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(LengthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LengthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LengthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LengthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LengthParserRULE_expression)

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
		p.Length()
	}
	{
		p.SetState(5)
		p.Match(LengthParserEOF)
	}

	return localctx
}

// ILengthContext is an interface to support dynamic dispatch.
type ILengthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLengthContext differentiates from other interfaces.
	IsLengthContext()
}

type LengthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthContext() *LengthContext {
	var p = new(LengthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LengthParserRULE_length
	return p
}

func (*LengthContext) IsLengthContext() {}

func NewLengthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthContext {
	var p = new(LengthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LengthParserRULE_length

	return p
}

func (s *LengthContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthContext) LENGTH() antlr.TerminalNode {
	return s.GetToken(LengthParserLENGTH, 0)
}

func (s *LengthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LengthListener); ok {
		listenerT.EnterLength(s)
	}
}

func (s *LengthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LengthListener); ok {
		listenerT.ExitLength(s)
	}
}

func (p *LengthParser) Length() (localctx ILengthContext) {
	this := p
	_ = this

	localctx = NewLengthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LengthParserRULE_length)

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
		p.Match(LengthParserLENGTH)
	}

	return localctx
}
