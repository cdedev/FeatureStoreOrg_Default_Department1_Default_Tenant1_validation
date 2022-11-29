// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707691291/PhoneNumber.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PhoneNumber

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

type PhoneNumberParser struct {
	*antlr.BaseParser
}

var phonenumberParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func phonenumberParserInit() {
	staticData := &phonenumberParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "PHONENUMBER", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "phonenumber",
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

// PhoneNumberParserInit initializes any static state used to implement PhoneNumberParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPhoneNumberParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PhoneNumberParserInit() {
	staticData := &phonenumberParserStaticData
	staticData.once.Do(phonenumberParserInit)
}

// NewPhoneNumberParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPhoneNumberParser(input antlr.TokenStream) *PhoneNumberParser {
	PhoneNumberParserInit()
	this := new(PhoneNumberParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &phonenumberParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "PhoneNumber.g4"

	return this
}

// PhoneNumberParser tokens.
const (
	PhoneNumberParserEOF         = antlr.TokenEOF
	PhoneNumberParserPHONENUMBER = 1
	PhoneNumberParserOWNKEY      = 2
	PhoneNumberParserSPLITKEY    = 3
	PhoneNumberParserWS          = 4
)

// PhoneNumberParser rules.
const (
	PhoneNumberParserRULE_expression  = 0
	PhoneNumberParserRULE_phonenumber = 1
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
	p.RuleIndex = PhoneNumberParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhoneNumberParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Phonenumber() IPhonenumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPhonenumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPhonenumberContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(PhoneNumberParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneNumberListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneNumberListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PhoneNumberParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PhoneNumberParserRULE_expression)

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
		p.Phonenumber()
	}
	{
		p.SetState(5)
		p.Match(PhoneNumberParserEOF)
	}

	return localctx
}

// IPhonenumberContext is an interface to support dynamic dispatch.
type IPhonenumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPhonenumberContext differentiates from other interfaces.
	IsPhonenumberContext()
}

type PhonenumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPhonenumberContext() *PhonenumberContext {
	var p = new(PhonenumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PhoneNumberParserRULE_phonenumber
	return p
}

func (*PhonenumberContext) IsPhonenumberContext() {}

func NewPhonenumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PhonenumberContext {
	var p = new(PhonenumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhoneNumberParserRULE_phonenumber

	return p
}

func (s *PhonenumberContext) GetParser() antlr.Parser { return s.parser }

func (s *PhonenumberContext) PHONENUMBER() antlr.TerminalNode {
	return s.GetToken(PhoneNumberParserPHONENUMBER, 0)
}

func (s *PhonenumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PhonenumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PhonenumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneNumberListener); ok {
		listenerT.EnterPhonenumber(s)
	}
}

func (s *PhonenumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhoneNumberListener); ok {
		listenerT.ExitPhonenumber(s)
	}
}

func (p *PhoneNumberParser) Phonenumber() (localctx IPhonenumberContext) {
	this := p
	_ = this

	localctx = NewPhonenumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PhoneNumberParserRULE_phonenumber)

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
		p.Match(PhoneNumberParserPHONENUMBER)
	}

	return localctx
}
