// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241758896/Bio.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Bio

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

type BioParser struct {
	*antlr.BaseParser
}

var bioParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func bioParserInit() {
	staticData := &bioParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "BIO", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "bio",
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

// BioParserInit initializes any static state used to implement BioParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBioParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BioParserInit() {
	staticData := &bioParserStaticData
	staticData.once.Do(bioParserInit)
}

// NewBioParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBioParser(input antlr.TokenStream) *BioParser {
	BioParserInit()
	this := new(BioParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &bioParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Bio.g4"

	return this
}

// BioParser tokens.
const (
	BioParserEOF      = antlr.TokenEOF
	BioParserBIO      = 1
	BioParserOWNKEY   = 2
	BioParserSPLITKEY = 3
	BioParserWS       = 4
)

// BioParser rules.
const (
	BioParserRULE_expression = 0
	BioParserRULE_bio        = 1
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
	p.RuleIndex = BioParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BioParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Bio() IBioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBioContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(BioParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BioListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BioListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *BioParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BioParserRULE_expression)

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
		p.Bio()
	}
	{
		p.SetState(5)
		p.Match(BioParserEOF)
	}

	return localctx
}

// IBioContext is an interface to support dynamic dispatch.
type IBioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBioContext differentiates from other interfaces.
	IsBioContext()
}

type BioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBioContext() *BioContext {
	var p = new(BioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BioParserRULE_bio
	return p
}

func (*BioContext) IsBioContext() {}

func NewBioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BioContext {
	var p = new(BioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BioParserRULE_bio

	return p
}

func (s *BioContext) GetParser() antlr.Parser { return s.parser }

func (s *BioContext) BIO() antlr.TerminalNode {
	return s.GetToken(BioParserBIO, 0)
}

func (s *BioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BioListener); ok {
		listenerT.EnterBio(s)
	}
}

func (s *BioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BioListener); ok {
		listenerT.ExitBio(s)
	}
}

func (p *BioParser) Bio() (localctx IBioContext) {
	this := p
	_ = this

	localctx = NewBioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BioParserRULE_bio)

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
		p.Match(BioParserBIO)
	}

	return localctx
}
