// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638199157/Trans_Date.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Trans_Date

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

type Trans_DateParser struct {
	*antlr.BaseParser
}

var trans_dateParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func trans_dateParserInit() {
	staticData := &trans_dateParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "TRANS_DATE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "trans_date",
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

// Trans_DateParserInit initializes any static state used to implement Trans_DateParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTrans_DateParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Trans_DateParserInit() {
	staticData := &trans_dateParserStaticData
	staticData.once.Do(trans_dateParserInit)
}

// NewTrans_DateParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTrans_DateParser(input antlr.TokenStream) *Trans_DateParser {
	Trans_DateParserInit()
	this := new(Trans_DateParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &trans_dateParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Trans_Date.g4"

	return this
}

// Trans_DateParser tokens.
const (
	Trans_DateParserEOF        = antlr.TokenEOF
	Trans_DateParserTRANS_DATE = 1
	Trans_DateParserOWNKEY     = 2
	Trans_DateParserSPLITKEY   = 3
	Trans_DateParserWS         = 4
)

// Trans_DateParser rules.
const (
	Trans_DateParserRULE_expression = 0
	Trans_DateParserRULE_trans_date = 1
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
	p.RuleIndex = Trans_DateParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Trans_DateParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Trans_date() ITrans_dateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrans_dateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrans_dateContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(Trans_DateParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Trans_DateListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Trans_DateListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Trans_DateParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Trans_DateParserRULE_expression)

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
		p.Trans_date()
	}
	{
		p.SetState(5)
		p.Match(Trans_DateParserEOF)
	}

	return localctx
}

// ITrans_dateContext is an interface to support dynamic dispatch.
type ITrans_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrans_dateContext differentiates from other interfaces.
	IsTrans_dateContext()
}

type Trans_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrans_dateContext() *Trans_dateContext {
	var p = new(Trans_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Trans_DateParserRULE_trans_date
	return p
}

func (*Trans_dateContext) IsTrans_dateContext() {}

func NewTrans_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Trans_dateContext {
	var p = new(Trans_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Trans_DateParserRULE_trans_date

	return p
}

func (s *Trans_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Trans_dateContext) TRANS_DATE() antlr.TerminalNode {
	return s.GetToken(Trans_DateParserTRANS_DATE, 0)
}

func (s *Trans_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Trans_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Trans_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Trans_DateListener); ok {
		listenerT.EnterTrans_date(s)
	}
}

func (s *Trans_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Trans_DateListener); ok {
		listenerT.ExitTrans_date(s)
	}
}

func (p *Trans_DateParser) Trans_date() (localctx ITrans_dateContext) {
	this := p
	_ = this

	localctx = NewTrans_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Trans_DateParserRULE_trans_date)

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
		p.Match(Trans_DateParserTRANS_DATE)
	}

	return localctx
}
