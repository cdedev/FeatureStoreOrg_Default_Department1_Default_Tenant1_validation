// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672373090178/BPStable.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BPStable

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

type BPStableParser struct {
	*antlr.BaseParser
}

var bpstableParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bpstableParserInit() {
	staticData := &bpstableParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BPSTABLE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bpstable",
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

// BPStableParserInit initializes any static state used to implement BPStableParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBPStableParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BPStableParserInit() {
	staticData := &bpstableParserStaticData
	staticData.once.Do(bpstableParserInit)
}

// NewBPStableParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBPStableParser(input antlr.TokenStream) *BPStableParser {
	BPStableParserInit()
	this := new(BPStableParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bpstableParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BPStable.g4"

	return this
}

// BPStableParser tokens.
const (
	BPStableParserEOF      = antlr.TokenEOF
	BPStableParserBPSTABLE = 1
	BPStableParserOWNKEY   = 2
	BPStableParserSPLITKEY = 3
	BPStableParserWS       = 4
)

// BPStableParser rules.
const (
	BPStableParserRULE_expression = 0
	BPStableParserRULE_bpstable   = 1
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
	p.RuleIndex = BPStableParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BPStableParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bpstable() IBpstableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBpstableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBpstableContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BPStableParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BPStableListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BPStableListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BPStableParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BPStableParserRULE_expression)

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
		p.Bpstable()
	}
	{
		p.SetState(5)
		p.Match(BPStableParserEOF)
	}

	return localctx
}

// IBpstableContext is an interface to support dynamic dispatch.
type IBpstableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBpstableContext differentiates from other interfaces.
	IsBpstableContext()
}

type BpstableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBpstableContext() *BpstableContext {
	var p = new(BpstableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BPStableParserRULE_bpstable
	return p
}

func (*BpstableContext) IsBpstableContext() {}

func NewBpstableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BpstableContext {
	var p = new(BpstableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BPStableParserRULE_bpstable

	return p
}

func (s *BpstableContext) GetParser() antlr.Parser { return s.parser }

func (s *BpstableContext) BPSTABLE() antlr.TerminalNode {
	return s.GetToken(BPStableParserBPSTABLE, 0)
}

func (s *BpstableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BpstableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BpstableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BPStableListener); ok {
		listenerT.EnterBpstable(s)
	}
}

func (s *BpstableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BPStableListener); ok {
		listenerT.ExitBpstable(s)
	}
}

func (p *BPStableParser) Bpstable() (localctx IBpstableContext) {
	this := p
	_ = this

	localctx = NewBpstableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BPStableParserRULE_bpstable)

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
		p.Match(BPStableParserBPSTABLE)
	}

	return localctx
}
