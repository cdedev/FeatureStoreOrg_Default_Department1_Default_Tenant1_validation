// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669693636336/Room.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Room

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

type RoomParser struct {
	*antlr.BaseParser
}

var roomParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func roomParserInit() {
	staticData := &roomParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "ROOM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "room",
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

// RoomParserInit initializes any static state used to implement RoomParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRoomParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RoomParserInit() {
	staticData := &roomParserStaticData
	staticData.once.Do(roomParserInit)
}

// NewRoomParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRoomParser(input antlr.TokenStream) *RoomParser {
	RoomParserInit()
	this := new(RoomParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &roomParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Room.g4"

	return this
}

// RoomParser tokens.
const (
	RoomParserEOF      = antlr.TokenEOF
	RoomParserROOM     = 1
	RoomParserOWNKEY   = 2
	RoomParserSPLITKEY = 3
	RoomParserWS       = 4
)

// RoomParser rules.
const (
	RoomParserRULE_expression = 0
	RoomParserRULE_room       = 1
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
	p.RuleIndex = RoomParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoomParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Room() IRoomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoomContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(RoomParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoomListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoomListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *RoomParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RoomParserRULE_expression)

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
		p.Room()
	}
	{
		p.SetState(5)
		p.Match(RoomParserEOF)
	}

	return localctx
}

// IRoomContext is an interface to support dynamic dispatch.
type IRoomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoomContext differentiates from other interfaces.
	IsRoomContext()
}

type RoomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoomContext() *RoomContext {
	var p = new(RoomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RoomParserRULE_room
	return p
}

func (*RoomContext) IsRoomContext() {}

func NewRoomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoomContext {
	var p = new(RoomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RoomParserRULE_room

	return p
}

func (s *RoomContext) GetParser() antlr.Parser { return s.parser }

func (s *RoomContext) ROOM() antlr.TerminalNode {
	return s.GetToken(RoomParserROOM, 0)
}

func (s *RoomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoomListener); ok {
		listenerT.EnterRoom(s)
	}
}

func (s *RoomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RoomListener); ok {
		listenerT.ExitRoom(s)
	}
}

func (p *RoomParser) Room() (localctx IRoomContext) {
	this := p
	_ = this

	localctx = NewRoomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RoomParserRULE_room)

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
		p.Match(RoomParserROOM)
	}

	return localctx
}
