// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672205472033/PoojaRoom.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PoojaRoom

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

type PoojaRoomParser struct {
	*antlr.BaseParser
}

var poojaroomParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func poojaroomParserInit() {
	staticData := &poojaroomParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "POOJAROOM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "poojaroom",
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

// PoojaRoomParserInit initializes any static state used to implement PoojaRoomParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPoojaRoomParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PoojaRoomParserInit() {
	staticData := &poojaroomParserStaticData
	staticData.once.Do(poojaroomParserInit)
}

// NewPoojaRoomParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPoojaRoomParser(input antlr.TokenStream) *PoojaRoomParser {
	PoojaRoomParserInit()
	this := new(PoojaRoomParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &poojaroomParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PoojaRoom.g4"

	return this
}

// PoojaRoomParser tokens.
const (
	PoojaRoomParserEOF       = antlr.TokenEOF
	PoojaRoomParserPOOJAROOM = 1
	PoojaRoomParserOWNKEY    = 2
	PoojaRoomParserSPLITKEY  = 3
	PoojaRoomParserWS        = 4
)

// PoojaRoomParser rules.
const (
	PoojaRoomParserRULE_expression = 0
	PoojaRoomParserRULE_poojaroom  = 1
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
	p.RuleIndex = PoojaRoomParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PoojaRoomParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Poojaroom() IPoojaroomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPoojaroomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPoojaroomContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PoojaRoomParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoojaRoomListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoojaRoomListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PoojaRoomParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PoojaRoomParserRULE_expression)

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
		p.Poojaroom()
	}
	{
		p.SetState(5)
		p.Match(PoojaRoomParserEOF)
	}

	return localctx
}

// IPoojaroomContext is an interface to support dynamic dispatch.
type IPoojaroomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPoojaroomContext differentiates from other interfaces.
	IsPoojaroomContext()
}

type PoojaroomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPoojaroomContext() *PoojaroomContext {
	var p = new(PoojaroomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PoojaRoomParserRULE_poojaroom
	return p
}

func (*PoojaroomContext) IsPoojaroomContext() {}

func NewPoojaroomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PoojaroomContext {
	var p = new(PoojaroomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PoojaRoomParserRULE_poojaroom

	return p
}

func (s *PoojaroomContext) GetParser() antlr.Parser { return s.parser }

func (s *PoojaroomContext) POOJAROOM() antlr.TerminalNode {
	return s.GetToken(PoojaRoomParserPOOJAROOM, 0)
}

func (s *PoojaroomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PoojaroomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PoojaroomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoojaRoomListener); ok {
		listenerT.EnterPoojaroom(s)
	}
}

func (s *PoojaroomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PoojaRoomListener); ok {
		listenerT.ExitPoojaroom(s)
	}
}

func (p *PoojaRoomParser) Poojaroom() (localctx IPoojaroomContext) {
	this := p
	_ = this

	localctx = NewPoojaroomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PoojaRoomParserRULE_poojaroom)

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
		p.Match(PoojaRoomParserPOOJAROOM)
	}

	return localctx
}
