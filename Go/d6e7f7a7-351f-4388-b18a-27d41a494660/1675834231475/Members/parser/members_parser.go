// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834231475/Members.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Members

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

type MembersParser struct {
	*antlr.BaseParser
}

var membersParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func membersParserInit() {
	staticData := &membersParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "MEMBERS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "members",
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

// MembersParserInit initializes any static state used to implement MembersParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMembersParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MembersParserInit() {
	staticData := &membersParserStaticData
	staticData.once.Do(membersParserInit)
}

// NewMembersParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMembersParser(input antlr.TokenStream) *MembersParser {
	MembersParserInit()
	this := new(MembersParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &membersParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Members.g4"

	return this
}

// MembersParser tokens.
const (
	MembersParserEOF      = antlr.TokenEOF
	MembersParserMEMBERS  = 1
	MembersParserOWNKEY   = 2
	MembersParserSPLITKEY = 3
	MembersParserWS       = 4
)

// MembersParser rules.
const (
	MembersParserRULE_expression = 0
	MembersParserRULE_members    = 1
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
	p.RuleIndex = MembersParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MembersParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Members() IMembersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMembersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMembersContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(MembersParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MembersListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MembersListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *MembersParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MembersParserRULE_expression)

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
		p.Members()
	}
	{
		p.SetState(5)
		p.Match(MembersParserEOF)
	}

	return localctx
}

// IMembersContext is an interface to support dynamic dispatch.
type IMembersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMembersContext differentiates from other interfaces.
	IsMembersContext()
}

type MembersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMembersContext() *MembersContext {
	var p = new(MembersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MembersParserRULE_members
	return p
}

func (*MembersContext) IsMembersContext() {}

func NewMembersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MembersContext {
	var p = new(MembersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MembersParserRULE_members

	return p
}

func (s *MembersContext) GetParser() antlr.Parser { return s.parser }

func (s *MembersContext) MEMBERS() antlr.TerminalNode {
	return s.GetToken(MembersParserMEMBERS, 0)
}

func (s *MembersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MembersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MembersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MembersListener); ok {
		listenerT.EnterMembers(s)
	}
}

func (s *MembersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MembersListener); ok {
		listenerT.ExitMembers(s)
	}
}

func (p *MembersParser) Members() (localctx IMembersContext) {
	this := p
	_ = this

	localctx = NewMembersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MembersParserRULE_members)

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
		p.Match(MembersParserMEMBERS)
	}

	return localctx
}
