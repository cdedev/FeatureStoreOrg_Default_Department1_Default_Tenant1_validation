// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377534929/FrontCamera.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FrontCamera

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

type FrontCameraParser struct {
	*antlr.BaseParser
}

var frontcameraParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func frontcameraParserInit() {
	staticData := &frontcameraParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "FRONTCAMERA", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "frontcamera",
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

// FrontCameraParserInit initializes any static state used to implement FrontCameraParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFrontCameraParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FrontCameraParserInit() {
	staticData := &frontcameraParserStaticData
	staticData.once.Do(frontcameraParserInit)
}

// NewFrontCameraParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFrontCameraParser(input antlr.TokenStream) *FrontCameraParser {
	FrontCameraParserInit()
	this := new(FrontCameraParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &frontcameraParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "FrontCamera.g4"

	return this
}

// FrontCameraParser tokens.
const (
	FrontCameraParserEOF         = antlr.TokenEOF
	FrontCameraParserFRONTCAMERA = 1
	FrontCameraParserOWNKEY      = 2
	FrontCameraParserSPLITKEY    = 3
	FrontCameraParserWS          = 4
)

// FrontCameraParser rules.
const (
	FrontCameraParserRULE_expression  = 0
	FrontCameraParserRULE_frontcamera = 1
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
	p.RuleIndex = FrontCameraParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FrontCameraParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Frontcamera() IFrontcameraContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFrontcameraContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFrontcameraContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(FrontCameraParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrontCameraListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrontCameraListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *FrontCameraParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FrontCameraParserRULE_expression)

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
		p.Frontcamera()
	}
	{
		p.SetState(5)
		p.Match(FrontCameraParserEOF)
	}

	return localctx
}

// IFrontcameraContext is an interface to support dynamic dispatch.
type IFrontcameraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrontcameraContext differentiates from other interfaces.
	IsFrontcameraContext()
}

type FrontcameraContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrontcameraContext() *FrontcameraContext {
	var p = new(FrontcameraContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FrontCameraParserRULE_frontcamera
	return p
}

func (*FrontcameraContext) IsFrontcameraContext() {}

func NewFrontcameraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FrontcameraContext {
	var p = new(FrontcameraContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FrontCameraParserRULE_frontcamera

	return p
}

func (s *FrontcameraContext) GetParser() antlr.Parser { return s.parser }

func (s *FrontcameraContext) FRONTCAMERA() antlr.TerminalNode {
	return s.GetToken(FrontCameraParserFRONTCAMERA, 0)
}

func (s *FrontcameraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FrontcameraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FrontcameraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrontCameraListener); ok {
		listenerT.EnterFrontcamera(s)
	}
}

func (s *FrontcameraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FrontCameraListener); ok {
		listenerT.ExitFrontcamera(s)
	}
}

func (p *FrontCameraParser) Frontcamera() (localctx IFrontcameraContext) {
	this := p
	_ = this

	localctx = NewFrontcameraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FrontCameraParserRULE_frontcamera)

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
		p.Match(FrontCameraParserFRONTCAMERA)
	}

	return localctx
}
