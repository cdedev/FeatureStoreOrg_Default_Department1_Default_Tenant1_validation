// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674580642195/Religiousness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Religiousness

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

type ReligiousnessParser struct {
	*antlr.BaseParser
}

var religiousnessParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func religiousnessParserInit() {
	staticData := &religiousnessParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "RELIGIOUSNESS", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "religiousness",
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

// ReligiousnessParserInit initializes any static state used to implement ReligiousnessParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewReligiousnessParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ReligiousnessParserInit() {
	staticData := &religiousnessParserStaticData
	staticData.once.Do(religiousnessParserInit)
}

// NewReligiousnessParser produces a new parser instance for the optional input antlr.TokenStream.
func NewReligiousnessParser(input antlr.TokenStream) *ReligiousnessParser {
	ReligiousnessParserInit()
	this := new(ReligiousnessParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &religiousnessParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Religiousness.g4"

	return this
}

// ReligiousnessParser tokens.
const (
	ReligiousnessParserEOF           = antlr.TokenEOF
	ReligiousnessParserRELIGIOUSNESS = 1
	ReligiousnessParserOWNKEY        = 2
	ReligiousnessParserSPLITKEY      = 3
	ReligiousnessParserWS            = 4
)

// ReligiousnessParser rules.
const (
	ReligiousnessParserRULE_expression    = 0
	ReligiousnessParserRULE_religiousness = 1
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
	p.RuleIndex = ReligiousnessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReligiousnessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Religiousness() IReligiousnessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReligiousnessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReligiousnessContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ReligiousnessParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReligiousnessListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReligiousnessListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ReligiousnessParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ReligiousnessParserRULE_expression)

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
		p.Religiousness()
	}
	{
		p.SetState(5)
		p.Match(ReligiousnessParserEOF)
	}

	return localctx
}

// IReligiousnessContext is an interface to support dynamic dispatch.
type IReligiousnessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReligiousnessContext differentiates from other interfaces.
	IsReligiousnessContext()
}

type ReligiousnessContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReligiousnessContext() *ReligiousnessContext {
	var p = new(ReligiousnessContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ReligiousnessParserRULE_religiousness
	return p
}

func (*ReligiousnessContext) IsReligiousnessContext() {}

func NewReligiousnessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReligiousnessContext {
	var p = new(ReligiousnessContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ReligiousnessParserRULE_religiousness

	return p
}

func (s *ReligiousnessContext) GetParser() antlr.Parser { return s.parser }

func (s *ReligiousnessContext) RELIGIOUSNESS() antlr.TerminalNode {
	return s.GetToken(ReligiousnessParserRELIGIOUSNESS, 0)
}

func (s *ReligiousnessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReligiousnessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReligiousnessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReligiousnessListener); ok {
		listenerT.EnterReligiousness(s)
	}
}

func (s *ReligiousnessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ReligiousnessListener); ok {
		listenerT.ExitReligiousness(s)
	}
}

func (p *ReligiousnessParser) Religiousness() (localctx IReligiousnessContext) {
	this := p
	_ = this

	localctx = NewReligiousnessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ReligiousnessParserRULE_religiousness)

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
		p.Match(ReligiousnessParserRELIGIOUSNESS)
	}

	return localctx
}
