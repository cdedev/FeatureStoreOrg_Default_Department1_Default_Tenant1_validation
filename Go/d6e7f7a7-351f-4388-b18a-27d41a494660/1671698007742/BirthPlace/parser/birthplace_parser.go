// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698007742/BirthPlace.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BirthPlace

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

type BirthPlaceParser struct {
	*antlr.BaseParser
}

var birthplaceParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func birthplaceParserInit() {
	staticData := &birthplaceParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BIRTHPLACE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "birthplace",
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

// BirthPlaceParserInit initializes any static state used to implement BirthPlaceParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBirthPlaceParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BirthPlaceParserInit() {
	staticData := &birthplaceParserStaticData
	staticData.once.Do(birthplaceParserInit)
}

// NewBirthPlaceParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBirthPlaceParser(input antlr.TokenStream) *BirthPlaceParser {
	BirthPlaceParserInit()
	this := new(BirthPlaceParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &birthplaceParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "BirthPlace.g4"

	return this
}

// BirthPlaceParser tokens.
const (
	BirthPlaceParserEOF        = antlr.TokenEOF
	BirthPlaceParserBIRTHPLACE = 1
	BirthPlaceParserOWNKEY     = 2
	BirthPlaceParserSPLITKEY   = 3
	BirthPlaceParserWS         = 4
)

// BirthPlaceParser rules.
const (
	BirthPlaceParserRULE_expression = 0
	BirthPlaceParserRULE_birthplace = 1
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
	p.RuleIndex = BirthPlaceParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirthPlaceParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Birthplace() IBirthplaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBirthplaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBirthplaceContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BirthPlaceParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirthPlaceListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirthPlaceListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BirthPlaceParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BirthPlaceParserRULE_expression)

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
		p.Birthplace()
	}
	{
		p.SetState(5)
		p.Match(BirthPlaceParserEOF)
	}

	return localctx
}

// IBirthplaceContext is an interface to support dynamic dispatch.
type IBirthplaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBirthplaceContext differentiates from other interfaces.
	IsBirthplaceContext()
}

type BirthplaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBirthplaceContext() *BirthplaceContext {
	var p = new(BirthplaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BirthPlaceParserRULE_birthplace
	return p
}

func (*BirthplaceContext) IsBirthplaceContext() {}

func NewBirthplaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BirthplaceContext {
	var p = new(BirthplaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BirthPlaceParserRULE_birthplace

	return p
}

func (s *BirthplaceContext) GetParser() antlr.Parser { return s.parser }

func (s *BirthplaceContext) BIRTHPLACE() antlr.TerminalNode {
	return s.GetToken(BirthPlaceParserBIRTHPLACE, 0)
}

func (s *BirthplaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BirthplaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BirthplaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirthPlaceListener); ok {
		listenerT.EnterBirthplace(s)
	}
}

func (s *BirthplaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BirthPlaceListener); ok {
		listenerT.ExitBirthplace(s)
	}
}

func (p *BirthPlaceParser) Birthplace() (localctx IBirthplaceContext) {
	this := p
	_ = this

	localctx = NewBirthplaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BirthPlaceParserRULE_birthplace)

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
		p.Match(BirthPlaceParserBIRTHPLACE)
	}

	return localctx
}
