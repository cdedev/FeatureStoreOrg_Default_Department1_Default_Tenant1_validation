// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771855533/SequenceLengthCount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SequenceLengthCount

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

type SequenceLengthCountParser struct {
	*antlr.BaseParser
}

var sequencelengthcountParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sequencelengthcountParserInit() {
	staticData := &sequencelengthcountParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "SEQUENCELENGTHCOUNT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "sequencelengthcount",
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

// SequenceLengthCountParserInit initializes any static state used to implement SequenceLengthCountParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSequenceLengthCountParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SequenceLengthCountParserInit() {
	staticData := &sequencelengthcountParserStaticData
	staticData.once.Do(sequencelengthcountParserInit)
}

// NewSequenceLengthCountParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSequenceLengthCountParser(input antlr.TokenStream) *SequenceLengthCountParser {
	SequenceLengthCountParserInit()
	this := new(SequenceLengthCountParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sequencelengthcountParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "SequenceLengthCount.g4"

	return this
}

// SequenceLengthCountParser tokens.
const (
	SequenceLengthCountParserEOF                 = antlr.TokenEOF
	SequenceLengthCountParserSEQUENCELENGTHCOUNT = 1
	SequenceLengthCountParserOWNKEY              = 2
	SequenceLengthCountParserSPLITKEY            = 3
	SequenceLengthCountParserWS                  = 4
)

// SequenceLengthCountParser rules.
const (
	SequenceLengthCountParserRULE_expression          = 0
	SequenceLengthCountParserRULE_sequencelengthcount = 1
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
	p.RuleIndex = SequenceLengthCountParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SequenceLengthCountParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Sequencelengthcount() ISequencelengthcountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISequencelengthcountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISequencelengthcountContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SequenceLengthCountParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SequenceLengthCountListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SequenceLengthCountListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SequenceLengthCountParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SequenceLengthCountParserRULE_expression)

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
		p.Sequencelengthcount()
	}
	{
		p.SetState(5)
		p.Match(SequenceLengthCountParserEOF)
	}

	return localctx
}

// ISequencelengthcountContext is an interface to support dynamic dispatch.
type ISequencelengthcountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequencelengthcountContext differentiates from other interfaces.
	IsSequencelengthcountContext()
}

type SequencelengthcountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequencelengthcountContext() *SequencelengthcountContext {
	var p = new(SequencelengthcountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SequenceLengthCountParserRULE_sequencelengthcount
	return p
}

func (*SequencelengthcountContext) IsSequencelengthcountContext() {}

func NewSequencelengthcountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequencelengthcountContext {
	var p = new(SequencelengthcountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SequenceLengthCountParserRULE_sequencelengthcount

	return p
}

func (s *SequencelengthcountContext) GetParser() antlr.Parser { return s.parser }

func (s *SequencelengthcountContext) SEQUENCELENGTHCOUNT() antlr.TerminalNode {
	return s.GetToken(SequenceLengthCountParserSEQUENCELENGTHCOUNT, 0)
}

func (s *SequencelengthcountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequencelengthcountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequencelengthcountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SequenceLengthCountListener); ok {
		listenerT.EnterSequencelengthcount(s)
	}
}

func (s *SequencelengthcountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SequenceLengthCountListener); ok {
		listenerT.ExitSequencelengthcount(s)
	}
}

func (p *SequenceLengthCountParser) Sequencelengthcount() (localctx ISequencelengthcountContext) {
	this := p
	_ = this

	localctx = NewSequencelengthcountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SequenceLengthCountParserRULE_sequencelengthcount)

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
		p.Match(SequenceLengthCountParserSEQUENCELENGTHCOUNT)
	}

	return localctx
}
