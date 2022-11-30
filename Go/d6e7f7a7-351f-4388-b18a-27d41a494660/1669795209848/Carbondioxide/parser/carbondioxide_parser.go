// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669795209848/Carbondioxide.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbondioxide

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

type CarbondioxideParser struct {
	*antlr.BaseParser
}

var carbondioxideParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func carbondioxideParserInit() {
	staticData := &carbondioxideParserStaticData
	staticData.literalNames = []string{
		"", "", "'$~$'", "' !$~$#%#&$&! '",
	}
	staticData.symbolicNames = []string{
		"", "CARBONDIOXIDE", "OWNKEY", "SPLITKEY", "WS",
	}
	staticData.ruleNames = []string{
		"expression", "carbondioxide",
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

// CarbondioxideParserInit initializes any static state used to implement CarbondioxideParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCarbondioxideParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CarbondioxideParserInit() {
	staticData := &carbondioxideParserStaticData
	staticData.once.Do(carbondioxideParserInit)
}

// NewCarbondioxideParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCarbondioxideParser(input antlr.TokenStream) *CarbondioxideParser {
	CarbondioxideParserInit()
	this := new(CarbondioxideParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &carbondioxideParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Carbondioxide.g4"

	return this
}

// CarbondioxideParser tokens.
const (
	CarbondioxideParserEOF           = antlr.TokenEOF
	CarbondioxideParserCARBONDIOXIDE = 1
	CarbondioxideParserOWNKEY        = 2
	CarbondioxideParserSPLITKEY      = 3
	CarbondioxideParserWS            = 4
)

// CarbondioxideParser rules.
const (
	CarbondioxideParserRULE_expression    = 0
	CarbondioxideParserRULE_carbondioxide = 1
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
	p.RuleIndex = CarbondioxideParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CarbondioxideParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Carbondioxide() ICarbondioxideContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICarbondioxideContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICarbondioxideContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(CarbondioxideParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbondioxideListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbondioxideListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CarbondioxideParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CarbondioxideParserRULE_expression)

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
		p.Carbondioxide()
	}
	{
		p.SetState(5)
		p.Match(CarbondioxideParserEOF)
	}

	return localctx
}

// ICarbondioxideContext is an interface to support dynamic dispatch.
type ICarbondioxideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCarbondioxideContext differentiates from other interfaces.
	IsCarbondioxideContext()
}

type CarbondioxideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCarbondioxideContext() *CarbondioxideContext {
	var p = new(CarbondioxideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CarbondioxideParserRULE_carbondioxide
	return p
}

func (*CarbondioxideContext) IsCarbondioxideContext() {}

func NewCarbondioxideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CarbondioxideContext {
	var p = new(CarbondioxideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CarbondioxideParserRULE_carbondioxide

	return p
}

func (s *CarbondioxideContext) GetParser() antlr.Parser { return s.parser }

func (s *CarbondioxideContext) CARBONDIOXIDE() antlr.TerminalNode {
	return s.GetToken(CarbondioxideParserCARBONDIOXIDE, 0)
}

func (s *CarbondioxideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CarbondioxideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CarbondioxideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbondioxideListener); ok {
		listenerT.EnterCarbondioxide(s)
	}
}

func (s *CarbondioxideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CarbondioxideListener); ok {
		listenerT.ExitCarbondioxide(s)
	}
}

func (p *CarbondioxideParser) Carbondioxide() (localctx ICarbondioxideContext) {
	this := p
	_ = this

	localctx = NewCarbondioxideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CarbondioxideParserRULE_carbondioxide)

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
		p.Match(CarbondioxideParserCARBONDIOXIDE)
	}

	return localctx
}
