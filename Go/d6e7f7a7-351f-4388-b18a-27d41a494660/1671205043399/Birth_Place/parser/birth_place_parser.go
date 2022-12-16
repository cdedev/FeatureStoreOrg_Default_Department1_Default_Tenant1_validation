// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671205043399/Birth_Place.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Birth_Place

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

type Birth_PlaceParser struct {
	*antlr.BaseParser
}

var birth_placeParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func birth_placeParserInit() {
	staticData := &birth_placeParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BIRTH_PLACE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "birth_place",
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

// Birth_PlaceParserInit initializes any static state used to implement Birth_PlaceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBirth_PlaceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Birth_PlaceParserInit() {
	staticData := &birth_placeParserStaticData
	staticData.once.Do(birth_placeParserInit)
}

// NewBirth_PlaceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBirth_PlaceParser(input antlr.TokenStream) *Birth_PlaceParser {
	Birth_PlaceParserInit()
	this := new(Birth_PlaceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &birth_placeParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Birth_Place.g4"

	return this
}

// Birth_PlaceParser tokens.
const (
	Birth_PlaceParserEOF         = antlr.TokenEOF
	Birth_PlaceParserBIRTH_PLACE = 1
	Birth_PlaceParserOWNKEY      = 2
	Birth_PlaceParserSPLITKEY    = 3
	Birth_PlaceParserWS          = 4
)

// Birth_PlaceParser rules.
const (
	Birth_PlaceParserRULE_expression  = 0
	Birth_PlaceParserRULE_birth_place = 1
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
	p.RuleIndex = Birth_PlaceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Birth_PlaceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Birth_place() IBirth_placeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBirth_placeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBirth_placeContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Birth_PlaceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Birth_PlaceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Birth_PlaceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Birth_PlaceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Birth_PlaceParserRULE_expression)

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
		p.Birth_place()
	}
	{
		p.SetState(5)
		p.Match(Birth_PlaceParserEOF)
	}

	return localctx
}

// IBirth_placeContext is an interface to support dynamic dispatch.
type IBirth_placeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBirth_placeContext differentiates from other interfaces.
	IsBirth_placeContext()
}

type Birth_placeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBirth_placeContext() *Birth_placeContext {
	var p = new(Birth_placeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Birth_PlaceParserRULE_birth_place
	return p
}

func (*Birth_placeContext) IsBirth_placeContext() {}

func NewBirth_placeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Birth_placeContext {
	var p = new(Birth_placeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Birth_PlaceParserRULE_birth_place

	return p
}

func (s *Birth_placeContext) GetParser() antlr.Parser { return s.parser }

func (s *Birth_placeContext) BIRTH_PLACE() antlr.TerminalNode {
	return s.GetToken(Birth_PlaceParserBIRTH_PLACE, 0)
}

func (s *Birth_placeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Birth_placeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Birth_placeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Birth_PlaceListener); ok {
		listenerT.EnterBirth_place(s)
	}
}

func (s *Birth_placeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Birth_PlaceListener); ok {
		listenerT.ExitBirth_place(s)
	}
}

func (p *Birth_PlaceParser) Birth_place() (localctx IBirth_placeContext) {
	this := p
	_ = this

	localctx = NewBirth_placeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Birth_PlaceParserRULE_birth_place)

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
		p.Match(Birth_PlaceParserBIRTH_PLACE)
	}

	return localctx
}
