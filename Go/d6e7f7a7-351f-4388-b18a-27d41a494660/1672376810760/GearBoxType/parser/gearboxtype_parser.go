// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376810760/GearBoxType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GearBoxType

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

type GearBoxTypeParser struct {
	*antlr.BaseParser
}

var gearboxtypeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gearboxtypeParserInit() {
	staticData := &gearboxtypeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "GEARBOXTYPE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "gearboxtype",
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

// GearBoxTypeParserInit initializes any static state used to implement GearBoxTypeParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGearBoxTypeParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GearBoxTypeParserInit() {
	staticData := &gearboxtypeParserStaticData
	staticData.once.Do(gearboxtypeParserInit)
}

// NewGearBoxTypeParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGearBoxTypeParser(input antlr.TokenStream) *GearBoxTypeParser {
	GearBoxTypeParserInit()
	this := new(GearBoxTypeParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &gearboxtypeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "GearBoxType.g4"

	return this
}

// GearBoxTypeParser tokens.
const (
	GearBoxTypeParserEOF         = antlr.TokenEOF
	GearBoxTypeParserGEARBOXTYPE = 1
	GearBoxTypeParserOWNKEY      = 2
	GearBoxTypeParserSPLITKEY    = 3
	GearBoxTypeParserWS          = 4
)

// GearBoxTypeParser rules.
const (
	GearBoxTypeParserRULE_expression  = 0
	GearBoxTypeParserRULE_gearboxtype = 1
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
	p.RuleIndex = GearBoxTypeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GearBoxTypeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Gearboxtype() IGearboxtypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGearboxtypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGearboxtypeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(GearBoxTypeParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearBoxTypeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearBoxTypeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *GearBoxTypeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GearBoxTypeParserRULE_expression)

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
		p.Gearboxtype()
	}
	{
		p.SetState(5)
		p.Match(GearBoxTypeParserEOF)
	}

	return localctx
}

// IGearboxtypeContext is an interface to support dynamic dispatch.
type IGearboxtypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGearboxtypeContext differentiates from other interfaces.
	IsGearboxtypeContext()
}

type GearboxtypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGearboxtypeContext() *GearboxtypeContext {
	var p = new(GearboxtypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GearBoxTypeParserRULE_gearboxtype
	return p
}

func (*GearboxtypeContext) IsGearboxtypeContext() {}

func NewGearboxtypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GearboxtypeContext {
	var p = new(GearboxtypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GearBoxTypeParserRULE_gearboxtype

	return p
}

func (s *GearboxtypeContext) GetParser() antlr.Parser { return s.parser }

func (s *GearboxtypeContext) GEARBOXTYPE() antlr.TerminalNode {
	return s.GetToken(GearBoxTypeParserGEARBOXTYPE, 0)
}

func (s *GearboxtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GearboxtypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GearboxtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearBoxTypeListener); ok {
		listenerT.EnterGearboxtype(s)
	}
}

func (s *GearboxtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GearBoxTypeListener); ok {
		listenerT.ExitGearboxtype(s)
	}
}

func (p *GearBoxTypeParser) Gearboxtype() (localctx IGearboxtypeContext) {
	this := p
	_ = this

	localctx = NewGearboxtypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GearBoxTypeParserRULE_gearboxtype)

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
		p.Match(GearBoxTypeParserGEARBOXTYPE)
	}

	return localctx
}
