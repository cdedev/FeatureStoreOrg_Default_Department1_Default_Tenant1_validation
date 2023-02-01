// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241389951/Victim.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Victim

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

type VictimParser struct {
	*antlr.BaseParser
}

var victimParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func victimParserInit() {
	staticData := &victimParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "VICTIM", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "victim",
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

// VictimParserInit initializes any static state used to implement VictimParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVictimParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func VictimParserInit() {
	staticData := &victimParserStaticData
	staticData.once.Do(victimParserInit)
}

// NewVictimParser produces a new parser instance for the optional input antlr.TokenStream.
func NewVictimParser(input antlr.TokenStream) *VictimParser {
	VictimParserInit()
	this := new(VictimParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &victimParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Victim.g4"

	return this
}

// VictimParser tokens.
const (
	VictimParserEOF      = antlr.TokenEOF
	VictimParserVICTIM   = 1
	VictimParserOWNKEY   = 2
	VictimParserSPLITKEY = 3
	VictimParserWS       = 4
)

// VictimParser rules.
const (
	VictimParserRULE_expression = 0
	VictimParserRULE_victim     = 1
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
	p.RuleIndex = VictimParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VictimParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Victim() IVictimContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVictimContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVictimContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(VictimParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VictimListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VictimListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *VictimParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VictimParserRULE_expression)

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
		p.Victim()
	}
	{
		p.SetState(5)
		p.Match(VictimParserEOF)
	}

	return localctx
}

// IVictimContext is an interface to support dynamic dispatch.
type IVictimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVictimContext differentiates from other interfaces.
	IsVictimContext()
}

type VictimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVictimContext() *VictimContext {
	var p = new(VictimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = VictimParserRULE_victim
	return p
}

func (*VictimContext) IsVictimContext() {}

func NewVictimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VictimContext {
	var p = new(VictimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = VictimParserRULE_victim

	return p
}

func (s *VictimContext) GetParser() antlr.Parser { return s.parser }

func (s *VictimContext) VICTIM() antlr.TerminalNode {
	return s.GetToken(VictimParserVICTIM, 0)
}

func (s *VictimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VictimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VictimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VictimListener); ok {
		listenerT.EnterVictim(s)
	}
}

func (s *VictimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VictimListener); ok {
		listenerT.ExitVictim(s)
	}
}

func (p *VictimParser) Victim() (localctx IVictimContext) {
	this := p
	_ = this

	localctx = NewVictimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, VictimParserRULE_victim)

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
		p.Match(VictimParserVICTIM)
	}

	return localctx
}
