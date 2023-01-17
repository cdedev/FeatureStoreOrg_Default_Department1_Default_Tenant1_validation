// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673929478843/Contact.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contact

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

type ContactParser struct {
	*antlr.BaseParser
}

var contactParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func contactParserInit() {
	staticData := &contactParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CONTACT", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "contact",
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

// ContactParserInit initializes any static state used to implement ContactParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewContactParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ContactParserInit() {
	staticData := &contactParserStaticData
	staticData.once.Do(contactParserInit)
}

// NewContactParser produces a new parser instance for the optional input antlr.TokenStream.
func NewContactParser(input antlr.TokenStream) *ContactParser {
	ContactParserInit()
	this := new(ContactParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &contactParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Contact.g4"

	return this
}

// ContactParser tokens.
const (
	ContactParserEOF      = antlr.TokenEOF
	ContactParserCONTACT  = 1
	ContactParserOWNKEY   = 2
	ContactParserSPLITKEY = 3
	ContactParserWS       = 4
)

// ContactParser rules.
const (
	ContactParserRULE_expression = 0
	ContactParserRULE_contact    = 1
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
	p.RuleIndex = ContactParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ContactParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Contact() IContactContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContactContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContactContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ContactParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContactListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContactListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ContactParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ContactParserRULE_expression)

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
		p.Contact()
	}
	{
		p.SetState(5)
		p.Match(ContactParserEOF)
	}

	return localctx
}

// IContactContext is an interface to support dynamic dispatch.
type IContactContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContactContext differentiates from other interfaces.
	IsContactContext()
}

type ContactContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContactContext() *ContactContext {
	var p = new(ContactContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ContactParserRULE_contact
	return p
}

func (*ContactContext) IsContactContext() {}

func NewContactContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContactContext {
	var p = new(ContactContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ContactParserRULE_contact

	return p
}

func (s *ContactContext) GetParser() antlr.Parser { return s.parser }

func (s *ContactContext) CONTACT() antlr.TerminalNode {
	return s.GetToken(ContactParserCONTACT, 0)
}

func (s *ContactContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContactContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContactContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContactListener); ok {
		listenerT.EnterContact(s)
	}
}

func (s *ContactContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ContactListener); ok {
		listenerT.ExitContact(s)
	}
}

func (p *ContactParser) Contact() (localctx IContactContext) {
	this := p
	_ = this

	localctx = NewContactContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ContactParserRULE_contact)

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
		p.Match(ContactParserCONTACT)
	}

	return localctx
}
