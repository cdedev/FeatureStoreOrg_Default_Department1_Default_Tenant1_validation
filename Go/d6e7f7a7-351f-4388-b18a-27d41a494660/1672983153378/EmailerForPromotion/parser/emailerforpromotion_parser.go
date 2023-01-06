// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983153378/EmailerForPromotion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EmailerForPromotion

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

type EmailerForPromotionParser struct {
	*antlr.BaseParser
}

var emailerforpromotionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func emailerforpromotionParserInit() {
	staticData := &emailerforpromotionParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "EMAILERFORPROMOTION", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "emailerforpromotion",
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

// EmailerForPromotionParserInit initializes any static state used to implement EmailerForPromotionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewEmailerForPromotionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func EmailerForPromotionParserInit() {
	staticData := &emailerforpromotionParserStaticData
	staticData.once.Do(emailerforpromotionParserInit)
}

// NewEmailerForPromotionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewEmailerForPromotionParser(input antlr.TokenStream) *EmailerForPromotionParser {
	EmailerForPromotionParserInit()
	this := new(EmailerForPromotionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &emailerforpromotionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "EmailerForPromotion.g4"

	return this
}

// EmailerForPromotionParser tokens.
const (
	EmailerForPromotionParserEOF                 = antlr.TokenEOF
	EmailerForPromotionParserEMAILERFORPROMOTION = 1
	EmailerForPromotionParserOWNKEY              = 2
	EmailerForPromotionParserSPLITKEY            = 3
	EmailerForPromotionParserWS                  = 4
)

// EmailerForPromotionParser rules.
const (
	EmailerForPromotionParserRULE_expression          = 0
	EmailerForPromotionParserRULE_emailerforpromotion = 1
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
	p.RuleIndex = EmailerForPromotionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmailerForPromotionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Emailerforpromotion() IEmailerforpromotionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmailerforpromotionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmailerforpromotionContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(EmailerForPromotionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmailerForPromotionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmailerForPromotionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *EmailerForPromotionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, EmailerForPromotionParserRULE_expression)

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
		p.Emailerforpromotion()
	}
	{
		p.SetState(5)
		p.Match(EmailerForPromotionParserEOF)
	}

	return localctx
}

// IEmailerforpromotionContext is an interface to support dynamic dispatch.
type IEmailerforpromotionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmailerforpromotionContext differentiates from other interfaces.
	IsEmailerforpromotionContext()
}

type EmailerforpromotionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmailerforpromotionContext() *EmailerforpromotionContext {
	var p = new(EmailerforpromotionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = EmailerForPromotionParserRULE_emailerforpromotion
	return p
}

func (*EmailerforpromotionContext) IsEmailerforpromotionContext() {}

func NewEmailerforpromotionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmailerforpromotionContext {
	var p = new(EmailerforpromotionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = EmailerForPromotionParserRULE_emailerforpromotion

	return p
}

func (s *EmailerforpromotionContext) GetParser() antlr.Parser { return s.parser }

func (s *EmailerforpromotionContext) EMAILERFORPROMOTION() antlr.TerminalNode {
	return s.GetToken(EmailerForPromotionParserEMAILERFORPROMOTION, 0)
}

func (s *EmailerforpromotionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmailerforpromotionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmailerforpromotionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmailerForPromotionListener); ok {
		listenerT.EnterEmailerforpromotion(s)
	}
}

func (s *EmailerforpromotionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(EmailerForPromotionListener); ok {
		listenerT.ExitEmailerforpromotion(s)
	}
}

func (p *EmailerForPromotionParser) Emailerforpromotion() (localctx IEmailerforpromotionContext) {
	this := p
	_ = this

	localctx = NewEmailerforpromotionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, EmailerForPromotionParserRULE_emailerforpromotion)

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
		p.Match(EmailerForPromotionParserEMAILERFORPROMOTION)
	}

	return localctx
}
