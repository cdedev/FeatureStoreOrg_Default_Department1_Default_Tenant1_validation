// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675937987570/Bandwidth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bandwidth

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

type BandwidthParser struct {
	*antlr.BaseParser
}

var bandwidthParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bandwidthParserInit() {
	staticData := &bandwidthParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BANDWIDTH", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bandwidth",
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

// BandwidthParserInit initializes any static state used to implement BandwidthParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBandwidthParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BandwidthParserInit() {
	staticData := &bandwidthParserStaticData
	staticData.once.Do(bandwidthParserInit)
}

// NewBandwidthParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBandwidthParser(input antlr.TokenStream) *BandwidthParser {
	BandwidthParserInit()
	this := new(BandwidthParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bandwidthParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bandwidth.g4"

	return this
}

// BandwidthParser tokens.
const (
	BandwidthParserEOF       = antlr.TokenEOF
	BandwidthParserBANDWIDTH = 1
	BandwidthParserOWNKEY    = 2
	BandwidthParserSPLITKEY  = 3
	BandwidthParserWS        = 4
)

// BandwidthParser rules.
const (
	BandwidthParserRULE_expression = 0
	BandwidthParserRULE_bandwidth  = 1
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
	p.RuleIndex = BandwidthParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BandwidthParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bandwidth() IBandwidthContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBandwidthContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBandwidthContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BandwidthParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BandwidthListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BandwidthListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BandwidthParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BandwidthParserRULE_expression)

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
		p.Bandwidth()
	}
	{
		p.SetState(5)
		p.Match(BandwidthParserEOF)
	}

	return localctx
}

// IBandwidthContext is an interface to support dynamic dispatch.
type IBandwidthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBandwidthContext differentiates from other interfaces.
	IsBandwidthContext()
}

type BandwidthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBandwidthContext() *BandwidthContext {
	var p = new(BandwidthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BandwidthParserRULE_bandwidth
	return p
}

func (*BandwidthContext) IsBandwidthContext() {}

func NewBandwidthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BandwidthContext {
	var p = new(BandwidthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BandwidthParserRULE_bandwidth

	return p
}

func (s *BandwidthContext) GetParser() antlr.Parser { return s.parser }

func (s *BandwidthContext) BANDWIDTH() antlr.TerminalNode {
	return s.GetToken(BandwidthParserBANDWIDTH, 0)
}

func (s *BandwidthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BandwidthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BandwidthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BandwidthListener); ok {
		listenerT.EnterBandwidth(s)
	}
}

func (s *BandwidthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BandwidthListener); ok {
		listenerT.ExitBandwidth(s)
	}
}

func (p *BandwidthParser) Bandwidth() (localctx IBandwidthContext) {
	this := p
	_ = this

	localctx = NewBandwidthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BandwidthParserRULE_bandwidth)

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
		p.Match(BandwidthParserBANDWIDTH)
	}

	return localctx
}
