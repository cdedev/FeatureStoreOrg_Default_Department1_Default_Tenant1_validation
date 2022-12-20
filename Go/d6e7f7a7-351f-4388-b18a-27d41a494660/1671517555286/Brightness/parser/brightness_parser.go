// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517555286/Brightness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Brightness

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

type BrightnessParser struct {
	*antlr.BaseParser
}

var brightnessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func brightnessParserInit() {
	staticData := &brightnessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BRIGHTNESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "brightness",
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

// BrightnessParserInit initializes any static state used to implement BrightnessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBrightnessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BrightnessParserInit() {
	staticData := &brightnessParserStaticData
	staticData.once.Do(brightnessParserInit)
}

// NewBrightnessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBrightnessParser(input antlr.TokenStream) *BrightnessParser {
	BrightnessParserInit()
	this := new(BrightnessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &brightnessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Brightness.g4"

	return this
}

// BrightnessParser tokens.
const (
	BrightnessParserEOF        = antlr.TokenEOF
	BrightnessParserBRIGHTNESS = 1
	BrightnessParserOWNKEY     = 2
	BrightnessParserSPLITKEY   = 3
	BrightnessParserWS         = 4
)

// BrightnessParser rules.
const (
	BrightnessParserRULE_expression = 0
	BrightnessParserRULE_brightness = 1
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
	p.RuleIndex = BrightnessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrightnessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Brightness() IBrightnessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrightnessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrightnessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BrightnessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrightnessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrightnessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BrightnessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BrightnessParserRULE_expression)

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
		p.Brightness()
	}
	{
		p.SetState(5)
		p.Match(BrightnessParserEOF)
	}

	return localctx
}

// IBrightnessContext is an interface to support dynamic dispatch.
type IBrightnessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBrightnessContext differentiates from other interfaces.
	IsBrightnessContext()
}

type BrightnessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrightnessContext() *BrightnessContext {
	var p = new(BrightnessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BrightnessParserRULE_brightness
	return p
}

func (*BrightnessContext) IsBrightnessContext() {}

func NewBrightnessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrightnessContext {
	var p = new(BrightnessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BrightnessParserRULE_brightness

	return p
}

func (s *BrightnessContext) GetParser() antlr.Parser { return s.parser }

func (s *BrightnessContext) BRIGHTNESS() antlr.TerminalNode {
	return s.GetToken(BrightnessParserBRIGHTNESS, 0)
}

func (s *BrightnessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrightnessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrightnessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrightnessListener); ok {
		listenerT.EnterBrightness(s)
	}
}

func (s *BrightnessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BrightnessListener); ok {
		listenerT.ExitBrightness(s)
	}
}

func (p *BrightnessParser) Brightness() (localctx IBrightnessContext) {
	this := p
	_ = this

	localctx = NewBrightnessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BrightnessParserRULE_brightness)

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
		p.Match(BrightnessParserBRIGHTNESS)
	}

	return localctx
}
