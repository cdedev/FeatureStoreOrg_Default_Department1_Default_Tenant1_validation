// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979225806/Epoch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Epoch

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

type EpochParser struct {
	*antlr.BaseParser
}

var epochParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func epochParserInit() {
	staticData := &epochParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EPOCH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "epoch",
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

// EpochParserInit initializes any static state used to implement EpochParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEpochParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EpochParserInit() {
	staticData := &epochParserStaticData
	staticData.once.Do(epochParserInit)
}

// NewEpochParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEpochParser(input antlr.TokenStream) *EpochParser {
	EpochParserInit()
	this := new(EpochParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &epochParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Epoch.g4"

	return this
}

// EpochParser tokens.
const (
	EpochParserEOF      = antlr.TokenEOF
	EpochParserEPOCH    = 1
	EpochParserOWNKEY   = 2
	EpochParserSPLITKEY = 3
	EpochParserWS       = 4
)

// EpochParser rules.
const (
	EpochParserRULE_expression = 0
	EpochParserRULE_epoch      = 1
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
	p.RuleIndex = EpochParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EpochParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Epoch() IEpochContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEpochContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEpochContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EpochParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpochListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpochListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EpochParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EpochParserRULE_expression)

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
		p.Epoch()
	}
	{
		p.SetState(5)
		p.Match(EpochParserEOF)
	}

	return localctx
}

// IEpochContext is an interface to support dynamic dispatch.
type IEpochContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEpochContext differentiates from other interfaces.
	IsEpochContext()
}

type EpochContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpochContext() *EpochContext {
	var p = new(EpochContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EpochParserRULE_epoch
	return p
}

func (*EpochContext) IsEpochContext() {}

func NewEpochContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpochContext {
	var p = new(EpochContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EpochParserRULE_epoch

	return p
}

func (s *EpochContext) GetParser() antlr.Parser { return s.parser }

func (s *EpochContext) EPOCH() antlr.TerminalNode {
	return s.GetToken(EpochParserEPOCH, 0)
}

func (s *EpochContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpochContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EpochContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpochListener); ok {
		listenerT.EnterEpoch(s)
	}
}

func (s *EpochContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EpochListener); ok {
		listenerT.ExitEpoch(s)
	}
}

func (p *EpochParser) Epoch() (localctx IEpochContext) {
	this := p
	_ = this

	localctx = NewEpochContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EpochParserRULE_epoch)

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
		p.Match(EpochParserEPOCH)
	}

	return localctx
}
